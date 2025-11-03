//go:build integration

package sample_test

import (
	"bytes"
	"context"
	"io"
	"testing"
	"time"

	"unit-test/internal/testkit"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func Test_LocalStack_S3_RoundTrip(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	cfg, err := testkit.NewAWSConfigLocal(ctx) // ca-central-1
	if err != nil {
		t.Fatal(err)
	}
	s3c := s3.NewFromConfig(cfg)

	bucket := "it-sample-bucket"
	key := "hello.txt"
	want := []byte("hello, localstack")

	// For non-us-east-1 regions, include CreateBucketConfiguration
	_, _ = s3c.CreateBucket(ctx, &s3.CreateBucketInput{
		Bucket: &bucket,
		CreateBucketConfiguration: &types.CreateBucketConfiguration{
			LocationConstraint: types.BucketLocationConstraintCaCentral1, 
		},
	})

	time.Sleep(200 * time.Millisecond)

	if _, err := s3c.PutObject(ctx, &s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &key,
		Body:   bytes.NewReader(want),
	}); err != nil {
		t.Fatalf("put object: %v", err)
	}

	out, err := s3c.GetObject(ctx, &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})
	if err != nil {
		t.Fatalf("get object: %v", err)
	}
	got, _ := io.ReadAll(out.Body)
	_ = out.Body.Close()

	if string(got) != string(want) {
		t.Fatalf("body mismatch: got=%q want=%q", got, want)
	}
}
