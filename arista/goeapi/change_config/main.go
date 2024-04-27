package main

import (
        "fmt"
        "github.com/aristanetworks/goeapi"
)

//hosts to configure
var HostList = []string{"ceos1", "ceos2", "ceos3", "ceos4", "ceos5", "ceos6"}

func main() {

        //cmds := []string{"show running-config"}
        cmds1 := []string{"show ip interface brief"}

        // Configuration commands
        cfg1 := "interface loopback 55"
        cfg2 := "description testing goeapi"
        cfg3 := "ip address 55.55.55.55/32"
        //cfg4 := "no interface loopback 55"

        for _, host := range HostList {
                node, err := goeapi.ConnectTo(host)
                if err != nil {
                        fmt.Println(err)
                }
                fmt.Printf("Connected to %v\n", host)
                fmt.Printf("\n\n\nPre-Change state:\n\n\n")
                fmt.Println(node.Enable(cmds1))
                fmt.Println(node.Config(cfg1, cfg2, cfg3))
                fmt.Printf("\n\n\nPost change state: \n\n\n")
                fmt.Println(node.Enable(cmds1))
                fmt.Printf("\n\n-------------------------------\n\n")
        }
}
