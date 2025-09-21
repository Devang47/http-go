package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
)

const PORT = ":42069"

func main() {
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("error listening for TCP traffic: %s\n", err.Error())
	}
	defer listener.Close()

	fmt.Println("Listening for TCP traffic on", PORT)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("error: %s\n", err)
		}
		fmt.Println("Accepted connection from", conn.RemoteAddr())

		for line := range getLinesChannel(conn) {
			fmt.Println(line)
		}
		fmt.Println("Connection to ", conn.RemoteAddr(), "closed")
	}
}

func getLinesChannel(f io.ReadCloser) <-chan string {
	out := make(chan string, 1)

	go func() {
		defer f.Close()
		defer close(out)
		curLine := ""
		for {
			buffer := make([]byte, 8)
			n, err := f.Read(buffer)
			if err != nil {
				break
			}

			if i := bytes.IndexByte(buffer, '\n'); i != -1 {
				curLine += string(buffer[:i])
				out <- curLine
				curLine = string(buffer[i+1:])
			} else {
				curLine += string(buffer[:n])
			}
		}

		if curLine != "" {
			out <- curLine
		}
	}()

	return out
}
