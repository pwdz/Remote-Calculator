package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	conn, _ := net.Dial("tcp", "localhost"+":"+"1540")
	for {
		var input string
		fmt.Println("Enter input")
		_, _ = fmt.Scanln(&input)
		fmt.Println(input)
		_, _ = conn.Write([]byte(string(input+"\n")))
		res,_ :=bufio.NewReader(conn).ReadString('\n')
		fmt.Println("result is: "+res)
		if strings.Compare(input, "finish") == 0 {
			break
		}
	}
}
