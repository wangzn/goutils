// Copyright 2018 TED@Sogou, Inc. All rights reserved.

// @Author: wangzhongning@sogou-inc.com
// @Date: 2018-08-22 15:07:21

package flock

import (
	"os"
	"strconv"
	"syscall"
)

// Flock defines the struct to lock with file
type Flock struct {
	fh *os.File
	fn string
}

// New returns a new Flock pointer
func New(fn string) *Flock {
	return &Flock{
		fn: fn,
	}
}

func (f *Flock) setFh() error {
	fh, err := os.OpenFile(f.fn, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	f.fh = fh
	return nil
}

// Lock tries to lock, block frontend
func (f *Flock) Lock() error {
	if f.fh == nil {
		if err := f.setFh(); err != nil {
			return err
		}
	}

	if err := syscall.Flock(int(f.fh.Fd()), syscall.LOCK_EX); err != nil {
		return err
	}

	f.WritePid()

	return nil
}

// TryLock tries to lock, returns error if flock is blocked by someone else
func (f *Flock) TryLock() (bool, error) {
	if f.fh == nil {
		if err := f.setFh(); err != nil {
			return false, err
		}
	}

	err := syscall.Flock(int(f.fh.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)

	switch err {
	case syscall.EWOULDBLOCK:
		return false, nil

	case nil:
		f.WritePid()
		return true, nil
	}

	return false, err
}

// WritePid writes a pid to file
func (f *Flock) WritePid() {
	f.fh.Truncate(0)
	f.fh.Write([]byte(strconv.Itoa(os.Getpid())))
}

// Unlock release lock
func (f *Flock) Unlock() error {
	if err := syscall.Flock(int(f.fh.Fd()), syscall.LOCK_UN); err != nil {
		return err
	}

	f.fh.Close()
	f.fh = nil

	return nil
}

// WaitLock waits lock
func WaitLock(fn string) (*Flock, error) {
	wl := New(fn + ".wait")
	ok, err := wl.TryLock()
	if !ok {
		return nil, err
	}

	defer wl.Unlock()

	fl := New(fn)

	if err := fl.Lock(); err != nil {
		return nil, err
	}

	return fl, nil
}
