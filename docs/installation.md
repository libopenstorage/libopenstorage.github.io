# Installing OpenStorage SDK

## Clients
Currently, OpenStorage provides Python, Ruby, Golang, and JavaScript bindings
under the [`api/client/sdk`](https://github.com/libopenstorage/openstorage/tree/master/api/client/sdk)
directory. If your client is not available, please feel free to generate one
from the [`api.proto`](https://github.com/libopenstorage/openstorage/blob/master/api/api.proto)
file on OpenStorage.

### Golang
For Golang, simply import `github.com/libopenstorage/openstorage/api` in your
sources. For a full example see [Tutorial: Golang](tutorial-golang.html).

### Python
Python libraries are available [`api/client/sdk/python`](https://github.com/libopenstorage/openstorage/tree/master/api/client/sdk/python).
You will need the files `api*.py` and `google/` copied to your path. We may
provide a python `pip` package, but do not have one at the moment.

### Ruby and Javascript
Ruby and Javascript clients are also located in [`api/client/sdk`](https://github.com/libopenstorage/openstorage/tree/master/api/client/sdk/python). 