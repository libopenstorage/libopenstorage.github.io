# OpenStorage API


## Introduction
Welcome to the _OpenStorage API_ documentation. Here you will find the
API functions, examples, and references to multiple client libraries which can
communicate with an OpenStorage server.

## Quick Example
Here is a quick example of how to get cluster information from an OpenStorage
server. In this example we will be using the mock SDK container as the
server:

* Start up the container

```
$ docker run -d -p 9100:9100 -p 9110:9110 openstorage/mock-sdk-server
```

* Running the gRPC server and using [Polyglot](https://github.com/grpc-ecosystem/polyglot) as the gRPC client.

```
$ wget https://github.com/grpc-ecosystem/polyglot/releases/download/v1.6.0/polyglot.jar
$ echo {} | java -jar polyglot.jar \
  --command=call \
  --endpoint=localhost:9100 \
  --full_method=openstorage.api.OpenStorageCluster/Enumerate
```

* Run using the REST gRPC Gateway

```
$ curl -X GET "http://localhost:9110/v1/cluster" \
     -H "accept: application/json" \
	 -H "Content-Type: application/json" \
	 -d "{}"
```

Results in:

```json
{
  "cluster": {
    "status": "STATUS_OK",
    "id": "deadbeeef",
    "node_id": "1",
    "node_ids": [
      "1"
    ]
  }
}
```

* Check out more REST information by looking at the [Swagger](https://swagger.io) located
at [http://localhost:9110/swagger-ui/](http://localhost:9110/swagger-ui/)

