package producer

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

const (
	deviceIDPrefix       = "dev-"
	deviceIDOffset       = 100_000
	baseTemperature      = 60.0
	temperatureVariance  = 10.0
	defaultParallelCount = 1
	runtimeDuration      = time.Minute
)

type Metric struct {
	MessageID   string  `json:"message_id"`
	DeviceID    string  `json:"device_id"`
	Temperature float64 `json:"temperature"`
	Timestamp   int64   `json:"timestamp"`
}

type Device struct {
	ID string
}

type Orchestrator struct {
	writer        *kafka.Writer
	deviceCount   int
	devices       []Device
	devicesLocker sync.Mutex
}

func NewOrchestrator(parallelism int, writer *kafka.Writer) *Orchestrator {
	if parallelism <= 0 {
		parallelism = defaultParallelCount
	}

	orc := &Orchestrator{
		writer:      writer,
		deviceCount: parallelism,
	}
	orc.bootstrapDevices()

	return orc
}

func (o *Orchestrator) Run(ctx context.Context) {
	var wg sync.WaitGroup
	wg.Add(len(o.devices))

	// ctxTimeout, cancel := context.WithTimeout(ctx, runtimeDuration)
	// defer cancel()

	for _, device := range o.devices {
		go func(d Device) {
			defer wg.Done()
			o.produce(ctx, d)
		}(device)
	}

	wg.Wait()
	log.Println("all producers stopped")
}

func (o *Orchestrator) bootstrapDevices() {
	for i := 0; i < o.deviceCount; i++ {
		device := Device{ID: deviceIDPrefix + strconv.Itoa(i+deviceIDOffset)}
		o.addDevice(device)
	}
}

func (o *Orchestrator) addDevice(device Device) {
	o.devicesLocker.Lock()
	o.devices = append(o.devices, device)
	o.devicesLocker.Unlock()
}

func (o *Orchestrator) produce(ctx context.Context, d Device) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		metric := generateMetric(d.ID)
		payload, err := json.Marshal(metric)
		if err != nil {
			log.Printf("failed to marshal metric for device %s: %v", d.ID, err)
			continue
		}

		err = o.writer.WriteMessages(ctx, kafka.Message{
			Key:   []byte(metric.DeviceID),
			Value: payload,
		})
		if err != nil {
			log.Printf("write error: %v", err)
		} else {
			log.Printf("produced metric for device %s: %s", d.ID, string(metric.MessageID))
		}

		time.Sleep(time.Microsecond * 100) // simulate some delay between messages
	}
}

func generateMetric(deviceID string) Metric {
	return Metric{
		MessageID:   uuid.New().String(),
		DeviceID:    deviceID,
		Temperature: baseTemperature + rand.Float64()*temperatureVariance,
		Timestamp:   time.Now().Unix(),
	}
}
