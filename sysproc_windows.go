//go:build windows
// +build windows

package main

import "os/exec"

func setSysProcAttr(cmd *exec.Cmd) {
    // Windows 下不需要设置
}
