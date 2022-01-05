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
curl http://192.168.202.128:8080/file/upload
level=info trace= Starting [service] assert
level=info trace= HTTP API Listening on [::]:43227
level=info trace= Registry [service] Registering node: assert-350d533e-8477-42f6-aab1-bcf46a9e06ae
level=info trace=9659111d-68eb-4de3-b5ce-cadf46e89f70 >>>>> Received /assert/file/upload request = {}
level=info trace=9659111d-68eb-4de3-b5ce-cadf46e89f70 <<<<< Finished /assert/file/upload 85ms response = {"name":"hello2012"}
level=info trace=d63d0671-6a0b-4708-bf33-8428397fd18c >>>>> Received /assert/file/upload request = {}
level=info trace=d63d0671-6a0b-4708-bf33-8428397fd18c <<<<< Finished /assert/file/upload 96ms response = {"name":"hello2012"}
level=info trace=124bfb05-ac6c-4a4f-9b75-0e8ad1a87016 >>>>> Received /assert/file/upload request = {}
level=info trace=124bfb05-ac6c-4a4f-9b75-0e8ad1a87016 <<<<< Finished /assert/file/upload 146ms response = {"name":"hello2012"}
```