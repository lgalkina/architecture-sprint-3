package consumer

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/IBM/sarama"
	"telemetry-service/internal/domains/entities"
	"telemetry-service/internal/domains/repository"
)

const (
	group = "telemetry_consumer_group"
	topic = "telemetry"
)

type consumer struct {
	consumerGroup sarama.ConsumerGroup
	telemetryRepo repository.ITelemetryRepository
}

func NewConsumer() IConsumerService {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Version = sarama.V2_6_0_0
	config.Consumer.Offsets.Initial = sarama.OffsetNewest

	cg, err := sarama.NewConsumerGroup(buildConnectionStr(), group, config)
	if err != nil {
		log.Fatalf("Error creating consumer group: %v", err)
	}

	c := &consumer{
		consumerGroup: cg,
		telemetryRepo: repository.NewTelemetryRepository(),
	}
	go c.consume()

	return c
}

func (c *consumer) Close() {
	if err := c.consumerGroup.Close(); err != nil {
		log.Printf("Error closing consumer: %v\n", err)
	}
}

func (c *consumer) consume() {
	ctx := context.Background()
	sConsumer := saramaConsumer{
		ready:    make(chan bool),
		messages: make(chan *sarama.ConsumerMessage),
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if err := c.consumerGroup.Consume(ctx, []string{topic}, &sConsumer); err != nil {
				log.Fatalf("Error consuming messages: %v", err)
			}
			if ctx.Err() != nil {
				return
			}
			sConsumer.ready = make(chan bool)
		}
	}()

	<-sConsumer.ready
	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-sConsumer.messages:
			c.processMessage(msg)
		}
	}
}

func (c *consumer) processMessage(msg *sarama.ConsumerMessage) {
	telemetryData := &entities.TelemetryData{}
	err := json.Unmarshal(msg.Value, telemetryData)
	if err != nil {
		log.Printf("Error unmarshalling message: %v", err)
		return
	}

	if err := c.telemetryRepo.SaveDeviceTelemetry(telemetryData); err != nil {
		log.Printf("Error saving telemetry: %v", err)
		return
	}
}

func buildConnectionStr() []string {
	return strings.Split(os.Getenv("KAFKA_SERVERS"), ",")
}

// Consumer represents a Sarama consumer group consumer
type saramaConsumer struct {
	ready    chan bool
	messages chan *sarama.ConsumerMessage
}

func (c *saramaConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		log.Printf("Message claimed: value = %s, timestamp = %v, topic = %s", string(message.Value), message.Timestamp, message.Topic)
		c.messages <- message
		session.MarkMessage(message, "")
	}
	return nil
}

func (c *saramaConsumer) Setup(sarama.ConsumerGroupSession) error {
	close(c.ready)
	return nil
}

func (c *saramaConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}
