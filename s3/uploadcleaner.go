package s3

import (
	"context"
	"fmt"
	"io"
)

// A S3UploadCleaner provides methods to clean the storage before
// uploading a file.
type S3UploadCleaner struct {
	u IS3Uploader
	c IS3Cleaner
}

// NewS3UploadCleaner is a constructor for S3UploadCleaner.
//
// It configures and instantiates configuration and s3-client.
// endpoint is an s3-api endpoint, e.g. https://example.com:443.
// accessKey and secretKey must grant access to bucket where files will be deleted.
// region must match your aws-region or might be fictios for other solutions.
// If you want to use TLS/SSL encrytion, set secure to true.
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

// CleanAndUpload cleans storage before uploading a file.
// Refer to [S3Uploader] and [S3Cleaner] for more information
func (uc S3UploadCleaner) CleanAndUpload(ctx context.Context, bucketName, backupDir string, maxBackupCount int, fileName string, fileContent io.Reader) error {
	err := uc.u.Upload(ctx, bucketName, fileName, fileContent)
	if err != nil {
		return fmt.Errorf("failed to upload object to S3: %+v", err)
	}

	err = uc.c.Clean(ctx, bucketName, backupDir, maxBackupCount)
	if err != nil {
		return fmt.Errorf("failed to clean S3: %+v", err)
	}

	return nil
}
