FROM golang:alpine as builder
RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

RUN adduser -D -g '' appuser
RUN mkdir /app-folder
COPY . /app-folder
WORKDIR /app-folder

RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /app

FROM scratch

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /app /app

USER appuser

ENTRYPOINT ["/app"]