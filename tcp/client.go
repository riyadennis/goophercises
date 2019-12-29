package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":8085")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	userInput(conn)
	checkResponse(conn)
}

func userInput(conn net.Conn) {
	for {
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
}

func checkResponse(conn net.Conn) {
	for {
		response, err := bufio.NewReader(conn).ReadString('\n')
		switch err {
		case nil:
			switch response {
			case "SUCCESS\n":
				fmt.Println("communication successful")
				break
			default:
				fmt.Println("communication not successful")
				break
			}
		case io.EOF:
			os.Exit(0)
		default:
			fmt.Fprintf(conn, "ERROR")
			os.Exit(1)
		}
	}
}
