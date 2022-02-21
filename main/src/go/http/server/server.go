package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close() // no
	for {
		var buf [128]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			// panic(err)
			fmt.Println("read from conn failed err:", err)
			break
		}
		fmt.Println(n, buf[:n])
	}
}
