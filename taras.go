package main

import (
	"fmt"
// 	"net"
	"os"
	"strings"
	"crypto/tls"
)

const (
	text = iota
	link = iota
	heading = iota
	sheading = iota
	ssheading = iota
	list = iota
	preformat = iota

	whitespace = iota
	newline = iota
	
	// implement if really needed, aka. > quote marks
	quote = iota

)

type Token struct {
	tokentype int
	value string
}


func parse(raw []rune) []Token {
	tokens := []Token{}

	for i := 0; i < len(raw); i++ {
		switch raw[i] {
			case '=':
				if raw[i+1] == '>' {
					i += 1	
					tokens = append(tokens, Token{link, "=>"})
				} else {
					tokens = append(tokens, Token{text, string(raw[i])})
				}
		
			case '#':
				if raw[i+1] == '#' {
					if raw[i+1] == '#' {
						tokens = append(tokens, Token{ssheading, "###"})
						i += 2
					} else {
						tokens = append(tokens, Token{sheading, "##"})
						i += 1
					}
				} else {
					tokens = append(tokens, Token{heading, "#"})
				}
			
			case '*':
				tokens = append(tokens, Token{list , "*"})

			case '\n':
				tokens = append(tokens, Token{newline , "\n"})

			case ' ':
			case '\t':
				tokens = append(tokens, Token{whitespace, string(raw[i])})
			
			case '`':
				if raw[i+1] == '`' {
					if raw[i+1] == '`' {
						tokens = append(tokens, Token{preformat, "```"})			
					} else {
						tokens = append(tokens, Token{text, string(raw[i])})
					}
				} else {
					tokens = append(tokens, Token{text, string(raw[i])})

				}

			default:
				tokens = append(tokens, Token{text, string(raw[i])})
			
		}	
	} 

	return tokens
}

// purely for debugging purposes
func tokentype(num int) string {
	return_v := ""
	switch num {
		case text:
			return_v = "Text"
		case link:
			return_v = "Link"
		case heading:
			return_v = "Heading"
		case sheading:
			return_v = "Sub-heading"
		case ssheading:
			return_v = "Sub-Sub-Heading"
		case list:
			return_v = "List"
		case preformat:
			return_v = "Preformated"
		case whitespace:
			return_v = "Whitespace"
		case newline:
			return_v = "Newline"	
	}

	return return_v
}

func print_tokens(tokens []Token) {
	for i := 0; i < len(tokens); i++ {
		fmt.Printf("{ TokenType: %v, Value: %v }, ", tokentype(tokens[i].tokentype), tokens[i].value)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Print("URL not provided\n")
		fmt.Printf("Usage: %v <url>\n", os.Args[0])
		os.Exit(0)
	}
	url := os.Args[1]
	
	if !strings.HasSuffix(url, "/") {
		url += "/"
	}

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
			output += string(reply[0:n])
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

    conn.Close()

//	tokens := parse([]rune(output))
	
	fmt.Println(output)
}


