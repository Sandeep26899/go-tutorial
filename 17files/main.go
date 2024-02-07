package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Files in golang")
	content := "This needs to go in a file - LearnCodeOnline.in"

	file, err := os.Create("./myLcoGoFile.txt")

	checkNilErr(err)

	length, err := io.WriteString(file, content)

	checkNilErr(err)

	fmt.Println("length is ", length)
	defer file.Close()
	readFile("./myLcoGoFile.txt")
}

func readFile(filename string) {
	dataBytes, err := os.ReadFile(filename)

	checkNilErr(err)

	fmt.Println("Text data inside the file is \n", string(dataBytes))
}

func checkNilErr(err error) {
	if err != nil {
		panic((err))
	}
}
