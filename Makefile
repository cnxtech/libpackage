GOFILES=$(shell find . -type f -name '*.go')

.DEFAULT: bin/impact

bin/impact: $(GOFILES)
	@echo "+ $@"
	@go build -o ./bin/impact ./cmd/impact

.PHONY: all binaries images clean vendor
