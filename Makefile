# Variables
APP_NAME = mynewapp
GO_FILES = $(shell find . -name '*.go')

# Default target
all: build

# Build the application
build: $(GO_FILES)
	go build -o $(APP_NAME)

# Run tests
test:
	go test ./...

# Clean up build artifacts
clean:
	rm -f $(APP_NAME)

# Install the application
install: build
	sudo mv $(APP_NAME) /usr/local/bin/

# Clean install the application
cleanInstall: clean install

# Run the application
run: build
	./$(APP_NAME) add --help