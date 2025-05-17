// Package metrics provides functionality to work with metrics.
//
// metrics should only be used in backupers and restorers to communicate with
// adapter core.
// metrics is basically a gRPC-client.
package metrics

import (
	"context"

	pt "github.com/AntonShadrinNN/oiler-backup-base/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// A MetricsReporter reports metrics over gRPC to CoreAddr.
// Optionally uses TLS/SSL encryption.
type MetricsReporter struct {
	CoreAddr string
	Secure   bool
}

// NewMetricsReporter is a constructor for MetricsReporter.
// coreAddr is an ip4-addres of domain name of an Kubernetes Operator core.
// secure forces to use TLS/SSL encryption.
func NewMetricsReporter(coreAddr string, secure bool) MetricsReporter { // coverage-ignore
	return MetricsReporter{
		CoreAddr: coreAddr,
		Secure:   secure,
	}
}

// ReportStatus sends a message to Kubernetes Operator core with parameters required to generate Prometheus Metrics.
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
