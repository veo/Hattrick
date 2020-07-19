package crypto

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

func SHA1_Encode(data string) string {
	h := sha1.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func SHA256_Encode(data string) string {
	h := sha256.New()
	return hex.EncodeToString(h.Sum(nil))
}

func SHA512_Encode(data string) string {
	h := sha512.New()
	return hex.EncodeToString(h.Sum(nil))
}
