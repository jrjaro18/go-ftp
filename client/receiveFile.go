package client

import (
	"fmt"
	"net"
	"os"
	"strconv"
)


func takeFolder(folderPath string) (string, error) {
	if _, err := os.Stat(string(folderPath)); os.IsNotExist(err) {
		fmt.Println("Folder does not exist")
		return "", err
	}
	return folderPath, nil
}

// ReceiveFile connects to the server and requests a file from the server's directory
//
// path: "C:/Users/username/Desktop/Test Folder",
// fileName: "TT.png",
// ip: "localhost",
// port: "20"
//
// returns an error if there is an error connecting, reading, or creating the file
//
func ReceiveFile(path string, fileName string, ip string, port string) error {
	folderPath, err := takeFolder(path)
	if err != nil {
		return err
	}
	_, err = os.Stat(path); os.IsNotExist(err)
	if(err != nil) {
		fmt.Println("Folder does not exist")
		return err
	}
	conn, err := net.Dial("tcp", ip + ":" + port)
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		return err
	}
	defer conn.Close()
	conn.Write([]byte(fileName + "\n"))
	var buf [128]byte
	n, err := conn.Read(buf[:])
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return err
	}
	fileSize, err := strconv.Atoi(string(buf[:n]))
	if err != nil {
		fmt.Println("Error converting to int:", err.Error())
		return err
	}
	file, err := os.Create(folderPath + "/" + fileName)
	if err != nil {
		fmt.Println("Error creating file:", err.Error())
		return err
	}
	defer file.Close()
	buffer := make([]byte, fileSize)
	n, err = conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return err
	}
	file.Write(buffer[:n])
	return nil
}
