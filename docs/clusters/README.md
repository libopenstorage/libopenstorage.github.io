# Clusters

These set of APIs manage cluster information.

## Overview
The cluster API is used to interact with the cluster as whole. You will need to create a cluster client to be able to interact with the server and utilize the functions described in this chapter.

### Creating a client
Below shows examples of how to create a client in multiple languages:

{% codetabs name="Golang", type="go" -%}
import (
    clusterclient "github.com/libopenstorage/openstorage/api/client/cluster"
	"github.com/libopenstorage/openstorage/cluster"
)

// Create a client setup to talk to a specific node URL
c, err := clusterclient.NewClusterClient(url, "v1")
if err != nil {
    return err
}

// Create a manager which is an implementation of the
// cluster.Cluster interface
m := clusterclient.ClusterManager(c)

// Example - Enumerate
resp, err := m.Enumerate()

{%- language name="Python", type="py" -%}
TBD
{%- endcodetabs %}


## Reference

* [Golang](https://godoc.org/github.com/libopenstorage/openstorage/cluster#Cluster)