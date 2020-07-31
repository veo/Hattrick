package crypto

func Unescape(data string) string {
	return "<script>document.write(unescape('" + UrlEncode_GBK(data) + "'));</script>"
}

func UnescapeAll(data string) string {
	return "<script>document.write(unescape('" + UrlEncode_ALL(data) + "'));</script>"
}
