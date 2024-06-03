package main

import (
	"bufio"
	"clock-app/config"
	"clock-app/internal/clock"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
)

// handleUserCommand processes a single user command.
func handleUserCommand(c *clock.Clock, completeCommand string) {

	// Splitting the input command into two parts: the command and the value.
	parts := strings.SplitN(completeCommand, " ", 2)
	if len(parts) == 2 {
		command, value := parts[0], parts[1]
		switch command {
		case "tick":
			// Set the tick message
			c.SetTick(value)
		case "tock":
			// Set the tock message
			c.SetTock(value)
		case "bong":
			// Set the bong message
			c.SetBong(value)
		default:
			fmt.Println("Unknown command:", command)
		}
	} else {
		fmt.Println("Invalid input")
	}
}

// handleInput reads user input from the console and processes commands.
func handleInput(c *clock.Clock, cancel context.CancelFunc, stopOnce *sync.Once) {

	// Creating a new reader for standard input
	reader := bufio.NewReader(os.Stdin)
	for {
		// Printing the instructions for the user
		fmt.Println("To change the tick, tock, and bong values, or quit, disable print first by toggling the print settings, enter the necessary commands,\nand enable print by toggling the print settings again.\nType any of the below given commands and press enter:\ntick [value]: For changing the default tick message,\ntock [value]: For changing the default tock message,\nbong [value]: For changing the default bong message,\nt : For toggling the print settings anytime,\nquit: For stopping the time counter any time that you would like")

		// Reading input from the user
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		input = strings.TrimSpace(input)
		switch input {
		case "quit":
			// Stop the application if the user types "quit", and only ensure that the cancel function is executed once to ensure there is no panic
			stopOnce.Do(cancel)
			return
		case "t":
			// Toggle the print settings between enable-print and disable-print if the user types "t", by default print is enabled
			c.TogglePrint()
		default:
			// Handle other commands
			handleUserCommand(c, input)
		}

	}
}

func main() {

	// Load configuration values from the environment or the .env file
	tick, tock, bong, clockLimit := config.LoadEnv()

	log.Println("Starting clock application...")

	// Create a new clock instance
	clock := clock.NewClock(tick, tock, bong, clockLimit)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//Start running the clock
	clock.Run(cancel)

	stopOnce := &sync.Once{}

	//start the go routine that handles the input commands
	go handleInput(clock, cancel, stopOnce)

	// Handle graceful shutdown on SIGINT and SIGTERM
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-sigChan:
		log.Println("Received shutdown signal")
		//execute the cancel function only once to ensure there is no panic
		stopOnce.Do(cancel)
	case <-ctx.Done():
		log.Println("Context canceled")
	case <-clock.Finished:
		log.Println("Finished Ticking for the time that was set")
	}

	//stop the ticker initialised in the Run function
	clock.StopTicker()
	log.Println("Shutting down gracefully...")
}
