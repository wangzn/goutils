// @Author: wangzn04@gmail.com
// @Date: 2018-08-14 15:04:33

package rpmver

// http://blog.jasonantman.com/2014/07/how-yum-and-rpm-compare-versions/
// https://github.com/rpm-software-management/rpm/blob/master/lib/rpmvercmp.c

import (
	"strings"
	"unicode"
)

// Compare compares two rpm version
func Compare(v1, v2 string) int {
	if v1 == v2 {
		return 0
	}

	b1 := []byte(v1)
	b2 := []byte(v2)

	for {
		var nb1 []byte
		var nb2 []byte

		for i, bb := range b1 {
			r := rune(bb)
			if bb == '~' || unicode.IsLetter(r) || unicode.IsNumber(r) {
				nb1 = b1[i:]
				break
			}
		}
		b1 = nb1

		for i, bb := range b2 {
			r := rune(bb)
			if bb == '~' || unicode.IsLetter(r) || unicode.IsNumber(r) {
				nb2 = b2[i:]
				break
			}
		}
		b2 = nb2

		tilde1 := len(b1) > 0 && b1[0] == '~'
		tilde2 := len(b2) > 0 && b2[0] == '~'

		if tilde1 || tilde2 {
			if !tilde1 {
				return 1
			}
			if !tilde2 {
				return -1
			}

			b1 = b1[1:]
			b2 = b2[1:]

			continue
		}

		if len(b1) == 0 || len(b2) == 0 {
			break
		}

		var isNum bool

		nb1 = nil
		nb2 = nil

		if unicode.IsNumber(rune(b1[0])) {
			for i, bb := range b1 {
				if !unicode.IsNumber(rune(bb)) {
					nb1 = b1[:i]
					b1 = b1[i:]
					break
				}
			}

			for i, bb := range b2 {
				if !unicode.IsNumber(rune(bb)) {
					nb2 = b2[:i]
					b2 = b2[i:]
					break
				}
			}

			isNum = true
		} else {
			for i, bb := range b1 {
				if !unicode.IsLetter(rune(bb)) {
					nb1 = b1[:i]
					b1 = b1[i:]
					break
				}
			}
			for i, bb := range b2 {
				if !unicode.IsLetter(rune(bb)) {
					nb2 = b2[:i]
					b2 = b2[i:]
					break
				}
			}
		}

		if nb1 == nil {
			nb1 = b1
			b1 = nil
		}
		if nb2 == nil {
			nb2 = b2
			b2 = nil
		}

		if len(nb1) == 0 {
			return -1
		}

		if len(nb2) == 0 {
			if isNum {
				return 1
			}

			return -1
		}

		if isNum {
			for len(nb1) > 0 && nb1[0] == '0' {
				nb1 = nb1[1:]
			}
			for len(nb2) > 0 && nb2[0] == '0' {
				nb2 = nb2[1:]
			}
			if len(nb1) > len(nb2) {
				return 1
			}
			if len(nb2) > len(nb1) {
				return -1
			}
		}

		rc := strings.Compare(string(nb1), string(nb2))

		if rc != 0 {
			return rc
		}
	}

	if len(b1) == 0 && len(b2) == 0 {
		return 0
	}

	if len(b1) == 0 {
		return -1
	}

	return 1
}
