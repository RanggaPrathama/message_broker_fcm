package lib

import (
	"fmt"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

func ConnectionRabbitMQ() (*amqp091.Connection, error) {

	conPattern := "amqp://%s:%s@%s:%s/"
	if LoadEnv("RABBITMQ_USER") == "" {
		log.Panicf("RABBITMQ_USER is not set")
	}


	conStr := fmt.Sprintf(conPattern, LoadEnv("RABBITMQ_USER"), LoadEnv("RABBITMQ_PASSWORD"), LoadEnv("RABBITMQ_HOST"), LoadEnv("RABBITMQ_PORT"))


	log.Printf("Connecting to RabbitMQ: %s", conStr)

	conn, err := amqp091.Dial(conStr)

	if err != nil {
		log.Panicf("Failed to connect to RabbitMQ: %v", err)
	}

	// defer conn.Close()

	log.Print("Connected to RabbitMQ")

	return conn, err

}

func ChannelRabbitMQ(conn *amqp091.Connection) (*amqp091.Channel ,amqp091.Queue) {

	ch, err := conn.Channel()
	if err != nil {
		log.Panicf("Failed to open a channel: %v", err)
	}

	// for quueue load balanching
	// q , err := ch.QueueDeclare(
	// 	"notif_queue", // Nama queue
	// 	false,         // Durable
	// 	false,         // Delete otomatis
	// 	false,         // Eksklusif
	// 	false,         // No wait
	// 	nil,           // Argumen tambahan
	// )
	q, err := ch.QueueDeclare(
		"",    // Queue kosong (RabbitMQ akan membuat queue unik)
		false, // Tidak durable
		true,  // Auto-delete saat consumer disconnect
		true,  // Eksklusif untuk satu consumer
		false,
		nil,
	)



	if err != nil {
		log.Fatal(err)
	}

	

	//defer ch.Close()

	return  ch,q
}


func PublishRabbitMQ(ch *amqp091.Channel, q amqp091.Queue,exchange string ,body string) {

	// queue load balanching
	// err := ch.Publish(
	// 	"", // Exchange
	// 	q.Name,        // Routing key
	// 	false,
	// 	false,
	// 	amqp091.Publishing{
	// 		ContentType: "text/plain",
	// 		Body:        []byte(body),
	// 	},
	// )

	err := ch.Publish(
		exchange, // Nama exchange
		"",       // Routing key kosong (karena fanout tidak butuh routing key)
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)

	log.Printf(" [x] Sent %s", body)
	if err != nil {
		log.Printf("Failed to publish a message: %v", err)
	}
}


func ConsumeRabbitMQ(ch *amqp091.Channel, q amqp091.Queue, exchange string) {

	log.Printf("isi dari queue %s", q.Name)
	
	msgs, err := ch.Consume(
		q.Name, // Nama queue
		"",     // Consumer
		true,   // Auto Ack
		false,  // Eksklusif
		false,  // No Local
		false,  // No Wait
		nil,    // Argumen tambahan
	)

	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	// Binding queue ke exchange fanout
	err = ch.QueueBind(
		q.Name,  // Nama queue
		"",      // Routing key (kosong karena fanout tidak pakai routing key)
		exchange, // Nama exchange
		false,
		nil,
	)

	if err != nil {
		log.Fatalf("Failed to bind queue: %v", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}


func DeclareExchange(ch *amqp091.Channel, exchange string) {
	err := ch.ExchangeDeclare(
		exchange, // Nama exchange
		"fanout", // Tipe exchange
		true,     // Durable
		false,    // Auto-deleted
		false,    // Internal
		false,    // No Wait
		nil,      // Argumen tambahan
	)

	if err != nil {
		log.Fatalf("Failed to declare an exchange: %v", err)
	}
}
