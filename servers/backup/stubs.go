package backup

import (
	"fmt"

	"github.com/google/uuid"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type EnvGetter interface {
	GetEnvs() []corev1.EnvVar
}

type EnvGetterMerger struct {
	envegetters []EnvGetter
}

func NewEnvGetterMerger(envgetters []EnvGetter) EnvGetterMerger {
	return EnvGetterMerger{
		envegetters: envgetters,
	}
}

func (egm EnvGetterMerger) GetEnvs() []corev1.EnvVar {
	envs := []corev1.EnvVar{}
	for _, getter := range egm.envegetters {
		envs = append(envs, getter.GetEnvs()...)
	}

	return envs
}

type JobsStub struct {
	namePsfx      string
	namespace     string
	backuperImage string
	restorerImage string

	restorerJob     *batchv1.Job
	backuperCronJob *batchv1.CronJob
}

func NewJobsStub(name, namespace, backuperImg, restorerImg string) JobsStub {
	return JobsStub{
		namePsfx:      name,
		namespace:     namespace,
		backuperImage: backuperImg,
		restorerImage: restorerImg,
	}
}

func (js JobsStub) BuildBackuperCj(schedule string, eg EnvGetter) *batchv1.CronJob {
	return &batchv1.CronJob{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("backup-%s-%s", uuid.NewString()[:8], js.namePsfx),
			Namespace: js.namespace,
		},
		Spec: batchv1.CronJobSpec{
			Schedule: schedule,
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

func (js JobsStub) BuildRestorerJob(eg EnvGetter) *batchv1.Job {
	return &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("restore-%s-%s", uuid.NewString()[:8], js.namePsfx),
			Namespace: js.restorerImage,
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
