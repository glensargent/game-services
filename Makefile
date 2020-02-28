GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
ENV=CGO_ENABLED=0 GOOS="linux" GOARCH=amd64

CREDS="src/configs/creds.json"
ORCHESTRATOR=./src/orchestrator

PROTODIR=./pb


# Run the service with google creds
run-orchestrator:
	go run ./src/orchestrator/*.go

build: 
	$(ENV) $(GOBUILD) -o $(ORCHESTRATOR) $(ORCHESTRATOR)

test: 
	$(GOTEST) -v ./...