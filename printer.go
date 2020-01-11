package main

import (
	"fmt"
	"runtime"
)

func colorPurple(term string) string {
	if runtime.GOOS != "windows" {
		return "\033[0;35m" + term + "\033[0m"
	}

	return term
}

func colorGreen(term string) string {
	if runtime.GOOS != "windows" {
		return "\033[0;32m" + term + "\033[0m"
	}

	return term
}

func colorBlue(term string) string {
	if runtime.GOOS != "windows" {
		return "\033[0;34m" + term + "\033[0m"
	}

	return term
}

func printMatches(matches []string, filename string) {
	for _, match := range matches {

		if len(filename) > 0 {
			filename = colorPurple(filename)
			fmt.Printf("%s:%s", filename, match)
		} else {
			fmt.Printf("%s", match)
		}
	}
}
