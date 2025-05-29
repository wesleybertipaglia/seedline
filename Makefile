BINARY_NAME=seedline

all: build

build:
	go build -o $(BINARY_NAME) main.go

run: build
	./$(BINARY_NAME)

clean:
	rm -f $(BINARY_NAME)

fmt:
	go fmt ./...

rebuild: clean build

help:
	@echo "Available commands:"
	@echo "  make build    - Build the game binary"
	@echo "  make run      - Build and run the game"
	@echo "  make clean    - Remove the binary"
	@echo "  make fmt      - Format the Go source files"
	@echo "  make rebuild  - Clean and build the binary again"
	@echo "  make help     - Show this help message"

.PHONY: all build run clean fmt rebuild help
