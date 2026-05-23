package report

import (
	"os"
	"encoding/csv"
	"fmt"
	"github.com/Bashsoft707/tweet-audit/src/models"
)

func WriteCSV(filename string, tweets []models.FlaggedTweet) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("create csv file: %w", err)
	}

	defer file.Close()

	csvWriter := csv.NewWriter(file)

	err = csvWriter.Write([]string{"tweet_url", "deleted"})
	if err != nil {
		return fmt.Errorf("write csv header: %w", err)
	}

	for _, tweet := range tweets {
		csvWriter.Write([]string{tweet.URL, "false"})

		if err != nil {
			return fmt.Errorf("write csv row: %w", err)
		}
	}

	csvWriter.Flush()

	if err := csvWriter.Error(); err != nil {
		return fmt.Errorf("flush csv writer: %w", err)
	}

	return nil
}