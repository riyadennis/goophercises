package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	l, err := net.Listen("tcp", ":8085")
	if err != nil {
		fmt.Println("ERROR")
		os.Exit(1)
	}
	defer l.Close()
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("ERROR")
		os.Exit(1)
	}
	for {
		str, err := bufio.NewReader(conn).ReadString('\n')
		switch err {
		case nil:
			fmt.Printf("from client >> %s", str)
			fmt.Fprint(conn, "SUCCESS\n")
			break
		case io.EOF:
			fmt.Fprint(conn, "SUCCESS\n")
			os.Exit(0)
		default:
			fmt.Fprint(conn, "ERROR\n")
			os.Exit(1)
		}
	}
}
