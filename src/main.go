package main

import (
	"github.com/Bashsoft707/tweet-audit/src/analyzer"
	"github.com/Bashsoft707/tweet-audit/src/archive"
	"github.com/Bashsoft707/tweet-audit/src/config"
	"github.com/Bashsoft707/tweet-audit/src/gemini"
	"github.com/Bashsoft707/tweet-audit/src/models"
	"github.com/Bashsoft707/tweet-audit/src/report"

	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println("Tweet Audit Started")

	cfg, err := config.LoadConfig("config.json")

	if err != nil {
		panic(err)
	}

	fmt.Println("Config Loaded Successfully")
	data, err := archive.ReadArchive(cfg.ArchivePath)
	if err != nil {
		panic(err)
	}

	fmt.Println("Archive Read Successful:", cfg.ArchivePath)

	tweets, err := archive.ParseArchive(data)

	if err != nil {
		panic(err)
	}

	flaggedTweets := analyzer.AnalyseTweets(tweets, cfg)
	flaggedIDs := make(map[string]bool)

	for _, tweet := range flaggedTweets {
		flaggedIDs[tweet.TweetID] = true
		// fmt.Printf(
		// 	"Tweet: %s\nReason: %s\nURL: %s\n\n",
		// 	tweet.Text,
		// 	tweet.Reason,
		// 	tweet.URL,
		// )
	}


	baseURL := "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash:generateContent"

	count := 0
	for _, tweet := range tweets {
		// test first 20 tweets with Gemini
		if count >= 20 {
			break
		}
		if flaggedIDs[tweet.ID] {
			continue
		}

		result, err := gemini.AnalyseTweet(tweet, cfg.GeminiAPIKey, baseURL, cfg)
		count++
		time.Sleep(4 * time.Second)

		if err != nil {
			fmt.Printf("Warning: failed to analyse tweet %s: %v\n", tweet.ID, err)
    		continue
		}
		if strings.HasPrefix(result, "FLAGGED") {
			fmt.Printf("Tweet ID: %s flagged by Gemini with reason: %s\n", tweet.ID, result)
			flaggedTweet := models.FlaggedTweet{
				TweetID: tweet.ID,
				Text: tweet.FullText,
				Reason: "Flagged by Gemini: " + result,
				URL: "https://x.com/" + cfg.Username + "/status/" + tweet.ID,
				Deleted: false,
			}
			flaggedTweets = append(flaggedTweets, flaggedTweet)
		}
	}
	
	err = report.WriteCSV(cfg.OutputPath, flaggedTweets)

	if err != nil {
		panic(err)
	}
	
	fmt.Printf("CSV report generated successfully")
}