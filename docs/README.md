# OpenStorage Library API


## Introduction
Welcome to the _OpenStorage Library API_ documentation. Here you will find the
API functions, examples, and references to multiple client libraries which can
communicate with a OpenStorage REST server.

## Quick Example
Here is a super quick example of how to get cluster information from an OpenStorage
server using a mock container running the gRPC server and using [Polyglot](https://github.com/grpc-ecosystem/polyglot) as the gRPC client.

```
$ docker run -d -p 9100:9100 openstorage/mock-sdk-server
$ wget https://github.com/grpc-ecosystem/polyglot/releases/download/v1.6.0/polyglot.jar
$ echo {} | java -jar polyglot.jar \
  --command=call \
  --endpoint=localhost:9100 \
  --full_method=openstorage.api.OpenStorageCluster/Enumerate
```

Results in:

```yaml
{ 
  "cluster": {
    "status": "STATUS_OK",
    "id": "deadbeeef",
    "nodeId": "1",
    "nodes": [{
      "id": "1",
      "cpu": 0.4993757802746567,
      "memTotal": "8320278528",
      "memUsed": "1895813120",
      "memFree": "6424465408",
      "status": "STATUS_OK",
      "mgmtIp": "172.17.0.2",
      "dataIp": "172.17.0.2",
      "hostname": "42f9d528292c"
    }]
  }
}
```
