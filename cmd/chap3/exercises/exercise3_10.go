package exercises

import "bytes"

// Write a non-recursive version of comma, using bytes.Buffer instead of string concatenation.
// func Comma(s string) string {
// 	n := len(s)
//
// 	if n <= 3 {
// 		return s
// 	}
// 	return Comma(s[:n-3]) + "," + s[n-3:]
// }

func Comma(s string) string {
	var buff bytes.Buffer
	for position := range len(s) {
		if position != 0 && (len(s)-position)%3 == 0 {
			buff.WriteByte(',')
		}
		buff.WriteByte(s[position])
	}
	return buff.String()
}
