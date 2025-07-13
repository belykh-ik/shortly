package broker

import (
	"api/shorturl/broker/models"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
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
	Routrer  *http.ServeMux
}

func NewConsumer(servers []string, groupId string, topic string, router *http.ServeMux) (*Consumer, error) {
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
		Routrer:  router,
	}, nil
}

func (c *Consumer) Start() {
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
				go c.handleMassage(msg.Value)
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

func (c *Consumer) handleMassage(msg []byte) {
	msgUrl := &models.MassageConsumer{
		Url: string(msg),
	}
	msgByte, err := json.Marshal(msgUrl)
	if err != nil {
		log.Println(err)
	}
	// Create *http.Request
	req := httptest.NewRequest("POST", "/create", bytes.NewReader(msgByte))
	req.Header.Set("Content-Type", "application/json")
	// Add Authorization
	req.Header.Set("Authorization", os.Getenv("TOKEN"))
	// Create *httptest.ResponseRecorder (ResponseWriter interface)
	rr := httptest.NewRecorder()

	c.Routrer.ServeHTTP(rr, req)
	time.Sleep(3 * time.Second)
	fmt.Println(rr.Body.String())
}
