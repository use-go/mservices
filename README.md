#### Install

Install micro with the following commands:

```shell
go install github.com/2637309949/go-service/comm/micro
```

#### Usage

```shell
micro new test && cd test-service && make proto && make up
```


#### Known Service

```
registry // :8000
broker   // :8003
network  // :8443
runtime  // :8088
config   // :8001
store    // :8002
events   // :unset
auth     // :8010
proxy    // :8081
api      // :8080
web      // :8082
```