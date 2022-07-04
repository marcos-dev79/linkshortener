package db

// This Struct representing the data to be stored
type Link struct {
	URL       string `json:"url"`
	Hash      string `json:"hash"`
	Shorthash string `json:"shorthash"`
	Created   int64  `json:"created"`
	Counter   int64  `json:"counter"`
}
