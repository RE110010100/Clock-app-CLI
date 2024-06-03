package clock

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestNewClock verifies the creation of a new Clock instance and checks its initial values.
func TestNewClock(t *testing.T) {
	c := NewClock("tick", "tock", "bong", 600)

	assert.NotNil(t, c)                // Check that the Clock instance is not nil
	assert.Equal(t, "tick", c.tick)    // Verify the tick value
	assert.Equal(t, "tock", c.tock)    // Verify the tock value
	assert.Equal(t, "bong", c.bong)    // Verify the bong value
	assert.Equal(t, 600, c.clockLimit) // Verify the clockLimit value
	assert.True(t, c.print)            // Check that printing is enabled
	assert.NotNil(t, c.Finished)       // Ensure the Finished channel is initialized
	assert.NotNil(t, c.stopTicker)     // Ensure the stopTicker channel is initialized
}

// TestClock_Run tests the Run method of the Clock, ensuring it operates correctly.
func TestClock_Run(t *testing.T) {

	c := NewClock("tick", "tock", "bong", 600)
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	go c.Run(cancel) // Start the clock in a goroutine

	time.Sleep(2 * time.Second) // Allow some ticks to be printed

	c.mutex.RLock()
	assert.Equal(t, "tick", c.tick) // Verify the tick value during operation
	assert.True(t, c.print)         // Ensure printing is enabled during operation
	c.mutex.RUnlock()

	c.StopTicker() // Stop the ticker
	cancel()
	time.Sleep(1 * time.Second) // Give some time to stop
}

// TestClock_TogglePrint verifies the TogglePrint method toggles the print state correctly.
func TestClock_TogglePrint(t *testing.T) {
	c := NewClock("tick", "tock", "bong", 600)

	// Initial state should be true
	c.mutex.RLock()
	assert.True(t, c.print)
	c.mutex.RUnlock()

	// Toggle to false
	c.TogglePrint()
	c.mutex.RLock()
	assert.False(t, c.print)
	c.mutex.RUnlock()

	// Toggle back to true
	c.TogglePrint()
	c.mutex.RLock()
	assert.True(t, c.print)
	c.mutex.RUnlock()
}

// TestClock_SetValues verifies that SetTick, SetTock, and SetBong methods update the clock's values correctly.
func TestClock_SetValues(t *testing.T) {
	c := NewClock("tick", "tock", "bong", 600)

	// Set new values
	c.SetTick("newtick")
	c.SetTock("newtock")
	c.SetBong("newbong")

	c.mutex.RLock()
	assert.Equal(t, "newtick", c.tick)
	assert.Equal(t, "newtock", c.tock)
	assert.Equal(t, "newbong", c.bong)
	c.mutex.RUnlock()
}

// TestClock_PrintFunctions verifies that the print functions operate correctly based on the print state.
func TestClock_PrintFunctions(t *testing.T) {
	c := NewClock("tick", "tock", "bong", 600)
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	go c.Run(cancel) // Start the clock in a goroutine

	time.Sleep(1 * time.Second) // Allow some ticks to be printed

	c.mutex.Lock()
	c.print = false // Disable printing
	c.mutex.Unlock()

	// With print disabled, nothing should be printed
	assert.NotPanics(t, func() { c.printTick() })
	assert.NotPanics(t, func() { c.printTock() })
	assert.NotPanics(t, func() { c.printBong() })

	c.mutex.Lock()
	c.print = true // Enable printing
	c.mutex.Unlock()

	// With print enabled, it should print without panicking
	assert.NotPanics(t, func() { c.printTick() })
	assert.NotPanics(t, func() { c.printTock() })
	assert.NotPanics(t, func() { c.printBong() })

	c.StopTicker()
	//cancel()
	time.Sleep(1 * time.Second) // Give some time to stop
}

// TestClock_StopTicker tests the closing of the stopTicker channel.
func TestClock_StopTicker(t *testing.T) {
	// Create a new Clock instance
	clk := NewClock("tick", "tock", "bong", 600)

	// Start the clock in a separate goroutine
	_, cancel := context.WithCancel(context.Background())
	go clk.Run(cancel)

	// Let the clock run for a few seconds
	time.Sleep(3 * time.Second)

	// Stop the ticker
	clk.StopTicker()

	// Give some time to ensure the ticker stops
	time.Sleep(1 * time.Second)

	// Check if the stopTicker channel is closed
	select {
	case _, ok := <-clk.stopTicker:
		if ok {
			t.Errorf("stopTicker channel is not closed")
		}
	default:
		t.Errorf("stopTicker channel is not closed")
	}
}

// TestClock_Finished tests the closing of the Finished channel.
func TestClock_Finished(t *testing.T) {
	// Create a new Clock instance
	clk := NewClock("tick", "tock", "bong", 10)

	// Start the clock in a separate goroutine
	_, cancel := context.WithCancel(context.Background())
	go clk.Run(cancel)

	// Let the clock run until it should stop itself
	select {
	case <-clk.Finished:
		// Check if the Finished channel is closed
		select {
		case _, ok := <-clk.Finished:
			if ok {
				t.Errorf("Finished channel is not closed")
			}
		default:
			t.Errorf("Finished channel is not closed")
		}
	case <-time.After(12 * time.Second):
		t.Errorf("Clock did not stop and close the Finished channel in time")
	}
}
