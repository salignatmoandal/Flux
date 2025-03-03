package analyzer

import (
	"testing"
)

func TestAnalyzeResources(t *testing.T) {
	tests := []struct {
		name     string
		provider string
		wantErr  bool
	}{
		{
			name:     "Test AWS provider",
			provider: "aws",
			wantErr:  false,
		},
		{
			name:     "Test invalid provider",
			provider: "invalid",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewResourceAnalyzer()
			_, err := a.AnalyzeResources(tt.provider, true)
			if (err != nil) != tt.wantErr {
				t.Errorf("AnalyzeResources() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
