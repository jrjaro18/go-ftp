package main

import (
	"github.com/jrjaro18/go-ftp/client"
	"github.com/jrjaro18/go-ftp/server"
)

func main() {
	go server.Start("C:/Users/", "localhost", "20")
	client.ReceiveFile("C:/Users/Test Folder", "send.xlsx", "localhost", "20")
	// client.ReceiveList("localhost", "20")
}
