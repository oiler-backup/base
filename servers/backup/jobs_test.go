package backup

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestJobsCreator_CreateJob_Success(t *testing.T) {
	mockKube := new(MockKubeClient)
	mockBatch := new(MockBatchV1Interface)
	mockJobIf := new(MockJobInterface)

	jobSpec := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-job",
			Namespace: "default",
		},
	}

	// Настройка мока для Get (чтобы вернуть nil + IsNotFound)
	mockJobIf.On("Get", mock.Anything, "test-job", mock.Anything).
		Return(nil, apierrors.NewNotFound(batchv1.Resource("job"), "test-job"))

	// Настройка мока для Create
	mockJobIf.On("Create", mock.Anything, jobSpec, mock.Anything).
		Return(jobSpec, nil)

	// Подготовка BatchV1
	mockBatch.On("Jobs", "default").Return(mockJobIf)
	mockBatch.On("RESTClient").Return(nil)

	// Подготовка KubeClient
	mockKube.On("BatchV1").Return(mockBatch)

	// Создание тестируемого объекта
	creator := NewJobsCreator(mockKube)

	// Вызов тестируемого метода
	name, ns, err := creator.CreateJob(context.Background(), jobSpec)

	// Проверки
	require.NoError(t, err)
	assert.Equal(t, "test-job", name)
	assert.Equal(t, "default", ns)
}
func TestJobsCreator_CreateJob_AlreadyExists(t *testing.T) {
	mockKube := new(MockKubeClient)
	mockBatch := MockBatchV1Interface{}
	mockJobIf := MockJobInterface{}

	jobSpec := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-job",
			Namespace: "default",
		},
	}

	mockJobIf.On("Get", mock.Anything, "test-job", mock.Anything).
		Return(nil, errors.New("AlreadyExists"))

	mockBatch.On("Jobs", "default").Return(&mockJobIf)
	mockKube.On("BatchV1").Return(&mockBatch)

	creator := NewJobsCreator(mockKube)

	name, ns, err := creator.CreateJob(context.Background(), jobSpec)

	require.Error(t, err)
	assert.Empty(t, name)
	assert.Empty(t, ns)
}

func TestJobsCreator_UpdateCronJob_Success(t *testing.T) {
	mockKube := new(MockKubeClient)
	mockBatch := new(MockBatchV1Interface)
	mockCronJobIf := new(MockCronJobInterface)

	cronJobName := "cronjob-1"
	cronJobNamespace := "default"

	newEnvs := []corev1.EnvVar{
		{Name: "NEW_ENV", Value: "value"},
	}

	// Настройка мока для Patch
	mockCronJobIf.On("Patch",
		mock.Anything,
		cronJobName,
		types.StrategicMergePatchType,
		mock.Anything,
		mock.Anything,
		mock.AnythingOfType("[]string"),
	).Return(&batchv1.CronJob{}, nil)

	// Подготовка BatchV1
	mockBatch.On("CronJobs", cronJobNamespace).Return(mockCronJobIf)
	mockBatch.On("RESTClient").Return(nil)

	// Подготовка KubeClient
	mockKube.On("BatchV1").Return(mockBatch)

	creator := NewJobsCreator(mockKube)

	err := creator.UpdateCronJob(context.Background(), cronJobName, cronJobNamespace, newEnvs)
	require.NoError(t, err)
}

func TestJobsCreator_CreateCronJob_Success(t *testing.T) {
	mockKube := new(MockKubeClient)
	mockBatch := new(MockBatchV1Interface)
	mockCronJobIf := new(MockCronJobInterface)

	cronJobSpec := &batchv1.CronJob{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-cronjob",
			Namespace: "default",
		},
	}

	// Настройка мока: Get возвращает NotFound → можно создавать
	mockCronJobIf.On("Get", mock.Anything, "test-cronjob", mock.Anything).
		Return(nil, apierrors.NewNotFound(batchv1.Resource("cronjob"), "test-cronjob"))

	// Настройка мока: Create возвращает успешный результат
	mockCronJobIf.On("Create", mock.Anything, cronJobSpec, mock.Anything).
		Return(cronJobSpec, nil)

	// Подготовка BatchV1
	mockBatch.On("CronJobs", "default").Return(mockCronJobIf)
	mockBatch.On("RESTClient").Return(nil)
	mockKube.On("BatchV1").Return(mockBatch)

	creator := NewJobsCreator(mockKube)

	name, ns, err := creator.CreateCronJob(context.Background(), cronJobSpec)

	require.NoError(t, err)
	assert.Equal(t, "test-cronjob", name)
	assert.Equal(t, "default", ns)
}

func TestJobsCreator_CreateCronJob_GetErrorIgnored(t *testing.T) {
	mockKube := new(MockKubeClient)
	mockBatch := new(MockBatchV1Interface)
	mockCronJobIf := new(MockCronJobInterface)

	cronJobSpec := &batchv1.CronJob{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "error-cronjob",
			Namespace: "default",
		},
	}

	// Get возвращает ошибку, отличную от NotFound
	mockCronJobIf.On("Get", mock.Anything, "error-cronjob", mock.Anything).
		Return(nil, errors.New("some unexpected error"))

	// Create не должен вызываться
	mockCronJobIf.On("Create", mock.Anything, mock.Anything, mock.Anything).Maybe()

	// Подготовка BatchV1
	mockBatch.On("CronJobs", "default").Return(mockCronJobIf)
	mockBatch.On("RESTClient").Return(nil)
	mockKube.On("BatchV1").Return(mockBatch)

	creator := NewJobsCreator(mockKube)

	name, ns, err := creator.CreateCronJob(context.Background(), cronJobSpec)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to check cj error-cronjob existence")
	assert.Empty(t, name)
	assert.Empty(t, ns)
}

func TestJobsCreator_CreateCronJob_CreateError(t *testing.T) {
	mockKube := new(MockKubeClient)
	mockBatch := new(MockBatchV1Interface)
	mockCronJobIf := new(MockCronJobInterface)

	cronJobSpec := &batchv1.CronJob{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "fail-cronjob",
			Namespace: "default",
		},
	}

	// Get возвращает NotFound → пытаемся создать
	mockCronJobIf.On("Get", mock.Anything, "fail-cronjob", mock.Anything).
		Return(nil, apierrors.NewNotFound(batchv1.Resource("cronjob"), "fail-cronjob"))

	// Create возвращает ошибку
	mockCronJobIf.On("Create", mock.Anything, cronJobSpec, mock.Anything).
		Return(nil, errors.New("create failed"))

	// Подготовка BatchV1
	mockBatch.On("CronJobs", "default").Return(mockCronJobIf)
	mockBatch.On("RESTClient").Return(nil)
	mockKube.On("BatchV1").Return(mockBatch)

	creator := NewJobsCreator(mockKube)

	name, ns, err := creator.CreateCronJob(context.Background(), cronJobSpec)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to create CronJob")
	assert.Empty(t, name)
	assert.Empty(t, ns)
}
