package main

import (
	"fmt"
// 	"net"
	"os"
	"strings"
	"crypto/tls"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Print("URL not provided\n")
		fmt.Printf("Usage: %v <url>\n", os.Args[0])
		os.Exit(0)
	}
	url := os.Args[1]

	protocol := "gemini://"	
	protocol_splited := strings.Split(url, protocol)
	
	// Get out the paths
	path_splited := strings.Split(protocol_splited[1], "/")

	// e.g. geminiprotocol.net:1965
	serverAddr := path_splited[0] + ":1965"
	//fmt.Println(serverAddr)
	
	/*
 	tcpAddr, err := net.ResolveTCPAddr("tcp", serverAddr)

	if err != nil {
		fmt.Println("Error acquiring TCPAddr: ", err.Error())
		os.Exit(1)
	}
	*/

	config := &tls.Config{
		// ServerName is required for certificate verification (e.g., "example.com")
		InsecureSkipVerify: true,
		ServerName: path_splited[0],
	}
	// fmt.Println(path_splited[0])

	conn, err := tls.Dial("tcp", serverAddr, config)
    if err != nil {
        println("Dial failed:", err.Error())
        os.Exit(1)
    }

    _, err = conn.Write([]byte(url+"\r\n"))
    if err != nil {
        println("Write to server failed:", err.Error())
        os.Exit(1)
    }

	output := ""
    reply := make([]byte, 1024)
	
	n, err := conn.Read(reply)

	for true {			
		if n != 0 {
			output += string(reply[1:n])
		} else {
			output += string(reply)
		}
		n, err = conn.Read(reply)

		if n == 0 {
			break;	
		}
		if err != nil {
        	println("Failed to read from server:", err.Error())
        	os.Exit(1)
	    }
	}

	fmt.Println(output)

    conn.Close()
}


