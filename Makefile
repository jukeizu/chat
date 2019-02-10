TAG=$(shell git describe --tags)
VERSION=$(TAG:v%=%)
REPO=jukeizu/chat
GO=GO111MODULE=on go
BUILD=GOARCH=amd64 $(GO) build -ldflags="-s -w -X main.Version=$(VERSION)" 

.PHONY: all deps test build build-linux docker-build docker-deploy clean

all: deps test build 
deps:
	$(GO) mod download

test:
	$(GO) vet
	$(GO) test -v -race

build:
	$(BUILD) -o bin/chat-$(VERSION)

build-linux:
	CGO_ENABLED=0 GOOS=linux $(BUILD) -a -installsuffix cgo -o bin/chat

docker-build:
	docker build -t $(REPO):$(VERSION) .

docker-deploy:
	docker push $(REPO):$(VERSION)

clean:
	@find bin -type f -delete -print
