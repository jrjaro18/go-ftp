# go-ftp
This repository contains a simple implementation of an FTP server and client in Go. The server allows clients to connect, request files, and retrieve a list of files from the server's directory. The client can connect to the server, download files, and retrieve a list of files from the server.

<img src="https://www.lattaharris.com/wp-content/uploads/2019/12/file-transfer.jpg" height="350px" width="900px">

## Features

- **Server**: 
  - Listens for incoming connections from clients.
  - Supports file download requests.
  - Provides a list of files in the server's directory.

- **Client**:
  - Connects to the server using TCP.
  - Requests file downloads from the server.
  - Retrieves a list of files from the server's directory.

## Usage

## To install the package, run the following command
```shell
go get github.com/jrjaro18/go-ftp
```

## Demo

#### Import the following package
```go
import "github.com/jrjaro18/go-ftp"
```

#### Sample Input
```go
const (
	addr = "path to your folder"
        //NOTE: without / in the end
	networkAddr = "localhost"
	port = "20"
)
```
## SERVER GUIDE
#### To start the FTP server, use the following code snippet:
```go
err := server.Start(addr, networkAddr, port)
if err != nil {
  fmt.Println(err)
}
```
## CLIENT GUIDE
#### To start the FTP client and receive the file, use the following code snippet:
```go
err := client.ReceiveFile(addr, fileName, networkAddr, port)
if err != nil {
  fmt.Println(err)
}
```
#### To start the FTP client and receive the list of files available in server, use the following code snippet:
```go
arr, err := client.ReceiveList(networkAddr, port)
if err != nil {
  fmt.Println(err)
}
```
---
###### This project is licensed under the MIT License - see the [License](LICENSE) file for details.

###### This README provides clear instructions on how to use both the server and client components of your Go FTP package. Adjust the paths and configuration parameters as needed for your specific setup.
