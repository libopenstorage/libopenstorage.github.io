# Tutorial

This section contains simple OpenStorage SDK tutorials for multiple
languages. Each tutorial will provide an introduction to programming
your OpenStorage system.

## Setting up the mock-sdk-server
For simplicity these tutorials will use the `mock-sdk-server` provided by the
OpenStorage project. It is an in-memory storage system implementation of the
OpenStorage SDK and it makes it possible to develop your client by simply
running a container on your laptop or system.

Run the following command to start the server:

```
$ docker run --rm --name sdk -d -p 9100:9100 -p 9110:9110 openstorage/mock-sdk-server
```

Running this docker command will expose port `9100` for the gRPC server, and
port `9110` for the gRPC REST Gateway connection. As shown in the [Architecture](arch.md),
the gRPC REST Gateway translates simple REST requests to gRPC requets.

This document assumes that you will be running the container on your system,
therefore it will refer to the connetion as `localhost`. If you have the container
running on another system, you will need to change the server location
accordingly.

Let's test the connection now before continuing by running the following to the
gRPC REST Gateway:

```
$ curl -X GET "http://localhost:9110/v1/clusters/current" \
     -H "accept: application/json" \
	 -H "Content-Type: application/json" \
	 -d "{}"
```

This should return information about the cluster in JSON format, for example:

```json
{"cluster":{"status":"STATUS_OK","id":"mock"}}
```

## Setting up security in the mock-sdk-server
This section will help you enable security in the _mock-sdk-server_.

### Enabling TLS
You will need the server _crt_ and _key_ files to provide them to the
mock-sdk-server. Since the mock-sdk-server is running in a container, you
will need to map into the container the directory where the certificates are
located. Here is an example:

```
$ ls ./certs/
server.crt server.key
$ docker run --rm --name sdk -d -p 9100:9100 -p 9110:9110 \
  -v $PWD/certs:/certs \
  openstorage/mock-sdk-server \
  -d --driver=name=fake \
  --tls-cert-file=/certs/server.crt \
  --tls-key-file=/certs/server.key
```

> NOTE: For development simplicity, cert files are provided under the
  [`examples/unsecure_certs`](https://github.com/libopenstorage/libopenstorage.github.io/tree/master/examples/unsecure_certs)

### Enabling Authorization Using A Shared Key
The simplest method to enable authorization is to use the shared secret key
model. To enable the server authorization with a shared secret, you will need
to pass the secret using the command line. In the following example, the secret
`mysecret` is passed to the server as the shared secret:

```
$ docker run --rm --name sdk -d -p 9100:9100 -p 9110:9110 \
  openstorage/mock-sdk-server \
  -d --driver=name=fake \
  --jwt-shared-secret=mysecret
```

RSA and ECDSA authorization using public keys are also supported. Check
out the help by typing:

```
$ docker run --rm --name sdk -d -p 9100:9100 -p 9110:9110 \
  openstorage/mock-sdk-server \
  --help
```