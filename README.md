# systemsTest1
:: Description::
This tool is a concurrent TCP port scanner, designed to efficiently probe one or more target hosts for open ports within a specified range or a custom list. It supports configurable worker threads, connection timeouts, banner grabbing. Users can run the scanner with various command-line flags to control the scan behavior and output format, including JSON for integration with other tools. 

Flag	        Type	    Description
-target	        string	    Comma-separated list of target hosts to scan (default: "scanme.nmap.org")
-start-port	    int	        Starting port number to scan (default: 1)
-end-port	    int	        Ending port number to scan (default: 1024)
-ports	        string	    Comma-separated list of specific ports to scan (overrides start-port & end-port)
-workers	    int	        Number of concurrent worker goroutines to use (default: 100)
-timeout	    int	        Timeout in seconds for each connection attempt (default: 5)
-json	        bool	    Output scan summary in JSON format (default: false)


How to build and run (examples provided):
run without building: make run FLAGS="-target=localhost -start-port=22 -end-port=25"
Build and run the compiled binary: make exec FLAGS="-target=scanme.nmap.org -ports=80,443"
Just build: make build
Clean up: make clean

Sample output:
Running with flags: -target=localhost -start-port=22 -end-port=25
go run main.go -target=localhost -start-port=22 -end-port=25
Scanning the targets:
localhost
Attempt 1 to localhost:22 failed. Waiting 1s...
Attempt 1 to localhost:24 failed. Waiting 1s...
Attempt 1 to localhost:25 failed. Waiting 1s...
Attempt 1 to localhost:23 failed. Waiting 1s...
Scanning 0 of 4 ports
Attempt 2 to localhost:22 failed. Waiting 2s...
Attempt 2 to localhost:23 failed. Waiting 2s...
Attempt 2 to localhost:25 failed. Waiting 2s...
Attempt 2 to localhost:24 failed. Waiting 2s...
Scanning 0 of 4 ports
Scanning 0 of 4 ports
Attempt 3 to localhost:24 failed. Waiting 4s...
Attempt 3 to localhost:23 failed. Waiting 4s...
Attempt 3 to localhost:22 failed. Waiting 4s...
Attempt 3 to localhost:25 failed. Waiting 4s...
Scanning 0 of 4 ports
Scanning 0 of 4 ports
Scanning 0 of 4 ports
Scanning 0 of 4 ports
Failed to connect to localhost:22 after 3 attempts
Failed to connect to localhost:24 after 3 attempts
Failed to connect to localhost:23 after 3 attempts
Failed to connect to localhost:25 after 3 attempts
Scanned 4 ports on 1 targets
Total time taken: 7.00 seconds

video file link:
https://youtu.be/3Vn9OJb0l1w