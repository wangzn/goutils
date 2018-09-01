// @Author: wangzn04@gmail.com
// @Date: 2017-08-30 13:47:39

package os

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"

	"github.com/wangzn/goutils/myhttp"
)

// FileGetContents reads data from a FS file or HTTP endpoint
func FileGetContents(fn string) ([]byte, error) {
	if strings.HasPrefix(fn, "http://") {
		_, bs, err := myhttp.HGet(fn, nil)
		return bs, err
	}
	fp, err := os.OpenFile(fn, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return []byte(""), err
	}
	defer fp.Close()
	reader := bufio.NewReader(fp)
	return ioutil.ReadAll(reader)
}

// FilePutContents puts data to file or HTTP endpoint
func FilePutContents(fn string, con []byte) ([]byte, int, error) {
	var size int
	if strings.HasPrefix(fn, "http://") {
		_, body, err := myhttp.HPost(fn, nil, con, "")
		return body, len(body), err
	}
	fp, err := os.OpenFile(fn, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, 0, err
	}
	defer fp.Close()
	size, err = fp.Write(con)
	return nil, size, err
}

// FileAppendContents appends data to file
func FileAppendContents(fn string, con []byte) (int, error) {
	size := 0
	fp, err := os.OpenFile(fn, os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		return size, err
	}
	defer fp.Close()
	size, err = fp.Write(con)
	return size, err
}

// FExist checks if file exist in local FS
func FExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}
	return false
}

// FNExist checks if file does not exit in local FS
func FNExist(path string) bool {
	return (!FExist(path))
}
