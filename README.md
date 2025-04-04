# systemsTest1

## Description

**systemsTest1** is a concurrent TCP port scanner designed to efficiently probe one or more target hosts for open ports within a specified range or a custom list. It supports configurable worker threads, connection timeouts, and banner grabbing. Users can run the scanner with various command-line flags to control the scan behavior and output format, including JSON for integration with other tools.

## Available Flags

| **Flag**       | **Type**   | **Description**                                                                        |
|---------------|-----------|----------------------------------------------------------------------------------------|
| `-target`     | `string`   | Comma-separated list of target hosts to scan (default: `scanme.nmap.org`)              |
| `-start-port` | `int`      | Starting port number to scan (default: `1`)                                            |
| `-end-port`   | `int`      | Ending port number to scan (default: `1024`)                                           |
| `-ports`      | `string`   | Comma-separated list of specific ports to scan (overrides `start-port` & `end-port`)   |
| `-workers`    | `int`      | Number of concurrent worker goroutines to use (default: `100`)                         |
| `-timeout`    | `int`      | Timeout in seconds for each connection attempt (default: `5`)                          |
| `-json`       | `bool`     | Output scan summary in JSON format (default: `false`)                                  |

## How to Build and Run

### Run Without Building
```sh
make run FLAGS="-target=localhost -start-port=22 -end-port=25"
```

### Build and Run the Compiled Binary
```sh
make exec FLAGS="-target=scanme.nmap.org -ports=80,443"
```

### Just Build
```sh
make build
```

### Clean Up
```sh
make clean
```

## Sample Output
```
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
```

## Video Demonstration
[![Watch the video](https://img.youtube.com/vi/3Vn9OJb0l1wh/0.jpg)](https://youtu.be/3Vn9OJb0l1wh)

Click the link to watch: [https://youtu.be/3Vn9OJb0l1wh](https://youtu.be/3Vn9OJb0l1wh)

---

This project provides a fast and efficient way to scan open ports with detailed logging and retry mechanisms. Feel free to contribute or report issues on GitHub!

