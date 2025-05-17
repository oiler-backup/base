package s3

import (
	"context"
	"errors"
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

func TestCleanAndUpload_CleanError(t *testing.T) {
	mockUploader := new(MockS3Uploader)
	mockCleaner := new(MockS3Cleaner)

	uploadCleaner := S3UploadCleaner{
		u: mockUploader,
		c: mockCleaner,
	}

	expectedErr := errors.New("clean failed")
	mockCleaner.On("Clean", mock.Anything, "test-bucket", "backups/", 5).Return(expectedErr)

	err := uploadCleaner.CleanAndUpload(context.Background(), "test-bucket", "backups/", 5, "backup.tar.gz", nil)

	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to clean S3")
	mockUploader.AssertNotCalled(t, "Upload")
}

func TestCleanAndUpload_UploadError(t *testing.T) {
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
	mockUploader.On("Upload", mock.Anything, bucketName, fileName, fileContent).Return(errors.New("upload failed"))

	err := uploadCleaner.CleanAndUpload(context.Background(), bucketName, backupDir, maxBackupCount, fileName, fileContent)

	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to upload object to S3")
	mockCleaner.AssertExpectations(t)
	mockUploader.AssertExpectations(t)
}
