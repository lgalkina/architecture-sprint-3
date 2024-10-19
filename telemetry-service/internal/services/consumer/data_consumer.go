package consumer

import "github.com/Shopify/sarama"

type dataConsumer struct {
	ready    chan bool
	messages chan *sarama.ConsumerMessage
}

func (consumer *dataConsumer) Setup(sarama.ConsumerGroupSession) error {
	close(consumer.ready)
	return nil
}

func (consumer *dataConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *dataConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		consumer.messages <- message
		session.MarkMessage(message, "")
	}
	return nil
}
