package test

import (
	"designmode/creativeDesign/builder"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuilder(t *testing.T) {
	tests := []struct {
		Name    string
		builder *builder.ResourcePoolConfigBuilder
		want    *builder.ResourcePoolConfig
		wantErr bool
	}{
		{
			Name: "Name empty",
			builder: &builder.ResourcePoolConfigBuilder{
				Name:     "",
				MaxTotal: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			Name: "MaxIdle < MinIdle",
			builder: &builder.ResourcePoolConfigBuilder{
				Name:     "test",
				MaxTotal: 0,
				MaxIdle:  10,
				MinIdle:  20,
			},
			want:    nil,
			wantErr: true,
		},
		{
			Name: "success",
			builder: &builder.ResourcePoolConfigBuilder{
				Name: "test",
			},
			want: &builder.ResourcePoolConfig{
				Name:     "test",
				MaxTotal: builder.DefaultMaxTotal,
				MaxIdle:  builder.DefaultMaxIdle,
				MinIdle:  builder.DefaultMinIdle,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			got, err := tt.builder.Build()
			require.Equalf(t, tt.wantErr, err != nil, "Build() error = %v, wantErr %v", err, tt.wantErr)
			assert.Equal(t, tt.want, got)

			fmt.Printf("%s\n", tt.Name)
		})
	}

	fmt.Printf("Builder end ...\n")

}
