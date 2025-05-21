package s3

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/stretchr/testify/mock"
)

type MockS3Client struct {
	mock.Mock
}

func (m *MockS3Client) CreateMultipartUpload(ctx context.Context, params *s3.CreateMultipartUploadInput, optFns ...func(*s3.Options)) (*s3.CreateMultipartUploadOutput, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(*s3.CreateMultipartUploadOutput), args.Error(1)
}

func (m *MockS3Client) UploadPart(ctx context.Context, params *s3.UploadPartInput, optFns ...func(*s3.Options)) (*s3.UploadPartOutput, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(*s3.UploadPartOutput), args.Error(1)
}
func (m *MockS3Client) CompleteMultipartUpload(ctx context.Context, params *s3.CompleteMultipartUploadInput, optFns ...func(*s3.Options)) (*s3.CompleteMultipartUploadOutput, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(*s3.CompleteMultipartUploadOutput), args.Error(1)
}

func (m *MockS3Client) ListObjectsV2(ctx context.Context, params *s3.ListObjectsV2Input, optFns ...func(*s3.Options)) (*s3.ListObjectsV2Output, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(*s3.ListObjectsV2Output), args.Error(1)
}

func (m *MockS3Client) DeleteObjects(ctx context.Context, params *s3.DeleteObjectsInput, optFns ...func(*s3.Options)) (*s3.DeleteObjectsOutput, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(*s3.DeleteObjectsOutput), args.Error(1)
}

type MockS3Uploader struct {
	mock.Mock
}

func (m *MockS3Uploader) Upload(ctx context.Context, bucketName, objectKey string, fileContent io.Reader) error {
	args := m.Called(ctx, bucketName, objectKey, fileContent)
	return args.Error(0)
}

type MockS3Cleaner struct {
	mock.Mock
}

func (m *MockS3Cleaner) Clean(ctx context.Context, bucketName, backupDir string, maxBackupCount int) error {
	args := m.Called(ctx, bucketName, backupDir, maxBackupCount)
	return args.Error(0)
}
