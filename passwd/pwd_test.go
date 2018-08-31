package utils

import (
	"testing"
)

func TestMain(t *testing.T) {
	for i := 1; i < 20; i++ {
		t.Error(string(GenPwd(i)))
	}
}
