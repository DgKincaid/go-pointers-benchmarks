.PHONY: benchmark
benchmark:
	go test -gcflags "-m -m" -bench . -benchmem ./internal/...
