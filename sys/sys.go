// @Author: wangzn04@gmail.com
// @Date: 2018-08-13 15:32:05

package sys

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"
)

// GetSN returns sn from 'sogou-agent', "" if err
func GetSN() string {
	out, err := exec.Command("/usr/local/sbin/sogou-agent", "sn").Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(out))
}

// GetOS returns os relase version, "" if err
func GetOS() string {
	def := ""
	out, err := exec.Command("/usr/bin/lsb_release", "-d").Output()
	if err != nil {
		return def
	}
	dat := string(out)
	if strings.HasPrefix(dat, "Description:") {
		return strings.TrimSpace(dat[12:])
	}
	return def
}

// Chmod tries to chmod of a regular file
func Chmod(dst, mod string) error {
	m, err := strconv.Atoi(mod)
	if err != nil {
		return fmt.Errorf("can't conv mode to int: %s", err.Error())
	}
	err = os.Chmod(dst, os.FileMode(m))
	if err != nil {
		return fmt.Errorf("fail to chmod, target: %d, err: %s", m, err.Error())
	}
	return nil
}

// Chown tries to chown of a regular file
func Chown(dst, own string) error {
	uid := 0
	gid := 0
	usr := ""
	grp := ""
	parts := strings.Split(own, ".")
	if len(parts) < 2 {
		// try to use : as delimeter
		parts = strings.Split(own, ":")
		if len(parts) == 1 {
			return fmt.Errorf("invalid owner format, could not get user and group")
		}
	}
	usr = parts[0]
	grp = parts[1]
	u, err := user.Lookup(usr)
	if err != nil {
		return fmt.Errorf("invalid user %s, err: %s", usr, err.Error())
	}
	g, err := user.LookupGroup(grp)
	if err != nil {
		return fmt.Errorf("invalid group %s, err: %s", grp, err.Error())
	}
	if u.Gid != g.Gid {
		return fmt.Errorf("user's gid does not equal to group's gid, u.gid: %s"+
			"g.gid: %s", u.Gid, g.Gid)
	}
	uid, err = strconv.Atoi(u.Uid)
	if err != nil {
		return fmt.Errorf("uid %s is not a number, err: %s", u.Uid, err.Error())
	}
	gid, err = strconv.Atoi(u.Gid)
	if err != nil {
		return fmt.Errorf("gid %s is not a number, err: %s", u.Gid, err.Error())
	}
	err = os.Chown(dst, uid, gid)
	if err != nil {
		return fmt.Errorf("fail to chown, uid: %d, gid: %d, err: %s", uid, gid,
			err.Error())
	}
	return nil
}
