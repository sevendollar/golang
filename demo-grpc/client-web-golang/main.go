package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	greetpb "github.com/sevendollar/demo-web-golang/proto"
	"google.golang.org/grpc"
)

var (
	conn       *grpc.ClientConn
	target     *string
	targetPort *string
	otps       []grpc.DialOption
)

func init() {
	target = flag.String("host", "localhost", "host to listen to")
	targetPort = flag.String("port", "50051", "port to listen to")
	flag.Parse()

	otps = append(otps, grpc.WithInsecure())
}

func main() {
	r := gin.Default()

	// for docker
	// r.LoadHTMLGlob("templates/*")

	// for kubernetes
	r.LoadHTMLFiles("templates/root.html")

	r.GET("/", func(c *gin.Context) {
		hostname, err := os.Hostname()
		if err != nil {
			log.Printf("Failed getting the hostname, %v\n", err)
			c.HTML(http.StatusOK, "root.html", gin.H{
				"hostname": "Not Found!",
			})
			return
		}

		c.HTML(http.StatusOK, "root.html", gin.H{
			"hostname": hostname,
		})
	})

	r.GET("/hello/:firstname/:lastname", func(c *gin.Context) {
		firstName := c.Param("firstname")
		lastName := c.Param("lastname")

		conn, err := grpc.Dial((*target)+":"+(*targetPort), otps...)
		if err != nil {
			log.Printf("Cannot connecte to server, %v\n", err)
			c.HTML(http.StatusInternalServerError, "root.html", gin.H{
				"hostname": "Internal Server Error, " + fmt.Sprintf("%v", err),
			})
			return
		}
		defer conn.Close()

		pbc := greetpb.NewGreetServiceClient(conn)

		req := &greetpb.HelloRequest{
			PersonalInfo: &greetpb.PersonalInfo{
				FirstName: firstName,
				LastName:  lastName,
			},
		}

		resp, err := pbc.Hello(context.Background(), req)
		if err != nil {
			log.Printf("Failed calling Helle() function, %v\n", err)
			c.HTML(http.StatusInternalServerError, "root.html", gin.H{
				"hostname": "Internal Server Error, " + fmt.Sprintf("%v", err),
			})
			return
		}

		log.Printf("Client Called, %v\n", resp)
		c.HTML(http.StatusOK, "root.html", gin.H{
			"hostname": resp.GetMessage(),
		})
	})

	r.Run(":8080")
}
