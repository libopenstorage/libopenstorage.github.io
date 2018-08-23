// Example code
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/golang/protobuf/ptypes"
	api "github.com/libopenstorage/openstorage-sdk-clients/sdk/golang"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

type OpenStorageSdkToken struct{}

func (t OpenStorageSdkToken) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	secret := "thisistheadminkey"

	// Create Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		// Set issuer
		"sub": "admin",
		// Set issued at time
		"iat": time.Now().Unix(),
		// Set expiration
		"exp": time.Now().Add(time.Minute * 5).Unix(),
		// Name
		"name": "Luis Pabon",
		// Email
		"email": "luis@portworx.com",
	})

	// Sign the token
	signedtoken, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"authorization": "bearer " + signedtoken,
	}, nil
}

func (t OpenStorageSdkToken) RequireTransportSecurity() bool {
	return false
}

func main() {

	ctx := context.Background()
	token := OpenStorageSdkToken{}

	// Setup connection
	grpccreds, err := credentials.NewClientTLSFromFile("cert.pem", "")
	if err != nil {
		fmt.Printf("Creds failed: %v", err)
		os.Exit(1)
	}

	// Setup a connection
	conn, err := grpc.Dial(":9100", grpc.WithTransportCredentials(grpccreds), grpc.WithPerRPCCredentials(token))
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	// Create a cluster client
	cluster := api.NewOpenStorageClusterClient(conn)

	// Print the cluster information
	clusterInfo, err := cluster.InspectCurrent(
		ctx,
		&api.SdkClusterInspectCurrentRequest{})
	if err != nil {
		gerr, _ := status.FromError(err)
		fmt.Printf("A Error Code[%d] Message[%s]\n",
			gerr.Code(), gerr.Message())
		os.Exit(1)
	}
	fmt.Printf("Connected to Cluster %s\n",
		clusterInfo.GetCluster().GetId())
	fmt.Println()

	// Create a 100Gi volume
	volumes := api.NewOpenStorageVolumeClient(conn)
	v, err := volumes.Create(
		ctx,
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
	fmt.Println()

	// Take a snapshot
	snap, err := volumes.SnapshotCreate(
		ctx,
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
	fmt.Println()

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
	fmt.Println()

	// Create a backup to a cloud provider of our volume
	cloudbackups := api.NewOpenStorageCloudBackupClient(conn)
	_, err = cloudbackups.Create(context.Background(),
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
	fmt.Printf("Backup started for volume %s\n", v.GetVolumeId())

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
		fmt.Printf("Backup status for volume: %s\n"+
			"Type: %s\n"+
			"Status: %s\n"+
			"Full JSON Response: %s\n",
			volID,
			status.GetOptype().String(),
			status.GetStatus().String(),
			string(b))
	}
	fmt.Println()

	// Backup History
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
	fmt.Println()
}
