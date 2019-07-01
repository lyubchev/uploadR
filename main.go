package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
)

var pathToMCMap = make(map[string]string)

func main() {

	pathToMCMap["windows"] = os.Getenv("APPDATA") + "\\.minecraft"
	pathToMCMap["linux"] = "~/.minecraft"
	pathToMCMap["darwin"] = "~/Library/Application Support/minecraft"

	pathToMC := pathToMCMap[runtime.GOOS]

	if exists, err := minecraftExists(pathToMC); !exists || err != nil {
		return
	}

	file, err := os.Open(pathToMC + "\\options.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buf := make([]byte, 32*1024) // define your buffer size here.

	for {
		n, err := file.Read(buf)

		if n > 0 {
			fmt.Printl((buf[:n]) // your read buffer.
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("read %d bytes: %v", n, err)
			break
		}
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

func check(e error) {
	if e != nil {
		panic(e)
	}
}
