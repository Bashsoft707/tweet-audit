package models

type Tweet struct {
    ID        string `json:"id_str"`
    FullText  string `json:"full_text"`
    CreatedAt string `json:"created_at"`
}

type TweetWrapper struct {
	Tweet Tweet `json:"tweet"`
}