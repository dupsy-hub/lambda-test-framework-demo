package awsiface

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3PutObject interface {
	PutObject(ctx context.Context, in *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error)
}

// This line makes sure the real S3 client actually implements the interface.
var _ S3PutObject = (*s3.Client)(nil)
