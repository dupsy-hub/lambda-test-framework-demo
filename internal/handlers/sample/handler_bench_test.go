package sample

import (
	"context"
	"strings"
	"testing"

	"unit-test/mocks"

	"github.com/golang/mock/gomock"
)

func BenchmarkProcess_Large(b *testing.B) {
	ctx := context.Background()
	ctrl := gomock.NewController(b)
	ms3 := mocks.NewMockS3PutObject(ctrl)
	ms3.EXPECT().PutObject(gomock.Any(), gomock.Any()).AnyTimes().Return(nil, nil)

	deps := Deps{S3: ms3, Bucket: "bench"}
	req := Request{ID: "B1", Data: strings.Repeat("x", 2<<20)} // ~2MB

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Process(ctx, deps, req)
	}
}
