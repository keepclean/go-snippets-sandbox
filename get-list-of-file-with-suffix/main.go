package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

const cgDir string = "/sys/fs/cgroup/system.slice/"

func main() {
	entries, err := ioutil.ReadDir(cgDir)
	if err != nil {
		log.Fatal(err)
	}

	var items []string
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		if !strings.HasSuffix(entry.Name(), ".service") {
			continue
		}
		items = append(items, entry.Name())
	}

	for _, item := range items {
		itemEntries, err := ioutil.ReadDir(filepath.Join(cgDir, item))
		if err != nil {
			log.Fatal(err)
		}

		var memoryFiles []string
		for _, itemEntry := range itemEntries {
			if itemEntry.IsDir() {
				continue
			}
			if !strings.HasPrefix(itemEntry.Name(), "memory.") {
				continue
			}
			memoryFiles = append(memoryFiles, itemEntry.Name())
		}

		fmt.Println(memoryFiles)
	}
}
