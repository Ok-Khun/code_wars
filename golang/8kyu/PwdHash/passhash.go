package eightKyu

import (
	"crypto/md5"
	"encoding/hex"
)

func PassHash(text string) string {
	hashedBytes := md5.Sum([]byte(text))
	return hex.EncodeToString(hashedBytes[:])
}
