package client

import (
	"context"
	"errors"
	io "io"
	"strings"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	s3Types "github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	gomock "go.uber.org/mock/gomock"
)

func TestPresignPutObject(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPresignApi := NewMockS3PresignApi(ctrl)
	client := &s3ClientImpl{
		presignApi: mockPresignApi,
	}

	ctx := context.Background()
	bucket := "test-bucket"
	key := "test-key"
	expiration := 15 * time.Minute
	expectedURL := "https://example.com/presigned-url"

	mockPresignApi.EXPECT().PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}, gomock.Any()).Return(&v4.PresignedHTTPRequest{
		URL: expectedURL,
	}, nil)

	url, err := client.PresignPutObject(ctx, bucket, key, expiration)
	assert.NoError(t, err)
	assert.Equal(t, expectedURL, url)
}

func TestPresignPutObject_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPresignApi := NewMockS3PresignApi(ctrl)
	client := &s3ClientImpl{
		presignApi: mockPresignApi,
	}

	ctx := context.Background()
	bucket := "test-bucket"
	key := "test-key"
	expiration := 15 * time.Minute
	expectedError := errors.New("presign error")

	mockPresignApi.EXPECT().PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}, gomock.Any()).Return(nil, expectedError)

	url, err := client.PresignPutObject(ctx, bucket, key, expiration)
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Empty(t, url)
}

func s3Setup(t *testing.T) (*gomock.Controller, *s3ClientImpl, context.Context) {
	ctrl := gomock.NewController(t)
	mockApi := NewMockS3Api(ctrl)
	client := &s3ClientImpl{
		api: mockApi,
	}
	ctx := context.Background()
	return ctrl, client, ctx
}

func TestWaitForObject_Success(t *testing.T) {
	t.Parallel()
	ctrl, client, ctx := s3Setup(t)
	defer ctrl.Finish()

	bucket := "test-bucket"
	key := "test-key"
	timeout := 2 * time.Second
	expectedBody := "test content"

	mockApi := client.api.(*MockS3Api)
	mockApi.EXPECT().HeadObject(gomock.Any(), &s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}, gomock.Any()).Return(&s3.HeadObjectOutput{}, nil)

	mockApi.EXPECT().GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}, gomock.Any()).Return(&s3.GetObjectOutput{
		Body: io.NopCloser(strings.NewReader(expectedBody)),
	}, nil)

	body, err := client.WaitForObject(ctx, bucket, key, timeout)
	require.NoError(t, err)

	defer body.Close()
	content, err := io.ReadAll(body)
	require.NoError(t, err)
	assert.Equal(t, expectedBody, string(content))
}

func TestWaitForObject_Timeout(t *testing.T) {
	t.Parallel()
	ctrl, client, ctx := s3Setup(t)
	defer ctrl.Finish()

	bucket := "test-bucket"
	key := "test-key"
	timeout := 2 * time.Second

	mockApi := client.api.(*MockS3Api)
	mockApi.EXPECT().HeadObject(gomock.Any(), &s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}, gomock.Any()).Return(nil, &s3Types.NotFound{}).AnyTimes()

	body, err := client.WaitForObject(ctx, bucket, key, timeout)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed waiting for object")
	assert.Nil(t, body)
}

func TestWaitForObject_Deleted(t *testing.T) {
	t.Parallel()
	ctrl, client, ctx := s3Setup(t)
	defer ctrl.Finish()

	bucket := "test-bucket"
	key := "test-key"
	timeout := 2 * time.Second
	expectedError := &s3Types.NoSuchKey{}

	mockApi := client.api.(*MockS3Api)
	mockApi.EXPECT().HeadObject(gomock.Any(), &s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}, gomock.Any()).Return(&s3.HeadObjectOutput{}, nil)

	mockApi.EXPECT().GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}, gomock.Any()).Return(nil, expectedError)

	body, err := client.WaitForObject(ctx, bucket, key, timeout)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "was deleted before it was retrieved")
	assert.Nil(t, body)
}

func TestWaitForObject_GetObjectError(t *testing.T) {
	t.Parallel()
	ctrl, client, ctx := s3Setup(t)
	defer ctrl.Finish()

	bucket := "test-bucket"
	key := "test-key"
	timeout := 2 * time.Second
	expectedError := errors.New("get object error")

	mockApi := client.api.(*MockS3Api)
	mockApi.EXPECT().HeadObject(gomock.Any(), &s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}, gomock.Any()).Return(&s3.HeadObjectOutput{}, nil)

	mockApi.EXPECT().GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}, gomock.Any()).Return(nil, expectedError)

	body, err := client.WaitForObject(ctx, bucket, key, timeout)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to get object")
	assert.Nil(t, body)
}
