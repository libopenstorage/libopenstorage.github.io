# OpenStorage SDK


## Introduction
Welcome to the _OpenStorage SDK_ documentation. Here you will find the
API functions, examples, and references to multiple client libraries which can
communicate with an OpenStorage server.

## Quick Example
Here is a quick example of how to get cluster information from an OpenStorage
server. In this example we will be using the mock SDK container as the
server. First let's start the container by running the following command:

```
$ docker run --rm --name sdk -d -p 9100:9100 -p 9110:9110 openstorage/mock-sdk-server
```

This container starts a tiny, in-memory OpenStorage SDK server able to be used
for development. More information can be found in the [Tutorial](tutorial.html).

Now we can send requests from the command line. To send requests to the gRPC
server, we will use [Polyglot 1.x.x](https://github.com/grpc-ecosystem/polyglot/releases)
as the gRPC client. To send requests to the gRPC REST Gateway, we will be
sending the same request as Polyglot, but using _curl_ instead:

{% codetabs name="gRPC", type="less" -%}
$ wget https://github.com/grpc-ecosystem/polyglot/releases/download/v1.6.0/polyglot.jar
$ echo {} | java -jar polyglot.jar \
  --command=call \
  --endpoint=localhost:9100 \
  --full_method=openstorage.api.OpenStorageCluster/InspectCurrent
{%- language name="REST", type="less" -%}
$ curl -X GET "http://localhost:9110/v1/clusters/inspectcurrent" \
     -H "accept: application/json" \
	 -H "Content-Type: application/json" \
	 -d "{}"
{%- endcodetabs %}

Results in:

```json
{
  "cluster": {
    "status": "STATUS_OK",
    "id": "mock"
  }
}
```

The `mock-sdk-server` container also comes with a full [Swagger](https://swagger.io)
UI located at [http://localhost:9110/swagger-ui/](http://localhost:9110/swagger-ui/).

## What's next

Check out:

* [Architecture](arch.html)
* [Installation](installation.html)
* [Tutorials](tutorial.html)


