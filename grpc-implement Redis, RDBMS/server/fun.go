package main

import (
	"context"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

// Get data from Database with Redis As Cache
//
// check the logic below:
//
// get the value from redis(GET):
//     - yes: return the value
//     - no: get the LOCK from redis(SETNX):
//         - yes:
//             - get value from DB
//             - set the value to redis(SET)
//             - publish to redis(PUB)
//             - delete the LOCK from redis(DEL)
//             - return the value
//         - no:
//             - subscript to redis(SUB)
//             - return the value
func getDBWithRedis(ctx context.Context, key string, rCh chan string, errCh chan error) {
	c := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	defer c.Close()
	defer close(rCh)
	defer close(errCh)

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
				errCh <- err
				rCh <- ""
				log.Println("SetNX: ", err)
				return
			}

			if lock == true { // get the LOCK from redis(SETNX): yes
				// get value from DB
				dbValue := getDB()

				// set the value to redis(SET)
				setRlt, err := c.Set(ctx, key, dbValue, time.Second*30).Result()
				if err != nil {
					errCh <- err
					rCh <- ""
					log.Println("Set: ", err)
					return
				}
				if setRlt != "OK" {
					// TODO: setting to redis has failed
				}
				// publish to redis(PUB)
				pubRlt, err := c.Publish(ctx, key, dbValue).Result()
				if err != nil {
					errCh <- err
					rCh <- ""
					log.Println("Publish: ", err)
					return
				}
				if pubRlt != 1 {
					// TODO: publish to redis has failed
				}

				// delete the LOCK from redis(DEL)
				delRlt, err := c.Del(ctx, keyLock).Result()
				if err != nil {
					errCh <- err
					rCh <- ""
					log.Println("Del: ", err)
					return
				}
				if delRlt != 1 {
					// TODO: deleting from redis has failed
				}

				// return the value
				errCh <- nil
				rCh <- dbValue + "(DB)"

			} else { // get the LOCK from redis(SETNX): no

				// subscript to redis(SUB)
				sub := c.Subscribe(ctx, key)
				defer sub.Close()

				rlt, ok := <-sub.Channel()
				if !ok {
					errCh <- err
					rCh <- ""
					log.Println("subscribing has failed")
					return
				}

				// return the value
				errCh <- nil
				rCh <- rlt.Payload + "(SUB)"

			}
			return
		}
		errCh <- err
		rCh <- ""
		log.Println("Get: ", err)
		return
	}

	// get the value from redis(GET): yes
	// return the value
	errCh <- nil
	rCh <- rlt + "(CACHE)"

}

// fake db process
func getDB() string {
	time.Sleep(time.Second * 5)

	l := []string{
		"Redis",
		"PostgreSQL",
		"gRPC",
		"World",
	}

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(l))

	return "Hello " + l[n]
}
