package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/smallfish/ftp"
)

func main() {
	ftp := new(ftp.FTP)
	// debug default false
	ftp.Debug = true
	ftp.Connect("localhost", 21)

	// login
	ftp.Login("bob", "bob")
	if ftp.Code == 530 {
		fmt.Println("error: login failure")
		os.Exit(-1)
	}

	// pwd
	ftp.Pwd()
	fmt.Println("code:", ftp.Code, ", message:", ftp.Message)

	// make dir
	ftp.Mkd("teste")

	ftp.Request("TYPE I")

	// stor file
	b, _ := ioutil.ReadFile("/tmp/a.txt")
	ftp.Stor("a.txt", b)

	ftp.Quit()
}
