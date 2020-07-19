package crypto

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5_Encode_32(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func MD5_Encode_16(data string) string {
	return MD5_Encode_32(data)[8:24]
}
