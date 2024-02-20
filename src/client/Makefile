PROTO_FILE := helloworld.proto
PROTO_FILE := ../$(PROTO_FILE)

all: init
	python3 greeter_client.py

init: $(PROTO_FILE)
	python3 -m grpc_tools.protoc -I../ --python_out=. --pyi_out=. --grpc_python_out=. $<