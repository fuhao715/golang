package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	fmt.Println("Welcome to the palyground .")
	fmt.Println("The time is ", time.Now())
	fmt.Println("And if you try to open a file :")
	file, err := os.Open("D:\\a.xml")
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 100000)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes: %q\n", count, data[:count])

	fmt.Println("Or access the network:")
	conn, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	// status, err := bufio.NewReader(conn).ReadString('\n')

}
