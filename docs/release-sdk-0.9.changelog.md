# Changelog

> NOTE: The SDK is still in tech preview. Once officially released, this changelog will also
> use the SDK version numbers.

## Releases

### v0.9.3 - Tech Preview (12/5/2018)

* Add new API for extracting volume capacity usage details.

### v0.9.2 - Tech Preview (10/15/2018)

* Changed value of SdkSchedulePolicyCreateRequest from `SchedulePolicy` to the
  correct name of `schedule_policy`. This will not impact Golang.

### v0.9.1 - Tech Preview (10/10/2018)

* (Backport from v0.15.0) Added support to set the snapshot schedule policy of a Volume
* (Backport from v0.14.0) Added support for periodic type in OpenStorageSchedulePolicy service

### v0.9.0 - Tech Preview (9/18/2018)

NOTE: This release has breaking chages for the Mount/Attach/Detach/Unmount calls

* New service called OpenStorageMountAttach has been created and documented to
  hold the mount/attach/detach/unmount calls.
* Mount/Attach/Detach/Unmount calls have been moved from the OpenStorageVolume
  service to the OpenStorageMountAttach service.

### v0.8.0 - Tech Preview (9/11/2018)

* SdkVolumeSnapshotEnumerateWithFilters all attributes are now optional. [#609](https://github.com/libopenstorage/openstorage/issues/609)

### v0.7.0 - Tech Preview (9/5/2018)

* Add `Name` to `StorageCluster`. This name will hold the name given to the cluster by the administrator. The `StorageCluster.Id` will now hold a unique id for the cluster.

### v0.6.0 - Tech Preview (8/30/2018)

* Remove unsupported FS Types from supported drivers [#593](https://github.com/libopenstorage/openstorage/issues/593)
* Remove SDK Alert calls as they will be redesinged [#596](https://github.com/libopenstorage/openstorage/issues/596)

### v0.5.0 - Tech Preview (8/25/2018)

* Added `queue_depth` to VolumeSpec and VolumeSpecUpdate
* Remove values from VolumeSpecUpdate which cannot be updated [#590](https://github.com/libopenstorage/openstorage/issues/590)

### v0.4.0 - Tech Preview (8/24/2018)

* Added bucket name and encryption key to SdkCredentialCreateRequest
* Added the ability to disable ssl connection to SdkAwsCredentialRequest

### v0.3.0 - Tech Preview (8/20/2018)

* Added SchedulerNodeName field to StorageNode object

### v0.2.0 - Tech Preview (8/16/2018)

* Changed Credentials.Create to take `name` as a required parameter

### Tech Preview (8/7/2018)

* Added [Identity.Version](https://libopenstorage.github.io/w/generated-api.html#methodopenstorageapiopenstorageidentityversion) Service

### Tech Preview (8/3/2018)

* Added [Idenity](https://libopenstorage.github.io/w/generated-api.html#serviceopenstorageapiopenstorageidentity) Service
* Added [Identity.Capabilities](https://libopenstorage.github.io/w/generated-api.html#methodopenstorageapiopenstorageidentitycapabilities) RPC
