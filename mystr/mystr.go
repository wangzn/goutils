// @Author: wangzn04@gmail.com
// @Date: 2018-08-13 15:33:45

package mystr

import "unicode"

// ToSnake translates a camel-case string to snake case
func ToSnake(s string) string {
	in := []rune(s)
	isLower := func(idx int) bool {
		return idx >= 0 && idx < len(in) &&
			unicode.IsLower(in[idx])
	}
	out := make([]rune, 0, len(in)+len(in)/2)
	for i, r := range in {
		if unicode.IsUpper(r) {
			r = unicode.ToLower(r)
			if i > 0 && in[i-1] != '_' && (isLower(i-1) || isLower(i+1)) {
				out = append(out, '_')
			}
		}
		out = append(out, r)
	}
	return string(out)
}

// SLower lowers the first letter of string s
// e.g.: AaeName => aeaName
func SLower(s string) string {
	if s == "" {
		return s
	}
	in := []rune(s)
	in[0] = unicode.ToLower(in[0])
	return string(in)
}

// IsUpperPrefix checks if s starts with upper letter
func IsUpperPrefix(s string) bool {
	if s == "" {
		return false
	}
	in := []rune(s)
	return unicode.IsUpper(in[0])
}
