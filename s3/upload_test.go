package s3

// import (
// 	"bytes"
// 	"context"
// 	"errors"
// 	"io"
// 	"testing"

// 	"github.com/aws/aws-sdk-go-v2/aws"
// 	"github.com/aws/aws-sdk-go-v2/service/s3"
// 	"github.com/stretchr/testify/mock"
// 	"github.com/stretchr/testify/require"
// )

// func TestUpload_Success(t *testing.T) {
// 	mockClient := new(MockS3Client)
// 	uploader := S3Uploader{client: mockClient}

// 	bucketName := "my-bucket"
// 	objectKey := "my-file.txt"
// 	fileContent := bytes.NewReader([]byte("test content"))

// 	mockClient.On("PutObject", mock.Anything, &s3.PutObjectInput{
// 		Bucket: aws.String(bucketName),
// 		Key:    aws.String(objectKey),
// 		Body:   fileContent,
// 	}).Return(&s3.PutObjectOutput{}, nil)

// 	err := uploader.Upload(context.Background(), bucketName, objectKey, fileContent)

// 	require.NoError(t, err)
// 	mockClient.AssertExpectations(t)
// }

// func TestUpload_Failure(t *testing.T) {
// 	mockClient := new(MockS3Client)
// 	uploader := S3Uploader{client: mockClient}

// 	bucketName := "my-bucket"
// 	objectKey := "my-file.txt"
// 	fileContent := io.NopCloser(bytes.NewReader([]byte("test content")))

// 	expectedErr := errors.New("aws error")

// 	mockClient.On("PutObject", mock.Anything, &s3.PutObjectInput{
// 		Bucket: aws.String(bucketName),
// 		Key:    aws.String(objectKey),
// 		Body:   fileContent,
// 	}).Return((*s3.PutObjectOutput)(nil), expectedErr)

// 	err := uploader.Upload(context.Background(), bucketName, objectKey, fileContent)

// 	require.Error(t, err)
// 	require.Contains(t, err.Error(), "failed to load file to S3")
// 	mockClient.AssertExpectations(t)
// }
