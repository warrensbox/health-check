EXE  := health-check
VER  := $(shell git describe --tags)
PATH := build:$(PATH)
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

.PHONY: darwin linux windows release clean dist-clean test

$(EXE): go.mod *.go
	go build -v -ldflags '-s -w -X main.Version=$(VER)' -o $@

release: $(EXE) darwin windows linux

darwin linux:
	GOOS=$@ go build -ldflags '-s -w -X main.Version=$(VER)' -o $(EXE)-$(VER)-$@-$(GOARCH)
	upx -v $(EXE)-$(VER)-$@-$(GOARCH)

# $(shell go env GOEXE) is evaluated in the context of the Makefile host (before GOOS is evaluated), so hard-code .exe
windows:
	GOOS=$@ go build -ldflags '-s -w -X main.Version=$(VER)' -o $(EXE)-$(VER)-$@-$(GOARCH).exe
	upx -v $(EXE)-$(VER)-$@-$(GOARCH).exe

clean:
	rm -f $(EXE) $(EXE)-*-*-*

dist-clean: clean
	rm -f go.sum

docker: clean linux Dockerfile
	docker build . -t $(EXE):$(VER)

test: $(EXE)
	mkdir -p build
	mv $(EXE) build
	go test -v ./...
	bundle install