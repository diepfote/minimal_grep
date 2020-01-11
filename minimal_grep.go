package main

import (
	"fmt"
	"flag"
	"os"
	"bufio"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"
	"runtime"
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
        fmt.Printf("matches: %t\n", matches)
	for _, match := range matches {

    if len(filename) > 0 {
      filename = colorPurple(filename)
      fmt.Printf("%s:%s", filename, match)
    } else {
      fmt.Printf("%s", match)
    }
	}
}


func readContent(filename string) string {
	reader, file := getReader(filename)
	defer file.Close()

	bytes, _ := ioutil.ReadAll(reader)
	return string(bytes)
}

type search_fn func(string, string, bool) []string

func recursiveSearch(pattern string, filenames []string, dirToExcludePtr *string, ignoreCase bool, search search_fn) {
  fmt.Printf("filenames: %v\n", filenames)
  dirname := filenames[0]

err := filepath.Walk(dirname, func(path string, fileinfo os.FileInfo, err error) error {
  if err != nil {
      fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
      return err
    }

    if fileinfo.IsDir() && fileinfo.Name() == *dirToExcludePtr {
      //fmt.Printf("skipping a dir without errors: %+v \n", fileinfo.Name())
      return filepath.SkipDir
    }


    if !fileinfo.IsDir() {
      //fmt.Printf("file: %q\n", path)
      matches := search(pattern, readContent(path), ignoreCase)
      printMatches(matches, fileinfo.Name())
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

	// fmt.Println(os.Args)
	//fmt.Println(content)
    var search search_fn
    if !*perlSyntaxPtr {
      search = func(pattern string, content string, ignoreCase bool)  ([]string) {
      if ignoreCase {
        pattern = strings.ToLower(pattern)
        content = strings.ToLower(content)
      }

        lines := strings.Split(content, "\n")
        var matches []string

        linenumber := 1
        for _, line := range lines {
          if strings.Contains(line, pattern) {
            matches = append(matches, line + "\n")
          }
          linenumber++
        }

        return matches
    }} else {
      search = func(pattern string, content string, ignoreCase bool) ([]string) {
        if ignoreCase {
          //panic("Not implemented")
          //pattern = "(?i)" + pattern
          pattern = strings.ToLower(pattern)
          content = strings.ToLower(content)
        }

        re := regexp.MustCompilePOSIX(pattern + ".*[\n]")
        return re.FindAllString(content, -1)
    }}



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

	//fmt.Println()
}

