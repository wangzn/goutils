// @Author: wangzn04@gmail.com
// @Date: 2018-08-13 15:34:43

package mystr

import "testing"

func TestToSnake(t *testing.T) {
	ss := []string{
		"A",
		"a",
		"Snake",
		"snake",
		"SnakeTest",
		"snake_test",
		"SnakeID",
		"snake_id",
		"LinuxMOTD",
		"linux_motd",
		"OMGWTFBBQ",
		"omgwtfbbq",
	}
	for i, s := range ss {
		if i%2 == 1 {
			continue
		}
		v := ToSnake(s)
		if v != ss[i+1] {
			t.Errorf("Invalid snake case: Got: %s Expect: %s",
				v, ss[i+1])
		}
	}
}
