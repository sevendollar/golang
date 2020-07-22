package main

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	pb "proto/webCalculator"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	r := gin.Default()

	r.GET("/add/:a/:b", func(c *gin.Context) {

		a, err := strconv.Atoi(c.Param("a"))
		if err != nil {
			c.String(http.StatusInternalServerError, "")
		}

		b, err := strconv.Atoi(c.Param("b"))
		if err != nil {
			c.String(http.StatusInternalServerError, "")
		}

		req := &pb.Request{
			A: int32(a),
			B: int32(b),
		}

		conn, err := grpc.Dial(":2343", grpc.WithInsecure())
		if err != nil {
			log.Printf("[ERROR] %v", err)
		}
		client := pb.NewCalculatorClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(1000*time.Millisecond))
		defer cancel()

		rlt, err := client.Add(ctx, req)
		if err != nil {
			log.Printf("[ERROR] %v", err)
		}

		c.JSON(http.StatusOK, gin.H{"result": rlt.GetResult()})
	})

	r.Run()
}
