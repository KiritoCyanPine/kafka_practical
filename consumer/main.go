package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/segmentio/kafka-go"
)

var (
	// messagesProcessed tracks the total number of Kafka messages successfully consumed.
	messagesProcessed = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "consumer_messages_total",
		Help: "Total number of messages consumed",
	})
	// processingLatency measures time spent per message from fetch to processing completion.
	processingLatency = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "consumer_processing_latency_seconds",
		Help:    "Processing time per message",
		Buckets: prometheus.DefBuckets,
	})
)

func init() {
	// Register the metrics with the global Prometheus registry once at startup.
	prometheus.MustRegister(messagesProcessed)
	prometheus.MustRegister(processingLatency)
}

func main() {
	// Launch an HTTP server exposing Prometheus metrics on /metrics for scraping.
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Println("metrics on :2112")
		_ = http.ListenAndServe(":2112", nil)
	}()

	// Configure a Kafka reader that joins consumer group consumer-1 on the metrics topic.
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "metrics",
		GroupID:  "consumer-1",
		MinBytes: 10e3,
		MaxBytes: 10e6,
	})
	defer r.Close()

	for {
		start := time.Now() // capture start time to measure read+process latency

		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			// On read failures, log and pause briefly before retrying.
			log.Printf("read error: %v", err)
			time.Sleep(time.Microsecond * 10)
			continue
		}

		log.Printf("message at offset %d: %s = %s", msg.Offset, string(msg.Key), string(msg.Value))

		// Update metrics for successful messages: increment total and record latency.
		messagesProcessed.Inc()
		processingLatency.Observe(time.Since(start).Seconds())
	}
}
