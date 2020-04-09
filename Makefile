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

# target dir
TARGETDEB=$(GOCMD)/deb/bterminal/usr/local/bin

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
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(TARGETDEB)/$(BINARY_NAME) -v $(GOBUILDDEB) 
build-arm:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm $(GOBUILD) -o $(TARGETDEB)/$(BINARY_NAME) -v $(GOBUILDDEB) 
build-arm_v5:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=5 $(GOBUILD) -o $(TARGETDEB)/$(BINARY_NAME) -v $(GOBUILDDEB) 
build-arm_v6:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=6 $(GOBUILD) -o $(TARGETDEB)/$(BINARY_NAME) -v $(GOBUILDDEB) 
build-arm_v7:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 $(GOBUILD) -o $(TARGETDEB)/$(BINARY_NAME) -v $(GOBUILDDEB) 
build-arm_v8:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=8 $(GOBUILD) -o $(TARGETDEB)/$(BINARY_NAME) -v $(GOBUILDDEB) 
build-arm64:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 $(GOBUILD) -o $(TARGETDEB)/$(BINARY_NAME) -v $(GOBUILDDEB) 
build-win:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -ldflags="-H windowsgui" -o $(BINARY_NAME).exe -v $(GOBUILDWIN)