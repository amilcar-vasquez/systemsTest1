# Makefile for building and running main.go with command-line flags

# Default to no flags unless specified by the user
FLAGS ?=

.PHONY: all build run clean

# Build the binary
build:
	@echo "Building the project..."
	go build -o scanner main.go
	@echo "Build complete. Run with ./scanner $(FLAGS)"

# Run the program directly with go run
run:
	@echo "Running with flags: $(FLAGS)"
	go run main.go $(FLAGS)

# Build and run the compiled binary
exec: build
	@echo "Executing the binary with flags: $(FLAGS)"
	./scanner $(FLAGS)

# Clean up the built binary
clean:
	@echo "Cleaning up..."
	rm -f scanner
	@echo "Clean complete."
