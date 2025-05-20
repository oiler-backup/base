package s3

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

const (
	partSize = 5 * 1024 * 1024
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
	createOutput, err := u.client.CreateMultipartUpload(ctx, &s3.CreateMultipartUploadInput{
		Bucket: &bucketName,
		Key:    &objectKey,
	})
	if err != nil {
		return err
	}

	uploadID := createOutput.UploadId
	var wg sync.WaitGroup
	var mu sync.Mutex
	completedParts := []types.CompletedPart{}

	buffer := make([]byte, partSize)
	partNumber := 1

	for {
		bytesRead, err := fileContent.Read(buffer)
		if bytesRead == 0 || err == io.EOF {
			break
		}
		if err != nil && err != io.EOF {
			fmt.Println("Ошибка чтения файла:", err)
			return err
		}

		wg.Add(1)
		go func(partNumber int32, data []byte, bytesRead int) {
			defer wg.Done()

			partInput := &s3.UploadPartInput{
				Bucket:     aws.String(bucketName),
				Key:        aws.String(objectKey),
				PartNumber: &partNumber,
				UploadId:   uploadID,
				Body:       bytes.NewReader(data[:bytesRead]),
			}
			partResp, err := u.client.UploadPart(ctx, partInput)
			if err != nil {
				fmt.Printf("Ошибка загрузки части %d: %v\n", partNumber, err)
				return
			}

			mu.Lock()
			completedParts = append(completedParts, types.CompletedPart{
				ETag:       partResp.ETag,
				PartNumber: &partNumber,
			})
			mu.Unlock()
		}(int32(partNumber), buffer, bytesRead)

		partNumber++
	}

	wg.Wait()

	completeInput := &s3.CompleteMultipartUploadInput{
		Bucket:   aws.String(bucketName),
		Key:      aws.String(objectKey),
		UploadId: uploadID,
		MultipartUpload: &types.CompletedMultipartUpload{
			Parts: completedParts,
		},
	}
	_, err = u.client.CompleteMultipartUpload(ctx, completeInput)
	if err != nil {
		fmt.Println("Ошибка завершения multipart upload:", err)
		return err
	}
	// _, err := u.client.PutObject(ctx, &s3.PutObjectInput{
	// 	Bucket: aws.String(bucketName),
	// 	Key:    aws.String(objectKey),
	// 	Body:   fileContent,
	// })
	// if err != nil {
	// 	return fmt.Errorf("failed to load file to S3: %w", err)
	// }

	return nil
}
