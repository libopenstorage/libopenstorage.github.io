import grpc
import jwt
import datetime
import hashlib
from openstorage import api_pb2
from openstorage import api_pb2_grpc


claims = {}
claims['sub'] = 'admin'
claims['iat'] = datetime.datetime.utcnow()
claims['exp'] = datetime.datetime.utcnow() + datetime.timedelta(minutes=10)
claims['name'] = 'Luis Pabon'
claims['email'] = 'luis@portworx.com'
secret='ZDBlMTY3NDhiYTlmZGMwYzQyNTQ2Y2I5Mjg4YzkwNjAgIC0K'
token = jwt.encode(claims, secret, algorithm='HS256')
metadata=[('authorization', 'bearer ' + token)]

channel = grpc.insecure_channel('localhost:9100')

try:
    # Cluster connection
    clusters = api_pb2_grpc.OpenStorageClusterStub(channel)
    ic_resp = clusters.InspectCurrent(api_pb2.SdkClusterInspectCurrentRequest(), metadata=metadata)
    print('Connected to {0} with status {1}'.format(
        ic_resp.cluster.id,
        api_pb2.Status.Name(ic_resp.cluster.status)
    ))

    # Create a volume
    volumes = api_pb2_grpc.OpenStorageVolumeStub(channel)
    v_resp = volumes.Create(api_pb2.SdkVolumeCreateRequest(
        name="myvol",
        spec=api_pb2.VolumeSpec(
            size=100*1024*1024*1024,
            ha_level=3,
        )
    ), metadata=metadata)
    print('Volume id is {0}'.format(v_resp.volume_id))

    # Create a snapshot
    snap = volumes.SnapshotCreate(api_pb2.SdkVolumeSnapshotCreateRequest(
        volume_id=v_resp.volume_id,
        name="mysnap"
    ), metadata=metadata)
    print('Snapshot created with id {0}'.format(snap.snapshot_id))

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
    ), metadata=metadata)
    print('Credential id is {0}'.format(cred_resp.credential_id))

    # Create backup
    backups = api_pb2_grpc.OpenStorageCloudBackupStub(channel)
    backup_resp = backups.Create(api_pb2.SdkCloudBackupCreateRequest(
        volume_id=v_resp.volume_id,
        credential_id=cred_resp.credential_id,
        full=False
    ), metadata=metadata)

    # Get status
    status_resp = backups.Status(api_pb2.SdkCloudBackupStatusRequest(
        volume_id=v_resp.volume_id
    ), metadata=metadata)
    backup_status = status_resp.statuses[v_resp.volume_id]
    print('Status of the backup is {0}'.format(
        api_pb2.SdkCloudBackupStatusType.Name(backup_status.status)
    ))

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

except grpc.RpcError as e:
    print('Failed: code={0} msg={1}'.format(e.code(), e.details()))
