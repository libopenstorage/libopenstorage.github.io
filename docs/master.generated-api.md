# API Reference

# Table of Contents


- Services
    - [OpenStorageAlerts](#serviceopenstorageapiopenstoragealerts)
    - [OpenStorageBucket](#serviceopenstorageapiopenstoragebucket)
    - [OpenStorageCloudBackup](#serviceopenstorageapiopenstoragecloudbackup)
    - [OpenStorageCluster](#serviceopenstorageapiopenstoragecluster)
    - [OpenStorageClusterDomains](#serviceopenstorageapiopenstorageclusterdomains)
    - [OpenStorageClusterPair](#serviceopenstorageapiopenstorageclusterpair)
    - [OpenStorageCredentials](#serviceopenstorageapiopenstoragecredentials)
    - [OpenStorageDiags](#serviceopenstorageapiopenstoragediags)
    - [OpenStorageFilesystemCheck](#serviceopenstorageapiopenstoragefilesystemcheck)
    - [OpenStorageFilesystemTrim](#serviceopenstorageapiopenstoragefilesystemtrim)
    - [OpenStorageIdentity](#serviceopenstorageapiopenstorageidentity)
    - [OpenStorageJob](#serviceopenstorageapiopenstoragejob)
    - [OpenStorageMigrate](#serviceopenstorageapiopenstoragemigrate)
    - [OpenStorageMountAttach](#serviceopenstorageapiopenstoragemountattach)
    - [OpenStorageNode](#serviceopenstorageapiopenstoragenode)
    - [OpenStorageObjectstore](#serviceopenstorageapiopenstorageobjectstore)
    - [OpenStoragePolicy](#serviceopenstorageapiopenstoragepolicy)
    - [OpenStoragePool](#serviceopenstorageapiopenstoragepool)
    - [OpenStorageRole](#serviceopenstorageapiopenstoragerole)
    - [OpenStorageSchedulePolicy](#serviceopenstorageapiopenstorageschedulepolicy)
    - [OpenStorageVolume](#serviceopenstorageapiopenstoragevolume)
  


- Messages
    - [ActiveRequest](#activerequest)
    - [ActiveRequest.ReqestKVEntry](#activerequestreqestkventry)
    - [ActiveRequests](#activerequests)
    - [Alert](#alert)
    - [Alerts](#alerts)
    - [BucketAccessCredentials](#bucketaccesscredentials)
    - [BucketCreateRequest](#bucketcreaterequest)
    - [BucketCreateResponse](#bucketcreateresponse)
    - [BucketDeleteRequest](#bucketdeleterequest)
    - [BucketDeleteResponse](#bucketdeleteresponse)
    - [BucketGrantAccessRequest](#bucketgrantaccessrequest)
    - [BucketGrantAccessResponse](#bucketgrantaccessresponse)
    - [BucketRevokeAccessRequest](#bucketrevokeaccessrequest)
    - [BucketRevokeAccessResponse](#bucketrevokeaccessresponse)
    - [CapacityUsageInfo](#capacityusageinfo)
    - [Catalog](#catalog)
    - [CatalogResponse](#catalogresponse)
    - [CloudDriveTransferJob](#clouddrivetransferjob)
    - [CloudMigrate](#cloudmigrate)
    - [CloudMigrateCancelRequest](#cloudmigratecancelrequest)
    - [CloudMigrateInfo](#cloudmigrateinfo)
    - [CloudMigrateInfoList](#cloudmigrateinfolist)
    - [CloudMigrateStartRequest](#cloudmigratestartrequest)
    - [CloudMigrateStartResponse](#cloudmigratestartresponse)
    - [CloudMigrateStatusRequest](#cloudmigratestatusrequest)
    - [CloudMigrateStatusResponse](#cloudmigratestatusresponse)
    - [CloudMigrateStatusResponse.InfoEntry](#cloudmigratestatusresponseinfoentry)
    - [ClusterPairCreateRequest](#clusterpaircreaterequest)
    - [ClusterPairCreateResponse](#clusterpaircreateresponse)
    - [ClusterPairGetResponse](#clusterpairgetresponse)
    - [ClusterPairInfo](#clusterpairinfo)
    - [ClusterPairInfo.OptionsEntry](#clusterpairinfooptionsentry)
    - [ClusterPairMode](#clusterpairmode)
    - [ClusterPairProcessRequest](#clusterpairprocessrequest)
    - [ClusterPairProcessResponse](#clusterpairprocessresponse)
    - [ClusterPairProcessResponse.OptionsEntry](#clusterpairprocessresponseoptionsentry)
    - [ClusterPairTokenGetResponse](#clusterpairtokengetresponse)
    - [ClusterPairsEnumerateResponse](#clusterpairsenumerateresponse)
    - [ClusterPairsEnumerateResponse.PairsEntry](#clusterpairsenumerateresponsepairsentry)
    - [ClusterResponse](#clusterresponse)
    - [CollectDiagsJob](#collectdiagsjob)
    - [DiagsCollectionStatus](#diagscollectionstatus)
    - [DiagsNodeSelector](#diagsnodeselector)
    - [DiagsVolumeSelector](#diagsvolumeselector)
    - [DrainAttachmentsSummary](#drainattachmentssummary)
    - [ExportSpec](#exportspec)
    - [FastpathConfig](#fastpathconfig)
    - [FastpathReplState](#fastpathreplstate)
    - [FilesystemCheck](#filesystemcheck)
    - [FilesystemTrim](#filesystemtrim)
    - [FstrimVolumeUsageInfo](#fstrimvolumeusageinfo)
    - [GraphDriverChanges](#graphdriverchanges)
    - [Group](#group)
    - [GroupSnapCreateRequest](#groupsnapcreaterequest)
    - [GroupSnapCreateRequest.LabelsEntry](#groupsnapcreaterequestlabelsentry)
    - [GroupSnapCreateResponse](#groupsnapcreateresponse)
    - [GroupSnapCreateResponse.SnapshotsEntry](#groupsnapcreateresponsesnapshotsentry)
    - [IoStrategy](#iostrategy)
    - [IoThrottle](#iothrottle)
    - [Job](#job)
    - [JobAudit](#jobaudit)
    - [JobSummary](#jobsummary)
    - [JobWorkSummary](#jobworksummary)
    - [LabelSelectorRequirement](#labelselectorrequirement)
    - [LocateResponse](#locateresponse)
    - [LocateResponse.DockeridsEntry](#locateresponsedockeridsentry)
    - [LocateResponse.MountsEntry](#locateresponsemountsentry)
    - [MountOptions](#mountoptions)
    - [MountOptions.OptionsEntry](#mountoptionsoptionsentry)
    - [NFSProxySpec](#nfsproxyspec)
    - [NodeDrainAttachmentOptions](#nodedrainattachmentoptions)
    - [NodeDrainAttachmentsJob](#nodedrainattachmentsjob)
    - [ObjectstoreInfo](#objectstoreinfo)
    - [Ownership](#ownership)
    - [Ownership.AccessControl](#ownershipaccesscontrol)
    - [Ownership.AccessControl.CollaboratorsEntry](#ownershipaccesscontrolcollaboratorsentry)
    - [Ownership.AccessControl.GroupsEntry](#ownershipaccesscontrolgroupsentry)
    - [Ownership.PublicAccessControl](#ownershippublicaccesscontrol)
    - [PXDProxySpec](#pxdproxyspec)
    - [ProxySpec](#proxyspec)
    - [PureBlockSpec](#pureblockspec)
    - [PureFileSpec](#purefilespec)
    - [RelaxedReclaimPurge](#relaxedreclaimpurge)
    - [ReplicaPlacementSpec](#replicaplacementspec)
    - [ReplicaSet](#replicaset)
    - [Report](#report)
    - [RestoreVolSnashotSchedule](#restorevolsnashotschedule)
    - [RestoreVolStoragePolicy](#restorevolstoragepolicy)
    - [RestoreVolumeSpec](#restorevolumespec)
    - [RuntimeStateMap](#runtimestatemap)
    - [RuntimeStateMap.RuntimeStateEntry](#runtimestatemapruntimestateentry)
    - [S3ProxySpec](#s3proxyspec)
    - [ScanPolicy](#scanpolicy)
    - [SchedulerTopology](#schedulertopology)
    - [SchedulerTopology.LabelsEntry](#schedulertopologylabelsentry)
    - [SdkAlertsAlertTypeQuery](#sdkalertsalerttypequery)
    - [SdkAlertsCountSpan](#sdkalertscountspan)
    - [SdkAlertsDeleteRequest](#sdkalertsdeleterequest)
    - [SdkAlertsDeleteResponse](#sdkalertsdeleteresponse)
    - [SdkAlertsEnumerateWithFiltersRequest](#sdkalertsenumeratewithfiltersrequest)
    - [SdkAlertsEnumerateWithFiltersResponse](#sdkalertsenumeratewithfiltersresponse)
    - [SdkAlertsOption](#sdkalertsoption)
    - [SdkAlertsQuery](#sdkalertsquery)
    - [SdkAlertsResourceIdQuery](#sdkalertsresourceidquery)
    - [SdkAlertsResourceTypeQuery](#sdkalertsresourcetypequery)
    - [SdkAlertsTimeSpan](#sdkalertstimespan)
    - [SdkAutoFSTrimStatusRequest](#sdkautofstrimstatusrequest)
    - [SdkAutoFSTrimStatusResponse](#sdkautofstrimstatusresponse)
    - [SdkAutoFSTrimStatusResponse.TrimStatusEntry](#sdkautofstrimstatusresponsetrimstatusentry)
    - [SdkAutoFSTrimUsageRequest](#sdkautofstrimusagerequest)
    - [SdkAutoFSTrimUsageResponse](#sdkautofstrimusageresponse)
    - [SdkAutoFSTrimUsageResponse.UsageEntry](#sdkautofstrimusageresponseusageentry)
    - [SdkAwsCredentialRequest](#sdkawscredentialrequest)
    - [SdkAwsCredentialResponse](#sdkawscredentialresponse)
    - [SdkAzureCredentialRequest](#sdkazurecredentialrequest)
    - [SdkAzureCredentialResponse](#sdkazurecredentialresponse)
    - [SdkCloudBackupCatalogRequest](#sdkcloudbackupcatalogrequest)
    - [SdkCloudBackupCatalogResponse](#sdkcloudbackupcatalogresponse)
    - [SdkCloudBackupCreateRequest](#sdkcloudbackupcreaterequest)
    - [SdkCloudBackupCreateRequest.LabelsEntry](#sdkcloudbackupcreaterequestlabelsentry)
    - [SdkCloudBackupCreateResponse](#sdkcloudbackupcreateresponse)
    - [SdkCloudBackupDeleteAllRequest](#sdkcloudbackupdeleteallrequest)
    - [SdkCloudBackupDeleteAllResponse](#sdkcloudbackupdeleteallresponse)
    - [SdkCloudBackupDeleteRequest](#sdkcloudbackupdeleterequest)
    - [SdkCloudBackupDeleteResponse](#sdkcloudbackupdeleteresponse)
    - [SdkCloudBackupEnumerateWithFiltersRequest](#sdkcloudbackupenumeratewithfiltersrequest)
    - [SdkCloudBackupEnumerateWithFiltersRequest.MetadataFilterEntry](#sdkcloudbackupenumeratewithfiltersrequestmetadatafilterentry)
    - [SdkCloudBackupEnumerateWithFiltersResponse](#sdkcloudbackupenumeratewithfiltersresponse)
    - [SdkCloudBackupGroupCreateRequest](#sdkcloudbackupgroupcreaterequest)
    - [SdkCloudBackupGroupCreateRequest.LabelsEntry](#sdkcloudbackupgroupcreaterequestlabelsentry)
    - [SdkCloudBackupGroupCreateResponse](#sdkcloudbackupgroupcreateresponse)
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
    - [SdkCloudBackupSchedUpdateRequest](#sdkcloudbackupschedupdaterequest)
    - [SdkCloudBackupSchedUpdateResponse](#sdkcloudbackupschedupdateresponse)
    - [SdkCloudBackupScheduleInfo](#sdkcloudbackupscheduleinfo)
    - [SdkCloudBackupScheduleInfo.LabelsEntry](#sdkcloudbackupscheduleinfolabelsentry)
    - [SdkCloudBackupSizeRequest](#sdkcloudbackupsizerequest)
    - [SdkCloudBackupSizeResponse](#sdkcloudbackupsizeresponse)
    - [SdkCloudBackupStateChangeRequest](#sdkcloudbackupstatechangerequest)
    - [SdkCloudBackupStateChangeResponse](#sdkcloudbackupstatechangeresponse)
    - [SdkCloudBackupStatus](#sdkcloudbackupstatus)
    - [SdkCloudBackupStatusRequest](#sdkcloudbackupstatusrequest)
    - [SdkCloudBackupStatusResponse](#sdkcloudbackupstatusresponse)
    - [SdkCloudBackupStatusResponse.StatusesEntry](#sdkcloudbackupstatusresponsestatusesentry)
    - [SdkCloudMigrateCancelRequest](#sdkcloudmigratecancelrequest)
    - [SdkCloudMigrateCancelResponse](#sdkcloudmigratecancelresponse)
    - [SdkCloudMigrateStartRequest](#sdkcloudmigratestartrequest)
    - [SdkCloudMigrateStartRequest.MigrateAllVolumes](#sdkcloudmigratestartrequestmigrateallvolumes)
    - [SdkCloudMigrateStartRequest.MigrateVolume](#sdkcloudmigratestartrequestmigratevolume)
    - [SdkCloudMigrateStartRequest.MigrateVolumeGroup](#sdkcloudmigratestartrequestmigratevolumegroup)
    - [SdkCloudMigrateStartResponse](#sdkcloudmigratestartresponse)
    - [SdkCloudMigrateStatusRequest](#sdkcloudmigratestatusrequest)
    - [SdkCloudMigrateStatusResponse](#sdkcloudmigratestatusresponse)
    - [SdkClusterDomainActivateRequest](#sdkclusterdomainactivaterequest)
    - [SdkClusterDomainActivateResponse](#sdkclusterdomainactivateresponse)
    - [SdkClusterDomainDeactivateRequest](#sdkclusterdomaindeactivaterequest)
    - [SdkClusterDomainDeactivateResponse](#sdkclusterdomaindeactivateresponse)
    - [SdkClusterDomainInspectRequest](#sdkclusterdomaininspectrequest)
    - [SdkClusterDomainInspectResponse](#sdkclusterdomaininspectresponse)
    - [SdkClusterDomainsEnumerateRequest](#sdkclusterdomainsenumeraterequest)
    - [SdkClusterDomainsEnumerateResponse](#sdkclusterdomainsenumerateresponse)
    - [SdkClusterInspectCurrentRequest](#sdkclusterinspectcurrentrequest)
    - [SdkClusterInspectCurrentResponse](#sdkclusterinspectcurrentresponse)
    - [SdkClusterPairCreateRequest](#sdkclusterpaircreaterequest)
    - [SdkClusterPairCreateResponse](#sdkclusterpaircreateresponse)
    - [SdkClusterPairDeleteRequest](#sdkclusterpairdeleterequest)
    - [SdkClusterPairDeleteResponse](#sdkclusterpairdeleteresponse)
    - [SdkClusterPairEnumerateRequest](#sdkclusterpairenumeraterequest)
    - [SdkClusterPairEnumerateResponse](#sdkclusterpairenumerateresponse)
    - [SdkClusterPairGetTokenRequest](#sdkclusterpairgettokenrequest)
    - [SdkClusterPairGetTokenResponse](#sdkclusterpairgettokenresponse)
    - [SdkClusterPairInspectRequest](#sdkclusterpairinspectrequest)
    - [SdkClusterPairInspectResponse](#sdkclusterpairinspectresponse)
    - [SdkClusterPairResetTokenRequest](#sdkclusterpairresettokenrequest)
    - [SdkClusterPairResetTokenResponse](#sdkclusterpairresettokenresponse)
    - [SdkCredentialCreateRequest](#sdkcredentialcreaterequest)
    - [SdkCredentialCreateResponse](#sdkcredentialcreateresponse)
    - [SdkCredentialDeleteReferencesRequest](#sdkcredentialdeletereferencesrequest)
    - [SdkCredentialDeleteReferencesResponse](#sdkcredentialdeletereferencesresponse)
    - [SdkCredentialDeleteRequest](#sdkcredentialdeleterequest)
    - [SdkCredentialDeleteResponse](#sdkcredentialdeleteresponse)
    - [SdkCredentialEnumerateRequest](#sdkcredentialenumeraterequest)
    - [SdkCredentialEnumerateResponse](#sdkcredentialenumerateresponse)
    - [SdkCredentialInspectRequest](#sdkcredentialinspectrequest)
    - [SdkCredentialInspectResponse](#sdkcredentialinspectresponse)
    - [SdkCredentialUpdateRequest](#sdkcredentialupdaterequest)
    - [SdkCredentialUpdateResponse](#sdkcredentialupdateresponse)
    - [SdkCredentialValidateRequest](#sdkcredentialvalidaterequest)
    - [SdkCredentialValidateResponse](#sdkcredentialvalidateresponse)
    - [SdkDiagsCollectRequest](#sdkdiagscollectrequest)
    - [SdkDiagsCollectResponse](#sdkdiagscollectresponse)
    - [SdkEnumerateJobsRequest](#sdkenumeratejobsrequest)
    - [SdkEnumerateJobsResponse](#sdkenumeratejobsresponse)
    - [SdkEnumerateRebalanceJobsRequest](#sdkenumeraterebalancejobsrequest)
    - [SdkEnumerateRebalanceJobsResponse](#sdkenumeraterebalancejobsresponse)
    - [SdkFilesystemCheckStartRequest](#sdkfilesystemcheckstartrequest)
    - [SdkFilesystemCheckStartResponse](#sdkfilesystemcheckstartresponse)
    - [SdkFilesystemCheckStatusRequest](#sdkfilesystemcheckstatusrequest)
    - [SdkFilesystemCheckStatusResponse](#sdkfilesystemcheckstatusresponse)
    - [SdkFilesystemCheckStopRequest](#sdkfilesystemcheckstoprequest)
    - [SdkFilesystemCheckStopResponse](#sdkfilesystemcheckstopresponse)
    - [SdkFilesystemTrimStartRequest](#sdkfilesystemtrimstartrequest)
    - [SdkFilesystemTrimStartResponse](#sdkfilesystemtrimstartresponse)
    - [SdkFilesystemTrimStatusRequest](#sdkfilesystemtrimstatusrequest)
    - [SdkFilesystemTrimStatusResponse](#sdkfilesystemtrimstatusresponse)
    - [SdkFilesystemTrimStopRequest](#sdkfilesystemtrimstoprequest)
    - [SdkFilesystemTrimStopResponse](#sdkfilesystemtrimstopresponse)
    - [SdkGetJobStatusRequest](#sdkgetjobstatusrequest)
    - [SdkGetJobStatusResponse](#sdkgetjobstatusresponse)
    - [SdkGetRebalanceJobStatusRequest](#sdkgetrebalancejobstatusrequest)
    - [SdkGetRebalanceJobStatusResponse](#sdkgetrebalancejobstatusresponse)
    - [SdkGoogleCredentialRequest](#sdkgooglecredentialrequest)
    - [SdkGoogleCredentialResponse](#sdkgooglecredentialresponse)
    - [SdkIdentityCapabilitiesRequest](#sdkidentitycapabilitiesrequest)
    - [SdkIdentityCapabilitiesResponse](#sdkidentitycapabilitiesresponse)
    - [SdkIdentityVersionRequest](#sdkidentityversionrequest)
    - [SdkIdentityVersionResponse](#sdkidentityversionresponse)
    - [SdkJobResponse](#sdkjobresponse)
    - [SdkNfsCredentialRequest](#sdknfscredentialrequest)
    - [SdkNfsCredentialResponse](#sdknfscredentialresponse)
    - [SdkNodeCordonAttachmentsRequest](#sdknodecordonattachmentsrequest)
    - [SdkNodeCordonAttachmentsResponse](#sdknodecordonattachmentsresponse)
    - [SdkNodeDrainAttachmentsRequest](#sdknodedrainattachmentsrequest)
    - [SdkNodeEnumerateRequest](#sdknodeenumeraterequest)
    - [SdkNodeEnumerateResponse](#sdknodeenumerateresponse)
    - [SdkNodeEnumerateWithFiltersRequest](#sdknodeenumeratewithfiltersrequest)
    - [SdkNodeEnumerateWithFiltersResponse](#sdknodeenumeratewithfiltersresponse)
    - [SdkNodeInspectCurrentRequest](#sdknodeinspectcurrentrequest)
    - [SdkNodeInspectCurrentResponse](#sdknodeinspectcurrentresponse)
    - [SdkNodeInspectRequest](#sdknodeinspectrequest)
    - [SdkNodeInspectResponse](#sdknodeinspectresponse)
    - [SdkNodeRelaxedReclaimPurgeRequest](#sdknoderelaxedreclaimpurgerequest)
    - [SdkNodeRelaxedReclaimPurgeResponse](#sdknoderelaxedreclaimpurgeresponse)
    - [SdkNodeUncordonAttachmentsRequest](#sdknodeuncordonattachmentsrequest)
    - [SdkNodeUncordonAttachmentsResponse](#sdknodeuncordonattachmentsresponse)
    - [SdkNodeVolumeUsageByNodeRequest](#sdknodevolumeusagebynoderequest)
    - [SdkNodeVolumeUsageByNodeResponse](#sdknodevolumeusagebynoderesponse)
    - [SdkObjectstoreCreateRequest](#sdkobjectstorecreaterequest)
    - [SdkObjectstoreCreateResponse](#sdkobjectstorecreateresponse)
    - [SdkObjectstoreDeleteRequest](#sdkobjectstoredeleterequest)
    - [SdkObjectstoreDeleteResponse](#sdkobjectstoredeleteresponse)
    - [SdkObjectstoreInspectRequest](#sdkobjectstoreinspectrequest)
    - [SdkObjectstoreInspectResponse](#sdkobjectstoreinspectresponse)
    - [SdkObjectstoreUpdateRequest](#sdkobjectstoreupdaterequest)
    - [SdkObjectstoreUpdateResponse](#sdkobjectstoreupdateresponse)
    - [SdkOpenStoragePolicyCreateRequest](#sdkopenstoragepolicycreaterequest)
    - [SdkOpenStoragePolicyCreateResponse](#sdkopenstoragepolicycreateresponse)
    - [SdkOpenStoragePolicyDefaultInspectRequest](#sdkopenstoragepolicydefaultinspectrequest)
    - [SdkOpenStoragePolicyDefaultInspectResponse](#sdkopenstoragepolicydefaultinspectresponse)
    - [SdkOpenStoragePolicyDeleteRequest](#sdkopenstoragepolicydeleterequest)
    - [SdkOpenStoragePolicyDeleteResponse](#sdkopenstoragepolicydeleteresponse)
    - [SdkOpenStoragePolicyEnumerateRequest](#sdkopenstoragepolicyenumeraterequest)
    - [SdkOpenStoragePolicyEnumerateResponse](#sdkopenstoragepolicyenumerateresponse)
    - [SdkOpenStoragePolicyInspectRequest](#sdkopenstoragepolicyinspectrequest)
    - [SdkOpenStoragePolicyInspectResponse](#sdkopenstoragepolicyinspectresponse)
    - [SdkOpenStoragePolicyReleaseRequest](#sdkopenstoragepolicyreleaserequest)
    - [SdkOpenStoragePolicyReleaseResponse](#sdkopenstoragepolicyreleaseresponse)
    - [SdkOpenStoragePolicySetDefaultRequest](#sdkopenstoragepolicysetdefaultrequest)
    - [SdkOpenStoragePolicySetDefaultResponse](#sdkopenstoragepolicysetdefaultresponse)
    - [SdkOpenStoragePolicyUpdateRequest](#sdkopenstoragepolicyupdaterequest)
    - [SdkOpenStoragePolicyUpdateResponse](#sdkopenstoragepolicyupdateresponse)
    - [SdkRole](#sdkrole)
    - [SdkRoleCreateRequest](#sdkrolecreaterequest)
    - [SdkRoleCreateResponse](#sdkrolecreateresponse)
    - [SdkRoleDeleteRequest](#sdkroledeleterequest)
    - [SdkRoleDeleteResponse](#sdkroledeleteresponse)
    - [SdkRoleEnumerateRequest](#sdkroleenumeraterequest)
    - [SdkRoleEnumerateResponse](#sdkroleenumerateresponse)
    - [SdkRoleInspectRequest](#sdkroleinspectrequest)
    - [SdkRoleInspectResponse](#sdkroleinspectresponse)
    - [SdkRoleUpdateRequest](#sdkroleupdaterequest)
    - [SdkRoleUpdateResponse](#sdkroleupdateresponse)
    - [SdkRule](#sdkrule)
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
    - [SdkSchedulePolicyIntervalPeriodic](#sdkschedulepolicyintervalperiodic)
    - [SdkSchedulePolicyIntervalWeekly](#sdkschedulepolicyintervalweekly)
    - [SdkSchedulePolicyUpdateRequest](#sdkschedulepolicyupdaterequest)
    - [SdkSchedulePolicyUpdateResponse](#sdkschedulepolicyupdateresponse)
    - [SdkServiceCapability](#sdkservicecapability)
    - [SdkServiceCapability.OpenStorageService](#sdkservicecapabilityopenstorageservice)
    - [SdkStoragePolicy](#sdkstoragepolicy)
    - [SdkStoragePool](#sdkstoragepool)
    - [SdkStoragePoolResizeRequest](#sdkstoragepoolresizerequest)
    - [SdkStoragePoolResizeResponse](#sdkstoragepoolresizeresponse)
    - [SdkStorageRebalanceRequest](#sdkstoragerebalancerequest)
    - [SdkStorageRebalanceResponse](#sdkstoragerebalanceresponse)
    - [SdkUpdateJobRequest](#sdkupdatejobrequest)
    - [SdkUpdateJobResponse](#sdkupdatejobresponse)
    - [SdkUpdateRebalanceJobRequest](#sdkupdaterebalancejobrequest)
    - [SdkUpdateRebalanceJobResponse](#sdkupdaterebalancejobresponse)
    - [SdkVersion](#sdkversion)
    - [SdkVolumeAttachOptions](#sdkvolumeattachoptions)
    - [SdkVolumeAttachRequest](#sdkvolumeattachrequest)
    - [SdkVolumeAttachRequest.DriverOptionsEntry](#sdkvolumeattachrequestdriveroptionsentry)
    - [SdkVolumeAttachResponse](#sdkvolumeattachresponse)
    - [SdkVolumeCapacityUsageRequest](#sdkvolumecapacityusagerequest)
    - [SdkVolumeCapacityUsageResponse](#sdkvolumecapacityusageresponse)
    - [SdkVolumeCatalogRequest](#sdkvolumecatalogrequest)
    - [SdkVolumeCatalogResponse](#sdkvolumecatalogresponse)
    - [SdkVolumeCloneRequest](#sdkvolumeclonerequest)
    - [SdkVolumeCloneResponse](#sdkvolumecloneresponse)
    - [SdkVolumeCreateRequest](#sdkvolumecreaterequest)
    - [SdkVolumeCreateRequest.LabelsEntry](#sdkvolumecreaterequestlabelsentry)
    - [SdkVolumeCreateResponse](#sdkvolumecreateresponse)
    - [SdkVolumeDeleteRequest](#sdkvolumedeleterequest)
    - [SdkVolumeDeleteResponse](#sdkvolumedeleteresponse)
    - [SdkVolumeDetachOptions](#sdkvolumedetachoptions)
    - [SdkVolumeDetachRequest](#sdkvolumedetachrequest)
    - [SdkVolumeDetachRequest.DriverOptionsEntry](#sdkvolumedetachrequestdriveroptionsentry)
    - [SdkVolumeDetachResponse](#sdkvolumedetachresponse)
    - [SdkVolumeEnumerateRequest](#sdkvolumeenumeraterequest)
    - [SdkVolumeEnumerateResponse](#sdkvolumeenumerateresponse)
    - [SdkVolumeEnumerateWithFiltersRequest](#sdkvolumeenumeratewithfiltersrequest)
    - [SdkVolumeEnumerateWithFiltersRequest.LabelsEntry](#sdkvolumeenumeratewithfiltersrequestlabelsentry)
    - [SdkVolumeEnumerateWithFiltersResponse](#sdkvolumeenumeratewithfiltersresponse)
    - [SdkVolumeInspectRequest](#sdkvolumeinspectrequest)
    - [SdkVolumeInspectResponse](#sdkvolumeinspectresponse)
    - [SdkVolumeInspectResponse.LabelsEntry](#sdkvolumeinspectresponselabelsentry)
    - [SdkVolumeInspectWithFiltersRequest](#sdkvolumeinspectwithfiltersrequest)
    - [SdkVolumeInspectWithFiltersRequest.LabelsEntry](#sdkvolumeinspectwithfiltersrequestlabelsentry)
    - [SdkVolumeInspectWithFiltersResponse](#sdkvolumeinspectwithfiltersresponse)
    - [SdkVolumeMountRequest](#sdkvolumemountrequest)
    - [SdkVolumeMountRequest.DriverOptionsEntry](#sdkvolumemountrequestdriveroptionsentry)
    - [SdkVolumeMountResponse](#sdkvolumemountresponse)
    - [SdkVolumeSnapshotCreateRequest](#sdkvolumesnapshotcreaterequest)
    - [SdkVolumeSnapshotCreateRequest.LabelsEntry](#sdkvolumesnapshotcreaterequestlabelsentry)
    - [SdkVolumeSnapshotCreateResponse](#sdkvolumesnapshotcreateresponse)
    - [SdkVolumeSnapshotEnumerateRequest](#sdkvolumesnapshotenumeraterequest)
    - [SdkVolumeSnapshotEnumerateResponse](#sdkvolumesnapshotenumerateresponse)
    - [SdkVolumeSnapshotEnumerateWithFiltersRequest](#sdkvolumesnapshotenumeratewithfiltersrequest)
    - [SdkVolumeSnapshotEnumerateWithFiltersRequest.LabelsEntry](#sdkvolumesnapshotenumeratewithfiltersrequestlabelsentry)
    - [SdkVolumeSnapshotEnumerateWithFiltersResponse](#sdkvolumesnapshotenumeratewithfiltersresponse)
    - [SdkVolumeSnapshotRestoreRequest](#sdkvolumesnapshotrestorerequest)
    - [SdkVolumeSnapshotRestoreResponse](#sdkvolumesnapshotrestoreresponse)
    - [SdkVolumeSnapshotScheduleUpdateRequest](#sdkvolumesnapshotscheduleupdaterequest)
    - [SdkVolumeSnapshotScheduleUpdateResponse](#sdkvolumesnapshotscheduleupdateresponse)
    - [SdkVolumeStatsRequest](#sdkvolumestatsrequest)
    - [SdkVolumeStatsResponse](#sdkvolumestatsresponse)
    - [SdkVolumeUnmountOptions](#sdkvolumeunmountoptions)
    - [SdkVolumeUnmountRequest](#sdkvolumeunmountrequest)
    - [SdkVolumeUnmountRequest.DriverOptionsEntry](#sdkvolumeunmountrequestdriveroptionsentry)
    - [SdkVolumeUnmountResponse](#sdkvolumeunmountresponse)
    - [SdkVolumeUpdateRequest](#sdkvolumeupdaterequest)
    - [SdkVolumeUpdateRequest.LabelsEntry](#sdkvolumeupdaterequestlabelsentry)
    - [SdkVolumeUpdateResponse](#sdkvolumeupdateresponse)
    - [Sharedv4FailoverStrategy](#sharedv4failoverstrategy)
    - [Sharedv4ServiceSpec](#sharedv4servicespec)
    - [Sharedv4Spec](#sharedv4spec)
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
    - [StoragePoolOperation](#storagepooloperation)
    - [StoragePoolOperation.ParamsEntry](#storagepooloperationparamsentry)
    - [StorageRebalanceAudit](#storagerebalanceaudit)
    - [StorageRebalanceJob](#storagerebalancejob)
    - [StorageRebalanceSummary](#storagerebalancesummary)
    - [StorageRebalanceTriggerThreshold](#storagerebalancetriggerthreshold)
    - [StorageRebalanceWorkSummary](#storagerebalanceworksummary)
    - [StorageResource](#storageresource)
    - [StorageVersion](#storageversion)
    - [StorageVersion.DetailsEntry](#storageversiondetailsentry)
    - [TopologyRequirement](#topologyrequirement)
    - [TopologyRequirement.LabelsEntry](#topologyrequirementlabelsentry)
    - [Volume](#volume)
    - [Volume.AttachInfoEntry](#volumeattachinfoentry)
    - [VolumeConsumer](#volumeconsumer)
    - [VolumeCreateRequest](#volumecreaterequest)
    - [VolumeCreateResponse](#volumecreateresponse)
    - [VolumeInfo](#volumeinfo)
    - [VolumeInspectOptions](#volumeinspectoptions)
    - [VolumeLocator](#volumelocator)
    - [VolumeLocator.VolumeLabelsEntry](#volumelocatorvolumelabelsentry)
    - [VolumePlacementSpec](#volumeplacementspec)
    - [VolumePlacementStrategy](#volumeplacementstrategy)
    - [VolumeResponse](#volumeresponse)
    - [VolumeServiceInstanceResponse](#volumeserviceinstanceresponse)
    - [VolumeServiceInstanceResponse.StatusEntry](#volumeserviceinstanceresponsestatusentry)
    - [VolumeServiceRequest](#volumeservicerequest)
    - [VolumeServiceRequest.SrvCmdParamsEntry](#volumeservicerequestsrvcmdparamsentry)
    - [VolumeServiceResponse](#volumeserviceresponse)
    - [VolumeSetRequest](#volumesetrequest)
    - [VolumeSetRequest.OptionsEntry](#volumesetrequestoptionsentry)
    - [VolumeSetResponse](#volumesetresponse)
    - [VolumeSpec](#volumespec)
    - [VolumeSpec.VolumeLabelsEntry](#volumespecvolumelabelsentry)
    - [VolumeSpecPolicy](#volumespecpolicy)
    - [VolumeSpecPolicy.VolumeLabelsEntry](#volumespecpolicyvolumelabelsentry)
    - [VolumeSpecUpdate](#volumespecupdate)
    - [VolumeStateAction](#volumestateaction)
    - [VolumeUsage](#volumeusage)
    - [VolumeUsageByNode](#volumeusagebynode)
    - [Xattr](#xattr)
  


- Enums
    - [AlertActionType](#alertactiontype)
    - [AnonymousBucketAccessMode](#anonymousbucketaccessmode)
    - [AttachState](#attachstate)
    - [CloudMigrate.OperationType](#cloudmigrateoperationtype)
    - [CloudMigrate.Stage](#cloudmigratestage)
    - [CloudMigrate.Status](#cloudmigratestatus)
    - [ClusterPairMode.Mode](#clusterpairmodemode)
    - [CosType](#costype)
    - [DiagsCollectionStatus.State](#diagscollectionstatusstate)
    - [DriverType](#drivertype)
    - [EnforcementType](#enforcementtype)
    - [ExportProtocol](#exportprotocol)
    - [FSType](#fstype)
    - [FastpathProtocol](#fastpathprotocol)
    - [FastpathStatus](#fastpathstatus)
    - [FilesystemCheck.FilesystemCheckStatus](#filesystemcheckfilesystemcheckstatus)
    - [FilesystemHealthStatus](#filesystemhealthstatus)
    - [FilesystemTrim.FilesystemTrimStatus](#filesystemtrimfilesystemtrimstatus)
    - [GraphDriverChangeType](#graphdriverchangetype)
    - [HardwareType](#hardwaretype)
    - [IoProfile](#ioprofile)
    - [Job.State](#jobstate)
    - [Job.Type](#jobtype)
    - [LabelSelectorRequirement.Operator](#labelselectorrequirementoperator)
    - [OperationFlags](#operationflags)
    - [Ownership.AccessType](#ownershipaccesstype)
    - [ProxyProtocol](#proxyprotocol)
    - [ResourceType](#resourcetype)
    - [RestoreParamBoolType](#restoreparambooltype)
    - [ScanPolicy.ScanAction](#scanpolicyscanaction)
    - [ScanPolicy.ScanTrigger](#scanpolicyscantrigger)
    - [SdkCloudBackupClusterType](#sdkcloudbackupclustertype)
    - [SdkCloudBackupOpType](#sdkcloudbackupoptype)
    - [SdkCloudBackupRequestedState](#sdkcloudbackuprequestedstate)
    - [SdkCloudBackupStatusType](#sdkcloudbackupstatustype)
    - [SdkServiceCapability.OpenStorageService.Type](#sdkservicecapabilityopenstorageservicetype)
    - [SdkStoragePool.OperationStatus](#sdkstoragepooloperationstatus)
    - [SdkStoragePool.OperationType](#sdkstoragepooloperationtype)
    - [SdkStoragePool.ResizeOperationType](#sdkstoragepoolresizeoperationtype)
    - [SdkTimeWeekday](#sdktimeweekday)
    - [SdkVersion.Version](#sdkversionversion)
    - [SeverityType](#severitytype)
    - [Sharedv4FailoverStrategy.Value](#sharedv4failoverstrategyvalue)
    - [Sharedv4ServiceSpec.ServiceType](#sharedv4servicespecservicetype)
    - [Status](#status)
    - [StorageMedium](#storagemedium)
    - [StorageNode.SecurityStatus](#storagenodesecuritystatus)
    - [StorageRebalanceAudit.StorageRebalanceAction](#storagerebalanceauditstoragerebalanceaction)
    - [StorageRebalanceJobState](#storagerebalancejobstate)
    - [StorageRebalanceTriggerThreshold.Metric](#storagerebalancetriggerthresholdmetric)
    - [StorageRebalanceTriggerThreshold.Type](#storagerebalancetriggerthresholdtype)
    - [StorageRebalanceWorkSummary.Type](#storagerebalanceworksummarytype)
    - [VolumeActionParam](#volumeactionparam)
    - [VolumeSpecPolicy.PolicyOp](#volumespecpolicypolicyop)
    - [VolumeState](#volumestate)
    - [VolumeStatus](#volumestatus)
    - [Xattr.Value](#xattrvalue)
  


- [Scalar Value Types](#scalar-value-types)




# OpenStorageAlerts {#serviceopenstorageapiopenstoragealerts}
OpenStorageAlerts defines rpc's for alerts.

## EnumerateWithFilters {#methodopenstorageapiopenstoragealertsenumeratewithfilters}

> **rpc** EnumerateWithFilters([SdkAlertsEnumerateWithFiltersRequest](#sdkalertsenumeratewithfiltersrequest))
    [SdkAlertsEnumerateWithFiltersResponse](#sdkalertsenumeratewithfiltersresponse)

Allows querying alerts.

EnumerateWithFilters allows 3 different types of queries as defined below:

* Query that takes only resource type as input
* Query that takes resource type and alert type as input and
* Query that takes resource id, alert type and resource type as input.

#### Input
SdkAlertsEnumerateRequest takes a list of such queries and the returned
output is a collective output from each of these queries. In that sense,
the filtering of these queries has a behavior of OR operation.
Each query also has a list of optional options. These options allow
narrowing down the scope of alerts search. These options have a
behavior of an AND operation.

#### Examples
To search by a resource type in a given time window would require
initializing SdkAlertsResourceTypeQuery query and pass in
SdkAlertsTimeSpan option into SdkAlertsQuery struct and finally
packing any other such queries into SdkAlertsEnumerateRequest object.
Alternatively, to search by both resource type and alert type, use
SdkAlertsAlertTypeQuery as query builder.
Finally to search all alerts of a given resource type and some
alerts of another resource type but with specific alert type,
use two queries, first initialized with SdkAlertsResourceTypeQuery
and second initialized with SdkAlertsAlertTypeQuery and both
eventually packed as list in SdkAlertsEnumerateRequest.
## Delete {#methodopenstorageapiopenstoragealertsdelete}

> **rpc** Delete([SdkAlertsDeleteRequest](#sdkalertsdeleterequest))
    [SdkAlertsDeleteResponse](#sdkalertsdeleteresponse)

Delete alerts

#### Delete
Delete allows 3 different types of queries as defined below:

* Query that takes only resource type as input
* Query that takes resource type and alert type as input and
* Query that takes resource id, alert type and resource type as input.

#### Input
SdkAlertsDeleteRequest takes a list of such queries and all alerts
that match at least one of the queries are deleted.
 <!-- end methods -->

# OpenStorageBucket {#serviceopenstorageapiopenstoragebucket}
BucketService to manage the bucket driver

## Create {#methodopenstorageapiopenstoragebucketcreate}

> **rpc** Create([BucketCreateRequest](#bucketcreaterequest))
    [BucketCreateResponse](#bucketcreateresponse)


## Delete {#methodopenstorageapiopenstoragebucketdelete}

> **rpc** Delete([BucketDeleteRequest](#bucketdeleterequest))
    [BucketDeleteResponse](#bucketdeleteresponse)


## GrantAccess {#methodopenstorageapiopenstoragebucketgrantaccess}

> **rpc** GrantAccess([BucketGrantAccessRequest](#bucketgrantaccessrequest))
    [BucketGrantAccessResponse](#bucketgrantaccessresponse)


## RevokeAccess {#methodopenstorageapiopenstoragebucketrevokeaccess}

> **rpc** RevokeAccess([BucketRevokeAccessRequest](#bucketrevokeaccessrequest))
    [BucketRevokeAccessResponse](#bucketrevokeaccessresponse)


 <!-- end methods -->

# OpenStorageCloudBackup {#serviceopenstorageapiopenstoragecloudbackup}
OpenStorageCloudBackup service manages backing up volumes to a cloud
location like Amazon, Google, or Azure.

#### Backup
To create a backup, you must first call the Create() call for a specified
volume. To see the status of this request, use Status() which returns
a map where the keys are the source volume id.

#### Restore
To restore, you would pass a `backup_id` of a successful backup.
`backup_id` can be retreived by calling Enumerate() for a specified volume.
Pass this `backup_id` and a new volume name to Restore() to start
restoring a new volume from an existing backup. To see the status of this
restore, pass volume id returned by Restore() to input to Status()

## Create {#methodopenstorageapiopenstoragecloudbackupcreate}

> **rpc** Create([SdkCloudBackupCreateRequest](#sdkcloudbackupcreaterequest))
    [SdkCloudBackupCreateResponse](#sdkcloudbackupcreateresponse)

Creates a backup request for a specified volume. Use
OpenStorageCloudBackup.Status() to get the current status of the
backup request.

Requires access AccessType.Read of volume
## GroupCreate {#methodopenstorageapiopenstoragecloudbackupgroupcreate}

> **rpc** GroupCreate([SdkCloudBackupGroupCreateRequest](#sdkcloudbackupgroupcreaterequest))
    [SdkCloudBackupGroupCreateResponse](#sdkcloudbackupgroupcreateresponse)

Creates a group backup request for a specified group. Use
OpenStorageCloudBackup.Status() to get the current status of the
backup request.

This will only backup volumes which the user has read_access to.
## Restore {#methodopenstorageapiopenstoragecloudbackuprestore}

> **rpc** Restore([SdkCloudBackupRestoreRequest](#sdkcloudbackuprestorerequest))
    [SdkCloudBackupRestoreResponse](#sdkcloudbackuprestoreresponse)

Restore creates a new volume from a backup id. The newly created volume
has an ha_level (number of replicas) of only 1. To increase the number of
replicas, use OpenStorageVolume.Set() to change the ha_level.
## Delete {#methodopenstorageapiopenstoragecloudbackupdelete}

> **rpc** Delete([SdkCloudBackupDeleteRequest](#sdkcloudbackupdeleterequest))
    [SdkCloudBackupDeleteResponse](#sdkcloudbackupdeleteresponse)

Deletes a backup stored in the cloud. If the backup is an incremental
backup and other backups are dependent on it, it will not be able to be deleted.
## DeleteAll {#methodopenstorageapiopenstoragecloudbackupdeleteall}

> **rpc** DeleteAll([SdkCloudBackupDeleteAllRequest](#sdkcloudbackupdeleteallrequest))
    [SdkCloudBackupDeleteAllResponse](#sdkcloudbackupdeleteallresponse)

DeleteAll deletes all the backups in the cloud for the specified volume.
## EnumerateWithFilters {#methodopenstorageapiopenstoragecloudbackupenumeratewithfilters}

> **rpc** EnumerateWithFilters([SdkCloudBackupEnumerateWithFiltersRequest](#sdkcloudbackupenumeratewithfiltersrequest))
    [SdkCloudBackupEnumerateWithFiltersResponse](#sdkcloudbackupenumeratewithfiltersresponse)

Return a list of backups for the specified volume
## Status {#methodopenstorageapiopenstoragecloudbackupstatus}

> **rpc** Status([SdkCloudBackupStatusRequest](#sdkcloudbackupstatusrequest))
    [SdkCloudBackupStatusResponse](#sdkcloudbackupstatusresponse)

Status returns the status of any cloud backups of a volume
## Catalog {#methodopenstorageapiopenstoragecloudbackupcatalog}

> **rpc** Catalog([SdkCloudBackupCatalogRequest](#sdkcloudbackupcatalogrequest))
    [SdkCloudBackupCatalogResponse](#sdkcloudbackupcatalogresponse)

Catalog returns a list of the contents in the backup
## History {#methodopenstorageapiopenstoragecloudbackuphistory}

> **rpc** History([SdkCloudBackupHistoryRequest](#sdkcloudbackuphistoryrequest))
    [SdkCloudBackupHistoryResponse](#sdkcloudbackuphistoryresponse)

History returns a list of backups for a specified volume
## StateChange {#methodopenstorageapiopenstoragecloudbackupstatechange}

> **rpc** StateChange([SdkCloudBackupStateChangeRequest](#sdkcloudbackupstatechangerequest))
    [SdkCloudBackupStateChangeResponse](#sdkcloudbackupstatechangeresponse)

StateChange can be used to stop, pause, and restart a backup
## SchedCreate {#methodopenstorageapiopenstoragecloudbackupschedcreate}

> **rpc** SchedCreate([SdkCloudBackupSchedCreateRequest](#sdkcloudbackupschedcreaterequest))
    [SdkCloudBackupSchedCreateResponse](#sdkcloudbackupschedcreateresponse)

Create cloud backup schedule
## SchedUpdate {#methodopenstorageapiopenstoragecloudbackupschedupdate}

> **rpc** SchedUpdate([SdkCloudBackupSchedUpdateRequest](#sdkcloudbackupschedupdaterequest))
    [SdkCloudBackupSchedUpdateResponse](#sdkcloudbackupschedupdateresponse)

Update existing cloud backup schedule
## SchedDelete {#methodopenstorageapiopenstoragecloudbackupscheddelete}

> **rpc** SchedDelete([SdkCloudBackupSchedDeleteRequest](#sdkcloudbackupscheddeleterequest))
    [SdkCloudBackupSchedDeleteResponse](#sdkcloudbackupscheddeleteresponse)

Delete cloud backup schedule
## SchedEnumerate {#methodopenstorageapiopenstoragecloudbackupschedenumerate}

> **rpc** SchedEnumerate([SdkCloudBackupSchedEnumerateRequest](#sdkcloudbackupschedenumeraterequest))
    [SdkCloudBackupSchedEnumerateResponse](#sdkcloudbackupschedenumerateresponse)

Enumerate cloud backup schedules
## Size {#methodopenstorageapiopenstoragecloudbackupsize}

> **rpc** Size([SdkCloudBackupSizeRequest](#sdkcloudbackupsizerequest))
    [SdkCloudBackupSizeResponse](#sdkcloudbackupsizeresponse)

Size returns the size of any cloud backups of a volume
 <!-- end methods -->

# OpenStorageCluster {#serviceopenstorageapiopenstoragecluster}
OpenStorageCluster service provides the methods to manage the cluster

## InspectCurrent {#methodopenstorageapiopenstorageclusterinspectcurrent}

> **rpc** InspectCurrent([SdkClusterInspectCurrentRequest](#sdkclusterinspectcurrentrequest))
    [SdkClusterInspectCurrentResponse](#sdkclusterinspectcurrentresponse)

InspectCurrent returns information about the current cluster
 <!-- end methods -->

# OpenStorageClusterDomains {#serviceopenstorageapiopenstorageclusterdomains}
OpenStorageClusterDomains is a service used to manage cluster domains in an openstorage cluster.
A single openstorage cluster can stretch across multiple cluster domains. In other words it can
span across multiple Container Orchestrator clusters. Each node in the cluster is assigned a
a cluster domain to start with. A cluster domain and in turn the nodes which are a part of it
are assigned a status indicating whether the cluster domain is actively participating in the
cluster or whether it is inactive.

## Enumerate {#methodopenstorageapiopenstorageclusterdomainsenumerate}

> **rpc** Enumerate([SdkClusterDomainsEnumerateRequest](#sdkclusterdomainsenumeraterequest))
    [SdkClusterDomainsEnumerateResponse](#sdkclusterdomainsenumerateresponse)

Enumerate returns names of all the cluster domains in the cluster
## Inspect {#methodopenstorageapiopenstorageclusterdomainsinspect}

> **rpc** Inspect([SdkClusterDomainInspectRequest](#sdkclusterdomaininspectrequest))
    [SdkClusterDomainInspectResponse](#sdkclusterdomaininspectresponse)

Inspect returns information about a cluster domain and a
status indicating whether the cluster domain is active
## Activate {#methodopenstorageapiopenstorageclusterdomainsactivate}

> **rpc** Activate([SdkClusterDomainActivateRequest](#sdkclusterdomainactivaterequest))
    [SdkClusterDomainActivateResponse](#sdkclusterdomainactivateresponse)

Activates a cluster domain in the cluster.
All the nodes which are part of an active cluster domain
will participate in cluster quorum calculation
## Deactivate {#methodopenstorageapiopenstorageclusterdomainsdeactivate}

> **rpc** Deactivate([SdkClusterDomainDeactivateRequest](#sdkclusterdomaindeactivaterequest))
    [SdkClusterDomainDeactivateResponse](#sdkclusterdomaindeactivateresponse)

Deactivates a cluster domain in the cluster.
All the nodes which are part of a deactivated cluster domain.
will not participate in cluster quorum calculation
 <!-- end methods -->

# OpenStorageClusterPair {#serviceopenstorageapiopenstorageclusterpair}
OpenStorageClusterPair service provides the methods to manage a cluster pair

## Create {#methodopenstorageapiopenstorageclusterpaircreate}

> **rpc** Create([SdkClusterPairCreateRequest](#sdkclusterpaircreaterequest))
    [SdkClusterPairCreateResponse](#sdkclusterpaircreateresponse)

Creates Pair with a remote cluster and returns details about the remote cluster

##### Example
{% codetabs name="Golang", type="go" -%}
id, err := client.Create(context.Background(), &api.SdkClusterPairCreateRequest {
  Request : &api.ClusterPairCreateRequest {
                 RemoteClusterIp: "127.0.0.1",
                 RemoteClusterPort: 12345,
                 RemoteClusterToken: "<Auth-Token>",
                 SetDefault: true,
             }
       })
{%- endcodetabs %}
## Inspect {#methodopenstorageapiopenstorageclusterpairinspect}

> **rpc** Inspect([SdkClusterPairInspectRequest](#sdkclusterpairinspectrequest))
    [SdkClusterPairInspectResponse](#sdkclusterpairinspectresponse)

Inspect information about a cluster pair
## Enumerate {#methodopenstorageapiopenstorageclusterpairenumerate}

> **rpc** Enumerate([SdkClusterPairEnumerateRequest](#sdkclusterpairenumeraterequest))
    [SdkClusterPairEnumerateResponse](#sdkclusterpairenumerateresponse)

Enumerate returns list of cluster pairs
## GetToken {#methodopenstorageapiopenstorageclusterpairgettoken}

> **rpc** GetToken([SdkClusterPairGetTokenRequest](#sdkclusterpairgettokenrequest))
    [SdkClusterPairGetTokenResponse](#sdkclusterpairgettokenresponse)

GetToken returns a auth token
## ResetToken {#methodopenstorageapiopenstorageclusterpairresettoken}

> **rpc** ResetToken([SdkClusterPairResetTokenRequest](#sdkclusterpairresettokenrequest))
    [SdkClusterPairResetTokenResponse](#sdkclusterpairresettokenresponse)

ResetToken returns a auth token
## Delete {#methodopenstorageapiopenstorageclusterpairdelete}

> **rpc** Delete([SdkClusterPairDeleteRequest](#sdkclusterpairdeleterequest))
    [SdkClusterPairDeleteResponse](#sdkclusterpairdeleteresponse)

Delete a cluster pair
 <!-- end methods -->

# OpenStorageCredentials {#serviceopenstorageapiopenstoragecredentials}
OpenStorageCredentials is a service used to manage the cloud credentials
which can then be used by the OpenStorageCloudBackup service

## Create {#methodopenstorageapiopenstoragecredentialscreate}

> **rpc** Create([SdkCredentialCreateRequest](#sdkcredentialcreaterequest))
    [SdkCredentialCreateResponse](#sdkcredentialcreateresponse)

Create is used to submit cloud credentials. It will return an
id of the credentials once they are verified to work.

##### Example
{% codetabs name="Golang", type="go" -%}
id, err := client.Create(context.Background(), &api.SdkCredentialCreateRequest{
    Name: "awscred",
    UseProxy: false,
    CredentialType: &api.SdkCredentialCreateRequest_AwsCredential{
      AwsCredential: &api.SdkAwsCredentialRequest{
      AccessKey: "dummy-access",
      SecretKey: "dummy-secret",
      Endpoint:  "dummy-endpoint",
      Region:    "dummy-region",
    },
  },
})
{%- language name="Python", type="py" -%}
en_resp = client.Create(api_pb2.SdkCredentialCreateRequest(
  name='awscred',
  aws_credential=api_pb2.SdkAwsCredentialRequest(
    access_key='dummy-access',
    secret_key='dumm-secret',
    endpoint='dummy-endpoint',
    region='dummy-region')))
{%- endcodetabs %}
## Update {#methodopenstorageapiopenstoragecredentialsupdate}

> **rpc** Update([SdkCredentialUpdateRequest](#sdkcredentialupdaterequest))
    [SdkCredentialUpdateResponse](#sdkcredentialupdateresponse)

input is very same as credential create
## Enumerate {#methodopenstorageapiopenstoragecredentialsenumerate}

> **rpc** Enumerate([SdkCredentialEnumerateRequest](#sdkcredentialenumeraterequest))
    [SdkCredentialEnumerateResponse](#sdkcredentialenumerateresponse)

Enumerate returns a list of credential ids
## Inspect {#methodopenstorageapiopenstoragecredentialsinspect}

> **rpc** Inspect([SdkCredentialInspectRequest](#sdkcredentialinspectrequest))
    [SdkCredentialInspectResponse](#sdkcredentialinspectresponse)

Inspect returns the information about a credential, but does not return the secret key.
## Delete {#methodopenstorageapiopenstoragecredentialsdelete}

> **rpc** Delete([SdkCredentialDeleteRequest](#sdkcredentialdeleterequest))
    [SdkCredentialDeleteResponse](#sdkcredentialdeleteresponse)

Delete a specified credential
## Validate {#methodopenstorageapiopenstoragecredentialsvalidate}

> **rpc** Validate([SdkCredentialValidateRequest](#sdkcredentialvalidaterequest))
    [SdkCredentialValidateResponse](#sdkcredentialvalidateresponse)

Validate is used to validate credentials
## DeleteReferences {#methodopenstorageapiopenstoragecredentialsdeletereferences}

> **rpc** DeleteReferences([SdkCredentialDeleteReferencesRequest](#sdkcredentialdeletereferencesrequest))
    [SdkCredentialDeleteReferencesResponse](#sdkcredentialdeletereferencesresponse)

DeleteReferences is used to remove references to credentials
 <!-- end methods -->

# OpenStorageDiags {#serviceopenstorageapiopenstoragediags}
OpenStorageDiags service provides methods to manage diagnostic bundles

## Collect {#methodopenstorageapiopenstoragediagscollect}

> **rpc** Collect([SdkDiagsCollectRequest](#sdkdiagscollectrequest))
    [SdkDiagsCollectResponse](#sdkdiagscollectresponse)

Collect starts a job to collect diagnostics from set of nodes that are selected based on the selectors provided
in the SdkDiagsCollectRequest. See SdkDiagsCollectRequest for more details on how to select the nodes
Returns SdkDiagsCollectResponse which has the job that is responsible for collecting the diags.
 <!-- end methods -->

# OpenStorageFilesystemCheck {#serviceopenstorageapiopenstoragefilesystemcheck}


## Start {#methodopenstorageapiopenstoragefilesystemcheckstart}

> **rpc** Start([SdkFilesystemCheckStartRequest](#sdkfilesystemcheckstartrequest))
    [SdkFilesystemCheckStartResponse](#sdkfilesystemcheckstartresponse)

Start a filesystem-check background operation on a unmounted volume.
## Status {#methodopenstorageapiopenstoragefilesystemcheckstatus}

> **rpc** Status([SdkFilesystemCheckStatusRequest](#sdkfilesystemcheckstatusrequest))
    [SdkFilesystemCheckStatusResponse](#sdkfilesystemcheckstatusresponse)

Get Status of a filesystem-check background operation on an unmounted
volume, if any
## Stop {#methodopenstorageapiopenstoragefilesystemcheckstop}

> **rpc** Stop([SdkFilesystemCheckStopRequest](#sdkfilesystemcheckstoprequest))
    [SdkFilesystemCheckStopResponse](#sdkfilesystemcheckstopresponse)

Stop a filesystem check background operation on an unmounted volume, if any
 <!-- end methods -->

# OpenStorageFilesystemTrim {#serviceopenstorageapiopenstoragefilesystemtrim}
## OpenStorageFilesystemTrim Service
This service provides methods to manage filesystem trim operation on a
volume.

This operation runs in the background on a **mounted volume**. If the volumes
are not mounted, these API return error.

Once the filesystem trim operation is started, the clients have to poll for
the status of the background operation using the
`OpenStorageFilesystemTrim.Status()` rpc request

A typical workflow involving filesystem trim would be as follows
1. Attach the volume
   `OpenStorageMountAttachClient.Attach()`
2. Mount the volume
   `OpenStorageMountAttachClient.Mount()`
3. Start the filesystem trim operation by issuing a grpc call to
   `OpenStorageFilesystemTrimClient.Start()`
   This call returns immediately with a status code indicating if the
   operation was successfully started or not.
4. To get the status of the Filesystem Trim operation, issue a grpc call to
   `OpenStorageFilesystemTrimClient.Status()`
5. To stop the Filesystem Trim operation, issue a grpc call to
   `OpenStorageFilesystemTrimClient.Stop()`

## Start {#methodopenstorageapiopenstoragefilesystemtrimstart}

> **rpc** Start([SdkFilesystemTrimStartRequest](#sdkfilesystemtrimstartrequest))
    [SdkFilesystemTrimStartResponse](#sdkfilesystemtrimstartresponse)

Start a filesystem Trim background operation on a mounted volume
## Status {#methodopenstorageapiopenstoragefilesystemtrimstatus}

> **rpc** Status([SdkFilesystemTrimStatusRequest](#sdkfilesystemtrimstatusrequest))
    [SdkFilesystemTrimStatusResponse](#sdkfilesystemtrimstatusresponse)

Status of a filesystem Trim background operation on a mounted
volume, if any
## AutoFSTrimStatus {#methodopenstorageapiopenstoragefilesystemtrimautofstrimstatus}

> **rpc** AutoFSTrimStatus([SdkAutoFSTrimStatusRequest](#sdkautofstrimstatusrequest))
    [SdkAutoFSTrimStatusResponse](#sdkautofstrimstatusresponse)


## AutoFSTrimUsage {#methodopenstorageapiopenstoragefilesystemtrimautofstrimusage}

> **rpc** AutoFSTrimUsage([SdkAutoFSTrimUsageRequest](#sdkautofstrimusagerequest))
    [SdkAutoFSTrimUsageResponse](#sdkautofstrimusageresponse)

Usage of a filesystem Trim background operation on all locally mounted
volume
## Stop {#methodopenstorageapiopenstoragefilesystemtrimstop}

> **rpc** Stop([SdkFilesystemTrimStopRequest](#sdkfilesystemtrimstoprequest))
    [SdkFilesystemTrimStopResponse](#sdkfilesystemtrimstopresponse)

Stop a filesystem Trim background operation on a mounted volume, if any
 <!-- end methods -->

# OpenStorageIdentity {#serviceopenstorageapiopenstorageidentity}
OpenStorageIdentity service provides methods to obtain information
about the cluster

## Capabilities {#methodopenstorageapiopenstorageidentitycapabilities}

> **rpc** Capabilities([SdkIdentityCapabilitiesRequest](#sdkidentitycapabilitiesrequest))
    [SdkIdentityCapabilitiesResponse](#sdkidentitycapabilitiesresponse)

Capabilities returns the supported services by the cluster.
This allows SDK implementations to advertise their supported
services as the API matures. With this information, clients
can determine supported services from storage clusters at
different versions.
## Version {#methodopenstorageapiopenstorageidentityversion}

> **rpc** Version([SdkIdentityVersionRequest](#sdkidentityversionrequest))
    [SdkIdentityVersionResponse](#sdkidentityversionresponse)

Version returns version information about the system.
 <!-- end methods -->

# OpenStorageJob {#serviceopenstorageapiopenstoragejob}
OpenstorageJob is a service that provides a common set of APIs for services
that use the asynchronous job framework

## Update {#methodopenstorageapiopenstoragejobupdate}

> **rpc** Update([SdkUpdateJobRequest](#sdkupdatejobrequest))
    [SdkUpdateJobResponse](#sdkupdatejobresponse)

Update updates an existing job's state
Only acceptable state values are
JobState_PAUSED - acceptable only from running state
JobState_CANCELLED - acceptable only from running/pause state
JobState_RUNNING - acceptable only from pause state
## GetStatus {#methodopenstorageapiopenstoragejobgetstatus}

> **rpc** GetStatus([SdkGetJobStatusRequest](#sdkgetjobstatusrequest))
    [SdkGetJobStatusResponse](#sdkgetjobstatusresponse)

GetStatus gets the status of a job
## Enumerate {#methodopenstorageapiopenstoragejobenumerate}

> **rpc** Enumerate([SdkEnumerateJobsRequest](#sdkenumeratejobsrequest))
    [SdkEnumerateJobsResponse](#sdkenumeratejobsresponse)

Enumerate returns all the jobs currently known to the system
 <!-- end methods -->

# OpenStorageMigrate {#serviceopenstorageapiopenstoragemigrate}
OpenStorageMigrate is a service used to manage migration of volumes

## Start {#methodopenstorageapiopenstoragemigratestart}

> **rpc** Start([SdkCloudMigrateStartRequest](#sdkcloudmigratestartrequest))
    [SdkCloudMigrateStartResponse](#sdkcloudmigratestartresponse)

Start a migration operation
## Cancel {#methodopenstorageapiopenstoragemigratecancel}

> **rpc** Cancel([SdkCloudMigrateCancelRequest](#sdkcloudmigratecancelrequest))
    [SdkCloudMigrateCancelResponse](#sdkcloudmigratecancelresponse)

Cancel a migration operation
## Status {#methodopenstorageapiopenstoragemigratestatus}

> **rpc** Status([SdkCloudMigrateStatusRequest](#sdkcloudmigratestatusrequest))
    [SdkCloudMigrateStatusResponse](#sdkcloudmigratestatusresponse)

Status for migration operation.
If status request is empty, status for all migration operation will be returned.
 <!-- end methods -->

# OpenStorageMountAttach {#serviceopenstorageapiopenstoragemountattach}
OpenStorageMountAttach is a service used to manage node access to a volume.
Note, these APIs are here for testing or diagnostics purposes only. In normal
operations, the Container Orchestration (CO) system is managing all mount
and attach calls through the CSI interface. The normal usage is once volumes
are created, to let the CO manage the node access functions to the volume.

## Attach {#methodopenstorageapiopenstoragemountattachattach}

> **rpc** Attach([SdkVolumeAttachRequest](#sdkvolumeattachrequest))
    [SdkVolumeAttachResponse](#sdkvolumeattachresponse)

Attach attaches device to the host that the client is communicating with.

Requires access AccessType.Write of volume
## Detach {#methodopenstorageapiopenstoragemountattachdetach}

> **rpc** Detach([SdkVolumeDetachRequest](#sdkvolumedetachrequest))
    [SdkVolumeDetachResponse](#sdkvolumedetachresponse)

Detaches a the volume from the host

Requires access AccessType.Write of volume
## Mount {#methodopenstorageapiopenstoragemountattachmount}

> **rpc** Mount([SdkVolumeMountRequest](#sdkvolumemountrequest))
    [SdkVolumeMountResponse](#sdkvolumemountresponse)

Mount mounts an attached volume in the host that the client is communicating with

Requires access AccessType.Write of volume
## Unmount {#methodopenstorageapiopenstoragemountattachunmount}

> **rpc** Unmount([SdkVolumeUnmountRequest](#sdkvolumeunmountrequest))
    [SdkVolumeUnmountResponse](#sdkvolumeunmountresponse)

Unmount unmounts a mounted volume in the host that the client is communicating with

Requires access AccessType.Write of volume
 <!-- end methods -->

# OpenStorageNode {#serviceopenstorageapiopenstoragenode}
OpenStorageNode is a service used to manage nodes in the cluster

## Inspect {#methodopenstorageapiopenstoragenodeinspect}

> **rpc** Inspect([SdkNodeInspectRequest](#sdknodeinspectrequest))
    [SdkNodeInspectResponse](#sdknodeinspectresponse)

Inspect returns information about the specified node
## InspectCurrent {#methodopenstorageapiopenstoragenodeinspectcurrent}

> **rpc** InspectCurrent([SdkNodeInspectCurrentRequest](#sdknodeinspectcurrentrequest))
    [SdkNodeInspectCurrentResponse](#sdknodeinspectcurrentresponse)

InspectCurrent returns information about the storage node
where the client is currently connected to.
## Enumerate {#methodopenstorageapiopenstoragenodeenumerate}

> **rpc** Enumerate([SdkNodeEnumerateRequest](#sdknodeenumeraterequest))
    [SdkNodeEnumerateResponse](#sdknodeenumerateresponse)

Enumerate returns the ids of all the nodes in the current cluster
## EnumerateWithFilters {#methodopenstorageapiopenstoragenodeenumeratewithfilters}

> **rpc** EnumerateWithFilters([SdkNodeEnumerateWithFiltersRequest](#sdknodeenumeratewithfiltersrequest))
    [SdkNodeEnumerateWithFiltersResponse](#sdknodeenumeratewithfiltersresponse)

EnumerateWithFilters returns all the nodes in the current cluster
## VolumeUsageByNode {#methodopenstorageapiopenstoragenodevolumeusagebynode}

> **rpc** VolumeUsageByNode([SdkNodeVolumeUsageByNodeRequest](#sdknodevolumeusagebynoderequest))
    [SdkNodeVolumeUsageByNodeResponse](#sdknodevolumeusagebynoderesponse)

Returns capacity usage of all volumes/snaps for a give node
## RelaxedReclaimPurge {#methodopenstorageapiopenstoragenoderelaxedreclaimpurge}

> **rpc** RelaxedReclaimPurge([SdkNodeRelaxedReclaimPurgeRequest](#sdknoderelaxedreclaimpurgerequest))
    [SdkNodeRelaxedReclaimPurgeResponse](#sdknoderelaxedreclaimpurgeresponse)

Triggers RelaxedReclaim purge for a give node
## DrainAttachments {#methodopenstorageapiopenstoragenodedrainattachments}

> **rpc** DrainAttachments([SdkNodeDrainAttachmentsRequest](#sdknodedrainattachmentsrequest))
    [SdkJobResponse](#sdkjobresponse)

DrainAttachments creates a task to drain volume attachments
from the provided node in the cluster.
## CordonAttachments {#methodopenstorageapiopenstoragenodecordonattachments}

> **rpc** CordonAttachments([SdkNodeCordonAttachmentsRequest](#sdknodecordonattachmentsrequest))
    [SdkNodeCordonAttachmentsResponse](#sdknodecordonattachmentsresponse)

CordonAttachments disables any new volume attachments
from the provided node in the cluster. Existing volume attachments
will stay on the node.
## UncordonAttachments {#methodopenstorageapiopenstoragenodeuncordonattachments}

> **rpc** UncordonAttachments([SdkNodeUncordonAttachmentsRequest](#sdknodeuncordonattachmentsrequest))
    [SdkNodeUncordonAttachmentsResponse](#sdknodeuncordonattachmentsresponse)

UncordonAttachments re-enables volume attachments
on the provided node in the cluster.
 <!-- end methods -->

# OpenStorageObjectstore {#serviceopenstorageapiopenstorageobjectstore}
OpenStorageObjectstore is a service used to manage object store services on volumes

## Inspect {#methodopenstorageapiopenstorageobjectstoreinspect}

> **rpc** Inspect([SdkObjectstoreInspectRequest](#sdkobjectstoreinspectrequest))
    [SdkObjectstoreInspectResponse](#sdkobjectstoreinspectresponse)

Inspect returns information about the object store endpoint
## Create {#methodopenstorageapiopenstorageobjectstorecreate}

> **rpc** Create([SdkObjectstoreCreateRequest](#sdkobjectstorecreaterequest))
    [SdkObjectstoreCreateResponse](#sdkobjectstorecreateresponse)

Creates creates an object store endpoint on specified volume
## Delete {#methodopenstorageapiopenstorageobjectstoredelete}

> **rpc** Delete([SdkObjectstoreDeleteRequest](#sdkobjectstoredeleterequest))
    [SdkObjectstoreDeleteResponse](#sdkobjectstoredeleteresponse)

Delete destroys the object store endpoint on the volume
## Update {#methodopenstorageapiopenstorageobjectstoreupdate}

> **rpc** Update([SdkObjectstoreUpdateRequest](#sdkobjectstoreupdaterequest))
    [SdkObjectstoreUpdateResponse](#sdkobjectstoreupdateresponse)

Updates provided objectstore status.
This call can be used to stop and start the server while maintaining the same
object storage id.
 <!-- end methods -->

# OpenStoragePolicy {#serviceopenstorageapiopenstoragepolicy}
OpenStoragePolicy service provides methods to manager storage policies.

Policies can be used to validate/ensure a set of volume configurations to be followed
while creating volumes.

## Create {#methodopenstorageapiopenstoragepolicycreate}

> **rpc** Create([SdkOpenStoragePolicyCreateRequest](#sdkopenstoragepolicycreaterequest))
    [SdkOpenStoragePolicyCreateResponse](#sdkopenstoragepolicycreateresponse)

Creates a storage policy
## Enumerate {#methodopenstorageapiopenstoragepolicyenumerate}

> **rpc** Enumerate([SdkOpenStoragePolicyEnumerateRequest](#sdkopenstoragepolicyenumeraterequest))
    [SdkOpenStoragePolicyEnumerateResponse](#sdkopenstoragepolicyenumerateresponse)

Enumerate returns a list of storage policies
## Inspect {#methodopenstorageapiopenstoragepolicyinspect}

> **rpc** Inspect([SdkOpenStoragePolicyInspectRequest](#sdkopenstoragepolicyinspectrequest))
    [SdkOpenStoragePolicyInspectResponse](#sdkopenstoragepolicyinspectresponse)

Inspect returns information about a specified storage policy
## Update {#methodopenstorageapiopenstoragepolicyupdate}

> **rpc** Update([SdkOpenStoragePolicyUpdateRequest](#sdkopenstoragepolicyupdaterequest))
    [SdkOpenStoragePolicyUpdateResponse](#sdkopenstoragepolicyupdateresponse)

Updates specified storage policy
## Delete {#methodopenstorageapiopenstoragepolicydelete}

> **rpc** Delete([SdkOpenStoragePolicyDeleteRequest](#sdkopenstoragepolicydeleterequest))
    [SdkOpenStoragePolicyDeleteResponse](#sdkopenstoragepolicydeleteresponse)

Deletes specified storage policy
## SetDefault {#methodopenstorageapiopenstoragepolicysetdefault}

> **rpc** SetDefault([SdkOpenStoragePolicySetDefaultRequest](#sdkopenstoragepolicysetdefaultrequest))
    [SdkOpenStoragePolicySetDefaultResponse](#sdkopenstoragepolicysetdefaultresponse)

Set specified storage policy as default policy
## DefaultInspect {#methodopenstorageapiopenstoragepolicydefaultinspect}

> **rpc** DefaultInspect([SdkOpenStoragePolicyDefaultInspectRequest](#sdkopenstoragepolicydefaultinspectrequest))
    [SdkOpenStoragePolicyDefaultInspectResponse](#sdkopenstoragepolicydefaultinspectresponse)

DefaultInspect return default storage policy if any, otherwise
empty response
## Release {#methodopenstorageapiopenstoragepolicyrelease}

> **rpc** Release([SdkOpenStoragePolicyReleaseRequest](#sdkopenstoragepolicyreleaserequest))
    [SdkOpenStoragePolicyReleaseResponse](#sdkopenstoragepolicyreleaseresponse)

Release specified storage policy constraint for volume
creation
 <!-- end methods -->

# OpenStoragePool {#serviceopenstorageapiopenstoragepool}
OpenStoragePool is a service used to manage storage pools in the cluster

## Resize {#methodopenstorageapiopenstoragepoolresize}

> **rpc** Resize([SdkStoragePoolResizeRequest](#sdkstoragepoolresizerequest))
    [SdkStoragePoolResizeResponse](#sdkstoragepoolresizeresponse)

Resize expands the specified storage pool based on the request parameters
## Rebalance {#methodopenstorageapiopenstoragepoolrebalance}

> **rpc** Rebalance([SdkStorageRebalanceRequest](#sdkstoragerebalancerequest))
    [SdkStorageRebalanceResponse](#sdkstoragerebalanceresponse)

Rebalance creates a new rebalance task
## UpdateRebalanceJobState {#methodopenstorageapiopenstoragepoolupdaterebalancejobstate}

> **rpc** UpdateRebalanceJobState([SdkUpdateRebalanceJobRequest](#sdkupdaterebalancejobrequest))
    [SdkUpdateRebalanceJobResponse](#sdkupdaterebalancejobresponse)

UpdateRebalanceJobState updates existing rebalance task state.
Only acceptable values are
StorageRebalanceJobState_PAUSED - acceptable only from running state
StorageRebalanceJobState_CANCELLED - acceptable only from running/pause state
StorageRebalanceJobState_RUNNING - acceptable only from pause state
## GetRebalanceJobStatus {#methodopenstorageapiopenstoragepoolgetrebalancejobstatus}

> **rpc** GetRebalanceJobStatus([SdkGetRebalanceJobStatusRequest](#sdkgetrebalancejobstatusrequest))
    [SdkGetRebalanceJobStatusResponse](#sdkgetrebalancejobstatusresponse)

GetRebalanceJobStatus returns rebalance status for specified job
## EnumerateRebalanceJobs {#methodopenstorageapiopenstoragepoolenumeraterebalancejobs}

> **rpc** EnumerateRebalanceJobs([SdkEnumerateRebalanceJobsRequest](#sdkenumeraterebalancejobsrequest))
    [SdkEnumerateRebalanceJobsResponse](#sdkenumeraterebalancejobsresponse)

EnumerateRebalanceJobs returns all rebalance jobs currently known to the system
 <!-- end methods -->

# OpenStorageRole {#serviceopenstorageapiopenstoragerole}
OpenStorageRole service provides methods to manage user roles

### Custom roles
The OpenStorage SDK server is equipped to handle customized authorization
roles. Using this model it allows administrators to customize the permission
rules of a role to be used by a user.

Creating a custom role, or an SdkRole, is done by setting up a set of allowed _rules_,
or SdkRules, directives which are sequentially scanned until a match is found. Rules
are created using the names of OpenStorage SDK services and APIs as follows:

The message SdkRules has the following properties:

* Services: Is the gRPC service name in `OpenStorage<service name>` in lowercase
* Apis: Is the API name in the service in lowercase

Please see SdkRule for more information on the format.

## Create {#methodopenstorageapiopenstoragerolecreate}

> **rpc** Create([SdkRoleCreateRequest](#sdkrolecreaterequest))
    [SdkRoleCreateResponse](#sdkrolecreateresponse)

Create a role for users in the system
## Enumerate {#methodopenstorageapiopenstorageroleenumerate}

> **rpc** Enumerate([SdkRoleEnumerateRequest](#sdkroleenumeraterequest))
    [SdkRoleEnumerateResponse](#sdkroleenumerateresponse)

List all roles
## Inspect {#methodopenstorageapiopenstorageroleinspect}

> **rpc** Inspect([SdkRoleInspectRequest](#sdkroleinspectrequest))
    [SdkRoleInspectResponse](#sdkroleinspectresponse)

Get information about a role
## Delete {#methodopenstorageapiopenstorageroledelete}

> **rpc** Delete([SdkRoleDeleteRequest](#sdkroledeleterequest))
    [SdkRoleDeleteResponse](#sdkroledeleteresponse)

Delete an existing role
## Update {#methodopenstorageapiopenstorageroleupdate}

> **rpc** Update([SdkRoleUpdateRequest](#sdkroleupdaterequest))
    [SdkRoleUpdateResponse](#sdkroleupdateresponse)

Update an existing role
 <!-- end methods -->

# OpenStorageSchedulePolicy {#serviceopenstorageapiopenstorageschedulepolicy}
OpenStorageSchedulePolicy service is used to manage the automated
snapshots for a volume

## Create {#methodopenstorageapiopenstorageschedulepolicycreate}

> **rpc** Create([SdkSchedulePolicyCreateRequest](#sdkschedulepolicycreaterequest))
    [SdkSchedulePolicyCreateResponse](#sdkschedulepolicycreateresponse)

Create creates a new snapshot schedule. They can be setup daily,
weekly, or monthly.
## Update {#methodopenstorageapiopenstorageschedulepolicyupdate}

> **rpc** Update([SdkSchedulePolicyUpdateRequest](#sdkschedulepolicyupdaterequest))
    [SdkSchedulePolicyUpdateResponse](#sdkschedulepolicyupdateresponse)

Update a snapshot schedule
## Enumerate {#methodopenstorageapiopenstorageschedulepolicyenumerate}

> **rpc** Enumerate([SdkSchedulePolicyEnumerateRequest](#sdkschedulepolicyenumeraterequest))
    [SdkSchedulePolicyEnumerateResponse](#sdkschedulepolicyenumerateresponse)

Enumerate returns a list of schedules
## Inspect {#methodopenstorageapiopenstorageschedulepolicyinspect}

> **rpc** Inspect([SdkSchedulePolicyInspectRequest](#sdkschedulepolicyinspectrequest))
    [SdkSchedulePolicyInspectResponse](#sdkschedulepolicyinspectresponse)

Inspect returns information about a specified schedule
## Delete {#methodopenstorageapiopenstorageschedulepolicydelete}

> **rpc** Delete([SdkSchedulePolicyDeleteRequest](#sdkschedulepolicydeleterequest))
    [SdkSchedulePolicyDeleteResponse](#sdkschedulepolicydeleteresponse)

Delete removes a snapshot schedule
 <!-- end methods -->

# OpenStorageVolume {#serviceopenstorageapiopenstoragevolume}
OpenStorageVolume is a service used to manage the volumes of a storage system

## Create {#methodopenstorageapiopenstoragevolumecreate}

> **rpc** Create([SdkVolumeCreateRequest](#sdkvolumecreaterequest))
    [SdkVolumeCreateResponse](#sdkvolumecreateresponse)

Create creates a volume according to the specification provided

Requires access AccessType.Read when cloning from a parent volume.

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
## Clone {#methodopenstorageapiopenstoragevolumeclone}

> **rpc** Clone([SdkVolumeCloneRequest](#sdkvolumeclonerequest))
    [SdkVolumeCloneResponse](#sdkvolumecloneresponse)

Clone creates a new writable volume cloned from an existing volume

Requires access AccessType.Read of volume
## Delete {#methodopenstorageapiopenstoragevolumedelete}

> **rpc** Delete([SdkVolumeDeleteRequest](#sdkvolumedeleterequest))
    [SdkVolumeDeleteResponse](#sdkvolumedeleteresponse)

Delete deletes the provided volume

Requires access AccessType.Admin of volume
## Inspect {#methodopenstorageapiopenstoragevolumeinspect}

> **rpc** Inspect([SdkVolumeInspectRequest](#sdkvolumeinspectrequest))
    [SdkVolumeInspectResponse](#sdkvolumeinspectresponse)

Inspect returns information about a volume

Requires access AccessType.Read of volume
## InspectWithFilters {#methodopenstorageapiopenstoragevolumeinspectwithfilters}

> **rpc** InspectWithFilters([SdkVolumeInspectWithFiltersRequest](#sdkvolumeinspectwithfiltersrequest))
    [SdkVolumeInspectWithFiltersResponse](#sdkvolumeinspectwithfiltersresponse)

Returns information for a list of volumes that match a filter.
This call is a helper function like calling
`OpenStorageVolume.EnumerateWithFilters` then having it
return the contents of each of those volumes
`OpenStorageVolume.Inspect()`. Take care in using this call
when requesting large number of volumes because it will
block until it has all the information requested before
returning.
## Update {#methodopenstorageapiopenstoragevolumeupdate}

> **rpc** Update([SdkVolumeUpdateRequest](#sdkvolumeupdaterequest))
    [SdkVolumeUpdateResponse](#sdkvolumeupdateresponse)

Update provides a method for manipulating the specification and attributes of a volume.
Set can be used to resize a volume, update labels, change replica count, and much more.

Requires access AccessType.Write of volume
## Stats {#methodopenstorageapiopenstoragevolumestats}

> **rpc** Stats([SdkVolumeStatsRequest](#sdkvolumestatsrequest))
    [SdkVolumeStatsResponse](#sdkvolumestatsresponse)

Stats returns the statistics for the requested volume

Requires access AccessType.Read of volume
## CapacityUsage {#methodopenstorageapiopenstoragevolumecapacityusage}

> **rpc** CapacityUsage([SdkVolumeCapacityUsageRequest](#sdkvolumecapacityusagerequest))
    [SdkVolumeCapacityUsageResponse](#sdkvolumecapacityusageresponse)

CapacityUsage returns volume/snapshot's capacity usage details

Requires access AccessType.Read of volume

##### Error codes:

* codes.Aborted : Command was aborted and only total_bytes field is valid
* code.Unimmplemented : Command is not suported this kernel.Only total_bytes
field is valid;
## Enumerate {#methodopenstorageapiopenstoragevolumeenumerate}

> **rpc** Enumerate([SdkVolumeEnumerateRequest](#sdkvolumeenumeraterequest))
    [SdkVolumeEnumerateResponse](#sdkvolumeenumerateresponse)

Enumerate returns a list of volume ids
## EnumerateWithFilters {#methodopenstorageapiopenstoragevolumeenumeratewithfilters}

> **rpc** EnumerateWithFilters([SdkVolumeEnumerateWithFiltersRequest](#sdkvolumeenumeratewithfiltersrequest))
    [SdkVolumeEnumerateWithFiltersResponse](#sdkvolumeenumeratewithfiltersresponse)

Enumerate returns a list of volume ids that match the labels if any are provided.
## SnapshotCreate {#methodopenstorageapiopenstoragevolumesnapshotcreate}

> **rpc** SnapshotCreate([SdkVolumeSnapshotCreateRequest](#sdkvolumesnapshotcreaterequest))
    [SdkVolumeSnapshotCreateResponse](#sdkvolumesnapshotcreateresponse)

SnapshotCreate creates a snapshot of a volume. This creates an immutable (read-only),
point-in-time snapshot of a volume. To create a new writable volume from
a snapshot, please use OpenStorageVolume.Clone().

Requires access AccessType.Read of volume
## SnapshotRestore {#methodopenstorageapiopenstoragevolumesnapshotrestore}

> **rpc** SnapshotRestore([SdkVolumeSnapshotRestoreRequest](#sdkvolumesnapshotrestorerequest))
    [SdkVolumeSnapshotRestoreResponse](#sdkvolumesnapshotrestoreresponse)

SnapshotRestore restores a volume to a specified snapshot

Requires access AccessType.Write of volume
## SnapshotEnumerate {#methodopenstorageapiopenstoragevolumesnapshotenumerate}

> **rpc** SnapshotEnumerate([SdkVolumeSnapshotEnumerateRequest](#sdkvolumesnapshotenumeraterequest))
    [SdkVolumeSnapshotEnumerateResponse](#sdkvolumesnapshotenumerateresponse)

SnapshotEnumerate returns a list of snapshots for a specific volume
## SnapshotEnumerateWithFilters {#methodopenstorageapiopenstoragevolumesnapshotenumeratewithfilters}

> **rpc** SnapshotEnumerateWithFilters([SdkVolumeSnapshotEnumerateWithFiltersRequest](#sdkvolumesnapshotenumeratewithfiltersrequest))
    [SdkVolumeSnapshotEnumerateWithFiltersResponse](#sdkvolumesnapshotenumeratewithfiltersresponse)

SnapshotEnumerate returns a list of snapshots.
To filter all the snapshots for a specific volume which may no longer exist,
specify a volume id.
Labels can also be used to filter the snapshot list.
If neither are provided all snapshots will be returned.
## SnapshotScheduleUpdate {#methodopenstorageapiopenstoragevolumesnapshotscheduleupdate}

> **rpc** SnapshotScheduleUpdate([SdkVolumeSnapshotScheduleUpdateRequest](#sdkvolumesnapshotscheduleupdaterequest))
    [SdkVolumeSnapshotScheduleUpdateResponse](#sdkvolumesnapshotscheduleupdateresponse)

Sets the snapshot schedules. This information is saved in the VolumeSpec.snapshot_schedule
as `policy=<name>,...`. This function will overwrite any policy values
in the volume. To delete the policies in the volume send no policies.

Requires access AccessType.Write of volume
## VolumeCatalog {#methodopenstorageapiopenstoragevolumevolumecatalog}

> **rpc** VolumeCatalog([SdkVolumeCatalogRequest](#sdkvolumecatalogrequest))
    [SdkVolumeCatalogResponse](#sdkvolumecatalogresponse)

Gets the volume catalog of an attached and mounted volume.
Returns the entire tree up to "n"  depth (default is all of it)
Takes a path that can be used as the new root for the catalog request.
 <!-- end methods -->
 <!-- end services -->

# Messages


## ActiveRequest {#activerequest}
Active Request


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


| Field | Type | Description |
| ----- | ---- | ----------- |
| RequestCount | [ int64](#int64) | none |
| ActiveRequest | [repeated ActiveRequest](#activerequest) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## Alert {#alert}
Alert is a structure that represents an alert object


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ int64](#int64) | Id for Alert |
| severity | [ SeverityType](#severitytype) | Severity of the Alert |
| alert_type | [ int64](#int64) | AlertType user defined alert type |
| message | [ string](#string) | Message describing the Alert |
| timestamp | [ google.protobuf.Timestamp](#googleprotobuftimestamp) | Timestamp when Alert occurred |
| resource_id | [ string](#string) | ResourceId where Alert occurred |
| resource | [ ResourceType](#resourcetype) | Resource where Alert occurred |
| cleared | [ bool](#bool) | Cleared Flag |
| ttl | [ uint64](#uint64) | Time-to-live in seconds for this Alert |
| unique_tag | [ string](#string) | UniqueTag helps identify a unique alert for a given resouce |
| count | [ int64](#int64) | Count of such alerts raised so far. |
| first_seen | [ google.protobuf.Timestamp](#googleprotobuftimestamp) | Timestamp when such alert was raised the very first time. |
 <!-- end Fields -->
 <!-- end HasFields -->


## Alerts {#alerts}
Alerts is an array of Alert objects


| Field | Type | Description |
| ----- | ---- | ----------- |
| alert | [repeated Alert](#alert) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## BucketAccessCredentials {#bucketaccesscredentials}
Defines the bucket access credential object


| Field | Type | Description |
| ----- | ---- | ----------- |
| access_key_id | [ string](#string) | Access key id |
| secret_access_key | [ string](#string) | Secret access key |
 <!-- end Fields -->
 <!-- end HasFields -->


## BucketCreateRequest {#bucketcreaterequest}
Defines a request to create a bucket.


| Field | Type | Description |
| ----- | ---- | ----------- |
| name | [ string](#string) | Unique name of the bucket. This will be used for idempotency. |
| region | [ string](#string) | Region in which bucket will be created. |
| endpoint | [ string](#string) | Endpoint to use when creating the bucket |
| anonymousBucketAccessMode | [ AnonymousBucketAccessMode](#anonymousbucketaccessmode) | Anonymous access policy for the bucket. |
 <!-- end Fields -->
 <!-- end HasFields -->


## BucketCreateResponse {#bucketcreateresponse}
Defines a response to the creation of a bucket


| Field | Type | Description |
| ----- | ---- | ----------- |
| bucket_id | [ string](#string) | Id of new bucket |
 <!-- end Fields -->
 <!-- end HasFields -->


## BucketDeleteRequest {#bucketdeleterequest}
Defines the request to delete a bucket


| Field | Type | Description |
| ----- | ---- | ----------- |
| bucket_id | [ string](#string) | Id of bucket to delete |
| region | [ string](#string) | Region in which bucket will be created. |
| endpoint | [ string](#string) | Endpoint to use when deleting the bucket |
| clear_bucket | [ bool](#bool) | Flag to allow non empty bucket deletion. |
 <!-- end Fields -->
 <!-- end HasFields -->


## BucketDeleteResponse {#bucketdeleteresponse}
Empty response

 <!-- end HasFields -->


## BucketGrantAccessRequest {#bucketgrantaccessrequest}
Defines a request to grant access to the bucket


| Field | Type | Description |
| ----- | ---- | ----------- |
| bucket_id | [ string](#string) | Id of the bucket |
| account_name | [ string](#string) | Name of the account to which access to be provided |
| access_policy | [ string](#string) | Access policy to be applied for the account |
 <!-- end Fields -->
 <!-- end HasFields -->


## BucketGrantAccessResponse {#bucketgrantaccessresponse}
Defines a response to the creation of a bucket


| Field | Type | Description |
| ----- | ---- | ----------- |
| account_id | [ string](#string) | This is the account_id that is being provided access. This will This will be required later to revoke access. |
| credentials | [ BucketAccessCredentials](#bucketaccesscredentials) | Credentials supplied for accessing the bucket ex: aws access key id and secret, etc. |
 <!-- end Fields -->
 <!-- end HasFields -->


## BucketRevokeAccessRequest {#bucketrevokeaccessrequest}
Defines the request to revoke access to the bucket


| Field | Type | Description |
| ----- | ---- | ----------- |
| bucket_id | [ string](#string) | Id of bucket to delete |
| account_id | [ string](#string) | AccountId that is having its access revoked. |
 <!-- end Fields -->
 <!-- end HasFields -->


## BucketRevokeAccessResponse {#bucketrevokeaccessresponse}
Empty response

 <!-- end HasFields -->


## CapacityUsageInfo {#capacityusageinfo}
Provides details on exclusive and shared storage used by
snapshot/volume specifically for copy-on-write(COW) snapshots. Deletion
of snapshots and overwrite of volume will affect the exclusive storage
used by the other dependent snaps and parent volume.


| Field | Type | Description |
| ----- | ---- | ----------- |
| exclusive_bytes | [ int64](#int64) | Storage consumed exclusively by this single snapshot. Deletion of this snapshot may increase the free storage available by this amount. |
| shared_bytes | [ int64](#int64) | Storage consumed by this snapshot that is shared with parent and children |
| total_bytes | [ int64](#int64) | TotalBytes used by this volume |
 <!-- end Fields -->
 <!-- end HasFields -->


## Catalog {#catalog}



| Field | Type | Description |
| ----- | ---- | ----------- |
| name | [ string](#string) | Name of the Directory/File |
| path | [ string](#string) | Full Path of the Directory/File |
| type | [ string](#string) | Type Directory or File |
| size | [ uint64](#uint64) | File or Directory Size |
| LastModified | [ google.protobuf.Timestamp](#googleprotobuftimestamp) | Last Modified |
| children | [repeated Catalog](#catalog) | Children |
 <!-- end Fields -->
 <!-- end HasFields -->


## CatalogResponse {#catalogresponse}



| Field | Type | Description |
| ----- | ---- | ----------- |
| root | [ Catalog](#catalog) | Root Catalog |
| report | [ Report](#report) | Report of total directories and files count |
 <!-- end Fields -->
 <!-- end HasFields -->


## CloudDriveTransferJob {#clouddrivetransferjob}



| Field | Type | Description |
| ----- | ---- | ----------- |
| source_driveset_id | [ string](#string) | SourceDrivesetID is the ID of the current driveset that needs to be transferred |
| destination_instance_id | [ string](#string) | DestinationInstanceID is the ID of the storageless instance that needs to take over the SourceDriveSetID |
| status | [ string](#string) | Status describes a helpful status of this operation |
 <!-- end Fields -->
 <!-- end HasFields -->


## CloudMigrate {#cloudmigrate}


 <!-- end HasFields -->


## CloudMigrateCancelRequest {#cloudmigratecancelrequest}
Request to stop a cloud migration


| Field | Type | Description |
| ----- | ---- | ----------- |
| task_id | [ string](#string) | The id of the task to cancel |
 <!-- end Fields -->
 <!-- end HasFields -->


## CloudMigrateInfo {#cloudmigrateinfo}



| Field | Type | Description |
| ----- | ---- | ----------- |
| task_id | [ string](#string) | Task id associated with this migration |
| cluster_id | [ string](#string) | ID of the cluster where the volume is being migrated |
| local_volume_id | [ string](#string) | ID of the volume on the local cluster |
| local_volume_name | [ string](#string) | Name of the volume on the local cluster |
| remote_volume_id | [ string](#string) | ID of the volume on the remote cluster |
| cloudbackup_id | [ string](#string) | ID of the cloudbackup used for the migration |
| current_stage | [ CloudMigrate.Stage](#cloudmigratestage) | Current stage of the volume migration |
| status | [ CloudMigrate.Status](#cloudmigratestatus) | Status of the current stage |
| last_update | [ google.protobuf.Timestamp](#googleprotobuftimestamp) | Last time the status was updated |
| error_reason | [ string](#string) | Contains the reason for the migration error |
| start_time | [ google.protobuf.Timestamp](#googleprotobuftimestamp) | StartTime indicates Op's start time |
| completed_time | [ google.protobuf.Timestamp](#googleprotobuftimestamp) | CompletedTime indicates Op's completed time |
| bytes_total | [ uint64](#uint64) | BytesTotal is the number of bytes being transferred |
| bytes_done | [ uint64](#uint64) | BytesDone is the number of bytes already transferred |
| eta_seconds | [ int64](#int64) | ETASeconds the time duration in seconds for cloud migration completion |
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
| task_id | [ string](#string) | (Optional) Unique TaskId assocaiated with this migration. If not provided one will be generated and returned in the response |
 <!-- end Fields -->
 <!-- end HasFields -->


## CloudMigrateStartResponse {#cloudmigratestartresponse}
Response to start a cloud migration


| Field | Type | Description |
| ----- | ---- | ----------- |
| task_id | [ string](#string) | TaskId assocaiated with the migration that was started |
 <!-- end Fields -->
 <!-- end HasFields -->


## CloudMigrateStatusRequest {#cloudmigratestatusrequest}
Request for cloud migration operation status


| Field | Type | Description |
| ----- | ---- | ----------- |
| task_id | [ string](#string) | Task id for which to return status |
| cluster_id | [ string](#string) | ID of the cluster for which to return migration statuses |
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


## ClusterPairCreateRequest {#clusterpaircreaterequest}
Used to send a request to create a cluster pair


| Field | Type | Description |
| ----- | ---- | ----------- |
| remote_cluster_ip | [ string](#string) | IP of the remote cluster |
| remote_cluster_port | [ uint32](#uint32) | Port for the remote cluster |
| remote_cluster_token | [ string](#string) | Token used to authenticate with the remote cluster |
| set_default | [ bool](#bool) | Set the new pair as the default |
| mode | [ ClusterPairMode.Mode](#clusterpairmodemode) | The mode to use for the cluster pair |
| credential_id | [ string](#string) | Use for the cluster pairing, if given credential id will be used in ClusterPairCreate service |
 <!-- end Fields -->
 <!-- end HasFields -->


## ClusterPairCreateResponse {#clusterpaircreateresponse}
Response for a pair request


| Field | Type | Description |
| ----- | ---- | ----------- |
| remote_cluster_id | [ string](#string) | ID of the remote cluster |
| remote_cluster_name | [ string](#string) | Name of the remote cluster |
 <!-- end Fields -->
 <!-- end HasFields -->


## ClusterPairGetResponse {#clusterpairgetresponse}
Response to get a cluster pair


| Field | Type | Description |
| ----- | ---- | ----------- |
| pair_info | [ ClusterPairInfo](#clusterpairinfo) | Info about the cluster pair |
 <!-- end Fields -->
 <!-- end HasFields -->


## ClusterPairInfo {#clusterpairinfo}
Information about a cluster pair


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ string](#string) | ID of the cluster |
| name | [ string](#string) | Name of the cluster |
| endpoint | [ string](#string) | The endpoint used for creating the pair |
| current_endpoints | [repeated string](#string) | Current endpoints of the cluster |
| secure | [ bool](#bool) | Flag used to determine if communication is over a secure channel |
| token | [ string](#string) | Token associated with cluster |
| options | [map ClusterPairInfo.OptionsEntry](#clusterpairinfooptionsentry) | Key/value pair of options associated with the cluster Opaque to openstorage and interpreted by the drivers |
| mode | [ ClusterPairMode.Mode](#clusterpairmodemode) | Mode for the cluster pair |
 <!-- end Fields -->
 <!-- end HasFields -->


## ClusterPairInfo.OptionsEntry {#clusterpairinfooptionsentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## ClusterPairMode {#clusterpairmode}


 <!-- end HasFields -->


## ClusterPairProcessRequest {#clusterpairprocessrequest}
Used to process a pair request from a remote cluster


| Field | Type | Description |
| ----- | ---- | ----------- |
| source_cluster_id | [ string](#string) | ID of the cluster requesting the pairing |
| remote_cluster_token | [ string](#string) | Token used to authenticate with the remote cluster |
| mode | [ ClusterPairMode.Mode](#clusterpairmodemode) | The mode to use for the cluster pair |
| credential_id | [ string](#string) | Use for the cluster pairing, if given credential id will be used in ClusterPairCreate service |
 <!-- end Fields -->
 <!-- end HasFields -->


## ClusterPairProcessResponse {#clusterpairprocessresponse}
Response after a pairing has been processed


| Field | Type | Description |
| ----- | ---- | ----------- |
| remote_cluster_id | [ string](#string) | ID of the cluster which processed the pair request |
| remote_cluster_name | [ string](#string) | Name of the cluster which processed the pair request |
| remote_cluster_endpoints | [repeated string](#string) | List of endpoints that can be used to communicate with the cluster |
| options | [map ClusterPairProcessResponse.OptionsEntry](#clusterpairprocessresponseoptionsentry) | Key/value pair of options returned on successful pairing. Opaque to openstorage and interpreted by the drivers |
 <!-- end Fields -->
 <!-- end HasFields -->


## ClusterPairProcessResponse.OptionsEntry {#clusterpairprocessresponseoptionsentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## ClusterPairTokenGetResponse {#clusterpairtokengetresponse}
Response to get the cluster token


| Field | Type | Description |
| ----- | ---- | ----------- |
| token | [ string](#string) | Token used to authenticate clusters |
 <!-- end Fields -->
 <!-- end HasFields -->


## ClusterPairsEnumerateResponse {#clusterpairsenumerateresponse}
Response to enumerate all the cluster pairs


| Field | Type | Description |
| ----- | ---- | ----------- |
| default_id | [ string](#string) | ID of the default cluster pair |
| pairs | [map ClusterPairsEnumerateResponse.PairsEntry](#clusterpairsenumerateresponsepairsentry) | Pairs Info about the cluster pairs |
 <!-- end Fields -->
 <!-- end HasFields -->


## ClusterPairsEnumerateResponse.PairsEntry {#clusterpairsenumerateresponsepairsentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ ClusterPairInfo](#clusterpairinfo) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## ClusterResponse {#clusterresponse}
ClusterResponse specifies a response that gets returned when requesting the cluster


| Field | Type | Description |
| ----- | ---- | ----------- |
| error | [ string](#string) | Error code

in: body |
 <!-- end Fields -->
 <!-- end HasFields -->


## CollectDiagsJob {#collectdiagsjob}



| Field | Type | Description |
| ----- | ---- | ----------- |
| request | [ SdkDiagsCollectRequest](#sdkdiagscollectrequest) | Request is the user request for this diags collection job |
| statuses | [repeated DiagsCollectionStatus](#diagscollectionstatus) | Statuses is a list of statuses for diags collection for each node that is part of the request |
 <!-- end Fields -->
 <!-- end HasFields -->


## DiagsCollectionStatus {#diagscollectionstatus}



| Field | Type | Description |
| ----- | ---- | ----------- |
| node | [ string](#string) | Node is the node that's collecting the diags |
| state | [ DiagsCollectionStatus.State](#diagscollectionstatusstate) | State is the current state of diags collection on the node |
| message | [ string](#string) | Message is a user friendly message for current status of diags collection |
 <!-- end Fields -->
 <!-- end HasFields -->


## DiagsNodeSelector {#diagsnodeselector}
DiagsNodeSelector allows selecting nodes for diags collection
User can select NodeLabelSelector AND/OR NodeIDs. If both are provided, the implementation will select nodes based on
both labels and IDs and also handle overlaps
If All is set to true, other selectors are ignored since it selects all nodes


| Field | Type | Description |
| ----- | ---- | ----------- |
| node_label_selector | [repeated LabelSelectorRequirement](#labelselectorrequirement) | NodeLabelSelector is a label selector used to select the nodes for which diags will be collected |
| node_ids | [repeated string](#string) | NodeIDs are unique IDs fo the nodes for which the diags will be collected |
| all | [ bool](#bool) | All selects all nodes for diags collection |
 <!-- end Fields -->
 <!-- end HasFields -->


## DiagsVolumeSelector {#diagsvolumeselector}
DiagsVolumeSelector allows selecting volumes for diags collection
User can select VolumeLabelSelector AND/OR VolumeIDs. If both are provided, the implementation will select nodes
based on both labels and IDs and also handle overlaps


| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_label_selector | [repeated LabelSelectorRequirement](#labelselectorrequirement) | VolumeLabelSelector selects volumes by their labels and then uses replica and attached nodes for those volumes for diags collection |
| volume_ids | [repeated string](#string) | VolumeIDs selects volumes by their unique IDs and then uses replica and attached nodes for those volumes for diags collection |
 <!-- end Fields -->
 <!-- end HasFields -->


## DrainAttachmentsSummary {#drainattachmentssummary}
DrainAttachments summary of the volumes whose attachments need to be drained
from a node


| Field | Type | Description |
| ----- | ---- | ----------- |
| num_volumes_total | [ uint64](#uint64) | Total number of volumes that need to be drained |
| num_volumes_done | [ uint64](#uint64) | Number of volumes which have been drained |
| num_volumes_pending | [ uint64](#uint64) | Number of volumes which have not been drained yet |
 <!-- end Fields -->
 <!-- end HasFields -->


## ExportSpec {#exportspec}
ExportSpec defines how the volume is exported..


| Field | Type | Description |
| ----- | ---- | ----------- |
| export_protocol | [ ExportProtocol](#exportprotocol) | ExportProtocol defines how the volume is exported. |
| export_options | [ string](#string) | ExportOptions options exporting the volume. |
 <!-- end Fields -->
 <!-- end HasFields -->


## FastpathConfig {#fastpathconfig}
FastpathConfig part of volume


| Field | Type | Description |
| ----- | ---- | ----------- |
| setup_on | [ int32](#int32) | fastpath setup on this node |
| promote | [ bool](#bool) | Fastpath temporary promotion during attach |
| status | [ FastpathStatus](#fastpathstatus) | Fastpath consolidated current status across replicas |
| replicas | [repeated FastpathReplState](#fastpathreplstate) | Fastpath replica state for each replica in replica set |
| dirty | [ bool](#bool) | Dirty flag on volume - was attached in userspace |
| coord_uuid | [ string](#string) | fastpath coordinator node uuid to enhance reporting |
| force_failover | [ bool](#bool) | fastpath force failover, disable auto promote to fastpath |
 <!-- end Fields -->
 <!-- end HasFields -->


## FastpathReplState {#fastpathreplstate}



| Field | Type | Description |
| ----- | ---- | ----------- |
| dev_id | [ uint64](#uint64) | none |
| node_id | [ uint32](#uint32) | none |
| protocol | [ FastpathProtocol](#fastpathprotocol) | none |
| acl | [ bool](#bool) | none |
| exported_device | [ string](#string) | target info |
| block | [ bool](#bool) | none |
| target | [ string](#string) | none |
| exported | [ bool](#bool) | none |
| imported | [ bool](#bool) | initiator info |
| devpath | [ string](#string) | none |
| node_uuid | [ string](#string) | node_uuid added to enhance UI reporting |
 <!-- end Fields -->
 <!-- end HasFields -->


## FilesystemCheck {#filesystemcheck}


 <!-- end HasFields -->


## FilesystemTrim {#filesystemtrim}


 <!-- end HasFields -->


## FstrimVolumeUsageInfo {#fstrimvolumeusageinfo}
FstrimVolUsageInfo


| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_name | [ string](#string) | Volume name |
| volume_size | [ uint64](#uint64) | Volume size |
| du_usage | [ uint64](#uint64) | Disk usage in bytes |
| px_usage | [ uint64](#uint64) | Disk usage seen in Portworx in bytes |
| perform_auto_fstrim | [ string](#string) | If auto fstrim is performed to the volume, if not, why |
 <!-- end Fields -->
 <!-- end HasFields -->


## GraphDriverChanges {#graphdriverchanges}
GraphDriverChanges represent a list of changes between the filesystem layers
specified by the ID and Parent.  // Parent may be an empty string, in which
case there is no parent.
Where the Path is the filesystem path within the layered filesystem


| Field | Type | Description |
| ----- | ---- | ----------- |
| path | [ string](#string) | none |
| kind | [ GraphDriverChangeType](#graphdriverchangetype) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## Group {#group}
Group represents VolumeGroup / namespace
All volumes in the same group share this object.


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ string](#string) | Id common identifier across volumes that have the same group. |
 <!-- end Fields -->
 <!-- end HasFields -->


## GroupSnapCreateRequest {#groupsnapcreaterequest}
GroupSnapCreateRequest specifies a request to create a snapshot of given group.


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ string](#string) | none |
| Labels | [map GroupSnapCreateRequest.LabelsEntry](#groupsnapcreaterequestlabelsentry) | none |
| volume_ids | [repeated string](#string) | none |
| delete_on_failure | [ bool](#bool) | none |
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


## IoStrategy {#iostrategy}
IoStrategy defines how I/O should be performed to backing storage media.


| Field | Type | Description |
| ----- | ---- | ----------- |
| async_io | [ bool](#bool) | AsyncIO enables kaio. |
| early_ack | [ bool](#bool) | EarlyAck enables acks for async I/O at the source. |
| direct_io | [ bool](#bool) | Enable direct I/O on the backing datastore |
 <!-- end Fields -->
 <!-- end HasFields -->


## IoThrottle {#iothrottle}
IoThrottle defines IO throttle limits for a volume
read_iops : maximum read iops this volume is allowed
write_iops : maximum write iops this volume is allowed
read_bw_mbytes  : maximum read bandwidth this volume is allowed in MegaBytes
write_bw_mbytes : maximum write bandwidth this volume is allowed in MegaBytes


| Field | Type | Description |
| ----- | ---- | ----------- |
| read_iops | [ uint32](#uint32) | none |
| write_iops | [ uint32](#uint32) | none |
| read_bw_mbytes | [ uint32](#uint32) | none |
| write_bw_mbytes | [ uint32](#uint32) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## Job {#job}
Job is a generic job object that can encapsulate other
messages which follow the job framework of APIs


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ string](#string) | ID of the job |
| state | [ Job.State](#jobstate) | State of the current job |
| type | [ Job.Type](#jobtype) | Type is the job type |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) job.drain_attachments | [ NodeDrainAttachmentsJob](#nodedrainattachmentsjob) | NodeDrainAttachmentsJob if selected this job desribes the task for removing volume attachments from a node |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) job.clouddrive_transfer | [ CloudDriveTransferJob](#clouddrivetransferjob) | CloudDriveTransferJob if selected describes the task to transfer a cloud driveset from one node to another |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) job.collect_diags | [ CollectDiagsJob](#collectdiagsjob) | CollectDiagsJob if selected describes the task to collect diagnostics for the cluster |
| create_time | [ google.protobuf.Timestamp](#googleprotobuftimestamp) | CreateTime is the time the job was created |
| last_update_time | [ google.protobuf.Timestamp](#googleprotobuftimestamp) | LastUpdateTime is the time the job was updated |
 <!-- end Fields -->
 <!-- end HasFields -->


## JobAudit {#jobaudit}
JobAudit is an audit entry for a job describing the different operations
performed as a part of the job


| Field | Type | Description |
| ----- | ---- | ----------- |
| summary | [repeated JobWorkSummary](#jobworksummary) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## JobSummary {#jobsummary}
JobSummary provides a summary of a job


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ string](#string) | ID of the job |
| total_runtime_seconds | [ uint64](#uint64) | Total runtime in seconds |
| work_summaries | [repeated JobWorkSummary](#jobworksummary) | Summary provides more information about the on-going job |
 <!-- end Fields -->
 <!-- end HasFields -->


## JobWorkSummary {#jobworksummary}
JobWorkSummary describes an action taken while performing the hob


| Field | Type | Description |
| ----- | ---- | ----------- |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) summary.drain_attachments_summary | [ DrainAttachmentsSummary](#drainattachmentssummary) | Summary summarizes drain attachment job |
 <!-- end Fields -->
 <!-- end HasFields -->


## LabelSelectorRequirement {#labelselectorrequirement}
LabelSelectorRequirement is a selector that contains values, a key, and an operator that
relates the key and values.


| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | Key is the label key that the selector applies to. |
| operator | [ LabelSelectorRequirement.Operator](#labelselectorrequirementoperator) | Operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist. |
| values | [repeated string](#string) | Values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch. |
 <!-- end Fields -->
 <!-- end HasFields -->


## LocateResponse {#locateresponse}
Locate response would be used to return a set of mounts
and/or Container IDs and their mount paths


| Field | Type | Description |
| ----- | ---- | ----------- |
| mounts | [map LocateResponse.MountsEntry](#locateresponsemountsentry) | Map of mounts <host>: /var/lib/osd/<volumemount> |
| dockerids | [map LocateResponse.DockeridsEntry](#locateresponsedockeridsentry) | Map of docker id's and their mounts <containerid>: /var/www |
 <!-- end Fields -->
 <!-- end HasFields -->


## LocateResponse.DockeridsEntry {#locateresponsedockeridsentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## LocateResponse.MountsEntry {#locateresponsemountsentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## MountOptions {#mountoptions}
MountOptions defines the mount options with which a volume is mounted.


| Field | Type | Description |
| ----- | ---- | ----------- |
| options | [map MountOptions.OptionsEntry](#mountoptionsoptionsentry) | Options are opaque key value pairs that are passed as mount options when a volume is mounted. If an empty value is provided only the key will be passed as an option If both key and value are provided then 'key=value' will be passed as an option |
 <!-- end Fields -->
 <!-- end HasFields -->


## MountOptions.OptionsEntry {#mountoptionsoptionsentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## NFSProxySpec {#nfsproxyspec}
NFSProxySpec is the spec for proxying an NFS share.


| Field | Type | Description |
| ----- | ---- | ----------- |
| export_path | [ string](#string) | ExportPath is the NFS export path on the NFS server |
| sub_path | [ string](#string) | SubPath is the sub-directory from an NFS share that should be reflected. |
 <!-- end Fields -->
 <!-- end HasFields -->


## NodeDrainAttachmentOptions {#nodedrainattachmentoptions}
Options for draining volume attachment from a node

 <!-- end HasFields -->


## NodeDrainAttachmentsJob {#nodedrainattachmentsjob}
NodeDrainAttachmentsJob describe a job to drain volume attachments from a node


| Field | Type | Description |
| ----- | ---- | ----------- |
| node_id | [ string](#string) | NodeID of the node for which this drain job is running |
| Status | [ string](#string) | Status describes a helpful status of this node drain operation |
| issuer | [ string](#string) | Issuer is a user friendly name for the caller who is invoking the API. It can be used by caller to filter out drain requests from a particular issuer |
| parameters | [ SdkNodeDrainAttachmentsRequest](#sdknodedrainattachmentsrequest) | Parameters is the original request params for this node drain operation This node drain job is applicable to only one of these node drain operations. |
| create_time | [ google.protobuf.Timestamp](#googleprotobuftimestamp) | CreateTime is the time the job was created |
| last_update_time | [ google.protobuf.Timestamp](#googleprotobuftimestamp) | LastUpdateTime is the time the job was updated |
 <!-- end Fields -->
 <!-- end HasFields -->


## ObjectstoreInfo {#objectstoreinfo}
ObjectstoreInfo is a structure that has current objectstore info


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


## Ownership {#ownership}
Ownership information for resource.
Administrators are users who belong to the group `*`, meaning, every group.


| Field | Type | Description |
| ----- | ---- | ----------- |
| owner | [ string](#string) | Username of owner.

The storage system uses the username taken from the security authorization token and is saved on this field. Only users with system administration can edit this value. |
| acls | [ Ownership.AccessControl](#ownershipaccesscontrol) | Permissions to share resource which can be set by the owner.

NOTE: To create an "admin" user which has access to any resource set the group value in the token of the user to `*`. |
 <!-- end Fields -->
 <!-- end HasFields -->


## Ownership.AccessControl {#ownershipaccesscontrol}



| Field | Type | Description |
| ----- | ---- | ----------- |
| groups | [map Ownership.AccessControl.GroupsEntry](#ownershipaccesscontrolgroupsentry) | Group access to resource which must match the group set in the authorization token. Can be set by the owner or the system administrator only. Possible values are: 1. no groups: Means no groups are given access. 2. `["*"]`: All groups are allowed. 3. `["group1", "group2"]`: Only certain groups are allowed. In this example only _group1_ and _group2_ are allowed. |
| collaborators | [map Ownership.AccessControl.CollaboratorsEntry](#ownershipaccesscontrolcollaboratorsentry) | Collaborator access to resource gives access to other user. Must be the username (unique id) set in the authorization token. The owner or the administrator can set this value. Possible values are: 1. no collaborators: Means no users are given access. 2. `["*"]`: All users are allowed. 3. `["username1", "username2"]`: Only certain usernames are allowed. In this example only _username1_ and _username2_ are allowed. |
| public | [ Ownership.PublicAccessControl](#ownershippublicaccesscontrol) | Public access to resource may be assigned for access by the public userd |
 <!-- end Fields -->
 <!-- end HasFields -->


## Ownership.AccessControl.CollaboratorsEntry {#ownershipaccesscontrolcollaboratorsentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ Ownership.AccessType](#ownershipaccesstype) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## Ownership.AccessControl.GroupsEntry {#ownershipaccesscontrolgroupsentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ Ownership.AccessType](#ownershipaccesstype) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## Ownership.PublicAccessControl {#ownershippublicaccesscontrol}
PublicAccessControl allows assigning public ownership


| Field | Type | Description |
| ----- | ---- | ----------- |
| type | [ Ownership.AccessType](#ownershipaccesstype) | AccessType declares which level of public access is allowed |
 <!-- end Fields -->
 <!-- end HasFields -->


## PXDProxySpec {#pxdproxyspec}
PXDProxySpec is the spec for proxying a Portworx volume


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ string](#string) | ID of the remote portworx volume |
 <!-- end Fields -->
 <!-- end HasFields -->


## ProxySpec {#proxyspec}
ProxySpec defines how this volume will reflect an external data source.


| Field | Type | Description |
| ----- | ---- | ----------- |
| proxy_protocol | [ ProxyProtocol](#proxyprotocol) | ProxyProtocol defines the protocol used for proxy. |
| endpoint | [ string](#string) | Endpoint is the external endpoint which can be used for accessing the external data source. |
| nfs_spec | [ NFSProxySpec](#nfsproxyspec) | NFSProxySpec is the spec for proxying an NFS share |
| s3_spec | [ S3ProxySpec](#s3proxyspec) | S3ProxySpec is the spec for proxying an external object store |
| pxd_spec | [ PXDProxySpec](#pxdproxyspec) | PXDProxySpec is the spec for proxying a Portworx volume |
| pure_block_spec | [ PureBlockSpec](#pureblockspec) | PureFileSpec is the spec for proxying a volume on pure_file backends |
| pure_file_spec | [ PureFileSpec](#purefilespec) | PureFileSpec is the spec for proxying a volume on pure_file backends |
 <!-- end Fields -->
 <!-- end HasFields -->


## PureBlockSpec {#pureblockspec}
PureBlockSpec is the spec for proxying a volume on pure_block backends


| Field | Type | Description |
| ----- | ---- | ----------- |
| serial_num | [ string](#string) | none |
| full_vol_name | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## PureFileSpec {#purefilespec}
PureFileSpec is the spec for proxying a volume on pure_file backends


| Field | Type | Description |
| ----- | ---- | ----------- |
| export_rules | [ string](#string) | none |
| full_vol_name | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## RelaxedReclaimPurge {#relaxedreclaimpurge}
Purges the RelaxedReclaim queue


| Field | Type | Description |
| ----- | ---- | ----------- |
| num_purged | [ uint64](#uint64) | num_purged returns number of volumes purged |
 <!-- end Fields -->
 <!-- end HasFields -->


## ReplicaPlacementSpec {#replicaplacementspec}



| Field | Type | Description |
| ----- | ---- | ----------- |
| weight | [ int64](#int64) | Weight defines the weight of the rule which allows to break the tie with other matching rules. A rule with higher weight wins over a rule with lower weight. (optional) |
| enforcement | [ EnforcementType](#enforcementtype) | Enforcement specifies the rule enforcement policy. Can take values: required or preferred. (optional) |
| affected_replicas | [ int32](#int32) | AffectedReplicas defines the number of volume replicas affected by this rule. If not provided, rule would affect all replicas (optional) |
| topology_key | [ string](#string) | TopologyKey key for the matching all segments of the cluster topology with the same key e.g If the key is failure-domain.beta.kubernetes.io/zone, this should match all nodes with the same value for this key (i.e in the same zone) |
| match_expressions | [repeated LabelSelectorRequirement](#labelselectorrequirement) | MatchExpressions is a list of label selector requirements. The requirements are ANDed. |
 <!-- end Fields -->
 <!-- end HasFields -->


## ReplicaSet {#replicaset}
ReplicaSet set of machine IDs (nodes) to which part of this volume is erasure
coded - for clustered storage arrays


| Field | Type | Description |
| ----- | ---- | ----------- |
| nodes | [repeated string](#string) | none |
| pool_uuids | [repeated string](#string) | Unique IDs of the storage pools for this replica set |
| id | [ uint32](#uint32) | ID is the unique ID of this replica set |
 <!-- end Fields -->
 <!-- end HasFields -->


## Report {#report}



| Field | Type | Description |
| ----- | ---- | ----------- |
| directories | [ int64](#int64) | Directory count |
| files | [ int64](#int64) | File count |
 <!-- end Fields -->
 <!-- end HasFields -->


## RestoreVolSnashotSchedule {#restorevolsnashotschedule}



| Field | Type | Description |
| ----- | ---- | ----------- |
| schedule | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## RestoreVolStoragePolicy {#restorevolstoragepolicy}



| Field | Type | Description |
| ----- | ---- | ----------- |
| policy | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## RestoreVolumeSpec {#restorevolumespec}
RestoreSpec allows some of the restore volume properties of to be modified
while restoring the cloud baackup. All pointer fields with nil value will
inherit corresponding field value from backup's spec.


| Field | Type | Description |
| ----- | ---- | ----------- |
| ha_level | [ int64](#int64) | HaLevel specifies the number of copies of data. |
| cos | [ CosType](#costype) | Cos specifies the relative class of service. |
| io_profile | [ IoProfile](#ioprofile) | IoProfile provides a hint about application using this volume. This field is ignored if IoProfileBkupSrc is set true |
| snapshot_interval | [ uint32](#uint32) | SnapshotInterval in minutes, set to 0 to disable snapshots |
| shared | [ RestoreParamBoolType](#restoreparambooltype) | Shared is true if this volume can be concurrently accessed by multiple users. |
| replica_set | [ ReplicaSet](#replicaset) | ReplicaSet is the desired set of nodes for the volume data. |
| aggregation_level | [ uint32](#uint32) | Aggregation level Specifies the number of parts the volume can be aggregated from. |
| snapshot_schedule | [ RestoreVolSnashotSchedule](#restorevolsnashotschedule) | SnapshotSchedule a well known string that specifies when snapshots should be taken. |
| sticky | [ RestoreParamBoolType](#restoreparambooltype) | Sticky volumes cannot be deleted until the flag is removed. |
| group | [ Group](#group) | Group identifies a consistency group |
| group_enforced | [ bool](#bool) | GroupEnforced is true if consistency group creation is enforced. |
| journal | [ RestoreParamBoolType](#restoreparambooltype) | Journal is true if data for the volume goes into the journal. |
| sharedv4 | [ RestoreParamBoolType](#restoreparambooltype) | Sharedv4 is true if this volume can be accessed via sharedv4. |
| queue_depth | [ uint32](#uint32) | QueueDepth defines the desired block device queue depth |
| nodiscard | [ RestoreParamBoolType](#restoreparambooltype) | Nodiscard specifies if the volume will be mounted with discard support disabled. i.e. FS will not release allocated blocks back to the backing storage pool. |
| io_strategy | [ IoStrategy](#iostrategy) | IoStrategy preferred strategy for I/O. |
| placement_strategy | [ VolumePlacementStrategy](#volumeplacementstrategy) | PlacementStrategy specifies a spec to indicate where to place the volume. |
| storage_policy | [ RestoreVolStoragePolicy](#restorevolstoragepolicy) | StoragePolicy if applied/specified while creating volume |
| ownership | [ Ownership](#ownership) | Ownership |
| export_spec | [ ExportSpec](#exportspec) | ExportSpec defines how the volume should be exported. |
| fp_preference | [ RestoreParamBoolType](#restoreparambooltype) | fastpath extensions |
| mount_options | [ MountOptions](#mountoptions) | MountOptions defines the options that should be used while mounting this volume |
| sharedv4_mount_options | [ MountOptions](#mountoptions) | Sharedv4MountOptions defines the options that will be used while mounting a sharedv4 volume from a node where the volume replica does not exist |
| proxy_write | [ RestoreParamBoolType](#restoreparambooltype) | Proxy_write is true if proxy write replication is enabled for the volume |
| io_profile_bkup_src | [ bool](#bool) | IoProfileBkupSrc indicates to inherit IoProfile from cloudbackup |
| proxy_spec | [ ProxySpec](#proxyspec) | ProxySpec indicates that this volume is used for proxying an external data source |
| sharedv4_service_spec | [ Sharedv4ServiceSpec](#sharedv4servicespec) | Sharedv4ServiceSpec specifies a spec for configuring a service for a sharedv4 volume |
| sharedv4_spec | [ Sharedv4Spec](#sharedv4spec) | Sharedv4Spec specifies common properties of sharedv4 and sharedv4 service volumes |
| auto_fstrim | [ RestoreParamBoolType](#restoreparambooltype) | Autofstrim is true if automatic fstrim is enabled for the volume |
| io_throttle | [ IoThrottle](#iothrottle) | IoThrottle specifies maximum io(iops/bandwidth) this volume is restricted to |
| readahead | [ RestoreParamBoolType](#restoreparambooltype) | Enable readahead for the volume |
 <!-- end Fields -->
 <!-- end HasFields -->


## RuntimeStateMap {#runtimestatemap}
RuntimeStateMap is a list of name value mapping of driver specific runtime
information.


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


## S3ProxySpec {#s3proxyspec}
S3ProxySpec is the spec for proxying an external object store.


| Field | Type | Description |
| ----- | ---- | ----------- |
| bucket_name | [ string](#string) | BucketName is the name of the bucket from the object store |
 <!-- end Fields -->
 <!-- end HasFields -->


## ScanPolicy {#scanpolicy}
ScanPolicy defines when a filesystem check is triggered and what action to take
User can specify *one* of the following valid policies
1. trigger=SCAN_TRIGGER_ON_MOUNT, action=SCAN_ACTION_SCAN_ONLY
2. trigger=SCAN_TRIGGER_ON_MOUNT, action=SCAN_ACTION_SCAN_REPAIR
3. trigger=SCAN_TRIGGER_ON_NEXT_MOUNT, action=SCAN_ACTION_SCAN_ONLY
4. trigger=SCAN_TRIGGER_ON_NEXT_MOUNT, action=SCAN_ACTION_SCAN_REPAIR
5. trigger=SCAN_TRIGGER_NONE, action=SCAN_ACTION_NONE
Note: When trigger == SCAN_TRIGGER_ON_NEXT_MOUNT, the associated action is
performed once on the next mount and the scanpolicy gets reset to
`trigger=SCAN_TRIGGER_NONE, action=SCAN_ACTION_NONE`, irrespective of the
output of the action.


| Field | Type | Description |
| ----- | ---- | ----------- |
| trigger | [ ScanPolicy.ScanTrigger](#scanpolicyscantrigger) | none |
| action | [ ScanPolicy.ScanAction](#scanpolicyscanaction) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## SchedulerTopology {#schedulertopology}
SchedulerTopology defines the topology information of the storage node
in scheduler context


| Field | Type | Description |
| ----- | ---- | ----------- |
| labels | [map SchedulerTopology.LabelsEntry](#schedulertopologylabelsentry) | Key-value pairs defining the topology of the node |
 <!-- end Fields -->
 <!-- end HasFields -->


## SchedulerTopology.LabelsEntry {#schedulertopologylabelsentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkAlertsAlertTypeQuery {#sdkalertsalerttypequery}
SdkAlertsAlertTypeQuery queries for alerts using alert type
and it requires that resource type be provided as well.


| Field | Type | Description |
| ----- | ---- | ----------- |
| resource_type | [ ResourceType](#resourcetype) | Resource type used to build query. |
| alert_type | [ int64](#int64) | Alert type used to build query. |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkAlertsCountSpan {#sdkalertscountspan}
SdkAlertsCountSpan to store count range information.


| Field | Type | Description |
| ----- | ---- | ----------- |
| min_count | [ int64](#int64) | Min count of such alerts raised so far. |
| max_count | [ int64](#int64) | Max count of such alerts raised so far. |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkAlertsDeleteRequest {#sdkalertsdeleterequest}
SdkAlertsDeleteRequest is a request message to delete alerts.


| Field | Type | Description |
| ----- | ---- | ----------- |
| queries | [repeated SdkAlertsQuery](#sdkalertsquery) | It takes a list of queries to find matching alerts. Matching alerts are deleted. |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkAlertsDeleteResponse {#sdkalertsdeleteresponse}
SdkAlertsDeleteResponse is empty.

 <!-- end HasFields -->


## SdkAlertsEnumerateWithFiltersRequest {#sdkalertsenumeratewithfiltersrequest}
SdkAlertsEnumerateRequest is a request message to enumerate alerts.


| Field | Type | Description |
| ----- | ---- | ----------- |
| queries | [repeated SdkAlertsQuery](#sdkalertsquery) | It is a list of queries to find matching alerts. Output of each of these queries is added to a global pool and returned as output of an RPC call. In that sense alerts are fetched if they match any of the queries. |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkAlertsEnumerateWithFiltersResponse {#sdkalertsenumeratewithfiltersresponse}
SdkAlertsEnumerateResponse is a list of alerts.


| Field | Type | Description |
| ----- | ---- | ----------- |
| alerts | [repeated Alert](#alert) | Response contains a list of alerts. |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkAlertsOption {#sdkalertsoption}
SdkAlertsOption contains options for filtering alerts.


| Field | Type | Description |
| ----- | ---- | ----------- |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) opt.min_severity_type | [ SeverityType](#severitytype) | Query using minimum severity type. |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) opt.is_cleared | [ bool](#bool) | Query using cleared flag. |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) opt.time_span | [ SdkAlertsTimeSpan](#sdkalertstimespan) | Query using a time span during which alert was last seen. |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) opt.count_span | [ SdkAlertsCountSpan](#sdkalertscountspan) | Query using a count span in which alert count exists. |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkAlertsQuery {#sdkalertsquery}
SdkAlertsQuery is one of the query types and a list of options.
Each query object is one of the three query types and a list of
options.


| Field | Type | Description |
| ----- | ---- | ----------- |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) query.resource_type_query | [ SdkAlertsResourceTypeQuery](#sdkalertsresourcetypequery) | Query only using resource type. |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) query.alert_type_query | [ SdkAlertsAlertTypeQuery](#sdkalertsalerttypequery) | Query using alert type and resource type. |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) query.resource_id_query | [ SdkAlertsResourceIdQuery](#sdkalertsresourceidquery) | Query using resource id, alert type and resource type. |
| opts | [repeated SdkAlertsOption](#sdkalertsoption) | Opts is a list of options associated with one of the queries. |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkAlertsResourceIdQuery {#sdkalertsresourceidquery}
SdkAlertsResourceIdQuery queries for alerts using resource id
and it requires that both alert type and resource type be provided
as well.


| Field | Type | Description |
| ----- | ---- | ----------- |
| resource_type | [ ResourceType](#resourcetype) | Resource type used to build query. |
| alert_type | [ int64](#int64) | Alert type used to build query. |
| resource_id | [ string](#string) | Resource ID used to build query. |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkAlertsResourceTypeQuery {#sdkalertsresourcetypequery}
SdkAlertsResourceTypeQuery queries for alerts using only resource id.


| Field | Type | Description |
| ----- | ---- | ----------- |
| resource_type | [ ResourceType](#resourcetype) | Resource type used to build query. |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkAlertsTimeSpan {#sdkalertstimespan}
SdkAlertsTimeSpan to store time window information.


| Field | Type | Description |
| ----- | ---- | ----------- |
| start_time | [ google.protobuf.Timestamp](#googleprotobuftimestamp) | Start timestamp when Alert occurred |
| end_time | [ google.protobuf.Timestamp](#googleprotobuftimestamp) | End timestamp when Alert occurred |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkAutoFSTrimStatusRequest {#sdkautofstrimstatusrequest}
SdkAutoFSTrimStatusRequest defines a request to get status of autofs trim operation

 <!-- end HasFields -->


## SdkAutoFSTrimStatusResponse {#sdkautofstrimstatusresponse}
SdkAutoFSTrimStatusResponse defines the response for a
SdkAutoFSTrimStatusRequest.


| Field | Type | Description |
| ----- | ---- | ----------- |
| trim_status | [map SdkAutoFSTrimStatusResponse.TrimStatusEntry](#sdkautofstrimstatusresponsetrimstatusentry) | map of volume id and the state of the filesystem trim operation |
| message | [ string](#string) | Text blob containing ASCII text providing details of the operation |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkAutoFSTrimStatusResponse.TrimStatusEntry {#sdkautofstrimstatusresponsetrimstatusentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ FilesystemTrim.FilesystemTrimStatus](#filesystemtrimfilesystemtrimstatus) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkAutoFSTrimUsageRequest {#sdkautofstrimusagerequest}
SdkAutoFSTrimUsageRequest defines a request to get status of autofs trim operation

 <!-- end HasFields -->


## SdkAutoFSTrimUsageResponse {#sdkautofstrimusageresponse}
SdkAutoFSTrimUsageResponse defines the response for a
SdkAutoFSTrimUsageRequest.


| Field | Type | Description |
| ----- | ---- | ----------- |
| usage | [map SdkAutoFSTrimUsageResponse.UsageEntry](#sdkautofstrimusageresponseusageentry) | map of fstrim disk usage and volume name |
| message | [ string](#string) | Text blob containing ASCII text providing details of the operation |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkAutoFSTrimUsageResponse.UsageEntry {#sdkautofstrimusageresponseusageentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ FstrimVolumeUsageInfo](#fstrimvolumeusageinfo) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkAwsCredentialRequest {#sdkawscredentialrequest}
Defines credentials for Aws/S3 endpoints


| Field | Type | Description |
| ----- | ---- | ----------- |
| access_key | [ string](#string) | Access key |
| secret_key | [ string](#string) | Secret key |
| endpoint | [ string](#string) | Endpoint |
| region | [ string](#string) | Region |
| disable_ssl | [ bool](#bool) | (optional) Disable SSL connection |
| disable_path_style | [ bool](#bool) | (optional) Disable path-style access |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkAwsCredentialResponse {#sdkawscredentialresponse}
Defines the response for AWS/S3 credentials


| Field | Type | Description |
| ----- | ---- | ----------- |
| access_key | [ string](#string) | Access key |
| endpoint | [ string](#string) | Endpoint |
| region | [ string](#string) | Region |
| disable_ssl | [ bool](#bool) | (optional) Disable SSL connection |
| disable_path_style | [ bool](#bool) | (optional) Disable path-style access |
| s3_storage_class | [ string](#string) | (optional) Storage class for s3 puts |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkAzureCredentialRequest {#sdkazurecredentialrequest}
Defines credentials for Azure


| Field | Type | Description |
| ----- | ---- | ----------- |
| account_name | [ string](#string) | Account name |
| account_key | [ string](#string) | Account key |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkAzureCredentialResponse {#sdkazurecredentialresponse}
Defines the response for Azure credentials


| Field | Type | Description |
| ----- | ---- | ----------- |
| account_name | [ string](#string) | Account name |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupCatalogRequest {#sdkcloudbackupcatalogrequest}
Defines a request to get catalog of a backup stored by a cloud provider


| Field | Type | Description |
| ----- | ---- | ----------- |
| backup_id | [ string](#string) | Id of the backup |
| credential_id | [ string](#string) | Credential id describe the credentials for the cloud |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupCatalogResponse {#sdkcloudbackupcatalogresponse}
Defines a response containing the contents of a backup stored by a cloud provider


| Field | Type | Description |
| ----- | ---- | ----------- |
| contents | [repeated string](#string) | Contents is listing of backup contents |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupCreateRequest {#sdkcloudbackupcreaterequest}
Defines a request to create a backup of a volume to the cloud


| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_id | [ string](#string) | VolumeID of the volume for which cloudbackup is requested |
| credential_id | [ string](#string) | Credential id refers to the cloud credentials needed to backup |
| full | [ bool](#bool) | Full indicates if full backup is desired even though incremental is possible |
| task_id | [ string](#string) | TaskId of the task performing this backup. This value can be used for idempotency. |
| labels | [map SdkCloudBackupCreateRequest.LabelsEntry](#sdkcloudbackupcreaterequestlabelsentry) | Labels are list of key value pairs to tag the cloud backup. These labels are stored in the metadata associated with the backup. |
| full_backup_frequency | [ uint32](#uint32) | FullBackupFrequency indicates number of incremental backup after which a fullbackup must be created. This is to override the default value for manual/user triggerred backups and not applicable for scheduled backups Value of 0 retains the default behavior. |
| delete_local | [ bool](#bool) | DeleteLocal indicates if local snap created for backup must be deleted after the backup is complete |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupCreateRequest.LabelsEntry {#sdkcloudbackupcreaterequestlabelsentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupCreateResponse {#sdkcloudbackupcreateresponse}
Empty response


| Field | Type | Description |
| ----- | ---- | ----------- |
| task_id | [ string](#string) | TaskId of the task performing the backup |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupDeleteAllRequest {#sdkcloudbackupdeleteallrequest}
Defines a request to delete all the backups stored by a cloud provider
for a specified volume


| Field | Type | Description |
| ----- | ---- | ----------- |
| src_volume_id | [ string](#string) | id of the volume for the request |
| credential_id | [ string](#string) | Credential id is the credential for cloud to be used for the request |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupDeleteAllResponse {#sdkcloudbackupdeleteallresponse}
Empty response

 <!-- end HasFields -->


## SdkCloudBackupDeleteRequest {#sdkcloudbackupdeleterequest}
Defines a request to delete a single backup stored by a cloud provider


| Field | Type | Description |
| ----- | ---- | ----------- |
| backup_id | [ string](#string) | ID is the ID of the cloud backup |
| credential_id | [ string](#string) | Credential id is the credential for cloud to be used for the request |
| force | [ bool](#bool) | Force Delete cloudbackup even if there are dependencies. This may be needed if the backup is an incremental backup and subsequent backups depend on this backup specified by `backup_id`. |
| bucket | [ string](#string) | Bucket name to which cloud backup belongs to |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupDeleteResponse {#sdkcloudbackupdeleteresponse}
Empty response

 <!-- end HasFields -->


## SdkCloudBackupEnumerateWithFiltersRequest {#sdkcloudbackupenumeratewithfiltersrequest}
Defines a request to list the backups stored by a cloud provider.
The following combinations can be used to get cloud backup information:

* For a specific volume in current cluster: Set `src_volume_id` to your desired volume id
and do not provide `cluster_id` and `all`.
* For a specific volume in a specific cluster: Set `src_volume_id` to your desired volume id
and specify `cluster_id`.
* For a specific volume in all clusters: Set `src_volume_id` to your desired volume id
and set `all` to true, do not provide `cluster_id`.
* For all volumes in current cluster: do not provide `cluster_id`, `volume_id` and `all`.
* For all volumes in a specific cluster: Set `cluster_id` to your desired cluster id
and do not provide `volume_id` and `all`.
* For all volumes in all clusters: Set `all` to true do not provide `volume_id` and `cluster_id`.


| Field | Type | Description |
| ----- | ---- | ----------- |
| src_volume_id | [ string](#string) | (optional) Source id of the volume for the request. |
| cluster_id | [ string](#string) | (optional) Cluster id specifies the cluster for the request |
| credential_id | [ string](#string) | Credential id is the credential for cloud to be used for the request |
| all | [ bool](#bool) | (optional) All indicates if the request should show cloud backups for all clusters or the current cluster. |
| status_filter | [ SdkCloudBackupStatusType](#sdkcloudbackupstatustype) | (optional) enumerates backups that have status specified by this type |
| metadata_filter | [map SdkCloudBackupEnumerateWithFiltersRequest.MetadataFilterEntry](#sdkcloudbackupenumeratewithfiltersrequestmetadatafilterentry) | (optional) Enumerates backups that have tags of this type |
| max_backups | [ uint64](#uint64) | (optional) if caller wished to limit number of backups returned by enumerate |
| continuation_token | [ string](#string) | Returned in the enumerate response if not all of the backups could be returned in the response. |
| cloud_backup_id | [ string](#string) | If one wants to enumerate known backup, set this field to the backup ID naming format :clusteruuidORbicketname/srcVolId-snapId(-incr) |
| missing_src_volumes | [ bool](#bool) | To enumerate cloudbackups for which source volumes do not exist in this cluster |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupEnumerateWithFiltersRequest.MetadataFilterEntry {#sdkcloudbackupenumeratewithfiltersrequestmetadatafilterentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupEnumerateWithFiltersResponse {#sdkcloudbackupenumeratewithfiltersresponse}
Defines a response which lists all the backups stored by a cloud provider


| Field | Type | Description |
| ----- | ---- | ----------- |
| backups | [repeated SdkCloudBackupInfo](#sdkcloudbackupinfo) | none |
| continuation_token | [ string](#string) | if this is not an empty string, callers must pass this to get next list of backups |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupGroupCreateRequest {#sdkcloudbackupgroupcreaterequest}
Defines a request to create a group backup of a group to the cloud


| Field | Type | Description |
| ----- | ---- | ----------- |
| group_id | [ string](#string) | GroupID of the volume for which cloudbackup is requested |
| volume_ids | [repeated string](#string) | VolumeIds are a list of volume IDs to use for the backup request. If multiple of GroupID, Labels or VolumeIDs are specified, volumes matching all of them are backed uup |
| credential_id | [ string](#string) | Credential id refers to the cloud credentials needed to backup |
| full | [ bool](#bool) | Full indicates if full backup is desired even though incremental is possible |
| labels | [map SdkCloudBackupGroupCreateRequest.LabelsEntry](#sdkcloudbackupgroupcreaterequestlabelsentry) | Labels are list of key value pairs to tag the cloud backup. These labels are stored in the metadata associated with the backup. |
| delete_local | [ bool](#bool) | DeleteLocal indicates if local snap created for backup must be deleted after the backup is complete |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupGroupCreateRequest.LabelsEntry {#sdkcloudbackupgroupcreaterequestlabelsentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupGroupCreateResponse {#sdkcloudbackupgroupcreateresponse}
Empty response


| Field | Type | Description |
| ----- | ---- | ----------- |
| group_cloud_backup_id | [ string](#string) | ID for this group of backups |
| task_ids | [repeated string](#string) | TaskIds of the tasks performing the group backup |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupHistoryItem {#sdkcloudbackuphistoryitem}
SdkCloudBackupHistoryItem contains information about a backup for a
specific volume


| Field | Type | Description |
| ----- | ---- | ----------- |
| src_volume_id | [ string](#string) | SrcVolumeID is volume ID which was backedup |
| timestamp | [ google.protobuf.Timestamp](#googleprotobuftimestamp) | TimeStamp is the time at which either backup completed/failed |
| status | [ SdkCloudBackupStatusType](#sdkcloudbackupstatustype) | Status indicates whether backup was completed/failed |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupHistoryRequest {#sdkcloudbackuphistoryrequest}
Defines a request to retrieve the history of the backups for
a specific volume to a cloud provider


| Field | Type | Description |
| ----- | ---- | ----------- |
| src_volume_id | [ string](#string) | This optional value defines which history of backups is being requested. If not provided, it will return the history for all volumes. |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupHistoryResponse {#sdkcloudbackuphistoryresponse}
Defines a response containing a list of history of backups to a cloud provider


| Field | Type | Description |
| ----- | ---- | ----------- |
| history_list | [repeated SdkCloudBackupHistoryItem](#sdkcloudbackuphistoryitem) | HistoryList is list of past backups on this volume |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupInfo {#sdkcloudbackupinfo}
SdkCloudBackupInfo has information about a backup stored by a cloud provider


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ string](#string) | This is the id as represented by the cloud provider |
| src_volume_id | [ string](#string) | Source volumeID of the backup |
| src_volume_name | [ string](#string) | Name of the sourceVolume of the backup |
| timestamp | [ google.protobuf.Timestamp](#googleprotobuftimestamp) | Timestamp is the timestamp at which the source volume was backed up to cloud |
| metadata | [map SdkCloudBackupInfo.MetadataEntry](#sdkcloudbackupinfometadataentry) | Metadata associated with the backup |
| status | [ SdkCloudBackupStatusType](#sdkcloudbackupstatustype) | Status indicates the status of the backup |
| cluster_type | [ SdkCloudBackupClusterType](#sdkcloudbackupclustertype) | cluster indicates if the cloudbackup belongs to current cluster, with older cluster this value may be unknown |
| namespace | [ string](#string) | k8s namespace to which this backup belongs to |
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
Defines a request to restore a volume from an existing backup stored by
a cloud provider


| Field | Type | Description |
| ----- | ---- | ----------- |
| backup_id | [ string](#string) | Backup ID being restored |
| restore_volume_name | [ string](#string) | Optional volume Name of the new volume to be created in the cluster for restoring the cloudbackup |
| credential_id | [ string](#string) | The credential to be used for restore operation |
| node_id | [ string](#string) | Optional for provisioning restore volume (ResoreVolumeName should not be specified) |
| task_id | [ string](#string) | TaskId of the task performing this restore |
| spec | [ RestoreVolumeSpec](#restorevolumespec) | Modifiable Restore volume spec |
| locator | [ VolumeLocator](#volumelocator) | RestoreVolume locator |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupRestoreResponse {#sdkcloudbackuprestoreresponse}
Defines a response when restoring a volume from a backup stored by
a cloud provider


| Field | Type | Description |
| ----- | ---- | ----------- |
| restore_volume_id | [ string](#string) | VolumeID to which the backup is being restored |
| task_id | [ string](#string) | TaskId of the task performing the restore |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupSchedCreateRequest {#sdkcloudbackupschedcreaterequest}
Defines a request to create a schedule for volume backups to a
cloud provider


| Field | Type | Description |
| ----- | ---- | ----------- |
| cloud_sched_info | [ SdkCloudBackupScheduleInfo](#sdkcloudbackupscheduleinfo) | Cloud Backup Schedule info |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupSchedCreateResponse {#sdkcloudbackupschedcreateresponse}
Defines a response containing the id of a schedule for a volume backup
to a cloud provider


| Field | Type | Description |
| ----- | ---- | ----------- |
| backup_schedule_id | [ string](#string) | Id of newly created backup schedule |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupSchedDeleteRequest {#sdkcloudbackupscheddeleterequest}
Defines a request to delete a backup schedule


| Field | Type | Description |
| ----- | ---- | ----------- |
| backup_schedule_id | [ string](#string) | Id of cloud backup to delete |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupSchedDeleteResponse {#sdkcloudbackupscheddeleteresponse}
Empty response

 <!-- end HasFields -->


## SdkCloudBackupSchedEnumerateRequest {#sdkcloudbackupschedenumeraterequest}
Empty request

 <!-- end HasFields -->


## SdkCloudBackupSchedEnumerateResponse {#sdkcloudbackupschedenumerateresponse}
Defines a response containing a map listing the schedules for volume
backups to a cloud provider


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


## SdkCloudBackupSchedUpdateRequest {#sdkcloudbackupschedupdaterequest}
Defines a request to update a schedule for volume backups to a
cloud provider


| Field | Type | Description |
| ----- | ---- | ----------- |
| cloud_sched_info | [ SdkCloudBackupScheduleInfo](#sdkcloudbackupscheduleinfo) | Cloud Backup Schedule info |
| sched_uuid | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupSchedUpdateResponse {#sdkcloudbackupschedupdateresponse}
Empty response

 <!-- end HasFields -->


## SdkCloudBackupScheduleInfo {#sdkcloudbackupscheduleinfo}
SdkCloudBackupScheduleInfo describes a schedule for volume backups to
a cloud provider


| Field | Type | Description |
| ----- | ---- | ----------- |
| src_volume_id | [ string](#string) | The schedule's source volume |
| credential_id | [ string](#string) | The cloud credential used with this schedule |
| schedules | [repeated SdkSchedulePolicyInterval](#sdkschedulepolicyinterval) | Schedules are the frequencies of the backup |
| max_backups | [ uint64](#uint64) | MaxBackups indicates when to force full backup to cloud. If RetentionDays is not specified or is 0 (older scheme), this is also the maximum number of scheduled backups retained in the cloud. Older backups are deleted |
| full | [ bool](#bool) | Full indicates if scheduled backups should always be full and never incremental. |
| retention_days | [ uint32](#uint32) | Number of days that Scheduled CloudBackups will be kept after which they are deleted |
| group_id | [ string](#string) | GroupId indicates the group of volumes for which this schedule applies |
| labels | [map SdkCloudBackupScheduleInfo.LabelsEntry](#sdkcloudbackupscheduleinfolabelsentry) | labels indicates group of volumes with similar labels for which this schedule applies |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupScheduleInfo.LabelsEntry {#sdkcloudbackupscheduleinfolabelsentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupSizeRequest {#sdkcloudbackupsizerequest}
Defines a request to retrieve the size of the volume for the
specificed volume


| Field | Type | Description |
| ----- | ---- | ----------- |
| backup_id | [ string](#string) | BackupId is a value which is used to get information on the size of the specified backup. |
| credential_id | [ string](#string) | Credential id describe the credentials for the cloud |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupSizeResponse {#sdkcloudbackupsizeresponse}
Defines a response containing the size of the volume


| Field | Type | Description |
| ----- | ---- | ----------- |
| size | [ uint64](#uint64) | Size is the size of the volume in bytes |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupStateChangeRequest {#sdkcloudbackupstatechangerequest}
Defines a request to change the state of a backup or restore to or
from a cloud provider


| Field | Type | Description |
| ----- | ---- | ----------- |
| task_id | [ string](#string) | Describes the backup/restore task state change is being requested |
| requested_state | [ SdkCloudBackupRequestedState](#sdkcloudbackuprequestedstate) | The desired state of the operation |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupStateChangeResponse {#sdkcloudbackupstatechangeresponse}
Empty response

 <!-- end HasFields -->


## SdkCloudBackupStatus {#sdkcloudbackupstatus}
SdkCloudBackupStatus defines the status of a backup stored by a cloud provider


| Field | Type | Description |
| ----- | ---- | ----------- |
| backup_id | [ string](#string) | This is the id as represented by the cloud provider |
| optype | [ SdkCloudBackupOpType](#sdkcloudbackupoptype) | OpType indicates if this is a backup or restore |
| status | [ SdkCloudBackupStatusType](#sdkcloudbackupstatustype) | State indicates if the op is currently active/done/failed |
| bytes_done | [ uint64](#uint64) | BytesDone indicates total Bytes uploaded/downloaded |
| start_time | [ google.protobuf.Timestamp](#googleprotobuftimestamp) | StartTime indicates Op's start time |
| completed_time | [ google.protobuf.Timestamp](#googleprotobuftimestamp) | CompletedTime indicates Op's completed time |
| node_id | [ string](#string) | NodeID is the ID of the node where this Op is active |
| src_volume_id | [ string](#string) | SourceVolumeID is the the volume that is either being backed up to cloud or target volume to which a backup is being restored |
| info | [repeated string](#string) | Info currently indicates the failure cause for failed backup/restore |
| credential_id | [ string](#string) | CredentialId is the credential used for cloud with this backup/restore op |
| bytes_total | [ uint64](#uint64) | BytesTotal is the total number of bytes being transferred |
| eta_seconds | [ int64](#int64) | ETASeconds is the number of seconds for cloud backup completion |
| group_id | [ string](#string) | string group_id volume's group id if this was group cloud backup |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupStatusRequest {#sdkcloudbackupstatusrequest}
Defines a request to retrieve the status of a backup or restore for a
specified volume


| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_id | [ string](#string) | (optional) VolumeId is a value which is used to get information on the status of a backup for the specified volume. If no volume id and task_id is provided, then status for all volumes is returned. |
| local | [ bool](#bool) | Local indicates if only those backups/restores that are active on current node must be returned |
| task_id | [ string](#string) | TaskId of the backup/restore task, if this is specified, volume_id is ignored. |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupStatusResponse {#sdkcloudbackupstatusresponse}
Defines a response containing the status of the backups for a specified volume


| Field | Type | Description |
| ----- | ---- | ----------- |
| statuses | [map SdkCloudBackupStatusResponse.StatusesEntry](#sdkcloudbackupstatusresponsestatusesentry) | Statuses is list of currently active/failed/done backup/restores where the key is the id of the task performing backup/restore. |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudBackupStatusResponse.StatusesEntry {#sdkcloudbackupstatusresponsestatusesentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ SdkCloudBackupStatus](#sdkcloudbackupstatus) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudMigrateCancelRequest {#sdkcloudmigratecancelrequest}
Defines a request to stop a cloud migration


| Field | Type | Description |
| ----- | ---- | ----------- |
| request | [ CloudMigrateCancelRequest](#cloudmigratecancelrequest) | Request containing the task id to be cancelled |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudMigrateCancelResponse {#sdkcloudmigratecancelresponse}
Empty Response

 <!-- end HasFields -->


## SdkCloudMigrateStartRequest {#sdkcloudmigratestartrequest}
Defines a migration request


| Field | Type | Description |
| ----- | ---- | ----------- |
| cluster_id | [ string](#string) | ID of the cluster to which volumes are to be migrated |
| task_id | [ string](#string) | Unique name assocaiated with this migration. This is a Optional field for idempotency |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) opt.volume | [ SdkCloudMigrateStartRequest.MigrateVolume](#sdkcloudmigratestartrequestmigratevolume) | Request to migrate a volume |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) opt.volume_group | [ SdkCloudMigrateStartRequest.MigrateVolumeGroup](#sdkcloudmigratestartrequestmigratevolumegroup) | Request to migrate a volume group |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) opt.all_volumes | [ SdkCloudMigrateStartRequest.MigrateAllVolumes](#sdkcloudmigratestartrequestmigrateallvolumes) | Request to migrate all volumes |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudMigrateStartRequest.MigrateAllVolumes {#sdkcloudmigratestartrequestmigrateallvolumes}
Defines a migration request for all volumes in a cluster

 <!-- end HasFields -->


## SdkCloudMigrateStartRequest.MigrateVolume {#sdkcloudmigratestartrequestmigratevolume}
Defines a migration request for a volume


| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_id | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudMigrateStartRequest.MigrateVolumeGroup {#sdkcloudmigratestartrequestmigratevolumegroup}
Defines a migration request for a volume group


| Field | Type | Description |
| ----- | ---- | ----------- |
| group_id | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudMigrateStartResponse {#sdkcloudmigratestartresponse}
Defines a response for the migration that was started


| Field | Type | Description |
| ----- | ---- | ----------- |
| result | [ CloudMigrateStartResponse](#cloudmigratestartresponse) | Result assocaiated with the migration that was started |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudMigrateStatusRequest {#sdkcloudmigratestatusrequest}
Request for cloud migration operation status


| Field | Type | Description |
| ----- | ---- | ----------- |
| request | [ CloudMigrateStatusRequest](#cloudmigratestatusrequest) | Request contains the task id and cluster id for which status should be returned |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCloudMigrateStatusResponse {#sdkcloudmigratestatusresponse}
Defines a response for the status request


| Field | Type | Description |
| ----- | ---- | ----------- |
| result | [ CloudMigrateStatusResponse](#cloudmigratestatusresponse) | Status of all migration requests |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkClusterDomainActivateRequest {#sdkclusterdomainactivaterequest}
Defines a request to activate a cluster domain


| Field | Type | Description |
| ----- | ---- | ----------- |
| cluster_domain_name | [ string](#string) | Name of the cluster domain to activate |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkClusterDomainActivateResponse {#sdkclusterdomainactivateresponse}
Empty response

 <!-- end HasFields -->


## SdkClusterDomainDeactivateRequest {#sdkclusterdomaindeactivaterequest}
Defines a request to deactivate a cluster domain


| Field | Type | Description |
| ----- | ---- | ----------- |
| cluster_domain_name | [ string](#string) | Name of the cluster domain to deactivate |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkClusterDomainDeactivateResponse {#sdkclusterdomaindeactivateresponse}
Empty response

 <!-- end HasFields -->


## SdkClusterDomainInspectRequest {#sdkclusterdomaininspectrequest}
Defines a request to inspect a cluster domain


| Field | Type | Description |
| ----- | ---- | ----------- |
| cluster_domain_name | [ string](#string) | Name of the cluster domain to inspect |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkClusterDomainInspectResponse {#sdkclusterdomaininspectresponse}
Defines a response to inspecting a cluster domain


| Field | Type | Description |
| ----- | ---- | ----------- |
| cluster_domain_name | [ string](#string) | Name of the cluster domain |
| is_active | [ bool](#bool) | IsActive indicates whether this cluster domain is active |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkClusterDomainsEnumerateRequest {#sdkclusterdomainsenumeraterequest}
Empty request

 <!-- end HasFields -->


## SdkClusterDomainsEnumerateResponse {#sdkclusterdomainsenumerateresponse}
Defines a response when enumerating cluster domains


| Field | Type | Description |
| ----- | ---- | ----------- |
| cluster_domain_names | [repeated string](#string) | List of names of all the cluster domains in a cluster |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkClusterInspectCurrentRequest {#sdkclusterinspectcurrentrequest}
Empty request

 <!-- end HasFields -->


## SdkClusterInspectCurrentResponse {#sdkclusterinspectcurrentresponse}
Defines a response when inspecting the current cluster


| Field | Type | Description |
| ----- | ---- | ----------- |
| cluster | [ StorageCluster](#storagecluster) | Cluster information |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkClusterPairCreateRequest {#sdkclusterpaircreaterequest}
Defines a request for creating a cluster pair


| Field | Type | Description |
| ----- | ---- | ----------- |
| request | [ ClusterPairCreateRequest](#clusterpaircreaterequest) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkClusterPairCreateResponse {#sdkclusterpaircreateresponse}
Defines a result of the cluster pair


| Field | Type | Description |
| ----- | ---- | ----------- |
| result | [ ClusterPairCreateResponse](#clusterpaircreateresponse) | Contains the information about cluster pair |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkClusterPairDeleteRequest {#sdkclusterpairdeleterequest}
Defines a delete request for a cluster pair


| Field | Type | Description |
| ----- | ---- | ----------- |
| cluster_id | [ string](#string) | ID of the cluster pair to be deleted |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkClusterPairDeleteResponse {#sdkclusterpairdeleteresponse}
Empty response

 <!-- end HasFields -->


## SdkClusterPairEnumerateRequest {#sdkclusterpairenumeraterequest}
Empty Request

 <!-- end HasFields -->


## SdkClusterPairEnumerateResponse {#sdkclusterpairenumerateresponse}
Defines a list of cluster pair


| Field | Type | Description |
| ----- | ---- | ----------- |
| result | [ ClusterPairsEnumerateResponse](#clusterpairsenumerateresponse) | List of all the cluster pairs |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkClusterPairGetTokenRequest {#sdkclusterpairgettokenrequest}
Empty request

 <!-- end HasFields -->


## SdkClusterPairGetTokenResponse {#sdkclusterpairgettokenresponse}
Defines a response for the token request


| Field | Type | Description |
| ----- | ---- | ----------- |
| result | [ ClusterPairTokenGetResponse](#clusterpairtokengetresponse) | Contains authentication token for the cluster |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkClusterPairInspectRequest {#sdkclusterpairinspectrequest}
Defines a cluster pair inspect request


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ string](#string) | ID of the cluster, if empty gets the default pair |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkClusterPairInspectResponse {#sdkclusterpairinspectresponse}
Defines a cluster pair inspect response


| Field | Type | Description |
| ----- | ---- | ----------- |
| result | [ ClusterPairGetResponse](#clusterpairgetresponse) | Information about cluster pair |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkClusterPairResetTokenRequest {#sdkclusterpairresettokenrequest}
Empty request

 <!-- end HasFields -->


## SdkClusterPairResetTokenResponse {#sdkclusterpairresettokenresponse}
Defines a response for the token request


| Field | Type | Description |
| ----- | ---- | ----------- |
| result | [ ClusterPairTokenGetResponse](#clusterpairtokengetresponse) | Contains authentication token for the cluster |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCredentialCreateRequest {#sdkcredentialcreaterequest}
Defines a request to create credentials


| Field | Type | Description |
| ----- | ---- | ----------- |
| name | [ string](#string) | Name of the credential |
| bucket | [ string](#string) | (optional) Name of bucket |
| encryption_key | [ string](#string) | (optional) Key used to encrypt the data |
| ownership | [ Ownership](#ownership) | Ownership of the credential. Collaborators and groups may be added here with their appropriate ACLS. |
| use_proxy | [ bool](#bool) | use_proxy indicates if a proxy must be used |
| iam_policy | [ bool](#bool) | iamPolicy indicates if IAM creds must be used for access |
| s3_storage_class | [ string](#string) | s3StorageClass for object puts, empty indicates default STANDARD |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) credential_type.aws_credential | [ SdkAwsCredentialRequest](#sdkawscredentialrequest) | Credentials for AWS/S3 |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) credential_type.azure_credential | [ SdkAzureCredentialRequest](#sdkazurecredentialrequest) | Credentials for Azure |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) credential_type.google_credential | [ SdkGoogleCredentialRequest](#sdkgooglecredentialrequest) | Credentials for Google |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) credential_type.nfs_credential | [ SdkNfsCredentialRequest](#sdknfscredentialrequest) | Credentials for NFS |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCredentialCreateResponse {#sdkcredentialcreateresponse}
Defines a response from creating a credential


| Field | Type | Description |
| ----- | ---- | ----------- |
| credential_id | [ string](#string) | Id of the credentials |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCredentialDeleteReferencesRequest {#sdkcredentialdeletereferencesrequest}
Defines a request to remove any references to credentials


| Field | Type | Description |
| ----- | ---- | ----------- |
| credential_id | [ string](#string) | Id of the credentials |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCredentialDeleteReferencesResponse {#sdkcredentialdeletereferencesresponse}
Empty response

 <!-- end HasFields -->


## SdkCredentialDeleteRequest {#sdkcredentialdeleterequest}
Defines the request to delete credentials


| Field | Type | Description |
| ----- | ---- | ----------- |
| credential_id | [ string](#string) | Id for credentials |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCredentialDeleteResponse {#sdkcredentialdeleteresponse}
Empty response

 <!-- end HasFields -->


## SdkCredentialEnumerateRequest {#sdkcredentialenumeraterequest}
Empty request

 <!-- end HasFields -->


## SdkCredentialEnumerateResponse {#sdkcredentialenumerateresponse}
Defines response for a enumeration of credentials


| Field | Type | Description |
| ----- | ---- | ----------- |
| credential_ids | [repeated string](#string) | List of credentials |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCredentialInspectRequest {#sdkcredentialinspectrequest}
Defines the request to inspection for credentials


| Field | Type | Description |
| ----- | ---- | ----------- |
| credential_id | [ string](#string) | Id of the credential |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCredentialInspectResponse {#sdkcredentialinspectresponse}
Defines the response to an inspection of a credential.
This response uses OneOf proto style. Depending on your programming language
you will need to check if the value of credential_type is one of the ones below.


| Field | Type | Description |
| ----- | ---- | ----------- |
| credential_id | [ string](#string) | Credential id |
| name | [ string](#string) | Name of the credential |
| bucket | [ string](#string) | (optional) Name of bucket |
| ownership | [ Ownership](#ownership) | Ownership of the credential |
| use_proxy | [ bool](#bool) | proxy flag for the credential |
| iam_policy | [ bool](#bool) | iamPolicy indicates if IAM creds must be used for access |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) credential_type.aws_credential | [ SdkAwsCredentialResponse](#sdkawscredentialresponse) | Aws credentials |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) credential_type.azure_credential | [ SdkAzureCredentialResponse](#sdkazurecredentialresponse) | Azure credentials |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) credential_type.google_credential | [ SdkGoogleCredentialResponse](#sdkgooglecredentialresponse) | Google credentials |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) credential_type.nfs_credential | [ SdkNfsCredentialResponse](#sdknfscredentialresponse) | NFS credentials |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCredentialUpdateRequest {#sdkcredentialupdaterequest}
Defines request for credential update


| Field | Type | Description |
| ----- | ---- | ----------- |
| credential_id | [ string](#string) | none |
| update_req | [ SdkCredentialCreateRequest](#sdkcredentialcreaterequest) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCredentialUpdateResponse {#sdkcredentialupdateresponse}
Defines response for credential update

 <!-- end HasFields -->


## SdkCredentialValidateRequest {#sdkcredentialvalidaterequest}
Defines a request to validate credentials


| Field | Type | Description |
| ----- | ---- | ----------- |
| credential_id | [ string](#string) | Id of the credentials |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkCredentialValidateResponse {#sdkcredentialvalidateresponse}
Empty response

 <!-- end HasFields -->


## SdkDiagsCollectRequest {#sdkdiagscollectrequest}
SdkDiagsCollectRequest is the request object that specifies what should be part of the diags that are collected
User can specify both Node and Volume or just one of them. If both are provided, the implementation will select
nodes based on both and also handle overlaps


| Field | Type | Description |
| ----- | ---- | ----------- |
| node | [ DiagsNodeSelector](#diagsnodeselector) | Node selects the node(s) for diags collection |
| volume | [ DiagsVolumeSelector](#diagsvolumeselector) | Volume selects the volume(s) for diags collection |
| profile_only | [ bool](#bool) | ProfileOnly is an optional flag if true will only collect the stack and heap profile of the driver and will skip other diag components |
| issuer | [ string](#string) | Issuer is an optional user friendly name for the caller invoking the API |
| timeout_mins | [ int64](#int64) | TimeoutMins is the timeout in minutes for the job. This is an optional field and if not provided, the implementation of the SDK will use a sane default |
| live | [ bool](#bool) | Live is an optional flag if true will collect live cores from running processes of the driver |
| filename | [ string](#string) | Filename is an optional flag only to be used for testing purposes. |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkDiagsCollectResponse {#sdkdiagscollectresponse}
SdkDiagsCollectResponse defines a response for an SDK request to collect diags


| Field | Type | Description |
| ----- | ---- | ----------- |
| job | [ Job](#job) | Job that was created for the SDK request |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkEnumerateJobsRequest {#sdkenumeratejobsrequest}
Defines a request to list all the  jobs


| Field | Type | Description |
| ----- | ---- | ----------- |
| type | [ Job.Type](#jobtype) | Type if specified will list the jobs of the provided type |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkEnumerateJobsResponse {#sdkenumeratejobsresponse}
Defines a response will all the known jobs


| Field | Type | Description |
| ----- | ---- | ----------- |
| jobs | [repeated Job](#job) | Jobs is the list of jobs in the response |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkEnumerateRebalanceJobsRequest {#sdkenumeraterebalancejobsrequest}


 <!-- end HasFields -->


## SdkEnumerateRebalanceJobsResponse {#sdkenumeraterebalancejobsresponse}



| Field | Type | Description |
| ----- | ---- | ----------- |
| jobs | [repeated StorageRebalanceJob](#storagerebalancejob) | Jobs is the list of rebalance jobs in the response |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkFilesystemCheckStartRequest {#sdkfilesystemcheckstartrequest}
SdkFilesystemCheckStartRequest defines a request to start a background
filesystem consistency check operation


| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_id | [ string](#string) | Id of the volume |
| mode | [ string](#string) | Mode of operation |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkFilesystemCheckStartResponse {#sdkfilesystemcheckstartresponse}
SdkFilesystemCheckStartResponse defines the response for a
SdkFilesystemCheckStartRequest.


| Field | Type | Description |
| ----- | ---- | ----------- |
| status | [ FilesystemCheck.FilesystemCheckStatus](#filesystemcheckfilesystemcheckstatus) | Status code representing the state of the filesystem check operation |
| message | [ string](#string) | Text blob containing ASCII text providing details of the operation |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkFilesystemCheckStatusRequest {#sdkfilesystemcheckstatusrequest}
SdkFilesystemCheckStatusRequest defines a request to get status of a
background filesystem check operation


| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_id | [ string](#string) | Id of the volume |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkFilesystemCheckStatusResponse {#sdkfilesystemcheckstatusresponse}
SdkFilesystemCheckStatusResponse defines the response for a
SdkFilesystemCheckStatusRequest.


| Field | Type | Description |
| ----- | ---- | ----------- |
| status | [ FilesystemCheck.FilesystemCheckStatus](#filesystemcheckfilesystemcheckstatus) | Status code representing the state of the filesystem check operation |
| health_status | [ FilesystemHealthStatus](#filesystemhealthstatus) | Status code representing the health of the filesystem after a checkHealth operation |
| mode | [ string](#string) | Text string representing the mode of filesystem check operation |
| message | [ string](#string) | Text blob containing ASCII text providing details of the operation |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkFilesystemCheckStopRequest {#sdkfilesystemcheckstoprequest}
SdkFilesystemCheckStopRequest defines a request to stop a background
filesystem check operation


| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_id | [ string](#string) | Id of the volume |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkFilesystemCheckStopResponse {#sdkfilesystemcheckstopresponse}
Empty response

 <!-- end HasFields -->


## SdkFilesystemTrimStartRequest {#sdkfilesystemtrimstartrequest}
SdkFilesystemTrimStartRequest defines a request to start a background filesystem trim operation


| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_id | [ string](#string) | Id of the volume |
| mount_path | [ string](#string) | Path where the volume is mounted |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkFilesystemTrimStartResponse {#sdkfilesystemtrimstartresponse}
SdkFilesystemTrimStartResponse defines the response for a
SdkFilesystemTrimStartRequest.


| Field | Type | Description |
| ----- | ---- | ----------- |
| status | [ FilesystemTrim.FilesystemTrimStatus](#filesystemtrimfilesystemtrimstatus) | Status code representing the state of the filesystem trim operation |
| message | [ string](#string) | Text blob containing ASCII text providing details of the operation |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkFilesystemTrimStatusRequest {#sdkfilesystemtrimstatusrequest}
SdkFilesystemTrimStatusRequest defines a request to get status of a
background filesystem trim operation


| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_id | [ string](#string) | Id of the volume |
| mount_path | [ string](#string) | Path where the volume is mounted |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkFilesystemTrimStatusResponse {#sdkfilesystemtrimstatusresponse}
SdkFilesystemTrimStatusResponse defines the response for a
SdkFilesystemTrimStatusRequest.


| Field | Type | Description |
| ----- | ---- | ----------- |
| status | [ FilesystemTrim.FilesystemTrimStatus](#filesystemtrimfilesystemtrimstatus) | Status code representing the state of the filesystem trim operation |
| message | [ string](#string) | Text blob containing ASCII text providing details of the operation |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkFilesystemTrimStopRequest {#sdkfilesystemtrimstoprequest}
SdkFilesystemTrimStopRequest defines a request to stop a background
filesystem trim operation


| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_id | [ string](#string) | Id of the volume |
| mount_path | [ string](#string) | Path where the volume is mounted |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkFilesystemTrimStopResponse {#sdkfilesystemtrimstopresponse}
Empty response

 <!-- end HasFields -->


## SdkGetJobStatusRequest {#sdkgetjobstatusrequest}
Defines a request to get the status of an existing job


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ string](#string) | ID of the job |
| type | [ Job.Type](#jobtype) | Type of the job |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkGetJobStatusResponse {#sdkgetjobstatusresponse}
Defines the status of an existing job


| Field | Type | Description |
| ----- | ---- | ----------- |
| job | [ Job](#job) | Job for this node drain operation. |
| summary | [ JobSummary](#jobsummary) | Summary of this job |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkGetRebalanceJobStatusRequest {#sdkgetrebalancejobstatusrequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ string](#string) | ID of the rebalance job |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkGetRebalanceJobStatusResponse {#sdkgetrebalancejobstatusresponse}



| Field | Type | Description |
| ----- | ---- | ----------- |
| job | [ StorageRebalanceJob](#storagerebalancejob) | Job for this rebalance |
| summary | [ StorageRebalanceSummary](#storagerebalancesummary) | Summary summarizes the rebalance job |
| actions | [repeated StorageRebalanceAudit](#storagerebalanceaudit) | Actions describe all the actions taken during this rebalance |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkGoogleCredentialRequest {#sdkgooglecredentialrequest}
Defines credentials for Google


| Field | Type | Description |
| ----- | ---- | ----------- |
| project_id | [ string](#string) | Project ID |
| json_key | [ string](#string) | JSON Key |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkGoogleCredentialResponse {#sdkgooglecredentialresponse}
Defines the response for Google credentials


| Field | Type | Description |
| ----- | ---- | ----------- |
| project_id | [ string](#string) | Project ID |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkIdentityCapabilitiesRequest {#sdkidentitycapabilitiesrequest}
Empty request

 <!-- end HasFields -->


## SdkIdentityCapabilitiesResponse {#sdkidentitycapabilitiesresponse}
Defines a response containing the capabilities of the cluster


| Field | Type | Description |
| ----- | ---- | ----------- |
| capabilities | [repeated SdkServiceCapability](#sdkservicecapability) | Provides all the capabilities supported by the cluster |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkIdentityVersionRequest {#sdkidentityversionrequest}
Empty request

 <!-- end HasFields -->


## SdkIdentityVersionResponse {#sdkidentityversionresponse}
Defines a response containing version information


| Field | Type | Description |
| ----- | ---- | ----------- |
| sdk_version | [ SdkVersion](#sdkversion) | OpenStorage SDK version used by the server |
| version | [ StorageVersion](#storageversion) | Version information about the storage system |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkJobResponse {#sdkjobresponse}
Defines a response for an SDK request that spins up a new job
to perform the request


| Field | Type | Description |
| ----- | ---- | ----------- |
| job | [ Job](#job) | Job that was created for the SDK request |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkNfsCredentialRequest {#sdknfscredentialrequest}
Defines credentials for NFS


| Field | Type | Description |
| ----- | ---- | ----------- |
| server | [ string](#string) | NFS Server address |
| sub_path | [ string](#string) | NFS export path |
| mount_opts | [ string](#string) | mount options for nfs mount |
| timeout_seconds | [ uint32](#uint32) | timeout for nfs IO in seconds |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkNfsCredentialResponse {#sdknfscredentialresponse}
Defines the response for NFS credential


| Field | Type | Description |
| ----- | ---- | ----------- |
| server | [ string](#string) | NFS Server Address |
| sub_path | [ string](#string) | NFS export path |
| mount_opts | [ string](#string) | mount options ( "," separated) |
| timeout_seconds | [ uint32](#uint32) | timeout in seconds |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkNodeCordonAttachmentsRequest {#sdknodecordonattachmentsrequest}
SdkNodeCordonAttachmentsRequest request for disabling new volume
attachments from a node


| Field | Type | Description |
| ----- | ---- | ----------- |
| node_id | [ string](#string) | Node ID on which any further volume attachments will be disabled |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkNodeCordonAttachmentsResponse {#sdknodecordonattachmentsresponse}
SdkNodeCordonAttachmentsRespinse response for disabling new volume
attachments from a node

 <!-- end HasFields -->


## SdkNodeDrainAttachmentsRequest {#sdknodedrainattachmentsrequest}
Defines a node drain volume attachments request


| Field | Type | Description |
| ----- | ---- | ----------- |
| node_id | [ string](#string) | Id of the node to drain |
| selector | [repeated LabelSelectorRequirement](#labelselectorrequirement) | Selector is used for selecting volumes whose attachment needs to be moved from this node. The selector could be a list of volume label key value pairs to select a subset of volumes. |
| only_sharedv4 | [ bool](#bool) | Drain only sharedv4 volumes from a node By default all volumes will be drained. |
| issuer | [ string](#string) | Issuer is a user friendly name for the caller who is invoking the API. It can be used by caller to filter out drain requests from a particular issuer |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkNodeEnumerateRequest {#sdknodeenumeraterequest}
Empty request

 <!-- end HasFields -->


## SdkNodeEnumerateResponse {#sdknodeenumerateresponse}
Defines a response with a list of node ids


| Field | Type | Description |
| ----- | ---- | ----------- |
| node_ids | [repeated string](#string) | List of all the node ids in the cluster |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkNodeEnumerateWithFiltersRequest {#sdknodeenumeratewithfiltersrequest}
Defines a request to list nodes with given filter. Currently there are
no filters and all the nodes will be returned.

 <!-- end HasFields -->


## SdkNodeEnumerateWithFiltersResponse {#sdknodeenumeratewithfiltersresponse}
Defines a response with a list of nodes


| Field | Type | Description |
| ----- | ---- | ----------- |
| nodes | [repeated StorageNode](#storagenode) | List of all the nodes in the cluster |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkNodeInspectCurrentRequest {#sdknodeinspectcurrentrequest}
Empty request

 <!-- end HasFields -->


## SdkNodeInspectCurrentResponse {#sdknodeinspectcurrentresponse}
Defines a response when inspecting a node


| Field | Type | Description |
| ----- | ---- | ----------- |
| node | [ StorageNode](#storagenode) | Node information |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkNodeInspectRequest {#sdknodeinspectrequest}
Defines a request when inspecting a node


| Field | Type | Description |
| ----- | ---- | ----------- |
| node_id | [ string](#string) | Id of node to inspect |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkNodeInspectResponse {#sdknodeinspectresponse}
Defines a response when inspecting a node


| Field | Type | Description |
| ----- | ---- | ----------- |
| node | [ StorageNode](#storagenode) | Node information |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkNodeRelaxedReclaimPurgeRequest {#sdknoderelaxedreclaimpurgerequest}
Defines request to trigger RelaxedReclaim purge
for a given node


| Field | Type | Description |
| ----- | ---- | ----------- |
| node_id | [ string](#string) | Id of the node to trigger the purge |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkNodeRelaxedReclaimPurgeResponse {#sdknoderelaxedreclaimpurgeresponse}
Defines response containing status of the trigger


| Field | Type | Description |
| ----- | ---- | ----------- |
| status | [ RelaxedReclaimPurge](#relaxedreclaimpurge) | status returns true on successful purge |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkNodeUncordonAttachmentsRequest {#sdknodeuncordonattachmentsrequest}
SdkNodeUncordonAttachmentsRequest request for re-enabling volume
attachments for a node


| Field | Type | Description |
| ----- | ---- | ----------- |
| node_id | [ string](#string) | Node ID on which any further volume attachments will be enabled |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkNodeUncordonAttachmentsResponse {#sdknodeuncordonattachmentsresponse}
SdkNodeUncordonAttachmentsRespinse response for enabling new volume
attachments from a node

 <!-- end HasFields -->


## SdkNodeVolumeUsageByNodeRequest {#sdknodevolumeusagebynoderequest}
Defines request to retrieve all volumes/snapshots capacity usage details
for a given node


| Field | Type | Description |
| ----- | ---- | ----------- |
| node_id | [ string](#string) | Id of the node to get snapshot/volumes capacity usage details |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkNodeVolumeUsageByNodeResponse {#sdknodevolumeusagebynoderesponse}
Defines response containing Node's volumes/snapshot capacity usage details


| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_usage_info | [ VolumeUsageByNode](#volumeusagebynode) | VolumeUsageByNode details |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkObjectstoreCreateRequest {#sdkobjectstorecreaterequest}
Defines a request to create an object store


| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_id | [ string](#string) | Volume on which objectstore will be running |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkObjectstoreCreateResponse {#sdkobjectstorecreateresponse}
Defines a response when an object store has been created for a
specified volume


| Field | Type | Description |
| ----- | ---- | ----------- |
| objectstore_status | [ ObjectstoreInfo](#objectstoreinfo) | Created objecstore status |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkObjectstoreDeleteRequest {#sdkobjectstoredeleterequest}
Defines a request to delete an object store service from a volume


| Field | Type | Description |
| ----- | ---- | ----------- |
| objectstore_id | [ string](#string) | Id of the object store to delete |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkObjectstoreDeleteResponse {#sdkobjectstoredeleteresponse}
Empty response

 <!-- end HasFields -->


## SdkObjectstoreInspectRequest {#sdkobjectstoreinspectrequest}
Defines a request to get information about an object store endpoint


| Field | Type | Description |
| ----- | ---- | ----------- |
| objectstore_id | [ string](#string) | Id of the object store |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkObjectstoreInspectResponse {#sdkobjectstoreinspectresponse}
Defines a response when inspecting an object store endpoint


| Field | Type | Description |
| ----- | ---- | ----------- |
| objectstore_status | [ ObjectstoreInfo](#objectstoreinfo) | Contains information about the object store requested |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkObjectstoreUpdateRequest {#sdkobjectstoreupdaterequest}
Defines a request to update an object store


| Field | Type | Description |
| ----- | ---- | ----------- |
| objectstore_id | [ string](#string) | Objectstore Id to update |
| enable | [ bool](#bool) | enable/disable objectstore |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkObjectstoreUpdateResponse {#sdkobjectstoreupdateresponse}
Empty response

 <!-- end HasFields -->


## SdkOpenStoragePolicyCreateRequest {#sdkopenstoragepolicycreaterequest}
Define a request to create storage policy


| Field | Type | Description |
| ----- | ---- | ----------- |
| storage_policy | [ SdkStoragePolicy](#sdkstoragepolicy) | storage policy to create |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkOpenStoragePolicyCreateResponse {#sdkopenstoragepolicycreateresponse}
Empty response

 <!-- end HasFields -->


## SdkOpenStoragePolicyDefaultInspectRequest {#sdkopenstoragepolicydefaultinspectrequest}
Empty Request

 <!-- end HasFields -->


## SdkOpenStoragePolicyDefaultInspectResponse {#sdkopenstoragepolicydefaultinspectresponse}
Define default storage policy response


| Field | Type | Description |
| ----- | ---- | ----------- |
| storage_policy | [ SdkStoragePolicy](#sdkstoragepolicy) | storage policy information which is set as default |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkOpenStoragePolicyDeleteRequest {#sdkopenstoragepolicydeleterequest}
Define a request to delete storage policy


| Field | Type | Description |
| ----- | ---- | ----------- |
| name | [ string](#string) | name of storage policy to delete |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkOpenStoragePolicyDeleteResponse {#sdkopenstoragepolicydeleteresponse}
Empty Response

 <!-- end HasFields -->


## SdkOpenStoragePolicyEnumerateRequest {#sdkopenstoragepolicyenumeraterequest}
Empty request

 <!-- end HasFields -->


## SdkOpenStoragePolicyEnumerateResponse {#sdkopenstoragepolicyenumerateresponse}
Define a storage policy enumerate response


| Field | Type | Description |
| ----- | ---- | ----------- |
| storage_policies | [repeated SdkStoragePolicy](#sdkstoragepolicy) | List of storage policies |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkOpenStoragePolicyInspectRequest {#sdkopenstoragepolicyinspectrequest}
Define a request to inspect storage policy


| Field | Type | Description |
| ----- | ---- | ----------- |
| name | [ string](#string) | name of storage policy to retrive |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkOpenStoragePolicyInspectResponse {#sdkopenstoragepolicyinspectresponse}
Define a storage policy inspect response


| Field | Type | Description |
| ----- | ---- | ----------- |
| storage_policy | [ SdkStoragePolicy](#sdkstoragepolicy) | storage policy information requested by name |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkOpenStoragePolicyReleaseRequest {#sdkopenstoragepolicyreleaserequest}
Empty Request

 <!-- end HasFields -->


## SdkOpenStoragePolicyReleaseResponse {#sdkopenstoragepolicyreleaseresponse}
Empty Response

 <!-- end HasFields -->


## SdkOpenStoragePolicySetDefaultRequest {#sdkopenstoragepolicysetdefaultrequest}
Define a request to set default storage policy


| Field | Type | Description |
| ----- | ---- | ----------- |
| name | [ string](#string) | name of policy to set as default storage policy for volume creation This policy will be used to validate/update volume configuration |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkOpenStoragePolicySetDefaultResponse {#sdkopenstoragepolicysetdefaultresponse}
Empty Response

 <!-- end HasFields -->


## SdkOpenStoragePolicyUpdateRequest {#sdkopenstoragepolicyupdaterequest}
Define a request to update storage policy


| Field | Type | Description |
| ----- | ---- | ----------- |
| storage_policy | [ SdkStoragePolicy](#sdkstoragepolicy) | storage policy to update |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkOpenStoragePolicyUpdateResponse {#sdkopenstoragepolicyupdateresponse}
Empty Response

 <!-- end HasFields -->


## SdkRole {#sdkrole}



| Field | Type | Description |
| ----- | ---- | ----------- |
| name | [ string](#string) | none |
| rules | [repeated SdkRule](#sdkrule) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkRoleCreateRequest {#sdkrolecreaterequest}
Defines a request for creating a role


| Field | Type | Description |
| ----- | ---- | ----------- |
| role | [ SdkRole](#sdkrole) | Role |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkRoleCreateResponse {#sdkrolecreateresponse}
Response contains informaiton about the creation of the role


| Field | Type | Description |
| ----- | ---- | ----------- |
| role | [ SdkRole](#sdkrole) | Role created |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkRoleDeleteRequest {#sdkroledeleterequest}
Defines a request to delete a role


| Field | Type | Description |
| ----- | ---- | ----------- |
| name | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkRoleDeleteResponse {#sdkroledeleteresponse}
Empty response

 <!-- end HasFields -->


## SdkRoleEnumerateRequest {#sdkroleenumeraterequest}
Empty request

 <!-- end HasFields -->


## SdkRoleEnumerateResponse {#sdkroleenumerateresponse}
Respose to enumerate all roles


| Field | Type | Description |
| ----- | ---- | ----------- |
| names | [repeated string](#string) | List of role names |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkRoleInspectRequest {#sdkroleinspectrequest}
Defines a request to inspect a role


| Field | Type | Description |
| ----- | ---- | ----------- |
| name | [ string](#string) | Name of role |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkRoleInspectResponse {#sdkroleinspectresponse}
Response to inspection request


| Field | Type | Description |
| ----- | ---- | ----------- |
| role | [ SdkRole](#sdkrole) | Role requested |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkRoleUpdateRequest {#sdkroleupdaterequest}
Defines a request to update an existing role


| Field | Type | Description |
| ----- | ---- | ----------- |
| role | [ SdkRole](#sdkrole) | New role update |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkRoleUpdateResponse {#sdkroleupdateresponse}
Response contains information about the updated role


| Field | Type | Description |
| ----- | ---- | ----------- |
| role | [ SdkRole](#sdkrole) | Role updated |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkRule {#sdkrule}
SdkRule is the message used to construct custom roles in the OpenStorage SDK.

### Format
The following shows the supported format for SdkRule:

* Services: Is the gRPC service name in `OpenStorage<service name>` in lowercase
* Apis: Is the API name in the service in lowercase

Values can also be set to `*`, or start or end with `*` to allow multiple matches in services or apis.

Services and APIs can also be denied by prefixing the value with a `!`. Note that on rule conflicts,
denial will always be chosen.

### Examples

* Allow any call:

```yaml
SdkRule:
  - Services: ["*"]
    Apis: ["*"]
```

* Allow only cluster operations:

```yaml
SdkRule:
  - services: ["cluster"]
    apis: ["*"]
```

* Allow inspection of any object and listings of only volumes

```yaml
SdkRule:
  - Services: ["volumes"]
    Apis: ["*enumerate*"]
  - Services: ["*"]
    Apis: ["inspect*"]
```

* Allow all volume call except create

```yaml
SdkRule:
  - Services: ["volumes"]
    Apis: ["*", "!create"]
```


| Field | Type | Description |
| ----- | ---- | ----------- |
| services | [repeated string](#string) | The gRPC service name in `OpenStorage<service name>` in lowercase |
| apis | [repeated string](#string) | The API name in the service in lowercase |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkSchedulePolicy {#sdkschedulepolicy}
Defines a schedule policy


| Field | Type | Description |
| ----- | ---- | ----------- |
| name | [ string](#string) | Name of the schedule policy |
| schedules | [repeated SdkSchedulePolicyInterval](#sdkschedulepolicyinterval) | Schedule policies |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkSchedulePolicyCreateRequest {#sdkschedulepolicycreaterequest}
Define a schedule policy request


| Field | Type | Description |
| ----- | ---- | ----------- |
| schedule_policy | [ SdkSchedulePolicy](#sdkschedulepolicy) | Schedule Policy |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkSchedulePolicyCreateResponse {#sdkschedulepolicycreateresponse}
Empty response

 <!-- end HasFields -->


## SdkSchedulePolicyDeleteRequest {#sdkschedulepolicydeleterequest}
Define schedule policy deletion request


| Field | Type | Description |
| ----- | ---- | ----------- |
| name | [ string](#string) | Name of the schedule policy |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkSchedulePolicyDeleteResponse {#sdkschedulepolicydeleteresponse}
Empty response

 <!-- end HasFields -->


## SdkSchedulePolicyEnumerateRequest {#sdkschedulepolicyenumeraterequest}
Empty request

 <!-- end HasFields -->


## SdkSchedulePolicyEnumerateResponse {#sdkschedulepolicyenumerateresponse}
Defines a schedule policy enumerate response


| Field | Type | Description |
| ----- | ---- | ----------- |
| policies | [repeated SdkSchedulePolicy](#sdkschedulepolicy) | List of Schedule Policy |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkSchedulePolicyInspectRequest {#sdkschedulepolicyinspectrequest}
Define a schedule policy inspection request


| Field | Type | Description |
| ----- | ---- | ----------- |
| name | [ string](#string) | Name of the schedule Policy |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkSchedulePolicyInspectResponse {#sdkschedulepolicyinspectresponse}
Defines a schedule policy inspection response


| Field | Type | Description |
| ----- | ---- | ----------- |
| policy | [ SdkSchedulePolicy](#sdkschedulepolicy) | List of Schedule Policy |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkSchedulePolicyInterval {#sdkschedulepolicyinterval}
Defines a schedule policy interval


| Field | Type | Description |
| ----- | ---- | ----------- |
| retain | [ int64](#int64) | Number of instances to retain |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) period_type.daily | [ SdkSchedulePolicyIntervalDaily](#sdkschedulepolicyintervaldaily) | Daily policy |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) period_type.weekly | [ SdkSchedulePolicyIntervalWeekly](#sdkschedulepolicyintervalweekly) | Weekly policy |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) period_type.monthly | [ SdkSchedulePolicyIntervalMonthly](#sdkschedulepolicyintervalmonthly) | Monthly policy |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) period_type.periodic | [ SdkSchedulePolicyIntervalPeriodic](#sdkschedulepolicyintervalperiodic) | Periodic policy |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkSchedulePolicyIntervalDaily {#sdkschedulepolicyintervaldaily}
Defines a daily schedule


| Field | Type | Description |
| ----- | ---- | ----------- |
| hour | [ int32](#int32) | Range: 0-23 |
| minute | [ int32](#int32) | Range: 0-59 |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkSchedulePolicyIntervalMonthly {#sdkschedulepolicyintervalmonthly}
Defines a monthly schedule


| Field | Type | Description |
| ----- | ---- | ----------- |
| day | [ int32](#int32) | Range: 1-28 |
| hour | [ int32](#int32) | Range: 0-59 |
| minute | [ int32](#int32) | Range: 0-59 |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkSchedulePolicyIntervalPeriodic {#sdkschedulepolicyintervalperiodic}
Defines a periodic schedule


| Field | Type | Description |
| ----- | ---- | ----------- |
| seconds | [ int64](#int64) | Specify the number of seconds between intervals |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkSchedulePolicyIntervalWeekly {#sdkschedulepolicyintervalweekly}
Defines a weekly schedule


| Field | Type | Description |
| ----- | ---- | ----------- |
| day | [ SdkTimeWeekday](#sdktimeweekday) | none |
| hour | [ int32](#int32) | Range: 0-23 |
| minute | [ int32](#int32) | Range: 0-59 |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkSchedulePolicyUpdateRequest {#sdkschedulepolicyupdaterequest}
Define a request to update a schedule policy


| Field | Type | Description |
| ----- | ---- | ----------- |
| schedule_policy | [ SdkSchedulePolicy](#sdkschedulepolicy) | Schedule Policy |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkSchedulePolicyUpdateResponse {#sdkschedulepolicyupdateresponse}
Empty response

 <!-- end HasFields -->


## SdkServiceCapability {#sdkservicecapability}
Defines a capability of he cluster


| Field | Type | Description |
| ----- | ---- | ----------- |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) type.service | [ SdkServiceCapability.OpenStorageService](#sdkservicecapabilityopenstorageservice) | service type supported by this cluster |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkServiceCapability.OpenStorageService {#sdkservicecapabilityopenstorageservice}



| Field | Type | Description |
| ----- | ---- | ----------- |
| type | [ SdkServiceCapability.OpenStorageService.Type](#sdkservicecapabilityopenstorageservicetype) | Type of service supported |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkStoragePolicy {#sdkstoragepolicy}
A SdkStoragePolicy represents minimum set of volume specs to be
follow while creating volumes.
If storage policy is set default in OpenStoragePolicy service, VolumeSpecPolicy will be
used before creating volume to validate volume specs or ensure minimum volume creation
rules followed


| Field | Type | Description |
| ----- | ---- | ----------- |
| name | [ string](#string) | Name of storage policy. |
| policy | [ VolumeSpecPolicy](#volumespecpolicy) | VolumeSpecs to apply while creating volume. |
| force | [ bool](#bool) | Force if set to true volume specs will be overwritten, otherwise volume creation will fail if the volume specifications are not inline with storage policy |
| allow_update | [ bool](#bool) | If set a volume can be updated without storage Policy restriction, otherwise volume update will be followed as per storage policy specification |
| ownership | [ Ownership](#ownership) | Owner info of storage policy |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkStoragePool {#sdkstoragepool}


 <!-- end HasFields -->


## SdkStoragePoolResizeRequest {#sdkstoragepoolresizerequest}
Defines a request when inspect a storage pool


| Field | Type | Description |
| ----- | ---- | ----------- |
| uuid | [ string](#string) | UUID of the storage pool to inspect |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) resize_factor.size | [ uint64](#uint64) | Size is the new desired size of the storage pool |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) resize_factor.percentage | [ uint64](#uint64) | Size is the new desired size of the storage pool |
| operation_type | [ SdkStoragePool.ResizeOperationType](#sdkstoragepoolresizeoperationtype) | OperationType is the operation that's used to resize the storage pool (optional) |
| skip_wait_for_clean_volumes | [ bool](#bool) | SkipWaitForCleanVolumes would skip the wait for all volumes on the pool to be clean before doing a resize |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkStoragePoolResizeResponse {#sdkstoragepoolresizeresponse}
Defines a response when resizing a storage pool

 <!-- end HasFields -->


## SdkStorageRebalanceRequest {#sdkstoragerebalancerequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| trigger_thresholds | [repeated StorageRebalanceTriggerThreshold](#storagerebalancetriggerthreshold) | TriggerThresholds defines thresholds that would trigger rebalance. For example, TriggerThreshold{ThresholdTypeAbsolutePercent, MetricTypeUsedSpace, 75, 10} would trigger rebalance on pools where used space is more than 75% or less than 10%. Similarly, TriggerThreshold{ThresholdTypeDeltaMeanPercent, MetricTypeUsedSpace, 15, 25} will trigger rebalance for pools where used space is more than 15% from the mean percent for used space for the entire cluster or less than 25% from the mean percent for used space for the entire cluster. |
| trial_run | [ bool](#bool) | TrialRun if true the job only produces steps that would be taken without making any changes |
| source_pool_selector | [repeated LabelSelectorRequirement](#labelselectorrequirement) | SourcePoolSelector allows selecting pools to which trigger thresholds will apply as source |
| target_pool_selector | [repeated LabelSelectorRequirement](#labelselectorrequirement) | TargetPoolSelector allows selecting pools to which trigger thresholds will apply as target |
| max_duration_minutes | [ uint64](#uint64) | MaxDurationMinutes defines how long operation should run when started at schedule. 0 values means no limit on duration |
| remove_repl_1_snapshots | [ bool](#bool) | RemoveRepl1Snapshots if true will instruct rebalance job to remove repl-1 snapshots |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkStorageRebalanceResponse {#sdkstoragerebalanceresponse}
SdkStorageRebalanceResponse is the response to a storage rebalance request


| Field | Type | Description |
| ----- | ---- | ----------- |
| job | [ StorageRebalanceJob](#storagerebalancejob) | Job for this rebalance |
| summary | [ StorageRebalanceSummary](#storagerebalancesummary) | Summary summarizes the rebalance job |
| actions | [repeated StorageRebalanceAudit](#storagerebalanceaudit) | Actions describe all the actions taken during this rebalance |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkUpdateJobRequest {#sdkupdatejobrequest}
Defines a request to update an existing job


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ string](#string) | ID of the job |
| type | [ Job.Type](#jobtype) | Type of the job |
| state | [ Job.State](#jobstate) | State is the new task state to update the job to |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkUpdateJobResponse {#sdkupdatejobresponse}
Defines the response for an update to an existing job

 <!-- end HasFields -->


## SdkUpdateRebalanceJobRequest {#sdkupdaterebalancejobrequest}



| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ string](#string) | ID of the rebalance job |
| state | [ StorageRebalanceJobState](#storagerebalancejobstate) | State is the new task state to update the job to |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkUpdateRebalanceJobResponse {#sdkupdaterebalancejobresponse}


 <!-- end HasFields -->


## SdkVersion {#sdkversion}
SDK version in Major.Minor.Patch format. The goal of this
message is to provide clients a method to determine the SDK
version run by an SDK server.


| Field | Type | Description |
| ----- | ---- | ----------- |
| major | [ int32](#int32) | SDK version major number |
| minor | [ int32](#int32) | SDK version minor number |
| patch | [ int32](#int32) | SDK version patch number |
| version | [ string](#string) | String representation of the SDK version. Must be in `major.minor.patch` format. |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeAttachOptions {#sdkvolumeattachoptions}
Options to attach device


| Field | Type | Description |
| ----- | ---- | ----------- |
| secret_name | [ string](#string) | Indicates the name of the secret stored in a secret store In case of Hashicorp's Vault, it will be the key from the key-value pair stored in its kv backend. In case of Kubernetes secret, it is the name of the secret object itself |
| secret_key | [ string](#string) | In case of Kubernetes, this will be the key stored in the Kubernetes secret |
| secret_context | [ string](#string) | It indicates the additional context which could be used to retrieve the secret. In case of Kubernetes, this is the namespace in which the secret is created. |
| fastpath | [ string](#string) | Indicates whether fastpath needs to be enabled during attach |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeAttachRequest {#sdkvolumeattachrequest}
Defines a request to attach a volume to the node receiving this request


| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_id | [ string](#string) | Id of volume |
| options | [ SdkVolumeAttachOptions](#sdkvolumeattachoptions) | Options to attach device |
| driver_options | [map SdkVolumeAttachRequest.DriverOptionsEntry](#sdkvolumeattachrequestdriveroptionsentry) | The following options are private to the driver plugin running the OpenStorage SDK. Contact your driver developer for any special values that need to be provided here. |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeAttachRequest.DriverOptionsEntry {#sdkvolumeattachrequestdriveroptionsentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeAttachResponse {#sdkvolumeattachresponse}
Defines a response from the node which received the request to attach


| Field | Type | Description |
| ----- | ---- | ----------- |
| device_path | [ string](#string) | Device path where device is exported |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeCapacityUsageRequest {#sdkvolumecapacityusagerequest}
Defines request to retrieve volume/snapshot capacity usage details


| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_id | [ string](#string) | Id of the snapshot/volume to get capacity usage details |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeCapacityUsageResponse {#sdkvolumecapacityusageresponse}
Defines response containing volume/snapshot capacity usage details


| Field | Type | Description |
| ----- | ---- | ----------- |
| capacity_usage_info | [ CapacityUsageInfo](#capacityusageinfo) | CapacityUsage details |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeCatalogRequest {#sdkvolumecatalogrequest}
Request message to get the volume catalog


| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_id | [ string](#string) | VolumeId of the volume that is getting it's catalog retrieved. |
| path | [ string](#string) | Path which will be used as root (default is the actual root) |
| depth | [ string](#string) | Depth of folders/files retrieved (default is all of it, 1 would only return 1 layer) |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeCatalogResponse {#sdkvolumecatalogresponse}
Response message to get volume catalog


| Field | Type | Description |
| ----- | ---- | ----------- |
| catalog | [ CatalogResponse](#catalogresponse) | Catalog |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeCloneRequest {#sdkvolumeclonerequest}
Defines a request to clone a volume or create a volume from a snapshot


| Field | Type | Description |
| ----- | ---- | ----------- |
| name | [ string](#string) | Unique name of the volume. This will be used for idempotency. |
| parent_id | [ string](#string) | Parent volume id or snapshot id will create a new volume as a clone of the parent. |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeCloneResponse {#sdkvolumecloneresponse}
Defines the response when creating a clone from a volume or a snapshot


| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_id | [ string](#string) | Id of new volume |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeCreateRequest {#sdkvolumecreaterequest}
Defines a request to create a volume. Use OpenStorageVolume.Update()
to update any labels on the volume.


| Field | Type | Description |
| ----- | ---- | ----------- |
| name | [ string](#string) | Unique name of the volume. This will be used for idempotency. |
| spec | [ VolumeSpec](#volumespec) | Volume specification |
| labels | [map SdkVolumeCreateRequest.LabelsEntry](#sdkvolumecreaterequestlabelsentry) | Labels to apply to the volume |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeCreateRequest.LabelsEntry {#sdkvolumecreaterequestlabelsentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeCreateResponse {#sdkvolumecreateresponse}
Defines a response to the creation of a volume


| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_id | [ string](#string) | Id of new volume |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeDeleteRequest {#sdkvolumedeleterequest}
Defines the request to delete a volume


| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_id | [ string](#string) | Id of volume to delete |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeDeleteResponse {#sdkvolumedeleteresponse}
Empty response

 <!-- end HasFields -->


## SdkVolumeDetachOptions {#sdkvolumedetachoptions}



| Field | Type | Description |
| ----- | ---- | ----------- |
| force | [ bool](#bool) | Forcefully detach device from the kernel |
| unmount_before_detach | [ bool](#bool) | Unmount the volume before detaching |
| redirect | [ bool](#bool) | redirect the request to the attached node |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeDetachRequest {#sdkvolumedetachrequest}
Defines a request to detach a volume


| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_id | [ string](#string) | Id of the volume |
| options | [ SdkVolumeDetachOptions](#sdkvolumedetachoptions) | Options to detach device |
| driver_options | [map SdkVolumeDetachRequest.DriverOptionsEntry](#sdkvolumedetachrequestdriveroptionsentry) | The following options are private to the driver plugin running the OpenStorage SDK. Contact your driver developer for any special values that need to be provided here. |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeDetachRequest.DriverOptionsEntry {#sdkvolumedetachrequestdriveroptionsentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeDetachResponse {#sdkvolumedetachresponse}
Empty response

 <!-- end HasFields -->


## SdkVolumeEnumerateRequest {#sdkvolumeenumeraterequest}
Defines a request to list volumes

 <!-- end HasFields -->


## SdkVolumeEnumerateResponse {#sdkvolumeenumerateresponse}
Defines the response when listing volumes


| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_ids | [repeated string](#string) | List of volumes matching label |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeEnumerateWithFiltersRequest {#sdkvolumeenumeratewithfiltersrequest}
Defines a request to list volumes


| Field | Type | Description |
| ----- | ---- | ----------- |
| name | [ string](#string) | (optional) Name to search |
| labels | [map SdkVolumeEnumerateWithFiltersRequest.LabelsEntry](#sdkvolumeenumeratewithfiltersrequestlabelsentry) | (optional) Labels to search |
| ownership | [ Ownership](#ownership) | (optional) Ownership to match |
| group | [ Group](#group) | (optional) Group to match |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeEnumerateWithFiltersRequest.LabelsEntry {#sdkvolumeenumeratewithfiltersrequestlabelsentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeEnumerateWithFiltersResponse {#sdkvolumeenumeratewithfiltersresponse}
Defines the response when listing volumes


| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_ids | [repeated string](#string) | List of volumes matching label |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeInspectRequest {#sdkvolumeinspectrequest}
Defines the request to inspect a volume


| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_id | [ string](#string) | Id of volume to inspect |
| options | [ VolumeInspectOptions](#volumeinspectoptions) | Options during inspection |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeInspectResponse {#sdkvolumeinspectresponse}
Defines the response when inspecting a volume


| Field | Type | Description |
| ----- | ---- | ----------- |
| volume | [ Volume](#volume) | Information about the volume |
| name | [ string](#string) | Name of volume |
| labels | [map SdkVolumeInspectResponse.LabelsEntry](#sdkvolumeinspectresponselabelsentry) | Volume labels |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeInspectResponse.LabelsEntry {#sdkvolumeinspectresponselabelsentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeInspectWithFiltersRequest {#sdkvolumeinspectwithfiltersrequest}
Defines the request to inspect volumes using a filter


| Field | Type | Description |
| ----- | ---- | ----------- |
| name | [ string](#string) | (optional) Name to search |
| labels | [map SdkVolumeInspectWithFiltersRequest.LabelsEntry](#sdkvolumeinspectwithfiltersrequestlabelsentry) | (optional) Labels to search |
| ownership | [ Ownership](#ownership) | (optional) Ownership to match |
| group | [ Group](#group) | (optional) Group to match |
| options | [ VolumeInspectOptions](#volumeinspectoptions) | Options during inspection |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeInspectWithFiltersRequest.LabelsEntry {#sdkvolumeinspectwithfiltersrequestlabelsentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeInspectWithFiltersResponse {#sdkvolumeinspectwithfiltersresponse}
Defines the response when inspecting volumes using a filter


| Field | Type | Description |
| ----- | ---- | ----------- |
| volumes | [repeated SdkVolumeInspectResponse](#sdkvolumeinspectresponse) | List of `SdkVolumeInspectResponse` objects describing the volumes |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeMountRequest {#sdkvolumemountrequest}
Defines a request to mount a volume to the node receiving this request


| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_id | [ string](#string) | Id of the volume |
| mount_path | [ string](#string) | Mount path for mounting the volume. |
| options | [ SdkVolumeAttachOptions](#sdkvolumeattachoptions) | Options to attach device |
| driver_options | [map SdkVolumeMountRequest.DriverOptionsEntry](#sdkvolumemountrequestdriveroptionsentry) | The following options are private to the driver plugin running the OpenStorage SDK. Contact your driver developer for any special values that need to be provided here. |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeMountRequest.DriverOptionsEntry {#sdkvolumemountrequestdriveroptionsentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeMountResponse {#sdkvolumemountresponse}
Empty response

 <!-- end HasFields -->


## SdkVolumeSnapshotCreateRequest {#sdkvolumesnapshotcreaterequest}
Defines the request when creating a snapshot from a volume.


| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_id | [ string](#string) | Id of volume to take the snapshot from |
| name | [ string](#string) | Name of the snapshot. |
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
Defines a response after creating a snapshot of a volume


| Field | Type | Description |
| ----- | ---- | ----------- |
| snapshot_id | [ string](#string) | Id of immutable snapshot |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeSnapshotEnumerateRequest {#sdkvolumesnapshotenumeraterequest}
Defines a request to list the snaphots


| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_id | [ string](#string) | Get the snapshots for this volume id |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeSnapshotEnumerateResponse {#sdkvolumesnapshotenumerateresponse}
Defines a response when listing snapshots


| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_snapshot_ids | [repeated string](#string) | List of immutable snapshots |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeSnapshotEnumerateWithFiltersRequest {#sdkvolumesnapshotenumeratewithfiltersrequest}
Defines a request to list the snaphots


| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_id | [ string](#string) | (optional) Get the snapshots for this volume id |
| labels | [map SdkVolumeSnapshotEnumerateWithFiltersRequest.LabelsEntry](#sdkvolumesnapshotenumeratewithfiltersrequestlabelsentry) | (optional) Get snapshots that match these labels |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeSnapshotEnumerateWithFiltersRequest.LabelsEntry {#sdkvolumesnapshotenumeratewithfiltersrequestlabelsentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeSnapshotEnumerateWithFiltersResponse {#sdkvolumesnapshotenumeratewithfiltersresponse}
Defines a response when listing snapshots


| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_snapshot_ids | [repeated string](#string) | List of immutable snapshots |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeSnapshotRestoreRequest {#sdkvolumesnapshotrestorerequest}
Defines a request to restore a volume to a snapshot


| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_id | [ string](#string) | Id of volume |
| snapshot_id | [ string](#string) | Snapshot id to apply to `volume_id` |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeSnapshotRestoreResponse {#sdkvolumesnapshotrestoreresponse}
Empty response

 <!-- end HasFields -->


## SdkVolumeSnapshotScheduleUpdateRequest {#sdkvolumesnapshotscheduleupdaterequest}
Defines a request to update the snapshot schedule of a volume


| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_id | [ string](#string) | Id of volume |
| snapshot_schedule_names | [repeated string](#string) | Names of schedule policies |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeSnapshotScheduleUpdateResponse {#sdkvolumesnapshotscheduleupdateresponse}
Empty response

 <!-- end HasFields -->


## SdkVolumeStatsRequest {#sdkvolumestatsrequest}
Defines a request to retrieve volume statistics


| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_id | [ string](#string) | Id of the volume to get statistics |
| not_cumulative | [ bool](#bool) | When set to false the stats are in /proc/diskstats style stats. When set to true the stats are stats for a specific duration. |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeStatsResponse {#sdkvolumestatsresponse}
Defines a response containing drive statistics


| Field | Type | Description |
| ----- | ---- | ----------- |
| stats | [ Stats](#stats) | Statistics for a single volume |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeUnmountOptions {#sdkvolumeunmountoptions}
Options to unmount device


| Field | Type | Description |
| ----- | ---- | ----------- |
| delete_mount_path | [ bool](#bool) | Delete the mount path on the node after unmounting |
| no_delay_before_deleting_mount_path | [ bool](#bool) | Do not wait for a delay before deleting path. Normally a storage driver may delay before deleting the mount path, which may be necessary to reduce the risk of race conditions. This choice will remove that delay. This value is only usable when `delete_mount_path` is set. |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeUnmountRequest {#sdkvolumeunmountrequest}
Defines a request to unmount a volume on the node receiving this request


| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_id | [ string](#string) | Id of volume |
| mount_path | [ string](#string) | MountPath for device |
| options | [ SdkVolumeUnmountOptions](#sdkvolumeunmountoptions) | Options to unmount device |
| driver_options | [map SdkVolumeUnmountRequest.DriverOptionsEntry](#sdkvolumeunmountrequestdriveroptionsentry) | The following options are private to the driver plugin running the OpenStorage SDK. Contact your driver developer for any special values that need to be provided here. |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeUnmountRequest.DriverOptionsEntry {#sdkvolumeunmountrequestdriveroptionsentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeUnmountResponse {#sdkvolumeunmountresponse}
Empty response

 <!-- end HasFields -->


## SdkVolumeUpdateRequest {#sdkvolumeupdaterequest}
This request is used to adjust or set new values in the volume


| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_id | [ string](#string) | Id of the volume to update |
| labels | [map SdkVolumeUpdateRequest.LabelsEntry](#sdkvolumeupdaterequestlabelsentry) | Change label values. Some of these values may not be able to be changed. New labels will be added to the current volume labels. To delete a label, set the value of the label to an empty string. |
| spec | [ VolumeSpecUpdate](#volumespecupdate) | VolumeSpecUpdate provides a method to request that certain values in the VolumeSpec are changed. This is necessary as a separate variable because values like int and bool in the VolumeSpec cannot be determined if they are being requested to change in gRPC proto3. Some of these values may not be able to be changed.

Here are a few examples of actions that can be accomplished using the VolumeSpec:

* To resize the volume: Set a new value in spec.size_opt.size. * To change number of replicas: Adjust spec.ha_level_opt.ha_level. * To change the I/O Profile: Adjust spec.io_profile_opt.io_profile. |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeUpdateRequest.LabelsEntry {#sdkvolumeupdaterequestlabelsentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## SdkVolumeUpdateResponse {#sdkvolumeupdateresponse}
Empty response

 <!-- end HasFields -->


## Sharedv4FailoverStrategy {#sharedv4failoverstrategy}
Sharedv4FailoverStrategy specifies how long to wait before failing over to a new server.

 <!-- end HasFields -->


## Sharedv4ServiceSpec {#sharedv4servicespec}
Sharedv4ServiceSpec when set, creates a service endpoint for accessing
a sharedv4 volume.


| Field | Type | Description |
| ----- | ---- | ----------- |
| name | [ string](#string) | Name of the service. If not provided the name of the volume will be used for setting up a service |
| type | [ Sharedv4ServiceSpec.ServiceType](#sharedv4servicespecservicetype) | Indicates what kind of service would be created for this volume. |
| external_access | [ bool](#bool) | Indicates whether the service needs to be accessed outside of the cluster |
 <!-- end Fields -->
 <!-- end HasFields -->


## Sharedv4Spec {#sharedv4spec}
Sharedv4Spec specifies common properties of sharedv4 and sharedv4 service volumes


| Field | Type | Description |
| ----- | ---- | ----------- |
| failover_strategy | [ Sharedv4FailoverStrategy.Value](#sharedv4failoverstrategyvalue) | Indicates how aggressively to fail over to a new server. |
 <!-- end Fields -->
 <!-- end HasFields -->


## SnapCreateRequest {#snapcreaterequest}
SnapCreateRequest specifies a request to create a snapshot of given volume.


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ string](#string) | volume id |
| locator | [ VolumeLocator](#volumelocator) | none |
| readonly | [ bool](#bool) | none |
| no_retry | [ bool](#bool) | NoRetry indicates not to retry snapshot creation in the background. |
 <!-- end Fields -->
 <!-- end HasFields -->


## SnapCreateResponse {#snapcreateresponse}
SnapCreateRequest specifies a response that get's returned when creating a snapshot.


| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_create_response | [ VolumeCreateResponse](#volumecreateresponse) | VolumeCreateResponse

in: body Required: true |
 <!-- end Fields -->
 <!-- end HasFields -->


## Source {#source}
Source is a structure that can be given to a volume
to seed the volume with data.


| Field | Type | Description |
| ----- | ---- | ----------- |
| parent | [ string](#string) | A volume id, if specified will create a clone of the parent. |
| seed | [ string](#string) | Seed will seed the volume from the specified URI Any additional config for the source comes from the labels in the spec |
 <!-- end Fields -->
 <!-- end HasFields -->


## Stats {#stats}
Stats is a structure that represents last collected stats for a volume


| Field | Type | Description |
| ----- | ---- | ----------- |
| reads | [ uint64](#uint64) | Reads completed successfully |
| read_ms | [ uint64](#uint64) | Time spent in reads in ms |
| read_bytes | [ uint64](#uint64) | Number of bytes read |
| writes | [ uint64](#uint64) | Writes completed successfully |
| write_ms | [ uint64](#uint64) | Time spent in writes in ms |
| write_bytes | [ uint64](#uint64) | Number of bytes written |
| io_progress | [ uint64](#uint64) | IOs curently in progress |
| io_ms | [ uint64](#uint64) | Time spent doing IOs ms |
| bytes_used | [ uint64](#uint64) | BytesUsed |
| interval_ms | [ uint64](#uint64) | Interval in ms during which stats were collected |
| discards | [ uint64](#uint64) | Discards completed successfully |
| discard_ms | [ uint64](#uint64) | Time spent in discards in ms |
| discard_bytes | [ uint64](#uint64) | Number of bytes discarded |
| unique_blocks | [ uint64](#uint64) | Unique Blocks |
 <!-- end Fields -->
 <!-- end HasFields -->


## StorageCluster {#storagecluster}
StorageCluster represents the state and information about the cluster


| Field | Type | Description |
| ----- | ---- | ----------- |
| status | [ Status](#status) | Status of the cluster |
| id | [ string](#string) | Id of the cluster |
| name | [ string](#string) | Name of the cluster |
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
| scheduler_node_name | [ string](#string) | SchedulerNodeName is name of the node in scheduler context. It can be empty if unable to get the name from the scheduler. |
| HWType | [ HardwareType](#hardwaretype) | HardwareType is the type of the hardware the node has |
| security_status | [ StorageNode.SecurityStatus](#storagenodesecuritystatus) | Determine if the node is secured |
| scheduler_topology | [ SchedulerTopology](#schedulertopology) | Topology information of the node in scheduler context |
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
| ID | [ int32](#int32) | Deprecated! Use `uuid` instead. ID pool ID |
| Cos | [ CosType](#costype) | Cos reflects the capabilities of this drive pool |
| Medium | [ StorageMedium](#storagemedium) | Medium underlying storage type |
| RaidLevel | [ string](#string) | RaidLevel storage raid level |
| TotalSize | [ uint64](#uint64) | TotalSize of the pool |
| Used | [ uint64](#uint64) | Used size of the pool |
| labels | [map StoragePool.LabelsEntry](#storagepoollabelsentry) | Labels is a list of user defined name-value pairs |
| uuid | [ string](#string) | UUID is the unique identifier for a storage pool |
| last_operation | [ StoragePoolOperation](#storagepooloperation) | LastOperation is the most recent operation being performed on a storage pool |
 <!-- end Fields -->
 <!-- end HasFields -->


## StoragePool.LabelsEntry {#storagepoollabelsentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## StoragePoolOperation {#storagepooloperation}
StoragePoolOperation defines an operation being performed on a storage pool


| Field | Type | Description |
| ----- | ---- | ----------- |
| type | [ SdkStoragePool.OperationType](#sdkstoragepooloperationtype) | Type is the type of the operation |
| msg | [ string](#string) | Msg is a user friendly message for the operation |
| params | [map StoragePoolOperation.ParamsEntry](#storagepooloperationparamsentry) | Params for the parameters for the operation |
| status | [ SdkStoragePool.OperationStatus](#sdkstoragepooloperationstatus) | Status is the status of the operation |
 <!-- end Fields -->
 <!-- end HasFields -->


## StoragePoolOperation.ParamsEntry {#storagepooloperationparamsentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## StorageRebalanceAudit {#storagerebalanceaudit}
StorageRebalanceAudit describes the action taken during rebalance


| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_id | [ string](#string) | VolumeID is the id of the volume which was rebalanced |
| name | [ string](#string) | Name is the name of the volumes which was rebalanced |
| action | [ StorageRebalanceAudit.StorageRebalanceAction](#storagerebalanceauditstoragerebalanceaction) | Action is the action executed |
| node | [ string](#string) | Node on which this action happened |
| pool | [ string](#string) | Pool on which this action happened |
| start_time | [ google.protobuf.Timestamp](#googleprotobuftimestamp) | StartTime is the time at which action was started |
| end_time | [ google.protobuf.Timestamp](#googleprotobuftimestamp) | EndTime is time time at which action ended |
| work_summary | [repeated StorageRebalanceWorkSummary](#storagerebalanceworksummary) | WorkSummary summarizes the work done |
| replication_set_id | [ uint64](#uint64) | ReplicationSetId is the ID of the replication set |
| state | [ StorageRebalanceJobState](#storagerebalancejobstate) | State is the current state of the rebalance action |
 <!-- end Fields -->
 <!-- end HasFields -->


## StorageRebalanceJob {#storagerebalancejob}
StorageRebalanceJob describes job input and current status


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ string](#string) | ID of the rebalance job |
| status | [ string](#string) | Status describes status of pools after rebalance if rebalance did not finish successfully |
| state | [ StorageRebalanceJobState](#storagerebalancejobstate) | State of the current job |
| parameters | [ SdkStorageRebalanceRequest](#sdkstoragerebalancerequest) | Parameters is the original request params for this rebalance operation |
| create_time | [ google.protobuf.Timestamp](#googleprotobuftimestamp) | CreateTime is the time the job was created |
| last_update_time | [ google.protobuf.Timestamp](#googleprotobuftimestamp) | LastUpdateTime is the time the job was updated |
 <!-- end Fields -->
 <!-- end HasFields -->


## StorageRebalanceSummary {#storagerebalancesummary}
StorageRebalanceSummary describes summary for the job


| Field | Type | Description |
| ----- | ---- | ----------- |
| total_run_time_seconds | [ uint64](#uint64) | TotalRunTimeSeconds is the total time rebalance is running |
| work_summary | [repeated StorageRebalanceWorkSummary](#storagerebalanceworksummary) | WorkSummary summarizes the work done |
 <!-- end Fields -->
 <!-- end HasFields -->


## StorageRebalanceTriggerThreshold {#storagerebalancetriggerthreshold}



| Field | Type | Description |
| ----- | ---- | ----------- |
| type | [ StorageRebalanceTriggerThreshold.Type](#storagerebalancetriggerthresholdtype) | Type defines type of threshold |
| metric | [ StorageRebalanceTriggerThreshold.Metric](#storagerebalancetriggerthresholdmetric) | Metric defines metric for which this threshold applies to. |
| over_load_trigger_threshold | [ uint64](#uint64) | OverLoadTriggerThreshold will select entity which is over this threshold. OverLoadTriggerThreshold threshold selects pools which act as source for reduction of load defined by the metric. |
| under_load_trigger_threshold | [ uint64](#uint64) | UnderLoadTriggerThreshold will select entity which is under this threshold. UnderLoadTriggerThreshold selects pools which act as targets for increasing load defined by metric. |
 <!-- end Fields -->
 <!-- end HasFields -->


## StorageRebalanceWorkSummary {#storagerebalanceworksummary}



| Field | Type | Description |
| ----- | ---- | ----------- |
| type | [ StorageRebalanceWorkSummary.Type](#storagerebalanceworksummarytype) | Type describes the type of summary. |
| done | [ uint64](#uint64) | Done is the amount of bytes/work done |
| pending | [ uint64](#uint64) | Pending is the amount of bytes/work pending. Done + Pending == Total |
 <!-- end Fields -->
 <!-- end HasFields -->


## StorageResource {#storageresource}
StorageResource groups properties of a storage device.


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
| cache | [ bool](#bool) | True if dedicated as cache |
| pool_metadata_dev | [ bool](#bool) | True if the resource is used as thin pool metadata disk |
 <!-- end Fields -->
 <!-- end HasFields -->


## StorageVersion {#storageversion}
Version information about the storage system


| Field | Type | Description |
| ----- | ---- | ----------- |
| driver | [ string](#string) | OpenStorage driver name |
| version | [ string](#string) | Version of the server |
| details | [map StorageVersion.DetailsEntry](#storageversiondetailsentry) | Extra information provided by the storage system |
 <!-- end Fields -->
 <!-- end HasFields -->


## StorageVersion.DetailsEntry {#storageversiondetailsentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## TopologyRequirement {#topologyrequirement}
TopologyRequirement defines the topology requirement for a volume


| Field | Type | Description |
| ----- | ---- | ----------- |
| labels | [map TopologyRequirement.LabelsEntry](#topologyrequirementlabelsentry) | Key-value pairs defining the required topology for a volume |
 <!-- end Fields -->
 <!-- end HasFields -->


## TopologyRequirement.LabelsEntry {#topologyrequirementlabelsentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## Volume {#volume}
Volume represents an abstract storage volume.


| Field | Type | Description |
| ----- | ---- | ----------- |
| id | [ string](#string) | Self referential volume ID. |
| source | [ Source](#source) | Source specified seed data for the volume. |
| group | [ Group](#group) | Group volumes in the same group have the same group id. |
| readonly | [ bool](#bool) | Readonly is true if this volume is to be mounted with readonly access. |
| locator | [ VolumeLocator](#volumelocator) | User specified locator |
| ctime | [ google.protobuf.Timestamp](#googleprotobuftimestamp) | Volume creation time |
| spec | [ VolumeSpec](#volumespec) | User specified VolumeSpec |
| usage | [ uint64](#uint64) | Usage is bytes consumed by this volume. |
| last_scan | [ google.protobuf.Timestamp](#googleprotobuftimestamp) | LastScan is the time when an integrity check was run. |
| format | [ FSType](#fstype) | Format specifies the filesytem for this volume. |
| status | [ VolumeStatus](#volumestatus) | VolumeStatus is the availability status of this volume. |
| state | [ VolumeState](#volumestate) | VolumeState is the current runtime state of this volume. |
| attached_on | [ string](#string) | AttachedOn is the node instance identifier for clustered systems. |
| attached_state | [ AttachState](#attachstate) | AttachState shows whether the device is attached for internal or external use. |
| device_path | [ string](#string) | DevicePath is the device exported by block device implementations. |
| secure_device_path | [ string](#string) | SecureDevicePath is the device path for an encrypted volume. |
| attach_path | [repeated string](#string) | AttachPath is the mounted path in the host namespace. |
| attach_info | [map Volume.AttachInfoEntry](#volumeattachinfoentry) | AttachInfo is a list of name value mappings that provides attach information. |
| replica_sets | [repeated ReplicaSet](#replicaset) | ReplicatSets storage for this volumefor clustered storage arrays. |
| runtime_state | [repeated RuntimeStateMap](#runtimestatemap) | RuntimeState is a lst of name value mapping of driver specific runtime information. |
| error | [ string](#string) | Error is the Last recorded error. |
| volume_consumers | [repeated VolumeConsumer](#volumeconsumer) | VolumeConsumers are entities that consume this volume |
| fs_resize_required | [ bool](#bool) | FsResizeRequired if an FS resize is required on the volume. |
| attach_time | [ google.protobuf.Timestamp](#googleprotobuftimestamp) | AttachTime time this device was last attached externally. |
| detach_time | [ google.protobuf.Timestamp](#googleprotobuftimestamp) | DetachTime time this device was detached. |
| fpConfig | [ FastpathConfig](#fastpathconfig) | Fastpath extensions |
| last_scan_fix | [ google.protobuf.Timestamp](#googleprotobuftimestamp) | LastScanFix is the time when an integrity check fixed errors in filesystem |
| last_scan_status | [ FilesystemHealthStatus](#filesystemhealthstatus) | LastScanStatus is the time when an integrity check fixed errors in filesystem |
| mount_options | [ MountOptions](#mountoptions) | MountOptions are the runtime mount options that will be used while mounting this volume |
| sharedv4_mount_options | [ MountOptions](#mountoptions) | Sharedv4MountOptions are the runtime mount options that will be used while mounting a sharedv4 volume from a node where the volume replica does not exist |
| derived_io_profile | [ IoProfile](#ioprofile) | DerivedIoProfile the IO profile determined from the pattern |
| in_trashcan | [ bool](#bool) | InTrashcan if the volume is in trashcan |
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


| Field | Type | Description |
| ----- | ---- | ----------- |
| locator | [ VolumeLocator](#volumelocator) | User specified volume name and labels |
| source | [ Source](#source) | Source to create volume |
| spec | [ VolumeSpec](#volumespec) | The storage spec for the volume |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeCreateResponse {#volumecreateresponse}
VolumeCreateResponse


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


| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_id | [ string](#string) | none |
| path | [ string](#string) | none |
| storage | [ VolumeSpec](#volumespec) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeInspectOptions {#volumeinspectoptions}
Options used for volume inspection


| Field | Type | Description |
| ----- | ---- | ----------- |
| deep | [ bool](#bool) | Deep inspection is used to collect more information about the volume. Setting this value may delay the request. |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeLocator {#volumelocator}
VolumeLocator is a structure that is attached to a volume
and is used to carry opaque metadata.


| Field | Type | Description |
| ----- | ---- | ----------- |
| name | [ string](#string) | User friendly identifier |
| volume_labels | [map VolumeLocator.VolumeLabelsEntry](#volumelocatorvolumelabelsentry) | A set of name-value pairs that acts as search filters |
| ownership | [ Ownership](#ownership) | Filter with ownership |
| group | [ Group](#group) | Filter by group |
| volume_ids | [repeated string](#string) | Volume Ids to match |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeLocator.VolumeLabelsEntry {#volumelocatorvolumelabelsentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumePlacementSpec {#volumeplacementspec}



| Field | Type | Description |
| ----- | ---- | ----------- |
| weight | [ int64](#int64) | Weight defines the weight of the rule which allows to break the tie with other matching rules. A rule with higher weight wins over a rule with lower weight. (optional) |
| enforcement | [ EnforcementType](#enforcementtype) | Enforcement specifies the rule enforcement policy. Can take values: required or preferred. (optional) |
| topology_key | [ string](#string) | TopologyKey key for the matching all segments of the cluster topology with the same key e.g If the key is failure-domain.beta.kubernetes.io/zone, this should match all nodes with the same value for this key (i.e in the same zone) |
| match_expressions | [repeated LabelSelectorRequirement](#labelselectorrequirement) | MatchExpressions is a list of label selector requirements. The requirements are ANDed. |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumePlacementStrategy {#volumeplacementstrategy}
VolumePlacementStrategy defines a strategy for placing volumes in the cluster which will be a series of rules
All the rules specified will be applied for volume placement.
Rules that have enforcement as "required" are strictly enforced while "preferred" are best effort.
In situations, where 2 or more rules conflict, the weight of the rules will dictate which wins.


| Field | Type | Description |
| ----- | ---- | ----------- |
| replica_affinity | [repeated ReplicaPlacementSpec](#replicaplacementspec) | ReplicaAffinity defines affinity rules between replicas within a volume |
| replica_anti_affinity | [repeated ReplicaPlacementSpec](#replicaplacementspec) | ReplicaAntiAffinity defines anti-affinity rules between replicas within a volume |
| volume_affinity | [repeated VolumePlacementSpec](#volumeplacementspec) | VolumeAffinity defines affinity rules between volumes |
| volume_anti_affinity | [repeated VolumePlacementSpec](#volumeplacementspec) | VolumeAntiAffinity defines anti-affinity rules between volumes |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeResponse {#volumeresponse}
VolumeResponse is a structure that wraps an error.


| Field | Type | Description |
| ----- | ---- | ----------- |
| error | [ string](#string) | Error message

in: body Required: true |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeServiceInstanceResponse {#volumeserviceinstanceresponse}



| Field | Type | Description |
| ----- | ---- | ----------- |
| error | [ string](#string) | Error code |
| status | [map VolumeServiceInstanceResponse.StatusEntry](#volumeserviceinstanceresponsestatusentry) | Status information exposed a map |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeServiceInstanceResponse.StatusEntry {#volumeserviceinstanceresponsestatusentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeServiceRequest {#volumeservicerequest}
VolumeServiceRequest provides details on what volume service command to
perform in background on the volume


| Field | Type | Description |
| ----- | ---- | ----------- |
| srv_cmd | [ string](#string) | User specified volume service command |
| srv_cmd_params | [map VolumeServiceRequest.SrvCmdParamsEntry](#volumeservicerequestsrvcmdparamsentry) | User specified volume service command's params |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeServiceRequest.SrvCmdParamsEntry {#volumeservicerequestsrvcmdparamsentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeServiceResponse {#volumeserviceresponse}
VolumeServiceResponse specifies the response to a Volume Service command
performed on a volumen


| Field | Type | Description |
| ----- | ---- | ----------- |
| vol_srv_rsp_obj_cnt | [ int32](#int32) | Number of VolumeServiceInstanceResponse returned as part of this response structure |
| vol_srv_rsp | [repeated VolumeServiceInstanceResponse](#volumeserviceinstanceresponse) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeSetRequest {#volumesetrequest}
VolumeSet specifies a request to update a volume.


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


| Field | Type | Description |
| ----- | ---- | ----------- |
| volume | [ Volume](#volume) | Volume |
| volume_response | [ VolumeResponse](#volumeresponse) | VolumeResponse |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeSpec {#volumespec}
VolumeSpec has the properties needed to create a volume.


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
| volume_labels | [map VolumeSpec.VolumeLabelsEntry](#volumespecvolumelabelsentry) | (deprecated, do not use) VolumeLabels configuration labels |
| shared | [ bool](#bool) | Shared is true if this volume can be concurrently accessed by multiple users. |
| replica_set | [ ReplicaSet](#replicaset) | ReplicaSet is the desired set of nodes for the volume data. |
| aggregation_level | [ uint32](#uint32) | Aggregation level Specifies the number of parts the volume can be aggregated from. |
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
| queue_depth | [ uint32](#uint32) | QueueDepth defines the desired block device queue depth |
| force_unsupported_fs_type | [ bool](#bool) | Use to force a file system type which is not recommended. The driver may still refuse to use the file system type. |
| nodiscard | [ bool](#bool) | Nodiscard specifies if the volume will be mounted with discard support disabled. i.e. FS will not release allocated blocks back to the backing storage pool. |
| io_strategy | [ IoStrategy](#iostrategy) | IoStrategy preferred strategy for I/O. |
| placement_strategy | [ VolumePlacementStrategy](#volumeplacementstrategy) | PlacementStrategy specifies a spec to indicate where to place the volume. |
| storage_policy | [ string](#string) | StoragePolicy if applied/specified while creating volume |
| ownership | [ Ownership](#ownership) | Ownership |
| export_spec | [ ExportSpec](#exportspec) | ExportSpec defines how the volume should be exported. |
| fp_preference | [ bool](#bool) | fastpath extensions |
| xattr | [ Xattr.Value](#xattrvalue) | Xattr specifies implementation specific volume attributes |
| scan_policy | [ ScanPolicy](#scanpolicy) | ScanPolicy specifies the filesystem check policy |
| mount_options | [ MountOptions](#mountoptions) | MountOptions defines the options that will be used while mounting this volume |
| sharedv4_mount_options | [ MountOptions](#mountoptions) | Sharedv4MountOptions defines the options that will be used while mounting a sharedv4 volume from a node where the volume replica does not exist |
| proxy_write | [ bool](#bool) | Proxy_write if true, per volume proxy write replication enabled |
| proxy_spec | [ ProxySpec](#proxyspec) | ProxySpec indicates that this volume is used for proxying an external data source |
| sharedv4_service_spec | [ Sharedv4ServiceSpec](#sharedv4servicespec) | Sharedv4ServiceSpec specifies a spec for configuring a service for a sharedv4 volume |
| sharedv4_spec | [ Sharedv4Spec](#sharedv4spec) | Sharedv4Spec specifies common properties of sharedv4 and sharedv4 service volumes |
| auto_fstrim | [ bool](#bool) | Autofstrim indicates that fstrim would be run on this volume automatically, without user intervention |
| io_throttle | [ IoThrottle](#iothrottle) | IoThrottle specifies maximum io(iops/bandwidth) this volume is restricted to |
| number_of_chunks | [ uint32](#uint32) | NumberOfChunks indicates how many chunks must be created, 0 and 1 both mean 1 |
| readahead | [ bool](#bool) | Enable readahead for the volume |
| topology_requirement | [ TopologyRequirement](#topologyrequirement) | TopologyRequirement topology requirement for this volume |
| winshare | [ bool](#bool) | winshare is true if this volume can be accessed from windows pods. |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeSpec.VolumeLabelsEntry {#volumespecvolumelabelsentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeSpecPolicy {#volumespecpolicy}
VolumeSpecPolicy provides a method to set volume storage policy


| Field | Type | Description |
| ----- | ---- | ----------- |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) size_opt.size | [ uint64](#uint64) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) ha_level_opt.ha_level | [ int64](#int64) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) cos_opt.cos | [ CosType](#costype) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) io_profile_opt.io_profile | [ IoProfile](#ioprofile) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) dedupe_opt.dedupe | [ bool](#bool) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) snapshot_interval_opt.snapshot_interval | [ uint32](#uint32) | none |
| volume_labels | [map VolumeSpecPolicy.VolumeLabelsEntry](#volumespecpolicyvolumelabelsentry) | VolumeLabels configuration labels |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) shared_opt.shared | [ bool](#bool) | none |
| replica_set | [ ReplicaSet](#replicaset) | ReplicaSet is the desired set of nodes for the volume data. |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) passphrase_opt.passphrase | [ string](#string) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) snapshot_schedule_opt.snapshot_schedule | [ string](#string) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) scale_opt.scale | [ uint32](#uint32) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) sticky_opt.sticky | [ bool](#bool) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) group_opt.group | [ Group](#group) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) journal_opt.journal | [ bool](#bool) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) sharedv4_opt.sharedv4 | [ bool](#bool) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) queue_depth_opt.queue_depth | [ uint32](#uint32) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) encrypted_opt.encrypted | [ bool](#bool) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) aggregation_level_opt.aggregation_level | [ uint32](#uint32) | none |
| size_operator | [ VolumeSpecPolicy.PolicyOp](#volumespecpolicypolicyop) | Operator to check size |
| ha_level_operator | [ VolumeSpecPolicy.PolicyOp](#volumespecpolicypolicyop) | Operator to check ha_level |
| scale_operator | [ VolumeSpecPolicy.PolicyOp](#volumespecpolicypolicyop) | Operator to check scale |
| snapshot_interval_operator | [ VolumeSpecPolicy.PolicyOp](#volumespecpolicypolicyop) | Operator to check snapshot_interval |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) nodiscard_opt.nodiscard | [ bool](#bool) | none |
| io_strategy | [ IoStrategy](#iostrategy) | IoStrategy preferred strategy for I/O. |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) export_spec_opt.export_spec | [ ExportSpec](#exportspec) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) scan_policy_opt.scan_policy | [ ScanPolicy](#scanpolicy) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) mount_opt.mount_opt_spec | [ MountOptions](#mountoptions) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) sharedv4_mount_opt.sharedv4_mount_opt_spec | [ MountOptions](#mountoptions) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) proxy_write_opt.proxy_write | [ bool](#bool) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) proxy_spec_opt.proxy_spec | [ ProxySpec](#proxyspec) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) fastpath_opt.fastpath | [ bool](#bool) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) sharedv4_service_spec_opt.sharedv4_service_spec | [ Sharedv4ServiceSpec](#sharedv4servicespec) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) sharedv4_spec_opt.sharedv4_spec | [ Sharedv4Spec](#sharedv4spec) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) auto_fstrim_opt.auto_fstrim | [ bool](#bool) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) io_throttle_opt.io_throttle | [ IoThrottle](#iothrottle) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) readahead_opt.readahead | [ bool](#bool) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) winshare_opt.winshare | [ bool](#bool) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeSpecPolicy.VolumeLabelsEntry {#volumespecpolicyvolumelabelsentry}



| Field | Type | Description |
| ----- | ---- | ----------- |
| key | [ string](#string) | none |
| value | [ string](#string) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeSpecUpdate {#volumespecupdate}
VolumeSpecUpdate provides a method to set any of the VolumeSpec of an existing volume


| Field | Type | Description |
| ----- | ---- | ----------- |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) size_opt.size | [ uint64](#uint64) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) ha_level_opt.ha_level | [ int64](#int64) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) cos_opt.cos | [ CosType](#costype) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) io_profile_opt.io_profile | [ IoProfile](#ioprofile) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) dedupe_opt.dedupe | [ bool](#bool) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) snapshot_interval_opt.snapshot_interval | [ uint32](#uint32) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) shared_opt.shared | [ bool](#bool) | none |
| replica_set | [ ReplicaSet](#replicaset) | ReplicaSet is the desired set of nodes for the volume data. |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) passphrase_opt.passphrase | [ string](#string) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) snapshot_schedule_opt.snapshot_schedule | [ string](#string) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) scale_opt.scale | [ uint32](#uint32) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) sticky_opt.sticky | [ bool](#bool) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) group_opt.group | [ Group](#group) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) journal_opt.journal | [ bool](#bool) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) sharedv4_opt.sharedv4 | [ bool](#bool) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) queue_depth_opt.queue_depth | [ uint32](#uint32) | none |
| ownership | [ Ownership](#ownership) | Ownership volume information to update. If the value of `owner` in the `ownership` message is an empty string then the value of `owner` in the `VolumeSpec.Ownership.owner` will not be updated. |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) nodiscard_opt.nodiscard | [ bool](#bool) | none |
| io_strategy | [ IoStrategy](#iostrategy) | IoStrategy preferred strategy for I/O. |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) export_spec_opt.export_spec | [ ExportSpec](#exportspec) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) fastpath_opt.fastpath | [ bool](#bool) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) xattr_opt.xattr | [ Xattr.Value](#xattrvalue) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) scan_policy_opt.scan_policy | [ ScanPolicy](#scanpolicy) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) mount_opt.mount_opt_spec | [ MountOptions](#mountoptions) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) sharedv4_mount_opt.sharedv4_mount_opt_spec | [ MountOptions](#mountoptions) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) proxy_write_opt.proxy_write | [ bool](#bool) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) proxy_spec_opt.proxy_spec | [ ProxySpec](#proxyspec) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) sharedv4_service_spec_opt.sharedv4_service_spec | [ Sharedv4ServiceSpec](#sharedv4servicespec) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) sharedv4_spec_opt.sharedv4_spec | [ Sharedv4Spec](#sharedv4spec) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) auto_fstrim_opt.auto_fstrim | [ bool](#bool) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) io_throttle_opt.io_throttle | [ IoThrottle](#iothrottle) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) readahead_opt.readahead | [ bool](#bool) | none |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) winshare_opt.winshare | [ bool](#bool) | none |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeStateAction {#volumestateaction}
VolumeStateAction specifies desired actions.


| Field | Type | Description |
| ----- | ---- | ----------- |
| attach | [ VolumeActionParam](#volumeactionparam) | Attach or Detach volume |
| mount | [ VolumeActionParam](#volumeactionparam) | Mount or unmount volume |
| mount_path | [ string](#string) | MountPath Path where the device is mounted |
| device_path | [ string](#string) | DevicePath Path returned in attach |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeUsage {#volumeusage}
Provides volume's exclusive bytes and its total usage. This cannot be
retrieved individually and is obtained as part node's usage for a given
node.


| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_id | [ string](#string) | id for the volume/snapshot |
| volume_name | [ string](#string) | name of the volume/snapshot |
| pool_uuid | [ string](#string) | uuid of the pool that this volume belongs to |
| exclusive_bytes | [ uint64](#uint64) | size in bytes exclusively used by the volume/snapshot |
| total_bytes | [ uint64](#uint64) | size in bytes by the volume/snapshot |
| local_cloud_snapshot | [ bool](#bool) | set to true if this volume is snapshot created by cloudbackups |
 <!-- end Fields -->
 <!-- end HasFields -->


## VolumeUsageByNode {#volumeusagebynode}
Provides capacity usage of a node in terms of volumes. Returns VolumeUsage for
all the volume/snapshot(s) in the node.


| Field | Type | Description |
| ----- | ---- | ----------- |
| volume_usage | [repeated VolumeUsage](#volumeusage) | VolumeUsage returns list of VolumeUsage for given node |
 <!-- end Fields -->
 <!-- end HasFields -->


## Xattr {#xattr}
Xattr defines implementation specific volume attribute

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




## AnonymousBucketAccessMode {#anonymousbucketaccessmode}


| Name | Number | Description |
| ---- | ------ | ----------- |
| UnknownBucketAccessMode | 0 | none |
| Private | 1 | Default, disallow uncredentialed access to the backend storage. |
| ReadOnly | 2 | Read only, uncredentialed users can call ListBucket and GetObject. |
| WriteOnly | 3 | Write only, uncredentialed users can only call PutObject. |
| ReadWrite | 4 | Read/Write, uncredentialed users can read objects as well as PutObject. |




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
| Canceled | 6 | none |




## ClusterPairMode.Mode {#clusterpairmodemode}


| Name | Number | Description |
| ---- | ------ | ----------- |
| Default | 0 | Default pairing mode |
| DisasterRecovery | 1 | Paired for DisasterRecovery |
| OneTimeMigration | 2 | Paired for one-time migration |




## CosType {#costype}


| Name | Number | Description |
| ---- | ------ | ----------- |
| NONE | 0 | none |
| LOW | 1 | none |
| MEDIUM | 2 | none |
| HIGH | 3 | none |




## DiagsCollectionStatus.State {#diagscollectionstatusstate}
State is an enum for state of diags collection on a given node

| Name | Number | Description |
| ---- | ------ | ----------- |
| UNSPECIFIED | 0 | Unspecified means uninitialized or unknown state |
| PENDING | 1 | Pending indicates the diags collection is pending and hasn't started |
| RUNNING | 2 | Running indicates diags collection is actively running |
| DONE | 3 | Done indicates diags collection has finished |
| FAILED | 4 | Failed indicates diags collection has failed |




## DriverType {#drivertype}


| Name | Number | Description |
| ---- | ------ | ----------- |
| DRIVER_TYPE_NONE | 0 | none |
| DRIVER_TYPE_FILE | 1 | none |
| DRIVER_TYPE_BLOCK | 2 | none |
| DRIVER_TYPE_OBJECT | 3 | none |
| DRIVER_TYPE_CLUSTERED | 4 | none |
| DRIVER_TYPE_GRAPH | 5 | none |




## EnforcementType {#enforcementtype}
Defines the types of enforcement on the given rules

| Name | Number | Description |
| ---- | ------ | ----------- |
| required | 0 | This specifies that the rule is required and must be strictly enforced |
| preferred | 1 | This specifies that the rule is preferred and can be best effort |




## ExportProtocol {#exportprotocol}
ExportProtocol defines how the device is exported..

| Name | Number | Description |
| ---- | ------ | ----------- |
| INVALID | 0 | Invalid uninitialized value |
| PXD | 1 | PXD the volume is exported over Portworx block interface. |
| ISCSI | 2 | ISCSI the volume is exported over ISCSI. |
| NFS | 3 | NFS the volume is exported over NFS. |
| CUSTOM | 4 | Custom the volume is exported over custom interface. |




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
| FS_TYPE_XFSv2 | 8 | none |




## FastpathProtocol {#fastpathprotocol}


| Name | Number | Description |
| ---- | ------ | ----------- |
| FASTPATH_PROTO_UNKNOWN | 0 | none |
| FASTPATH_PROTO_NVMEOF_TCP | 1 | none |
| FASTPATH_PROTO_ISCSI | 2 | none |
| FASTPATH_PROTO_LOCAL | 3 | none |




## FastpathStatus {#fastpathstatus}
fastpath extensions

| Name | Number | Description |
| ---- | ------ | ----------- |
| FASTPATH_UNKNOWN | 0 | backward compatible state |
| FASTPATH_ACTIVE | 1 | fastpath activated and enabled |
| FASTPATH_INACTIVE | 2 | fastpath inactive |
| FASTPATH_UNSUPPORTED | 3 | fastpath unsupported, disabled |
| FASTPATH_PENDING | 4 | fastpath operation setup in progress |
| FASTPATH_ERRORED | 5 | fastpath error'd, possible stuck failure |




## FilesystemCheck.FilesystemCheckStatus {#filesystemcheckfilesystemcheckstatus}
FilesystemChecktatus represents the status codes returned from
OpenStorageFilesystemCheck service APIs()

| Name | Number | Description |
| ---- | ------ | ----------- |
| FS_CHECK_UNKNOWN | 0 | Filesystem Check operation is an unknown state |
| FS_CHECK_NOT_RUNNING | 1 | FilesystemCheck operation not running for the specified volume |
| FS_CHECK_STARTED | 2 | FilesystemCheck operation started for the specified volume |
| FS_CHECK_INPROGRESS | 3 | FilesystemCheck operation is in progress |
| FS_CHECK_STOPPED | 4 | FilesystemCheck operation was stopped by the user |
| FS_CHECK_COMPLETED | 5 | FilesystemCheck operation completed successfully |
| FS_CHECK_FAILED | 6 | FilesystemCheck operation failed due to internal error |




## FilesystemHealthStatus {#filesystemhealthstatus}


| Name | Number | Description |
| ---- | ------ | ----------- |
| FS_HEALTH_STATUS_UNKNOWN | 0 | filesystem health status is unknown |
| FS_HEALTH_STATUS_HEALTHY | 1 | filesystem is a healthy |
| FS_HEALTH_STATUS_SAFE_TO_FIX | 2 | filesystem has errors, but can be fixed safely |
| FS_HEALTH_STATUS_NEEDS_INSPECTION | 3 | filesystem has errors, these cannot be fixed automatically, user needs to review the reported errors and take appropriate action |




## FilesystemTrim.FilesystemTrimStatus {#filesystemtrimfilesystemtrimstatus}
FilesystemTrimStatus represents the status codes returned from
OpenStorageFilesystemTrim service APIs()

| Name | Number | Description |
| ---- | ------ | ----------- |
| FS_TRIM_UNKNOWN | 0 | Filesystem Trim operation is an unknown state |
| FS_TRIM_NOT_RUNNING | 1 | Filesystem Trim operation is not running for the specified volume |
| FS_TRIM_STARTED | 2 | Filesystem Trim operation started for the specified volume |
| FS_TRIM_INPROGRESS | 3 | Filesystem Trim operation is in progress for the specified volume |
| FS_TRIM_STOPPED | 4 | Filesystem Trim operation was stopped by the user for the specified volume |
| FS_TRIM_COMPLETED | 5 | Filesystem Trim operation completed successfully for the specified volume |
| FS_TRIM_FAILED | 6 | Filesystem Trim operation failed due to internal error for the specified volume |




## GraphDriverChangeType {#graphdriverchangetype}


| Name | Number | Description |
| ---- | ------ | ----------- |
| GRAPH_DRIVER_CHANGE_TYPE_NONE | 0 | none |
| GRAPH_DRIVER_CHANGE_TYPE_MODIFIED | 1 | none |
| GRAPH_DRIVER_CHANGE_TYPE_ADDED | 2 | none |
| GRAPH_DRIVER_CHANGE_TYPE_DELETED | 3 | none |




## HardwareType {#hardwaretype}


| Name | Number | Description |
| ---- | ------ | ----------- |
| UnknownMachine | 0 | When we do not know the machine type alsp the default |
| VirtualMachine | 1 | when we are running on virtual machine |
| BareMetalMachine | 2 | when we are running on physical hardware |




## IoProfile {#ioprofile}


| Name | Number | Description |
| ---- | ------ | ----------- |
| IO_PROFILE_SEQUENTIAL | 0 | none |
| IO_PROFILE_RANDOM | 1 | none |
| IO_PROFILE_DB | 2 | none |
| IO_PROFILE_DB_REMOTE | 3 | none |
| IO_PROFILE_CMS | 4 | none |
| IO_PROFILE_SYNC_SHARED | 5 | none |
| IO_PROFILE_AUTO | 6 | none |
| IO_PROFILE_NONE | 7 | none |




## Job.State {#jobstate}
State is an enum for state of a node drain operation

| Name | Number | Description |
| ---- | ------ | ----------- |
| UNSPECIFIED_STATE | 0 | Unspecified |
| PENDING | 1 | Pending indicates job is still pending and has not started work |
| RUNNING | 2 | Running indicates job is actively running |
| DONE | 3 | Done indicates job has finished processing |
| PAUSED | 4 | Paused indicates job is paused |
| CANCELLED | 5 | Cancelled indicates job is cancelled |
| FAILED | 6 | Failed indicates job has failed |




## Job.Type {#jobtype}
Type are the supported job types

| Name | Number | Description |
| ---- | ------ | ----------- |
| UNSPECIFIED_TYPE | 0 | Unspecified |
| NONE | 1 | None |
| DRAIN_ATTACHMENTS | 2 | Job for draining volume attachments |
| CLOUD_DRIVE_TRANSFER | 3 | Job for transferring cloud drives between nodes |
| COLLECT_DIAGS | 4 | Job for collecting diags from the cluster nodes |




## LabelSelectorRequirement.Operator {#labelselectorrequirementoperator}
This defines operator types used in a label matching rule

| Name | Number | Description |
| ---- | ------ | ----------- |
| In | 0 | In means the value for 'key' should be in one of the given value(s) |
| NotIn | 1 | NotIn means the value for 'key' should NOT be in one of the given value(s) |
| Exists | 2 | Exists means the 'key' should just exist regardless of the value |
| DoesNotExist | 3 | DoesNotExist means the 'key' should NOT exist |
| Gt | 4 | Gt means the 'key' should be greater than the value(s) |
| Lt | 5 | Lt means the 'key' should be less than the value(s) |




## OperationFlags {#operationflags}


| Name | Number | Description |
| ---- | ------ | ----------- |
| OP_FLAGS_UNKNOWN | 0 | none |
| OP_FLAGS_NONE | 1 | none |
| OP_FLAGS_DETACH_FORCE | 2 | Perform a force_detach during detach operation |




## Ownership.AccessType {#ownershipaccesstype}
Access types can be set by owner to have different levels of access to
a resource.

It is up to the resource to interpret what the types mean and are
used for.

| Name | Number | Description |
| ---- | ------ | ----------- |
| Read | 0 | Read access only and cannot affect the resource. |
| Write | 1 | Write access and can affect the resource. This type automatically provides Read access also. |
| Admin | 2 | Administrator access. This type automatically provides Read and Write access also. |




## ProxyProtocol {#proxyprotocol}
ProxyProtocol defines the protocol used for proxy.

| Name | Number | Description |
| ---- | ------ | ----------- |
| PROXY_PROTOCOL_INVALID | 0 | Invalid uninitialized value |
| PROXY_PROTOCOL_NFS | 1 | NFS indicates that the external data source is NFS and the volume will be reflected over NFS protocol |
| PROXY_PROTOCOL_S3 | 2 | S3 indicates that the external data source is an object store. |
| PROXY_PROTOCOL_PXD | 3 | PXD indicates that the external data source is a Portworx block volume. |
| PROXY_PROTOCOL_PURE_BLOCK | 4 | PURE_BLOCK indicates that the external data source is a pure_block volume. |
| PROXY_PROTOCOL_PURE_FILE | 5 | PURE_FILE indicates that the external data source is a pure_file volume. |




## ResourceType {#resourcetype}


| Name | Number | Description |
| ---- | ------ | ----------- |
| RESOURCE_TYPE_NONE | 0 | none |
| RESOURCE_TYPE_VOLUME | 1 | none |
| RESOURCE_TYPE_NODE | 2 | none |
| RESOURCE_TYPE_CLUSTER | 3 | none |
| RESOURCE_TYPE_DRIVE | 4 | none |
| RESOURCE_TYPE_POOL | 5 | none |




## RestoreParamBoolType {#restoreparambooltype}


| Name | Number | Description |
| ---- | ------ | ----------- |
| PARAM_BKUPSRC | 0 | Default: whateever was cloudbakup's option for the parameter |
| PARAM_FALSE | 1 | none |
| PARAM_TRUE | 2 | none |




## ScanPolicy.ScanAction {#scanpolicyscanaction}


| Name | Number | Description |
| ---- | ------ | ----------- |
| SCAN_ACTION_NONE | 0 | none |
| SCAN_ACTION_SCAN_ONLY | 1 | none |
| SCAN_ACTION_SCAN_REPAIR | 2 | none |




## ScanPolicy.ScanTrigger {#scanpolicyscantrigger}


| Name | Number | Description |
| ---- | ------ | ----------- |
| SCAN_TRIGGER_NONE | 0 | none |
| SCAN_TRIGGER_ON_MOUNT | 1 | none |
| SCAN_TRIGGER_ON_NEXT_MOUNT | 2 | none |




## SdkCloudBackupClusterType {#sdkcloudbackupclustertype}
CloudBackup operations types

| Name | Number | Description |
| ---- | ------ | ----------- |
| SdkCloudBackupClusterUnknown | 0 | Unknown |
| SdkCloudBackupClusterCurrent | 1 | Beongs to this cluster |
| SdkCloudBackupClusterOther | 2 | not this. other cluster |




## SdkCloudBackupOpType {#sdkcloudbackupoptype}
CloudBackup operations types

| Name | Number | Description |
| ---- | ------ | ----------- |
| SdkCloudBackupOpTypeUnknown | 0 | Unknown |
| SdkCloudBackupOpTypeBackupOp | 1 | Backup |
| SdkCloudBackupOpTypeRestoreOp | 2 | Restore |




## SdkCloudBackupRequestedState {#sdkcloudbackuprequestedstate}
SdkCloudBackupRequestedState defines states to set a specified backup or restore
to or from a cloud provider

| Name | Number | Description |
| ---- | ------ | ----------- |
| SdkCloudBackupRequestedStateUnknown | 0 | Unknown state |
| SdkCloudBackupRequestedStatePause | 1 | Pause the backup or restore |
| SdkCloudBackupRequestedStateResume | 2 | Resume the backup or restore |
| SdkCloudBackupRequestedStateStop | 3 | Stop a backup or restore |




## SdkCloudBackupStatusType {#sdkcloudbackupstatustype}
CloudBackup status types

| Name | Number | Description |
| ---- | ------ | ----------- |
| SdkCloudBackupStatusTypeUnknown | 0 | Unknown |
| SdkCloudBackupStatusTypeNotStarted | 1 | Not started |
| SdkCloudBackupStatusTypeDone | 2 | Done |
| SdkCloudBackupStatusTypeAborted | 3 | Aborted |
| SdkCloudBackupStatusTypePaused | 4 | Paused |
| SdkCloudBackupStatusTypeStopped | 5 | Stopped |
| SdkCloudBackupStatusTypeActive | 6 | Active |
| SdkCloudBackupStatusTypeFailed | 7 | Failed |
| SdkCloudBackupStatusTypeQueued | 8 | Queued |
| SdkCloudBackupStatusTypeInvalid | 9 | Invalid, used by enumerate, includes failed, stopped and aborted |




## SdkServiceCapability.OpenStorageService.Type {#sdkservicecapabilityopenstorageservicetype}


| Name | Number | Description |
| ---- | ------ | ----------- |
| UNKNOWN | 0 | Unknown service |
| CLUSTER | 1 | Cluster management |
| CLOUD_BACKUP | 2 | Cloud backup of volumes management |
| CREDENTIALS | 3 | Credentials management |
| NODE | 4 | Node management |
| OBJECT_STORAGE | 5 | Object Storage management |
| SCHEDULE_POLICY | 6 | Schedule policy management |
| VOLUME | 7 | Volume management |
| ALERTS | 8 | Alert enumeration |
| MOUNT_ATTACH | 9 | Mount/Attach Support |
| ROLE | 10 | Role service |
| CLUSTER_PAIR | 11 | Cluster Pair service |
| MIGRATE | 12 | Migrate service |
| STORAGE_POLICY | 13 | StoragePolicy Service |




## SdkStoragePool.OperationStatus {#sdkstoragepooloperationstatus}
OperationStatus captures the various statuses of a storage pool operation

| Name | Number | Description |
| ---- | ------ | ----------- |
| OPERATION_PENDING | 0 | Operation pending |
| OPERATION_IN_PROGRESS | 1 | Operation is in progress |
| OPERATION_SUCCESSFUL | 2 | Operation is successful |
| OPERATION_FAILED | 3 | Operation failed |




## SdkStoragePool.OperationType {#sdkstoragepooloperationtype}
OperationType defines the various operations that are performed on a storage pool

| Name | Number | Description |
| ---- | ------ | ----------- |
| OPERATION_RESIZE | 0 | Resize operation |




## SdkStoragePool.ResizeOperationType {#sdkstoragepoolresizeoperationtype}
Defines the operation types available to resize a storage pool

| Name | Number | Description |
| ---- | ------ | ----------- |
| RESIZE_TYPE_AUTO | 0 | Automatically pick the optimum resize operation type |
| RESIZE_TYPE_ADD_DISK | 1 | Add a new drive to resize the pool |
| RESIZE_TYPE_RESIZE_DISK | 2 | Resize existing drives to resize the pool |




## SdkTimeWeekday {#sdktimeweekday}
Defines times of day

| Name | Number | Description |
| ---- | ------ | ----------- |
| SdkTimeWeekdaySunday | 0 | Sunday |
| SdkTimeWeekdayMonday | 1 | Monday |
| SdkTimeWeekdayTuesday | 2 | Tuesday |
| SdkTimeWeekdayWednesday | 3 | Wednesday |
| SdkTimeWeekdayThursday | 4 | Thursday |
| SdkTimeWeekdayFriday | 5 | Friday |
| SdkTimeWeekdaySaturday | 6 | Saturday |




## SdkVersion.Version {#sdkversionversion}
These values are constants that can be used by the
client and server applications

| Name | Number | Description |
| ---- | ------ | ----------- |
| MUST_HAVE_ZERO_VALUE | 0 | Must be set in the proto file; ignore. |
| Major | 0 | SDK version major value of this specification |
| Minor | 151 | SDK version minor value of this specification |
| Patch | 0 | SDK version patch value of this specification |




## SeverityType {#severitytype}


| Name | Number | Description |
| ---- | ------ | ----------- |
| SEVERITY_TYPE_NONE | 0 | none |
| SEVERITY_TYPE_ALARM | 1 | none |
| SEVERITY_TYPE_WARNING | 2 | none |
| SEVERITY_TYPE_NOTIFY | 3 | none |




## Sharedv4FailoverStrategy.Value {#sharedv4failoverstrategyvalue}


| Name | Number | Description |
| ---- | ------ | ----------- |
| UNSPECIFIED | 0 | Unspecified |
| AGGRESSIVE | 1 | Fail over to the new NFS server quickly. This is appropriate for the apps with sharedv4 service volumes. It can also be used for the apps with sharedv4 volumes when such apps are able to restart quickly. |
| NORMAL | 2 | Give the old NFS server more time to come back up before triggering the failover. This is appropriate for the apps with sharedv4 volumes (non-service). |




## Sharedv4ServiceSpec.ServiceType {#sharedv4servicespecservicetype}
Type of sharedv4 service. Values are governed by the different types
of services supported by container orchestrators such as Kubernetes.

| Name | Number | Description |
| ---- | ------ | ----------- |
| UNSPECIFIED | 0 | Unspecified |
| NODEPORT | 1 | Export the sharedv4 service on each Node's IP. In this mode the sharedv4 volume can be accessed from outside the cluster using one of the node's IPs. |
| CLUSTERIP | 2 | Export the shared4 service on an internal cluster IP. In this mode the sharedv4 volume will only be accessible within the cluster via this service. |
| LOADBALANCER | 3 | Expose the sharedv4 service on cloud provider's load balancer. Applicable when running in cloud. In this mode the sharedv4 volume can be accessed from outside the cluster. |
| NONE | 4 | Do not use sharedv4 service feature. Use legacy sharedv4 volumes instead. |




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
| STATUS_POOLMAINTENANCE | 14 | none |
| STATUS_MAX | 15 | Add statuses before MAX and update the number for MAX |




## StorageMedium {#storagemedium}


| Name | Number | Description |
| ---- | ------ | ----------- |
| STORAGE_MEDIUM_MAGNETIC | 0 | Magnetic spinning disk. |
| STORAGE_MEDIUM_SSD | 1 | SSD disk |
| STORAGE_MEDIUM_NVME | 2 | NVME disk |




## StorageNode.SecurityStatus {#storagenodesecuritystatus}


| Name | Number | Description |
| ---- | ------ | ----------- |
| UNSPECIFIED | 0 | Security status type is unknown |
| UNSECURED | 1 | Node is unsecure |
| SECURED | 2 | Node is secured with authentication and authorization |
| SECURED_ALLOW_SECURITY_REMOVAL | 3 | Node is secured, but in the process of removing security. This state allows other unsecured nodes to join the cluster since the cluster is in the process of removing security authentication and authorization. |




## StorageRebalanceAudit.StorageRebalanceAction {#storagerebalanceauditstoragerebalanceaction}
StorageRebalanceAction describes type of rebalance action

| Name | Number | Description |
| ---- | ------ | ----------- |
| ADD_REPLICA | 0 | Indicates new replica was added |
| REMOVE_REPLICA | 1 | Indicates existing replica was removed |




## StorageRebalanceJobState {#storagerebalancejobstate}
StorageRebalanceJobState is an enum for state of the current rebalance operation

| Name | Number | Description |
| ---- | ------ | ----------- |
| PENDING | 0 | Pending indicates job is still pending and has not started work |
| RUNNING | 1 | Running indicates job is actively running |
| DONE | 2 | Done indicates job has finished processing |
| PAUSED | 3 | Paused indicates job is paused |
| CANCELLED | 4 | Cancelled indicates job is cancelled |




## StorageRebalanceTriggerThreshold.Metric {#storagerebalancetriggerthresholdmetric}
Metric is an enum that defines the metric to use for rebalance

| Name | Number | Description |
| ---- | ------ | ----------- |
| PROVISION_SPACE | 0 | ProvisionSpace indicates rebalance for provisioned space |
| USED_SPACE | 1 | UsedSpace indicates rebalance for used space |




## StorageRebalanceTriggerThreshold.Type {#storagerebalancetriggerthresholdtype}
Type is an enum that defines the type fo the trigger threshold

| Name | Number | Description |
| ---- | ------ | ----------- |
| ABSOLUTE_PERCENT | 0 | AbsolutePercent indicates absolute percent comparison. Example, 75 % used of capacity, or 50 % provisioned of capacity. |
| DELTA_MEAN_PERCENT | 1 | DeltaMeanPercent indicates mean percent comparison threshold. Example, 10 % more than mean for cluster. |




## StorageRebalanceWorkSummary.Type {#storagerebalanceworksummarytype}
Type is an enum to indicate the type of work summary

| Name | Number | Description |
| ---- | ------ | ----------- |
| UnbalancedPools | 0 | summary for unbalanced pools |
| UnbalancedVolumes | 1 | summary for unbalanced volumes |
| UnbalancedProvisionedSpaceBytes | 2 | summary for unbalanced provisioned space |
| UnbalancedUsedSpaceBytes | 3 | summary for unbalanced used space |




## VolumeActionParam {#volumeactionparam}


| Name | Number | Description |
| ---- | ------ | ----------- |
| VOLUME_ACTION_PARAM_NONE | 0 | none |
| VOLUME_ACTION_PARAM_OFF | 1 | Maps to the boolean value false |
| VOLUME_ACTION_PARAM_ON | 2 | Maps to the boolean value true. |




## VolumeSpecPolicy.PolicyOp {#volumespecpolicypolicyop}
This defines an operator for the policy comparisons

| Name | Number | Description |
| ---- | ------ | ----------- |
| Equal | 0 | Policy will make sure the value must be equal |
| Minimum | 1 | Policy will make sure the requested value must be greater than or equal |
| Maximum | 2 | Policy will make sure the requested value must be less than or equal |




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




## Xattr.Value {#xattrvalue}


| Name | Number | Description |
| ---- | ------ | ----------- |
| UNSPECIFIED | 0 | Value is uninitialized or unknown |
| COW_ON_DEMAND | 1 | Enable on-demand copy-on-write on the volume |


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

