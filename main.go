package main

import (
	"github.com/jrjaro18/go-ftp/client"
	"github.com/jrjaro18/go-ftp/server"
	"fmt"
)

const (
	addr = "C:/Users/username"
	fileName = "doremifasolatido.mp3"
	network_addr = "localhost"
	port = "20"
)

// example of how to use the server and client
func main() {

	go func() {
		err := server.Start( addr + "/Desktop", network_addr, port )
		if(err != nil) {
			fmt.Println(err)
		}
	}()

	err := client.ReceiveFile( addr + "/Downloads", fileName, network_addr, port )
	if(err != nil) {
		fmt.Println(err)
	}

}
