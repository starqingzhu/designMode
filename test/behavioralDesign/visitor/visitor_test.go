package test

import (
	"designmode/behavioralDesign/visitor"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCommpressVisit(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		wantErr string
	}{
		{
			name: "pdf",
			path: "./xx.pdf",
		},
		{
			name: "ppt",
			path: "./xx.ppt",
		},
		{
			name:    "404",
			path:    "./xx.xx",
			wantErr: "not found file type",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, err := visitor.NewResourceFile(tt.path)
			if tt.wantErr != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.wantErr)
				return
			}

			require.NoError(t, err)
			compressor := &visitor.Compressor{}
			f.Accept(compressor)
		})
	}
}

// 不用 Accept 其实也是可以的
func TestCompressorVisit2(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		wantErr string
	}{
		{
			name: "pdf",
			path: "./xx.pdf",
		},
		{
			name: "ppt",
			path: "./xx.ppt",
		},
		{
			name:    "404",
			path:    "./xx.xx",
			wantErr: "not found file type",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, err := visitor.NewResourceFile(tt.path)
			if tt.wantErr != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.wantErr)
				return
			}

			require.NoError(t, err)
			compressor := &visitor.Compressor{}
			compressor.Visit(f)
		})
	}
}
