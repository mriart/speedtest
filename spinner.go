package main

import (
	"fmt"
	"sync"
	"time"
)

// spinner displays a spinning animation until the wait group is signaled.
func spinner(wg *sync.WaitGroup, stopCh <-chan struct{}) {
	defer wg.Done() // ensure the wait group counter is decremented when the spinner stops

	// Define the characters for the spinner animation.
	chars := []rune{'\\', '|', '/', '-'}
	i := 0

	// Loop indefinitely until a signal is received on the stop channel.
	for {
		select {
		case <-stopCh:
			// Received signal to stop, print a clear line and exit the goroutine.
			fmt.Print("\r                      \r")
			return
		default:
			// Print the current spinner character and immediately reset the cursor to the start of the line.
			// \r is the carriage return character.
			fmt.Printf("\rProcessing speedtest... %c", chars[i])
			i = (i + 1) % len(chars)

			// Control the speed of the spinner
			time.Sleep(100 * time.Millisecond)
		}
	}
}
