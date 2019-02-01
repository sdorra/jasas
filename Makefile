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

GITHUB_USER="sdorra"
GITHUB_REPO="jasas"

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-"-X main.Version=${VERSION} -X main.BuildTime=${BUILD_TIME} -X main.CommitID=${COMMIT_ID}"

.DEFAULT_GOAL: build

.PHONY: build
build: $(GOPATH)/bin/gox
	@echo "building go binaries ..."
	@go generate
	@mkdir -p $(DIST_DIR)
	@gox -osarch ${OSARCH} -ldflags ${LDFLAGS} -output "dist/${BINARY}_{{.OS}}_{{.Arch}}" ./...
	@cd $(DIST_DIR); shasum -a 256 * > ${BINARY}.sha256sums
	@cd $(DIST_DIR); gpg --armor --detach-sign jasas.sha256sums
	@echo "... binaries can be found at $(DIST_DIR)"

.PHONY: release
release: $(GOPATH)/bin/github-release build push
	@echo "creating release ..."
	@git tag -s -m "release v${VERSION}" v${VERSION}
	@git push origin master --tags
	@github-release release \
  		--user ${GITHUB_USER} \
  		--repo ${GITHUB_REPO} \
  		--tag v${VERSION} \
  		--name v${VERSION} \
  		--description "release version ${VERSION}"
	@cd ${DIST_DIR}; ls -1 | xargs -n1 -I{} -- \
	 									github-release upload \
																	--user ${GITHUB_USER} \
																	--repo ${GITHUB_REPO} \
																	--tag v${VERSION} \
																	--name {} \
																	--file {}

$(GOPATH)/bin/github-release:
	@echo installing github-release
	@go get github.com/aktau/github-release

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
	@rm -f templates/templates_prod.go
	@rm -rf ${DIST_DIR}
