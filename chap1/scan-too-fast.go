package main

import (
	//Standard IO package
	"fmt"
	//Standard networking package
	"net"
)

func main() {
	for i := 1; i <= 1024; i++ {
		go func(j int) {
		address := fmt.Sprintf("scanme.nmap.org:%d", j)


		//Fetch from a website using .Dial
		//2 Parameters <Connection type: String, URL: string>
		conn, err := net.Dial("tcp", address)
		//Check if connection is made
		if err != nil {
			fmt.Println("Connection Failed")
			return
		}
		conn.Close()
		fmt.Printf("%d open\n", j)
	}(i)
	}
}