package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	""
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
	//buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	//req, err := conn.Read(buf)
	for {
		netData, err := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("you sent this:"+netData)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
		}
		if strings.Compare(string(netData), "finish") == 0 {
			break
		}
		fmt.Println("calculating...")
		_, _ = conn.Write([]byte(string(main.calculate(netData))))
	}
	_ = conn.Close()
}