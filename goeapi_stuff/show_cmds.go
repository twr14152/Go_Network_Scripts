package main

import (
	"fmt"
	"github.com/aristanetworks/goeapi"
	"os"
	"strings"
)

func main() {
	// This will run the commands after your program
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	cmds := strings.Split(s, ",")
	fmt.Println(cmds)
	for _, host := range HostList {
		node, err := goeapi.ConnectTo(host)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("#############################")
		fmt.Printf("Connected to %v\n", host)
		fmt.Println(node.Enable(cmds))
		fmt.Println("############################")
	}
}
