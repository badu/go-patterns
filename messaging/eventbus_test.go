package messaging_test

import (
	"testing"
	"time"

	"github.com/pearai/PaymentsDraft/pkg/payments"
	"github.com/stretchr/testify/assert"
)

const (
	eventDeployment = payments.EventID("deployment")
	eventSmokeTest  = payments.EventID("smoke_test")
)

type deploymentEvent struct {
	duration time.Duration
}

func (e *deploymentEvent) EventID() payments.EventID {
	return eventDeployment
}

type smokeTestEvent struct {
	duration time.Duration
}

func (e *smokeTestEvent) EventID() payments.EventID {
	return eventSmokeTest
}

func TestBusSubscribePublish(t *testing.T) {
	bus := payments.NewEventBus()
	hadEvent := false
	duration := 100 * time.Second

	bus.Subscribe(eventDeployment, func(e payments.Event) {
		assert.Equal(t, e.EventID(), eventDeployment)
		se := e.(*deploymentEvent)
		assert.Equal(t, se.duration, duration)
		hadEvent = true
	})
	bus.Subscribe(eventSmokeTest, func(e payments.Event) {
		t.Fatalf("should never be called")
	})
	assert.Equal(t, hadEvent, false)

	bus.Publish(&deploymentEvent{
		duration: duration,
	})
	assert.Equal(t, hadEvent, true)
}

func TestBusPublishIncompatibleEvent(t *testing.T) {
	bus := payments.NewEventBus()
	duration := 100 * time.Second
	bus.Subscribe(eventSmokeTest, func(e payments.Event) {
		t.Fatalf("should never be called")
	})

	bus.Publish(&deploymentEvent{
		duration: duration,
	})
}

func TestBusSubscribeUnsubscribe(t *testing.T) {
	bus := payments.NewEventBus()
	hadEvent := false
	duration := 42 * time.Millisecond

	id := bus.Subscribe(eventSmokeTest, func(e payments.Event) {
		assert.Equal(t, e.EventID(), eventSmokeTest)
		se := e.(*smokeTestEvent)
		assert.Equal(t, se.duration, duration)
		hadEvent = true
	})
	bus.Publish(&smokeTestEvent{
		duration: duration,
	})
	assert.Equal(t, hadEvent, true)

	hadEvent = false
	bus.Unsubscribe(id)
	bus.Publish(&smokeTestEvent{
		duration: duration,
	})
	assert.Equal(t, hadEvent, false)
}

func TestBusSubscribeMultiple(t *testing.T) {
	smokeCount := 0
	smokeTestDuration := 16 * time.Second
	onSmokeEvent := func(e payments.Event) {
		assert.Equal(t, e.EventID(), eventSmokeTest)
		smokeCount++
		se := e.(*smokeTestEvent)
		assert.Equal(t, se.duration, smokeTestDuration)
	}

	deploymentCount := 0
	deploymentDuration := 77 * time.Millisecond
	onDeploymentEvent := func(e payments.Event) {
		assert.Equal(t, e.EventID(), eventDeployment)
		deploymentCount++
		se := e.(*deploymentEvent)
		assert.Equal(t, se.duration, deploymentDuration)
	}

	bus := payments.NewEventBus()
	publishSmoke := func() {
		bus.Publish(&smokeTestEvent{
			duration: smokeTestDuration,
		})
	}
	publishDeployment := func() {
		bus.Publish(&deploymentEvent{
			duration: deploymentDuration,
		})
	}

	id1 := bus.Subscribe(eventSmokeTest, onSmokeEvent)
	id2 := bus.Subscribe(eventDeployment, onDeploymentEvent)
	id3 := bus.Subscribe(eventSmokeTest, onSmokeEvent)

	publishSmoke()
	assert.Equal(t, smokeCount, 2)
	assert.Equal(t, deploymentCount, 0)

	publishDeployment()
	assert.Equal(t, smokeCount, 2)
	assert.Equal(t, deploymentCount, 1)

	bus.Unsubscribe(id1)
	bus.Unsubscribe(id2)
	publishSmoke()
	publishDeployment()
	assert.Equal(t, smokeCount, 3)
	assert.Equal(t, deploymentCount, 1)

	bus.Unsubscribe(id3)
	publishSmoke()
	publishDeployment()
	assert.Equal(t, smokeCount, 3)
	assert.Equal(t, deploymentCount, 1)
}

func TestBusPublishRecursive(t *testing.T) {
	smokeCount := 0

	bus := payments.NewEventBus()
	smokeFn := func(duration time.Duration) {
		bus.Publish(&smokeTestEvent{
			duration: duration,
		})
	}

	onSmokeEvent := func(e payments.Event) {
		assert.Equal(t, e.EventID(), eventSmokeTest)
		smokeCount++
		se := e.(*smokeTestEvent)

		if se.duration < 16*time.Second {
			smokeFn(2 * se.duration)
		}
	}

	bus.Subscribe(eventSmokeTest, onSmokeEvent)
	smokeFn(1 * time.Second)

	assert.Equal(t, smokeCount, 5)
}
