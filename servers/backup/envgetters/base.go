package envgetters

import (
	corev1 "k8s.io/api/core/v1"
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
		if getter != nil {
			envs = append(envs, getter.GetEnvs()...)
		}
	}

	return envs
}
