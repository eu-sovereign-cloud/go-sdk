package gosdk_test

import (
	"context"
	"errors"
	"io"
	"testing"

	gosdk "github.com/eu-sovereign-cloud/go-sdk"
)

func TestIterator_Next(t *testing.T) {
	tests := []struct {
		name       string
		fn         gosdk.IteratorFunc[int]
		wantValues []int
		wantErr    error
	}{
		{
			name: "successful iteration",
			fn: func(ctx context.Context, skipToken *string) ([]int, *string, error) {
				if skipToken == nil {
					return []int{1, 2, 3}, nil, nil
				}
				return nil, nil, io.EOF
			},
			wantValues: []int{1, 2, 3},
			wantErr:    io.EOF,
		},
		{
			name: "error during iteration",
			fn: func(ctx context.Context, skipToken *string) ([]int, *string, error) {
				return nil, nil, errors.New("fetch error")
			},
			wantValues: nil,
			wantErr:    errors.New("fetch error"),
		},
		{
			name: "single call to fn",
			fn: func(ctx context.Context, skipToken *string) ([]int, *string, error) {
				if skipToken == nil {
					return []int{1}, nil, nil
				}
				return nil, nil, io.EOF
			},
			wantValues: []int{1},
			wantErr:    io.EOF,
		},
		{
			name: "two calls to fn",
			fn: func(ctx context.Context, skipToken *string) ([]int, *string, error) {
				if skipToken == nil {
					token := "next"
					return []int{1, 2}, &token, nil
				}
				return []int{3, 4}, nil, nil
			},
			wantValues: []int{1, 2, 3, 4},
			wantErr:    io.EOF,
		},
		{
			name: "three calls to fn",
			fn: func(ctx context.Context, skipToken *string) ([]int, *string, error) {
				if skipToken == nil {
					token := "next1"
					return []int{1, 2}, &token, nil
				}
				if *skipToken == "next1" {
					token := "next2"
					return []int{3, 4}, &token, nil
				}
				return []int{5, 6}, nil, nil
			},
			wantValues: []int{1, 2, 3, 4, 5, 6},
			wantErr:    io.EOF,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iterator := gosdk.NewIterator(tt.fn)
			var gotValues []int
			for {
				val, err := iterator.Next(context.Background())
				if err == io.EOF {
					break
				}
				if err != nil {
					if err.Error() != tt.wantErr.Error() {
						t.Errorf("Iterator.Next() error = %v, wantErr %v", err, tt.wantErr)
					}
					return
				}
				gotValues = append(gotValues, *val)
			}
			if len(gotValues) != len(tt.wantValues) {
				t.Errorf("Iterator.Next() gotValues = %v, want %v", gotValues, tt.wantValues)
			}
			for i, v := range gotValues {
				if v != tt.wantValues[i] {
					t.Errorf("Iterator.Next() gotValues[%d] = %v, want %v", i, v, tt.wantValues[i])
				}
			}
		})
	}
}
