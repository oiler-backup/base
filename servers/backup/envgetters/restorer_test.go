// restorer_env_getter_test.go
package envgetters

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRestorerEnvGetter_GetEnvs(t *testing.T) {
	tests := []struct {
		name           string
		backupRevision string
		expectedValue  string
	}{
		{
			name:           "Valid revision",
			backupRevision: "rev-12345",
			expectedValue:  "rev-12345",
		},
		{
			name:           "Empty revision",
			backupRevision: "",
			expectedValue:  "",
		},
		{
			name:           "Whitespace-only revision",
			backupRevision: "   ",
			expectedValue:  "   ",
		},
		{
			name:           "Special chars",
			backupRevision: "!@#$%^&*()",
			expectedValue:  "!@#$%^&*()",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reg := RestorerEnvGetter{
				BackupRevision: tt.backupRevision,
			}

			envs := reg.GetEnvs()

			require.Len(t, envs, 1)
			assert.Equal(t, "BACKUP_REVISION", envs[0].Name)
			assert.Equal(t, tt.expectedValue, envs[0].Value)
		})
	}
}
