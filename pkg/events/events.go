package events

import "time"

type Direction int

const (
	Up = iota
	Down
)

func (dir Direction) String() string {
	return []string{"up", "down"}[dir]
}

type PriceUpdateEvent struct {
	coin      string
	price     float64
	timestamp time.Time
	payload   interface{}
}

type ThresholdCrossedEvent struct {
	coin      string
	price     float64
	direction Direction
	timestamp time.Time
	payload   interface{}
}

func (e PriceUpdateEvent) GetName() string {
	return "PriceUpdateEvent"
}

func (e PriceUpdateEvent) GetDateTime() time.Time {
	return e.timestamp
}

func (e PriceUpdateEvent) GetPayload() interface{} {
	return e.payload
}

func NewPriceUpdateEvent(coin string, price float64, timestamp time.Time, payload interface{}) *PriceUpdateEvent {
	return &PriceUpdateEvent{
		coin:      coin,
		price:     price,
		timestamp: timestamp,
		payload:   payload,
	}
}

func (e ThresholdCrossedEvent) GetName() string {
	return "ThresholdCrossedEvent"
}

func (e ThresholdCrossedEvent) GetDateTime() time.Time {
	return e.timestamp
}

func (e ThresholdCrossedEvent) GetPayload() interface{} {
	return e.payload
}

func NewThresholdCrossedEvent(coin string, price float64, direction Direction, timestamp time.Time, payload interface{}) *ThresholdCrossedEvent {
	return &ThresholdCrossedEvent{
		coin:      coin,
		price:     price,
		direction: direction,
		timestamp: timestamp,
		payload:   payload,
	}
}
