package example

import (
	"runtime"

	"github.com/re-cinq/go-bus"
)

func main() {

	// Init the bus
	myBus := bus.NewEventBus(12, runtime.NumCPU())
	myBus.Start()

	// Init the handler
	myHandler := MyEventHandler{}

	// Subscribe it and from this moment on all events published on the
	// EmissionsCalculatedTopic will reach this handler
	// You can have multiple handlers subscribed to the same topic as well
	myBus.Subscribe(EmissionsCalculatedTopic, &myHandler)

	// Publish an event
	myBus.Publish(MetericsCollectedEvent{
		Id:     "my_unique_id",
		Cpu:    "4",
		Memory: "64",
	})

	myBus.Stop()
}
