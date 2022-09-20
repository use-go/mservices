### Installation

#### Installing protoc

```shell
PROTOC_ZIP=protoc-3.14.0-linux-x86_64.zip
curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.14.0/$PROTOC_ZIP
sudo unzip -o $PROTOC_ZIP -d /usr/local bin/protoc
sudo unzip -o $PROTOC_ZIP -d /usr/local 'include/*'
rm -f $PROTOC_ZIP
```

#### Installing micro

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
wget -q  https://raw.githubusercontent.com/2637309949/micro/master/scripts/install.sh -O - | /bin/bash
```

#### Run micro server

```shell
micro server
```

[Multi-cluster Deployment](http://hbchen.com/post/microservice/2019-11-15-go-micro-network/)

### Usage

```shell
micro new test && cd test-service && make proto && make up
```

### VerifyAccess

```shell
micro auth create rule --access=granted --scope='*' --resource="*:*:*" onlyloggedin
micro auth create rule --access=granted --resource="service:auth:*" auth
micro auth create rule --access=granted --resource="service:micro.:*" micro
micro auth create rule --access=granted --resource="service:assert:*" assert
micro auth create rule --access=granted --resource="service:quicktype:Quicktype.Call" quicktype
micro auth create rule --access=granted --resource="service:cas:*" cas
micro auth delete rule default
```
