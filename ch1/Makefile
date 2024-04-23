.DEFAULT_GOAL := build

.PHONY:fmt vet build

fmt:
				go fmt ./...

vet: fmt
				go vet ./...

build: vet
				go build -o world_hello

clean:
		go clean
		rm -rf world_hello


