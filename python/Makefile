ARTIFACTS=

ARTIFACTS += vstreamer_protos/commander/commander_pb2.py
ARTIFACTS += vstreamer_protos/commander/commander_pb2_grpc.py
ARTIFACTS += vstreamer_protos/commander/commander_pb2.pyi

PYTHON ?= .venv/bin/python

.PHONY: all
all: ${ARTIFACTS}

vstreamer_protos/commander/commander_pb2.py vstreamer_protos/commander/commander_pb2_grpc.py vstreamer_protos/commander/commander_pb2.pyi: ../protos/vstreamer_protos/commander/commander.proto
	${PYTHON} -m grpc_tools.protoc -I../protos --python_out=. --pyi_out=. --grpc_python_out=. ../protos/vstreamer_protos/commander/commander.proto



.PHONY: clean
clean:
	rm -f ${ARTIFACTS}
