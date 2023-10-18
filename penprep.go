package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
)

const VERSION = "v1.0.4"

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

func createExternalList() (directories, files []string) {
	directories = []string{
		"pentest/external",
		"pentest/external/scans/nmap",
		"pentest/external/scans/other",
		"pentest/external/loot/screenshots",
		"pentest/external/exploits",
		"pentest/external/downloads",
		"pentest/external/tools",
		"pentest/external/tmp",
		"pentest/external/pub",
	}

	files = []string{
		"pentest/external/notes.txt",
		"pentest/external/loot/credentials.txt",
	}

	return directories, files
}

func createInternalList() (directories, files []string) {
	directories = []string{
		"pentest/internal",
		"pentest/internal/scans/nmap",
		"pentest/internal/scans/other",
		"pentest/internal/loot/hashdumps",
		"pentest/internal/loot/tcpdump",
		"pentest/internal/loot/bloodhound",
		"pentest/internal/loot/screenshots",
		"pentest/internal/exploits",
		"pentest/internal/downloads",
		"pentest/internal/tools",
		"pentest/internal/tmp",
		"pentest/internal/pub",
	}

	files = []string{
		"pentest/internal/notes.txt",
		"pentest/internal/loot/credentials.txt",
	}

	return directories, files
}

func main() {
	var extFlag = flag.Bool("e", false, "create environ for external audit only")
	var intFlag = flag.Bool("i", false, "create environ for internal audit only")
	flag.Parse()

	directories := []string{}
	files := []string{}

	switch {
	case *extFlag:
		directories, files = createExternalList()
	case *intFlag:
		directories, files = createInternalList()
	default:
		d_ext, f_ext := createExternalList()
		d_int, f_int := createInternalList()
		for _, v := range d_ext {
			directories = append(directories, v)
		}

		for _, v := range d_int {
			directories = append(directories, v)
		}

		for _, v := range f_ext {
			files = append(files, v)
		}

		for _, v := range f_int {
			files = append(files, v)
		}
	}

	for _, d := range directories {
		createDirectory(d)
	}

	for _, f := range files {
		createFile(f)
	}

	fmt.Println("[+] done.")

}
