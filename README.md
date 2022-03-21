#### Install

Install micro with the following commands:

```shell
go install github.com/2637309949/micro/v3@v3.8.2
```

Run micro server with the following commands:

```shell
micro server
```

```shell
go get github.com/aoldershaw/proclimit/cmd/proclimit
```

#### Usage

```shell
micro new test && cd test-service && make proto && make up
```

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
