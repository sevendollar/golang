package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	hellopb "github.com/sevendollar/pb_test/proto"
	"google.golang.org/grpc"
)

func m() {
	// fmt.Println(get())
}

func get() (rlt string) {
	c := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	defer c.Close()

	ctx := context.Background()

	key := "likes:jef"

	keyList := strings.Split(key, ":")
	topic := keyList[0]
	username := keyList[1]
	keyLock := topic + "/lock:" + username

	// get the value from redis(GET):
	rlt, err := c.Get(ctx, key).Result()
	if err != nil {
		// get the value from redis(GET): no
		if err == redis.Nil {
			// get the LOCK from redis(SETNX):
			lock, err := c.SetNX(ctx, keyLock, true, 0).Result()
			if err != nil {
				log.Fatal("SetNX: ", err)
			}

			if lock == true { // get the LOCK from redis(SETNX): yes
				// get value from DB
				time.Sleep(time.Second * 5)
				dbValue := "hello world"

				// set the value to redis(SET)
				setRlt, err := c.Set(ctx, key, dbValue, time.Second*30).Result()
				if err != nil {
					log.Fatal("Set: ", err)
				}
				if setRlt != "OK" {
					// TODO: setting to redis has failed
				}
				// publish to redis(PUB)
				pubRlt, err := c.Publish(ctx, key, dbValue).Result()
				if err != nil {
					log.Fatal("Publish: ", err)
				}
				if pubRlt != 1 {
					// TODO: publish to redis has failed
				}

				// delete the LOCK from redis(DEL)
				delRlt, err := c.Del(ctx, keyLock).Result()
				if err != nil {
					log.Fatal("Del: ", err)
				}
				if delRlt != 1 {
					// TODO: deleting from redis has failed
				}

				// return the value
				return dbValue + "(DB)"
				// fmt.Printf("result: %s(from DB)\n", dbValue)

			} else { // get the LOCK from redis(SETNX): no

				// get the subscription from redis(SUB)
				sub := c.Subscribe(ctx, key)
				rlt, ok := <-sub.Channel()
				if !ok {
					log.Fatal("subscribing has failed")
				}

				// return the value
				return rlt.Payload + "(SUB)"
				// fmt.Printf("result: %s(from SUBSCRIBE)\n", rlt.Payload)

			}
			return
		}
		log.Fatal("Get: ", err)
	}

	// get the value from redis(GET): yes
	// return the value
	return rlt + "(CACHE)"
	// fmt.Printf("result: %s(from CACHE)\n", rlt)

}

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	c := hellopb.NewHelloServiceClient(conn)

	r, err := c.RedisGet(context.Background(), &hellopb.RedisGetRequest{
		Key: "likes:jef",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(r.GetValue())

}
