package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("*.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "ajax.html", nil)

	})

	r.GET("/ajax", func(c *gin.Context) {
		// CORS(Cross-Origin Resource Sharing)
		c.Header("Access-Control-Allow-Origin", "*")

		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "hello world",
		})
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Println("server is shutting down...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("server exiting...")

}
