// kube_client_mock.go
package backup

import (
	"context"

	"github.com/stretchr/testify/mock"
	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	applyconfigurationsbatchv1 "k8s.io/client-go/applyconfigurations/batch/v1"
	batchv1types "k8s.io/client-go/kubernetes/typed/batch/v1"
	"k8s.io/client-go/rest"
)

// ----------------------
// 🧪 MockJobInterface
// ----------------------
type MockJobInterface struct {
	mock.Mock
}

func (m *MockJobInterface) Create(ctx context.Context, job *batchv1.Job, opts metav1.CreateOptions) (*batchv1.Job, error) {
	args := m.Called(ctx, job, opts)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*batchv1.Job), args.Error(1)
}

func (m *MockJobInterface) Update(ctx context.Context, job *batchv1.Job, opts metav1.UpdateOptions) (*batchv1.Job, error) {
	args := m.Called(ctx, job, opts)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*batchv1.Job), args.Error(1)
}

func (m *MockJobInterface) UpdateStatus(ctx context.Context, job *batchv1.Job, opts metav1.UpdateOptions) (*batchv1.Job, error) {
	args := m.Called(ctx, job, opts)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*batchv1.Job), args.Error(1)
}

func (m *MockJobInterface) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	args := m.Called(ctx, name, opts)
	return args.Error(0)
}

func (m *MockJobInterface) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	args := m.Called(ctx, opts, listOpts)
	return args.Error(0)
}

func (m *MockJobInterface) Get(ctx context.Context, name string, opts metav1.GetOptions) (*batchv1.Job, error) {
	args := m.Called(ctx, name, opts)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*batchv1.Job), args.Error(1)
}

func (m *MockJobInterface) List(ctx context.Context, opts metav1.ListOptions) (*batchv1.JobList, error) {
	args := m.Called(ctx, opts)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*batchv1.JobList), args.Error(1)
}

func (m *MockJobInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	args := m.Called(ctx, opts)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(watch.Interface), args.Error(1)
}

func (m *MockJobInterface) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*batchv1.Job, error) {
	args := m.Called(ctx, name, pt, data, opts, subresources)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*batchv1.Job), args.Error(1)
}

// Apply и ApplyStatus требуют импорта applyconfigurationsbatchv1
// Если у тебя не поддерживается — можно удалить
func (m *MockJobInterface) Apply(ctx context.Context, job *applyconfigurationsbatchv1.JobApplyConfiguration, opts metav1.ApplyOptions) (*batchv1.Job, error) {
	args := m.Called(ctx, job, opts)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*batchv1.Job), args.Error(1)
}

func (m *MockJobInterface) ApplyStatus(ctx context.Context, job *applyconfigurationsbatchv1.JobApplyConfiguration, opts metav1.ApplyOptions) (*batchv1.Job, error) {
	args := m.Called(ctx, job, opts)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*batchv1.Job), args.Error(1)
}

// -----------------------------
// 🧪 MockCronJobInterface
// -----------------------------
type MockCronJobInterface struct {
	mock.Mock
}

func (m *MockCronJobInterface) Create(ctx context.Context, cronJob *batchv1.CronJob, opts metav1.CreateOptions) (*batchv1.CronJob, error) {
	args := m.Called(ctx, cronJob, opts)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*batchv1.CronJob), args.Error(1)
}

func (m *MockCronJobInterface) Update(ctx context.Context, cronJob *batchv1.CronJob, opts metav1.UpdateOptions) (*batchv1.CronJob, error) {
	args := m.Called(ctx, cronJob, opts)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*batchv1.CronJob), args.Error(1)
}

func (m *MockCronJobInterface) UpdateStatus(ctx context.Context, cronJob *batchv1.CronJob, opts metav1.UpdateOptions) (*batchv1.CronJob, error) {
	args := m.Called(ctx, cronJob, opts)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*batchv1.CronJob), args.Error(1)
}

func (m *MockCronJobInterface) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	args := m.Called(ctx, name, opts)
	return args.Error(0)
}

func (m *MockCronJobInterface) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	args := m.Called(ctx, opts, listOpts)
	return args.Error(0)
}

func (m *MockCronJobInterface) Get(ctx context.Context, name string, opts metav1.GetOptions) (*batchv1.CronJob, error) {
	args := m.Called(ctx, name, opts)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*batchv1.CronJob), args.Error(1)
}

func (m *MockCronJobInterface) List(ctx context.Context, opts metav1.ListOptions) (*batchv1.CronJobList, error) {
	args := m.Called(ctx, opts)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*batchv1.CronJobList), args.Error(1)
}

func (m *MockCronJobInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	args := m.Called(ctx, opts)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(watch.Interface), args.Error(1)
}

func (m *MockCronJobInterface) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*batchv1.CronJob, error) {
	args := m.Called(ctx, name, pt, data, opts, subresources)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*batchv1.CronJob), args.Error(1)
}

func (m *MockCronJobInterface) Apply(ctx context.Context, cj *applyconfigurationsbatchv1.CronJobApplyConfiguration, opts metav1.ApplyOptions) (*batchv1.CronJob, error) {
	args := m.Called(ctx, cj, opts)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*batchv1.CronJob), args.Error(1)
}

func (m *MockCronJobInterface) ApplyStatus(ctx context.Context, cj *applyconfigurationsbatchv1.CronJobApplyConfiguration, opts metav1.ApplyOptions) (*batchv1.CronJob, error) {
	args := m.Called(ctx, cj, opts)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*batchv1.CronJob), args.Error(1)
}

// -----------------------------
// 🧪 MockBatchV1Interface
// -----------------------------
type MockBatchV1Interface struct {
	mock.Mock
}

func (m *MockBatchV1Interface) Jobs(namespace string) batchv1types.JobInterface {
	args := m.Called(namespace)
	return args.Get(0).(batchv1types.JobInterface)
}

func (m *MockBatchV1Interface) CronJobs(namespace string) batchv1types.CronJobInterface {
	args := m.Called(namespace)
	return args.Get(0).(batchv1types.CronJobInterface)
}

func (m *MockBatchV1Interface) RESTClient() rest.Interface {
	args := m.Called()
	return args.Get(0).(rest.Interface)
}

// -----------------------------
// 🧪 MockKubeClient
// -----------------------------
type MockKubeClient struct {
	mock.Mock
}

func (m *MockKubeClient) BatchV1() batchv1types.BatchV1Interface {
	args := m.Called()
	return args.Get(0).(batchv1types.BatchV1Interface)
}
