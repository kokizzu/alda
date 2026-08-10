//go:build windows

package system

import "syscall"

// sysProcAttr returns the syscall.SysProcAttr to use when spawning player
// processes.
//
// Windows does not have POSIX process groups, and Ctrl+C is handled
// differently there, so no detachment is needed.
func sysProcAttr() *syscall.SysProcAttr {
	return nil
}
