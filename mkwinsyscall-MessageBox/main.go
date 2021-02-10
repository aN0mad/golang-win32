package main

// Generate command to "auto gen" the win32 calls
//go:generate go run golang.org/x/sys/windows/mkwinsyscall -output zsyscall_windows.go syscalls_windows.go

// Box types as defined here: https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-messagebox
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

func main() {

	boxTitle := "TitleBar"          // Title for the message box
	boxMSG := "Hello from mkwinsys" // Message displayed in the message box

	// Only need these if we are passing the strings as a *uint16. Go will automatically convert for us if we just pass a string
	/*
		ptrboxTitle, err := windows.UTF16PtrFromString(boxTitle)
		if err != nil {
			panic(err.Error())
		}

		ptrboxMSG, err := windows.UTF16PtrFromString(boxMSG)
		if err != nil {
			panic(err.Error())
		}

	*/
	MessageBox(0, boxMSG, boxTitle, MB_OK)
}
