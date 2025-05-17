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

type S3Cleaner struct {
	client IS3Client
}

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

func ensureTrailingSlash(s string) string {
	trailingSlash := "/"
	if !strings.HasSuffix(s, trailingSlash) {
		s = fmt.Sprint(s, trailingSlash)
	}
	return s
}
