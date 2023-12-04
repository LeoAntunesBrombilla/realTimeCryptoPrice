package events

import (
	"fmt"
	"sync"
)

type PriceUpdateEventHandler struct {
}

func NewPriceUpdateEventHandler() *PriceUpdateEventHandler {
	return &PriceUpdateEventHandler{}
}

func (h *PriceUpdateEventHandler) Handle(event EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()

	priceUpdateEvent, ok := event.(*PriceUpdateEvent)
	if !ok {
		return
	}

	fmt.Println("New price for %s: %f \n", priceUpdateEvent.coin, priceUpdateEvent.price)
}
