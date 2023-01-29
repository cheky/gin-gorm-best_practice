package libraries

import (
	"crypto/sha256"
	"encoding/hex"
)

func GetPassword(salt, original_password string) string {
	password := salt + original_password
	h := sha256.New()
	h.Write([]byte(password))
	encript := hex.EncodeToString(h.Sum(nil))
	return encript
}
