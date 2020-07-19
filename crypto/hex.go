package crypto

import "encoding/hex"

func HexEncode(data string) string {
	return hex.EncodeToString([]byte(data))
}
