package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	cmdstring := strings.Join(os.Args[1:], " ")
	fmt.Println(cmdstring)
	cmdslice := strings.Split(cmdstring, ",")
	fmt.Println(cmdslice)
}
