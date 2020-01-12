package main

import (
	"fmt"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

func colorTerm(line string, term string, perlSyntax bool, ignoreCase bool) string {
	if runtime.GOOS != "windows" {

		if ignoreCase {
			// TODO this is half-assed

			term = strings.ToLower(term)
			line = strings.ToLower(line)
		}

		if perlSyntax {
			// TODO fix match replaced with regex

			re := regexp.MustCompilePOSIX(term)
			return re.ReplaceAllString(line, "\033[1;31m"+term+"\033[0m")
		} else {
			return strings.ReplaceAll(line, term, "\033[1;31m"+term+"\033[0m")
		}
	}

	return line
}

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

func printMatches(matches map[int]string, pattern string, filename string, lineByLine bool, perlSyntax bool, ignoreCase bool) {
	blue_colon := colorBlue(":")

	for linenumber, match := range matches {

		match = colorTerm(match, pattern, perlSyntax, ignoreCase)

		if len(filename) > 0 {
			filename = colorPurple(filename)

			if lineByLine {
				fmt.Printf("%s%s%s%s%s", filename, blue_colon, colorGreen(strconv.Itoa(linenumber+1)),
					blue_colon, match)
			} else {
				fmt.Printf("%s%s%s", filename, blue_colon, match)
			}
		} else {
			if lineByLine {
				fmt.Printf("%s%s%s", colorGreen(strconv.Itoa(linenumber+1)), blue_colon, match)
			} else {
				fmt.Printf("%s", match)
			}
		}
	}
}
