package s3

import (
	"context"
	"fmt"
	"io"
)

type S3UploadCleaner struct {
	u IS3Uploader
	c IS3Cleaner
}

func NewS3UploadCleaner(ctx context.Context, endpoint, accessKey, secretKey, region string, secure bool) (S3UploadCleaner, error) { // coverage-ignore
	uploader, err := NewS3Uploader(ctx, endpoint, accessKey, secretKey, region, secure)
	if err != nil {
		return S3UploadCleaner{}, fmt.Errorf("failed to initialize uploader: %+v", err)
	}
	cleaner, err := NewS3Cleaner(ctx, endpoint, accessKey, secretKey, region, secure)
	if err != nil {
		return S3UploadCleaner{}, fmt.Errorf("failed to initialize cleaner: %+v", err)
	}

	return S3UploadCleaner{
		u: uploader,
		c: cleaner,
	}, nil
}

func (uc S3UploadCleaner) CleanAndUpload(ctx context.Context, bucketName, backupDir string, maxBackupCount int, fileName string, fileContent io.Reader) error {
	err := uc.c.Clean(ctx, bucketName, backupDir, maxBackupCount)
	if err != nil {
		return fmt.Errorf("failed to clean S3: %+v", err)
	}
	err = uc.u.Upload(ctx, bucketName, fileName, fileContent)
	if err != nil {
		return fmt.Errorf("failed to upload object to S3: %+v", err)
	}

	return nil
}
