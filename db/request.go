package db

// This struct represents the payload to be posted to the shortener link
type Request struct {
	URL string `json:"url"`
}
