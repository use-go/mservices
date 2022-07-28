#### Install

Install micro with the following commands:

```shell
go clone git@github.com:2637309949/micro.git
go install
```

Run micro server with the following commands:

```shell
micro server
```
#### Usage

```shell
micro new test && cd test-service && make proto && make up
```

[Multi-cluster Deployment](http://hbchen.com/post/microservice/2019-11-15-go-micro-network/)

#### VerifyAccess

```shell
micro auth delete rule default
micro auth create rule --access=granted --scope='*' --resource="*:*:*" onlyloggedin
micro auth create rule --access=granted --resource="service:auth:*" auth-public
micro auth create rule --access=granted --resource="service:micro.:*" micro-public
micro auth create rule --access=granted --resource="service:assert:*" assert-public
micro auth create rule --access=granted --resource="service:quicktype:Quicktype.Call" quicktype-public
micro auth create rule --access=granted --resource="service:cas:*" cas-public
```

#### Service List

- assert-service
- cache-service
- cas-service
- email-service
- helloworld-service
- id-service
- quicktype-service
- screenshot-service
- subscribe-service
- user-service
