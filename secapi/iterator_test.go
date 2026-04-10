package secapi_test

import (
	"context"
	"errors"
	"io"
	"testing"

	"github.com/eu-sovereign-cloud/go-sdk/pkg/spec/schema"
	"github.com/eu-sovereign-cloud/go-sdk/secapi"
)

func TestIterator_Next(t *testing.T) {
	tests := []struct {
		name       string
		fn         secapi.IteratorFunc[schema.Region]
		wantValues []schema.Region
		wantErr    error
	}{
		{
			name: "successful iteration",
			fn: func(ctx context.Context, skipToken *string) ([]schema.Region, *schema.ResponseMetadata, error) {
				if skipToken == nil {
					return []schema.Region{{Metadata: &schema.GlobalResourceMetadata{Name: "region-1"}},
						{Metadata: &schema.GlobalResourceMetadata{Name: "region-2"}},
						{Metadata: &schema.GlobalResourceMetadata{Name: "region-3"}},
					}, &schema.ResponseMetadata{SkipToken: skipToken}, nil
				}
				return nil, nil, io.EOF
			},
			wantValues: []schema.Region{{Metadata: &schema.GlobalResourceMetadata{Name: "region-1"}},
				{Metadata: &schema.GlobalResourceMetadata{Name: "region-2"}},
				{Metadata: &schema.GlobalResourceMetadata{Name: "region-3"}},
			}, wantErr: io.EOF,
		},
		{
			name: "error during iteration",
			fn: func(ctx context.Context, skipToken *string) ([]schema.Region, *schema.ResponseMetadata, error) {
				return nil, nil, errors.New("fetch error")
			},
			wantValues: nil,
			wantErr:    errors.New("fetch error"),
		},
		{
			name: "single call to fn",
			fn: func(ctx context.Context, skipToken *string) ([]schema.Region, *schema.ResponseMetadata, error) {
				if skipToken == nil {
					return []schema.Region{{Metadata: &schema.GlobalResourceMetadata{Name: "region-1"}}}, &schema.ResponseMetadata{SkipToken: skipToken}, nil
				}
				return nil, nil, io.EOF
			},
			wantValues: []schema.Region{{Metadata: &schema.GlobalResourceMetadata{Name: "region-1"}}},
			wantErr:    io.EOF,
		},
		{
			name: "two calls to fn",
			fn: func(ctx context.Context, skipToken *string) ([]schema.Region, *schema.ResponseMetadata, error) {
				if skipToken == nil {
					token := "next"
					return []schema.Region{{Metadata: &schema.GlobalResourceMetadata{Name: "region-1"}},
						{Metadata: &schema.GlobalResourceMetadata{Name: "region-2"}},
					}, &schema.ResponseMetadata{SkipToken: &token}, nil
				}
				return []schema.Region{{Metadata: &schema.GlobalResourceMetadata{Name: "region-3"}},
					{Metadata: &schema.GlobalResourceMetadata{Name: "region-4"}},
				}, &schema.ResponseMetadata{SkipToken: nil}, nil
			},
			wantValues: []schema.Region{{Metadata: &schema.GlobalResourceMetadata{Name: "region-1"}},
				{Metadata: &schema.GlobalResourceMetadata{Name: "region-2"}},
				{Metadata: &schema.GlobalResourceMetadata{Name: "region-3"}},
				{Metadata: &schema.GlobalResourceMetadata{Name: "region-4"}},
			},
			wantErr: io.EOF,
		},
		{
			name: "three calls to fn",
			fn: func(ctx context.Context, skipToken *string) ([]schema.Region, *schema.ResponseMetadata, error) {
				if skipToken == nil {
					token := "next1"
					return []schema.Region{{Metadata: &schema.GlobalResourceMetadata{Name: "region-1"}},
						{Metadata: &schema.GlobalResourceMetadata{Name: "region-2"}},
					}, &schema.ResponseMetadata{SkipToken: &token}, nil
				}
				if *skipToken == "next1" {
					token := "next2"
					return []schema.Region{{Metadata: &schema.GlobalResourceMetadata{Name: "region-3"}},
						{Metadata: &schema.GlobalResourceMetadata{Name: "region-4"}},
					}, &schema.ResponseMetadata{SkipToken: &token}, nil
				}
				return []schema.Region{{Metadata: &schema.GlobalResourceMetadata{Name: "region-5"}},
					{Metadata: &schema.GlobalResourceMetadata{Name: "region-6"}},
				}, &schema.ResponseMetadata{SkipToken: nil}, nil
			},

			wantValues: []schema.Region{{Metadata: &schema.GlobalResourceMetadata{Name: "region-1"}},
				{Metadata: &schema.GlobalResourceMetadata{Name: "region-2"}},
				{Metadata: &schema.GlobalResourceMetadata{Name: "region-3"}},
				{Metadata: &schema.GlobalResourceMetadata{Name: "region-4"}},
				{Metadata: &schema.GlobalResourceMetadata{Name: "region-5"}},
				{Metadata: &schema.GlobalResourceMetadata{Name: "region-6"}},
			},
			wantErr: io.EOF,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iterator := secapi.NewIterator(tt.fn)
			var gotValues []schema.Region
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
			for i, val := range gotValues {
				if val.Metadata.Name != tt.wantValues[i].Metadata.Name {
					t.Errorf("Iterator.Next() gotValues[%d] = %v, want %v", i, val, tt.wantValues[i])
				}
			}
		})
	}
}
