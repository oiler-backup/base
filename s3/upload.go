package s3

import (
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// A S3Uploader provides methods to upload file to s3-compatible storage.
type S3Uploader struct {
	client IS3Client
}

// NewS3Uploader is a constructor for S3Uploader.
//
// It configures and instantiates configuration and s3-client.
// endpoint is an s3-api endpoint, e.g. https://example.com:443.
// accessKey and secretKey must grant access to bucket where files will be deleted.
// region must match your aws-region or might be fictios for other solutions.
// If you want to use TLS/SSL encrytion, set secure to true.
func NewS3Uploader(ctx context.Context, endpoint, accessKey, secretKey, region string, secure bool) (S3Uploader, error) { // coverage-ignore
	client, err := NewS3Client(ctx, endpoint, accessKey, secretKey, region, secure)
	if err != nil {
		return S3Uploader{}, err
	}

	return S3Uploader{
		client: client,
	}, nil
}

// Upload uploads a single file to storage.
func (u S3Uploader) Upload(ctx context.Context, bucketName, objectKey string, fileContent io.Reader) error {
	_, err := u.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   fileContent,
	})
	if err != nil {
		return fmt.Errorf("failed to load file to S3: %w", err)
	}

	return nil
}
