package metrics

import (
	"context"
	"fmt"

	pt "github.com/AntonShadrinNN/oiler-backup-base/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func reportStatus(ctx context.Context, coreAddr string, name string, success bool, timeStamp int64, secure bool) error {
	opts := make([]grpc.DialOption, 0, 1)
	if !secure {
		opts = append(opts,
			grpc.WithTransportCredentials(
				insecure.NewCredentials(),
			),
		)
	}
	conn, err := grpc.NewClient(
		coreAddr,
		opts...,
	)

	if err != nil {
		return fmt.Errorf("failed to connect to %s: %w", coreAddr, err)
	}
	defer conn.Close()

	client := pt.NewBackupMetricsServiceClient(conn)

	req := pt.BackupMetrics{
		BackupName: name,
		Success:    success,
		Timestamp:  timeStamp,
	}

	_, err = client.ReportSuccessfulBackup(ctx, &req)
	if err != nil {
		return err
	}

	return nil
}
