package broker

import (
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

const (
	noTimeout = -1
)

type Consumer struct {
	Consumer *kafka.Consumer
}

func NewConsumer(servers []string, groupId string, topic string) (*Consumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":        strings.Join(servers, ","),
		"group.id":                 groupId,
		"auto.offset.reset":        "earliest",
		"enable.auto.offset.store": false,
		"enable.auto.commit":       true,
		"auto.commit.interval.ms":  7000,
	})

	if err != nil {
		return nil, err
	}

	time.Sleep(5 * time.Second)
	err = c.SubscribeTopics([]string{topic, "^aRegex.*[Tt]opic"}, nil)

	if err != nil {
		return nil, err
	}

	return &Consumer{
		Consumer: c,
	}, nil
}

func (c *Consumer) Start(in chan string) {
	// Готовимся корректно завершить при SIGINT/SIGTERM
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	run := true

	for run {
		select {
		case sig := <-sigchan:
			log.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			msg, err := c.Consumer.ReadMessage(noTimeout)
			if msg == nil {
				log.Println("Massage is nil")
				continue
			}
			if err == nil {
				in <- string(msg.Value)
				if _, err = c.Consumer.StoreMessage(msg); err != nil {
					log.Println(err)
					break
				}
			} else if !err.(kafka.Error).IsTimeout() {
				log.Printf("Consumer error: %v (%v)\n", err, msg)
				break
			}
		}
	}
	c.Consumer.Close()
	log.Println("Consumer Close!!!!!")
}
