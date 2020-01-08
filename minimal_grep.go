package main

import (
	"fmt"

	"flag"
	"os"

	"bufio"
	"io/ioutil"

	"path/filepath"

	"regexp"
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

  fmt.Printf("all remainig Args: %t\n", flag.Args())
	if len(flag.Args()) > 0 {
		pattern = flag.Args()[0]
	}

  filenames := []string{""}
	if len(flag.Args()) > 1 {
    filenames = flag.Args()[1:len(flag.Args())]
	}

	return recursivePtr, perlSyntaxPtr, lineByLinePtr, ignoreCasePtr, dirToExcludePtr, pattern, filenames
}

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

func readAll() {

}

func readLineByLine() {

}

func printMatches(matches []string) {
	for _, match := range matches {
		fmt.Printf("%s", match)
	}
}

func search(re *regexp.Regexp, content string)  ([]string) {
  return re.FindAllString(content, -1)
}

func readContent(filename string) string {
	reader, file := getReader(filename)
	defer file.Close()

	bytes, _ := ioutil.ReadAll(reader)
	return string(bytes)
}

func recursiveSearch(re *regexp.Regexp, filenames []string, dirToExcludePtr *string) {
  fmt.Printf("filenames: %v\n", filenames)
  dirname := filenames[0]
  err := filepath.Walk(dirname, func(path string, info os.FileInfo, err error) error {
    if err != nil {
      fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
      return err
    }
    if info.IsDir() && info.Name() == *dirToExcludePtr {
      //fmt.Printf("skipping a dir without errors: %+v \n", info.Name())
      return filepath.SkipDir
    }


    if !info.IsDir() {
      //fmt.Printf("file: %q\n", path)
      matches := search(re, readContent(path))
      printMatches(matches)
    }


    //fmt.Printf("visited file or dir: %q\n", path)
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

	fmt.Printf("recursive: %t\n", *recursivePtr)
	fmt.Printf("perlSyntax: %t\n", *perlSyntaxPtr)
	fmt.Printf("lineByLine: %t\n", *lineByLinePtr)
	fmt.Printf("ignoreCase: %t\n", *ignoreCasePtr)
	fmt.Printf("dirToExclude: %s\n", *dirToExcludePtr)
	fmt.Printf("pattern: %s\n", pattern)
	fmt.Printf("filename/dirname: %v\n", filenames)
	fmt.Println("")
	fmt.Println("printMatches:")


	for _, char := range pattern {
		fmt.Println(char)
	}
	//re := regexp.MustCompile("[\n]{0,1}.*" + pattern + ".*[\n]{0,1}")
  re := regexp.MustCompilePOSIX("[\n]{0,1}.*" + pattern + ".*[\n]{0,1}")

	// fmt.Println(os.Args)
	//fmt.Println(content)



	if *recursivePtr == true {
    recursiveSearch(re, filenames, dirToExcludePtr)
	} else {
    // accomodate for filename globbing
    for _, filename := range filenames {
      content := readContent(filename)
      matches := search(re, content)
      printMatches(matches)
    }
  }

	//fmt.Println()

}

