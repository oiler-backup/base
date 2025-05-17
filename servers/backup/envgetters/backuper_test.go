// envgetters_test.go
package envgetters

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBackuperEnvGetter_GetEnvs(t *testing.T) {
	tests := []struct {
		name           string
		maxBackupCount int
		expectedValue  string
	}{
		{
			name:           "Positive count",
			maxBackupCount: 5,
			expectedValue:  "5",
		},
		{
			name:           "Zero",
			maxBackupCount: 0,
			expectedValue:  "0",
		},
		{
			name:           "Negative count",
			maxBackupCount: -1,
			expectedValue:  "-1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			beg := BackuperEnvGetter{
				MaxBackupCount: tt.maxBackupCount,
			}

			envs := beg.GetEnvs()

			require.Len(t, envs, 1)
			assert.Equal(t, "MAX_BACKUP_COUNT", envs[0].Name)
			assert.Equal(t, tt.expectedValue, envs[0].Value)
		})
	}
}
