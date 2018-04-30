# Architecture

![arch.png](images/arch.png)

## Overview
The software development kit for OpenStorage is based on [gRPC](https://grpc.io/),
allowing clients to be automatically generated in multiple languages.

Different OpenStorage drivers may run the OpenStorage gRPC server on different
locations. Please, refer to your deployed driver documentation for instructions
on how to setup a connection to the gRPC server.

In this document, examples will refer to both the [OpenStorage SDK Mock](installing.md)
and Portworx drivers.

## Clients
Currently, OpenStorage provides Python, Ruby, Golang, and JavaScript bindings
under the [`api/client/sdk`](https://github.com/libopenstorage/openstorage/tree/master/api/client/sdk)
directory. If your client is not available, please feel free to generate one
from the [`api.proto`](https://github.com/libopenstorage/openstorage/blob/master/api/api.proto)
file on OpenStorage.
