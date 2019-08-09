package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)
const(
	PORT = "1540"
	HOST = "localhost"
)
func main()  {
	socket,err := net.Listen("tcp",HOST+":"+PORT)
	if err!=nil{
		fmt.Println("error listening")
		os.Exit(-1)
	}
	defer socket.Close()
	fmt.Println("Listening on " + HOST+ ":" + PORT)
	fmt.Println(socket)
	for{
		conn,err:=socket.Accept()
		if err!=nil{
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		fmt.Println("Connected!")
		go handleRequest(conn)
	}
}
func handleRequest(conn net.Conn)  {
	// Read the incoming connection into the buffer.
	for {
		netData, err := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("you sent this:"+netData)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
		}
		if strings.Compare(string(netData), "finish\n") == 0 {
			fmt.Println("byeeeeeeeeeeee")
			break
		}
		_, _ = conn.Write([]byte(strconv.Itoa(Calculate1(string(netData)))+"\n"))
	}
	_ = conn.Close()
}
func splitter1(input string) []string {
	return strings.FieldsFunc(input, Split1)
}
func Split1(r rune) bool {
	return r == '+' || r == '-'
}
func Calculate1(input string) int {
	elements := splitter1(input)
	var index int8
	var result int
	result = 0
	for _, element := range elements {
		element = strings.TrimSpace(element)
		num,_ := strconv.ParseInt(element,10,64)
		index = int8(strings.Index(input, element))
		if index > 0 {
			switch input[index-1] {
			case '-':
				num *=-1
			case '+':
				//nothing :)
			}
		}
		result+=int(num)
	}
	return result
}