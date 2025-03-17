package main


import (
	"fmt"

	"github.com/RanggaPrathama/message_broker_fcm/lib"
)

func main(){
	conn, _ := lib.ConnectionRabbitMQ()
	channel, q := lib.ChannelRabbitMQ(conn)

	exchangeName := "notif_exchange"
	lib.DeclareExchange(channel, exchangeName)

	type Message struct {
		IDMessage int
		Body      string
	}

	message := Message{
		Body: "Hello World 1",
	}

	for i := 0; i<10; i++{
		lib.PublishRabbitMQ(channel, q, exchangeName, message.Body)
		message.Body = fmt.Sprintf("Hello World %d", i+2)
	}

	// lib.PublishRabbitMQ(channel, q, exchangeName, message.Body)
}