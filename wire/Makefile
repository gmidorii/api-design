.PHONY: wire build run run-reload

BIN=wire-app
AUTO_PORT=3333
APP_PORT=8080
WIRE_BIN=wire
TAG=latest

wire:
	$(WIRE_BIN) gen github.com/gmidorii/api-design/wire/provider

build: wire
	go build -o $(BIN)

run: build
	./$(BIN) -p $(APP_PORT)

run-reload:
	gin -a $(APP_PORT) -p $(AUTO_PORT)


d-build:
	docker build -t wire-app:$(TAG) .

d-run:
	docker run --name wire-app -p 8080:8080 -d wire-app:$(TAG)

d-rm:
	docker rm -f wire-app