package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	cf := ConfigFile{FileAsset{Path: "xxxx"}}
	if !cf.Exist() {
		if err := cf.Create(); err != nil {
			log.Fatalln(err)
		}
	}

	if err := cf.Read(); err != nil {
		log.Println(err)
	}
}

type FileAsset struct {
	Path string
}

func (f FileAsset) String() string {
	return f.Path
}

func (f FileAsset) Exist() bool {
	_, err := os.Stat(f.ExpandPath())
	if err != nil && os.IsNotExist(err) {
		fmt.Println(f)
		return false
	} else if err != nil {
		log.Println("non-IsNotExist error upon calling os.Stat:", err)
		return false
	}

	return true
}

func (f FileAsset) ExpandPath() string {
	return os.ExpandEnv(f.Path)
}

func (f FileAsset) DirName() string {
	return filepath.Dir(f.ExpandPath())
}

func (f FileAsset) Read() error {
	data, err := ioutil.ReadFile(f.ExpandPath())
	if err != nil {
		return err
	}

	fmt.Print(string(data))
	return nil
}

type ConfigFile struct {
	FileAsset
}

func (cf ConfigFile) Create() error {
	fd, err := os.Create(cf.Path)
	if err != nil {
		return err
	}
	defer fd.Close()

	log.Printf("file %s has been created\n", cf.Path)
	return nil
}
