ARTIFACTS=

ARTIFACTS += vstreamer_protos/commander/commander_pb2.py
ARTIFACTS += vstreamer_protos/commander/commander_pb2_grpc.py

PYTHON ?= .venv/bin/python

.PHONY: all
all: ${ARTIFACTS}

vstreamer_protos/commander/commander_pb2.py vstreamer_protos/commander/commander_pb2_grpc.py: ../protos/vstreamer_protos/commander/commander.proto
	protoc -I../protos --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ..\protos\vstreamer_protos\commander\commander.proto

.PHONY: clean
clean:
	rm -f ${ARTIFACTS}
