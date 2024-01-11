package main

import (
	"bufio"
	"fmt"
	"log"
	"net"

	"github.com/CarlCao17/go-gears/pkg/bufferpool"
)

func main() {
	listener, err := net.Listen("tcp", ":8899")
	if err != nil {
		panic(fmt.Sprintf("ftp server listen socket tcp :8899 failed: %v", err))
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("client set connection to server failed: %v", err)
			continue
		}
		go ftp(conn)
	}
}

var (
	bp = bufferpool.NewBytesPool()
)

func ftp(conn net.Conn) {
	defer conn.Close()

	if buf == nil {
		panic("have no bytes buffer")
	}
	r := bufio.NewReader(conn)
	if !auth(conn) {
		log.Printf("ftp: failed to pass the auth, client_addr=%s\n", conn.RemoteAddr())
		return
	}

	for {
		cmd := parseCmd()
		switch cmd {
		case 'CD':

		}
	}
}



func parseCmd() (string {

}

func auth(conn net.Conn) bool {

}
