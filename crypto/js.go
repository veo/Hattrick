package crypto

import (
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
)

func FromChar(data string) string {
	var result string
	array := []rune(data)
	n := len(array)
	for i := 0; i < n; i++ {
		if i == 0 {
			result += fmt.Sprint(array[i])
		} else {
			result += fmt.Sprint(",", array[i])
		}
	}
	return "String.fromCharCode(" + result + ")"
}

func Unicode(data string) string {
	var result string
	array := []rune(data)
	n := len(array)
	for i := 0; i < n; i++ {
		result += "\\u" + fmt.Sprintf("%X", array[i])
	}
	return result
}

func JsHex(data string) string {
	encoder := simplifiedchinese.GBK.NewEncoder()
	reader := transform.NewReader(bytes.NewReader([]byte(data)), encoder)
	d, err := ioutil.ReadAll(reader)
	if err != nil {
		return err.Error()
	}
	var result string
	n := len(d)
	for i := 0; i < n; i++ {
		result += "\\x" + fmt.Sprintf("%X", d[i])
	}
	return result
}

func JsUnicode(data string) string {
	var result string
	array := []rune(data)
	n := len(array)
	for i := 0; i < n; i++ {
		if len(fmt.Sprintf("%X", array[i])) == 2 {
			result += "\\u00" + fmt.Sprintf("%X", array[i])
		} else {
			result += "\\u" + fmt.Sprintf("%X", array[i])
		}

	}
	return result
}
