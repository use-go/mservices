package template

var (
	Makefile = `include ../.env
NAME := {{.Alias}}
VARS:=$(shell sed -ne 's/ *\#.*$$//; /./ s/=.*$$// p' ../.env)
GOPATH:=$(shell go env GOPATH)
FILES:=$(wildcard ../proto/$(NAME)/*.proto)
$(foreach v,$(VARS),$(eval $(shell echo export $(v)="$($(v))")))
.PHONY: init
init:
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get github.com/micro/micro/v3/cmd/protoc-gen-micro
	go get github.com/micro/micro/v3/cmd/protoc-gen-openapi

.PHONY: api
api:
	protoc --openapi_out=. --proto_path=.. $(FILES)
	mv api-$(NAME).json api.json

.PHONY: proto
proto:
	$(foreach file,$(FILES),$(eval $(shell protoc --proto_path=.. --micro_out=../proto/$(NAME) --go_out=:../proto/$(NAME) $(file))))
	mv ../proto/$(NAME)/proto/* ../proto/$(NAME)
	rm ../proto/$(NAME)/proto -rf

.PHONY: build
build:
	go build -o $(NAME) *.go

.PHONY: up
up:
	go run . server

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t $(NAME):latest	
`
)
