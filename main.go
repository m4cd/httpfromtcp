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
	s := ""
	for {
		b := make([]byte, 8)
		_, err = f.Read(b)
		
		parts := strings.Split(string(b),"\n")
		
		s = s + parts[0]
		
		if len(parts) == 2 {
			fmt.Printf("read: %s\n", s)
			s = parts[1]
		}

		if err == io.EOF {
			fmt.Printf("read: %s\n", s)
			break
		}
	}
}
