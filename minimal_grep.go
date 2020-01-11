package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func getArgs() (*bool, *bool, *bool, *bool, *string, string, []string) {

	recursivePtr := flag.Bool("r", false, "recursive search, first filename/dirname will be taken as start-off point")
	perlSyntaxPtr := flag.Bool("P", false, "PATTERN in perl syntax")
	lineByLinePtr := flag.Bool("n", false, "search line by line")
	ignoreCasePtr := flag.Bool("i", false, "ignore case")
	dirToExcludePtr := flag.String("exclude-dir", "", "DIR to exclude")

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

	return recursivePtr, perlSyntaxPtr, lineByLinePtr, ignoreCasePtr, dirToExcludePtr, pattern, filenames
}

type search_fn func(string, string, bool) []string

func recursiveSearch(pattern string, filenames []string, dirToExcludePtr *string, ignoreCase bool, search search_fn) {
	dirname := filenames[0]

	err := filepath.Walk(dirname, func(path string, fileinfo os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		if fileinfo.IsDir() && fileinfo.Name() == *dirToExcludePtr {
			return filepath.SkipDir
		}

		if !fileinfo.IsDir() {
			matches := search(pattern, readContent(path), ignoreCase)
			printMatches(matches, fileinfo.Name())
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

	recursivePtr, perlSyntaxPtr, lineByLinePtr, ignoreCasePtr, dirToExcludePtr, pattern, filenames := getArgs()

  fmt.Printf("lineByLine: %v\n", *lineByLinePtr)

	var search search_fn
	if !*perlSyntaxPtr {
		search = searchString
	} else {
		search = searchRegex
	}

	if *recursivePtr == true {
		recursiveSearch(pattern, filenames, dirToExcludePtr, *ignoreCasePtr, search)
	} else {
		// filename globbing search
		for _, filename := range filenames {
			content := readContent(filename)
			matches := search(pattern, content, *ignoreCasePtr)
			printMatches(matches, filename)
		}
	}

}
