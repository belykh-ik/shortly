package broker

import (
	"api/shorturl/broker/handleMessage"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

const (
	noTimeout = -1
)

type Consumer struct {
	Consumer *kafka.Consumer
	h        *handleMessage.HandleMessageDeps
}

func NewConsumer(servers []string, groupId string, topic string, h *handleMessage.HandleMessageDeps) (*Consumer, error) {
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
		h:        h,
	}, nil
}

func (c *Consumer) Start() {
	// Готовимся корректно завершить при SIGINT/SIGTERM
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	wg := sync.WaitGroup{}
	check := false

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
				wg.Add(1)
				go func() {
					defer wg.Done()
					rr := c.h.HandleMassage(msg.Value)
					Produce(rr.Body.String())
					check = true
				}()
				wg.Wait()
				if check {
					if _, err = c.Consumer.StoreMessage(msg); err != nil {
						log.Println(err)
						break
					}
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
