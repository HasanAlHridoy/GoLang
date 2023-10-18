package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	var content string = "This is my new file"
	file, err := os.Create("./newFile.txt")
	checkNilErr(err)
	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println(length)
	defer file.Close()
	readFile("./newFile.txt")

}
func readFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	checkNilErr(err)
	fmt.Println(string(data))
}
func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
