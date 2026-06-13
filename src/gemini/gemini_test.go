package gemini

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Bashsoft707/tweet-audit/src/config"
	"github.com/Bashsoft707/tweet-audit/src/models"
)

func TestAnalyseTweet(t *testing.T) {

	tests := []struct {
		name string
		tweet string
		mockResponse string
		expected string
	}{
		{
        	name:         "Valid response from Gemini",
			tweet:        "This is a test tweet.",
			mockResponse: `{"candidates": [{"content": {"parts": [{"text": "FLAGGED"}]}}]}`,
			expected:     "FLAGGED",
    	},
		{
			name:         "Empty candidates from Gemini",
			tweet:        "This is a test tweet.",
			mockResponse: `{"candidates": []}`,
			expected:     "",
		},

	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprintln(w, tt.mockResponse)
			}))

			defer server.Close()

			result, err := AnalyseTweet(models.Tweet{FullText: tt.tweet}, "fake-api-key", server.URL, &config.Config{
				Criteria: config.Criteria{
					ForbiddenWords: []string{"crypto"},
					ProfessionalCheck: false,
				},
			})
			if err != nil {
				t.Fatalf("AnalyseTweet returned an error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Expected '%s', got '%s'", tt.expected, result)
			}	
		})
	}
}