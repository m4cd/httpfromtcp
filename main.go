package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal(err)
	}
	
	channel := getLinesChannel(f)

	for line := range channel {
		fmt.Printf("read: %s\n", line)
	}
}

func getLinesChannel(f io.ReadCloser) <-chan string {
	channel := make(chan string)

	go func(){
		s := ""
		for {
			b := make([]byte, 8)
			_, err := f.Read(b)
			
			parts := strings.Split(string(b),"\n")
			
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