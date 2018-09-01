// @Author: wangzn04@gmail.com
// @Date: 2018-09-01 08:17:12

package utils

import (
	"math/rand"
	"time"
)

func GenPwd(l int) []byte {
	var num int
	filled := false
	seed := time.Now().UnixNano()
	res := make([]byte, l)
	rand.Seed(seed)
	m := []int{0, 0, 0, 0}
	n := [][]int{
		{48, 58},
		{65, 91},
		{97, 123},
		{33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47,
			58, 59, 60, 61, 62, 63, 64,
			91, 92, 93, 94, 95, 96,
			123, 124, 125, 126},
	}
	for i := 0; i < l; i++ {
		r := rand.Int() % len(m)
		m[r] = 1
		j := r
		if !filled && i > l/2 {
			for k := 0; k < len(m); k++ {
				if m[k] == 0 {
					m[k] = 1
					j = k
					break
				}
				if k == len(m)-1 {
					filled = true
				}
			}
		}
		if len(n[j]) == 2 {
			num = rand.Intn(n[j][1]-n[j][0]) + n[j][0]
		} else {
			num = n[j][rand.Intn(len(n[j]))]
		}
		res[i] = byte(num)
	}
	return res
}
