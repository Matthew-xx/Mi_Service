
GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/PostUserAuth/PostUserAuth.proto

.PHONY: build
build: proto

	go build -o PostUserAuth-srv *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t PostUserAuth-srv:latest
