package crypto

import (
	"bytes"
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
	d, err := ioutil.ReadAll(reader)
	if err != nil {
		return err.Error()
	}
	return url.PathEscape(string(d))
}

func UrlEncode_ALL(data string) string {
	encoder := simplifiedchinese.GBK.NewEncoder()
	reader := transform.NewReader(bytes.NewReader([]byte(data)), encoder)
	d, err := ioutil.ReadAll(reader)
	if err != nil {
		return err.Error()
	}
	return urlencodeall(string(d))
}

func UrlDecode_UTF8(data string) string {
	s, err := url.PathUnescape(data)
	if err != nil {
		return err.Error()
	} else {
		return s
	}
}

func UrlDecode_GBK(data string) string {
	s, err := url.PathUnescape(data)
	if err != nil {
		return err.Error()
	}
	decoder := simplifiedchinese.GBK.NewDecoder()
	reader := transform.NewReader(bytes.NewReader([]byte(s)), decoder)
	d, err := ioutil.ReadAll(reader)
	if err != nil {
		return err.Error()
	}
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
