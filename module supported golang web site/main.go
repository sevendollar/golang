package main

import (
	"github.com/gin-gonic/gin"

	"github.com/sevendollar/lab/router"
)

func main() {
	r := gin.Default()

	r.GET("/ping", router.Ping)
	r.GET("/add/:a/:b", router.Add)

	r.Run()
}
