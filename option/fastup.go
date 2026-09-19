package option

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

const fastupPasswordSuffix = "#fastup"

// NormalizeFastupPassword removes the subscription marker and derives the
// Trojan password expected by Fastup servers. Ordinary passwords are unchanged.
func NormalizeFastupPassword(password, mpw string) (string, bool) {
	if !strings.HasSuffix(password, fastupPasswordSuffix) {
		return password, false
	}
	password = strings.TrimSuffix(password, fastupPasswordSuffix)
	if mpw == "" {
		mpw = string([]byte{110, 121, 97, 50, 48, 50, 52, 49, 50, 48, 57})
	}
	digest := md5.Sum([]byte(password + mpw))
	return hex.EncodeToString(digest[:]), true
}
