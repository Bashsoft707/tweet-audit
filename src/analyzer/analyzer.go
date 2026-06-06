package analyzer

import (
	"strings"

	"github.com/Bashsoft707/tweet-audit/src/config"
	"github.com/Bashsoft707/tweet-audit/src/models"
)

func AnalyseTweets(tweets []models.Tweet, cfg *config.Config ) []models.FlaggedTweet {
	forbiddenWords := cfg.Criteria.ForbiddenWords
	username := cfg.Username

	var flagged []models.FlaggedTweet

	for _, tweet := range tweets {
		text := strings.ToLower(tweet.FullText)

		for _, word := range forbiddenWords {
			if strings.Contains(text, word) {
				flaggedTweet := models.FlaggedTweet{
					TweetID: tweet.ID,
					Text: tweet.FullText,
					Reason: "contains forbidden word: " + word,
					URL: "https://x.com/" + username + "/status/" + tweet.ID,
					Deleted: false,
				}

				flagged = append(flagged, flaggedTweet)
				break
			}
		}
	}

	return flagged
}