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

	if exists, err := minecraftExists(pathToMC); !exists || err != nil {
		return
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

func minecraftExists(path string) (bool, error) {
	exists, err := dirExists(path)

	if err != nil {
		log.Fatal(err)
		return true, err
	}

	if !exists {
		println("Minecraft directory not found, quitting")
		return false, err
	}

	return true, nil
}
