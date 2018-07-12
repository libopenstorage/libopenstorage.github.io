// Example code
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/libopenstorage/openstorage/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func main() {

	// Setup a connection
	conn, err := grpc.Dial("localhost:9100", grpc.WithInsecure())
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

	// Take a snapshot
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

}
