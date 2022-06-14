module proto

go 1.17

require (
	github.com/2637309949/micro/v3 v3.8.6
	github.com/golang/protobuf v1.5.2
	google.golang.org/protobuf v1.27.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/evanphx/json-patch/v5 v5.0.0 // indirect
	github.com/go-acme/lego/v3 v3.4.0 // indirect
	github.com/google/uuid v1.1.2 // indirect
	github.com/oxtoacart/bpool v0.0.0-20190530202638-03653db5a59c // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	golang.org/x/net v0.0.0-20211123203042-d83791d6bcd9 // indirect
	golang.org/x/text v0.3.6 // indirect
	google.golang.org/grpc v1.40.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.27.1
