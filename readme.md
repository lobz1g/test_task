# It is just test task

## How to use

```bash
git clone https://github.com/lobz1g/test_task .

docker build -t "push" .

docker run -d -p 3000:3000 --env=DB_DSN="host=localhost user=postgres password=mysecretpassword dbname=postgres port=1234 sslmode=disable TimeZone=Etc/UTC" --name push-app push -migrate=false -token=key -user=key
```

### Args

1. `migrate` - the flag for migration schema. If this one is `true` migration will start after connection to database. Else schema won't be updated.
2. `token` - API Token/Key for pushover.net
3. `user` - your user key

### Env

1. `DB_DSN` - credentials for connecting to databse
   * `host` - address to database server
   * `user` - username
   * `password` - password
   * `dbname` - database name
   * `port` - port of database server
   * `sslmode` - SSL encryption
   * `TimeZone` - timezone of database (default is Etc/UTC)