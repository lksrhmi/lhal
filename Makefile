APP_NAME=lhal
CMD_PATH=./cmd/lhal
BIN_DIR=./bin

build:
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(APP_NAME) $(CMD_PATH)

run: build
	./$(BIN_DIR)/$(APP_NAME)

install:
	go install $(CMD_PATH)

clean:
	rm -rf $(BIN_DIR)

dev:
	go run $(CMD_PATH)