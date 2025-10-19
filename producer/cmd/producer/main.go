package main

import (
	"context"
	"log"
	"time"

	producer "khafka-producer/internal/producer"

	"github.com/segmentio/kafka-go"
)

func main() {
	khafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers:          []string{"localhost:9092"},
		Topic:            "metrics",
		Balancer:         &kafka.LeastBytes{},
		BatchSize:        1000,
		BatchTimeout:     10 * time.Millisecond,
		CompressionCodec: kafka.Snappy.Codec(),
	})
	defer func() {
		if err := khafkaWriter.Close(); err != nil {
			log.Printf("failed to close writer: %v", err)
		}
	}()

	orchestrator := producer.NewOrchestrator(100, khafkaWriter)

	orchestrator.Run(context.Background())
}
