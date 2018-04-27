// +build windows

package ctc

import (
	"syscall"
	"unsafe"
)

func (c Color) String() string {
	if initScreenInfo == nil { // No console info - Ex: stdout redirection
		return ""
	}
	setConsoleTextAttribute(hStdout, uint16(c))
	return ""
}

var (
	kernel32                               = syscall.NewLazyDLL("kernel32.dll")
	procGetStdHandle                       = kernel32.NewProc("GetStdHandle")
	procSetConsoleTextAttribute            = kernel32.NewProc("SetConsoleTextAttribute")
	procGetConsoleScreenBufferInfo         = kernel32.NewProc("GetConsoleScreenBufferInfo")
	stdOutputHandle                uintptr = -11 & 0xFFFFFFFF
	hStdout, _, _                          = procGetStdHandle.Call(stdOutputHandle)
	initScreenInfo                         = getConsoleScreenBufferInfo(hStdout)
)

func getConsoleScreenBufferInfo(hStdout uintptr) *consoleScreenBufferInfo {
	var csbi consoleScreenBufferInfo
	if ret, _, _ := procGetConsoleScreenBufferInfo.Call(hStdout, uintptr(unsafe.Pointer(&csbi))); ret == 0 {
		return nil
	}
	return &csbi
}

func setConsoleTextAttribute(hStdout uintptr, wAttributes uint16) bool {
	ret, _, _ := procSetConsoleTextAttribute.Call(
		hStdout,
		uintptr(wAttributes))
	return ret != 0
}

type coord struct {
	X, Y int16
}

type smallRect struct {
	Left, Top, Right, Bottom int16
}

type consoleScreenBufferInfo struct {
	DwSize              coord
	DwCursorPosition    coord
	WAttributes         uint16
	SrWindow            smallRect
	DwMaximumWindowSize coord
}