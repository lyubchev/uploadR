package main

import (
	"log"
	"os"
	"runtime"
)

var pathToMCMap = make(map[string]string)

func main() {

	pathToMCMap["windows"] = os.Getenv("APPDATA") + "\\.minecraft"
	pathToMCMap["linux"] = "~/.minecraft"
	pathToMCMap["macOS"] = "~/Library/Application Support/minecraft"

	pathToMC := pathToMCMap[runtime.GOOS]

	exists, err := dirExists(pathToMC)

	if err != nil {
		log.Fatal(err)
		return
	}

	if !exists {
		println("Minecraft directory not found, quitting")
		return
	} else {
		println("Minecraft directory found, uploading your pack. Please wait...")
	}
}

func dirExists(path string) (bool, error) {
	_, err := os.Stat(path)

	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return true, err
}
