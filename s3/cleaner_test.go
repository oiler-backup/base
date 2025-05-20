package s3

// import (
// 	"context"
// 	"fmt"
// 	"testing"
// 	"time"

// 	"github.com/aws/aws-sdk-go-v2/aws"
// 	"github.com/aws/aws-sdk-go-v2/service/s3"
// 	"github.com/aws/aws-sdk-go-v2/service/s3/types"
// 	"github.com/stretchr/testify/mock"
// 	"github.com/stretchr/testify/require"
// )

// func TestClean_NoCleanupNeeded(t *testing.T) {
// 	mockClient := new(MockS3Client)
// 	cleaner := S3Cleaner{client: mockClient}
// 	ctx := context.Background()

// 	bucketName := "test-bucket"
// 	backupDir := "backups/"
// 	maxBackupCount := 5

// 	objects := []types.Object{
// 		{Key: aws.String("backups/file1"), LastModified: aws.Time(time.Now())},
// 		{Key: aws.String("backups/file2"), LastModified: aws.Time(time.Now())},
// 	}

// 	mockClient.On("ListObjectsV2", ctx, &s3.ListObjectsV2Input{
// 		Bucket: aws.String(bucketName),
// 		Prefix: aws.String(backupDir),
// 	}).Return(&s3.ListObjectsV2Output{Contents: objects}, nil)

// 	err := cleaner.Clean(ctx, bucketName, backupDir, maxBackupCount)

// 	require.NoError(t, err)
// 	mockClient.AssertExpectations(t)
// }

// func TestClean_DeleteOldBackups(t *testing.T) {
// 	mockClient := new(MockS3Client)
// 	cleaner := S3Cleaner{client: mockClient}

// 	bucketName := "test-bucket"
// 	backupDir := "backups/"
// 	maxBackupCount := 2

// 	now := time.Now()
// 	objects := []types.Object{
// 		{Key: aws.String(fmt.Sprintf("%s/file1", backupDir)), LastModified: aws.Time(now.Add(-4 * time.Hour))},
// 		{Key: aws.String(fmt.Sprintf("%s/file2", backupDir)), LastModified: aws.Time(now.Add(-3 * time.Hour))},
// 		{Key: aws.String(fmt.Sprintf("%s/file3", backupDir)), LastModified: aws.Time(now.Add(-2 * time.Hour))},
// 		{Key: aws.String(fmt.Sprintf("%s/file4", backupDir)), LastModified: aws.Time(now.Add(-1 * time.Hour))},
// 	}

// 	toDelete := []types.ObjectIdentifier{
// 		{Key: aws.String(fmt.Sprintf("%s/file1", backupDir))},
// 		{Key: aws.String(fmt.Sprintf("%s/file2", backupDir))},
// 	}

// 	mockClient.On("ListObjectsV2", mock.Anything, &s3.ListObjectsV2Input{
// 		Bucket: aws.String(bucketName),
// 		Prefix: aws.String(backupDir),
// 	}).Return(&s3.ListObjectsV2Output{Contents: objects}, nil)

// 	mockClient.On("DeleteObjects", mock.Anything, &s3.DeleteObjectsInput{
// 		Bucket: aws.String(bucketName),
// 		Delete: &types.Delete{
// 			Objects: toDelete,
// 		},
// 	}).Return(&s3.DeleteObjectsOutput{}, nil)

// 	err := cleaner.Clean(context.Background(), bucketName, backupDir, maxBackupCount)

// 	require.NoError(t, err)
// 	mockClient.AssertExpectations(t)
// }

// func TestClean_ListError(t *testing.T) {
// 	mockClient := new(MockS3Client)
// 	cleaner := S3Cleaner{client: mockClient}

// 	bucketName := "test-bucket"
// 	backupDir := "backups/"
// 	maxBackupCount := 2

// 	mockClient.On("ListObjectsV2", mock.Anything, mock.Anything).
// 		Return((*s3.ListObjectsV2Output)(nil), fmt.Errorf("list error"))

// 	err := cleaner.Clean(context.Background(), bucketName, backupDir, maxBackupCount)

// 	require.Error(t, err)
// 	require.Contains(t, err.Error(), "failed to list objects")
// 	mockClient.AssertExpectations(t)
// }

// func TestClean_DeleteError(t *testing.T) {
// 	mockClient := new(MockS3Client)
// 	cleaner := S3Cleaner{client: mockClient}

// 	bucketName := "test-bucket"
// 	backupDir := "backups/"
// 	maxBackupCount := 1

// 	objects := []types.Object{
// 		{Key: aws.String(fmt.Sprintf("%s/file1", backupDir)), LastModified: aws.Time(time.Now().Add(-time.Hour))},
// 		{Key: aws.String(fmt.Sprintf("%s/file2", backupDir)), LastModified: aws.Time(time.Now().Add(-2 * time.Hour))},
// 	}

// 	toDelete := []types.ObjectIdentifier{
// 		{Key: aws.String(fmt.Sprintf("%s/file2", backupDir))},
// 	}

// 	mockClient.On("ListObjectsV2", mock.Anything, &s3.ListObjectsV2Input{
// 		Bucket: aws.String(bucketName),
// 		Prefix: aws.String(backupDir),
// 	}).Return(&s3.ListObjectsV2Output{Contents: objects}, nil)

// 	mockClient.On("DeleteObjects", mock.Anything, &s3.DeleteObjectsInput{
// 		Bucket: aws.String(bucketName),
// 		Delete: &types.Delete{
// 			Objects: toDelete,
// 		},
// 	}).Return((*s3.DeleteObjectsOutput)(nil), fmt.Errorf("delete error"))

// 	err := cleaner.Clean(context.Background(), bucketName, backupDir, maxBackupCount)

// 	require.Error(t, err)
// 	require.Contains(t, err.Error(), "failure during objects deletion")
// 	mockClient.AssertExpectations(t)
// }

// func TestEnsureTrailingSlash(t *testing.T) {
// 	backupDir := "backups"
// 	res := ensureTrailingSlash(backupDir)

// 	expected := fmt.Sprint(backupDir, "/")
// 	require.Equal(t, expected, res)
// }
