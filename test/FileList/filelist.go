package main

import (
	"os"
	"io/ioutil"
	"fmt"
	"regexp"
)

func GetFileList(dir string) []os.FileInfo{
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Get file list error: %v\n", err)
		return nil
	}
	return files
}

func GetApps(dir string)(apps []string){

	reg := regexp.MustCompile(`app[\d]`)
	files := GetFileList(dir)
	for _, v := range files{
		if reg.FindAllString(v.Name(),-1) != nil{
			apps = append(apps,v.Name())
		}

	}
	return apps
}


func main() {

	fmt.Println(GetFileList(os.Args[1]))
	files := GetFileList(os.Args[1])
	for _, v := range files{
		fmt.Println(v.Name())
	}
	fmt.Println(GetApps(os.Args[1]))
}
