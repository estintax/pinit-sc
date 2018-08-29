package main

import "fmt"

func ActionsHandler(args []string) bool {
	switch args[1] {
	case "ping":
		conn := ConnectToPinit()
		if conn != nil {
			SendPing(conn)
			conn.Close()
			return true
		}
		return false
	case "start":
		if len(args) > 2 {
			conn := ConnectToPinit()
			if conn != nil {
				StartService(conn, args[2])
				conn.Close()
				return true
			}
		} else {
			fmt.Println(args[0] + ": error: missed service name")
			return false
		}
		return false
	default:
		fmt.Println(args[0] + ": error: unknown action")
		return false
	}
}
