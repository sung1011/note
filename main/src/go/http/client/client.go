package main

import (
	"bufio"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", ":9090")
	defer conn.Close()
	if err != nil {
		panic(err)
	}

	r := bufio.NewReader(os.Stdin)
	for {
		input, err := r.ReadString('\n')
		if err != nil {
			panic(err)
		}
		trimmedInput := strings.TrimSpace(input)
		if trimmedInput == "q" || trimmedInput == "Q" {
			break
		}
		_, err = conn.Write([]byte(trimmedInput))
		if err != nil {
			panic(err)
		}
	}
}
