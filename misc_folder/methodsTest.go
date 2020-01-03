package main

import "fmt"

type Switch struct {
	Name        string
	Interface   string
	IntDescr    string
	Ip          string
	RoutingProt string
	Network     string
}

// Testing Method
func (sw *Switch) confInfo() {
	fmt.Println("hostname", sw.Name)
	fmt.Println(" interface", sw.Interface)
	fmt.Println(sw.IntDescr)
	fmt.Println("", sw.Ip, "255.255.255.255")
	fmt.Println("router", sw.RoutingProt)
	fmt.Println(" network", sw.Network, "255.255.255.0")
}

func main() {
	core1 := Switch{Name: "CORE_SW1", Interface: "loopback0", IntDescr: " management interface CORE_SW1", Ip: "65.90.100.1", RoutingProt: "bgp 65001", Network: "10.0.0.0"}
	core2 := Switch{Name: "CORE_SW2", Interface: "loopback0", IntDescr: " management interface CORE_SW2", Ip: "65.90.100.2", RoutingProt: "bgp 65001", Network: "10.0.0.0"}
	// Using positional parameters for var assignment in Switch struct
	core3 := Switch{"CORE_SW3", "Loopback1", " Test Loopback on CORE_SW3", "10.10.10.2", "ospf 100", "10.10.10.0"}

	core1.confInfo()
	fmt.Println("+++++++++++++++++++")
	core2.confInfo()
	fmt.Println("+++++++++++++++++++")
	core3.confInfo()
}
