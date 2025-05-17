package envgetters

import (
	corev1 "k8s.io/api/core/v1"
)

// RestorerEnvGetter describes specific variables for restorer instances.
type RestorerEnvGetter struct {
	BackupRevision string // Name of the backup in a storage.
}

func (reg RestorerEnvGetter) GetEnvs() []corev1.EnvVar {
	return []corev1.EnvVar{
		{
			Name:  "BACKUP_REVISION",
			Value: reg.BackupRevision,
		},
	}
}
