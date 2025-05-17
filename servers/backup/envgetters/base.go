package envgetters

import (
	corev1 "k8s.io/api/core/v1"
)

// EnvGetter provide functionality to pass Kubernetes environment variables.
type EnvGetter interface {
	GetEnvs() []corev1.EnvVar
}

// EnvGetterMerger allows to merge several EnvGetter instances in one.
type EnvGetterMerger struct {
	envgetters []EnvGetter
}

// NewEnvGetterMerger is a constructor for EnvGetterMerger.
// envgetters should not be nil.
func NewEnvGetterMerger(envgetters []EnvGetter) EnvGetterMerger {
	return EnvGetterMerger{
		envgetters: envgetters,
	}
}

// GetEnvs calls GetEnvs() on each EnvGetter provided and returns list of corev1.EnvVar
func (egm EnvGetterMerger) GetEnvs() []corev1.EnvVar {
	envs := []corev1.EnvVar{}
	for _, getter := range egm.envgetters {
		if getter != nil {
			envs = append(envs, getter.GetEnvs()...)
		}
	}

	return envs
}
