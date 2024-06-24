# If no ran target is specified, then run build
.DEFAULT_GOAL := build

# removes executable
clean: 
	go clean

# formats all Go files in current directory
fmt: clean
	go fmt ./...

# performs static analysis on Go code
vet: fmt
	go vet ./...

# compiles Go code; creates executable
build: fmt
	go build