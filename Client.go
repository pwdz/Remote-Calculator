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
		//reader := bufio.NewReader(os.Stdin)
		var input string
		fmt.Println("Enter input")
		_, _ = fmt.Scanln(&input)
		//text, _ := reader.ReadString('\n')
		// send to socket
		//fmt.Print(conn, text + "\n")
		// listen for reply
		//message, _ := bufio.NewReader(conn).ReadString('\n')

		//fmt.Print("Message from server: " + message)
		fmt.Println(input)
		_, _ = conn.Write([]byte(string(input+"\n")))
		res,_ :=bufio.NewReader(conn).ReadString('\n')
		fmt.Println("result is: "+res)
		if strings.Compare(input, "finish") == 0 {
			break
		}
	}
}
