package analyzer

import (
	"testing"

	"github.com/Bashsoft707/tweet-audit/src/config"
	"github.com/Bashsoft707/tweet-audit/src/models"
)

func TestAnalyseTweets(t *testing.T) {
	cfg := &config.Config{
		Username: "testuser",
		Criteria: config.Criteria{
			ForbiddenWords: []string{"crypto"},
		},
	}

	tests := []struct {
		name string
		tweets []models.Tweet
		expectedCount int
	}{
		{
			name: "No forbidden words",
			tweets: []models.Tweet{
				{ID: "1", FullText: "This is a normal tweet."},
				{ID: "2", FullText: "Another normal tweet."},
			},
			expectedCount: 0,
		},
		{
			name: "One forbidden word",
			tweets: []models.Tweet{
				{ID: "1", FullText: "This is a tweet about crypto."},
				{ID: "2", FullText: "This is a normal tweet."},
			},
			expectedCount: 1,
		},
		{
			name: "Multiple tweets with forbidden words",
			tweets: []models.Tweet{
				{ID: "1", FullText: "This is a tweet about crypto."},
				{ID: "2", FullText: "This is a tweet about crypto and NFTs."},
			},
			expectedCount: 2,
		},
		{
			name: "Case insensitivity",
			tweets: []models.Tweet{
				{ID: "1", FullText: "This is a tweet about Crypto."},
				{ID: "2", FullText: "This is a normal tweet."},
			},
			expectedCount: 1,
		},
		{
			name: "Empty tweet list",
			tweets: []models.Tweet{},
			expectedCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flagged := AnalyseTweets(tt.tweets, cfg)
			if len(flagged) != tt.expectedCount {
				t.Errorf("Expected %d flagged tweets, got %d", tt.expectedCount, len(flagged))
			}
		})
	}
}

