package s3

import (
	"context"
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Downloader struct {
	client IS3Client
}

func NewS3Downloader(ctx context.Context, endpoint, accessKey, secretKey, region string, secure bool) (S3Downloader, error) {
	client, err := NewS3Client(ctx, endpoint, accessKey, secretKey, region, secure)
	if err != nil {
		return S3Downloader{}, err
	}

	return S3Downloader{
		client: client,
	}, nil
}

// downloadBackupFromS3 скачивает выбранный бэкап из S3 в локальный файл
func (d S3Downloader) Download(ctx context.Context, bucketName, databaseName, backupRevisionStr, localFilePath string) error {
	var selectedBackupKey string

	backupRevision, err := strconv.Atoi(backupRevisionStr)
	// Если backupRevisionStr - число, выбираем бэкап по индексу
	if err == nil && backupRevision >= 0 {
		selectedBackupKey, err = d.GetBackupByRevision(ctx, backupRevision, databaseName, bucketName)
		if err != nil {
			return fmt.Errorf("failed to list backup files from S3: %v", err)
		}
	} else {
		// Если backupRevisionStr - строка, ищем бэкап с таким именем
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

	file, err := os.Create(localFilePath)
	if err != nil {
		return fmt.Errorf("failed to create local file: %v", err)
	}
	defer file.Close()

	_, err = file.ReadFrom(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write S3 object to local file: %v", err)
	}

	return nil
}

// listBackupFiles получает список файлов бэкапов из S3
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
