// Filename: main.go
// Purpose: This program demonstrates how to create a TCP network connection using Go

package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// Global variable to track the number of scanned ports
var portsScanned int32

// Worker function to scan ports
func worker(wg *sync.WaitGroup, tasks chan string, dialer net.Dialer) {
	defer wg.Done()
	maxRetries := 3
	for addr := range tasks {
		var success bool
		for i := 0; i < maxRetries; i++ {
			conn, err := dialer.Dial("tcp", addr)
			if err == nil {
				fmt.Printf("Connection to %s was successful\n", addr)
				go grabBanner(conn, addr)
				conn.Close()
				success = true
				break
			}
			backoff := time.Duration(1<<i) * time.Second
			fmt.Printf("Attempt %d to %s failed. Waiting %v...\n", i+1, addr, backoff)
			time.Sleep(backoff)
		}
		atomic.AddInt32(&portsScanned, 1)
		if !success {
			fmt.Printf("Failed to connect to %s after %d attempts\n", addr, maxRetries)
		}
	}
}

// Function to grab the banner from a service
func grabBanner(conn net.Conn, addr string) {
	defer conn.Close()
	buffer := make([]byte, 1024)
	conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Printf("Error reading from %s: %v\n", addr, err)
		return
	}
	banner := string(buffer[:n])
	fmt.Printf("Banner for %s: %s\n", addr, banner)
}

// Function to print scan summary
func printSummary(start time.Time, endPort, startPort int, targets []string, jsonOutput bool) {
	if jsonOutput {
		// Output in JSON format
		fmt.Printf("{\"targets\": %v, \"ports_scanned\": %d, \"time_taken\": %.2f}\n",
			targets, (endPort-startPort+1)*len(targets), time.Since(start).Seconds())
	} else {
		// Print summary with number of open ports, time taken, and total ports scanned
		fmt.Printf("Scanned %d ports on %d targets\n", (endPort-startPort+1), len(targets))
		fmt.Printf("Total time taken: %.2f seconds\n", time.Since(start).Seconds())
	}
}

func main() {
	startTime := time.Now()
	var wg sync.WaitGroup
	tasks := make(chan string, 100)

	// Define a string flag for multiple targets (comma-separated)
	targetsStr := flag.String("target", "scanme.nmap.org", "Comma-separated list of target hosts to scan")

	// Define port range flags
	startPort := flag.Int("start-port", 1, "Start port to scan")
	endPort := flag.Int("end-port", 1024, "End port to scan")

	// Define a flag for number of workers
	workers := flag.Int("workers", 100, "Number of workers to use")

	// Define a flag for timeout duration
	timeout := flag.Int("timeout", 5, "Timeout in seconds")

	// Define a flag for JSON output
	jsonOutput := flag.Bool("json", false, "Output in JSON format")

	//Define a flag for specific ports specified by comma separated values
	portsStr := flag.String("ports", "", "Comma-separated list of specific ports to scan")

	// Parse the command-line flags
	flag.Parse()

	// Convert target string into a slice
	targets := strings.Split(*targetsStr, ",")

	var ports []int
	// If specific ports are provided, parse them
	if *portsStr != "" {
		// Split the ports string by comma
		for _, p := range strings.Split(*portsStr, ",") {
			// Convert each port to an integer
			port, err := strconv.Atoi(strings.TrimSpace(p))
			if err == nil {
				ports = append(ports, port)
			}
		}
	} else {
		// If no specific ports are provided, use the range
		for p:= *startPort; p <= *endPort; p++ {
			ports = append(ports, p)
		}
	}

	// Print all targets
	fmt.Println("Scanning the targets:")
	for _, target := range targets {
		fmt.Println(target)
	}

	dialer := net.Dialer{
		Timeout: time.Duration(*timeout) * time.Second,
	}

	// Start the progress indicator
	go func() {
		for {
			time.Sleep(1 * time.Second)
			scannedPorts := atomic.LoadInt32(&portsScanned)
			fmt.Printf("\rScanning %d of %d ports\n", scannedPorts, len(ports)*len(targets))
			if scannedPorts >= int32(len(ports)*len(targets)) {
				break
			}
		}
	}()


	// Start worker goroutines
	for i := 1; i <= *workers; i++ {
		wg.Add(1)
		go worker(&wg, tasks, dialer)
	}

	// Enqueue tasks for each target and port
	for _, target := range targets {
		for _, port := range ports {
			address := net.JoinHostPort((strings.TrimSpace(target)), strconv.Itoa(port))
			tasks <- address
		}
	}
	close(tasks)
	wg.Wait()

	// Print the summary
	printSummary(startTime, *endPort, *startPort, targets, *jsonOutput)
}
