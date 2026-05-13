package main

import (
	"github.com/Bashsoft707/tweet-audit/src/config"

	"fmt"
)

func main() {
	fmt.Println("Tweet Audit Started")

	cfg, err := config.LoadConfig("config.json")

	if err != nil {
		panic(err)
	}

	fmt.Println("Config Loaded Successfully")
	fmt.Println("Archive path:", cfg.ArchivePath)
}