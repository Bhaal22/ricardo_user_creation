package events

import (
    "log"
    "github.com/streadway/amqp"
)


type RMQ struct {
    ConnectionString string
}

func (r *RMQ) Publish(event *Event) error {
    conn, err := amqp.Dial(r.ConnectionString)

    if err != nil {
        log.Printf("Failed to connect to RabbitMQ: %s", err)
        return err
    }
    defer conn.Close()

    channel, err := conn.Channel()
    if err != nil {
        log.Printf("Failed to open a channel: %s", err)
        return err
    }

    queue, err := channel.QueueDeclare(
        "user_creation", // name
        true,   // durable
        false,   // delete when unused
        false,   // exclusive
        false,   // no-wait
        nil,     // arguments
      )

      if err != nil {
        log.Printf("Failed to open a channel: %s", err)
        return err
    }


    log.Printf("publishing on %s", queue.Name)
    err = channel.Publish(
        "",     // exchange
        queue.Name, // routing key
        false,  // mandatory
        false,  // immediate
        amqp.Publishing {
            DeliveryMode: amqp.Persistent,
            ContentType: "text/plain",
            Body:        []byte(event.ToJson()),
        })

    if err != nil {
        log.Printf("Failed to publish a message")
        return err
    }

    return nil
}

func failOnError(err error, msg string) {
    if err != nil {
        log.Printf("%s: %s", msg, err)
    }
}