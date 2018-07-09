# Tutorial: Golang

In this tutorial you will learn to setup your environment, create a volume,
then take its snapshot.

## Setting up your environment
To setup your environment, you will need to import the following to gain
access to `api.pb.go`:

```go
import (
    "github.com/libopenstorage/openstorage/api"

    "google.golang.org/grpc"
	"google.golang.org/grpc/status"
)
```

Once you have imported this library into your project, you will have access
to the OpenStorage SDK Golang client which was generated from the protobuf
file [`api.proto`](https://github.com/libopenstorage/openstorage/blob/master/api/api.proto).

With the [`mock-sdk-server`](tutorial.html#setting-up-the-mock-sdk-server)
running, the following steps will provide an introduction to programming
with the OpenStorage SDK.

## Creating a connection
To use any of the gRPC functions, you must first create a connection with
the OpenStorage SDK server:

```go
	// Setup a connection
	conn, err := grpc.Dial(":9100", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
```

## Cluster operations
Now that we have made a connection, we can use the `conn` object to create
clients for each of the services we would like to use. Let's use the [OpenStorageCluster](generated-api.html#openstorageapiopenstoragecluster)
service to print the `id` of the cluster and how many nodes it has:

```go
	// Create a cluster client
	cluster := api.NewOpenStorageClusterClient(conn)

	// Print the cluster information
	clusterEnum, err := cluster.Enumerate(
		context.Background(),
		&api.SdkClusterEnumerateRequest{})
	if err != nil {
		gerr, _ := status.FromError(err)
		fmt.Printf("Error Code[%d] Message[%s]\n",
			gerr.Code(), gerr.Message())
		os.Exit(1)
	}
	fmt.Printf("Connected to Cluster %s with %d node(s)\n",
		clusterEnum.GetCluster().GetId(),
		len(clusterEnum.GetCluster().GetNodeIds()))
```

Notice the `status.FromError()` call above. As mentioned in the
[Architecture](arch.html#error-handling) all errors are encoded the standard
gRPC status. To gain access to the error code and its message you must
use `status.FromError()` which decodes the error value and the message.

Notice also the use of accessors `GetXXX()` above. These convenient functions
are provided automatically be the golang protobuf generator.

## Volume Operations
Now that we have connected to the cluster, let's go ahead and create a
volume of size 100Gi:

```go
	// Create a 100Gi volume
	volumes := api.NewOpenStorageVolumeClient(conn)
	v, err := volumes.Create(
		context.Background(),
		&api.SdkVolumeCreateRequest{
			Name: "myvol",
			Spec: &api.VolumeSpec{
				Size: 100 * 1024 * 1024 * 1024,
			},
		})
	if err != nil {
		gerr, _ := status.FromError(err)
		fmt.Printf("Error Code[%d] Message[%s]\n",
			gerr.Code(), gerr.Message())
		os.Exit(1)
	}
	fmt.Printf("Volume 100Gi created with id %s\n", v.GetVolumeId())
```

Notice the value of `Name` provided above. This value is important since
it allows for the function to be idempotent. In other words, this function
will always return the same volume id for the volume of same name.

You can now create a snapshot of this volume:

```go
	snap, err := volumes.SnapshotCreate(
		context.Background(),
		&api.SdkVolumeSnapshotCreateRequest{
			VolumeId: v.GetVolumeId(),
		},
	)
	if err != nil {
		gerr, _ := status.FromError(err)
		fmt.Printf("Error Code[%d] Message[%s]\n",
			gerr.Code(), gerr.Message())
		os.Exit(1)
	}
	fmt.Printf("Snapshot with id %s was create for volume %s\n",
		snap.GetSnapshotId(),
		v.GetVolumeId())
```

As you can see from the above, working with OpenStorage SDK is quite easy,
fun, and powerful. Please refer to the [API Reference](generated-api.html)
for a complete list of services.
