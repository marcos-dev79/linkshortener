package lib

import (
	"crypto/sha1"
	"encoding/hex"
	"linkshortener/db"
	"time"
)

// This method generates the hash
func ShortenLink(url db.Request) db.Link {

	now := time.Now()
	sec := now.Unix()
	nsec := now.UnixNano()
	hash := sha1.Sum([]byte(url.URL + string(rune(nsec))))
	shorthash := hash[0:8]

	var link db.Link
	link.Created = sec
	link.Hash = hex.EncodeToString(hash[:])
	link.Shorthash = hex.EncodeToString(shorthash[:])
	link.URL = url.URL

	return link

}
