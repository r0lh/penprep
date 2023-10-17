package main

import (
	"io/fs"
	"log"
	"os"
)

// set filemask
var filemode fs.FileMode = 0750

func checkError(err error) {
	if err != nil {
		log.Println(err)
	}
}

func createDirectory(d string) {
	err := os.Mkdir(d, filemode)
	checkError(err)
}

func createFile(f string) {
	file, err := os.Create(f)
	checkError(err)
	defer file.Close()
}

func main() {
	directories := []string{
		"pentest",
		"pentest/external",
		"pentest/external/scans",
		"pentest/external/loot",
		"pentest/external/exploits",
		"pentest/external/downloads",
		"pentest/external/tools",
		"pentest/external/tmp",
		"pentest/external/pub",
		"pentest/internal",
		"pentest/internal/scans",
		"pentest/internal/loot",
		"pentest/internal/loot/hashdumps",
		"pentest/internal/loot/screenshots/",
		"pentest/internal/exploits",
		"pentest/internal/downloads",
		"pentest/internal/tools",
		"pentest/internal/tmp",
		"pentest/internal/pub",
	}

	files := []string{
		"pentest/external/notes.txt",
		"pentest/external/loot/credentials.txt",
		"pentest/internal/notes.txt",
		"pentest/internal/loot/credentials.txt",
	}

	for _, d := range directories {
		createDirectory(d)
	}

	for _, f := range files {
		createFile(f)
	}

}
