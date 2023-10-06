package main

import (
	"fmt"
	"github.com/dev-sareno/ginamus/gin/handler"
	"github.com/dev-sareno/ginamus/gin/mq"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	"net/http"
	"os"
)

func main() {
	// setup RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	mq.FailOnError(err, "Failed to connect to RabbitMQ")
	defer func() {
		if err := conn.Close(); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}()

	ch, err := conn.Channel()
	mq.FailOnError(err, "Failed to open a channel")
	defer func() {
		if err := ch.Close(); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}()

	// init Gin web server
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/", func(c *gin.Context) {
		handler.Post(c, ch)
	})

	r.GET("/", func(c *gin.Context) {
		handler.Get(c, ch)
	})

	if err := r.Run(":8000"); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
