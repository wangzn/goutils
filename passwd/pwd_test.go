// @Author: wangzn04@gmail.com
// @Date: 2018-09-01 08:17:00

package utils

import (
	"testing"
)

func TestMain(t *testing.T) {
	for i := 1; i < 20; i++ {
		t.Error(string(GenPwd(i)))
	}
}
