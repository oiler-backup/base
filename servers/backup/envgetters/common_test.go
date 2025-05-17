package envgetters

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCommonEnvGetter_GetEnvs(t *testing.T) {
	testCases := []struct {
		name         string
		input        CommonEnvGetter
		expectedEnvs map[string]string
	}{
		{
			name: "All values set",
			input: CommonEnvGetter{
				DbUri:        "localhost",
				DbPort:       "5432",
				DbUser:       "postgres",
				DbPass:       "secret",
				DbName:       "maindb",
				S3Endpoint:   "http://minio.local",
				S3AccessKey:  "access123",
				S3SecretKey:  "secret456",
				S3BucketName: "mybucket",
				CoreAddr:     "core-service:50051",
			},
			expectedEnvs: map[string]string{
				"DB_HOST":        "localhost",
				"DB_PORT":        "5432",
				"DB_USER":        "postgres",
				"DB_PASSWORD":    "secret",
				"DB_NAME":        "maindb",
				"S3_ENDPOINT":    "http://minio.local",
				"S3_ACCESS_KEY":  "access123",
				"S3_SECRET_KEY":  "secret456",
				"S3_BUCKET_NAME": "mybucket",
				"CORE_ADDR":      "core-service:50051",
			},
		},
		{
			name:  "Empty values",
			input: CommonEnvGetter{},
			expectedEnvs: map[string]string{
				"DB_HOST":        "",
				"DB_PORT":        "",
				"DB_USER":        "",
				"DB_PASSWORD":    "",
				"DB_NAME":        "",
				"S3_ENDPOINT":    "",
				"S3_ACCESS_KEY":  "",
				"S3_SECRET_KEY":  "",
				"S3_BUCKET_NAME": "",
				"CORE_ADDR":      "",
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			envs := tt.input.GetEnvs()

			require.Len(t, envs, len(tt.expectedEnvs))

			for _, env := range envs {
				expectedValue, exists := tt.expectedEnvs[env.Name]
				assert.True(t, exists, "Unexpected environment variable: %s", env.Name)
				assert.Equal(t, expectedValue, env.Value)
			}
		})
	}
}
