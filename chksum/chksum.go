// Copyright 2017 TED@Sogou, Inc. All rights reserved.
//
// @Author: wangzhongning@sogou-inc.com
// @Date: 2017-09-07 15:24:01

package chksum

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
)

// SHA1 returns a sha1 checksum of bytes
func SHA1(bs []byte) string {
	h := sha1.New()
	h.Write(bs)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// SHA256 returns a sha256 checksum of bytes
func SHA256(bs []byte) string {
	sum := sha256.Sum256(bs)
	return fmt.Sprintf("%x", sum)
}

// MD5 returns md5 checksum of bytes
func MD5(bs []byte) string {
	return fmt.Sprintf("%x", md5.Sum(bs))
}
