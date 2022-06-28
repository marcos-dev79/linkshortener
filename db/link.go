package db

type Link struct {
	URL       string `json:"url"`
	Hash      string `json:"hash"`
	Shorthash string `json:"shorthash"`
	Created   int64  `json:"created"`
}
