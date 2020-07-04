package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:6942")
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	fmt.Fprintf(conn, "[0, 1, \"\", \"\"]")
	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	fmt.Printf("status: %s\n", string(status))
}
