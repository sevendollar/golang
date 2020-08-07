package router

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sevendollar/lab/functions"
)

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func Add(c *gin.Context) {
	a, _ := strconv.Atoi(c.Param("a"))
	b, _ := strconv.Atoi(c.Param("b"))

	rlt := functions.Add(a, b)
	c.String(http.StatusOK, "%d", rlt)
}
