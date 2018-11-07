# Tutorial: Golang

In this tutorial you will learn to setup your environment, create a volume,
then take its snapshot, create cloud credentials, and create a backup of the
volume.

The sources for this tutorial are available in the [`examples/golang`](https://github.com/libopenstorage/libopenstorage.github.io/tree/master/examples/golang) directory in the repo for this website.

#### Recorded demo
Check out the [step by step tutorial](https://asciinema.org/a/auGZOUd4R5osIgS8xEbRUX8Ah)
on how to setup the golang client using [Golang dep](https://github.com/golang/dep).

## Setting up your environment
To setup your environment, you will need to import the following to gain
access to `api.pb.go`:

```go
import (
    api "github.com/libopenstorage/openstorage-sdk-clients/sdk/golang"
    "github.com/golang/protobuf/ptypes"
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
	conn, err := grpc.Dial("localhost:9100", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
```

> **NOTE**: Notice the call [`grpc.WithInsecure()`](https://grpc.io/docs/guides/auth.html). The mock-sdk-server
currently does not support HTTPS or authentication, so the tutorial will
only focus on an insecure connection. The SDK project will add HTTPS and
authentication support in following releases.

## Cluster operations
Now that we have made a connection, we can use the `conn` object to create
clients for each of the services we would like to use. Let's use the [OpenStorageCluster](generated-api.html#openstorageapiopenstoragecluster)
service to print the `id` of the cluster:

```go
	// Create a cluster client
	cluster := api.NewOpenStorageClusterClient(conn)

	// Print the cluster information
	clusterInfo, err := cluster.InspectCurrent(
		context.Background(),
		&api.SdkClusterInspectCurrentRequest{})
	if err != nil {
		gerr, _ := status.FromError(err)
		fmt.Printf("Error Code[%d] Message[%s]\n",
			gerr.Code(), gerr.Message())
		os.Exit(1)
	}
	fmt.Printf("Connected to Cluster %s\n",
		clusterInfo.GetCluster().GetId())
```

Notice the `status.FromError()` call above. As mentioned in the
[Architecture](arch.html#error-handling) all errors are encoded using the
standard gRPC status. To gain access to the error code and its message you
must use `status.FromError()` which decodes the error value and the message.

Notice also the use of accessors `GetXXX()` above. These convenient functions
are provided automatically by the golang protobuf generator.

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
				Size:    100 * 1024 * 1024 * 1024,
				HaLevel: 3,
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
			Name:     fmt.Sprintf("snap-%v", time.Now().Unix()),
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

## Backup
In this section we will be making a backup of our volume to a cloud provider.
To do this, we must first create a set of credentials which will enable
the storage system to save the backup in the cloud.

```go
	// Create mock credentials
	creds := api.NewOpenStorageCredentialsClient(conn)
	credResponse, err := creds.Create(context.Background(),
		&api.SdkCredentialCreateRequest{
			Name: "aws",
			CredentialType: &api.SdkCredentialCreateRequest_AwsCredential{
				AwsCredential: &api.SdkAwsCredentialRequest{
					AccessKey: "dummy-access",
					SecretKey: "dummy-secret",
					Endpoint:  "dummy-endpoint",
					Region:    "dummy-region",
				},
			},
		})
	if err != nil {
		gerr, _ := status.FromError(err)
		fmt.Printf("Error Code[%d] Message[%s]\n",
			gerr.Code(), gerr.Message())
		os.Exit(1)
	}
	credID := credResponse.GetCredentialId()
	fmt.Printf("Credentials created with id %s\n", credID)
```

Notice above the `CredentialType.AwsCredential`. This is struct was
generated from the protobuf concept of [`oneof`](https://developers.google.com/protocol-buffers/docs/proto3#oneof). Oneof
states that _one of these types_ will be used. In SdkCredentialCreateRequest, `oneof` is used as a method of determining which
type of cloud provider is being requested.

The backup of the volume can now be started with the newly acquired
credential id:

```go
	// Create a backup to a cloud provider of our volume
	cloudbackups := api.NewOpenStorageCloudBackupClient(conn)
	backupCreateResp, err := cloudbackups.Create(context.Background(),
		&api.SdkCloudBackupCreateRequest{
			VolumeId:     v.GetVolumeId(),
			CredentialId: credID,
		})
	if err != nil {
		gerr, _ := status.FromError(err)
		fmt.Printf("Error Code[%d] Message[%s]\n",
			gerr.Code(), gerr.Message())
		os.Exit(1)
	}
	fmt.Printf("Backup started for volume %s with task id %s\n",
		v.GetVolumeId(),
		backupCreateResp.GetTaskId())
```

This request will not block while the backup is running. Instead you should
call OpenStorageCloudBackup.Status() to get information about the backup:

```go
	// Now check the status of the backup
	backupStatus, err := cloudbackups.Status(context.Background(),
		&api.SdkCloudBackupStatusRequest{
			VolumeId: v.GetVolumeId(),
		})
	if err != nil {
		gerr, _ := status.FromError(err)
		fmt.Printf("Error Code[%d] Message[%s]\n",
			gerr.Code(), gerr.Message())
		os.Exit(1)
	}
	for volID, status := range backupStatus.GetStatuses() {
		// There will be only one value in the map, but we use
		// a for-loop as an example.
		b, _ := json.MarshalIndent(status, "", "  ")
		fmt.Printf("Backup status for taskId: %s\n"+
			"Volume: %s\n"+
			"Type: %s\n"+
			"Status: %s\n"+
			"Full JSON Response: %s\n",
			taskId,
			status.GetSrcVolumeId(),
			status.GetOptype().String(),
			status.GetStatus().String(),
			string(b))
	}
```

Lastly, once the backup is complete, we can get a history of this and any
other backups we have created from our volume:

```go
	historyResp, err := cloudbackups.History(context.Background(),
		&api.SdkCloudBackupHistoryRequest{
			SrcVolumeId: v.GetVolumeId(),
		})
	if err != nil {
		gerr, _ := status.FromError(err)
		fmt.Printf("Error Code[%d] Message[%s]\n",
			gerr.Code(), gerr.Message())
		os.Exit(1)
	}

	fmt.Printf("Backup history for volume %s:\n", v.GetVolumeId())
	for _, history := range historyResp.GetHistoryList() {

		timestamp, _ := ptypes.Timestamp(history.GetTimestamp())
		fmt.Printf("Volume:%s \tttime:%v \tstatus:%v\n",
			history.GetSrcVolumeId(),
			timestamp,
			history.GetStatus())
	}
```

## Example output
Below is an example output run of this tutorial:

```
Connected to Cluster mock

Volume 100Gi created with id 42908221-1e24-4927-b67a-468c8b826b0b

Snapshot with id e24efabc-a2d8-4f11-aa4a-a283623536ae was create for volume 42908221-1e24-4927-b67a-468c8b826b0b

Credentials created with id 3c73dc9f-dce7-4852-b870-79a988ac4440

Backup started for volume 42908221-1e24-4927-b67a-468c8b826b0b
Backup status for volume: 42908221-1e24-4927-b67a-468c8b826b0b
Type: SdkCloudBackupOpTypeBackupOp
Status: SdkCloudBackupStatusTypeDone
Full JSON Response: {
  "backup_id": "ae23bc53-79b9-49a6-9eee-fac66221528e",
  "optype": 1,
  "status": 2,
  "bytes_done": 107374182400,
  "start_time": {
    "seconds": 1531438150,
    "nanos": 856373600
  },
  "completed_time": {
    "seconds": 1531438151,
    "nanos": 856397100
  },
  "node_id": "1"
} 

Backup history for volume 42908221-1e24-4927-b67a-468c8b826b0b:
Volume:42908221-1e24-4927-b67a-468c8b826b0b     ttime:2018-07-12 23:29:11.8563971 +0000 UTC     status:SdkCloudBackupStatusTypeDone
```

## Next
As you can see from the above, working with OpenStorage SDK is quite easy,
fun, and powerful. Please refer to the [API Reference](reference.md)
for a complete list of services.
