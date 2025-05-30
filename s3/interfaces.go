package s3

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// An IS3Client provides functionality to work with s3-compatible storage.
//
// See https://pkg.go.dev/github.com/aws/aws-sdk-go for more information.
type IS3Client interface {
	// ListObjectsV2 returns list of objects in a specified bucket.
	ListObjectsV2(ctx context.Context, params *s3.ListObjectsV2Input, optFns ...func(*s3.Options)) (*s3.ListObjectsV2Output, error)
	GetObject(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error)
	// DeleteObjects deletes specified files.
	DeleteObjects(ctx context.Context, params *s3.DeleteObjectsInput, optFns ...func(*s3.Options)) (*s3.DeleteObjectsOutput, error)
	CreateMultipartUpload(ctx context.Context, params *s3.CreateMultipartUploadInput, optFns ...func(*s3.Options)) (*s3.CreateMultipartUploadOutput, error)
	UploadPart(ctx context.Context, params *s3.UploadPartInput, optFns ...func(*s3.Options)) (*s3.UploadPartOutput, error)
	CompleteMultipartUpload(ctx context.Context, params *s3.CompleteMultipartUploadInput, optFns ...func(*s3.Options)) (*s3.CompleteMultipartUploadOutput, error)
	AbortMultipartUpload(ctx context.Context, params *s3.AbortMultipartUploadInput, optFns ...func(*s3.Options)) (*s3.AbortMultipartUploadOutput, error)
	PutObject(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error)
}

// An IS3Uploader provides functionality to Upload data to s3-compatible storage.
type IS3Uploader interface {
	// Upload uploads a single file.
	Upload(ctx context.Context, bucketName, objectKey string, fileContent io.Reader) error
}

// An IS3Uploader provides functionality to delete data from s3-compatible storage.
type IS3Cleaner interface {
	// Clean deletes files according to maxBackupCount.
	Clean(ctx context.Context, bucketName, backupDir string, maxBackupCount int) error
}
