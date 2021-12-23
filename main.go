package main

import (
	"fmt"
	"runtime"
)

// Infos are set at build time use ldflags.
var (
	Version   string
	Revision  string
	BuildUser string
	BuildDate string
	GoVersion = runtime.Version()
)

func main() {
	fmt.Printf("构建的Go版本: %s\n", GoVersion)
	fmt.Printf("应用当前版本: %s\n", Version)
	fmt.Printf("应用当前提交: %s\n", Revision)
	fmt.Printf("应用构建时间: %s\n", BuildDate)
	fmt.Printf("应用构建用户: %s\n", BuildUser)
	fmt.Println("Hello world!")
}
