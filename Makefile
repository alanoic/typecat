# Variables for easy updates
BINARY_NAME=typecat
BUILD_DIR=bin

# The default 'make' command
all: build

# Build the binary
build:
	@echo "Building..."
	go build -o $(BUILD_DIR)/$(BINARY_NAME) .

# Run the application
run: build
	./$(BUILD_DIR)/$(BINARY_NAME)

# Clean up the binary folder
clean:
	@echo "Cleaning up..."
	rm -rf $(BUILD_DIR)

# Install the binary to $GOPATH/bin
install:
	@echo "Installing typecat..."
	go install .

# Help command to show what's available
help:
	@echo "Available commands:"
	@echo "  make build   - Build the executable"
	@echo "  make run     - Build and run the executable"
	@echo "  make install - Install to \$GOPATH/bin"
	@echo "  make clean   - Remove the bin folder"