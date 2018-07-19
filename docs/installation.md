# Installing OpenStorage SDK

## Clients
Currently, OpenStorage provides Python, Ruby, Golang, and JavaScript bindings
in the [github.com/libopenstorage/openstorage-sdk-clients](https://github.com/libopenstorage/openstorage-sdk-clients)
repo. If your client is not available, please feel free to generate one
from the [`api.proto`](https://github.com/libopenstorage/openstorage/blob/master/api/api.proto)
file on OpenStorage.

### Golang
For Golang, simply add:

```go
import (
    api "github.com/libopenstorage/openstorage-sdk-clients/sdk/golang"
    "github.com/golang/protobuf/ptypes"
    "google.golang.org/grpc"
    "google.golang.org/grpc/status"
)
```
in your imports. For a full example see [Tutorial: Golang](tutorial-golang.html).

### Python
Python libraries are available [`sdk/python`](https://github.com/libopenstorage/openstorage-sdk-clients/tree/master/sdk/python).
You will need the files `api*.py` and `google/` copied to your path. We may
provide a python `pip` package, but do not have one at the moment.

### Ruby and Javascript
Ruby and Javascript clients are also located in the
[`sdk`](https://github.com/libopenstorage/openstorage-sdk-clients/tree/master/sdk) directory
in `openstorage-sdk-clients` repo. 