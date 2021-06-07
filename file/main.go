package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "This is the content for the file"
	writeTofile("./test.txt", content)
	defer readFile("./test.txt")
}

func writeTofile(fileName, data string) {
	file, err := os.Create(fileName)
	checkError(err)
	length, err := io.WriteString(file, data)
	checkError(err)
	fmt.Printf("Wrote a file with %v characters.\n", length)
	//close file. it is very important
	defer file.Close()
}

func readFile(filename string) {
	data, err := ioutil.ReadFile(filename)
	checkError(err)
	fmt.Println(string(data))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
