package s3

import (
	"context"
	"fmt"
	"io"
	"sort"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// S3Downloader represents a downloader for files stored in AWS S3.
type S3Downloader struct {
	client IS3Client
}

// NewS3Downloader creates and returns a new instance of S3Downloader.
// It initializes the underlying AWS S3 client with the provided configuration.
func NewS3Downloader(ctx context.Context, endpoint, accessKey, secretKey, region string, secure bool) (S3Downloader, error) {
	client, err := NewS3Client(ctx, endpoint, accessKey, secretKey, region, secure)
	if err != nil {
		return S3Downloader{}, err
	}

	return S3Downloader{
		client: client,
	}, nil
}

// Download downloads the specified backup from S3 to a local file.
// If backupRevisionStr is a number, it selects the backup by index.
// If backupRevisionStr is a string, it looks for a backup with that name.
func (d S3Downloader) Download(ctx context.Context, bucketName, databaseName, backupRevisionStr string, fileContent io.WriteCloser) error {
	defer fileContent.Close()
	var selectedBackupKey string

	backupRevision, err := strconv.Atoi(backupRevisionStr)
	if err == nil && backupRevision >= 0 {
		selectedBackupKey, err = d.GetBackupByRevision(ctx, backupRevision, databaseName, bucketName)
		if err != nil {
			return fmt.Errorf("failed to list backup files from S3: %v", err)
		}
	} else {
		selectedBackupKey = backupRevisionStr
	}

	resp, err := d.client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(selectedBackupKey),
	})
	if err != nil {
		return fmt.Errorf("failed to get S3 object: %v", err)
	}
	defer resp.Body.Close()

	buffer := make([]byte, partSize)

	for {
		bytesRead, _ := resp.Body.Read(buffer)
		if bytesRead == 0 {
			break
		}

		_, err := fileContent.Write(buffer[:bytesRead])
		if err != nil {
			return fmt.Errorf("failed to write S3 object to file: %v", err)
		}
	}

	return nil
}

// GetBackupByRevision retrieves the key of the backup file at the specified revision index.
// It lists all backup files in the specified directory and sorts them by modification time in descending order.
func (d S3Downloader) GetBackupByRevision(ctx context.Context, backupRevision int, backupDir, bucketName string) (string, error) {
	listOutput, err := d.client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
		Prefix: aws.String(ensureTrailingSlash(backupDir)),
	})
	if err != nil {
		return "", fmt.Errorf("failed to list objects: %+v", err)
	}

	objects := listOutput.Contents
	var selectedBackupKey string
	if backupRevision < len(objects) {
		sort.Slice(objects, func(i, j int) bool {
			return objects[j].LastModified.Before(*objects[i].LastModified)
		})
		selectedBackupKey = *objects[backupRevision].Key
	} else {
		return "", fmt.Errorf("BACKUP_REVISION (%d) is out of range. Available backups: %d", backupRevision, len(objects))
	}

	return selectedBackupKey, nil
}
