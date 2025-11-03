package sample

import (
	"context"
	"errors"
	"testing"

	"unit-test/mocks"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestProcess(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	const bucket = "test-bucket"

	tests := []struct {
		name      string
		req       Request
		mockS3Err error
		wantErr   bool
	}{
		{"ok", Request{ID: "123", Data: "hello"}, nil, false},
		{"missing id", Request{Data: "no-id"}, nil, true},
		{"s3 failure", Request{ID: "999", Data: "x"}, errors.New("boom"), true},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ms3 := mocks.NewMockS3PutObject(ctrl)

			// Only expect PutObject when validation passes (ID present).
			if tc.req.ID != "" {
				ms3.EXPECT().
					PutObject(gomock.Any(), gomock.AssignableToTypeOf(&s3.PutObjectInput{})).
					Return(nil, tc.mockS3Err)
			}

			deps := Deps{S3: ms3, Bucket: bucket}
			got, err := Process(ctx, deps, tc.req)

			if tc.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, "ok", got)
		})
	}
}
