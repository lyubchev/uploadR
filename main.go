package main

import (
	"fmt"
	"runtime"
)

var envPathMap = make(map[string]string)

func main() {

	envPathMap["windows"] = "%APPDATA%\\.minecraft"
	envPathMap["linux"] = "~/.minecraft"
	envPathMap["macOS"] = "~/Library/Application Support/minecraft"

	fmt.Println(runtime.GOOS)
}
