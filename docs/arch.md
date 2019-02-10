# Architecture

![arch.png](images/arch.png)

## Overview
The software development kit for OpenStorage is based on [gRPC](https://grpc.io/),
allowing clients to be automatically generated in multiple languages. The SDK
also provides a gRPC REST Gateway which automatically changes REST requests
to gRPC.

Different OpenStorage drivers may run the OpenStorage gRPC and REST servers on
different ports. Please, refer to your deployed driver documentation for
instructions on how to setup a connection to the gRPC server.

OpenStorage API calls are ment to be idempotent to ensure consistency across calls, unless
otherwise specified.

## Protocol Buffers gRPC Source File
The gRPC bindings are created from [`api/api.proto`](https://github.com/libopenstorage/openstorage/blob/master/api/api.proto) file
available in the OpenStorage github repo.

## REST gRPC Gateway
REST clients can benefit from the SDK server's [gRPC REST Gateway](https://github.com/grpc-ecosystem/grpc-gateway)
which is able to translate REST calls to gRPC requests.

[Swagger](https://swagger.io/) is also provided by going to the REST gRPC Gateway
and passing the path `/swagger-ui`.

## Security
Starting at _v0.38.0_, the OpenStorage SDK supports security based on the
following set of technologies:

* **TLS**: The SDK now supports TLS to secure client server communications.
* **Authentication**: Support for authentication based on JWT and OpenID Connect(OIDC)
* **Authorization**: Support for role based access control (RBAC) access to
  OpenStorage SDK API calls
* **Ownership**: Support for resource access control based on owner, groups,
  and collaborators.

## Error Handling
All API calls use the [standard gRPC status](https://github.com/grpc/grpc/blob/master/src/proto/grpc/status/status.proto).

## OpenStorage SDK Implementations
In this document, examples will refer to both the [OpenStorage SDK Mock](tutorial.html)
and Portworx OpenStorage drivers.

