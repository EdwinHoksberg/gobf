//go:build !darwin && !freebsd && !openbsd && !netbsd

package main

import "golang.org/x/sys/unix"

const (
	getTermios = unix.TCGETS
	setTermios = unix.TCSETS
)
