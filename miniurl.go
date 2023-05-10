// Package miniurl provides buiding blocks for url shortener app.
package miniurl

import (
	"crypto/md5"
	"encoding/hex"
)

var i int

// Hash generates 32 byte long hex encoded string.
func Hash(input string) string {
	// i++
	// if i == 10000 {
	// 	return ""
	// }
	hash := md5.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}
