package config

import (
	"os"
	"testing"
)

// Tests the GetEnvWithStringValues function
func TestGetEnvWithStringValues(t *testing.T) {
	key := "TEST_STRING_KEY"
	defaultValue := "default_value"

	// Test case when environment variable is set
	expectedValue := "test_value"
	os.Setenv(key, expectedValue)
	value := GetEnvWithStringValues(key, defaultValue)
	if value != expectedValue {
		t.Errorf("Expected %s, got %s", expectedValue, value)
	}

	// Test case when environment variable is not set
	os.Unsetenv(key)
	value = GetEnvWithStringValues(key, defaultValue)
	if value != defaultValue {
		t.Errorf("Expected %s, got %s", defaultValue, value)
	}
}

// Tests the GetEnvWithIntegerValues function
func TestGetEnvWithIntegerValues(t *testing.T) {
	key := "TEST_INT_KEY"
	defaultValue := 42

	// Test case when environment variable is set and is a valid integer
	expectedValue := 100
	os.Setenv(key, "100")
	value := GetEnvWithIntegerValues(key, defaultValue)
	if value != expectedValue {
		t.Errorf("Expected %d, got %d", expectedValue, value)
	}

	// Test case when environment variable is set and is not a valid integer
	os.Setenv(key, "invalid")
	value = GetEnvWithIntegerValues(key, defaultValue)
	if value != defaultValue {
		t.Errorf("Expected %d, got %d", defaultValue, value)
	}

	// Test case when environment variable is not set
	os.Unsetenv(key)
	value = GetEnvWithIntegerValues(key, defaultValue)
	if value != defaultValue {
		t.Errorf("Expected %d, got %d", defaultValue, value)
	}
}

// TestLoadEnv tests the LoadEnv function
func TestLoadEnv(t *testing.T) {
	// Set environment variables for the test
	os.Setenv("TICK_VALUE", "tick_test")
	os.Setenv("TOCK_VALUE", "tock_test")
	os.Setenv("BONG_VALUE", "bong_test")
	os.Setenv("CLOCK_LIMIT", "3600")

	defer os.Unsetenv("TICK_VALUE")
	defer os.Unsetenv("TOCK_VALUE")
	defer os.Unsetenv("BONG_VALUE")
	defer os.Unsetenv("CLOCK_LIMIT")

	tick, tock, bong, clockLimit := LoadEnv()

	if tick != "tick_test" {
		t.Errorf("expected 'tick_test', got '%s'", tick)
	}
	if tock != "tock_test" {
		t.Errorf("expected 'tock_test', got '%s'", tock)
	}
	if bong != "bong_test" {
		t.Errorf("expected 'bong_test', got '%s'", bong)
	}
	if clockLimit != 3600 {
		t.Errorf("expected 3600, got %d", clockLimit)
	}
}
