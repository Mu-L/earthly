// +build !windows

package pllb

import (
	"syscall"
)

func getInodeBestEffort(path string) uint64 {
	var stat syscall.Stat_t
	inode := uint64(0)
	if err := syscall.Stat(path, &stat); err == nil {
		inode = uint64(stat.Ino)
	}
	return inode
}
