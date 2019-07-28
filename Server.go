package main

import (
	"fmt"
	"net"
	"os"
)
const(
	ConnectionPort = "1540"
	ConnectionHost = "localhost"
)
func main()  {
	socket,err := net.Listen("tcp",ConnectionHost+":"+ConnectionPort)
	if err!=nil{
		fmt.Println("error listening")
		os.Exit(-1)
	}
	defer socket.Close()
	fmt.Println("Listening on " + ConnectionHost + ":" + ConnectionPort)
	fmt.Println(socket)
	for{
		conn,err:=socket.Accept()
		if err!=nil{
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		go handleRequest(conn)
	}
}
func handleRequest(conn net.Conn)  {
	
}