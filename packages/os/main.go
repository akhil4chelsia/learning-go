package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {

	//File exist
	check := isFileExist("sample.txt")
	fmt.Println("Check status: ", check)

	//File does not exist
	check = isFileExist("sample1.txt")
	fmt.Println("Check status: ", check)

	//Get file status
	stats := getFileStats("sample.txt")
	fmt.Println("Updated: ", stats.ModTime())
	fmt.Println("Mode: ", stats.Mode())
	fmt.Println("Size: ", stats.Size())
	fmt.Println("Directory?: ", stats.IsDir())
}

func getFileStats(path string) fs.FileInfo {
	stat, err := os.Stat(path)
	if err != nil {
		panic("Error while getting file status.")
	}
	return stat
}

//Check file exists or not
func isFileExist(path string) bool {
	//os.Stat() => to get file stats
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
