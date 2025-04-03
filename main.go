// Filename: main.go
// Purpose: This program demonstrates how to create a TCP network connection using Go

package main

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
	"flag"
)

func worker(wg *sync.WaitGroup, tasks chan string, dialer net.Dialer) {
	defer wg.Done()
	maxRetries := 3
    for addr := range tasks {
		var success bool
		for i := range maxRetries {      
		conn, err := dialer.Dial("tcp", addr)
		if err == nil {
			conn.Close()
			fmt.Printf("Connection to %s was successful\n", addr)
			success = true
			break
		}
		backoff := time.Duration(1<<i) * time.Second
		fmt.Printf("Attempt %d to %s failed. Waiting %v...\n", i+1,  addr, backoff)
		time.Sleep(backoff)
	    }
		if !success {
			fmt.Printf("Failed to connect to %s after %d attempts\n", addr, maxRetries)
		}
	}
}

func main() {

	var wg sync.WaitGroup
	tasks := make(chan string, 100)

    // add a target flag with a default value of "scanme.nmap.org"
	target := flag.String("target", "scanme.nmap.org", "Target host to scan")

	//add a -start-port and -end-port flag 1-1024
	startPort := flag.Int("start-port", 1, "Start port to scan")
	endPort := flag.Int("end-port", 1024, "End port to scan")

	//adjusteable number of workers flag default 100
	workers := flag.Int("workers", 100, "Number of workers to use")

	//add a -timeout flag with a default value of 5 seconds
	timeout := flag.Int("timeout", 5, "Timeout in seconds")


	//parse all the flags
	flag.Parse()

	dialer := net.Dialer {
		Timeout: time.Duration(*timeout) * time.Second,
	}
  

    for i := 1; i <= *workers; i++ {
		wg.Add(1)
		go worker(&wg, tasks, dialer)
	}


	for p := *startPort; p <= *endPort; p++ {
		port := strconv.Itoa(p)
        address := net.JoinHostPort(*target, port)
		tasks <- address
	}
	close(tasks)
	wg.Wait()
	//create variable to store start time
	start := time.Now()

	//print summary with number of open ports, time taken, and total ports scanned
	fmt.Printf("Scanned %d ports on %s\n", *endPort-*startPort+1, *target)
	fmt.Printf("Total time taken: %v\n", time.Since(start))
}