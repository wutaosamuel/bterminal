# Go Parameters
GOCC=go
GOBUILD=$(GOCC) build
GOCLEAN=$(GOCC) clean
GOTEST=$(GOCC) test
#GOGET=$(GOCC) get
BINARY_NAME=bterminal

# CMD dir
GOCMD=./cmd
GOBUILDWIN=$(GOCMD)/bterminalWin
GOBUILDLNX=$(GOCMD)/bterminal
GOBUILDDEB=$(GOCMD)/deb

all: test build
build:
	$(GOBUILD) -o $(BINARY_NAME) -v $(GOBUILDLNX)
test:
	$(GOTEST) v ./...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
# run:
deps:
	$(GOGET) -u 

# Cross compilation
build-deb:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME) -v $(GOBUILDDEB) 
build-win:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -ldflags="-H windowsgui" -o $(BINARY_NAME).exe -v $(GOBUILDWIN)