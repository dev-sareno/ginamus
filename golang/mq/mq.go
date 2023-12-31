package mq

import (
	"fmt"
	"github.com/dev-sareno/ginamus/codec"
	"github.com/dev-sareno/ginamus/dto"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
)

func GetChannel() (*amqp.Channel, func(), bool) {
	// setup RabbitMQ
	conn, err := amqp.Dial(os.Getenv("RMQ_URL"))
	if err != nil {
		log.Printf("failed to connect to RabbitMQ. %s\n", err)
		return nil, func() {}, false
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Printf("failed to open a RabbitMQ channel. %s\n", err)
		return nil, func() {}, false
	}

	return ch, func() {
		log.Println("closing RabbitMQ channel")
		if err := ch.Close(); err != nil {
			fmt.Println(err.Error())
		}

		log.Println("closing RabbitMQ connection")
		if err := conn.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}, true
}

func GetLookupAQueue(ch *amqp.Channel) *amqp.Queue {
	return getQueue(ch, "lookup-a")
}

func GetLookupCnameQueue(ch *amqp.Channel) *amqp.Queue {
	return getQueue(ch, "lookup-cname")
}

func PublishToLookupA(ch *amqp.Channel, job *dto.Job) bool {
	q := GetLookupAQueue(ch)
	encodedJob := codec.Encode(job)
	return publish(ch, q.Name, encodedJob)
}

func PublishToLookupCname(ch *amqp.Channel, job *dto.Job) bool {
	q := GetLookupCnameQueue(ch)
	encodedJob := codec.Encode(job)
	return publish(ch, q.Name, encodedJob)
}
