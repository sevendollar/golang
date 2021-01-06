package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

const (
	LIMIT = 3
)

func main() {
	// setup redis connection
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	defer rdb.Close()

	// setup context alone with cancel
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// init the GAME
	initGame(ctx, rdb)

	// Setup the web
	r := gin.Default()

	// Load functions
	r.SetFuncMap(template.FuncMap{
		"increase": func(i int) int {
			i++
			return i
		},
	})

	// Load templates
	r.LoadHTMLGlob("templates/*.html")

	// Default error joute
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusServiceUnavailable, "unavailable.html", nil)
	})

	// Setup routes
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "root.html", nil)
	})

	r.GET("/game", func(c *gin.Context) {
		rdb := redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
			DB:   0,
		})
		defer rdb.Close()

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		isover, err := rdb.Get(ctx, "isOver").Int()
		if err != nil {
			log.Fatal("[REDIS][ERROR] Error getting \"isOver\" value,", err)
		}

		if isover == 1 {
			// show results when game is over
			names, err := rdb.ZRange(ctx, "game", 0, LIMIT-1).Result()
			if err != nil {
				log.Fatal("[REDIS][ERROR] Error getting \"game\" value,", err)
			}

			payload := gin.H{
				"names": names,
			}

			c.HTML(http.StatusOK, "game-report.html", payload)

			return
		}
		// go to the game page
		c.HTML(http.StatusOK, "game.html", nil)
	})

	// route game_check
	{
		r.GET("/game_check", func(c *gin.Context) {
			payload := "whoa, this is kinda cheating!"

			c.HTML(http.StatusOK, "template.html", payload)
		})

		r.POST("/game_check", func(c *gin.Context) {
			rdb := redis.NewClient(&redis.Options{
				Addr: "localhost:6379",
				DB:   0,
			})
			defer rdb.Close()

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			name := c.PostForm("name")

			if err := rdb.ZRank(ctx, "game", name).Err(); err != nil {
				if err == redis.Nil {
					// member doesn't exist

					// increase the gameCount key
					gamecount, err := rdb.Incr(ctx, "gameCount").Result()
					if err != nil {
						log.Println("[REDIS][ERROR] Error incresing the gameCount value,", err)
						return
					}

					// if gameCount hits the LIMIT, tag the isOver key
					if gamecount >= LIMIT {
						err := rdb.Set(ctx, "isOver", 1, 0).Err()
						if err != nil {
							log.Println("[REDIS][ERROR] Error setting the isOver value,", err)
							return
						}
					}

					// add the "gameCount" as rank to its corresponding name
					if err := rdb.ZAdd(ctx, "game", &redis.Z{
						Score:  float64(gamecount),
						Member: name,
					}).Err(); err != nil {
						log.Println("[REDIS][ERROR] Error adding the sorted-set member,", err)
						return
					}

					// show ranking
					ranking, err := rdb.ZRank(ctx, "game", name).Result()
					if err != nil {
						log.Println("[REDIS][ERROR] Error when getting the sorted-set value for ranking,", err)
						return
					}
					ranking++

					payload := gin.H{
						"name":    name,
						"ranking": ranking,
					}

					// render the HTML
					c.HTML(http.StatusOK, "game-ranking.html", payload)

					return
				}
				log.Println("[REDIS][ERROR] Error when getting the sorted-set value for ranking,", err)
				c.Redirect(http.StatusPermanentRedirect, "/unavailable")

				return
			}
			// member existed
			ranking, err := rdb.ZRank(ctx, "game", name).Result()
			if err != nil {
				log.Println("[REDIS][ERROR] Error when getting the sorted-set value for ranking,", err)
				return
			}
			ranking++

			payload := gin.H{
				"name":    name,
				"ranking": ranking,
			}

			c.HTML(http.StatusOK, "game-played.html", payload)
		})
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}

func initGame(ctx context.Context, rdb *redis.Client) {
	if isover, err := rdb.Get(ctx, "isOver").Int(); err != nil {
		if err == redis.Nil {
			// handel nil
			if err := rdb.MSet(ctx, "gameCount", 0, "isOver", 0).Err(); err != nil {
				log.Fatal("[REDIS][ERROR]", err)

			}
			if err := rdb.Del(ctx, "game").Err(); err != nil {
				log.Fatal("[REDIS][ERROR]", err)

			}
			log.Println("[REDIS][INFO] gameCount and isOver re-created!")
			return
		}
		log.Fatal("[REDIS][ERROR]", err)
	} else if isover == 1 {
		fmt.Println("The Game is over, do you wanna start over? (Yes or No)")

		startover := ""
		fmt.Scan(&startover)

		// fmt.Println("Answer:", startover)
		startover = strings.ToLower(startover)
		if startover == "yes" {
			if err := rdb.MSet(ctx, "gameCount", 0, "isOver", 0).Err(); err != nil {
				log.Fatal("[REDIS][ERROR]", err)

			}
			if err := rdb.Del(ctx, "game").Err(); err != nil {
				log.Fatal("[REDIS][ERROR]", err)

			}
			log.Println("[REDIS][INFO] gameCount and isOver re-created!")
			return

		} else if startover == "no" {
			fmt.Println("Do nothing")
			return

		} else {
			fmt.Println("Wrong answer, Bye!")
			os.Exit(0)

		}
	}
}
