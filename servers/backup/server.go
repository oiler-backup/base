package backup

import (
	"context"
	"errors"
	"fmt"
	"log"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	pb "github.com/AntonShadrinNN/oiler-backup-base/proto"
	batchv1 "k8s.io/api/batch/v1"
)

type ErrBackupServer = error

var ErrAlreadyExists ErrBackupServer = fmt.Errorf("Already exists")

type BackupServer struct {
	pb.UnimplementedBackupServiceServer
	kubeClient *kubernetes.Clientset
}

func NewBackupServer(backupCjSpec *batchv1.CronJob, restoreJobSpec *batchv1.Job) (*BackupServer, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load Kubernetes config: %w", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create Kubernetes client: %w", err)
	}

	return &BackupServer{kubeClient: clientset}, nil
}

func (s *BackupServer) Backup(ctx context.Context, backupCjSpec *batchv1.CronJob) (*pb.BackupResponse, error) {
	name, namespace, err := s.createCronJob(ctx, backupCjSpec)
	if errors.Is(err, ErrAlreadyExists) {
		return &pb.BackupResponse{
			Status:           "Exists",
			CronjobName:      name,
			CronjobNamespace: namespace,
		}, nil
	}
	if err != nil {
		log.Printf("Failed to create CronJob: %v", err)
		return &pb.BackupResponse{Status: "Failed to create CronJob"}, nil
	}

	return &pb.BackupResponse{
		Status:           "CronJob created successfully",
		CronjobName:      name,
		CronjobNamespace: namespace,
	}, nil
}

func (s *BackupServer) Restore(ctx context.Context, restoreJobSpec *batchv1.Job) (*pb.BackupRestoreResponse, error) {
	name, namespace, err := s.createJob(ctx, restoreJobSpec)
	if errors.Is(err, ErrAlreadyExists) {
		return &pb.BackupRestoreResponse{
			Status:       "Exists",
			JobName:      name,
			JobNamespace: namespace,
		}, nil
	}
	if err != nil {
		log.Printf("Failed to create Job: %v", err)
		return &pb.BackupRestoreResponse{Status: "Failed to create Job"}, nil
	}

	return &pb.BackupRestoreResponse{
		Status:       "Job created successfully",
		JobName:      name,
		JobNamespace: namespace,
	}, nil
}
