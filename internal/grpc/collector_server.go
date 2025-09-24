package grpc

import (
	"context"

	"github.com/Observe86/lib-proto/gen"
)

type CollectorServer struct {
	gen.UnimplementedCollectorServiceServer
	Producer *KafkaProducer
}

func NewCollectorServer(producer *KafkaProducer) *CollectorServer {
	return &CollectorServer{Producer: producer}
}

func (s *CollectorServer) SendMetrics(ctx context.Context, batch *gen.MetricBatch) (*gen.Ack, error) {
	// serialize batch and send to Kafka
	data, _ := batch.Marshal()
	_ = s.Producer.SendMetrics(data)
	return &gen.Ack{Success: true}, nil
}

func (s *CollectorServer) SendLogs(ctx context.Context, batch *gen.LogBatch) (*gen.Ack, error) {
	data, _ := batch.Marshal()
	_ = s.Producer.SendLogs(data)
	return &gen.Ack{Success: true}, nil
}

func (s *CollectorServer) SendTraces(ctx context.Context, batch *gen.TraceBatch) (*gen.Ack, error) {
	data, _ := batch.Marshal()
	_ = s.Producer.SendTraces(data)
	return &gen.Ack{Success: true}, nil
}
