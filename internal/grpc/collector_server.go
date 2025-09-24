package grpc

import (
	"context"

	"github.com/Observe86/collector-service/internal/service"
	"github.com/Observe86/lib-proto/gen"
	"google.golang.org/protobuf/proto"
)

type CollectorServer struct {
	gen.UnimplementedCollectorServiceServer
	Producer *service.KafkaProducer
}

func NewCollectorServer(producer *service.KafkaProducer) *CollectorServer {
	return &CollectorServer{Producer: producer}
}

func (s *CollectorServer) SendMetrics(ctx context.Context, batch *gen.MetricBatch) (*gen.Ack, error) {
	data, _ := proto.Marshal(batch)
	_ = s.Producer.SendMetrics(data)
	return &gen.Ack{Success: true}, nil
}

func (s *CollectorServer) SendLogs(ctx context.Context, batch *gen.LogBatch) (*gen.Ack, error) {
	data, _ := proto.Marshal(batch)
	_ = s.Producer.SendLogs(data)
	return &gen.Ack{Success: true}, nil
}

func (s *CollectorServer) SendTraces(ctx context.Context, batch *gen.TraceBatch) (*gen.Ack, error) {
	data, _ := proto.Marshal(batch)
	_ = s.Producer.SendTraces(data)
	return &gen.Ack{Success: true}, nil
}
