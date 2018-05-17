PACKAGES := \
	github.com/marcoarthur/netcoord
DEPENDENCIES := \
	github.com/hashicorp/serf/coordinate \
	github.com/davecgh/go-spew/spew

all: build silent-test

build:
	go build netcoord.go

test:
	go test -v $(PACKAGES)

silent-test:
	go test $(PACKAGES)

format:
	go fmt $(PACKAGES)

deps:
	go get $(DEPENDENCIES)
