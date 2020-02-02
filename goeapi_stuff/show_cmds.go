package main

import (
	"fmt"
	"github.com/aristanetworks/goeapi"
	"os"
	"strings"
)

func main() {
	// This will create the list of the commands that you will run
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	cmds := strings.Split(s, ",")
	fmt.Printf("%T ", cmds)
	fmt.Println(cmds)
	for _, host := range HostList {
		node, err := goeapi.ConnectTo(host)
		fmt.Println("#############################")
		fmt.Printf("Connected to %v\n", host)
		if err != nil {
			panic(err)
		}
		fmt.Println(node.Enable(cmds))
		fmt.Println("############################")
	}
}
