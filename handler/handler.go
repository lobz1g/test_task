package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"test_task/database"
	"test_task/push"
)

type (
	handler struct {
	}
)

func New() *handler {
	return &handler{}
}

func (handler) Send(ctx *gin.Context) {
	p := push.New()
	err := ctx.ShouldBindJSON(p)
	if err != nil {
		log.Println(err)
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, _ := p.Push()

	s := database.NewStatus(resp.Status)
	err = s.Set()
	if err != nil {
		log.Println(err)
		ctx.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, "OK")
}

func (handler) GetCount(ctx *gin.Context) {
	date, err := getParamDateInTime(ctx)
	if err != nil {
		log.Println(err)
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	d := database.New()
	result, err := d.GetCount(date)
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusOK, result)
	return
}

func getParamDateInTime(ctx *gin.Context) (time.Time, error) {
	dateByte, err := ctx.GetRawData()
	if err != nil {
		return time.Time{}, err
	}
	dateTime, err := time.ParseInLocation("2006-01-02 15:04:05", string(dateByte), time.UTC)
	return dateTime, err
}
