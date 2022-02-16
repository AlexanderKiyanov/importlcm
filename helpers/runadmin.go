package helpers

import (
	"fmt"
	"golang.org/x/sys/windows"
	"os"
	"strings"
	"syscall"
)

func AmIAdmin() bool {
	// if not elevated, relaunch by shell execute with runas verb set
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	if err != nil {
		fmt.Println("Am I admin? No.\n\n")
		return false
	}
	fmt.Print("Am I admin? Yes.\n\n")

	return true
}

// Switch to administrator
func ElevateAsAdmin() {
	verb := "runas"
	exe, _ := os.Executable()
	cwd, _ := os.Getwd()
	args := strings.Join(os.Args[1:], " ")

	fmt.Println("Arguments: ", os.Args[1:])

	verbPtr, _ := syscall.UTF16PtrFromString(verb)
	exePtr, _ := syscall.UTF16PtrFromString(exe)
	cwdPtr, _ := syscall.UTF16PtrFromString(cwd)
	argPtr, _ := syscall.UTF16PtrFromString(args)

	var showCmd int32 = 1 //SW_NORMAL

	err := windows.ShellExecute(0, verbPtr, exePtr, argPtr, cwdPtr, showCmd)
	if err != nil {
		fmt.Println(err)
	}
}
