include ../.env
NAME := helloworld
VARS:=$(shell sed -ne 's/ *\#.*$$//; /./ s/=.*$$// p' ../.env)
GOPATH:=$(shell go env GOPATH)
PROTO_FLAGS=--go_opt=paths=source_relative --micro_opt=paths=source_relative
PROTO_PATH=..:.
SRC_DIR=..
$(foreach v,$(VARS),$(eval $(shell echo export $(v)="$($(v))")))
.PHONY: all
all: proto api build

.PHONY: proto
proto:
	find ../proto/ -name '*.proto' -exec protoc --proto_path=$(PROTO_PATH) $(PROTO_FLAGS) --micro_out=$(SRC_DIR) --go_out=:$(SRC_DIR) {} \;

.PHONY: api
api:
	protoc --openapi_out=. --proto_path=.. $(wildcard ../proto/$(NAME)/*.proto)
	mv api-$(NAME).json api.json

.PHONY: build
build:
	go build -o $(NAME) *.go

.PHONY: up
up:
	go run . server --name=helloworld

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t $(NAME):latest
