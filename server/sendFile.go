package server

import (
	"fmt"
	"net"
	"os"
	"bufio"
	"strconv"
)

var files []string

func takeFolder(folderPath string) error {
	if _, err := os.Stat(string(folderPath)); os.IsNotExist(err) {
		fmt.Println("Folder does not exist")
		return err
	}
	fileList(folderPath)
	return nil
}

func fileList(folderPath string) ([]string, error) {
	filesOs, err := os.ReadDir(folderPath)
	if err != nil {
		fmt.Println("Error reading directory")
		return nil, err
	}
	for _, fileDetails := range filesOs {
		//append only files not directories
		if fileDetails.IsDir() {
			continue
		}
		files = append(files, fileDetails.Name())
	}
	files = append(files, "quit")
	return files, nil
}

// start the server and listen for incoming connections.
// Once a connection is established, it will handle the request.
//
// folderPath: "C:/Users/username/Desktop", 
// ip: "localhost", 
// port: "20"
//
// returns an error if there is an error listening or accepting or sending file
func Start(folderPath string, ip string, port string) error {
	err := takeFolder(folderPath)
	if err != nil {
		fmt.Println("Error taking folder")
		return err
	}
	listener, err := net.Listen("tcp", ip + ":" + port)
	if err != nil {
		fmt.Println("Error listening")
		return err
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting")
			return err
		}
		go handleRequest(conn, folderPath)
	}
}

func handleRequest(conn net.Conn, folderPath string) error {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		input := scanner.Text()
		if(input == "LIST") {
			for _, file := range files {
				fmt.Fprintln(conn, file)
			}
		} else {
			fileStat, err := os.Stat(string(folderPath) + "/" + input)
			if err != nil {
				fmt.Println("Error reading file")
				return err
			}
			
			fmt.Fprint(conn, strconv.Itoa(int(fileStat.Size())))
			file, err := os.ReadFile(string(folderPath) + "/" + input)
			if err != nil {
				fmt.Println("Error reading file")
				return err
			}
			
			conn.Write(file)
		}
	}
	return nil
}
