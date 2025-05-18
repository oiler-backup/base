package backup

import (
	"context"

	"github.com/oiler-backup/base/servers/backup/envgetters"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/client-go/kubernetes/typed/batch/v1"
)

// IKubeClient describes necessary methods from Kubernetes client.
// Mainly for testing purposes.
type IKubeClient interface {
	BatchV1() v1.BatchV1Interface
}

// IJobStub provides methods to create manifests of Kubernetes Jobs and CronJobs.
type IJobStub interface {
	// BuildBackuperCj returns manifest of a backuper CronJob passing envs defined by
	// eg to underlying CronJob.
	BuildBackuperCj(schedule string, eg envgetters.EnvGetter) *batchv1.CronJob
	// BuildRestorerJob returns manifest of a restorer Job passing envs defined by
	// eg to underlying Job.
	BuildRestorerJob(eg envgetters.EnvGetter) *batchv1.Job
}

// IJobsCreator creates Kubernetes Jobs and CronJobs and applies them to cluster.
type IJobsCreator interface {
	// CreateJob creates Kubernetes Job and returns its name and namespace.
	CreateJob(ctx context.Context, jobSpec *batchv1.Job) (string, string, error)
	// CreateCronJob creates Kubernetes Job and returns its name and namespace.
	CreateCronJob(ctx context.Context, cronJobSpec *batchv1.CronJob) (string, string, error)
	// UpdateCronJob updated CronJob cronJobName in cronJobNamespace by passing newEnvs to environment variables.
	// Overrides old keys.
	UpdateCronJob(ctx context.Context, cronJobName, cronJobNamespace string, newEnvs []corev1.EnvVar) error
}
