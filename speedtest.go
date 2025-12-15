package main

import (
	"fmt"
	"sync"

	"github.com/showwin/speedtest-go/speedtest"
)

// To customize the server string representation. Verbose mode.
type MyServer speedtest.Server

func (s MyServer) String() string {
	return fmt.Sprintf(
		"URL:\t\t%s\n"+
			"Lat:\t\t%s\n"+
			"Lon:\t\t%s\n"+
			"Name:\t\t%s\n"+
			"Country:\t%s\n"+
			"Sponsor:\t%s\n"+
			"Host:\t\t%s\n"+
			"ID:\t\t%s\n"+
			"Distance:\t%f\n"+
			"Jitter:\t\t%s (latency fluctuation)\n"+
			"TestDuration:\t%s\n",
		s.URL, s.Lat, s.Lon, s.Name, s.Country, s.Sponsor, s.Host, s.ID, s.Distance, s.Jitter, s.TestDuration,
	)
}

// runSpeedTest performs the speed test and prints results.
func runSpeedTest(wg *sync.WaitGroup, verbose bool) {
	defer wg.Done() // ensure the wait group counter is decremented when the spinner stops

	// Initialize speedtest client and perform tests.
	var speedtestClient = speedtest.New()

	serverList, err := speedtestClient.FetchServers() // fetch server list
	if err != nil {
		fmt.Println("Error fetching server list:", err)
		return
	}

	targets, err := serverList.FindServer([]int{}) // returns 1 target, the most convenient
	if err != nil {
		fmt.Println("Error finding server:", err)
		return
	}

	s := targets[0]
	s.PingTest(nil)
	s.DownloadTest()
	s.UploadTest()

	fmt.Println()
	fmt.Println(s)
	fmt.Printf("Latency: %s, Download: %s, Upload: %s\n", s.Latency, s.DLSpeed, s.ULSpeed)

	if verbose {
		var mys MyServer = MyServer(*s)
		fmt.Printf("\nDetailed server info:\n%s\n", mys)
	}
}
