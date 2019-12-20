package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func ConnectToPinit() net.Conn {
	conn, err := net.Dial("tcp", "127.0.0.1:49001")
	if err != nil {
		fmt.Println(os.Args[0] + ": error: failed to connect to pinit")
		return nil
	}

	response, err := bufio.NewReader(conn).ReadString('\n')
	if strings.Split(response, ": ")[0] == "INFO" {
		return conn
	} else {
		return nil
	}
}

func StartService(conn net.Conn, service string) bool {
	conn.Write([]byte("START " + service + "\n"))
	response, err := bufio.NewReader(conn).ReadString('\n')
	response = strings.Trim(response, "\n")
	if err != nil {
		fmt.Println(os.Args[0] + ": error: " + err.Error())
	}
	params := strings.Split(response, " ")
	switch params[1] {
	case "FAIL":
		fmt.Println(os.Args[0] + ": start service failed")
		return false
	case "SUCCESS":
		return true
	default:
		fmt.Println(os.Args[0] + ": unknown error")
		return false
	}
}

func StopService(conn net.Conn, service string) bool {
	conn.Write([]byte("STOP " + service + "\n"))
	response, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println(os.Args[0] + ": error: " + err.Error())
	}
	response = strings.Trim(response, "\n")
	params := strings.Split(response, " ")
	switch params[1] {
	case "FAIL":
		fmt.Println(os.Args[0] + ": stop service failed. Maybe, service already stopped")
		return false
	case "SUCCESS":
		return true
	default:
		fmt.Println(os.Args[0] + ": unknown error")
		return false
	}
}

func SendPing(conn net.Conn) bool {
	conn.Write([]byte("PING\n"))
	response, err := bufio.NewReader(conn).ReadString('\n')
	response = strings.Trim(response, "\n")
	if err != nil {
		fmt.Println(os.Args[0] + ": error: " + err.Error())
		return false
	}

	fmt.Println("Response:", response)
	return true
}
