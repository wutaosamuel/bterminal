# Go Parameters
GOCC=go
GOBUILD=$(GOCC) build
GOCLEAN=$(GOCC) clean
GOTEST=$(GOCC) test
#GOGET=$(GOCC) get
BINARY_NAME=bterminal

# CMD dir
GOCMD=cmd
GOBUILDWIN=$(GOCMD)/bterminalWin
GOBUILDLNX=$(GOCMD)/bterminal

all: test build
build:
	$(GOBUILD) $(GOBUILDLNX)/. -o $(BINARY_NAME) -v
test:
	$(GOTEST) v ./...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
# run:
# deps:

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) $(GOBUILDLNX)/. -o $(BINARY_NAME) -v
build-win:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) $(GOBUILDWIN)/. -o $(BINARY_NAME) -v