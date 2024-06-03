package main

import (
	"clock-app/internal/clock"
	"testing"
)

// TestHandleUserCommand tests the handleUserCommand function to ensure it processes user commands correctly.
func TestHandleUserCommand(t *testing.T) {
	// Create a new Clock instance
	c := clock.NewClock("tick", "tock", "bong", 60)

	// Test setting the tick value
	handleUserCommand(c, "tick new_tick")
	if c.GetTick() != "new_tick" {
		t.Errorf("Expected new_tick, got %s", c.GetTick())
	}

	// Test setting the tock value
	handleUserCommand(c, "tock new_tock")
	if c.GetTock() != "new_tock" {
		t.Errorf("Expected new_tock, got %s", c.GetTock())
	}

	// Test setting the bong value
	handleUserCommand(c, "bong new_bong")
	if c.GetBong() != "new_bong" {
		t.Errorf("Expected new_bong, got %s", c.GetBong())
	}

	// Test handling an unknown command
	handleUserCommand(c, "unknown command")
	// Verify that unknown command does not change clock values
	if c.GetTick() != "new_tick" || c.GetTock() != "new_tock" || c.GetBong() != "new_bong" {
		t.Errorf("Unknown command should not change clock values")
	}
}
