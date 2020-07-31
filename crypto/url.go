package crypto

import (
	"bytes"
	"encoding/base64"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/url"
)

func UrlEncode_UTF8(data string) string {
	return url.PathEscape(data)
}

func UrlEncode_GBK(data string) string {
	encoder := simplifiedchinese.GBK.NewEncoder()
	reader := transform.NewReader(bytes.NewReader([]byte(data)), encoder)
	d, _ := ioutil.ReadAll(reader)
	return url.PathEscape(string(d))
}

func UrlEncode_ALL(data string) string {
	return urlencodeall(data)
}

func UrlDecode_UTF8(data string) string {
	m, _ := base64.StdEncoding.DecodeString(data)
	return string(m)
}

func UrlDecode_GBK(data string) string {
	m, _ := base64.StdEncoding.DecodeString(data)
	decoder := simplifiedchinese.GBK.NewDecoder()
	reader := transform.NewReader(bytes.NewReader(m), decoder)
	d, _ := ioutil.ReadAll(reader)
	return string(d)
}

func urlencodeall(s string) string {
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
