package envgetters

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
)

// BackuperEnvGetter describes specific variables for backuper instances.
type BackuperEnvGetter struct {
	MaxBackupCount int // Max allowed number of backups to be stored in storage.
}

func (beg BackuperEnvGetter) GetEnvs() []corev1.EnvVar {
	return []corev1.EnvVar{
		{
			Name:  "MAX_BACKUP_COUNT",
			Value: fmt.Sprint(beg.MaxBackupCount),
		},
	}
}
