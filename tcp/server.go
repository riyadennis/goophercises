package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func main(){
	l, err := net.Listen("tcp", ":8085")
	if err != nil{
		panic(err)
	}
	defer l.Close()

	for{
		conn, err := l.Accept()
		if err != nil{
			panic(err)
		}
		go func(){
			str, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil{
				if err == io.EOF{
					fmt.Println("client sent exit message")
					return
				}
				panic(err)
				return
			}
			fmt.Printf("got -> %s", str)
		}()
	}
}
