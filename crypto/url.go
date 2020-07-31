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
	return url.QueryEscape(data)
}

func UrlEncode_GBK(data string) string {
	encoder := simplifiedchinese.GBK.NewEncoder()
	reader := transform.NewReader(bytes.NewReader([]byte(data)), encoder)
	d, _ := ioutil.ReadAll(reader)
	return url.QueryEscape(string(d))
}

func UrlEncode_ALL(data string) string {
	return url.PathEscape("123")
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
