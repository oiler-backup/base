package backup

import (
	"fmt"

	"github.com/AntonShadrinNN/oiler-backup-base/servers/backup/envgetters"
	"github.com/google/uuid"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// JobsStub  provides methods to create manifests of Kubernetes Jobs and CronJobs.
type JobsStub struct {
	namePsfx      string
	namespace     string
	backuperImage string
	restorerImage string
}

// NewJobsStub is a constructor for JobsStub.
// It requires to specify name which will be a postfix for generated resources name.
// backuperImg and restorerImg will be placed in image section of a manifest.
func NewJobsStub(name, namespace, backuperImg, restorerImg string) JobsStub {
	return JobsStub{
		namePsfx:      name,
		namespace:     namespace,
		backuperImage: backuperImg,
		restorerImage: restorerImg,
	}
}

// BuildBackuperCj returns manifest of a backuper CronJob passing envs defined by
// eg to underlying CronJob.
func (js JobsStub) BuildBackuperCj(schedule string, eg envgetters.EnvGetter) *batchv1.CronJob {
	var historyLimit int32 = 1
	return &batchv1.CronJob{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("backup-%s-%s", uuid.NewString()[:8], js.namePsfx),
			Namespace: js.namespace,
		},
		Spec: batchv1.CronJobSpec{
			SuccessfulJobsHistoryLimit: &historyLimit,
			FailedJobsHistoryLimit:     &historyLimit,
			Schedule:                   schedule,
			JobTemplate: batchv1.JobTemplateSpec{
				Spec: batchv1.JobSpec{
					Template: corev1.PodTemplateSpec{
						Spec: corev1.PodSpec{
							Containers: []corev1.Container{
								{
									Name:            "backup-job",
									Image:           js.backuperImage,
									ImagePullPolicy: corev1.PullAlways,
									Env:             eg.GetEnvs(),
								},
							},
							RestartPolicy: corev1.RestartPolicyNever,
						},
					},
				},
			},
		},
	}
}

// BuildRestorerJob returns manifest of a restorer Job passing envs defined by
// eg to underlying Job.
func (js JobsStub) BuildRestorerJob(eg envgetters.EnvGetter) *batchv1.Job {
	return &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("restore-%s-%s", uuid.NewString()[:8], js.namePsfx),
			Namespace: js.namespace,
		},
		Spec: batchv1.JobSpec{
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:            "backup-restore-job",
							Image:           js.restorerImage,
							ImagePullPolicy: corev1.PullAlways,
							Env:             eg.GetEnvs(),
						},
					},
					RestartPolicy: corev1.RestartPolicyOnFailure,
				},
			},
		},
	}
}
