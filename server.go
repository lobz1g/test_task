package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"test_task/database"
	"test_task/handler"
)

func init() {
	var DB_MIGRATE = flag.String("migrate", "", "flag for migration command")
	var PUSH_TOKEN = flag.String("token", "", "your application's API token")
	var PUSH_USER = flag.String("user", "", "the user/group key")
	flag.Parse()
	_ = os.Setenv("DB_MIGRATE", *DB_MIGRATE)
	_ = os.Setenv("PUSH_TOKEN", *PUSH_TOKEN)
	_ = os.Setenv("PUSH_USER", *PUSH_USER)

	fmt.Println(os.Getenv("PUSH_USER"))
	fmt.Println(os.Getenv("PUSH_TOKEN"))
	fmt.Println(os.Getenv("DB_MIGRATE"))
}

func main() {
	err := database.OpenConnection()
	if err != nil {
		log.Fatal(err)
	}

	gin.SetMode(gin.ReleaseMode)
	g := gin.Default()
	g.Use(gin.Recovery())

	h := handler.New()
	g.POST("/send", h.Send)
	g.POST("/count", h.GetCount)

	err = g.Run(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
