PROTO_FILE := helloworld.proto

SERVER_DIR := server
CLIENT_DIR := client

all: go python

go: $(PROTO_FILE)
	protoc --go_out=./$(SERVER_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=./$(SERVER_DIR) --go-grpc_opt=paths=source_relative $<

python: $(PROTO_FILE)
	python3 -m grpc_tools.protoc -I. --python_out=./$(CLIENT_DIR) --pyi_out=./$(CLIENT_DIR) \
		--grpc_python_out=./$(CLIENT_DIR) $<
