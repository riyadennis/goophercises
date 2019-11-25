package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	scanner2 "go/scanner"
	"go/token"
	"io/ioutil"

	"github.com/prometheus/common/log"
)

func main() {
	scan("../quiz/main.go")
	_, cont := fileSet("../quiz/main.go")
	fileSet := token.NewFileSet()
	af, err := parser.ParseFile(fileSet, "test.go", cont, 0)
	if err != nil {
		log.Errorf("unable to parse file :: %v", err)
	}
	ast.Print(fileSet, af)
}

func fileSet(fileName string) (*token.File, []byte) {
	fc, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Errorf("unable tp open the file :: %v", err)
	}
	fileSet := token.NewFileSet()
	return fileSet.AddFile("test.go", fileSet.Base(), len(fc)), fc
}

func scan(fileName string) {
	var scanner scanner2.Scanner
	file, fc := fileSet(fileName)
	scanner.Init(file, fc, nil, 0)
	for {
		pos, tok, str := scanner.Scan()
		fmt.Printf("%v-%v-%v", pos, tok, str)
		if tok == token.EOF {
			break
		}
	}
}
