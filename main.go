package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/SoMuchForSubtlety/fileupload"
)

var pathToMCMap = map[string]string{}
var resourcePack string

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

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentLine := scanner.Text()

		if strings.HasPrefix(currentLine, "resourcePacks:") {
			rpNameToBeParsed := strings.Split(currentLine, ":")[1]
			rpNames := strings.Split(rpNameToBeParsed[2:len(rpNameToBeParsed)-2], "\",\"")

			for _, rp := range rpNames {
				currentResourcePackPath := pathToMC + "\\resourcepacks\\" + rp

				file, err := os.Open(currentResourcePackPath)

				if err != nil {
					log.Fatal(err)
				}

				url, err := fileupload.UploadToHost("https://0x0.st", file)

				if err != nil {
					log.Fatal(err)
				}

				fmt.Printf("%v\n", url)
			}

			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
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

func isDir(path string) (bool, error) {
	fileInfo, err := os.Stat(path)

	if err != nil {
		return false, err
	}

	return fileInfo.IsDir(), err
}

func minecraftExists(pathToMC string) (bool, error) {
	exists, err := dirExists(pathToMC)

	if err != nil {
		log.Fatal(err)
		return true, err
	}

	if !exists {
		fmt.Println("Minecraft directory not found, quitting")
		return false, err
	}

	return true, nil
}
