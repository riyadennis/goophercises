package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

const Port = ":8085"

func main() {
	conn, err := net.Dial("tcp", Port)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	for {
		userInput(conn)
		checkResponse(conn)
	}
}

func userInput(conn net.Conn) {
	fmt.Print(">>")
	userInput, err := bufio.NewReader(os.Stdin).ReadString('\n')
	switch err {
	case nil:
		fmt.Fprintf(conn, "%s\n", userInput)
	case io.EOF:
		os.Exit(0)
	default:
		fmt.Fprintf(conn, "ERROR")
		os.Exit(1)
	}

}

func checkResponse(conn net.Conn) {
	response, err := bufio.NewReader(conn).ReadString('\n')
	switch err {
	case nil:
		fmt.Println(response)
		switch response {
		case "SUCCESS\n":
			fmt.Println("communication successful")
		default:
			fmt.Println("communication not successful")
		}
	case io.EOF:
		os.Exit(0)
	default:
		fmt.Fprintf(conn, "ERROR")
		os.Exit(1)
	}

}
