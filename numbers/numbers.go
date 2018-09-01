// @Author: wangzn04@gmail.com
// @Date: 2017-08-11 17:52:26

package numbers

import "strconv"

// StringMustInt must return a int from string with default 0
func StringMustInt(s string) int {
	return StringMustIntDefault(s, 0)
}

// StringMustIntDefault must return a int from string with supplied default
// value
func StringMustIntDefault(s string, d int) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return d
	}
	return i
}

// StringMustFloat64 must return a float64 from string with default 0.0
func StringMustFloat64(s string) float64 {
	return StringMustFloat64Default(s, 0)
}

// StringMustFloat64Default must return a float64 from string with supplied
// default value
func StringMustFloat64Default(s string, d float64) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return d
	}
	return f
}
