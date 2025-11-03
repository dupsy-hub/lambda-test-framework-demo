package sample

import (
	"bytes" // ✅ add this
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"unit-test/internal/awsiface"
)

// Deps defines external dependencies injected into the function.
type Deps struct {
	S3     awsiface.S3PutObject
	Bucket string
}

type Request struct {
	ID   string `json:"id"`
	Data string `json:"data"`
}

// Process validates input, marshals it, and writes to S3.
// It uses the injected S3 interface so it’s easy to mock in tests.
func Process(ctx context.Context, d Deps, r Request) (string, error) {
	if r.ID == "" {
		return "", errors.New("missing id")
	}
	b, err := json.Marshal(r)
	if err != nil {
		return "", fmt.Errorf("marshal: %w", err)
	}

	_, err = d.S3.PutObject(ctx, &s3.PutObjectInput{
		Bucket: &d.Bucket,
		Key:    ptr(fmt.Sprintf("items/%s.json", r.ID)),
		Body:   bytes.NewReader(b),
	})
	if err != nil {
		return "", fmt.Errorf("s3 put: %w", err)
	}
	return "ok", nil
}

// Tiny generic pointer helper
func ptr[T any](v T) *T { return &v }
