VERSION ?= unknown

# Packages
TOP_PACKAGE_DIR := github.com/diskordanz
PACKAGE_LIST := weather_notifier

all: deps test build

test:
	go test -v --cover --race -short `glide novendor | grep -v ./proto`

build:
	CGO_ENABLED=0 go build -a -ldflags '-s -X main.serviceVersion=$(VERSION)' -installsuffix cgo -o main ./cmd/notifier/

deps:
	glide install

lint: 
	@for p in $(PACKAGE_LIST); do \
		echo "==> Lint $$p ..."; \
		golint $(TOP_PACKAGE_DIR)/$$p; \
	done

run:
	API_URL=http://localhost:8002 \
	API_KEY=b04ad8db6f75cbd1a02e6e4c8e1e1272 \
	./main
