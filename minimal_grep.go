package main

import (
  "fmt"

  "flag"

   "os"
  //"path/filepath"

  // "regexp"
  // "io/ioutil"
)


func getArgs() (bool, bool, bool, bool, string, string) {


  recursive := *flag.Bool("r", false, "recursive search")
  perlSyntax := *flag.Bool("P", false, "PATTERN in perl syntax")
  lineByLine := *flag.Bool("n", false, "search line by line")
  ignoreCase := *flag.Bool("i", false, "ignore case")

  flag.Parse()


  if len(os.Args) == 1 {
    var usage = func() {
      fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
      flag.PrintDefaults()
      os.Exit(0)
    }

    usage()
  }


  fmt.Println(recursive)
  pattern:= ""
  filename := ""

  if len(flag.Args()) > 0 {
    pattern = flag.Args()[0]
  }

  if len(flag.Args()) > 1 {
   filename = flag.Args()[1]
  }

  return recursive, perlSyntax, lineByLine, ignoreCase, pattern, filename
}

func getFile(filename string) {

  //fmt.Printf("filename %s\n", filename)
}


func readAll() {

}


func readLineByLine() {

}


// just use FindAllString() after MustCompile; do linenumbers using readline...

func main() {

  recursive, perlSyntax, lineByLine, ignoreCase, pattern, filename := getArgs()

  fmt.Println(recursive)
  fmt.Println(perlSyntax)
  fmt.Println(lineByLine)
  fmt.Println(ignoreCase)
  fmt.Println(pattern)
  fmt.Println(filename)

  getFile(filename)

  // bytes, _ := ioutil.ReadFile("test_regex.go")


  // content := string(bytes)
  // re := regexp.MustCompile("[A-Za-z]+")

  // fmt.Println(os.Args)
  // //fmt.Println(content)
  // fmt.Println("")
  // //fmt.Println(re.FindAllString(content, -1))
}

