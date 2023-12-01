COVERPROFILE=coverage.out

build:
	go build

install:
	go install

bi: 
	go build
	go install

cover:
	go test ./... -coverprofile=$(COVERPROFILE)
	go tool cover -html=$(COVERPROFILE)
	del $(COVERPROFILE)

test:
	go test -i ./...
	go test -v ./...

.PHONY: coverage dependencies test
.PHONY: all test clean