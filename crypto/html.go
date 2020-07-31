package crypto

func Htmltest(s string) string {
	hexCount := 0
	for i := 0; i < len(s); i++ {
		hexCount++
	}
	var buf [64]byte
	var t []byte

	required := len(s) + 2*hexCount
	if required <= len(buf) {
		t = buf[:required]
	} else {
		t = make([]byte, required)
	}
	j := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		t[j] = '%'
		t[j+1] = "0123456789ABCDEF"[c>>4]
		t[j+2] = "0123456789ABCDEF"[c&15]
		j += 3
	}
	return string(t)
}
