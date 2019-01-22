BIN=wire-app
AUTO_PORT=3333
APP_PORT=8080

build:
	go build -o $(BIN)

run: build
	./$(BIN) -p $(APP_PORT)

run-reload:
	gin -a $(APP_PORT) -p $(AUTO_PORT)
