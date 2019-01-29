// Example code
package main

import (
	"context"
	"crypto/x509"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/golang/protobuf/ptypes"
	api "github.com/libopenstorage/openstorage-sdk-clients/sdk/golang"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

const (
	Bytes = uint64(1)
	KB    = Bytes * uint64(1024)
	MB    = KB * uint64(1024)
	GB    = MB * uint64(1024)
)

var (
	useTls  = flag.Bool("usetls", false, "Connect to server using TLS. Loads CA from the system")
	token   = flag.String("token", "", "Authorization token if any")
	address = flag.String("address", "127.0.0.1:9100", "Address to server as <address>:<port>")
)

type OpenStorageSdkToken struct{}

func (t OpenStorageSdkToken) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": "bearer " + *token,
	}, nil
}

func (t OpenStorageSdkToken) RequireTransportSecurity() bool {
	return *useTls
}

func main() {
	flag.Parse()

	// There are two ways to setup a token:
	//   - One is to setup a client interceptor which adds the token
	//     to every call automatically using grpc.WithPerRPCCredentials().
	//   - Second way is just to add it to the context directly as follows:
	//   import "google.golang.org/grpc/metadata"
	//   md := metadata.New(map[string]string{
	//		"authorization": "bearer" + token,
	//	 })
	//   ctx := metadata.NewOutgoingContext(context.Background(), md)
	//
	//   We will be using the more complicated first model to show how it can be done
	//
	//   To accomplish this, we first need to create an object that satisfies the
	//   interface needed by grpc.WithPerRPCCredentials(..)
	contextToken := OpenStorageSdkToken{}

	dialOptions := []grpc.DialOption{grpc.WithInsecure()}
	if *useTls {
		// Setup a connection
		capool, err := x509.SystemCertPool()
		if err != nil {
			fmt.Printf("Failed to load system certs: %v\n")
			os.Exit(1)
		}
		dialOptions = []grpc.DialOption{grpc.WithTransportCredentials(
			credentials.NewClientTLSFromCert(capool, ""),
		)}
	}

	if len(*token) != 0 {
		// Add token interceptor
		dialOptions = append(dialOptions, grpc.WithPerRPCCredentials(contextToken))
	}

	conn, err := grpc.Dial(*address, dialOptions...)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

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

	// --- Get Cluster capacity ---
	// First, get all node node IDs in this cluster
	nodeclient := api.NewOpenStorageNodeClient(conn)
	nodeEnumResp, err := nodeclient.Enumerate(
		context.Background(),
		&api.SdkNodeEnumerateRequest{})
	if err != nil {
		gerr, _ := status.FromError(err)
		fmt.Printf("Error Code[%d] Message[%s]\n",
			gerr.Code(), gerr.Message())
		os.Exit(1)
	}

	// Initialize the variable
	totalCapacity := uint64(0)

	// For each node ID, get its information
	for _, nodeID := range nodeEnumResp.GetNodeIds() {
		node, err := nodeclient.Inspect(
			context.Background(),
			&api.SdkNodeInspectRequest{
				NodeId: nodeID,
			},
		)
		if err != nil {
			gerr, _ := status.FromError(err)
			fmt.Printf("Error Code[%d] Message[%s]\n",
				gerr.Code(), gerr.Message())
			os.Exit(1)
		}

		// Get size from the pools
		// Use Pool instead of the disks, because disks could be in a RAID
		// configuration. The Pool returns the usable size.
		for _, pool := range node.GetNode().GetPools() {
			totalCapacity += pool.GetTotalSize()
		}
	}
	fmt.Printf("Cluster total capacity is %d Gi\n", totalCapacity/GB)

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
	fmt.Println()

	// Take a snapshot
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
	for taskId, status := range backupStatus.GetStatuses() {
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
