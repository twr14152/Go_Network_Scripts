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
        c1 := "interface eth1"
        c2 := "no switchport"
        c3 := "ip address 10.0.0.1/24"
        c4 := "no shutdown"
        c5 := "router ospf 1"
        c6 := "network 10.0.0.0/24 area 0"
        c7 := "router bgp 100"
        c8 := "neighbor 2.2.2.2 remote-as 100"
        c9 := "neighbor 2.2.2.2 update-source loopback0"
        c10 := "router-id 1.1.1.1"
        c11 := "network 1.1.1.1 mask 255.255.255.255"
        c12 := "interface eth3"
        c13 := "no switchport"
        c14 := "no shutdown"
        c15 := "ip address 157.130.1.1/30"
        c16 := "router bgp 100"
        c17 := "neighbor 157.130.2.2 remote-as 400"
        c18 := "redistribute connected"



        node, err := goeapi.ConnectTo(host)
        if err != nil {
                log.Fatal(err)
        }
        fmt.Printf("Connected to %v\n", host)
        fmt.Printf("\n\n\nPre-Change state:\n\n\n")
        fmt.Println(node.Enable(cmds1))
        fmt.Println(node.ConfigWithErr(c1, c2, c3, c4, c5, c6, c7, c8, c9, c10, c11, c12, c13, c14, c15, c16, c17, c18))
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
        c1 := "interface ethernet 1"
        c2 := "no switchport"
        c3 := "ip address 10.0.0.2/24"
        c4 := "no shutdown"
        c5 := "router ospf 1"
        c6 := "network 10.0.0.0/24 area 0"
        c7 := "router bgp 100"
        c8 := "neighbor 1.1.1.1 remote-as 100"
        c9 := "neighbor 1.1.1.1  update-source loopback 0"
        c10 := "router-id 2.2.2.2"
        c11 := "network 2.2.2.2 mask 255.255.255.255"
        c12 := "interface eth3"
        c13 := "no switchport"
        c14 := "no shutdown"
        c15 := "ip address 157.130.2.1/30"
        c16 := "router bgp 100"
        c17 := "neighbor 157.130.2.2 remote-as 400"
        c18 := "redistribute connected"


        node, err := goeapi.ConnectTo(host)
        if err != nil {
                log.Fatal(err)
        }
        fmt.Printf("Connected to %v\n", host)
        fmt.Printf("\n\n\nPre-Change state:\n\n\n")
        fmt.Println(node.Enable(cmds1))
        fmt.Println(node.ConfigWithErr(c1, c2, c3, c4, c5, c6, c7, c8, c9, c10, c11, c12, c13, c14, c15, c16, c17, c18))
        fmt.Printf("\n\n\nPost-change state: \n\n\n")
        duration := time.Duration(15)*time.Second
        time.Sleep(duration)
        fmt.Println(node.Enable(cmds1))
        fmt.Println(node.Enable(cmds2))
        fmt.Printf("\n\n\n\n")

}

func deploy_ceos3() {

        host := "ceos3"

        cmds1 := []string{"show running-config"}
        cmds2 := []string{"show ip interface brief | json", "show ip ospf neighbor | json", "show ip bgp summary | json"}

        // Configuration commands
        c1 := "interface ethernet 1"
        c2 := "no switchport"
        c3 := "ip address 157.130.1.2/30"
        c4 := "no shutdown"
        c5 := "router bgp 300"
        c6 := "network 3.3.3.3 mask 255.255.255.255"
        c7 := "router-id 3.3.3.3"
        c8 := "neighbor 157.130.1.1 remote-as 100"
        c9 := "redistribute connected"



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



func deploy_ceos4() {

        host := "ceos4"

        cmds1 := []string{"show running-config"}
        cmds2 := []string{"show ip interface brief | json", "show ip ospf neighbor | json", "show ip bgp summary | json"}

        // Configuration commands
        c1 := "interface ethernet 1"
        c2 := "no switchport"
        c3 := "ip address 157.130.2.2/30"
        c4 := "no shutdown"
        c5 := "router bgp 400"
        c6 := "network 4.4.4.4 mask 255.255.255.255"
        c7 := "router-id 4.4.4.4"
        c8 := "neighbor 157.130.2.1 remote-as 100"
        c9 := "redistribute connected"



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


func main() {
        deploy_ceos1()
        deploy_ceos2()
        deploy_ceos3()
        deploy_ceos4()
}
