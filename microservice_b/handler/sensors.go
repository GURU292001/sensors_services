package handler

import (
	"context"
	"log"
	"math/rand"
	"sync"
	"time"

	pb "sensors/sensorpb" // adjust the import path
)

var TemperatureDuration time.Duration
var MotionDuration time.Duration
var HumidityDuration time.Duration
var mu sync.RWMutex

type LetterGenerator struct {
	current int
}

func NewLetterGenerator() *LetterGenerator {
	return &LetterGenerator{current: 0}
}

func (lg *LetterGenerator) Next() string {
	letter := 'A' + rune(lg.current%26) // cycle from a–z
	lg.current++
	return string(letter)
}

func StartSensorDataGenerator(c pb.SensorServiceClient, sensorType string) {
	Id1 := NewLetterGenerator()

	for { // outer loop restarts ticker if duration changes
		interval := GetDuration(sensorType)
		ticker := time.NewTicker(interval)
		log.Printf("Started ---%v sensor---", sensorType)
		for t := range ticker.C {
			// If interval changed → restart with new ticker
			if GetDuration(sensorType) != interval {
				ticker.Stop()
				break
			}

			var value float32
			var sensorTypeLabel = sensorType

			switch sensorType {
			case "TEMPERATURE":
				value = 20 + rand.Float32()*10
			case "MOTION":
				value = float32(rand.Intn(2))
			case "HUMIDITY":
				value = 40 + rand.Float32()*40
			default:
				log.Printf("Unknown sensor type: %s", sensorType)
				continue
			}

			data := &pb.SensorData{
				Value:     value,
				Type:      sensorTypeLabel,
				Id1:       Id1.Next(),
				Id2:       rand.Int31n(100),
				Timestamp: t.Format(time.RFC3339),
			}

			res, err := c.SendSensorData(context.Background(), data)
			if err != nil {
				log.Printf("[%s] Error sending data: %v", sensorTypeLabel, err)
				continue
			}
			log.Printf("[%s] Ack: %s (Value: %.2f)", sensorTypeLabel, res.Status, value)
		}
	}
}

func SetDuration(sensor string, d time.Duration) {
	mu.Lock()
	defer mu.Unlock()
	switch sensor {
	case "TEMPERATURE":
		TemperatureDuration = d
	case "MOTION":
		MotionDuration = d
	case "HUMIDITY":
		HumidityDuration = d
	}
}

func GetDuration(sensor string) time.Duration {
	mu.RLock()
	defer mu.RUnlock()
	switch sensor {
	case "TEMPERATURE":
		return TemperatureDuration
	case "MOTION":
		return MotionDuration
	case "HUMIDITY":
		return HumidityDuration
	default:
		return 5 * time.Second
	}
}
