.DEFAULT_GOAL := build

.PHONY:fmt vet build

fmt:
				go fmt ./...

vet: fmt
				go vet ./...

build: vet
				go build -o var2

clean:
		go clean
		rm -rf var


