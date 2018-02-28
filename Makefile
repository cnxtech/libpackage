GOFILES=$(shell find . -type f -name '*.go')

DOCKERFILES=$(shell find . -type f -name 'Dockerfile')

.DEFAULT: impact

image: $(GOFILES) $(DOCKERFILES)
	@echo "+ $@"
	@docker build -t hinshun/impact -f cmd/impact/Dockerfile .

bin/impact: $(GOFILES)
	@echo "+ $@"
	@go build -o ./bin/impact ./cmd/impact


.PHONY: image
