package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal(err)
	}
	for {
		b := make([]byte, 8)
		_, err = f.Read(b)
		if err == io.EOF {
			break
		}
		fmt.Printf("read: %s\n", b)
	}
}
