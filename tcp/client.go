package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)


func main(){
	conn, err := net.Dial("tcp", ":8085")
	if err != nil{
		panic(err)
	}
	defer conn.Close()
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text+"\n")

		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil{
			panic(err)
		}
		fmt.Printf("got message %s", message)
	}
}
