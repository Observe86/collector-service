package service

import (
	"context"
	"os"

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

	ctx := context.Background()
	return writer.WriteMessages(ctx, kafka.Message{Value: value})
}

func (k *KafkaProducer) SendMetrics(data []byte) error {
	topic := os.Getenv("KAFKA_METRICS_TOPIC")
	if topic == "" {
		topic = "metrics"
	}
	return k.SendMessage(topic, data)
}
func (k *KafkaProducer) SendLogs(data []byte) error {
	topic := os.Getenv("KAFKA_LOGS_TOPIC")
	if topic == "" {
		topic = "logs"
	}
	return k.SendMessage(topic, data)
}
func (k *KafkaProducer) SendTraces(data []byte) error {
	topic := os.Getenv("KAFKA_TRACES_TOPIC")
	if topic == "" {
		topic = "traces"
	}
	return k.SendMessage(topic, data)
}
