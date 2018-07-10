# API Reference

# Table of Contents


- Services
    - [OpenStorageCloudBackup](#openstorageapiopenstoragecloudbackup)
    - [OpenStorageCluster](#openstorageapiopenstoragecluster)
    - [OpenStorageCredentials](#openstorageapiopenstoragecredentials)
    - [OpenStorageObjectstore](#openstorageapiopenstorageobjectstore)
    - [OpenStorageSchedulePolicy](#openstorageapiopenstorageschedulepolicy)
    - [OpenStorageVolume](#openstorageapiopenstoragevolume)
  


- Messages
    - [ActiveRequest](#activerequest)
    - [ActiveRequest.ReqestKVEntry](#activerequestreqestkventry)
    - [ActiveRequests](#activerequests)
    - [Alert](#alert)
    - [Alerts](#alerts)
    - [CloudMigrate](#cloudmigrate)
    - [CloudMigrateCancelRequest](#cloudmigratecancelrequest)
    - [CloudMigrateInfo](#cloudmigrateinfo)
    - [CloudMigrateInfoList](#cloudmigrateinfolist)
    - [CloudMigrateStartRequest](#cloudmigratestartrequest)
    - [CloudMigrateStatusResponse](#cloudmigratestatusresponse)
    - [CloudMigrateStatusResponse.InfoEntry](#cloudmigratestatusresponseinfoentry)
    - [ClusterResponse](#clusterresponse)
    - [GraphDriverChanges](#graphdriverchanges)
    - [Group](#group)
    - [GroupSnapCreateRequest](#groupsnapcreaterequest)
    - [GroupSnapCreateRequest.LabelsEntry](#groupsnapcreaterequestlabelsentry)
    - [GroupSnapCreateResponse](#groupsnapcreateresponse)
    - [GroupSnapCreateResponse.SnapshotsEntry](#groupsnapcreateresponsesnapshotsentry)
    - [ObjectstoreInfo](#objectstoreinfo)
    - [ReplicaSet](#replicaset)
    - [RuntimeStateMap](#runtimestatemap)
    - [RuntimeStateMap.RuntimeStateEntry](#runtimestatemapruntimestateentry)
    - [SdkAwsCredentialRequest](#sdkawscredentialrequest)
    - [SdkAwsCredentialResponse](#sdkawscredentialresponse)
    - [SdkAzureCredentialRequest](#sdkazurecredentialrequest)
    - [SdkAzureCredentialResponse](#sdkazurecredentialresponse)
    - [SdkCloudBackupCatalogRequest](#sdkcloudbackupcatalogrequest)
    - [SdkCloudBackupCatalogResponse](#sdkcloudbackupcatalogresponse)
    - [SdkCloudBackupCreateRequest](#sdkcloudbackupcreaterequest)
    - [SdkCloudBackupCreateResponse](#sdkcloudbackupcreateresponse)
    - [SdkCloudBackupDeleteAllRequest](#sdkcloudbackupdeleteallrequest)
    - [SdkCloudBackupDeleteAllResponse](#sdkcloudbackupdeleteallresponse)
    - [SdkCloudBackupDeleteRequest](#sdkcloudbackupdeleterequest)
    - [SdkCloudBackupDeleteResponse](#sdkcloudbackupdeleteresponse)
    - [SdkCloudBackupEnumerateRequest](#sdkcloudbackupenumeraterequest)
    - [SdkCloudBackupEnumerateResponse](#sdkcloudbackupenumerateresponse)
    - [SdkCloudBackupHistoryItem](#sdkcloudbackuphistoryitem)
    - [SdkCloudBackupHistoryRequest](#sdkcloudbackuphistoryrequest)
    - [SdkCloudBackupHistoryResponse](#sdkcloudbackuphistoryresponse)
    - [SdkCloudBackupInfo](#sdkcloudbackupinfo)
    - [SdkCloudBackupInfo.MetadataEntry](#sdkcloudbackupinfometadataentry)
    - [SdkCloudBackupRestoreRequest](#sdkcloudbackuprestorerequest)
    - [SdkCloudBackupRestoreResponse](#sdkcloudbackuprestoreresponse)
    - [SdkCloudBackupSchedCreateRequest](#sdkcloudbackupschedcreaterequest)
    - [SdkCloudBackupSchedCreateResponse](#sdkcloudbackupschedcreateresponse)
    - [SdkCloudBackupSchedDeleteRequest](#sdkcloudbackupscheddeleterequest)
    - [SdkCloudBackupSchedDeleteResponse](#sdkcloudbackupscheddeleteresponse)
    - [SdkCloudBackupSchedEnumerateRequest](#sdkcloudbackupschedenumeraterequest)
    - [SdkCloudBackupSchedEnumerateResponse](#sdkcloudbackupschedenumerateresponse)
    - [SdkCloudBackupSchedEnumerateResponse.CloudSchedListEntry](#sdkcloudbackupschedenumerateresponsecloudschedlistentry)
    - [SdkCloudBackupScheduleInfo](#sdkcloudbackupscheduleinfo)
    - [SdkCloudBackupStateChangeRequest](#sdkcloudbackupstatechangerequest)
    - [SdkCloudBackupStateChangeResponse](#sdkcloudbackupstatechangeresponse)
    - [SdkCloudBackupStatus](#sdkcloudbackupstatus)
    - [SdkCloudBackupStatusRequest](#sdkcloudbackupstatusrequest)
    - [SdkCloudBackupStatusResponse](#sdkcloudbackupstatusresponse)
    - [SdkCloudBackupStatusResponse.StatusesEntry](#sdkcloudbackupstatusresponsestatusesentry)
    - [SdkClusterAlertClearRequest](#sdkclusteralertclearrequest)
    - [SdkClusterAlertClearResponse](#sdkclusteralertclearresponse)
    - [SdkClusterAlertDeleteRequest](#sdkclusteralertdeleterequest)
    - [SdkClusterAlertDeleteResponse](#sdkclusteralertdeleteresponse)
    - [SdkClusterAlertEnumerateRequest](#sdkclusteralertenumeraterequest)
    - [SdkClusterAlertEnumerateResponse](#sdkclusteralertenumerateresponse)
    - [SdkClusterEnumerateRequest](#sdkclusterenumeraterequest)
    - [SdkClusterEnumerateResponse](#sdkclusterenumerateresponse)
    - [SdkClusterInspectRequest](#sdkclusterinspectrequest)
    - [SdkClusterInspectResponse](#sdkclusterinspectresponse)
    - [SdkCredentialCreateRequest](#sdkcredentialcreaterequest)
    - [SdkCredentialCreateResponse](#sdkcredentialcreateresponse)
    - [SdkCredentialDeleteRequest](#sdkcredentialdeleterequest)
    - [SdkCredentialDeleteResponse](#sdkcredentialdeleteresponse)
    - [SdkCredentialEnumerateRequest](#sdkcredentialenumeraterequest)
    - [SdkCredentialEnumerateResponse](#sdkcredentialenumerateresponse)
    - [SdkCredentialInspectRequest](#sdkcredentialinspectrequest)
    - [SdkCredentialInspectResponse](#sdkcredentialinspectresponse)
    - [SdkCredentialValidateRequest](#sdkcredentialvalidaterequest)
    - [SdkCredentialValidateResponse](#sdkcredentialvalidateresponse)
    - [SdkGoogleCredentialRequest](#sdkgooglecredentialrequest)
    - [SdkGoogleCredentialResponse](#sdkgooglecredentialresponse)
    - [SdkObjectstoreCreateRequest](#sdkobjectstorecreaterequest)
    - [SdkObjectstoreCreateResponse](#sdkobjectstorecreateresponse)
    - [SdkObjectstoreDeleteRequest](#sdkobjectstoredeleterequest)
    - [SdkObjectstoreDeleteResponse](#sdkobjectstoredeleteresponse)
    - [SdkObjectstoreInspectRequest](#sdkobjectstoreinspectrequest)
    - [SdkObjectstoreInspectResponse](#sdkobjectstoreinspectresponse)
    - [SdkObjectstoreUpdateRequest](#sdkobjectstoreupdaterequest)
    - [SdkObjectstoreUpdateResponse](#sdkobjectstoreupdateresponse)
    - [SdkSchedulePolicy](#sdkschedulepolicy)
    - [SdkSchedulePolicyCreateRequest](#sdkschedulepolicycreaterequest)
    - [SdkSchedulePolicyCreateResponse](#sdkschedulepolicycreateresponse)
    - [SdkSchedulePolicyDeleteRequest](#sdkschedulepolicydeleterequest)
    - [SdkSchedulePolicyDeleteResponse](#sdkschedulepolicydeleteresponse)
    - [SdkSchedulePolicyEnumerateRequest](#sdkschedulepolicyenumeraterequest)
    - [SdkSchedulePolicyEnumerateResponse](#sdkschedulepolicyenumerateresponse)
    - [SdkSchedulePolicyInspectRequest](#sdkschedulepolicyinspectrequest)
    - [SdkSchedulePolicyInspectResponse](#sdkschedulepolicyinspectresponse)
    - [SdkSchedulePolicyInterval](#sdkschedulepolicyinterval)
    - [SdkSchedulePolicyIntervalDaily](#sdkschedulepolicyintervaldaily)
    - [SdkSchedulePolicyIntervalMonthly](#sdkschedulepolicyintervalmonthly)
    - [SdkSchedulePolicyIntervalWeekly](#sdkschedulepolicyintervalweekly)
    - [SdkSchedulePolicyUpdateRequest](#sdkschedulepolicyupdaterequest)
    - [SdkSchedulePolicyUpdateResponse](#sdkschedulepolicyupdateresponse)
    - [SdkVolumeAttachRequest](#sdkvolumeattachrequest)
    - [SdkVolumeAttachRequest.OptionsEntry](#sdkvolumeattachrequestoptionsentry)
    - [SdkVolumeAttachResponse](#sdkvolumeattachresponse)
    - [SdkVolumeCloneRequest](#sdkvolumeclonerequest)
    - [SdkVolumeCloneResponse](#sdkvolumecloneresponse)
    - [SdkVolumeCreateRequest](#sdkvolumecreaterequest)
    - [SdkVolumeCreateResponse](#sdkvolumecreateresponse)
    - [SdkVolumeDeleteRequest](#sdkvolumedeleterequest)
    - [SdkVolumeDeleteResponse](#sdkvolumedeleteresponse)
    - [SdkVolumeDetachRequest](#sdkvolumedetachrequest)
    - [SdkVolumeDetachResponse](#sdkvolumedetachresponse)
    - [SdkVolumeEnumerateRequest](#sdkvolumeenumeraterequest)
    - [SdkVolumeEnumerateResponse](#sdkvolumeenumerateresponse)
    - [SdkVolumeInspectRequest](#sdkvolumeinspectrequest)
    - [SdkVolumeInspectResponse](#sdkvolumeinspectresponse)
    - [SdkVolumeMountRequest](#sdkvolumemountrequest)
    - [SdkVolumeMountRequest.OptionsEntry](#sdkvolumemountrequestoptionsentry)
    - [SdkVolumeMountResponse](#sdkvolumemountresponse)
    - [SdkVolumeSetRequest](#sdkvolumesetrequest)
    - [SdkVolumeSetResponse](#sdkvolumesetresponse)
    - [SdkVolumeSnapshotCreateRequest](#sdkvolumesnapshotcreaterequest)
    - [SdkVolumeSnapshotCreateRequest.LabelsEntry](#sdkvolumesnapshotcreaterequestlabelsentry)
    - [SdkVolumeSnapshotCreateResponse](#sdkvolumesnapshotcreateresponse)
    - [SdkVolumeSnapshotEnumerateRequest](#sdkvolumesnapshotenumeraterequest)
    - [SdkVolumeSnapshotEnumerateRequest.LabelsEntry](#sdkvolumesnapshotenumeraterequestlabelsentry)
    - [SdkVolumeSnapshotEnumerateResponse](#sdkvolumesnapshotenumerateresponse)
    - [SdkVolumeSnapshotRestoreRequest](#sdkvolumesnapshotrestorerequest)
    - [SdkVolumeSnapshotRestoreResponse](#sdkvolumesnapshotrestoreresponse)
    - [SdkVolumeUnmountRequest](#sdkvolumeunmountrequest)
    - [SdkVolumeUnmountRequest.OptionsEntry](#sdkvolumeunmountrequestoptionsentry)
    - [SdkVolumeUnmountResponse](#sdkvolumeunmountresponse)
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
    - [VolumeCreateRequest](#volumecreaterequest)
    - [VolumeCreateResponse](#volumecreateresponse)
    - [VolumeInfo](#volumeinfo)
    - [VolumeLocator](#volumelocator)
    - [VolumeLocator.VolumeLabelsEntry](#volumelocatorvolumelabelsentry)
    - [VolumeResponse](#volumeresponse)
    - [VolumeSetRequest](#volumesetrequest)
    - [VolumeSetRequest.OptionsEntry](#volumesetrequestoptionsentry)
    - [VolumeSetResponse](#volumesetresponse)
    - [VolumeSpec](#volumespec)
    - [VolumeSpec.VolumeLabelsEntry](#volumespecvolumelabelsentry)
    - [VolumeSpecSet](#volumespecset)
    - [VolumeSpecSet.VolumeLabelsEntry](#volumespecsetvolumelabelsentry)
    - [VolumeStateAction](#volumestateaction)
  


- Enums
    - [AlertActionType](#alertactiontype)
    - [AttachState](#attachstate)
    - [CloudMigrate.OperationType](#cloudmigrateoperationtype)
    - [CloudMigrate.Stage](#cloudmigratestage)
    - [CloudMigrate.Status](#cloudmigratestatus)
    - [ClusterNotify](#clusternotify)
    - [CosType](#costype)
    - [DriverType](#drivertype)
    - [FSType](#fstype)
    - [GraphDriverChangeType](#graphdriverchangetype)
    - [IoProfile](#ioprofile)
    - [OperationFlags](#operationflags)
    - [ResourceType](#resourcetype)
    - [SdkCloudBackupOpType](#sdkcloudbackupoptype)
    - [SdkCloudBackupRequestedState](#sdkcloudbackuprequestedstate)
    - [SdkCloudBackupStatusType](#sdkcloudbackupstatustype)
    - [SdkTimeWeekday](#sdktimeweekday)
    - [SeverityType](#severitytype)
    - [Status](#status)
    - [StorageMedium](#storagemedium)
    - [VolumeActionParam](#volumeactionparam)
    - [VolumeState](#volumestate)
    - [VolumeStatus](#volumestatus)
  


- [Scalar Value Types](#scalar-value-types)



# OpenStorageCloudBackup {#openstorageapiopenstoragecloudbackup}
OpenStorageCloudBackup service manages backing up volumes to a cloud
location like Amazon, Google, or Azure.

## Create

> **rpc** Create([SdkCloudBackupCreateRequest](#sdkcloudbackupcreaterequest))
    [SdkCloudBackupCreateResponse](#sdkcloudbackupcreateresponse)

Creates a backup request for a specified volume. Use
OpenStorageCloudBackup.Status() to get the current status of the
backup request.
## Restore

> **rpc** Restore([SdkCloudBackupRestoreRequest](#sdkcloudbackuprestorerequest))
    [SdkCloudBackupRestoreResponse](#sdkcloudbackuprestoreresponse)

Restore creates a new volume from a backup id. The newly created volume
has an ha_level (number of replicas) of only 1. To increase the number of
replicas, use OpenStorageVolume.Set() to change the ha_level.
## Delete

> **rpc** Delete([SdkCloudBackupDeleteRequest](#sdkcloudbackupdeleterequest))
    [SdkCloudBackupDeleteResponse](#sdkcloudbackupdeleteresponse)

Delete deletes a backup stored in the cloud. If the backup is an incremental
backup and other backups are dependent on it, it will not be able to be deleted.
## DeleteAll

> **rpc** DeleteAll([SdkCloudBackupDeleteAllRequest](#sdkcloudbackupdeleteallrequest))
    [SdkCloudBackupDeleteAllResponse](#sdkcloudbackupdeleteallresponse)

DeleteAll deletes all the backups in the cloud for the specified volume.
## Enumerate

> **rpc** Enumerate([SdkCloudBackupEnumerateRequest](#sdkcloudbackupenumeraterequest))
    [SdkCloudBackupEnumerateResponse](#sdkcloudbackupenumerateresponse)

Return a list of backups for the specified volume
## Status

> **rpc** Status([SdkCloudBackupStatusRequest](#sdkcloudbackupstatusrequest))
    [SdkCloudBackupStatusResponse](#sdkcloudbackupstatusresponse)

Status returns the status of any cloud backups of a volume
## Catalog

> **rpc** Catalog([SdkCloudBackupCatalogRequest](#sdkcloudbackupcatalogrequest))
    [SdkCloudBackupCatalogResponse](#sdkcloudbackupcatalogresponse)

Catalog returns a list of the contents in the backup
## History

> **rpc** History([SdkCloudBackupHistoryRequest](#sdkcloudbackuphistoryrequest))
    [SdkCloudBackupHistoryResponse](#sdkcloudbackuphistoryresponse)

History returns a list of backups for a specified volume
## StateChange

> **rpc** StateChange([SdkCloudBackupStateChangeRequest](#sdkcloudbackupstatechangerequest))
    [SdkCloudBackupStateChangeResponse](#sdkcloudbackupstatechangeresponse)

StateChange can be used to stop, pause, and restart a backup
## SchedCreate

> **rpc** SchedCreate([SdkCloudBackupSchedCreateRequest](#sdkcloudbackupschedcreaterequest))
    [SdkCloudBackupSchedCreateResponse](#sdkcloudbackupschedcreateresponse)

Create cloud backup schedule
## SchedDelete

> **rpc** SchedDelete([SdkCloudBackupSchedDeleteRequest](#sdkcloudbackupscheddeleterequest))
    [SdkCloudBackupSchedDeleteResponse](#sdkcloudbackupscheddeleteresponse)

Delete cloud backup schedule
## SchedEnumerate

> **rpc** SchedEnumerate([SdkCloudBackupSchedEnumerateRequest](#sdkcloudbackupschedenumeraterequest))
    [SdkCloudBackupSchedEnumerateResponse](#sdkcloudbackupschedenumerateresponse)

Enumerate cloud backup schedules
 <!-- end methods -->
# OpenStorageCluster {#openstorageapiopenstoragecluster}
OpenStorageCluster service provides the methods to manage the cluster

## Enumerate

> **rpc** Enumerate([SdkClusterEnumerateRequest](#sdkclusterenumeraterequest))
    [SdkClusterEnumerateResponse](#sdkclusterenumerateresponse)

Enumerate returns information about the cluster and the unique ids of
all the nodes in the cluster.
## Inspect

> **rpc** Inspect([SdkClusterInspectRequest](#sdkclusterinspectrequest))
    [SdkClusterInspectResponse](#sdkclusterinspectresponse)

Inspect returns information about the specified node
## AlertEnumerate

> **rpc** AlertEnumerate([SdkClusterAlertEnumerateRequest](#sdkclusteralertenumeraterequest))
    [SdkClusterAlertEnumerateResponse](#sdkclusteralertenumerateresponse)

AlertEnumerate returns a list of alerts from the storage cluster
## AlertClear

> **rpc** AlertClear([SdkClusterAlertClearRequest](#sdkclusteralertclearrequest))
    [SdkClusterAlertClearResponse](#sdkclusteralertclearresponse)

AlertClear clears the alert for a given resource
## AlertDelete

> **rpc** AlertDelete([SdkClusterAlertDeleteRequest](#sdkclusteralertdeleterequest))
    [SdkClusterAlertDeleteResponse](#sdkclusteralertdeleteresponse)

AlertDelete deletes an alert for all resources
 <!-- end methods -->
# OpenStorageCredentials {#openstorageapiopenstoragecredentials}
OpenStorageCredentials is a service used to manage the cloud credentials
which can then be used by the OpenStorageCloudBackup service

## Create

> **rpc** Create([SdkCredentialCreateRequest](#sdkcredentialcreaterequest))
    [SdkCredentialCreateResponse](#sdkcredentialcreateresponse)

Create is used to submit cloud credentials. It will return an
id of the credentials once they are verified to work.

##### Example
{% codetabs name="Golang", type="go" -%}
id, err := client.Create(context.Background(), &api.SdkCredentialCreateRequest{
  CredentialType: &api.SdkCredentialCreateRequest_AwsCredential{
    AwsCredential: &api.SdkAwsCredentialRequest{
    AccessKey: "dummy-access",
    SecretKey: "dummy-secret",
    Endpoint:  "dummy-endpoint",
    Region:    "dummy-region",
  },
})
{%- language name="Python", type="py" -%}
en_resp = client.Create(api_pb2.SdkCredentialCreateRequest(
  aws_credential=api_pb2.SdkAwsCredentialRequest(
    access_key='dummy-access',
    secret_key='dumm-secret',
    endpoint='dummy-endpoint',
    region='dummy-region')))
{%- endcodetabs %}
## Enumerate

> **rpc** Enumerate([SdkCredentialEnumerateRequest](#sdkcredentialenumeraterequest))
    [SdkCredentialEnumerateResponse](#sdkcredentialenumerateresponse)

Enumerate returns a list of credential ids
## Inspect

> **rpc** Inspect([SdkCredentialInspectRequest](#sdkcredentialinspectrequest))
    [SdkCredentialInspectResponse](#sdkcredentialinspectresponse)

Inspect returns the information about a credential, but does not return the secret key.
## Delete

> **rpc** Delete([SdkCredentialDeleteRequest](#sdkcredentialdeleterequest))
    [SdkCredentialDeleteResponse](#sdkcredentialdeleteresponse)

Delete a specified credential
## Validate

> **rpc** Validate([SdkCredentialValidateRequest](#sdkcredentialvalidaterequest))
    [SdkCredentialValidateResponse](#sdkcredentialvalidateresponse)

Validate is used to validate credentials
 <!-- end methods -->
# OpenStorageObjectstore {#openstorageapiopenstorageobjectstore}
OpenStorageObjectstore is a service used to manage object store services on volumes

## Inspect

> **rpc** Inspect([SdkObjectstoreInspectRequest](#sdkobjectstoreinspectrequest))
    [SdkObjectstoreInspectResponse](#sdkobjectstoreinspectresponse)

Inspect returns information about the object store endpoint
## Create

> **rpc** Create([SdkObjectstoreCreateRequest](#sdkobjectstorecreaterequest))
    [SdkObjectstoreCreateResponse](#sdkobjectstorecreateresponse)

Creates creates an object store endpoint on specified volume
## Delete

> **rpc** Delete([SdkObjectstoreDeleteRequest](#sdkobjectstoredeleterequest))
    [SdkObjectstoreDeleteResponse](#sdkobjectstoredeleteresponse)

Delete destroys the object store endpoint on the volume
## Update

> **rpc** Update([SdkObjectstoreUpdateRequest](#sdkobjectstoreupdaterequest))
    [SdkObjectstoreUpdateResponse](#sdkobjectstoreupdateresponse)

Updates provided objectstore status.
This call can be used to stop and start the server while maintaining the same
object storage id.
 <!-- end methods -->
# OpenStorageSchedulePolicy {#openstorageapiopenstorageschedulepolicy}
OpenStorageSchedulePolicy service is used to manage the automated
snapshots for a volume

## Create

> **rpc** Create([SdkSchedulePolicyCreateRequest](#sdkschedulepolicycreaterequest))
    [SdkSchedulePolicyCreateResponse](#sdkschedulepolicycreateresponse)

Create creates a new snapshot schedule. They can be setup daily,
weekly, or monthly.
## Update

> **rpc** Update([SdkSchedulePolicyUpdateRequest](#sdkschedulepolicyupdaterequest))
    [SdkSchedulePolicyUpdateResponse](#sdkschedulepolicyupdateresponse)

Update a snapshot schedule
## Enumerate

> **rpc** Enumerate([SdkSchedulePolicyEnumerateRequest](#sdkschedulepolicyenumeraterequest))
    [SdkSchedulePolicyEnumerateResponse](#sdkschedulepolicyenumerateresponse)

Enumerate returns a list of schedules
## Inspect

> **rpc** Inspect([SdkSchedulePolicyInspectRequest](#sdkschedulepolicyinspectrequest))
    [SdkSchedulePolicyInspectResponse](#sdkschedulepolicyinspectresponse)

Inspect returns information about a specified schedule
## Delete

> **rpc** Delete([SdkSchedulePolicyDeleteRequest](#sdkschedulepolicydeleterequest))
    [SdkSchedulePolicyDeleteResponse](#sdkschedulepolicydeleteresponse)

Delete removes a snapshot schedule
 <!-- end methods -->
# OpenStorageVolume {#openstorageapiopenstoragevolume}
OpenStorageVolume is a service used to manage the volumes of a storage system

## Create

> **rpc** Create([SdkVolumeCreateRequest](#sdkvolumecreaterequest))
    [SdkVolumeCreateResponse](#sdkvolumecreateresponse)

Create creates a volume according to the specification provided

##### Example
{% codetabs name="Golang", type="go" -%}
id, err := client.Create(context.Background(), &api.SdkVolumeCreateRequest{
  Name: "volume-12345-east",
  Spec: &api.VolumeSpec {
    Size: 1234567,
  },
})
{%- language name="Python", type="py" -%}
en_resp = client.Create(api_pb2.SdkVolumeCreateRequest(
  name="volume-12345-east",
  spec=api_pb2.VolumeSpec(size=1234567)))
{%- endcodetabs %}
## Clone

> **rpc** Clone([SdkVolumeCloneRequest](#sdkvolumeclonerequest))
    [SdkVolumeCloneResponse](#sdkvolumecloneresponse)

Clone creates a new writable volume cloned from an existing volume
## Delete

> **rpc** Delete([SdkVolumeDeleteRequest](#sdkvolumedeleterequest))
    [SdkVolumeDeleteResponse](#sdkvolumedeleteresponse)

Delete deletes the provided volume
## Inspect

> **rpc** Inspect([SdkVolumeInspectRequest](#sdkvolumeinspectrequest))
    [SdkVolumeInspectResponse](#sdkvolumeinspectresponse)

Inspect returns information about a volume
## Set

> **rpc** Set([SdkVolumeSetRequest](#sdkvolumesetrequest))
    [SdkVolumeSetResponse](#sdkvolumesetresponse)

Set provides a method for manipulating the specification and attributes of a volume.
Set can be used to resize a volume, update labels, change replica count, and much more.
## Enumerate

> **rpc** Enumerate([SdkVolumeEnumerateRequest](#sdkvolumeenumeraterequest))
    [SdkVolumeEnumerateResponse](#sdkvolumeenumerateresponse)

Enumerate returns a list of volume ids that match the labels if any are provided.
## SnapshotCreate

> **rpc** SnapshotCreate([SdkVolumeSnapshotCreateRequest](#sdkvolumesnapshotcreaterequest))
    [SdkVolumeSnapshotCreateResponse](#sdkvolumesnapshotcreateresponse)

SnapshotCreate creates a snapshot of a volume. This creates an immutable (read-only),
point-in-time snapshot of a volume.
## SnapshotRestore

> **rpc** SnapshotRestore([SdkVolumeSnapshotRestoreRequest](#sdkvolumesnapshotrestorerequest))
    [SdkVolumeSnapshotRestoreResponse](#sdkvolumesnapshotrestoreresponse)

SnapshotRestore restores a volume to a specified snapshot
## SnapshotEnumerate

> **rpc** SnapshotEnumerate([SdkVolumeSnapshotEnumerateRequest](#sdkvolumesnapshotenumeraterequest))
    [SdkVolumeSnapshotEnumerateResponse](#sdkvolumesnapshotenumerateresponse)

SnapshotEnumerate returns a list of snapshots for a specific volume
that match the labels provided if any.
## Attach

> **rpc** Attach([SdkVolumeAttachRequest](#sdkvolumeattachrequest))
    [SdkVolumeAttachResponse](#sdkvolumeattachresponse)

Attach attaches device to the host that the client is communicating with.
NOTE: Please see [#381](https://github.com/libopenstorage/openstorage/issues/381) for more
information about a new feature to allow attachment to any node.
## Detach

> **rpc** Detach([SdkVolumeDetachRequest](#sdkvolumedetachrequest))
    [SdkVolumeDetachResponse](#sdkvolumedetachresponse)

Detaches a the volume from the host that the client is communicating with
NOTE: Please see [#381](https://github.com/libopenstorage/openstorage/issues/381) for more
information about a new feature to allow attachment to any node.
## Mount

> **rpc** Mount([SdkVolumeMountRequest](#sdkvolumemountrequest))
    [SdkVolumeMountResponse](#sdkvolumemountresponse)

Mount mounts an attached volume in the host that the client is communicating with
NOTE: Please see [#381](https://github.com/libopenstorage/openstorage/issues/381) for more
information about a new feature to allow attachment to any node.
## Unmount

> **rpc** Unmount([SdkVolumeUnmountRequest](#sdkvolumeunmountrequest))
    [SdkVolumeUnmountResponse](#sdkvolumeunmountresponse)

Unmount unmounts a mounted volume in the host that the client is communicating with
NOTE: Please see [#381](https://github.com/libopenstorage/openstorage/issues/381) for more
information about a new feature to allow attachment to any node.
 <!-- end methods -->
 <!-- end services -->

# Messages


## ActiveRequest {#activerequest}
Active Request
swagger:model


| Field | Type | Description |
| ----- | ---- | ----------- |
| ReqestKV | [map ActiveRequest.ReqestKVEntry](#activerequestreqestkventry) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## ActiveRequest.ReqestKVEntry {#activerequestreqestkventry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ int64](#int64) | none |
| value | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## ActiveRequests {#activerequests}
Active Requests
swagger:model


| Field | Type | Description |
| ----- | ---- | ----------- |
| RequestCount | [ int64](#int64) | none |
| ActiveRequest | [repeated ActiveRequest](#activerequest) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## Alert {#alert}
Alert is a structure that represents an alert object
swagger:model


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](#int64) | Id for Alert |
| severity | [ SeverityType](#severitytype) | Severity of the Alert |
| alert_type | [ int64](#int64) | AlertType user defined alert type |
| message | [ string](#string) | Message describing the Alert |
| timestamp | [ google.protobuf.Timestamp](#googleprotobuftimestamp) | Timestamp when Alert occured |
| resource_id | [ string](#string) | ResourceId where Alert occured |
| resource | [ ResourceType](#resourcetype) | Resource where Alert occured |
| cleared | [ bool](#bool) | Cleared Flag |
| ttl | [ uint64](#uint64) | Time-to-live in seconds for this Alert |
| unique_tag | [ string](#string) | UniqueTag helps identify a unique alert for a given resouce |
| count | [ int64](#int64) | Count of such alerts raised so far. |
| first_seen | [ google.protobuf.Timestamp](#googleprotobuftimestamp) | Timestamp when such alert was raised the very first time. |
 <!-- end Fields -->
 <!-- end HasFields -->


## Alerts {#alerts}
Alerts is an array of Alert objects
swagger:model


| Field | Type | Description |
| ----- | ---- | ----------- |
| alert | [repeated Alert](#alert) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## CloudMigrate {#cloudmigrate}


 <!-- end HasFields -->


## CloudMigrateCancelRequest {#cloudmigratecancelrequest}
Request to stop a cloud migration


| Field | Type | Description |
| ----- | ---- | ----------- |
| operation | [ CloudMigrate.OperationType](#cloudmigrateoperationtype) | The type of operation to cancel |
| cluster_id | [ string](#string) | ID of the cluster to which migration should be cancelled |
| target_id | [ string](#string) | Depending on the operation type this can be a VolumeID or VolumeGroupID |
 <!-- end Fields -->
 <!-- end HasFields -->


## CloudMigrateInfo {#cloudmigrateinfo}



| Field | Type | Description |
| ----- | ---- | ----------- |
| cluster_id | [ string](#string) | ID of the cluster where the volume is being migrated |
| local_volume_id | [ string](#string) | ID of the volume on the local cluster |
| local_volume_name | [ string](#string) | Name of the volume on the local cluster |
| remote_volume_id | [ string](#string) | ID of the volume on the remote cluster |
| cloudbackup_id | [ string](#string) | ID of the cloudbackup used for the migration |
| current_stage | [ CloudMigrate.Stage](#cloudmigratestage) | Current stage of the volume migration |
| status | [ CloudMigrate.Status](#cloudmigratestatus) | Status of the current stage |
| last_update | [ google.protobuf.Timestamp](#googleprotobuftimestamp) | Last time the status was updated |
| last_success | [ google.protobuf.Timestamp](#googleprotobuftimestamp) | Time of the last successful migration of the volume |
 <!-- end Fields -->
 <!-- end HasFields -->


## CloudMigrateInfoList {#cloudmigrateinfolist}



| Field | Type | Description |
| ----- | ---- | ----------- |
| list | [repeated CloudMigrateInfo](#cloudmigrateinfo) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## CloudMigrateStartRequest {#cloudmigratestartrequest}
Request to start a cloud migration


| Field | Type | Description |
| ----- | ---- | ----------- |
| operation | [ CloudMigrate.OperationType](#cloudmigrateoperationtype) | The type of operation to start |
| cluster_id | [ string](#string) | ID of the cluster to which volumes are to be migrated |
| target_id | [ string](#string) | Depending on the operation type this can be a VolumeID or VolumeGroupID |
 <!-- end Fields -->
 <!-- end HasFields -->


## CloudMigrateStatusResponse {#cloudmigratestatusresponse}
Response with a status of the cloud migration operations


| Field | Type | Description |
| ----- | ---- | ----------- |
| info | [map CloudMigrateStatusResponse.InfoEntry](#cloudmigratestatusresponseinfoentry) | Map of cluster id to the status of volumes being migrated |
 <!-- end Fields -->
 <!-- end HasFields -->


## CloudMigrateStatusResponse.InfoEntry {#cloudmigratestatusresponseinfoentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ CloudMigrateInfoList](#cloudmigrateinfolist) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## ClusterResponse {#clusterresponse}
ClusterResponse specifies a response that gets returned when requesting the cluster
swagger:response clusterResponse


| Field | Type | Description |
| ----- | ---- | ----------- |
| error | [ string](#string) | Error code

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
| path | [ string](#string) | none |
| kind | [ GraphDriverChangeType](#graphdriverchangetype) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## Group {#group}
Group represents VolumeGroup / namespace
All volumes in the same group share this object.
swagger:model


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ string](#string) | Id common identifier across volumes that have the same group. |
 <!-- end Fields -->
 <!-- end HasFields -->


## GroupSnapCreateRequest {#groupsnapcreaterequest}
GroupSnapCreateRequest specifies a request to create a snapshot of given group.
swagger:model


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ string](#string) | none |
| Labels | [map GroupSnapCreateRequest.LabelsEntry](#groupsnapcreaterequestlabelsentry) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## GroupSnapCreateRequest.LabelsEntry {#groupsnapcreaterequestlabelsentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## GroupSnapCreateResponse {#groupsnapcreateresponse}
GroupSnapCreateRequest specifies a response that get's returned when creating a group snapshot.
swagger:response groupSnapCreateResponse


| Field | Type | Description |
| ----- | ---- | ----------- |
| snapshots | [map GroupSnapCreateResponse.SnapshotsEntry](#groupsnapcreateresponsesnapshotsentry) | Created snapshots

in: body Required: true |
| error | [ string](#string) | Error message

in: body Required: true |
 <!-- end Fields -->
 <!-- end HasFields -->


## GroupSnapCreateResponse.SnapshotsEntry {#groupsnapcreateresponsesnapshotsentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ SnapCreateResponse](#snapcreateresponse) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## ObjectstoreInfo {#objectstoreinfo}
ObjectstoreInfo is a structure that has current objectstore info
swagger:model


| Field | Type | Description |
| ----- | ---- | ----------- |
| uuid | [ string](#string) | UUID of objectstore |
| volume_id | [ string](#string) | VolumeID of volume used by object store |
| enabled | [ bool](#bool) | Enable/Disable created objectstore |
| status | [ string](#string) | Status of objectstore running/failed |
| action | [ int64](#int64) | Action being taken on this objectstore |
| access_key | [ string](#string) | AccessKey for login into objectstore |
| secret_key | [ string](#string) | SecretKey for login into objectstore |
| endpoints | [repeated string](#string) | Endpoints for accessing objectstore |
| current_endpoint | [ string](#string) | CurrentEndpoint on which objectstore server is accessible |
| access_port | [ int64](#int64) | AccessPort is objectstore server port |
| region | [ string](#string) | Region for this objectstore |
 <!-- end Fields -->
 <!-- end HasFields -->


## ReplicaSet {#replicaset}
ReplicaSet set of machine IDs (nodes) to which part of this volume is erasure
coded - for clustered storage arrays
swagger:model


| Field | Type | Description |
| ----- | ---- | ----------- |
| nodes | [repeated string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## RuntimeStateMap {#runtimestatemap}
RuntimeStateMap is a list of name value mapping of driver specific runtime
information.
swagger:model


| Field | Type | Description |
| ----- | ---- | ----------- |
| runtime_state | [map RuntimeStateMap.RuntimeStateEntry](#runtimestatemapruntimestateentry) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## RuntimeStateMap.RuntimeStateEntry {#runtimestatemapruntimestateentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkAwsCredentialRequest {#sdkawscredentialrequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| access_key | [ string](#string) | Access key |
| secret_key | [ string](#string) | Secret key |
| endpoint | [ string](#string) | Endpoint |
| region | [ string](#string) | Region |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkAwsCredentialResponse {#sdkawscredentialresponse}



| Field | Type | Description |
| ----- | ---- | ----------- |
| credential_id | [ string](#string) | Credential Id |
| access_key | [ string](#string) | Access key |
| endpoint | [ string](#string) | Endpoint |
| region | [ string](#string) | Region |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkAzureCredentialRequest {#sdkazurecredentialrequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| account_name | [ string](#string) | Account name |
| account_key | [ string](#string) | Account key |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkAzureCredentialResponse {#sdkazurecredentialresponse}



| Field | Type | Description |
| ----- | ---- | ----------- |
| credential_id | [ string](#string) | Credential Id |
| account_name | [ string](#string) | Account name |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupCatalogRequest {#sdkcloudbackupcatalogrequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| backup_id | [ string](#string) | Id of the backup |
| credential_uuid | [ string](#string) | is the credential for cloud |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupCatalogResponse {#sdkcloudbackupcatalogresponse}



| Field | Type | Description |
| ----- | ---- | ----------- |
| contents | [repeated string](#string) | Contents is listing of backup contents |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupCreateRequest {#sdkcloudbackupcreaterequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_id | [ string](#string) | VolumeID of the volume for which cloudbackup is requested |
| credential_uuid | [ string](#string) | CredentialUUID is cloud credential to be used for backup |
| full | [ bool](#bool) | Full indicates if full backup is desired even though incremental is possible |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupCreateResponse {#sdkcloudbackupcreateresponse}


 <!-- end HasFields -->


## SdkCloudBackupDeleteAllRequest {#sdkcloudbackupdeleteallrequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| src_volume_id | [ string](#string) | id of the volume for the request |
| credential_uuid | [ string](#string) | CredentialUUID is the credential for cloud to be used for the request |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupDeleteAllResponse {#sdkcloudbackupdeleteallresponse}


 <!-- end HasFields -->


## SdkCloudBackupDeleteRequest {#sdkcloudbackupdeleterequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| backup_id | [ string](#string) | ID is the ID of the cloud backup |
| credential_uuid | [ string](#string) | CredentialUUID is the credential for cloud to be used for the request |
| force | [ bool](#bool) | Force Delete cloudbackup even if there are dependencies |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupDeleteResponse {#sdkcloudbackupdeleteresponse}


 <!-- end HasFields -->


## SdkCloudBackupEnumerateRequest {#sdkcloudbackupenumeraterequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| src_volume_id | [ string](#string) | Optional source id of the volume for the request |
| cluster_id | [ string](#string) | ClusterID is the optional clusterID for the request |
| credential_uuid | [ string](#string) | CredentialUUID is the credential for cloud to be used for the request |
| all | [ bool](#bool) | All if set to true, backups for all clusters in the cloud are processed |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupEnumerateResponse {#sdkcloudbackupenumerateresponse}



| Field | Type | Description |
| ----- | ---- | ----------- |
| backups | [repeated SdkCloudBackupInfo](#sdkcloudbackupinfo) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupHistoryItem {#sdkcloudbackuphistoryitem}



| Field | Type | Description |
| ----- | ---- | ----------- |
| src_volume_id | [ string](#string) | SrcVolumeID is volume ID which was backedup |
| timestamp | [ google.protobuf.Timestamp](#googleprotobuftimestamp) | TimeStamp is the time at which either backup completed/failed |
| status | [ SdkCloudBackupStatusType](#sdkcloudbackupstatustype) | Status indicates whether backup was completed/failed |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupHistoryRequest {#sdkcloudbackuphistoryrequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| src_volume_id | [ string](#string) | volumeID for which history of backup/restore is being requested (optional) If not provided, it will return the history for all volumes. |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupHistoryResponse {#sdkcloudbackuphistoryresponse}



| Field | Type | Description |
| ----- | ---- | ----------- |
| history_list | [repeated SdkCloudBackupHistoryItem](#sdkcloudbackuphistoryitem) | HistoryList is list of past backup/restores in the cluster |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupInfo {#sdkcloudbackupinfo}



| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ string](#string) | ID is the ID of the cloud backup |
| src_volume_id | [ string](#string) | Source volumeID of the backup |
| src_volume_name | [ string](#string) | Name of the sourceVolume of the backup |
| timestamp | [ google.protobuf.Timestamp](#googleprotobuftimestamp) | Timestamp is the timestamp at which the source volume was backed up to cloud |
| metadata | [map SdkCloudBackupInfo.MetadataEntry](#sdkcloudbackupinfometadataentry) | Metadata associated with the backup |
| status | [ SdkCloudBackupStatusType](#sdkcloudbackupstatustype) | Status indicates the status of the backup |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupInfo.MetadataEntry {#sdkcloudbackupinfometadataentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupRestoreRequest {#sdkcloudbackuprestorerequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| backup_id | [ string](#string) | Backup ID being restored |
| restore_volume_name | [ string](#string) | Optional volume Name of the new volume to be created in the cluster for restoring the cloudbackup |
| credential_uuid | [ string](#string) | The credential to be used for restore operation |
| node_id | [ string](#string) | Optional for provisioning restore volume (ResoreVolumeName should not be specified) |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupRestoreResponse {#sdkcloudbackuprestoreresponse}



| Field | Type | Description |
| ----- | ---- | ----------- |
| restore_volume_id | [ string](#string) | VolumeID to which the backup is being restored |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupSchedCreateRequest {#sdkcloudbackupschedcreaterequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| cloud_sched_info | [ SdkCloudBackupScheduleInfo](#sdkcloudbackupscheduleinfo) | Cloud Backup Schedule info |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupSchedCreateResponse {#sdkcloudbackupschedcreateresponse}



| Field | Type | Description |
| ----- | ---- | ----------- |
| uuid | [ string](#string) | UUID of newly created backup schedule |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupSchedDeleteRequest {#sdkcloudbackupscheddeleterequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| uuid | [ string](#string) | UUID of cloud backup to delete |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupSchedDeleteResponse {#sdkcloudbackupscheddeleteresponse}


 <!-- end HasFields -->


## SdkCloudBackupSchedEnumerateRequest {#sdkcloudbackupschedenumeraterequest}


 <!-- end HasFields -->


## SdkCloudBackupSchedEnumerateResponse {#sdkcloudbackupschedenumerateresponse}



| Field | Type | Description |
| ----- | ---- | ----------- |
| cloud_sched_list | [map SdkCloudBackupSchedEnumerateResponse.CloudSchedListEntry](#sdkcloudbackupschedenumerateresponsecloudschedlistentry) | Returns list of backup schedules |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupSchedEnumerateResponse.CloudSchedListEntry {#sdkcloudbackupschedenumerateresponsecloudschedlistentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ SdkCloudBackupScheduleInfo](#sdkcloudbackupscheduleinfo) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupScheduleInfo {#sdkcloudbackupscheduleinfo}



| Field | Type | Description |
| ----- | ---- | ----------- |
| src_volume_id | [ string](#string) | SrcVolumeID is the schedule's source volume |
| credential_uuid | [ string](#string) | CredentialUUID is the cloud credential used with this schedule |
| schedule | [ SdkSchedulePolicyInterval](#sdkschedulepolicyinterval) | Schedule is the frequence of backup |
| max_backups | [ uint64](#uint64) | MaxBackups are the maximum number of backups retained in cloud.Older backups are deleted |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupStateChangeRequest {#sdkcloudbackupstatechangerequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| src_volume_id | [ string](#string) | SrcVolumeID is volume ID on which backup/restore state change is being requested |
| requested_state | [ SdkCloudBackupRequestedState](#sdkcloudbackuprequestedstate) | RequestedState is desired state of the op can be pause/resume/stop |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupStateChangeResponse {#sdkcloudbackupstatechangeresponse}


 <!-- end HasFields -->


## SdkCloudBackupStatus {#sdkcloudbackupstatus}



| Field | Type | Description |
| ----- | ---- | ----------- |
| backup_id | [ string](#string) | ID is the ID for the operation |
| optype | [ SdkCloudBackupOpType](#sdkcloudbackupoptype) | OpType indicates if this is a backup or restore |
| status | [ SdkCloudBackupStatusType](#sdkcloudbackupstatustype) | State indicates if the op is currently active/done/failed |
| bytes_done | [ uint64](#uint64) | BytesDone indicates total Bytes uploaded/downloaded |
| start_time | [ google.protobuf.Timestamp](#googleprotobuftimestamp) | StartTime indicates Op's start time |
| completed_time | [ google.protobuf.Timestamp](#googleprotobuftimestamp) | CompletedTime indicates Op's completed time |
| node_id | [ string](#string) | NodeID is the ID of the node where this Op is active |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupStatusRequest {#sdkcloudbackupstatusrequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| src_volume_id | [ string](#string) | SrcVolumeID optional volumeID to list status of backup/restore |
| local | [ bool](#bool) | Local indicates if only those backups/restores that are active on current node must be returned |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupStatusResponse {#sdkcloudbackupstatusresponse}



| Field | Type | Description |
| ----- | ---- | ----------- |
| statuses | [map SdkCloudBackupStatusResponse.StatusesEntry](#sdkcloudbackupstatusresponsestatusesentry) | statuses is list of currently active/failed/done backup/restores |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupStatusResponse.StatusesEntry {#sdkcloudbackupstatusresponsestatusesentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ SdkCloudBackupStatus](#sdkcloudbackupstatus) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkClusterAlertClearRequest {#sdkclusteralertclearrequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| resource | [ ResourceType](#resourcetype) | Type of resource (required) |
| alert_id | [ int64](#int64) | Id of alert as returned by ClusterEnumerateAlertResponse (required) |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkClusterAlertClearResponse {#sdkclusteralertclearresponse}


 <!-- end HasFields -->


## SdkClusterAlertDeleteRequest {#sdkclusteralertdeleterequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| resource | [ ResourceType](#resourcetype) | Type of resource (required) |
| alert_id | [ int64](#int64) | Id of alert as returned by ClusterEnumerateAlertResponse (required) |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkClusterAlertDeleteResponse {#sdkclusteralertdeleteresponse}


 <!-- end HasFields -->


## SdkClusterAlertEnumerateRequest {#sdkclusteralertenumeraterequest}
This request contains the information needed to get alerts from
the storage system


| Field | Type | Description |
| ----- | ---- | ----------- |
| time_start | [ google.protobuf.Timestamp](#googleprotobuftimestamp) | Start time of alerts |
| time_end | [ google.protobuf.Timestamp](#googleprotobuftimestamp) | End time of alerts |
| resource | [ ResourceType](#resourcetype) | Type of resource |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkClusterAlertEnumerateResponse {#sdkclusteralertenumerateresponse}
This response provides a list of the alerts at the specified time.


| Field | Type | Description |
| ----- | ---- | ----------- |
| alerts | [repeated Alert](#alert) | Information on the alerts requested |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkClusterEnumerateRequest {#sdkclusterenumeraterequest}
This request is an empty request

 <!-- end HasFields -->


## SdkClusterEnumerateResponse {#sdkclusterenumerateresponse}
This response returns information about the cluster


| Field | Type | Description |
| ----- | ---- | ----------- |
| cluster | [ StorageCluster](#storagecluster) | Cluster information |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkClusterInspectRequest {#sdkclusterinspectrequest}
Contains the request when inspecting a node


| Field | Type | Description |
| ----- | ---- | ----------- |
| node_id | [ string](#string) | Id of node to inspect (required) |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkClusterInspectResponse {#sdkclusterinspectresponse}
This response contains information about the respective node requested


| Field | Type | Description |
| ----- | ---- | ----------- |
| node | [ StorageNode](#storagenode) | Node information |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCredentialCreateRequest {#sdkcredentialcreaterequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) credential_type.aws_credential | [ SdkAwsCredentialRequest](#sdkawscredentialrequest) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) credential_type.azure_credential | [ SdkAzureCredentialRequest](#sdkazurecredentialrequest) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) credential_type.google_credential | [ SdkGoogleCredentialRequest](#sdkgooglecredentialrequest) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCredentialCreateResponse {#sdkcredentialcreateresponse}



| Field | Type | Description |
| ----- | ---- | ----------- |
| credential_id | [ string](#string) | Id of the credentials |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCredentialDeleteRequest {#sdkcredentialdeleterequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| credential_id | [ string](#string) | ID for credentials |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCredentialDeleteResponse {#sdkcredentialdeleteresponse}


 <!-- end HasFields -->


## SdkCredentialEnumerateRequest {#sdkcredentialenumeraterequest}


 <!-- end HasFields -->


## SdkCredentialEnumerateResponse {#sdkcredentialenumerateresponse}



| Field | Type | Description |
| ----- | ---- | ----------- |
| credential_ids | [repeated string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCredentialInspectRequest {#sdkcredentialinspectrequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| credential_id | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCredentialInspectResponse {#sdkcredentialinspectresponse}
This response uses OneOf proto style. Depending on your programming language
you will need to check if the value of credential_type is one of the ones below.


| Field | Type | Description |
| ----- | ---- | ----------- |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) credential_type.aws_credential | [ SdkAwsCredentialResponse](#sdkawscredentialresponse) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) credential_type.azure_credential | [ SdkAzureCredentialResponse](#sdkazurecredentialresponse) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) credential_type.google_credential | [ SdkGoogleCredentialResponse](#sdkgooglecredentialresponse) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCredentialValidateRequest {#sdkcredentialvalidaterequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| credential_id | [ string](#string) | Id of the credentials |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCredentialValidateResponse {#sdkcredentialvalidateresponse}


 <!-- end HasFields -->


## SdkGoogleCredentialRequest {#sdkgooglecredentialrequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| project_id | [ string](#string) | Project ID |
| json_key | [ string](#string) | JSON Key |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkGoogleCredentialResponse {#sdkgooglecredentialresponse}



| Field | Type | Description |
| ----- | ---- | ----------- |
| credential_id | [ string](#string) | Credential Id |
| project_id | [ string](#string) | Project ID |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkObjectstoreCreateRequest {#sdkobjectstorecreaterequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_id | [ string](#string) | Volume on which objectstore will be running |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkObjectstoreCreateResponse {#sdkobjectstorecreateresponse}



| Field | Type | Description |
| ----- | ---- | ----------- |
| objectstore_status | [ ObjectstoreInfo](#objectstoreinfo) | Created objecstore status |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkObjectstoreDeleteRequest {#sdkobjectstoredeleterequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| objectstore_id | [ string](#string) | Objectstore ID to delete |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkObjectstoreDeleteResponse {#sdkobjectstoredeleteresponse}


 <!-- end HasFields -->


## SdkObjectstoreInspectRequest {#sdkobjectstoreinspectrequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| objectstore_id | [ string](#string) | ObjecstoreID to query objestore status |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkObjectstoreInspectResponse {#sdkobjectstoreinspectresponse}



| Field | Type | Description |
| ----- | ---- | ----------- |
| objectstore_status | [ ObjectstoreInfo](#objectstoreinfo) | Objectstore status |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkObjectstoreUpdateRequest {#sdkobjectstoreupdaterequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| objectstore_id | [ string](#string) | Objectstore Id to update |
| enable | [ bool](#bool) | enable/disable objectstore |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkObjectstoreUpdateResponse {#sdkobjectstoreupdateresponse}


 <!-- end HasFields -->


## SdkSchedulePolicy {#sdkschedulepolicy}



| Field | Type | Description |
| ----- | ---- | ----------- |
| name | [ string](#string) | Name of the schedule policy |
| schedule | [ SdkSchedulePolicyInterval](#sdkschedulepolicyinterval) | Schedule policy |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkSchedulePolicyCreateRequest {#sdkschedulepolicycreaterequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| SchedulePolicy | [ SdkSchedulePolicy](#sdkschedulepolicy) | Schedule Policy |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkSchedulePolicyCreateResponse {#sdkschedulepolicycreateresponse}


 <!-- end HasFields -->


## SdkSchedulePolicyDeleteRequest {#sdkschedulepolicydeleterequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| name | [ string](#string) | Name of the schedule policy |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkSchedulePolicyDeleteResponse {#sdkschedulepolicydeleteresponse}


 <!-- end HasFields -->


## SdkSchedulePolicyEnumerateRequest {#sdkschedulepolicyenumeraterequest}


 <!-- end HasFields -->


## SdkSchedulePolicyEnumerateResponse {#sdkschedulepolicyenumerateresponse}



| Field | Type | Description |
| ----- | ---- | ----------- |
| policies | [repeated SdkSchedulePolicy](#sdkschedulepolicy) | List of Schedule Policy |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkSchedulePolicyInspectRequest {#sdkschedulepolicyinspectrequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| name | [ string](#string) | Name of the schedule Policy |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkSchedulePolicyInspectResponse {#sdkschedulepolicyinspectresponse}



| Field | Type | Description |
| ----- | ---- | ----------- |
| policy | [ SdkSchedulePolicy](#sdkschedulepolicy) | List of Schedule Policy |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkSchedulePolicyInterval {#sdkschedulepolicyinterval}



| Field | Type | Description |
| ----- | ---- | ----------- |
| retain | [ int64](#int64) | Number of instances to retain |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) period_type.daily | [ SdkSchedulePolicyIntervalDaily](#sdkschedulepolicyintervaldaily) | Daily policy |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) period_type.weekly | [ SdkSchedulePolicyIntervalWeekly](#sdkschedulepolicyintervalweekly) | Weekly policy |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) period_type.monthly | [ SdkSchedulePolicyIntervalMonthly](#sdkschedulepolicyintervalmonthly) | Monthly policy |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkSchedulePolicyIntervalDaily {#sdkschedulepolicyintervaldaily}



| Field | Type | Description |
| ----- | ---- | ----------- |
| hour | [ int32](#int32) | Range: 0-23 |
| minute | [ int32](#int32) | Range: 0-59 |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkSchedulePolicyIntervalMonthly {#sdkschedulepolicyintervalmonthly}



| Field | Type | Description |
| ----- | ---- | ----------- |
| day | [ int32](#int32) | Range: 1-28 |
| hour | [ int32](#int32) | Range: 0-59 |
| minute | [ int32](#int32) | Range: 0-59 |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkSchedulePolicyIntervalWeekly {#sdkschedulepolicyintervalweekly}



| Field | Type | Description |
| ----- | ---- | ----------- |
| day | [ SdkTimeWeekday](#sdktimeweekday) | none |
| hour | [ int32](#int32) | Range: 0-23 |
| minute | [ int32](#int32) | Range: 0-59 |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkSchedulePolicyUpdateRequest {#sdkschedulepolicyupdaterequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| SchedulePolicy | [ SdkSchedulePolicy](#sdkschedulepolicy) | Schedule Policy |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkSchedulePolicyUpdateResponse {#sdkschedulepolicyupdateresponse}


 <!-- end HasFields -->


## SdkVolumeAttachRequest {#sdkvolumeattachrequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_id | [ string](#string) | Id of volume |
| options | [map SdkVolumeAttachRequest.OptionsEntry](#sdkvolumeattachrequestoptionsentry) | Options for attaching volume, right now only passphrase options is supported |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeAttachRequest.OptionsEntry {#sdkvolumeattachrequestoptionsentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeAttachResponse {#sdkvolumeattachresponse}



| Field | Type | Description |
| ----- | ---- | ----------- |
| device_path | [ string](#string) | Device path where device is exported |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeCloneRequest {#sdkvolumeclonerequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| name | [ string](#string) | Unique name of the volume. This will be used for idempotency. |
| parent_id | [ string](#string) | Parent volume id, if specified will create a new volume as a clone of the parent. |
| spec | [ VolumeSpec](#volumespec) | Volume specification |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeCloneResponse {#sdkvolumecloneresponse}



| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_id | [ string](#string) | Id of new volume |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeCreateRequest {#sdkvolumecreaterequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| name | [ string](#string) | Unique name of the volume. This will be used for idempotency. |
| spec | [ VolumeSpec](#volumespec) | Volume specification |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeCreateResponse {#sdkvolumecreateresponse}



| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_id | [ string](#string) | Id of new volume |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeDeleteRequest {#sdkvolumedeleterequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_id | [ string](#string) | Id of volume to delete |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeDeleteResponse {#sdkvolumedeleteresponse}


 <!-- end HasFields -->


## SdkVolumeDetachRequest {#sdkvolumedetachrequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_id | [ string](#string) | Id of the volume |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeDetachResponse {#sdkvolumedetachresponse}


 <!-- end HasFields -->


## SdkVolumeEnumerateRequest {#sdkvolumeenumeraterequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| locator | [ VolumeLocator](#volumelocator) | Volumes to match to this locator. If not provided, all volumes will be returned. |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeEnumerateResponse {#sdkvolumeenumerateresponse}



| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_ids | [repeated string](#string) | List of volumes matching label |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeInspectRequest {#sdkvolumeinspectrequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_id | [ string](#string) | Id of volume to inspect |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeInspectResponse {#sdkvolumeinspectresponse}



| Field | Type | Description |
| ----- | ---- | ----------- |
| volume | [ Volume](#volume) | Information about the volume |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeMountRequest {#sdkvolumemountrequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_id | [ string](#string) | Id of the volume |
| mount_path | [ string](#string) | Mount path for mounting the volume. |
| options | [map SdkVolumeMountRequest.OptionsEntry](#sdkvolumemountrequestoptionsentry) | Additional options |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeMountRequest.OptionsEntry {#sdkvolumemountrequestoptionsentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeMountResponse {#sdkvolumemountresponse}


 <!-- end HasFields -->


## SdkVolumeSetRequest {#sdkvolumesetrequest}
This request is used to adjust or set new values in the volume


| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_id | [ string](#string) | Id of the volume to update |
| locator | [ VolumeLocator](#volumelocator) | Change locator values. Some of these values may not be able to be changed. New labels will be added to the current volume labels. To delete a label, set the value of the label to an empty string. |
| spec | [ VolumeSpecSet](#volumespecset) | VolumeSpecSet provides a method to request that certain values in the VolumeSpec are changed. This is necessary as a separate variable because values like int and bool in the VolumeSpec cannot be determined if they are being requested to change in gRPC proto3. Some of these values may not be able to be changed. Here are a few examples of actions that can be accomplished using the VolumeSpec. To resize the volume: Set a new value in spec.size. To change number of replicas: Adjust spec.ha_level. To change the I/O Profile: Adjust spec.io_profile. |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeSetResponse {#sdkvolumesetresponse}
The response returns no data

 <!-- end HasFields -->


## SdkVolumeSnapshotCreateRequest {#sdkvolumesnapshotcreaterequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_id | [ string](#string) | Id of volume to take the snapshot from |
| labels | [map SdkVolumeSnapshotCreateRequest.LabelsEntry](#sdkvolumesnapshotcreaterequestlabelsentry) | Labels to apply to snapshot |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeSnapshotCreateRequest.LabelsEntry {#sdkvolumesnapshotcreaterequestlabelsentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeSnapshotCreateResponse {#sdkvolumesnapshotcreateresponse}



| Field | Type | Description |
| ----- | ---- | ----------- |
| snapshot_id | [ string](#string) | Id of immutable snapshot |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeSnapshotEnumerateRequest {#sdkvolumesnapshotenumeraterequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_id | [ string](#string) | Id of volume |
| labels | [map SdkVolumeSnapshotEnumerateRequest.LabelsEntry](#sdkvolumesnapshotenumeraterequestlabelsentry) | Labels from snapshot |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeSnapshotEnumerateRequest.LabelsEntry {#sdkvolumesnapshotenumeraterequestlabelsentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeSnapshotEnumerateResponse {#sdkvolumesnapshotenumerateresponse}



| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_snapshot_ids | [repeated string](#string) | List of immutable snapshots |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeSnapshotRestoreRequest {#sdkvolumesnapshotrestorerequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_id | [ string](#string) | Id of volume |
| snapshot_id | [ string](#string) | Snapshot id to apply to `volume_id` |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeSnapshotRestoreResponse {#sdkvolumesnapshotrestoreresponse}


 <!-- end HasFields -->


## SdkVolumeUnmountRequest {#sdkvolumeunmountrequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_id | [ string](#string) | Id of volume |
| mount_path | [ string](#string) | MountPath for device |
| options | [map SdkVolumeUnmountRequest.OptionsEntry](#sdkvolumeunmountrequestoptionsentry) | Options to unmount device |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeUnmountRequest.OptionsEntry {#sdkvolumeunmountrequestoptionsentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeUnmountResponse {#sdkvolumeunmountresponse}


 <!-- end HasFields -->


## SnapCreateRequest {#snapcreaterequest}
SnapCreateRequest specifies a request to create a snapshot of given volume.
swagger:parameters snapVolume


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ string](#string) | volume id |
| locator | [ VolumeLocator](#volumelocator) | none |
| readonly | [ bool](#bool) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## SnapCreateResponse {#snapcreateresponse}
SnapCreateRequest specifies a response that get's returned when creating a snapshot.
swagger:response snapCreateResponse


| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_create_response | [ VolumeCreateResponse](#volumecreateresponse) | VolumeCreateResponse

in: body Required: true |
 <!-- end Fields -->
 <!-- end HasFields -->


## Source {#source}
Source is a structure that can be given to a volume
to seed the volume with data.
swagger:model


| Field | Type | Description |
| ----- | ---- | ----------- |
| parent | [ string](#string) | A volume id, if specified will create a clone of the parent. |
| seed | [ string](#string) | Seed will seed the volume from the specified URI Any additional config for the source comes from the labels in the spec |
 <!-- end Fields -->
 <!-- end HasFields -->


## Stats {#stats}
Stats is a structure that represents last collected stats for a volume
swagger:model


| Field | Type | Description |
| ----- | ---- | ----------- |
| reads | [ uint64](#uint64) | Reads completed successfully |
| read_ms | [ uint64](#uint64) | Time spent in reads in ms |
| read_bytes | [ uint64](#uint64) | none |
| writes | [ uint64](#uint64) | Writes completed successfully |
| write_ms | [ uint64](#uint64) | Time spent in writes in ms |
| write_bytes | [ uint64](#uint64) | none |
| io_progress | [ uint64](#uint64) | IOs curently in progress |
| io_ms | [ uint64](#uint64) | Time spent doing IOs ms |
| bytes_used | [ uint64](#uint64) | BytesUsed |
| interval_ms | [ uint64](#uint64) | Interval in ms during which stats were collected |
 <!-- end Fields -->
 <!-- end HasFields -->


## StorageCluster {#storagecluster}
StorageCluster represents the state and information about the cluster


| Field | Type | Description |
| ----- | ---- | ----------- |
| status | [ Status](#status) | Status of the cluster |
| id | [ string](#string) | Id of the cluster |
| node_id | [ string](#string) | The id of the node servicing these requests |
| node_ids | [repeated string](#string) | List of all the node ids in the cluster |
 <!-- end Fields -->
 <!-- end HasFields -->


## StorageNode {#storagenode}
StorageNode describes the state of the node


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ string](#string) | Id of the node |
| cpu | [ double](#double) | Cpu usage of the node |
| mem_total | [ uint64](#uint64) | Total memory of the node |
| mem_used | [ uint64](#uint64) | Used memory of the node |
| mem_free | [ uint64](#uint64) | Free memory of the node |
| avg_load | [ int64](#int64) | Average load (percentage) |
| status | [ Status](#status) | Node status |
| disks | [map StorageNode.DisksEntry](#storagenodedisksentry) | List of disks on the node |
| pools | [repeated StoragePool](#storagepool) | List of storage pools this node supports |
| mgmt_ip | [ string](#string) | Management IP |
| data_ip | [ string](#string) | Data IP |
| hostname | [ string](#string) | Hostname of the node |
| node_labels | [map StorageNode.NodeLabelsEntry](#storagenodenodelabelsentry) | User defined labels for the node |
 <!-- end Fields -->
 <!-- end HasFields -->


## StorageNode.DisksEntry {#storagenodedisksentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ StorageResource](#storageresource) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## StorageNode.NodeLabelsEntry {#storagenodenodelabelsentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## StoragePool {#storagepool}
StoragePool groups different storage devices based on their CosType


| Field | Type | Description |
| ----- | ---- | ----------- |
| ID | [ int32](#int32) | ID pool ID |
| Cos | [ CosType](#costype) | Cos reflects the capabilities of this drive pool |
| Medium | [ StorageMedium](#storagemedium) | Medium underlying storage type |
| RaidLevel | [ string](#string) | RaidLevel storage raid level |
| TotalSize | [ uint64](#uint64) | TotalSize of the pool |
| Used | [ uint64](#uint64) | Used size of the pool |
| labels | [map StoragePool.LabelsEntry](#storagepoollabelsentry) | Labels is a list of user defined name-value pairs |
 <!-- end Fields -->
 <!-- end HasFields -->


## StoragePool.LabelsEntry {#storagepoollabelsentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## StorageResource {#storageresource}
StorageResource groups properties of a storage device.
swagger:model


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ string](#string) | Id is the LUN identifier. |
| path | [ string](#string) | Path device path for this storage resource. |
| medium | [ StorageMedium](#storagemedium) | Storage medium. |
| online | [ bool](#bool) | True if this device is online. |
| iops | [ uint64](#uint64) | IOPS |
| seq_write | [ double](#double) | SeqWrite |
| seq_read | [ double](#double) | SeqRead |
| randRW | [ double](#double) | RandRW |
| size | [ uint64](#uint64) | Total size in bytes. |
| used | [ uint64](#uint64) | Physical Bytes used. |
| rotation_speed | [ string](#string) | True if this device is rotational. |
| last_scan | [ google.protobuf.Timestamp](#googleprotobuftimestamp) | Timestamp of last time this device was scanned. |
| metadata | [ bool](#bool) | True if dedicated for metadata. |
 <!-- end Fields -->
 <!-- end HasFields -->


## Volume {#volume}
Volume represents an abstract storage volume.
Volume represents an abstract storage volume.
swagger:model


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ string](#string) | Self referential volume ID. |
| source | [ Source](#source) | Source specified seed data for the volume. |
| group | [ Group](#group) | Group volumes in the same group have the same group id. |
| readonly | [ bool](#bool) | Readonly is true if this volume is to be mounted with readonly access. |
| locator | [ VolumeLocator](#volumelocator) | User specified locator |
| ctime | [ google.protobuf.Timestamp](#googleprotobuftimestamp) | Volume creation time |
| spec | [ VolumeSpec](#volumespec) | User specified VolumeSpec |
| usage | [ uint64](#uint64) | Usage is bytes consumed by vtheis volume. |
| last_scan | [ google.protobuf.Timestamp](#googleprotobuftimestamp) | LastScan is the time when an integrity check was run. |
| format | [ FSType](#fstype) | Format specifies the filesytem for this volume. |
| status | [ VolumeStatus](#volumestatus) | Status is the availability status of this volume. |
| state | [ VolumeState](#volumestate) | State is the current runtime state of this volume. |
| attached_on | [ string](#string) | AttachedOn is the node instance identifier for clustered systems. |
| attached_state | [ AttachState](#attachstate) | AttachedState shows whether the device is attached for internal or external use. |
| device_path | [ string](#string) | DevicePath is the device exported by block device implementations. |
| secure_device_path | [ string](#string) | SecureDevicePath is the device path for an encrypted volume. |
| attach_path | [repeated string](#string) | AttachPath is the mounted path in the host namespace. |
| attach_info | [map Volume.AttachInfoEntry](#volumeattachinfoentry) | AttachInfo is a list of name value mappings that provides attach information. |
| replica_sets | [repeated ReplicaSet](#replicaset) | ReplicatSets storage for this volumefor clustered storage arrays. |
| runtime_state | [repeated RuntimeStateMap](#runtimestatemap) | RuntimeState is a lst of name value mapping of driver specific runtime information. |
| error | [ string](#string) | Error is the Last recorded error. |
| volume_consumers | [repeated VolumeConsumer](#volumeconsumer) | VolumeConsumers are entities that consume this volume |
 <!-- end Fields -->
 <!-- end HasFields -->


## Volume.AttachInfoEntry {#volumeattachinfoentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeConsumer {#volumeconsumer}
VolumeConsumer identifies a consumer for a Volume. An example of a VolumeConsumer
would be a Pod in Kubernetes who has mounted the PersistentVolumeClaim for the
Volume
swagger: model


| Field | Type | Description |
| ----- | ---- | ----------- |
| name | [ string](#string) | Name is the name of the volume consumer |
| namespace | [ string](#string) | Namespace is the namespace of the volume consumer |
| type | [ string](#string) | Type is the type of the consumer. E.g a Kubernetes pod |
| node_id | [ string](#string) | NodeID is the identifier of the node on which the consumer is running. This identifier would be from the perspective of the container runtime or orchestrator under which the volume consumer resides. For example, NodeID can be name of a minion in Kubernetes. |
| owner_name | [ string](#string) | OwnerName is the name of the entity who owns this volume consumer |
| owner_type | [ string](#string) | OwnerType is the type of the entity who owns this volume consumer. The type would be from the perspective of the container runtime or the orchestrator under which the volume consumer resides. For e.g OwnerType can be a Deployment in Kubernetes. |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeCreateRequest {#volumecreaterequest}
VolumeCreateRequest is a structure that has the locator, source and spec
to create a volume
swagger:model


| Field | Type | Description |
| ----- | ---- | ----------- |
| locator | [ VolumeLocator](#volumelocator) | User specified volume name and labels |
| source | [ Source](#source) | Source to create volume |
| spec | [ VolumeSpec](#volumespec) | The storage spec for the volume |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeCreateResponse {#volumecreateresponse}
VolumeCreateResponse
swagger:response volumeCreateResponse


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ string](#string) | ID of the newly created volume

in: body Required: true |
| volume_response | [ VolumeResponse](#volumeresponse) | Volume Response

in: body Required: true |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeInfo {#volumeinfo}
VolumeInfo
swagger:model


| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_id | [ string](#string) | none |
| path | [ string](#string) | none |
| storage | [ VolumeSpec](#volumespec) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeLocator {#volumelocator}
VolumeLocator is a structure that is attached to a volume
and is used to carry opaque metadata.
swagger:model


| Field | Type | Description |
| ----- | ---- | ----------- |
| name | [ string](#string) | User friendly identifier |
| volume_labels | [map VolumeLocator.VolumeLabelsEntry](#volumelocatorvolumelabelsentry) | A set of name-value pairs that acts as search filters |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeLocator.VolumeLabelsEntry {#volumelocatorvolumelabelsentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeResponse {#volumeresponse}
VolumeResponse is a structure that wraps an error.
swagger:response volumeResponse


| Field | Type | Description |
| ----- | ---- | ----------- |
| error | [ string](#string) | Error message

in: body Required: true |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeSetRequest {#volumesetrequest}
VolumeSet specifies a request to update a volume.
swagger:model


| Field | Type | Description |
| ----- | ---- | ----------- |
| locator | [ VolumeLocator](#volumelocator) | User specified volume name and labels |
| spec | [ VolumeSpec](#volumespec) | The storage spec for the volume |
| action | [ VolumeStateAction](#volumestateaction) | State modification on this volume. |
| options | [map VolumeSetRequest.OptionsEntry](#volumesetrequestoptionsentry) | additional options required for the Set operation. |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeSetRequest.OptionsEntry {#volumesetrequestoptionsentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeSetResponse {#volumesetresponse}
VolumeSetResponse
swagger:response volumeSetResponse


| Field | Type | Description |
| ----- | ---- | ----------- |
| volume | [ Volume](#volume) | Volume

in: body Required: true |
| volume_response | [ VolumeResponse](#volumeresponse) | VolumeResponse

in: body Required: true |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeSpec {#volumespec}
VolumeSpec has the properties needed to create a volume.
swagger:model


| Field | Type | Description |
| ----- | ---- | ----------- |
| ephemeral | [ bool](#bool) | Ephemeral storage |
| size | [ uint64](#uint64) | Size specifies the thin provisioned volume size in bytes |
| format | [ FSType](#fstype) | Format specifies the filesystem for this volume. |
| block_size | [ int64](#int64) | BlockSize for the filesystem. |
| ha_level | [ int64](#int64) | HaLevel specifies the number of copies of data. |
| cos | [ CosType](#costype) | Cos specifies the relative class of service. |
| io_profile | [ IoProfile](#ioprofile) | IoProfile provides a hint about application using this volume. |
| dedupe | [ bool](#bool) | Dedupe specifies if the volume data is to be de-duplicated. |
| snapshot_interval | [ uint32](#uint32) | SnapshotInterval in minutes, set to 0 to disable snapshots |
| volume_labels | [map VolumeSpec.VolumeLabelsEntry](#volumespecvolumelabelsentry) | VolumeLabels configuration labels |
| shared | [ bool](#bool) | Shared is true if this volume can be remotely accessed. |
| replica_set | [ ReplicaSet](#replicaset) | ReplicaSet is the desired set of nodes for the volume data. |
| aggregation_level | [ uint32](#uint32) | Aggregatiokn level Specifies the number of parts the volume can be aggregated from. |
| encrypted | [ bool](#bool) | Encrypted is true if this volume will be cryptographically secured. |
| passphrase | [ string](#string) | Passphrase for an encrypted volume |
| snapshot_schedule | [ string](#string) | SnapshotSchedule a well known string that specifies when snapshots should be taken. |
| scale | [ uint32](#uint32) | Scale allows autocreation of volumes. |
| sticky | [ bool](#bool) | Sticky volumes cannot be deleted until the flag is removed. |
| group | [ Group](#group) | Group identifies a consistency group |
| group_enforced | [ bool](#bool) | GroupEnforced is true if consistency group creation is enforced. |
| compressed | [ bool](#bool) | Compressed is true if this volume is to be compressed. |
| cascaded | [ bool](#bool) | Cascaded is true if this volume can be populated on any node from an external source. |
| journal | [ bool](#bool) | Journal is true if data for the volume goes into the journal. |
| sharedv4 | [ bool](#bool) | Sharedv4 is true if this volume can be accessed via sharedv4. |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeSpec.VolumeLabelsEntry {#volumespecvolumelabelsentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeSpecSet {#volumespecset}
VolumeSpecSet provides a method to set any of the VolumeSpec of an existing volume


| Field | Type | Description |
| ----- | ---- | ----------- |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) ephemeral_opt.ephemeral | [ bool](#bool) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) size_opt.size | [ uint64](#uint64) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) format_opt.format | [ FSType](#fstype) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) block_size_opt.block_size | [ int64](#int64) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) ha_level_opt.ha_level | [ int64](#int64) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) cos_opt.cos | [ CosType](#costype) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) io_profile_opt.io_profile | [ IoProfile](#ioprofile) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) dedupe_opt.dedupe | [ bool](#bool) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) snapshot_interval_opt.snapshot_interval | [ uint32](#uint32) | none |
| volume_labels | [map VolumeSpecSet.VolumeLabelsEntry](#volumespecsetvolumelabelsentry) | VolumeLabels configuration labels |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) shared_opt.shared | [ bool](#bool) | none |
| replica_set | [ ReplicaSet](#replicaset) | ReplicaSet is the desired set of nodes for the volume data. |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) aggregation_level_opt.aggregation_level | [ uint32](#uint32) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) encrypted_opt.encrypted | [ bool](#bool) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) passphrase_opt.passphrase | [ string](#string) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) snapshot_schedule_opt.snapshot_schedule | [ string](#string) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) scale_opt.scale | [ uint32](#uint32) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) sticky_opt.sticky | [ bool](#bool) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) group_opt.group | [ Group](#group) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) group_enforced_opt.group_enforced | [ bool](#bool) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) compressed_opt.compressed | [ bool](#bool) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) cascaded_opt.cascaded | [ bool](#bool) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) journal_opt.journal | [ bool](#bool) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) sharedv4_opt.sharedv4 | [ bool](#bool) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeSpecSet.VolumeLabelsEntry {#volumespecsetvolumelabelsentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeStateAction {#volumestateaction}
VolumeStateAction specifies desired actions.
swagger:model


| Field | Type | Description |
| ----- | ---- | ----------- |
| attach | [ VolumeActionParam](#volumeactionparam) | Attach or Detach volume |
| mount | [ VolumeActionParam](#volumeactionparam) | Mount or unmount volume |
| mount_path | [ string](#string) | MountPath Path where the device is mounted |
| device_path | [ string](#string) | DevicePath Path returned in attach |
 <!-- end Fields -->
 <!-- end HasFields -->
 <!-- end messages -->

# Enums


## AlertActionType {#alertactiontype}


| Name | Number | Description |
| ---- | ------ | ----------- |
| ALERT_ACTION_TYPE_NONE | 0 | none |
| ALERT_ACTION_TYPE_DELETE | 1 | none |
| ALERT_ACTION_TYPE_CREATE | 2 | none |
| ALERT_ACTION_TYPE_UPDATE | 3 | none |




## AttachState {#attachstate}


| Name | Number | Description |
| ---- | ------ | ----------- |
| ATTACH_STATE_EXTERNAL | 0 | Attached and available externally |
| ATTACH_STATE_INTERNAL | 1 | Attached but only available internally |
| ATTACH_STATE_INTERNAL_SWITCH | 2 | Switching from External to Internal |




## CloudMigrate.OperationType {#cloudmigrateoperationtype}


| Name | Number | Description |
| ---- | ------ | ----------- |
| InvalidType | 0 | none |
| MigrateCluster | 1 | Migrate all volumes in the cluster |
| MigrateVolume | 2 | Migrate a single volume |
| MigrateVolumeGroup | 3 | Migrate a group of volumes |




## CloudMigrate.Stage {#cloudmigratestage}


| Name | Number | Description |
| ---- | ------ | ----------- |
| InvalidStage | 0 | none |
| Backup | 1 | none |
| Restore | 2 | none |
| VolumeUpdate | 3 | none |
| Done | 4 | none |




## CloudMigrate.Status {#cloudmigratestatus}


| Name | Number | Description |
| ---- | ------ | ----------- |
| InvalidStatus | 0 | none |
| Queued | 1 | none |
| Initialized | 2 | none |
| InProgress | 3 | none |
| Failed | 4 | none |
| Complete | 5 | none |




## ClusterNotify {#clusternotify}


| Name | Number | Description |
| ---- | ------ | ----------- |
| CLUSTER_NOTIFY_DOWN | 0 | Node is down |




## CosType {#costype}


| Name | Number | Description |
| ---- | ------ | ----------- |
| NONE | 0 | none |
| LOW | 1 | none |
| MEDIUM | 2 | none |
| HIGH | 3 | none |




## DriverType {#drivertype}


| Name | Number | Description |
| ---- | ------ | ----------- |
| DRIVER_TYPE_NONE | 0 | none |
| DRIVER_TYPE_FILE | 1 | none |
| DRIVER_TYPE_BLOCK | 2 | none |
| DRIVER_TYPE_OBJECT | 3 | none |
| DRIVER_TYPE_CLUSTERED | 4 | none |
| DRIVER_TYPE_GRAPH | 5 | none |




## FSType {#fstype}


| Name | Number | Description |
| ---- | ------ | ----------- |
| FS_TYPE_NONE | 0 | none |
| FS_TYPE_BTRFS | 1 | none |
| FS_TYPE_EXT4 | 2 | none |
| FS_TYPE_FUSE | 3 | none |
| FS_TYPE_NFS | 4 | none |
| FS_TYPE_VFS | 5 | none |
| FS_TYPE_XFS | 6 | none |
| FS_TYPE_ZFS | 7 | none |




## GraphDriverChangeType {#graphdriverchangetype}


| Name | Number | Description |
| ---- | ------ | ----------- |
| GRAPH_DRIVER_CHANGE_TYPE_NONE | 0 | none |
| GRAPH_DRIVER_CHANGE_TYPE_MODIFIED | 1 | none |
| GRAPH_DRIVER_CHANGE_TYPE_ADDED | 2 | none |
| GRAPH_DRIVER_CHANGE_TYPE_DELETED | 3 | none |




## IoProfile {#ioprofile}


| Name | Number | Description |
| ---- | ------ | ----------- |
| IO_PROFILE_SEQUENTIAL | 0 | none |
| IO_PROFILE_RANDOM | 1 | none |
| IO_PROFILE_DB | 2 | none |
| IO_PROFILE_DB_REMOTE | 3 | none |
| IO_PROFILE_CMS | 4 | none |




## OperationFlags {#operationflags}


| Name | Number | Description |
| ---- | ------ | ----------- |
| OP_FLAGS_UNKNOWN | 0 | none |
| OP_FLAGS_NONE | 1 | none |
| OP_FLAGS_DETACH_FORCE | 2 | Perform a force_detach during detach operation |




## ResourceType {#resourcetype}


| Name | Number | Description |
| ---- | ------ | ----------- |
| RESOURCE_TYPE_NONE | 0 | none |
| RESOURCE_TYPE_VOLUME | 1 | none |
| RESOURCE_TYPE_NODE | 2 | none |
| RESOURCE_TYPE_CLUSTER | 3 | none |
| RESOURCE_TYPE_DRIVE | 4 | none |




## SdkCloudBackupOpType {#sdkcloudbackupoptype}


| Name | Number | Description |
| ---- | ------ | ----------- |
| SdkCloudBackupOpTypeUnknown | 0 | none |
| SdkCloudBackupOpTypeBackupOp | 1 | none |
| SdkCloudBackupOpTypeRestoreOp | 2 | none |




## SdkCloudBackupRequestedState {#sdkcloudbackuprequestedstate}


| Name | Number | Description |
| ---- | ------ | ----------- |
| SdkCloudBackupRequestedStateUnknown | 0 | none |
| SdkCloudBackupRequestedStatePause | 1 | none |
| SdkCloudBackupRequestedStateResume | 2 | none |
| SdkCloudBackupRequestedStateStop | 3 | none |




## SdkCloudBackupStatusType {#sdkcloudbackupstatustype}


| Name | Number | Description |
| ---- | ------ | ----------- |
| SdkCloudBackupStatusTypeUnknown | 0 | none |
| SdkCloudBackupStatusTypeNotStarted | 1 | none |
| SdkCloudBackupStatusTypeDone | 2 | none |
| SdkCloudBackupStatusTypeAborted | 3 | none |
| SdkCloudBackupStatusTypePaused | 4 | none |
| SdkCloudBackupStatusTypeStopped | 5 | none |
| SdkCloudBackupStatusTypeActive | 6 | none |
| SdkCloudBackupStatusTypeFailed | 7 | none |




## SdkTimeWeekday {#sdktimeweekday}


| Name | Number | Description |
| ---- | ------ | ----------- |
| SdkTimeWeekdaySunday | 0 | none |
| SdkTimeWeekdayMonday | 1 | none |
| SdkTimeWeekdayTuesday | 2 | none |
| SdkTimeWeekdayWednesday | 3 | none |
| SdkTimeWeekdayThursday | 4 | none |
| SdkTimeWeekdayFriday | 5 | none |
| SdkTimeWeekdaySaturday | 6 | none |




## SeverityType {#severitytype}


| Name | Number | Description |
| ---- | ------ | ----------- |
| SEVERITY_TYPE_NONE | 0 | none |
| SEVERITY_TYPE_ALARM | 1 | none |
| SEVERITY_TYPE_WARNING | 2 | none |
| SEVERITY_TYPE_NOTIFY | 3 | none |




## Status {#status}


| Name | Number | Description |
| ---- | ------ | ----------- |
| STATUS_NONE | 0 | none |
| STATUS_INIT | 1 | none |
| STATUS_OK | 2 | none |
| STATUS_OFFLINE | 3 | none |
| STATUS_ERROR | 4 | none |
| STATUS_NOT_IN_QUORUM | 5 | none |
| STATUS_DECOMMISSION | 6 | none |
| STATUS_MAINTENANCE | 7 | none |
| STATUS_STORAGE_DOWN | 8 | none |
| STATUS_STORAGE_DEGRADED | 9 | none |
| STATUS_NEEDS_REBOOT | 10 | none |
| STATUS_STORAGE_REBALANCE | 11 | none |
| STATUS_STORAGE_DRIVE_REPLACE | 12 | none |
| STATUS_NOT_IN_QUORUM_NO_STORAGE | 13 | none |
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
| VOLUME_ACTION_PARAM_NONE | 0 | none |
| VOLUME_ACTION_PARAM_OFF | 1 | Maps to the boolean value false |
| VOLUME_ACTION_PARAM_ON | 2 | Maps to the boolean value true. |




## VolumeState {#volumestate}
VolumeState represents the state of a volume.

| Name | Number | Description |
| ---- | ------ | ----------- |
| VOLUME_STATE_NONE | 0 | none |
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
| VOLUME_STATUS_NONE | 0 | none |
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

