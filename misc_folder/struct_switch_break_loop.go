//Structin my stuff playin with switches while breakin loops

package main

import "fmt"

// taking a structure
type device_id struct {
	// taking variables
	name string
	cmds []string
}

// Main Function
func main() {

	// creating the instance of the
	// device_id struct type
	r1 := &device_id{
		name: "cmh1.r1",
		cmds: []string{"config t", "router ospf 1", "network 10.0.0.0/24"},
	}
	r2 := &device_id{
		name: "cmh1.r2",
		cmds: []string{"config t", "router ospf 1", "network 10.0.1.0/24"},
	}
	r3 := &device_id{
		name: "cmh1.r3",
		cmds: []string{"config t", "router ospf 1", "network 10.0.2.0/24"},
	}

	routerId := []string{(*r1).name, (*r2).name, (*r3).name}
	routerCmds := [][]string{(*r1).cmds, (*r2).cmds, (*r3).cmds}

	//fmt.Println(routerId)
	//fmt.Println((*r3).cmds)
	//fmt.Println(routerCmds)

  target:= routerId[1]
  fmt.Println("Switch target is:", target)
Loop:
	for i := 0; i < len(routerId); i++ {
		fmt.Println("Times through loop:", i)
		switch target {
		case routerId[0]:
			fmt.Println("Target is: ", routerId[0])
			for i := range routerCmds[0] {
				fmt.Println(routerCmds[0][i])
			}
			fmt.Println("R1 done")
			break Loop
		case routerId[1]:
			fmt.Println("Target is: ", routerId[1])
			for i := range routerCmds[1] {
				fmt.Println(routerCmds[1][i])
			}
			fmt.Println("R2 done")
			break Loop
		case routerId[2]:
			fmt.Println("Target is: ", routerId[2])
			for i := range routerCmds[2] {
				fmt.Println(routerCmds[2][i])
			}
			fmt.Println("R3 done")
			break Loop

		}
	}
}

