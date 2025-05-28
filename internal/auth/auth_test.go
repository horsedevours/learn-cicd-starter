package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	testCases := []struct {
		name      string
		input     string
		want      string
		wantError bool
	}{
		{
			name:      "all good",
			input:     "ApiKey 123456",
			want:      "123456",
			wantError: false,
		}, {
			name:      "bad",
			input:     "Bearer123456",
			want:      "",
			wantError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			headers := http.Header{}
			headers.Set("Authorization", tc.input)

			got, err := GetAPIKey(headers)
			if (err != nil) && !tc.wantError {
				t.Errorf("unexpected error: %v", err)
			}
			if got != tc.want {
				t.Errorf("got %s, want %s", got, tc.want)
			}
		})
	}
}
