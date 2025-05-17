package envgetters

import (
	corev1 "k8s.io/api/core/v1"
)

// CommonEnvGetter describes common variables for each adapter.
type CommonEnvGetter struct {
	DbUri        string // Database uri
	DbPort       string // Database port
	DbUser       string // Database user
	DbPass       string // Database password
	DbName       string // Database name
	S3Endpoint   string // S3-compatible storage endpoint
	S3AccessKey  string // S3-compatible storage access key
	S3SecretKey  string // S3-compatible storage secret key
	S3BucketName string // S3-compatible storage bucket name
	CoreAddr     string // Uri of Kubernetes Operator core, includes port
}

func (ceg CommonEnvGetter) GetEnvs() []corev1.EnvVar {
	return []corev1.EnvVar{
		{
			Name:  "DB_HOST",
			Value: ceg.DbUri,
		},
		{
			Name:  "DB_PORT",
			Value: ceg.DbPort,
		},
		{
			Name:  "DB_USER",
			Value: ceg.DbUser,
		},
		{
			Name:  "DB_PASSWORD",
			Value: ceg.DbPass,
		},
		{
			Name:  "DB_NAME",
			Value: ceg.DbName,
		},
		{
			Name:  "S3_ENDPOINT",
			Value: ceg.S3Endpoint,
		},
		{
			Name:  "S3_ACCESS_KEY",
			Value: ceg.S3AccessKey,
		},
		{
			Name:  "S3_SECRET_KEY",
			Value: ceg.S3SecretKey,
		},
		{
			Name:  "S3_BUCKET_NAME",
			Value: ceg.S3BucketName,
		},
		{
			Name:  "CORE_ADDR",
			Value: ceg.CoreAddr,
		},
	}
}
