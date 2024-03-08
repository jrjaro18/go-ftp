package client

import (
	"bufio"
	"fmt"
	"net"
)


// ReceiveList connects to the server and requests a list of files in the server's directory
//
// ip: "localhost", port: "20"
//
func ReceiveList(ip string, port string) ([]string, error) {
	conn, err := net.Dial("tcp", ip + ":" + port)
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		return nil, err
	}
	defer conn.Close()
	conn.Write([]byte("LIST\n"))
	list := make([]string, 0)
	scanner := bufio.NewScanner(conn)		
	for scanner.Scan() {
		if(scanner.Text() == "quit") {
			break
		}
		list = append(list, scanner.Text())
	}
	return list, nil
}
