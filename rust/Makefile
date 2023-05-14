ARTIFACTS=

ARTIFACTS += dist/src/voicerecog.rs

.PHONY: all
all: ${ARTIFACTS}

dist/src/voicerecog.rs: ../protos/vstreamer_protos/commander/commander.proto
	cargo run



.PHONY: clean
clean:
	rm -f ${ARTIFACTS}
