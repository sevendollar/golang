package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/sevendollar/go-zabbix"
)

func t(w http.ResponseWriter, r *http.Request) {
	c := r.Context()

	rlt := make(chan []byte)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func(c context.Context) {
		go func() {
			user := "jeffl"
			passwd := ""
			uri := "http://192.168.15.230/zabbix/api_jsonrpc.php"

			s, err := zabbix.NewSession(user, passwd, uri)
			if err != nil {
				log.Printf("[ERROR]: %v\n", err)
				w.WriteHeader(404)
				return
			}

			resp, err := s.Do(zabbix.NewRequest(
				"item.gt",
				map[string]interface{}{
					"output": "extend",
				},
			))
			if err != nil {
				log.Printf("[ERROR]: %v\n", err)
				w.WriteHeader(404)
				return
			}

			zabbix.JsonPretty(&resp)

			rlt <- resp

		}()

		for {
			select {
			case <-c.Done():
				log.Println("[INFO]: zabbix call has been canceled")
				return
			}
		}
	}(ctx)

	select {
	case <-c.Done():
		cancel()
		log.Printf("[USER INTERUPTION]: %v\n", c.Err())
		w.WriteHeader(404)
	case rlt := <-rlt:
		w.Header().Set("Content-Type", "application/json")
		w.Write(rlt)
	case <-time.After(time.Duration(15 * time.Second)):
		w.Write([]byte("timeout!"))
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", t)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("[ERROR]: %v\n", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	srv.Shutdown(ctx)

	time.Sleep(500 * time.Millisecond)
	log.Print("[WARNNING]: service shutdown!")
	os.Exit(0)
}
