package msgbox

import (
	"syscall"
	"unsafe"
)

var (
	libuser32  = syscall.NewLazyDLL("user32.dll")
	messageBox = libuser32.NewProc("MessageBoxW")
)

type style int

const (
	OK                style = 0x00000000
	OKCANCEL          style = 0x00000001
	ABORTRETRYIGNORE  style = 0x00000002
	YESNOCANCEL       style = 0x00000003
	YESNO             style = 0x00000004
	RETRYCANCEL       style = 0x00000005
	CANCELTRYCONTINUE style = 0x00000006
	ICONHAND          style = 0x00000010
	ICONQUESTION      style = 0x00000020
	ICONEXCLAMATION   style = 0x00000030
	ICONASTERISK      style = 0x00000040
	USERICON          style = 0x00000080
	ICONWARNING       style = ICONEXCLAMATION
	ICONERROR         style = ICONHAND
	ICONINFORMATION   style = ICONASTERISK
	ICONSTOP          style = ICONHAND
	DEFBUTTON1        style = 0x00000000
	DEFBUTTON2        style = 0x00000100
	DEFBUTTON3        style = 0x00000200
	DEFBUTTON4        style = 0x00000300
)

func Show(hwnd syscall.Handle, message, title string, style style) int {
	r1, _, _ := messageBox.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(message))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))),
		uintptr(uint32(style)))
	return int(r1)
}
