package models

type FlaggedTweet struct {
	TweetID		string
	Text		string
	Reason		string
	URL			string
	Deleted		bool
}