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
[double@double] curl http://192.168.202.128:8080/file/upload
level=info trace= Starting [service] assert
level=info trace= HTTP API Listening on [::]:43227
level=info trace= Registry [service] Registering node: assert-350d533e-8477-42f6-aab1-bcf46a9e06ae
level=info trace=9659111d-68eb-4de3-b5ce-cadf46e89f70 >>>>> Received /assert/file/upload request = {}
level=info trace=058ec4c4-5741-43e6-9f13-cc6960dfd02e admin Do FileUpload
level=info trace=9659111d-68eb-4de3-b5ce-cadf46e89f70 <<<<< Finished /assert/file/upload 85ms response = {"name":"hello2012"}
```