package crypto

import "fmt"

func HtmlEncode_Dec(data string) string {
	var result string
	array := []rune(data)
	n := len(array)
	for i := 0; i < n; i++ {
		result += fmt.Sprint("&#", array[i]) + ";"
	}
	return result
}

func HtmlEncode_Hex(data string) string {
	var result string
	array := []rune(data)
	n := len(array)
	for i := 0; i < n; i++ {
		result += "&#x" + fmt.Sprintf("%X", array[i]) + ";"
	}
	return result
}
