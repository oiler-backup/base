package backup

import (
	"context"
	"fmt"

	batchv1 "k8s.io/api/batch/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type JobsCreationError = error

var ErrAlreadyExists JobsCreationError = fmt.Errorf("already exists")

type JobsCreator struct {
	kubeClient *kubernetes.Clientset
}

func NewJobsCreator(kubeClient *kubernetes.Clientset) JobsCreator {
	return JobsCreator{
		kubeClient: kubeClient,
	}
}

func (jc JobsCreator) CreateJob(ctx context.Context, jobSpec *batchv1.Job) (string, string, error) {
	exCj, err := jc.kubeClient.BatchV1().Jobs(jobSpec.Namespace).Get(ctx, jobSpec.Name, metav1.GetOptions{})
	if apierrors.IsAlreadyExists(err) {
		return exCj.Name, exCj.Namespace, ErrAlreadyExists
	}
	if err != nil && !apierrors.IsNotFound(err) {
		return "", "", fmt.Errorf("Failed to check cj %s existence: %w", jobSpec.Name, err)
	}
	generatedJob, err := jc.kubeClient.BatchV1().Jobs(jobSpec.Namespace).Create(ctx, jobSpec, metav1.CreateOptions{})
	if err != nil {
		return "", "", fmt.Errorf("failed to create Job: %w", err)
	}

	return generatedJob.Name, generatedJob.Namespace, nil
}

func (jc JobsCreator) CreateCronJob(ctx context.Context, cronJobSpec *batchv1.CronJob) (string, string, error) {
	exCj, err := jc.kubeClient.BatchV1().CronJobs(cronJobSpec.Namespace).Get(ctx, cronJobSpec.Name, metav1.GetOptions{})
	if apierrors.IsAlreadyExists(err) {
		return exCj.Name, exCj.Namespace, ErrAlreadyExists
	}
	if err != nil && !apierrors.IsNotFound(err) {
		return "", "", fmt.Errorf("Failed to check cj %s existence: %+v", cronJobSpec.Name, err)
	}
	generatedJob, err := jc.kubeClient.BatchV1().CronJobs(cronJobSpec.Namespace).Create(ctx, cronJobSpec, metav1.CreateOptions{})
	if apierrors.IsAlreadyExists(err) {
		return exCj.Name, exCj.Namespace, ErrAlreadyExists
	} else if err != nil {
		return "", "", fmt.Errorf("failed to create CronJob: %w", err)
	}

	return generatedJob.Name, generatedJob.Namespace, nil
}
