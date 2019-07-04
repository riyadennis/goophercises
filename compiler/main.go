package main

import (
	"fmt"
	"github.com/prometheus/common/log"
	scanner2 "go/scanner"
	"go/token"
	"io/ioutil"
)

func main(){
	scan("../quiz/main.go")
}

func scan(fileName string){
	var scanner scanner2.Scanner
	fc, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Errorf("unable tp open the file :: %v", err)
	}
	fileSet := token.NewFileSet()
	file := fileSet.AddFile("test.go", fileSet.Base(), len(fc))
	scanner.Init(file, fc, nil, 0)
	for {
		pos, tok, str := scanner.Scan()
		fmt.Printf("%v-%v-%v", pos, tok, str)
		if tok == token.EOF{
			break
		}
	}
}

