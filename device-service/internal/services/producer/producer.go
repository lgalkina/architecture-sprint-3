package producer

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"device-service/internal/domains/entities"
	"github.com/IBM/sarama"
)

const (
	topic = "device-commands"
)

type producer struct {
	kafkaProducer sarama.SyncProducer
}

func NewCommandProducer() ICommandProducer {
	// Kafka configuration
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Consumer.Return.Errors = true
	config.Version = sarama.V2_6_0_0

	// Create a new producer
	p, err := sarama.NewSyncProducer(buildConnectionStr(), config)
	if err != nil {
		log.Fatalf("Error creating producer: %v", err)
	}
	return &producer{
		kafkaProducer: p,
	}
}

func (p *producer) SendCommand(command entities.DeviceCommand) error {
	commandJSON, err := json.Marshal(command)
	if err != nil {
		return err
	}
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(commandJSON),
	}
	partition, offset, err := p.kafkaProducer.SendMessage(msg)
	if err != nil {
		return err
	}
	fmt.Printf("Message sent to partition %d at offset %d\n", partition, offset)
	return nil
}

func (p *producer) Close() {
	if err := p.kafkaProducer.Close(); err != nil {
		log.Println("Error closing producer: %v", err)
	}
}

func buildConnectionStr() []string {
	return strings.Split(os.Getenv("KAFKA_SERVERS"), ",")
}
