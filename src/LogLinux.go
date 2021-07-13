// +build linux

package main

import (
	"fmt"
	"os"
)

func initConsoleColor() {
}

func log(aFormat string, aParms ...interface{}) {
	str := aFormat + "\n"
	fmt.Fprintf(os.Stdout, str, aParms...)
}

func logErr(aFormat string, aParms ...interface{}) {
	str := "[ERROR] " + aFormat + "\n"
	fmt.Fprintf(os.Stdout, str, aParms...)
}
