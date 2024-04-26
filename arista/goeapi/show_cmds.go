package main

import (
	"fmt"
	"github.com/aristanetworks/goeapi"
	"os"
	"strings"
)

var HostList = []string{"ceos1", "ceos2", "ceos3", "ceos4", "ceos5", "ceos6"}

func main() {
	// This will run the commands listed after the filename
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
