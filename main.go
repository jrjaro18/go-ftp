package main

import (
	"github.com/jrjaro18/go-ftp/client"
	"github.com/jrjaro18/go-ftp/server"
	"fmt"
)

func main() {
	go server.Start("C:/Users/", "localhost", "20")
	err:= client.ReceiveFile("C:/Users/rohan/Desktop/Test Folder", "send.xlsx", "localhost", "20")
	if(err != nil) {
		fmt.Println(err)
	}
	// client.ReceiveList("localhost", "20")
}
