package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Usage:", os.Args[0], "ACTION [SERVICE_NAME]")
		fmt.Println("\nActions list:\nstart - Start service\nstop - Stop service\nping - Test request to pinit")
		return
	}

	ActionsHandler(os.Args)
}
