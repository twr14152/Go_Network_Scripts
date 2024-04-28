// These are partial configs for lab build out
package main

import (
        "fmt"
        "time"
        "log"
        "github.com/aristanetworks/goeapi"
)


func deploy_ceos1() {

        host := "ceos1"

        cmds1 := []string{"show running-config"}
        cmds2 := []string{"show ip interface brief | json", "show ip ospf neighbor | json", "show ip bgp summary | json"}

        // Configuration commands
        c1 := "interface ethernet 2"
        c2 := "no switchport"
        c3 := "ip address 10.0.1.1/24"
        c4 := "no shutdown"
        c5 := "router ospf 1"
        c6 := "network 10.0.1.0/24 area 0"
        c7 := "router bgp 100"
        c8 := "neighbor 6.6.6.6 remote-as 100"
        c9 := "neighbor 6.6.6.6 update-source loopback0"
        //c10 := ""
        //c11 := ""
        //c12 := ""
        //c13 := ""


        node, err := goeapi.ConnectTo(host)
        if err != nil {
                log.Fatal(err)
        }
        fmt.Printf("Connected to %v\n", host)
        fmt.Printf("\n\n\nPre-Change state:\n\n\n")
        fmt.Println(node.Enable(cmds1))
        fmt.Println(node.ConfigWithErr(c1, c2, c3, c4, c5, c6, c7, c8, c9))
        fmt.Printf("\n\n\nPost-change state: \n\n\n")
        duration := time.Duration(15)*time.Second
        time.Sleep(duration)
        fmt.Println(node.Enable(cmds1))
        fmt.Println(node.Enable(cmds2))
        fmt.Printf("\n\n\n\n")
}

func deploy_ceos2() {

        host := "ceos2"

        cmds1 := []string{"show running-config"}
        cmds2 := []string{"show ip interface brief | json", "show ip ospf neighbor | json", "show ip bgp summary | json"}

        // Configuration commands
        c1 := "interface ethernet 2"
        c2 := "no switchport"
        c3 := "ip address 10.0.2.2/24"
        c4 := "no shutdown"
        c5 := "router ospf 1"
        c6 := "network 10.0.2.0/24 area 0"
        c7 := "router bgp 100"
        c8 := "neighbor 6.6.6.6 remote-as 100"
        c9 := "neighbor 6.6.6.6 update-source loopback 0"
        //c10 := ""
        //c11 := ""
        //c12 := ""
        //c13 := ""


        node, err := goeapi.ConnectTo(host)
        if err != nil {
                log.Fatal(err)
        }
        fmt.Printf("Connected to %v\n", host)
        fmt.Printf("\n\n\nPre-Change state:\n\n\n")
        fmt.Println(node.Enable(cmds1))
        fmt.Println(node.ConfigWithErr(c1, c2, c3, c4, c5, c6, c7, c8, c9))
        fmt.Printf("\n\n\nPost-change state: \n\n\n")
        duration := time.Duration(15)*time.Second
        time.Sleep(duration)
        fmt.Println(node.Enable(cmds1))
        fmt.Println(node.Enable(cmds2))
        fmt.Printf("\n\n\n\n")

}

func deploy_ceos6() {

        host := "ceos6"

        cmds1 := []string{"show running-config"}
        cmds2 := []string{"show ip interface brief | json", "show ip ospf neighbor | json", "show ip bgp summary | json"}

        // Configuration commands
        c1 := "interface ethernet 1"
        c2 := "no switchport"
        c3 := "ip address 10.0.1.6/24"
        c4 := "no shutdown"
        c5 := "router ospf 1"
        c6 := "network 10.0.1.0/24 area 0"
        c7 := "network 10.0.2.0/24 area 0"
        c8 := "router bgp 100"
        c9 := "router-id 6.6.6.6"
        c10 := "network 6.6.6.6 mask 255.255.255.255"
        c11 := "redistribute connected"
        c12 := "neighbor 1.1.1.1 remote-as 100"
        c13 := "neighbor 1.1.1.1 update-source loopback0"
        c14 := "neighbor 2.2.2.2 remote-as 100"
        c15 := "neighbor 2.2.2.2 update-source loopback0"
        c16 := "interface ethernet 2"
        c17 := "no switchport"
        c18 := "ip address 10.0.2.6/24"
        c19 := "no shutdown"



        node, err := goeapi.ConnectTo(host)
        if err != nil {
                log.Fatal(err)
        }
        fmt.Printf("Connected to %v\n", host)
        fmt.Printf("\n\n\nPre-Change state:\n\n\n")
        fmt.Println(node.Enable(cmds1))
        fmt.Println(node.ConfigWithErr(c1, c2, c3, c4, c5, c6, c7, c8, c9, c10, c11, c12, c13, c14, c15, c16, c17, c18, c19))
        fmt.Printf("\n\n\nPost-change state: \n\n\n")
        duration := time.Duration(15)*time.Second
        time.Sleep(duration)
        fmt.Println(node.Enable(cmds1))
        fmt.Println(node.Enable(cmds2))
        fmt.Printf("\n\n\n\n")
}

func main() {
        deploy_ceos1()
        deploy_ceos2()
        deploy_ceos6()
}


