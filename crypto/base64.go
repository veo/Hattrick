package crypto

import (
	"bytes"
	"encoding/base64"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
)

func Base64Encode_UTF8(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func Base64Encode_GBK(data string) string {
	encoder := simplifiedchinese.GBK.NewEncoder()
	reader := transform.NewReader(bytes.NewReader([]byte(data)), encoder)
	d, err := ioutil.ReadAll(reader)
	if err != nil {
		return err.Error()
	}
	return base64.StdEncoding.EncodeToString(d)
}

func Base64Decode_UTF8(data string) string {
	m, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return err.Error()
	}
	return string(m)
}

func Base64Decode_GBK(data string) string {
	m, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return err.Error()
	}
	decoder := simplifiedchinese.GBK.NewDecoder()
	reader := transform.NewReader(bytes.NewReader(m), decoder)
	d, err := ioutil.ReadAll(reader)
	if err != nil {
		return err.Error()
	}
	return string(d)
}
