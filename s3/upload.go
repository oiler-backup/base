package s3

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Uploader struct {
	client *s3.Client
}

func NewS3Uploader(ctx context.Context, endpoint, accessKey, secretKey, region string, secure bool) (S3Uploader, error) {
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(region),
		config.WithCredentialsProvider(aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
			return aws.Credentials{
				AccessKeyID:     accessKey,
				SecretAccessKey: secretKey,
			}, nil
		})),
	)
	if err != nil {
		return S3Uploader{}, err
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
		o.BaseEndpoint = aws.String(endpoint)
		if !secure {
			o.HTTPClient = &http.Client{
				Transport: &http.Transport{
					TLSClientConfig: &tls.Config{
						InsecureSkipVerify: true,
					},
				},
			}
		}
	})

	// opts := "d"
	return S3Uploader{
		client: client,
	}, nil
}

func (u S3Uploader) Upload(ctx context.Context, bucketName, filePath, objectKey string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failure during opening file: %w", err)
	}
	defer file.Close()

	_, err = u.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   file,
	})
	if err != nil {
		return fmt.Errorf("failed to load file to S3: %w", err)
	}

	return nil
}
