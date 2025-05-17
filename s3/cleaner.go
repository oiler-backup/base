package s3

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"sort"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

// A S3Cleaner deletes files from s3-bucket according to specified policy.
type S3Cleaner struct {
	client IS3Client
}

// NewS3Cleaner is a constructor for S3Cleaner.
//
// It configures and instantiates configuration and s3-client.
// endpoint is an s3-api endpoint, e.g. https://example.com:443.
// accessKey and secretKey must grant access to bucket where files will be deleted.
// region must match your aws-region or might be fictios for other solutions.
// If you want to use TLS/SSL encrytion, set secure to true.
func NewS3Cleaner(ctx context.Context, endpoint, accessKey, secretKey, region string, secure bool) (S3Cleaner, error) { // coverage-ignore
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
		return S3Cleaner{}, err
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

	return S3Cleaner{
		client: client,
	}, nil
}

// Clean deletes oldest files to match maxBackupCount.
// backupDir might be either with or without trailing slash.
func (u S3Cleaner) Clean(ctx context.Context, bucketName, backupDir string, maxBackupCount int) error {
	listOutput, err := u.client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
		Prefix: aws.String(ensureTrailingSlash(backupDir)),
	})
	if err != nil {
		return fmt.Errorf("failed to list objects: %+v", err)
	}

	objects := listOutput.Contents

	if len(objects) > maxBackupCount {
		sort.Slice(objects, func(i, j int) bool {
			return objects[i].LastModified.Before(*objects[j].LastModified)
		})

		toDelete := objects[:len(objects)-maxBackupCount]

		deleteObjects := []types.ObjectIdentifier{}
		for _, obj := range toDelete {
			deleteObjects = append(deleteObjects, types.ObjectIdentifier{Key: obj.Key})
		}

		_, err = u.client.DeleteObjects(context.TODO(), &s3.DeleteObjectsInput{
			Bucket: aws.String(bucketName),
			Delete: &types.Delete{
				Objects: deleteObjects,
			},
		})
		if err != nil {
			return fmt.Errorf("failure during objects deletion: %+v", err)
		}
	}

	return nil
}

// ensureTrailingSlash adds trailing slash to s if it is not added yet.
func ensureTrailingSlash(s string) string {
	trailingSlash := "/"
	if !strings.HasSuffix(s, trailingSlash) {
		s = fmt.Sprint(s, trailingSlash)
	}
	return s
}
