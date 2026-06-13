# Why I Chose Go

I chose Go for this project because I have wanted to learn the language deeply for a long time, and this project provides a practical opportunity to build something real while improving my backend engineering skills. Its simplicity, strong concurrency model, and suitability for CLI and data-processing applications also made it a good fit for this tool.

## Why I Used Test Data Instead of a Real X Archive

I used simulated test data during development because downloading my X archive failed repeatedly and its time-consuming. Using local test data allowed me to continue building and learning Go without being blocked by external dependencies.

This approach also made development faster because I could fully control the dataset structure, quickly test edge cases, and repeatedly run the parser without waiting for archive downloads.

## Why Keyword Analysis Runs Before Gemini

The pipeline runs a local keyword pass first, before sending any tweets to the Gemini API. Tweets flagged by keywords are tracked in a `flaggedIDs` map and skipped in the Gemini loop. This avoids redundant API calls for tweets that are already confirmed as flagged, reducing cost and latency. The keyword pass is free and instant; the Gemini pass costs API quota and time.

## Why `baseURL` Is a Parameter in `AnalyseTweet`

The Gemini API URL is passed into `AnalyseTweet` as a parameter rather than hardcoded inside the function. This makes the function testable without hitting the real Gemini API. Tests inject a mock URL using `httptest.NewServer`, which spins up a local HTTP server that returns controlled responses. If the URL were hardcoded, tests would depend on network access and a valid API key, making them slow, flaky, and expensive.

## Why `time.Sleep` Between Gemini API Calls

The Gemini free tier enforces a per-minute request limit. Without a delay between calls, rapid successive requests exhaust the quota immediately and every call returns a `429 RESOURCE_EXHAUSTED` error. A 4-second sleep between requests keeps throughput at roughly 15 requests per minute, which stays within the free-tier rate limit.