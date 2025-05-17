package metrics

import (
	"context"

	pt "github.com/AntonShadrinNN/oiler-backup-base/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type MetricsReporter struct {
	CoreAddr string
	Secure   bool
}

func NewMetricsReporter(coreAddr string, secure bool) MetricsReporter { // coverage-ignore
	return MetricsReporter{
		CoreAddr: coreAddr,
		Secure:   secure,
	}
}

func (mr MetricsReporter) ReportStatus(ctx context.Context, metricName string, success bool, timeStamp int64) error { // coverage-ignore
	opts := make([]grpc.DialOption, 0, 1)
	if !mr.Secure {
		opts = append(opts,
			grpc.WithTransportCredentials(
				insecure.NewCredentials(),
			),
		)
	}
	conn, err := grpc.NewClient(
		mr.CoreAddr,
		opts...,
	)

	if err != nil {
		return err
	}

	client := pt.NewBackupMetricsServiceClient(conn)

	req := pt.BackupMetrics{
		BackupName: metricName,
		Success:    success,
		Timestamp:  timeStamp,
	}

	_, err = client.ReportSuccessfulBackup(ctx, &req)
	if err != nil {
		return err
	}

	return nil
}
