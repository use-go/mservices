## Usage

### Setting config
```shell
micro config set helloworld.db.spamd_address 127.0.0.1:783
micro config set helloworld.db.addr smtp.qq.com:25
micro config set helloworld.db.username 2637309949@qq.com
micro config set helloworld.db.identity 
micro config set helloworld.db.password jhprqpetmlfteabe
micro config set helloworld.db.host smtp.qq.com
```

Generate the proto code

```
make proto
```

Run the service

```
make up
```

Request the service

```
[double@double] curl http://81.71.122.245:8080/helloworld/helloworld/UpdateInfo
level=info Starting [service] helloworld
level=info Server [grpc] Listening on [::]:35653
level=info Registry [service] Registering node: helloworld-964a58bf-ffeb-47c2-a4f5-a18491c7e644
level=info >>>>> Received Helloworld.UpdateInfo request = {"operator_id":"admin","operator_name":"admin"}
level=info <<<<< Finished Helloworld.UpdateInfo 110ms response = {"name":"Hello"}
level=info >>>>> Received Helloworld.UpdateInfo request = {"operator_id":"admin","operator_name":"admin"}
level=info <<<<< Finished Helloworld.UpdateInfo 93ms response = {"name":"Hello"}
```