ARTIFACTS=

ARTIFACTS += vstreamer_protos/src/voicerecog.rs

.PHONY: all
all: ${ARTIFACTS}

vstreamer_protos/src/voicerecog.rs: ../protos/vstreamer_protos/commander/commander.proto
	cargo run -r -p vstp_builder



.PHONY: clean
clean:
	rm -f ${ARTIFACTS}
