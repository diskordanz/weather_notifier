VERSION ?= unknown

all: deps test build

test:
	go test -v --cover --race -short `glide novendor | grep -v ./proto`

deps:
	glide install

lint: 
	golint $$(go list ./... | grep -v /vendor/)

build: 
	CGO_ENABLED=0 go build -a -ldflags '-s -X main.serviceVersion=$(VERSION)' -installsuffix cgo -o ./cmd/notifier/notifier ./cmd/notifier/
		
run:
	API_URL=http://localhost:8002 \
	API_KEY=b04ad8db6f75cbd1a02e6e4c8e1e1272 \
	./cmd/notifier/notifier
