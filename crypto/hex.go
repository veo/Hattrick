package crypto

import "encoding/hex"

func HexEncode(data string) string {
	return "0x" + hex.EncodeToString([]byte(data))
}
