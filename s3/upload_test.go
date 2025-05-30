package s3

import (
	"bytes"
	"context"
	"errors"
	"io"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestUpload_Success(t *testing.T) {
	mockClient := new(MockS3Client)
	uploader := S3Uploader{client: mockClient}

	bucketName := "my-bucket"
	objectKey := "my-file.txt"
	fileContent := bytes.NewReader([]byte("test content"))
	uploadId := "x"
	eTag := "tag"
	mockClient.On("CreateMultipartUpload", mock.Anything, mock.Anything).Return(&s3.CreateMultipartUploadOutput{UploadId: &uploadId}, nil)
	mockClient.On("UploadPart", mock.Anything, mock.Anything).Return(&s3.UploadPartOutput{
		ETag: &eTag,
	}, nil)
	mockClient.On("CompleteMultipartUpload", mock.Anything, mock.Anything).Return(&s3.CompleteMultipartUploadOutput{}, nil)

	err := uploader.Upload(context.Background(), bucketName, objectKey, fileContent)

	require.NoError(t, err)
	mockClient.AssertExpectations(t)
}

func TestUpload_Failure(t *testing.T) {
	mockClient := new(MockS3Client)
	uploader := S3Uploader{client: mockClient}

	bucketName := "my-bucket"
	objectKey := "my-file.txt"
	fileContent := io.NopCloser(bytes.NewReader([]byte("test content")))

	expectedErr := errors.New("aws error")

	mockClient.On("CreateMultipartUpload", mock.Anything, mock.Anything).Return(&s3.CreateMultipartUploadOutput{}, expectedErr)

	err := uploader.Upload(context.Background(), bucketName, objectKey, fileContent)

	require.Error(t, err)
	require.Contains(t, err.Error(), "aws error")
	mockClient.AssertExpectations(t)
}
