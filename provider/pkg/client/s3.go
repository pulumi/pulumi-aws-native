// Copyright 2016-2024, Pulumi Corporation.

package client

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	signerV4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	s3Types "github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/pkg/errors"
)

//go:generate mockgen -package client -source s3.go -destination mock_s3.go S3Client,S3Api,S3PresignApi
type S3Client interface {
	// PresignPutObject generates a pre-signed URL for uploading an object to an S3 bucket.
	// The URL will be valid for the specified expiration duration.
	PresignPutObject(ctx context.Context, bucket, key string, expiration time.Duration) (string, error)

	// WaitForObject waits for an object to exist in an S3 bucket and returns a reader for the object.
	// The function will block until the object exists or the timeout is reached.
	// If the object does not exist after the timeout, an error will be returned.
	// The caller is responsible for closing the reader when done.
	WaitForObject(ctx context.Context, bucket string, key string, timeout time.Duration) (io.ReadCloser, error)
}

type S3Api interface {
	HeadObject(context.Context, *s3.HeadObjectInput, ...func(*s3.Options)) (*s3.HeadObjectOutput, error)
	GetObject(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error)
}

type S3PresignApi interface {
	PresignPutObject(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.PresignOptions)) (*signerV4.PresignedHTTPRequest, error)
}

type s3ClientImpl struct {
	api        S3Api
	presignApi S3PresignApi
}

func NewS3Client(api S3Api, presignApi S3PresignApi) S3Client {
	return &s3ClientImpl{
		api:        api,
		presignApi: presignApi,
	}
}

func (c *s3ClientImpl) PresignPutObject(ctx context.Context, bucket string, key string, expiration time.Duration) (string, error) {
	request, err := c.presignApi.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}, func(options *s3.PresignOptions) {
		options.Expires = expiration
	})
	if err != nil {
		return "", err
	}
	return request.URL, nil
}

func (c *s3ClientImpl) WaitForObject(ctx context.Context, bucket string, key string, timeout time.Duration) (io.ReadCloser, error) {
	err := s3.NewObjectExistsWaiter(c.api, func(o *s3.ObjectExistsWaiterOptions) {
		// We already aggressively retry SDK errors with plenty retry attempts and
		// throttling exceptions will be taken care of by the SDK
		o.MinDelay = time.Second
		o.MaxDelay = 5 * time.Second
	}).Wait(ctx, &s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}, timeout)

	if err != nil {
		return nil, errors.Wrapf(err, "failed waiting for object %q in bucket %q", key, bucket)
	}

	getResponse, err := c.api.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	if err != nil {
		var noKey *s3Types.NoSuchKey
		if errors.As(err, &noKey) {
			// this could happen if the response object got deleted before we can read it
			return nil, fmt.Errorf("object %q in bucket %q was deleted before it was retrieved", key, bucket)
		} else {
			return nil, errors.Wrapf(err, "failed to get object %q from bucket %q", key, bucket)
		}
	}

	return getResponse.Body, nil
}
