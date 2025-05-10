package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {

	listener, err := net.Listen("tcp4", "127.0.0.1:42069")
	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Connection accepted...")
		
		channel := getLinesChannel(conn)
		
		for line := range channel {
			fmt.Printf("%s\n", line)
		}
		fmt.Println("Connection closed...")
	}

}

func getLinesChannel(f io.ReadCloser) <-chan string {
	channel := make(chan string)

	go func() {
		s := ""
		for {
			b := make([]byte, 8)
			_, err := f.Read(b)

			parts := strings.Split(string(b), "\n")

			s = s + parts[0]

			if len(parts) == 2 {
				channel <- s
				s = parts[1]
			}

			if err == io.EOF {
				channel <- s
				break
			}
		}
		defer close(channel)
	}()
	return channel
}
