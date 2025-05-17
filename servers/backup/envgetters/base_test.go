package envgetters

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

func TestEnvGetterMerger_GetEnvs(t *testing.T) {
	tests := []struct {
		name        string
		getters     []EnvGetter
		expectedEnv []corev1.EnvVar
	}{
		{
			name: "Merge two getters",
			getters: []EnvGetter{
				mockEnvGetter{
					envs: []corev1.EnvVar{
						{Name: "VAR1", Value: "val1"},
						{Name: "VAR2", Value: "val2"},
					},
				},
				mockEnvGetter{
					envs: []corev1.EnvVar{
						{Name: "VAR3", Value: "val3"},
					},
				},
			},
			expectedEnv: []corev1.EnvVar{
				{Name: "VAR1", Value: "val1"},
				{Name: "VAR2", Value: "val2"},
				{Name: "VAR3", Value: "val3"},
			},
		},
		{
			name: "Empty getters list",
			getters: []EnvGetter{
				mockEnvGetter{envs: []corev1.EnvVar{}},
				mockEnvGetter{envs: []corev1.EnvVar{}},
			},
			expectedEnv: []corev1.EnvVar{},
		},
		{
			name: "Nil input",
			getters: []EnvGetter{
				nil,
				mockEnvGetter{envs: []corev1.EnvVar{}},
			},
			expectedEnv: []corev1.EnvVar{},
		},
		{
			name: "One getter with empty slice",
			getters: []EnvGetter{
				mockEnvGetter{envs: []corev1.EnvVar{}},
				mockEnvGetter{
					envs: []corev1.EnvVar{
						{Name: "VAR1", Value: "val1"},
					},
				},
			},
			expectedEnv: []corev1.EnvVar{
				{Name: "VAR1", Value: "val1"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merger := NewEnvGetterMerger(tt.getters)
			result := merger.GetEnvs()

			require.Equal(t, len(tt.expectedEnv), len(result))

			for i := range tt.expectedEnv {
				assert.Equal(t, tt.expectedEnv[i], result[i])
			}
		})
	}
}
