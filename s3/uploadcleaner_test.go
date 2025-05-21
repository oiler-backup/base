package s3

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCleanAndUpload_Success(t *testing.T) {
	mockUploader := new(MockS3Uploader)
	mockCleaner := new(MockS3Cleaner)

	uploadCleaner := S3UploadCleaner{
		u: mockUploader,
		c: mockCleaner,
	}

	bucketName := "test-bucket"
	backupDir := "/backups"
	maxBackupCount := 5
	fileName := "backup.tar.gz"
	fileContent := strings.NewReader("some content")

	mockCleaner.On("Clean", mock.Anything, bucketName, backupDir, maxBackupCount).Return(nil)
	mockUploader.On("Upload", mock.Anything, bucketName, fileName, fileContent).Return(nil)

	err := uploadCleaner.CleanAndUpload(context.Background(), bucketName, backupDir, maxBackupCount, fileName, fileContent)

	require.NoError(t, err)
	mockCleaner.AssertExpectations(t)
	mockUploader.AssertExpectations(t)
}

func TestCleanAndUpload_UploadError(t *testing.T) {
	mockUploader := new(MockS3Uploader)
	mockCleaner := new(MockS3Cleaner)

	uploadCleaner := S3UploadCleaner{
		u: mockUploader,
		c: mockCleaner,
	}

	expectedErr := fmt.Errorf("upload failed")
	mockUploader.On("Upload", mock.Anything, "test-bucket", "backup.tar.gz", nil).Return(expectedErr)
	err := uploadCleaner.CleanAndUpload(context.Background(), "test-bucket", "backups/", 5, "backup.tar.gz", nil)

	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to upload object to S3")
	mockCleaner.AssertNotCalled(t, "Clean")
}

func TestCleanAndUpload_CleanError(t *testing.T) {
	mockUploader := new(MockS3Uploader)
	mockCleaner := new(MockS3Cleaner)

	uploadCleaner := S3UploadCleaner{
		u: mockUploader,
		c: mockCleaner,
	}

	bucketName := "test-bucket"
	backupDir := "/backups"
	maxBackupCount := 5
	fileName := "backup.tar.gz"
	fileContent := strings.NewReader("some content")
	expectedErr := fmt.Errorf("clean error")
	mockUploader.On("Upload", mock.Anything, bucketName, fileName, fileContent).Return(nil)
	mockCleaner.On("Clean", mock.Anything, bucketName, backupDir, maxBackupCount).Return(expectedErr)

	err := uploadCleaner.CleanAndUpload(context.Background(), bucketName, backupDir, maxBackupCount, fileName, fileContent)

	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to clean S3")
	mockUploader.AssertExpectations(t)
	mockCleaner.AssertExpectations(t)
}
