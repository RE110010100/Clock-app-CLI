package clock

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Clock represents a configurable clock with tick, tock, and bong sounds.
type Clock struct {
	tick        string        // String for each second
	tock        string        // String for each minute
	bong        string        // String for each hour
	mutex       sync.RWMutex  // Mutex to protect shared resources
	print       bool          // Flag to enable/disable printing
	timeElapsed int           // Elapsed time in seconds
	clockLimit  int           // Limit for the clock in seconds
	Finished    chan struct{} // Channel to signal when the clock is finished
	stopTicker  chan int      // Channel to signal when to stop the ticker
}

// NewClock creates a new Clock instance with the given tick, tock, and bong values and a clock limit.
func NewClock(tick, tock, bong string, clockLimit int) *Clock {

	return &Clock{
		tick:        tick,
		tock:        tock,
		bong:        bong,
		print:       true,
		clockLimit:  clockLimit,
		timeElapsed: 0,
		Finished:    make(chan struct{}),
		stopTicker:  make(chan int),
	}
}

// Run starts the clock and prints tick, tock, and bong at appropriate intervals.
// It stops automatically when the clock limit is reached or when Stop is called.
func (c *Clock) Run(cancel context.CancelFunc) {

	// Create a ticker that ticks every second
	tickerSecond := time.NewTicker(1 * time.Second)

	//Go routine to handle the clock logic
	go func() {
		c.timeElapsed = 0
		for {
			select {

			case <-tickerSecond.C: //Handle Ticker ticks, and print tick, tock, and bong values at appropriate intervals
				c.timeElapsed++
				if c.timeElapsed%3600 == 0 {
					c.printBong()
				} else if c.timeElapsed%60 == 0 {
					c.printTock()
				} else {
					c.printTick()
				}
				if c.timeElapsed >= c.clockLimit {
					// Stop the clock when the limit is reached
					c.Stop()
					return
				}

			case <-c.stopTicker:
				//Stop the ticker
				tickerSecond.Stop()
				return

			}

		}
	}()
}

// printTick prints the tick value along with the elapsed time in seconds.
func (c *Clock) printTick() {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	if c.print {
		fmt.Println(c.tick, "[", c.timeElapsed, "seconds have elapsed ]")
	}
}

// printTock prints the tock value along with the elapsed time in minutes.
func (c *Clock) printTock() {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	if c.print {
		fmt.Println(c.tock, "[", (c.timeElapsed)/60, "minutes have elapsed ]")
	}
}

// printBong prints the bong value along with the elapsed time in hours.
func (c *Clock) printBong() {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	if c.print {
		fmt.Println(c.bong, "[", (c.timeElapsed)/3600, "hours have elapsed ]")
	}

}

// GetBong returns the bong value.
func (c *Clock) GetBong() string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.bong
}

// GetTock returns the tock value.
func (c *Clock) GetTock() string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.tock
}

// GetTick returns the tick value.
func (c *Clock) GetTick() string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.tick
}

// SetTick sets a new tick value.
func (c *Clock) SetTick(value string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.tick = value
}

// SetTock sets a new tock value.
func (c *Clock) SetTock(value string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.tock = value
}

// SetBong sets a new bong value.
func (c *Clock) SetBong(value string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.bong = value
}

// TogglePrint toggles the print status of the clock. It toggles between enable-print and disable-print
func (c *Clock) TogglePrint() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.print = !c.print
}

// StopTicker stops the ticker goroutine.
func (c *Clock) StopTicker() {
	close(c.stopTicker)
}

// Stop stops the clock by closing the Finished channel.
func (c *Clock) Stop() {
	close(c.Finished)
}
