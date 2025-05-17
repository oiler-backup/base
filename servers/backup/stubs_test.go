package backup

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
)

type mockEnvGetter struct {
	envs []corev1.EnvVar
}

func (m mockEnvGetter) GetEnvs() []corev1.EnvVar {
	return m.envs
}

func TestJobsStub_BuildBackuperCj(t *testing.T) {
	const (
		nameSuffix     = "test-backup"
		namespace      = "default"
		backuperImage  = "backuper:latest"
		schedule       = "0 0 * * *"
		expectedPrefix = "backup-"
	)

	testEnvs := []corev1.EnvVar{
		{Name: "TEST_ENV", Value: "test_value"},
	}
	getter := mockEnvGetter{envs: testEnvs}

	js := NewJobsStub(nameSuffix, namespace, backuperImage, "")

	cronJob := js.BuildBackuperCj(schedule, getter)

	require.NotNil(t, cronJob)
	assert.Equal(t, namespace, cronJob.Namespace)
	assert.Contains(t, cronJob.Name, expectedPrefix)
	assert.Equal(t, schedule, cronJob.Spec.Schedule)
	assert.Equal(t, corev1.PullAlways, cronJob.Spec.JobTemplate.Spec.Template.Spec.Containers[0].ImagePullPolicy)
	assert.Equal(t, testEnvs, cronJob.Spec.JobTemplate.Spec.Template.Spec.Containers[0].Env)
	assert.Equal(t, corev1.RestartPolicyNever, cronJob.Spec.JobTemplate.Spec.Template.Spec.RestartPolicy)

	historyLimit := int32(1)
	assert.Equal(t, &historyLimit, cronJob.Spec.SuccessfulJobsHistoryLimit)
	assert.Equal(t, &historyLimit, cronJob.Spec.FailedJobsHistoryLimit)
}

func TestJobsStub_BuildRestorerJob(t *testing.T) {
	const (
		nameSuffix     = "test-restore"
		namespace      = "default"
		restorerImage  = "restorer:latest"
		expectedPrefix = "restore-"
	)

	testEnvs := []corev1.EnvVar{
		{Name: "RESTORE_ENV", Value: "restore_value"},
	}
	getter := mockEnvGetter{envs: testEnvs}

	js := NewJobsStub(nameSuffix, namespace, "", restorerImage)

	job := js.BuildRestorerJob(getter)

	require.NotNil(t, job)
	assert.Equal(t, namespace, job.Namespace)
	assert.Contains(t, job.Name, expectedPrefix)
	assert.Equal(t, corev1.PullAlways, job.Spec.Template.Spec.Containers[0].ImagePullPolicy)
	assert.Equal(t, testEnvs, job.Spec.Template.Spec.Containers[0].Env)
	assert.Equal(t, corev1.RestartPolicyOnFailure, job.Spec.Template.Spec.RestartPolicy)
}
