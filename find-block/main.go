package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var devices map[string]string

func main() {
	sysBlockDir := "/sys/block/"
	entries, err := ioutil.ReadDir(sysBlockDir)
	if err != nil {
		log.Println(err)
	}

	devices = make(map[string]string)

	for _, entry := range entries {
		if entry.Mode()&os.ModeSymlink == 0 {
			continue
		}

		file, err := ioutil.ReadFile(filepath.Join(sysBlockDir, entry.Name(), "dev"))
		if err != nil {
			log.Println(err)
			continue
		}

		devices[strings.TrimSpace(string(file))] = entry.Name()
	}

	fmt.Println(devices)
}
