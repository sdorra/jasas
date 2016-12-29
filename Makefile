DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

BINARY=jasas
DIST_DIR=dist
IMAGE=sdorra/jasas

# These are the values we want to pass for Version and BuildTime
VERSION=0.1.0
BUILD_TIME=`date +%FT%T%z`
COMMIT_ID=`git rev-parse HEAD`

PACKAGES=$(shell go list ./... | grep -v /vendor/)

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.BuildTime=${BUILD_TIME} -X main.CommitID=${COMMIT_ID}"

.DEFAULT_GOAL: $(BINARY)

WEBUI:
	@echo "building webui ..."
	@yarn run build-prod

$(BINARY): 
	@echo "building go binary ..."
	@mkdir -p $(DIST_DIR) $(BINARY_DIR)
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -installsuffix cgo $(LDFLAGS) -o $(DIST_DIR)/$(BINARY)
	@echo "... binary can be found at $(DIST_DIR)/$(BINARY)"

docker: WEBUI $(BINARY)
	@docker build -t $(IMAGE):$(VERSION) .

push: docker
	@docker push $(IMAGE):$(VERSION)
