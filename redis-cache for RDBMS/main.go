// Redis Cache for slow OR expensive operations like RDBMS

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	DefaultDBValue      = "Hello Redis"
	DefaultRedisAddress = "localhost"
	DefaultRedisPort    = "6379"
)

var (
	ADDR = ""
	PORT = ""
)

func main() {
	// initialize the redis client

	ADDR = DefaultRedisAddress
	PORT = DefaultRedisPort
	if v := os.Getenv("ADDR"); v != "" {
		ADDR = v
	}
	if v := os.Getenv("PORT"); v != "" {
		ADDR = v
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: ADDR + ":" + PORT,
		DB:   0,
	})
	defer rdb.Close()

	// setup the context
	ctx := context.Background()

	// get value from Redis
	//     yes -> return the value
	//     no  -> get the lock
	//         yes ->
	//		       get value from RDBMS
	//             SET the value to Redis
	//             PUBLISH
	//             DEL the lock
	//             return the value
	//         no  ->
	//             get SUBSCRIBE -> get value from Redis -> return the value
	likesJef, likesJefErr := rdb.Get(ctx, "likes:jef").Result()
	if likesJefErr != nil {
		if likesJefErr == redis.Nil {
			fmt.Println("[Redis] Miss the cache!")

			// get the lock
			lock, err := rdb.SetNX(ctx, "likes/lock:jef", "true", 0).Result()
			if err != nil {
				log.Println("[ERRO] error getting the value:", err)
				return
			}
			// fmt.Println("got the lock?", lock)
			if lock == true {
				// GET value fro RDBMS
				time.Sleep(time.Second * 5)
				dbValue := DefaultDBValue

				// SET the value to Redis
				if _, err := rdb.Set(ctx, "likes:jef", dbValue, time.Second*30).Result(); err != nil {
					log.Println("[ERRO] error getting the value:", err)
					return
				}

				// PUBLISH
				_, err := rdb.Publish(ctx, "likes:jef", "OK").Result()
				if err != nil {
					log.Println("[ERRO] error getting the value:", err)
					return
				}

				// DEL the lock
				if _, err := rdb.Del(ctx, "likes/lock:jef").Result(); err != nil {
					log.Println("[ERRO] error getting the value:", err)
					return
				}
				// fmt.Println("Del the lock!")

				// return the value
				fmt.Println("likes:jef(DB)", dbValue)

			} else if lock == false {
				fmt.Println("[Redis] Wait for when the cache is ready.")
				// get subscribe channel
				for {
					pubsub := rdb.Subscribe(ctx, "likes:jef")
					msg, ok := <-pubsub.Channel()
					if !ok {
						log.Println("[ERRO] error getting the value:", err)
						return
					}
					defer pubsub.Close()

					if msg.Payload == "OK" {
						// get value
						r, err := rdb.Get(ctx, "likes:jef").Result()
						if err != nil {
							log.Println("[ERRO] error getting the value:", err)
							return
						}

						fmt.Println("likes:jef(CACHE:waited):", r)
						return
					}
				}

			}
			return
		}
		log.Println("[ERRO] error getting the value:", likesJefErr)
		return
	}

	fmt.Println("likes:jef(CACHE):", likesJef)
}
