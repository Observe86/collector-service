package main

import (
	"log"
	"net"
	"os"

	"github.com/Observe86/collector-service/internal/grpc"
	"github.com/Observe86/collector-service/internal/service"
	"github.com/Observe86/lib-proto/gen"
	"github.com/joho/godotenv"
	grpcPkg "google.golang.org/grpc"
)

func main() {
	godotenv.Load()

	kafkaBrokers := os.Getenv("KAFKA_BROKERS")
	if kafkaBrokers == "" {
		kafkaBrokers = "kafka:9092"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	producer := service.NewKafkaProducer(kafkaBrokers)
	server := grpc.NewCollectorServer(producer)

	grpcServer := grpcPkg.NewServer()
	gen.RegisterCollectorServiceServer(grpcServer, server)

	log.Printf("Collector-Gateway listening on %s, sending to Kafka %s", port, kafkaBrokers)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
