// This script will run multiple show commands on a device. 
// The issue is that it needs to reconnect with each command it iterates through.
// Using Map instead of slice or array. May come in handy down the road.

package main

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
)

func main() {

	commands := map[string]string{"a":"show ip arp", "b":"show version"}

	config := &ssh.ClientConfig{
		User:            "root",
		Auth:            []ssh.AuthMethod{ssh.Password("D_Vay!_10&")},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	for _, v := range commands {
            // Issues new connection with each command....Not sure what to think about that.....
               client, err := ssh.Dial("tcp", "ios-xe-mgmt.cisco.com:8181", config)
               if err != nil {
                        log.Fatal("Failed to dial: ", err)
               }
            // Each ClientConn can support multiple interactive sessions,
            // represented by a Session.
                session, err := client.NewSession()
                if err != nil {
                        log.Fatal("Failed to create session: ", err)
                }
                defer session.Close()
                var cmd_output bytes.Buffer
                session.Stdout = &cmd_output
                fmt.Printf("**** %s *** ", v)
                session.Run(v)
                fmt.Println(cmd_output.String())
	fmt.Printf("###########\n")
}
}
