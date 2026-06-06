package main

import (
	"github.com/Bashsoft707/tweet-audit/src/analyzer"
	"github.com/Bashsoft707/tweet-audit/src/archive"
	"github.com/Bashsoft707/tweet-audit/src/config"
	"github.com/Bashsoft707/tweet-audit/src/report"

	"fmt"
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
	
	// Prints all Tweets
	if len(tweets) > 0 {
		for _, tweet := range tweets {
			fmt.Printf("Tweet ID: %s\nText: %s\nCreated At: %s\n\n", tweet.ID, tweet.FullText, tweet.CreatedAt)
		}
	}

	flaggedTweets := analyzer.AnalyseTweets(tweets, cfg)

	for _, tweet := range flaggedTweets {
		fmt.Printf(
			"Tweet: %s\nReason: %s\nURL: %s\n\n",
			tweet.Text,
			tweet.Reason,
			tweet.URL,
		)
	}

	err = report.WriteCSV(cfg.OutputPath, flaggedTweets)

	if err != nil {
		panic(err)
	}
	
	fmt.Printf("CSV report generated successfully")
}