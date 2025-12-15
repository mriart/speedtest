// Speedtest CLI to know the speed of your internet connection.
// With spinner and verbose output option.
// MRS 202512

package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	// Check for -v verbose flag.
	verbose := false
	switch {
	case len(os.Args) == 1:
		verbose = false
	case len(os.Args) == 2 && os.Args[1] == "-v":
		verbose = true
	default:
		fmt.Println("Usage: speedtest [-v]  (for verbose output)")
		return
	}

	// Prepare to run speed test with spinner, all channels and wait groups.
	var speedTestWg sync.WaitGroup
	speedTestWg.Add(1)
	stopSpinnerCh := make(chan struct{})

	var spinnerWg sync.WaitGroup
	spinnerWg.Add(1)

	// Start the spinner and the speed test in their own goroutines.
	go spinner(&spinnerWg, stopSpinnerCh)
	go runSpeedTest(&speedTestWg, verbose)

	// Wait for the long-running task to complete and signal the spinner to stop.
	speedTestWg.Wait()
	close(stopSpinnerCh)

	// Wait for the spinner goroutine to finish its cleanup.
	spinnerWg.Wait()
}
