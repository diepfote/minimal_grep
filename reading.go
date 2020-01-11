package main

import (
	"bufio"
	"io/ioutil"
	"os"
)

func getReader(filename string) (*bufio.Reader, *os.File) {

	if filename == "-" || filename == "" {
		reader := bufio.NewReader(os.Stdin)
		var file *os.File

		return reader, file
	} else {
		file, _ := os.Open(filename)
		reader := bufio.NewReader(file)

		return reader, file
	}
}

func readContent(filename string) string {
	reader, file := getReader(filename)
	defer file.Close()

	bytes, _ := ioutil.ReadAll(reader)
	return string(bytes)
}
