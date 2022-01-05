## Usage

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
curl http://192.168.202.128:8080/helloworld/helloworld/UpdateInfo
level=info Starting [service] helloworld
level=info Server [grpc] Listening on [::]:35653
level=info Registry [service] Registering node: helloworld-964a58bf-ffeb-47c2-a4f5-a18491c7e644
level=info trace=de9d3f7b-e842-4bfd-bbb0-a8ff695db366 >>>>> Received Helloworld.UpdateInfo request = {"operator_id":"admin","operator_name":"admin"}
level=info trace=de9d3f7b-e842-4bfd-bbb0-a8ff695db366 <<<<< Finished Helloworld.UpdateInfo 110ms response = {"name":"Hello"}
level=info trace=ff8498a5-261c-41f2-af4a-25258e0b2fd1 >>>>> Received Helloworld.UpdateInfo request = {"operator_id":"admin","operator_name":"admin"}
level=info trace=ff8498a5-261c-41f2-af4a-25258e0b2fd1 <<<<< Finished Helloworld.UpdateInfo 93ms response = {"name":"Hello"}
```