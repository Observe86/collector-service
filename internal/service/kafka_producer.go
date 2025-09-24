package grpc

import (
	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	Brokers string
}

func NewKafkaProducer(brokers string) *KafkaProducer {
	return &KafkaProducer{Brokers: brokers}
}

func (k *KafkaProducer) SendMessage(topic string, value []byte) error {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{k.Brokers},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
	defer writer.Close()

	return writer.WriteMessages(nil, kafka.Message{Value: value})
}

func (k *KafkaProducer) SendMetrics(data []byte) error {
	return k.SendMessage("metrics", data)
}
func (k *KafkaProducer) SendLogs(data []byte) error {
	return k.SendMessage("logs", data)
}
func (k *KafkaProducer) SendTraces(data []byte) error {
	return k.SendMessage("traces", data)
}
