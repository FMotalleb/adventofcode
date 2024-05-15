package main

import (
	"log"
	"net"
	"os"
	"strconv"
	"time"

	dotenv "github.com/fmotalleb/adventofcode/dot_env"
	"github.com/gin-gonic/gin"
)

func getListenAddress() (addr net.TCPAddr, err error) {
	port, e := strconv.Atoi(os.Getenv("PORT"))
	log.Println(os.Getenv("HOST"), port)
	switch e {
	case nil:
		addr = net.TCPAddr{
			IP:   net.ParseIP(os.Getenv("HOST")),
			Port: port,
			Zone: "",
		}
	default:
		err = e
	}
	return
}
func main() {
	dotenv.Load()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, Gin!",
		})
	})

	runServer(router)

	for {
		time.Sleep(time.Second * 5)
	}
}

func runServer(engine *gin.Engine) {
	for {
		addr, err := getListenAddress()
		if err != nil {
			panic(err)
		}
		listener, err := net.ListenTCP("tcp", &addr)
		if err != nil {
			panic(err)
		}
		go engine.RunListener(listener)
		<-dotenv.Watch()
		listener.Close()
	}

}
