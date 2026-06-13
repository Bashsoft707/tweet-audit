package archive

import "testing"

func TestParseArchive(t *testing.T) {
	tests := []struct {
		name string
		input []byte
		expected int
		expectError bool
		expectedID string
		expectedText string
	}{
		{
			name: "Valid archive data",
			input: []byte(`window.YTD.tweets.part0 = [{"tweet": {"id_str": "123", "full_text": "Hello world!"}}]`),
			expected: 1,
			expectError: false,
			expectedID: "123",
			expectedText: "Hello world!",
		},
		{
			name: "Invalid JSON format",
			input: []byte(`window.YTD.tweets.part0 = [{"tweet": {"id_str": "123", "full_text": "Hello world!"}}`),
			expected: 0,
			expectError: true,
			expectedID: "",
			expectedText: "",
		},
		{
			name: "Empty archive data",
			input: []byte(`window.YTD.tweets.part0 = []`),
			expected: 0,
			expectError: false,
			expectedID: "",
			expectedText: "",
		},
		{
			name: "Valid JSON without JS prefix",
			input: []byte(`[{"tweet": {"id_str": "123", "full_text": "Raw JSON"}}]`),
			expected: 1,
			expectError: false,
			expectedID: "123",
			expectedText: "Raw JSON",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tweets, err := ParseArchive(tt.input)
		
			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Fatalf("Unexpected error: %v", err)
				}
				if len(tweets) != tt.expected {
					t.Errorf("Expected %d tweets, got %d", tt.expected, len(tweets))
				}
				if tt.expectedID != "" && tweets[0].ID != tt.expectedID {
					t.Errorf("Expected ID '%s', got '%s'", tt.expectedID, tweets[0].ID)
				}
				if tt.expectedText != "" && tweets[0].FullText != tt.expectedText {
					t.Errorf("Expected text '%s', got '%s'", tt.expectedText, tweets[0].FullText)
				}
			}
		})
	}
}