package messageBox

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

const (
	MB_OK                = 0x00000000
	MB_OKCANCEL          = 0x00000001
	MB_ABORTRETRYIGNORE  = 0x00000002
	MB_YESNOCANCEL       = 0x00000003
	MB_YESNO             = 0x00000004
	MB_RETRYCANCEL       = 0x00000005
	MB_CANCELTRYCONTINUE = 0x00000006
	MB_ICONHAND          = 0x00000010
	MB_ICONQUESTION      = 0x00000020
	MB_ICONEXCLAMATION   = 0x00000030
	MB_ICONASTERISK      = 0x00000040
	MB_USERICON          = 0x00000080
	MB_ICONWARNING       = MB_ICONEXCLAMATION
	MB_ICONERROR         = MB_ICONHAND
	MB_ICONINFORMATION   = MB_ICONASTERISK
	MB_ICONSTOP          = MB_ICONHAND

	MB_DEFBUTTON1 = 0x00000000
	MB_DEFBUTTON2 = 0x00000100
	MB_DEFBUTTON3 = 0x00000200
	MB_DEFBUTTON4 = 0x00000300
)

var user32DLL = windows.NewLazySystemDLL("user32.dll")
var procMessageBox = user32DLL.NewProc("MessageBoxW")

func main() {
	boxTitle := "TestBox"
	boxMSG := "Hello from golang"

	ptrboxTitle, err := windows.UTF16PtrFromString(boxTitle)
	if err != nil {
		panic(err.Error())
	}

	ptrboxMSG, err := windows.UTF16PtrFromString(boxMSG)
	if err != nil {
		panic(err.Error())
	}

	boxType := MB_OK

	r0, _, e1 := procMessageBox.Call(uintptr(unsafe.Pointer(nil)), uintptr(unsafe.Pointer(ptrboxMSG)), uintptr(unsafe.Pointer(ptrboxTitle)), uintptr(boxType))

	ret := int32(r0)
	if ret == 0 {
		panic((e1.Error()))

	}

	// hWnd, lpText, lpCaption, uType
	// hWnd - A handle to the owner window of the message box to be created. If this parameter is NULL, the message box has no owner window.
	// lpText - The message to be displayed.
	// lpCation - The dialog box title. If this parameter is NULL, the default title is Error.
	// uType - The contents and behavior of the dialog box.

	// Null, LPCTSTR, LPCTSTR, UINT

	// LPCTSTR - A pointer to a constant null-terminated string of 8-bit Windows (ANSI) characters.
	// UINT - Becomes unin32 go type
	// uType - MB_OK 0x00000000L The message box contains one push button: OK. This is the default.
}
