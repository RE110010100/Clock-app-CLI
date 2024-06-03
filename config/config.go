package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// GetEnvWithStringValues retrieves a string value from the environment variables
// or returns the default value if the key is not found.
func GetEnvWithStringValues(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// GetEnvWithIntegerValues retrieves an integer value from the environment variables
// or returns the default value if the key is not found or the value is not an integer.
func GetEnvWithIntegerValues(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		// Convert string to integer
		num, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println("Error with clock limit env variable, it does not seem to be an integer, hence using default value of 3hrs")
			return defaultValue
		}
		return num
	}
	return defaultValue
}

// LoadEnv loads the environment variables from a .env file (if present) and returns
// the values for TICK_VALUE, TOCK_VALUE, BONG_VALUE, and CLOCK_LIMIT.
func LoadEnv() (string, string, string, int) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found. Using default environment variables.")
	}
	tick := GetEnvWithStringValues("TICK_VALUE", "tick")
	tock := GetEnvWithStringValues("TOCK_VALUE", "tock")
	bong := GetEnvWithStringValues("BONG_VALUE", "bong")
	clockLimit := GetEnvWithIntegerValues("CLOCK_LIMIT", 3*3600)

	return tick, tock, bong, clockLimit
}
