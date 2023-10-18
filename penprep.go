package main

import (
	"fmt"
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
	err := os.MkdirAll(d, filemode)
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
		"pentest/external/scans/nmap",
		"pentest/external/scans/other",
		"pentest/external/loot/screenshots",
		"pentest/external/exploits",
		"pentest/external/downloads",
		"pentest/external/tools",
		"pentest/external/tmp",
		"pentest/external/pub",
		"pentest/internal",
		"pentest/internal/scans/nmap",
		"pentest/internal/scans/other",
		"pentest/internal/loot",
		"pentest/internal/loot/hashdumps",
		"pentest/internal/loot/tcpdump",
		"pentest/internal/loot/bloodhound",
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

	fmt.Println("[+] done.")

}
