package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		want    string
		wantErr bool
	}{
		{
			name: "Malformed Authorization header",
			headers: http.Header{
				"Authorization": []string{"InValidApiKey token"},
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Valid Header",
			headers: http.Header{
				"Authorization": []string{"ApiKey valid_token"},
			},
			want:    "valid_token",
			wantErr: false,
		},
		{
			name:    "No Header",
			headers: http.Header{},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotToken, err := GetAPIKey(tt.headers)
			if (err != nil) != tt.wantErr {
				t.Errorf("name: %v, wantErr %v", tt.name, err)
				return
			}

			if gotToken != tt.want {
				t.Errorf("name: %v, wantErr %v", tt.name, err)
				return
			}
		})
	}
}
