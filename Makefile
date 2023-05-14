.PHONY: all
all: go python rust

.PHONY: go python rust
go:
	$(MAKE) -C go
python:
	$(MAKE) -C python
rust:
	$(MAKE) -C rust

.PHONY: clean
clean:
	$(MAKE) -C go clean
	$(MAKE) -C python clean
	$(MAKE) -C rust clean
