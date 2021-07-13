// +build windows
package main

import (
	"fmt"
	"os"
	"syscall"
)

var gProc *syscall.LazyProc

func initConsoleColor() {
	gProc = syscall.NewLazyDLL("kernel32.dll").NewProc("SetConsoleTextAttribute")
}

func log(aFormat string, aParms ...interface{}) {
	gProc.Call(uintptr(syscall.Stdout), 7)
	str := aFormat + "\n"
	fmt.Fprintf(os.Stdout, str, aParms...)
}

func logErr(aFormat string, aParms ...interface{}) {
	gProc.Call(uintptr(syscall.Stdout), 12) //red 12
	str := "[ERROR] " + aFormat + "\n"
	fmt.Fprintf(os.Stdout, str, aParms...)
	gProc.Call(uintptr(syscall.Stdout), 7) //white 12
}
