package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Shopify/sarama"
	"telemetry-service/internal/domains/entities"
	"telemetry-service/internal/domains/repository"
)

const (
	group = "telemetry-consumer-group"
	topic = "telemetry-topic"
)

type consumer struct {
	consumerGroup sarama.ConsumerGroup
	telemetryRepo repository.ITelemetryRepository
	ready         chan bool
	messages      chan *sarama.ConsumerMessage
}

func NewConsumer() IConsumerService {
	// Set up Kafka consumer configuration
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Version = sarama.V2_6_0_0

	// Create a new consumer group
	cg, err := sarama.NewConsumerGroup(buildConnectionStr(), group, config)
	if err != nil {
		log.Fatalf("Error creating consumer group: %v", err)
	}

	c := &consumer{
		consumerGroup: cg,
		telemetryRepo: repository.NewTelemetryRepository(),
		ready:         make(chan bool),
		messages:      make(chan *sarama.ConsumerMessage),
	}
	go c.consume()

	return c
}

func (c *consumer) Close() {
	if err := c.consumerGroup.Close(); err != nil {
		log.Println("Error closing consumer: %v", err)
	}
}

func (c *consumer) consume() {
	<-c.ready
	ctx := context.Background()
	// Consume messages from the Kafka topic
	for {
		if err := c.consumerGroup.Consume(ctx, []string{topic}, c); err != nil {
			log.Fatalf("Error consuming messages: %v", err)
		}

		// Process messages
		select {
		case <-ctx.Done():
			return
		case msg := <-c.messages:
			c.processMessage(msg)
		}
	}
}

// processMessage processes a single Kafka message
func (c *consumer) processMessage(msg *sarama.ConsumerMessage) {
	// Deserialize the message into a TelemetryData struct
	var telemetryData *entities.TelemetryData
	err := json.Unmarshal(msg.Value, telemetryData)
	if err != nil {
		log.Printf("Error unmarshalling message: %v", err)
		return
	}

	// Process the telemetry data
	fmt.Printf("Received telemetry data: %+v\n", telemetryData)

	if err := c.telemetryRepo.SaveDeviceTelemetry(telemetryData); err != nil {
		log.Printf("Error saving telemetry: %v", err)
		return
	}
}

func (c *consumer) Setup(sarama.ConsumerGroupSession) error {
	close(c.ready)
	return nil
}

func (c *consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (c *consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		c.messages <- message
		session.MarkMessage(message, "")
	}
	return nil
}

func buildConnectionStr() []string {
	return strings.Split(os.Getenv("KAFKA_SERVERS"), ",")
}
