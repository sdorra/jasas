DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

BINARY=jasas
DIST_DIR=dist
IMAGE=sdorra/jasas

OSARCH="darwin/amd64 linux/amd64 linux/arm windows/amd64"

# These are the values we want to pass for Version and BuildTime
VERSION=0.1.0
BUILD_TIME=`date +%FT%T%z`
COMMIT_ID=`git rev-parse HEAD`

PACKAGES=$(shell go list ./... | grep -v /vendor/)

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.BuildTime=${BUILD_TIME} -X main.CommitID=${COMMIT_ID}"

.DEFAULT_GOAL: build

.PHONY: build
build: $(GOPATH)/bin/gox
	@echo "building go binaries ..."
	@go generate
	@mkdir -p $(DIST_DIR)
	@gox -osarch ${OSARCH} -output "dist/${BINARY}_{{.OS}}_{{.Arch}}" ./...
	@cd $(DIST_DIR); shasum -a 256 * > ${BINARY}.sha256sums
	@cd $(DIST_DIR);  gpg --armor --detach-sign jasas.sha256sums
	@echo "... binaries can be found at $(DIST_DIR)"

$(GOPATH)/bin/gox:
	@echo installing gox
	@go get github.com/mitchellh/gox

.PHONY: docker
docker:
	@docker build -t $(IMAGE):$(VERSION) .

.PHONY: push
push: docker
	@docker push $(IMAGE):$(VERSION)

.PHONY: clean
clean:
	@rm -rf ${DIST_DIR}
