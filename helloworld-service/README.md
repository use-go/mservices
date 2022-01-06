### Usage

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
[double@double] curl http://192.168.202.128:8080/helloworld/helloworld/UpdateInfo
level=info Starting [service] helloworld
level=info Server [grpc] Listening on [::]:35653
level=info Registry [service] Registering node: helloworld-964a58bf-ffeb-47c2-a4f5-a18491c7e644
level=info >>>>> Received Helloworld.UpdateInfo request = {"operator_id":"admin","operator_name":"admin"}
level=info <<<<< Finished Helloworld.UpdateInfo 110ms response = {"name":"Hello"}
level=info >>>>> Received Helloworld.UpdateInfo request = {"operator_id":"admin","operator_name":"admin"}
level=info <<<<< Finished Helloworld.UpdateInfo 93ms response = {"name":"Hello"}
```