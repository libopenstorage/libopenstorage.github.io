# Tutorial: Python

In this tutorial you will learn to setup your environment, create a volume,
then take its snapshot, create cloud credentials, and create a backup of the
volume.

The sources for this tutorial are available in the [`examples/python`](https://github.com/libopenstorage/libopenstorage.github.io/tree/master/examples/python) directory in the repo for this website.

If you would like to try this tutorial on the web, check out our
[Katacoda Scenario: OpenStorage SDK Python Client Tutorial](https://www.katacoda.com/lpabon/scenarios/tutorial-python).

## Setting up your environment
To setup your environment, run the following:

```
$ python3 -m venv sdk
$ source sdk/bin/activate
$ pip3 install libopenstorage-openstorage
```

You will now have access to the OpenStorage SDK Python client which was
generated from the protobuf file [`api.proto`](https://github.com/libopenstorage/openstorage/blob/master/api/api.proto).

With the [`mock-sdk-server`](tutorial.html#setting-up-the-mock-sdk-server)
running, the following steps will provide an introduction to programming
with the OpenStorage SDK.

## Import dependencies
For most of the examples below, you will need the following `import`'s.

```python
import grpc
from openstorage import api_pb2
from openstorage import api_pb2_grpc
from openstorage import connector
```

## Creating a connection
To use any of the gRPC functions, you must first create a connection with
the OpenStorage SDK server:

```python
c = connector.Connector('localhost:9100')
```

>The `connector.Connector()` object supports not only insecure access,
>but also TLS, Token Authentication, and fetching tokens from Kubernetes secrets. For
>more information, type:
>
>   ```
>   $ pydoc3 openstorage.connector
>   ```
>

## Cluster operations
Now that we have made a connection, we can use the `channel` object to create
clients for each of the services we would like to use. Let's use the [OpenStorageCluster](generated-api.html#openstorageapiopenstoragecluster)
service to print the `id` of the cluster:

```python
try:
    # Cluster connection
    clusters = api_pb2_grpc.OpenStorageClusterStub(channel)
    ic_resp = clusters.InspectCurrent(api_pb2.SdkClusterInspectCurrentRequest())
    print('Connected to {0} with status {1}'.format(
        ic_resp.cluster.id,
        api_pb2.Status.Name(ic_resp.cluster.status)
    ))
except grpc.RpcError as e:
    print('Failed: args={0} msg={1}'.format(e.args, e.message))
```

Notice the `except` above. As mentioned in the
[Architecture](arch.html#error-handling) all errors are encoded using the
standard gRPC status. To gain access to the error code and its message you
must use `except grpc.RpcError as e` which decodes the error value and the message.

## Volume Operations
Now that we have connected to the cluster, let's go ahead and create a
volume of size 100Gi:

```python
    volumes = api_pb2_grpc.OpenStorageVolumeStub(channel)
    v_resp = volumes.Create(api_pb2.SdkVolumeCreateRequest(
        name="myvol",
        spec=api_pb2.VolumeSpec(
            size=100*1024*1024*1024,
            ha_level=3,
        )
    ))
    print('Volume id is {0}'.format(v_resp.volume_id))
```

Notice the value of `name` provided above. This value is important since
it allows for the function to be idempotent. In other words, this function
will always return the same volume id for the volume of same name.

You can now create a snapshot of this volume:

```python
    # Create a snapshot
    snap = volumes.SnapshotCreate(api_pb2.SdkVolumeSnapshotCreateRequest(
        volume_id=v_resp.volume_id,
        name="mysnap"
    ))
    print('Snapshot created with id {0}'.format(snap.snapshot_id))
```

## Backup
In this section we will be making a backup of our volume to a cloud provider.
To do this, we must first create a set of credentials which will enable
the storage system to save the backup in the cloud.

```python
    # Create a credentials
    creds = api_pb2_grpc.OpenStorageCredentialsStub(channel)
    cred_resp = creds.Create(api_pb2.SdkCredentialCreateRequest(
	    name='aws',
        aws_credential=api_pb2.SdkAwsCredentialRequest(
            access_key='dummy',
            secret_key='dummy',
            endpoint='dummy',
            region='dummy'
        )
    ))
    print('Credential id is {0}'.format(cred_resp.credential_id))
```

Notice above the [`api_pb2.SdkCredentialCreateRequest`](https://libopenstorage.github.io/w/generated-api.html#sdkcredentialcreaterequest).
This struct in protobuf uses [`oneof`](https://developers.google.com/protocol-buffers/docs/proto3#oneof).
Oneof states that _one of these types_ will be used. In SdkCredentialCreateRequest,
only `aws_credential` is set, showing the server which type of credentials
are being registered.

The backup of the volume can now be started with the newly acquired
credential id:

```python
    # Create backup
    backups = api_pb2_grpc.OpenStorageCloudBackupStub(channel)
    backup_resp = backups.Create(api_pb2.SdkCloudBackupCreateRequest(
        volume_id=v_resp.volume_id,
        credential_id=cred_resp.credential_id,
        full=False
    ))
```

This request will not block while the backup is running. Instead you should
call [OpenStorageCloudBackup.Status()](https://libopenstorage.github.io/w/generated-api.html#methodstatus)
to get information about the backup:

```python
    status_resp = backups.Status(api_pb2.SdkCloudBackupStatusRequest(
        volume_id=v_resp.volume_id
    ))
    backup_status = status_resp.statuses[v_resp.volume_id]
    print('Status of the backup is {0}'.format(
        api_pb2.SdkCloudBackupStatusType.Name(backup_status.status)
    ))
```

Lastly, once the backup is complete, we can get a history of this and any
other backups we have created from our volume:

```python
    # Show history
    history = backups.History(api_pb2.SdkCloudBackupHistoryRequest(
        src_volume_id=v_resp.volume_id
    ))
    print('Backup history for volume {0}'.format(v_resp.volume_id))
    for item in history.history_list:
        print('Time:{0} Status:{1}'.format(
            item.timestamp.ToJsonString(),
            api_pb2.SdkCloudBackupStatusType.Name(item.status)
        ))
```

## Example output
Below is an example output run of this tutorial:

```
Conntected to portworx with status STATUS_OK
Volume id is 9b64012f-db53-4a7b-8393-94b3d1ba0b02
Snapshot created with id 26a7262c-3eec-4352-a0be-3710d25f8335
Credential id is 716ed48f-f089-453d-ba07-7072eacd2ba0
Status of the backup is SdkCloudBackupStatusTypeDone
Backup history for volume 9b64012f-db53-4a7b-8393-94b3d1ba0b02
Time:2018-07-18T02:04:53.278579951Z Status:SdkCloudBackupStatusTypeDone
```

## Using the Python SDK Documentation

The python SDK generated documentation in the [API reference](reference.md)
provides some necessary information on the Python bindings for the OpenStorage
SDK, but it does not provide a full model. This is due to the python protobuf
generator. You will need to use both the gRPC and Python API references
to achieve this.

Here is an example model of how to to access the UUID of a cluster.

1. Go to the OpenStorageCluster service in the gRPC reference. Here you will notice
   that the _SdkClusterInspectCurrentResponse_ is returned with a field named
   `cluster` of type `StorageCluster`.
1. Now go to the Python API reference and open the `api_pb2` module. Here you
   can search for _SdkClusterInspectCurrentResponse_ and notice that it returns
   a `var cluster`. This is the gRPC `field` named `cluster` in the gRPC
   reference, but the python document does not provide the type of the variable
   so that you can determine its instance variables. As you noticed on the
   previous step, `cluster` is of type `StorageCluster`.
1. In the same `api_pb2` python document, look up `StorageCluster` class. You
   will notice that it shows three instance variables: `id`, `name`, and
   `status`.
1. `cluster.id` and `cluster.name` are `strings` as shown in the gRPC
   documentation, but `cluster.status` is an enum of type `Status`.
1. If you print the value from `cluster.status` you will get the `int` value of
   the enum, but if you want the string value of the enum, protobuf provides a
   conversion funcion called: `api_pb2.<enum>.Name(value)`. You can print the
   string value of the enum `cluster.status` by calling
   `api_pb2.Status.Name(cluster.status)`.

## Next
As you can see from the above, working with OpenStorage SDK is quite easy,
fun, and powerful. Please refer to the [API Reference](reference.md)
for a complete list of services.
