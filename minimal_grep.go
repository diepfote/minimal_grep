package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func getArgs() (*bool, *bool, *bool, *bool, *string, string, []string) {

	recursivePtr := flag.Bool("r", false, "recursive search, first filename/dirname will be taken as start-off point")
	perlSyntaxPtr := flag.Bool("P", false, "PATTERN in perl syntax")
	lineByLinePtr := flag.Bool("n", false, "search line by line")
	ignoreCasePtr := flag.Bool("i", false, "ignore case")
	dirsToExclude := flag.String("exclude-dirs", "", "DIRs to exclude (separated by commas ',')")

	flag.Parse()

	if len(os.Args) == 1 {
		var usage = func() {
			fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
			flag.PrintDefaults()
			os.Exit(0)
		}

		usage()
	}

	pattern := ""

	if len(flag.Args()) > 0 {
		pattern = flag.Args()[0]
	}

	filenames := []string{""}
	if len(flag.Args()) > 1 {
		filenames = flag.Args()[1:len(flag.Args())]
	}

	return recursivePtr, perlSyntaxPtr, lineByLinePtr, ignoreCasePtr, dirsToExclude, pattern, filenames
}

type search_fn func(string, string, bool, bool) map[int]string

func recursiveSearch(pattern string, filenames []string, dirsToExclude *string, ignoreCase bool, lineByLine bool, perlSyntax bool, search search_fn) {
	dirname := filenames[0]

	if len(dirname) < 1 {
		dirname = "."
	}

	err := filepath.Walk(dirname, func(path string, fileinfo os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		for _, dirname := range strings.Split(*dirsToExclude, ",") {
			if fileinfo.IsDir() && fileinfo.Name() == dirname {
				return filepath.SkipDir
			}
		}

		if !fileinfo.IsDir() {
			matches := search(pattern, readContent(path), ignoreCase, lineByLine)
			printMatches(matches, pattern, fileinfo.Name(), lineByLine, perlSyntax, ignoreCase)
		}

		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", dirname, err)
		return
	}
}

// just use FindAllString() after MustCompile; do linenumbers using readline...

func main() {

	recursivePtr, perlSyntaxPtr, lineByLinePtr, ignoreCasePtr, dirsToExclude, pattern, filenames := getArgs()

	var search search_fn
	if !*perlSyntaxPtr {
		search = searchString
	} else {
		search = searchRegex
	}

	if *recursivePtr == true {
		recursiveSearch(pattern, filenames, dirsToExclude, *ignoreCasePtr, *lineByLinePtr, *perlSyntaxPtr, search)
	} else {
		// filename globbing search
		for _, filename := range filenames {
			content := readContent(filename)
			matches := search(pattern, content, *ignoreCasePtr, *lineByLinePtr)
			printMatches(matches, pattern, filename, *lineByLinePtr, *perlSyntaxPtr, *ignoreCasePtr)
		}
	}

}
