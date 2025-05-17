package backup

import (
	"context"

	"github.com/AntonShadrinNN/oiler-backup-base/servers/backup/envgetters"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/client-go/kubernetes/typed/batch/v1"
)

type KubeClient interface {
	BatchV1() v1.BatchV1Interface
}

type IJobStub interface {
	BuildBackuperCj(schedule string, eg envgetters.EnvGetter) *batchv1.CronJob
	BuildRestorerJob(eg envgetters.EnvGetter) *batchv1.Job
}

type IJobsCreator interface {
	CreateJob(ctx context.Context, jobSpec *batchv1.Job) (string, string, error)
	CreateCronJob(ctx context.Context, cronJobSpec *batchv1.CronJob) (string, string, error)
	UpdateCronJob(ctx context.Context, cronJobName, cronJobNamespace string, newEnvs []corev1.EnvVar) error
}
