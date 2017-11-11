MAIN=./cmd/tmng/main.go

build:
	-go build -v -i $(MAIN)

deps:
	-dep ensure
