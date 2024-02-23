package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("working with reading and writing in files")
	content := "Golang is a procedural and statically typed programming language having the syntax similar to C programming language. Sometimes it is termed as Go Programming Language. It provides a rich standard library, garbage collection, and dynamic-typing capability. It was developed in 2007 by Robert Griesemer, Rob Pike, and Ken Thompson at Google but launched in 2009 as an open-source programming language and mainly used in Google production systems. Golang is one of the most trending programming languages among developers."
	file, err := os.Create("gofile.txt")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	fmt.Println("length is", length)
	defer file.Close()
	readfile, err := ioutil.ReadFile("gofile.txt")
	fmt.Println(string(readfile))
}
