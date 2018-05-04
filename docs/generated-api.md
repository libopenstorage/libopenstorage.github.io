# API Reference

# Table of Contents


- Services
    - [OpenStorageCluster](#openstorageapiopenstoragecluster)
    - [OpenStorageVolume](#openstorageapiopenstoragevolume)
  


- Messages
    - [ActiveRequest](#activerequest)
    - [ActiveRequest.ReqestKVEntry](#activerequestreqestkventry)
    - [ActiveRequests](#activerequests)
    - [Alert](#alert)
    - [Alerts](#alerts)
    - [ClusterAlertClearRequest](#clusteralertclearrequest)
    - [ClusterAlertClearResponse](#clusteralertclearresponse)
    - [ClusterAlertEnumerateRequest](#clusteralertenumeraterequest)
    - [ClusterAlertEnumerateResponse](#clusteralertenumerateresponse)
    - [ClusterAlertEraseRequest](#clusteralerteraserequest)
    - [ClusterAlertEraseResponse](#clusteralerteraseresponse)
    - [ClusterEnumerateRequest](#clusterenumeraterequest)
    - [ClusterEnumerateResponse](#clusterenumerateresponse)
    - [ClusterInspectRequest](#clusterinspectrequest)
    - [ClusterInspectResponse](#clusterinspectresponse)
    - [ClusterResponse](#clusterresponse)
    - [GraphDriverChanges](#graphdriverchanges)
    - [Group](#group)
    - [GroupSnapCreateRequest](#groupsnapcreaterequest)
    - [GroupSnapCreateRequest.LabelsEntry](#groupsnapcreaterequestlabelsentry)
    - [GroupSnapCreateResponse](#groupsnapcreateresponse)
    - [GroupSnapCreateResponse.SnapshotsEntry](#groupsnapcreateresponsesnapshotsentry)
    - [OpenStorageVolumeCreateRequest](#openstoragevolumecreaterequest)
    - [OpenStorageVolumeCreateResponse](#openstoragevolumecreateresponse)
    - [ReplicaSet](#replicaset)
    - [RuntimeStateMap](#runtimestatemap)
    - [RuntimeStateMap.RuntimeStateEntry](#runtimestatemapruntimestateentry)
    - [SnapCreateRequest](#snapcreaterequest)
    - [SnapCreateResponse](#snapcreateresponse)
    - [Source](#source)
    - [Stats](#stats)
    - [StorageCluster](#storagecluster)
    - [StorageNode](#storagenode)
    - [StorageNode.DisksEntry](#storagenodedisksentry)
    - [StorageNode.NodeLabelsEntry](#storagenodenodelabelsentry)
    - [StoragePool](#storagepool)
    - [StoragePool.LabelsEntry](#storagepoollabelsentry)
    - [StorageResource](#storageresource)
    - [Volume](#volume)
    - [Volume.AttachInfoEntry](#volumeattachinfoentry)
    - [VolumeConsumer](#volumeconsumer)
    - [VolumeCreateFromVolumeIDRequest](#volumecreatefromvolumeidrequest)
    - [VolumeCreateFromVolumeIDResponse](#volumecreatefromvolumeidresponse)
    - [VolumeCreateRequest](#volumecreaterequest)
    - [VolumeCreateResponse](#volumecreateresponse)
    - [VolumeDeleteRequest](#volumedeleterequest)
    - [VolumeDeleteResponse](#volumedeleteresponse)
    - [VolumeEnumerateRequest](#volumeenumeraterequest)
    - [VolumeEnumerateResponse](#volumeenumerateresponse)
    - [VolumeInfo](#volumeinfo)
    - [VolumeInspectRequest](#volumeinspectrequest)
    - [VolumeInspectResponse](#volumeinspectresponse)
    - [VolumeLocator](#volumelocator)
    - [VolumeLocator.VolumeLabelsEntry](#volumelocatorvolumelabelsentry)
    - [VolumeResponse](#volumeresponse)
    - [VolumeSetRequest](#volumesetrequest)
    - [VolumeSetRequest.OptionsEntry](#volumesetrequestoptionsentry)
    - [VolumeSetResponse](#volumesetresponse)
    - [VolumeSnapshotCreateRequest](#volumesnapshotcreaterequest)
    - [VolumeSnapshotCreateRequest.LabelsEntry](#volumesnapshotcreaterequestlabelsentry)
    - [VolumeSnapshotCreateResponse](#volumesnapshotcreateresponse)
    - [VolumeSnapshotEnumerateRequest](#volumesnapshotenumeraterequest)
    - [VolumeSnapshotEnumerateRequest.LabelsEntry](#volumesnapshotenumeraterequestlabelsentry)
    - [VolumeSnapshotEnumerateResponse](#volumesnapshotenumerateresponse)
    - [VolumeSnapshotRestoreRequest](#volumesnapshotrestorerequest)
    - [VolumeSnapshotRestoreResponse](#volumesnapshotrestoreresponse)
    - [VolumeSpec](#volumespec)
    - [VolumeSpec.VolumeLabelsEntry](#volumespecvolumelabelsentry)
    - [VolumeStateAction](#volumestateaction)
  


- Enums
    - [AlertActionType](#alertactiontype)
    - [AttachState](#attachstate)
    - [ClusterNotify](#clusternotify)
    - [CosType](#costype)
    - [DriverType](#drivertype)
    - [FSType](#fstype)
    - [GraphDriverChangeType](#graphdriverchangetype)
    - [IoProfile](#ioprofile)
    - [OperationFlags](#operationflags)
    - [ResourceType](#resourcetype)
    - [SeverityType](#severitytype)
    - [Status](#status)
    - [StorageMedium](#storagemedium)
    - [VolumeActionParam](#volumeactionparam)
    - [VolumeState](#volumestate)
    - [VolumeStatus](#volumestatus)
  


- [Scalar Value Types](#scalar-value-types)



# OpenStorageCluster {#openstorageapiopenstoragecluster}


## Enumerate 

> **rpc** Enumerate([ClusterEnumerateRequest](#clusterenumeraterequest))
    **returns** [ClusterEnumerateResponse](#clusterenumerateresponse)
   
Enumerate lists all the nodes in the cluster.
## Inspect 

> **rpc** Inspect([ClusterInspectRequest](#clusterinspectrequest))
    **returns** [ClusterInspectResponse](#clusterinspectresponse)
   
Inspect the node given a UUID.
## AlertEnumerate 

> **rpc** AlertEnumerate([ClusterAlertEnumerateRequest](#clusteralertenumeraterequest))
    **returns** [ClusterAlertEnumerateResponse](#clusteralertenumerateresponse)
   
Get a list of alerts from the storage cluster
## AlertClear 

> **rpc** AlertClear([ClusterAlertClearRequest](#clusteralertclearrequest))
    **returns** [ClusterAlertClearResponse](#clusteralertclearresponse)
   
Clear the alert for a given resource
## AlertErase 

> **rpc** AlertErase([ClusterAlertEraseRequest](#clusteralerteraserequest))
    **returns** [ClusterAlertEraseResponse](#clusteralerteraseresponse)
   
Erases an alert for a given resource
 <!-- end methods -->
# OpenStorageVolume {#openstorageapiopenstoragevolume}


## Create 

> **rpc** Create([OpenStorageVolumeCreateRequest](#openstoragevolumecreaterequest))
    **returns** [OpenStorageVolumeCreateResponse](#openstoragevolumecreateresponse)
   
Creates a new volume
## CreateFromVolumeID 

> **rpc** CreateFromVolumeID([VolumeCreateFromVolumeIDRequest](#volumecreatefromvolumeidrequest))
    **returns** [VolumeCreateFromVolumeIDResponse](#volumecreatefromvolumeidresponse)
   
CreateFromVolumeID creates a new volume cloned from an existing volume
## Delete 

> **rpc** Delete([VolumeDeleteRequest](#volumedeleterequest))
    **returns** [VolumeDeleteResponse](#volumedeleteresponse)
   
Delete a volume
## Inspect 

> **rpc** Inspect([VolumeInspectRequest](#volumeinspectrequest))
    **returns** [VolumeInspectResponse](#volumeinspectresponse)
   
Get information on a volume
## Enumerate 

> **rpc** Enumerate([VolumeEnumerateRequest](#volumeenumeraterequest))
    **returns** [VolumeEnumerateResponse](#volumeenumerateresponse)
   
Get a list of volumes
## SnapshotCreate 

> **rpc** SnapshotCreate([VolumeSnapshotCreateRequest](#volumesnapshotcreaterequest))
    **returns** [VolumeSnapshotCreateResponse](#volumesnapshotcreateresponse)
   
Create a snapshot of a volume. This creates an immutable (read-only),
point-in-time snapshot of a volume.
## SnapshotRestore 

> **rpc** SnapshotRestore([VolumeSnapshotRestoreRequest](#volumesnapshotrestorerequest))
    **returns** [VolumeSnapshotRestoreResponse](#volumesnapshotrestoreresponse)
   
Restores a volume to a specified snapshot
## SnapshotEnumerate 

> **rpc** SnapshotEnumerate([VolumeSnapshotEnumerateRequest](#volumesnapshotenumeraterequest))
    **returns** [VolumeSnapshotEnumerateResponse](#volumesnapshotenumerateresponse)
   
List the number of snapshots for a specific volume
 <!-- end methods -->
 <!-- end services -->

# Messages


## ActiveRequest {#activerequest}
Active Request
swagger:model


| Field | Type | Description |
| ----- | ---- | ----------- |
| ReqestKV | repeated [ActiveRequest.ReqestKVEntry](#activerequestreqestkventry) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## ActiveRequest.ReqestKVEntry {#activerequestreqestkventry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key |  [int64](#int64) | none |
| value |  [string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## ActiveRequests {#activerequests}
Active Requests
swagger:model


| Field | Type | Description |
| ----- | ---- | ----------- |
| RequestCount |  [int64](#int64) | none |
| ActiveRequest | repeated [ActiveRequest](#activerequest) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## Alert {#alert}
Alert is a structure that represents an alert object
swagger:model


| Field | Type | Description |
| ----- | ---- | ----------- |
| id |  [int64](#int64) | Id for Alert |
| severity |  [SeverityType](#severitytype) | Severity of the Alert |
| alert_type |  [int64](#int64) | AlertType user defined alert type |
| message |  [string](#string) | Message describing the Alert |
| timestamp |  [google.protobuf.Timestamp](#googleprotobuftimestamp) | Timestamp when Alert occured |
| resource_id |  [string](#string) | ResourceId where Alert occured |
| resource |  [ResourceType](#resourcetype) | Resource where Alert occured |
| cleared |  [bool](#bool) | Cleared Flag |
| ttl |  [uint64](#uint64) | TTL in seconds for this Alert |
| unique_tag |  [string](#string) | UniqueTag helps identify a unique alert for a given resouce |
 <!-- end Fields -->
 <!-- end HasFields -->


## Alerts {#alerts}
Alerts is an array of Alert objects
swagger:model


| Field | Type | Description |
| ----- | ---- | ----------- |
| alert | repeated [Alert](#alert) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## ClusterAlertClearRequest {#clusteralertclearrequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| resource |  [ResourceType](#resourcetype) | Type of resource (required) |
| alert_id |  [int64](#int64) | Id of alert as returned by ClusterEnumerateAlertResponse (required) |
 <!-- end Fields -->
 <!-- end HasFields -->


## ClusterAlertClearResponse {#clusteralertclearresponse}


 <!-- end HasFields -->


## ClusterAlertEnumerateRequest {#clusteralertenumeraterequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| time_start |  [google.protobuf.Timestamp](#googleprotobuftimestamp) | Start time of alerts (required) |
| time_end |  [google.protobuf.Timestamp](#googleprotobuftimestamp) | End time of alerts (required) |
| resource |  [ResourceType](#resourcetype) | Type of resource (required) |
 <!-- end Fields -->
 <!-- end HasFields -->


## ClusterAlertEnumerateResponse {#clusteralertenumerateresponse}



| Field | Type | Description |
| ----- | ---- | ----------- |
| alerts |  [Alerts](#alerts) | Information on the alerts requested |
 <!-- end Fields -->
 <!-- end HasFields -->


## ClusterAlertEraseRequest {#clusteralerteraserequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| resource |  [ResourceType](#resourcetype) | Type of resource (required) |
| alert_id |  [int64](#int64) | Id of alert as returned by ClusterEnumerateAlertResponse (required) |
 <!-- end Fields -->
 <!-- end HasFields -->


## ClusterAlertEraseResponse {#clusteralerteraseresponse}


 <!-- end HasFields -->


## ClusterEnumerateRequest {#clusterenumeraterequest}


 <!-- end HasFields -->


## ClusterEnumerateResponse {#clusterenumerateresponse}



| Field | Type | Description |
| ----- | ---- | ----------- |
| cluster |  [StorageCluster](#storagecluster) | Cluster information |
 <!-- end Fields -->
 <!-- end HasFields -->


## ClusterInspectRequest {#clusterinspectrequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| node_id |  [string](#string) | Id of node to inspect (required) |
 <!-- end Fields -->
 <!-- end HasFields -->


## ClusterInspectResponse {#clusterinspectresponse}



| Field | Type | Description |
| ----- | ---- | ----------- |
| node |  [StorageNode](#storagenode) | Node information |
 <!-- end Fields -->
 <!-- end HasFields -->


## ClusterResponse {#clusterresponse}
ClusterResponse specifies a response that gets returned when requesting the cluster
swagger:response clusterResponse


| Field | Type | Description |
| ----- | ---- | ----------- |
| error |  [string](#string) | Error code

in: body |
 <!-- end Fields -->
 <!-- end HasFields -->


## GraphDriverChanges {#graphdriverchanges}
GraphDriverChanges represent a list of changes between the filesystem layers
specified by the ID and Parent.  // Parent may be an empty string, in which
case there is no parent.
Where the Path is the filesystem path within the layered filesystem
swagger:model


| Field | Type | Description |
| ----- | ---- | ----------- |
| path |  [string](#string) | none |
| kind |  [GraphDriverChangeType](#graphdriverchangetype) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## Group {#group}
Group represents VolumeGroup / namespace
All volumes in the same group share this object.
swagger:model


| Field | Type | Description |
| ----- | ---- | ----------- |
| id |  [string](#string) | Id common identifier across volumes that have the same group. |
 <!-- end Fields -->
 <!-- end HasFields -->


## GroupSnapCreateRequest {#groupsnapcreaterequest}
GroupSnapCreateRequest specifies a request to create a snapshot of given group.
swagger:model


| Field | Type | Description |
| ----- | ---- | ----------- |
| id |  [string](#string) | none |
| Labels | repeated [GroupSnapCreateRequest.LabelsEntry](#groupsnapcreaterequestlabelsentry) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## GroupSnapCreateRequest.LabelsEntry {#groupsnapcreaterequestlabelsentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key |  [string](#string) | none |
| value |  [string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## GroupSnapCreateResponse {#groupsnapcreateresponse}
GroupSnapCreateRequest specifies a response that get's returned when creating a group snapshot.
swagger:response groupSnapCreateResponse


| Field | Type | Description |
| ----- | ---- | ----------- |
| snapshots | repeated [GroupSnapCreateResponse.SnapshotsEntry](#groupsnapcreateresponsesnapshotsentry) | Created snapshots

in: body Required: true |
| error |  [string](#string) | Error message

in: body Required: true |
 <!-- end Fields -->
 <!-- end HasFields -->


## GroupSnapCreateResponse.SnapshotsEntry {#groupsnapcreateresponsesnapshotsentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key |  [string](#string) | none |
| value |  [SnapCreateResponse](#snapcreateresponse) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## OpenStorageVolumeCreateRequest {#openstoragevolumecreaterequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| name |  [string](#string) | Unique name of the volume. This will be used for idempotency. |
| spec |  [VolumeSpec](#volumespec) | Volume specification |
 <!-- end Fields -->
 <!-- end HasFields -->


## OpenStorageVolumeCreateResponse {#openstoragevolumecreateresponse}



| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_id |  [string](#string) | Id of new volume |
 <!-- end Fields -->
 <!-- end HasFields -->


## ReplicaSet {#replicaset}
ReplicaSet set of machine IDs (nodes) to which part of this volume is erasure
coded - for clustered storage arrays
swagger:model


| Field | Type | Description |
| ----- | ---- | ----------- |
| nodes | repeated [string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## RuntimeStateMap {#runtimestatemap}
RuntimeStateMap is a list of name value mapping of driver specific runtime
information.
swagger:model


| Field | Type | Description |
| ----- | ---- | ----------- |
| runtime_state | repeated [RuntimeStateMap.RuntimeStateEntry](#runtimestatemapruntimestateentry) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## RuntimeStateMap.RuntimeStateEntry {#runtimestatemapruntimestateentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key |  [string](#string) | none |
| value |  [string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## SnapCreateRequest {#snapcreaterequest}
SnapCreateRequest specifies a request to create a snapshot of given volume.
swagger:parameters snapVolume


| Field | Type | Description |
| ----- | ---- | ----------- |
| id |  [string](#string) | volume id |
| locator |  [VolumeLocator](#volumelocator) | none |
| readonly |  [bool](#bool) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## SnapCreateResponse {#snapcreateresponse}
SnapCreateRequest specifies a response that get's returned when creating a snapshot.
swagger:response snapCreateResponse


| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_create_response |  [VolumeCreateResponse](#volumecreateresponse) | VolumeCreateResponse

in: body Required: true |
 <!-- end Fields -->
 <!-- end HasFields -->


## Source {#source}
Source is a structure that can be given to a volume
to seed the volume with data.
swagger:model


| Field | Type | Description |
| ----- | ---- | ----------- |
| parent |  [string](#string) | A volume id, if specified will create a clone of the parent. |
| seed |  [string](#string) | Seed will seed the volume from the specified URI Any additional config for the source comes from the labels in the spec |
 <!-- end Fields -->
 <!-- end HasFields -->


## Stats {#stats}
Stats is a structure that represents last collected stats for a volume
swagger:model


| Field | Type | Description |
| ----- | ---- | ----------- |
| reads |  [uint64](#uint64) | Reads completed successfully |
| read_ms |  [uint64](#uint64) | Time spent in reads in ms |
| read_bytes |  [uint64](#uint64) | none |
| writes |  [uint64](#uint64) | Writes completed successfully |
| write_ms |  [uint64](#uint64) | Time spent in writes in ms |
| write_bytes |  [uint64](#uint64) | none |
| io_progress |  [uint64](#uint64) | IOs curently in progress |
| io_ms |  [uint64](#uint64) | Time spent doing IOs ms |
| bytes_used |  [uint64](#uint64) | BytesUsed |
| interval_ms |  [uint64](#uint64) | Interval in ms during which stats were collected |
 <!-- end Fields -->
 <!-- end HasFields -->


## StorageCluster {#storagecluster}
StorageCluster represents the state of the cluster


| Field | Type | Description |
| ----- | ---- | ----------- |
| status |  [Status](#status) | Status of the cluster |
| id |  [string](#string) | Id of the cluster |
| node_id |  [string](#string) | NodeId is the id of the node servicing these requests |
| nodes | repeated [StorageNode](#storagenode) | Nodes are a list of all the nodes on the cluster |
 <!-- end Fields -->
 <!-- end HasFields -->


## StorageNode {#storagenode}
StorageNode describes the state of the node


| Field | Type | Description |
| ----- | ---- | ----------- |
| id |  [string](#string) | Id of the node |
| cpu |  [double](#double) | Cpu usage of the node |
| mem_total |  [uint64](#uint64) | Total memory of the node |
| mem_used |  [uint64](#uint64) | Used memory of the node |
| mem_free |  [uint64](#uint64) | Free memory of the node |
| avg_load |  [int64](#int64) | Average load (percentage) |
| status |  [Status](#status) | Node status |
| disks | repeated [StorageNode.DisksEntry](#storagenodedisksentry) | List of disks on the node |
| pools | repeated [StoragePool](#storagepool) | List of storage pools this node supports |
| mgmt_ip |  [string](#string) | Management IP |
| data_ip |  [string](#string) | Data IP |
| hostname |  [string](#string) | Hostname of the node |
| node_labels | repeated [StorageNode.NodeLabelsEntry](#storagenodenodelabelsentry) | User defined labels for the node |
 <!-- end Fields -->
 <!-- end HasFields -->


## StorageNode.DisksEntry {#storagenodedisksentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key |  [string](#string) | none |
| value |  [StorageResource](#storageresource) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## StorageNode.NodeLabelsEntry {#storagenodenodelabelsentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key |  [string](#string) | none |
| value |  [string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## StoragePool {#storagepool}
StoragePool groups different storage devices based on their CosType


| Field | Type | Description |
| ----- | ---- | ----------- |
| ID |  [int32](#int32) | ID pool ID |
| Cos |  [CosType](#costype) | Cos reflects the capabilities of this drive pool |
| Medium |  [StorageMedium](#storagemedium) | Medium underlying storage type |
| RaidLevel |  [string](#string) | RaidLevel storage raid level |
| TotalSize |  [uint64](#uint64) | TotalSize of the pool |
| Used |  [uint64](#uint64) | Used size of the pool |
| labels | repeated [StoragePool.LabelsEntry](#storagepoollabelsentry) | Labels is a list of user defined name-value pairs |
 <!-- end Fields -->
 <!-- end HasFields -->


## StoragePool.LabelsEntry {#storagepoollabelsentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key |  [string](#string) | none |
| value |  [string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## StorageResource {#storageresource}
StorageResource groups properties of a storage device.
swagger:model


| Field | Type | Description |
| ----- | ---- | ----------- |
| id |  [string](#string) | Id is the LUN identifier. |
| path |  [string](#string) | Path device path for this storage resource. |
| medium |  [StorageMedium](#storagemedium) | Storage medium. |
| online |  [bool](#bool) | True if this device is online. |
| iops |  [uint64](#uint64) | IOPS |
| seq_write |  [double](#double) | SeqWrite |
| seq_read |  [double](#double) | SeqRead |
| randRW |  [double](#double) | RandRW |
| size |  [uint64](#uint64) | Total size in bytes. |
| used |  [uint64](#uint64) | Physical Bytes used. |
| rotation_speed |  [string](#string) | True if this device is rotational. |
| last_scan |  [google.protobuf.Timestamp](#googleprotobuftimestamp) | Timestamp of last time this device was scanned. |
| metadata |  [bool](#bool) | True if dedicated for metadata. |
 <!-- end Fields -->
 <!-- end HasFields -->


## Volume {#volume}
Volume represents an abstract storage volume.
Volume represents an abstract storage volume.
swagger:model


| Field | Type | Description |
| ----- | ---- | ----------- |
| id |  [string](#string) | Self referential volume ID. |
| source |  [Source](#source) | Source specified seed data for the volume. |
| group |  [Group](#group) | Group volumes in the same group have the same group id. |
| readonly |  [bool](#bool) | Readonly is true if this volume is to be mounted with readonly access. |
| locator |  [VolumeLocator](#volumelocator) | User specified locator |
| ctime |  [google.protobuf.Timestamp](#googleprotobuftimestamp) | Volume creation time |
| spec |  [VolumeSpec](#volumespec) | User specified VolumeSpec |
| usage |  [uint64](#uint64) | Usage is bytes consumed by vtheis volume. |
| last_scan |  [google.protobuf.Timestamp](#googleprotobuftimestamp) | LastScan is the time when an integrity check was run. |
| format |  [FSType](#fstype) | Format specifies the filesytem for this volume. |
| status |  [VolumeStatus](#volumestatus) | Status is the availability status of this volume. |
| state |  [VolumeState](#volumestate) | State is the current runtime state of this volume. |
| attached_on |  [string](#string) | AttachedOn is the node instance identifier for clustered systems. |
| attached_state |  [AttachState](#attachstate) | AttachedState shows whether the device is attached for internal or external use. |
| device_path |  [string](#string) | DevicePath is the device exported by block device implementations. |
| secure_device_path |  [string](#string) | SecureDevicePath is the device path for an encrypted volume. |
| attach_path | repeated [string](#string) | AttachPath is the mounted path in the host namespace. |
| attach_info | repeated [Volume.AttachInfoEntry](#volumeattachinfoentry) | AttachInfo is a list of name value mappings that provides attach information. |
| replica_sets | repeated [ReplicaSet](#replicaset) | ReplicatSets storage for this volumefor clustered storage arrays. |
| runtime_state | repeated [RuntimeStateMap](#runtimestatemap) | RuntimeState is a lst of name value mapping of driver specific runtime information. |
| error |  [string](#string) | Error is the Last recorded error. |
| volume_consumers | repeated [VolumeConsumer](#volumeconsumer) | VolumeConsumers are entities that consume this volume |
 <!-- end Fields -->
 <!-- end HasFields -->


## Volume.AttachInfoEntry {#volumeattachinfoentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key |  [string](#string) | none |
| value |  [string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeConsumer {#volumeconsumer}
VolumeConsumer identifies a consumer for a Volume. An example of a VolumeConsumer
would be a Pod in Kubernetes who has mounted the PersistentVolumeClaim for the
Volume
swagger: model


| Field | Type | Description |
| ----- | ---- | ----------- |
| name |  [string](#string) | Name is the name of the volume consumer |
| namespace |  [string](#string) | Namespace is the namespace of the volume consumer |
| type |  [string](#string) | Type is the type of the consumer. E.g a Kubernetes pod |
| node_id |  [string](#string) | NodeID is the identifier of the node on which the consumer is running. This identifier would be from the perspective of the container runtime or orchestrator under which the volume consumer resides. For example, NodeID can be name of a minion in Kubernetes. |
| owner_name |  [string](#string) | OwnerName is the name of the entity who owns this volume consumer |
| owner_type |  [string](#string) | OwnerType is the type of the entity who owns this volume consumer. The type would be from the perspective of the container runtime or the orchestrator under which the volume consumer resides. For e.g OwnerType can be a Deployment in Kubernetes. |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeCreateFromVolumeIDRequest {#volumecreatefromvolumeidrequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| name |  [string](#string) | Unique name of the volume. This will be used for idempotency. |
| parent_id |  [string](#string) | Parent volume id, if specified will create a new volume as a clone of the parent. |
| spec |  [VolumeSpec](#volumespec) | Volume specification |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeCreateFromVolumeIDResponse {#volumecreatefromvolumeidresponse}



| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_id |  [string](#string) | Id of new volume |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeCreateRequest {#volumecreaterequest}
VolumeCreateRequest is a structure that has the locator, source and spec
to create a volume
swagger:model


| Field | Type | Description |
| ----- | ---- | ----------- |
| locator |  [VolumeLocator](#volumelocator) | User specified volume name and labels |
| source |  [Source](#source) | Source to create volume |
| spec |  [VolumeSpec](#volumespec) | The storage spec for the volume |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeCreateResponse {#volumecreateresponse}
VolumeCreateResponse
swagger:response volumeCreateResponse


| Field | Type | Description |
| ----- | ---- | ----------- |
| id |  [string](#string) | ID of the newly created volume

in: body Required: true |
| volume_response |  [VolumeResponse](#volumeresponse) | Volume Response

in: body Required: true |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeDeleteRequest {#volumedeleterequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_id |  [string](#string) | Id of volume to delete |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeDeleteResponse {#volumedeleteresponse}


 <!-- end HasFields -->


## VolumeEnumerateRequest {#volumeenumeraterequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| locator |  [VolumeLocator](#volumelocator) | Volumes to match to this locator. If not provided, all volumes will be returned. |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeEnumerateResponse {#volumeenumerateresponse}



| Field | Type | Description |
| ----- | ---- | ----------- |
| volumes | repeated [Volume](#volume) | List of volumes matching label |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeInfo {#volumeinfo}
VolumeInfo
swagger:model


| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_id |  [string](#string) | none |
| path |  [string](#string) | none |
| storage |  [VolumeSpec](#volumespec) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeInspectRequest {#volumeinspectrequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_id |  [string](#string) | Id of volume to inspect |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeInspectResponse {#volumeinspectresponse}



| Field | Type | Description |
| ----- | ---- | ----------- |
| volume |  [Volume](#volume) | Information about the volume |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeLocator {#volumelocator}
VolumeLocator is a structure that is attached to a volume
and is used to carry opaque metadata.
swagger:model


| Field | Type | Description |
| ----- | ---- | ----------- |
| name |  [string](#string) | User friendly identifier |
| volume_labels | repeated [VolumeLocator.VolumeLabelsEntry](#volumelocatorvolumelabelsentry) | A set of name-value pairs that acts as search filters |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeLocator.VolumeLabelsEntry {#volumelocatorvolumelabelsentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key |  [string](#string) | none |
| value |  [string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeResponse {#volumeresponse}
VolumeResponse is a structure that wraps an error.
swagger:response volumeResponse


| Field | Type | Description |
| ----- | ---- | ----------- |
| error |  [string](#string) | Error message

in: body Required: true |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeSetRequest {#volumesetrequest}
VolumeSet specifies a request to update a volume.
swagger:model


| Field | Type | Description |
| ----- | ---- | ----------- |
| locator |  [VolumeLocator](#volumelocator) | User specified volume name and labels |
| spec |  [VolumeSpec](#volumespec) | The storage spec for the volume |
| action |  [VolumeStateAction](#volumestateaction) | State modification on this volume. |
| options | repeated [VolumeSetRequest.OptionsEntry](#volumesetrequestoptionsentry) | additional options required for the Set operation. |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeSetRequest.OptionsEntry {#volumesetrequestoptionsentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key |  [string](#string) | none |
| value |  [string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeSetResponse {#volumesetresponse}
VolumeSetResponse
swagger:response volumeSetResponse


| Field | Type | Description |
| ----- | ---- | ----------- |
| volume |  [Volume](#volume) | Volume

in: body Required: true |
| volume_response |  [VolumeResponse](#volumeresponse) | VolumeResponse

in: body Required: true |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeSnapshotCreateRequest {#volumesnapshotcreaterequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_id |  [string](#string) | Id of volume to take the snapshot from |
| labels | repeated [VolumeSnapshotCreateRequest.LabelsEntry](#volumesnapshotcreaterequestlabelsentry) | Labels to apply to snapshot |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeSnapshotCreateRequest.LabelsEntry {#volumesnapshotcreaterequestlabelsentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key |  [string](#string) | none |
| value |  [string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeSnapshotCreateResponse {#volumesnapshotcreateresponse}



| Field | Type | Description |
| ----- | ---- | ----------- |
| snapshot_id |  [string](#string) | Id of immutable snapshot |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeSnapshotEnumerateRequest {#volumesnapshotenumeraterequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_id |  [string](#string) | Id of volume |
| labels | repeated [VolumeSnapshotEnumerateRequest.LabelsEntry](#volumesnapshotenumeraterequestlabelsentry) | Labels from snapshot |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeSnapshotEnumerateRequest.LabelsEntry {#volumesnapshotenumeraterequestlabelsentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key |  [string](#string) | none |
| value |  [string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeSnapshotEnumerateResponse {#volumesnapshotenumerateresponse}



| Field | Type | Description |
| ----- | ---- | ----------- |
| snapshots | repeated [Volume](#volume) | List of immutable snapshots |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeSnapshotRestoreRequest {#volumesnapshotrestorerequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_id |  [string](#string) | Id of volume |
| snapshot_id |  [string](#string) | Snapshot id to apply to `volume_id` |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeSnapshotRestoreResponse {#volumesnapshotrestoreresponse}


 <!-- end HasFields -->


## VolumeSpec {#volumespec}
VolumeSpec has the properties needed to create a volume.
swagger:model


| Field | Type | Description |
| ----- | ---- | ----------- |
| ephemeral |  [bool](#bool) | Ephemeral storage |
| size |  [uint64](#uint64) | Size specifies the thin provisioned volume size. |
| format |  [FSType](#fstype) | Format specifies the filesystem for this volume. |
| block_size |  [int64](#int64) | BlockSize for the filesystem. |
| ha_level |  [int64](#int64) | HaLevel specifies the number of copies of data. |
| cos |  [CosType](#costype) | Cos specifies the relative class of service. |
| io_profile |  [IoProfile](#ioprofile) | IoProfile provides a hint about application using this volume. |
| dedupe |  [bool](#bool) | Dedupe specifies if the volume data is to be de-duplicated. |
| snapshot_interval |  [uint32](#uint32) | SnapshotInterval in minutes, set to 0 to disable snapshots |
| volume_labels | repeated [VolumeSpec.VolumeLabelsEntry](#volumespecvolumelabelsentry) | VolumeLabels configuration labels |
| shared |  [bool](#bool) | Shared is true if this volume can be remotely accessed. |
| replica_set |  [ReplicaSet](#replicaset) | ReplicaSet is the desired set of nodes for the volume data. |
| aggregation_level |  [uint32](#uint32) | Aggregatiokn level Specifies the number of parts the volume can be aggregated from. |
| encrypted |  [bool](#bool) | Encrypted is true if this volume will be cryptographically secured. |
| passphrase |  [string](#string) | Passphrase for an encrypted volume |
| snapshot_schedule |  [string](#string) | SnapshotSchedule a well known string that specifies when snapshots should be taken. |
| scale |  [uint32](#uint32) | Scale allows autocreation of volumes. |
| sticky |  [bool](#bool) | Sticky volumes cannot be deleted until the flag is removed. |
| group |  [Group](#group) | Group identifies a consistency group |
| group_enforced |  [bool](#bool) | GroupEnforced is true if consistency group creation is enforced. |
| compressed |  [bool](#bool) | Compressed is true if this volume is to be compressed. |
| cascaded |  [bool](#bool) | Cascaded is true if this volume can be populated on any node from an external source. |
| journal |  [bool](#bool) | Journal is true if data for the volume goes into the journal. |
| sharedv4 |  [bool](#bool) | Sharedv4 is true if this volume can be accessed via sharedv4. |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeSpec.VolumeLabelsEntry {#volumespecvolumelabelsentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key |  [string](#string) | none |
| value |  [string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeStateAction {#volumestateaction}
VolumeStateAction specifies desired actions.
swagger:model


| Field | Type | Description |
| ----- | ---- | ----------- |
| attach |  [VolumeActionParam](#volumeactionparam) | Attach or Detach volume |
| mount |  [VolumeActionParam](#volumeactionparam) | Mount or unmount volume |
| mount_path |  [string](#string) | MountPath Path where the device is mounted |
| device_path |  [string](#string) | DevicePath Path returned in attach |
 <!-- end Fields -->
 <!-- end HasFields -->
 <!-- end messages -->

# Enums


## AlertActionType {#alertactiontype}


| Name | Number | Description |
| ---- | ------ | ----------- |
| ALERT_ACTION_TYPE_NONE | 0 |  |
| ALERT_ACTION_TYPE_DELETE | 1 |  |
| ALERT_ACTION_TYPE_CREATE | 2 |  |
| ALERT_ACTION_TYPE_UPDATE | 3 |  |




## AttachState {#attachstate}


| Name | Number | Description |
| ---- | ------ | ----------- |
| ATTACH_STATE_EXTERNAL | 0 | Attached and available externally |
| ATTACH_STATE_INTERNAL | 1 | Attached but only available internally |
| ATTACH_STATE_INTERNAL_SWITCH | 2 | Switching from External to Internal |




## ClusterNotify {#clusternotify}


| Name | Number | Description |
| ---- | ------ | ----------- |
| CLUSTER_NOTIFY_DOWN | 0 | Node is down |




## CosType {#costype}


| Name | Number | Description |
| ---- | ------ | ----------- |
| NONE | 0 |  |
| LOW | 1 |  |
| MEDIUM | 2 |  |
| HIGH | 3 |  |




## DriverType {#drivertype}


| Name | Number | Description |
| ---- | ------ | ----------- |
| DRIVER_TYPE_NONE | 0 |  |
| DRIVER_TYPE_FILE | 1 |  |
| DRIVER_TYPE_BLOCK | 2 |  |
| DRIVER_TYPE_OBJECT | 3 |  |
| DRIVER_TYPE_CLUSTERED | 4 |  |
| DRIVER_TYPE_GRAPH | 5 |  |




## FSType {#fstype}


| Name | Number | Description |
| ---- | ------ | ----------- |
| FS_TYPE_NONE | 0 |  |
| FS_TYPE_BTRFS | 1 |  |
| FS_TYPE_EXT4 | 2 |  |
| FS_TYPE_FUSE | 3 |  |
| FS_TYPE_NFS | 4 |  |
| FS_TYPE_VFS | 5 |  |
| FS_TYPE_XFS | 6 |  |
| FS_TYPE_ZFS | 7 |  |




## GraphDriverChangeType {#graphdriverchangetype}


| Name | Number | Description |
| ---- | ------ | ----------- |
| GRAPH_DRIVER_CHANGE_TYPE_NONE | 0 |  |
| GRAPH_DRIVER_CHANGE_TYPE_MODIFIED | 1 |  |
| GRAPH_DRIVER_CHANGE_TYPE_ADDED | 2 |  |
| GRAPH_DRIVER_CHANGE_TYPE_DELETED | 3 |  |




## IoProfile {#ioprofile}


| Name | Number | Description |
| ---- | ------ | ----------- |
| IO_PROFILE_SEQUENTIAL | 0 |  |
| IO_PROFILE_RANDOM | 1 |  |
| IO_PROFILE_DB | 2 |  |
| IO_PROFILE_DB_REMOTE | 3 |  |
| IO_PROFILE_CMS | 4 |  |




## OperationFlags {#operationflags}


| Name | Number | Description |
| ---- | ------ | ----------- |
| OP_FLAGS_UNKNOWN | 0 |  |
| OP_FLAGS_NONE | 1 |  |
| OP_FLAGS_DETACH_FORCE | 2 | Perform a force_detach during detach operation |




## ResourceType {#resourcetype}


| Name | Number | Description |
| ---- | ------ | ----------- |
| RESOURCE_TYPE_NONE | 0 |  |
| RESOURCE_TYPE_VOLUME | 1 |  |
| RESOURCE_TYPE_NODE | 2 |  |
| RESOURCE_TYPE_CLUSTER | 3 |  |
| RESOURCE_TYPE_DRIVE | 4 |  |




## SeverityType {#severitytype}


| Name | Number | Description |
| ---- | ------ | ----------- |
| SEVERITY_TYPE_NONE | 0 |  |
| SEVERITY_TYPE_ALARM | 1 |  |
| SEVERITY_TYPE_WARNING | 2 |  |
| SEVERITY_TYPE_NOTIFY | 3 |  |




## Status {#status}


| Name | Number | Description |
| ---- | ------ | ----------- |
| STATUS_NONE | 0 |  |
| STATUS_INIT | 1 |  |
| STATUS_OK | 2 |  |
| STATUS_OFFLINE | 3 |  |
| STATUS_ERROR | 4 |  |
| STATUS_NOT_IN_QUORUM | 5 |  |
| STATUS_DECOMMISSION | 6 |  |
| STATUS_MAINTENANCE | 7 |  |
| STATUS_STORAGE_DOWN | 8 |  |
| STATUS_STORAGE_DEGRADED | 9 |  |
| STATUS_NEEDS_REBOOT | 10 |  |
| STATUS_STORAGE_REBALANCE | 11 |  |
| STATUS_STORAGE_DRIVE_REPLACE | 12 |  |
| STATUS_NOT_IN_QUORUM_NO_STORAGE | 13 |  |
| STATUS_MAX | 14 | Add statuses before MAX and update the number for MAX |




## StorageMedium {#storagemedium}


| Name | Number | Description |
| ---- | ------ | ----------- |
| STORAGE_MEDIUM_MAGNETIC | 0 | Magnetic spinning disk. |
| STORAGE_MEDIUM_SSD | 1 | SSD disk |
| STORAGE_MEDIUM_NVME | 2 | NVME disk |




## VolumeActionParam {#volumeactionparam}


| Name | Number | Description |
| ---- | ------ | ----------- |
| VOLUME_ACTION_PARAM_NONE | 0 |  |
| VOLUME_ACTION_PARAM_OFF | 1 | Maps to the boolean value false |
| VOLUME_ACTION_PARAM_ON | 2 | Maps to the boolean value true. |




## VolumeState {#volumestate}
VolumeState represents the state of a volume.

| Name | Number | Description |
| ---- | ------ | ----------- |
| VOLUME_STATE_NONE | 0 |  |
| VOLUME_STATE_PENDING | 1 | Volume is transitioning to new state |
| VOLUME_STATE_AVAILABLE | 2 | Volume is ready to be assigned to a container |
| VOLUME_STATE_ATTACHED | 3 | Volume is attached to container |
| VOLUME_STATE_DETACHED | 4 | Volume is detached but associated with a container |
| VOLUME_STATE_DETATCHING | 5 | Volume detach is in progress |
| VOLUME_STATE_ERROR | 6 | Volume is in error state |
| VOLUME_STATE_DELETED | 7 | Volume is deleted, it will remain in this state while resources are asynchronously reclaimed |
| VOLUME_STATE_TRY_DETACHING | 8 | Volume is trying to be detached |
| VOLUME_STATE_RESTORE | 9 | Volume is undergoing restore |




## VolumeStatus {#volumestatus}
VolumeStatus represents a health status for a volume.

| Name | Number | Description |
| ---- | ------ | ----------- |
| VOLUME_STATUS_NONE | 0 |  |
| VOLUME_STATUS_NOT_PRESENT | 1 | Volume is not present |
| VOLUME_STATUS_UP | 2 | Volume is healthy |
| VOLUME_STATUS_DOWN | 3 | Volume is in fail mode |
| VOLUME_STATUS_DEGRADED | 4 | Volume is up but with degraded performance In a RAID group, this may indicate a problem with one or more drives |


 <!-- end Enums -->
 <!-- end Files -->

# Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <div><h4 id="double" /></div><a name="double" /> double |  | double | double | float |
| <div><h4 id="float" /></div><a name="float" /> float |  | float | float | float |
| <div><h4 id="int32" /></div><a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <div><h4 id="int64" /></div><a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <div><h4 id="uint32" /></div><a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <div><h4 id="uint64" /></div><a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <div><h4 id="sint32" /></div><a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <div><h4 id="sint64" /></div><a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <div><h4 id="fixed32" /></div><a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <div><h4 id="fixed64" /></div><a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <div><h4 id="sfixed32" /></div><a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <div><h4 id="sfixed64" /></div><a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <div><h4 id="bool" /></div><a name="bool" /> bool |  | bool | boolean | boolean |
| <div><h4 id="string" /></div><a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <div><h4 id="bytes" /></div><a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |
