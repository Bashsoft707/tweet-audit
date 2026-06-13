package archive

import (
	"encoding/json"

	"strings"

	"github.com/Bashsoft707/tweet-audit/src/models"
)

func ParseArchive(data []byte) ([]models.Tweet, error) {
	// Convert bytes to string
	content := string(data)

	// Remove Twitter JS wrapper
	content = strings.TrimPrefix(content, "window.YTD.tweets.part0 = ")

	// Create wrappers
	var wrappers []models.TweetWrapper

	// Parse JSON
	err := json.Unmarshal([]byte(content), &wrappers)
	if err != nil {
		return nil, err
	}

	// Create tweets array, extract tweet from wrapper and append to tweets array
	var tweets []models.Tweet
	for _, wrapper := range wrappers {
		tweets = append(tweets, wrapper.Tweet)
	}

	return tweets, nil
}