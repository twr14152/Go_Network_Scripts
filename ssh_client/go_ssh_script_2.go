// This script is used to run show commands on a cisco devnet device
// I used a map to create the command set that will be run on the device
// The for loop allows us to run multiple commands on the device.
// (c) 2019 Todd Riemenschneider

package main

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
)

func main() {
        // Used a map to build my command set
	commands := map[string]string{"a":"show ip arp", "b":"show version", "c":"show ip route summ"}
        // SSH Creds to log into remote device
	config := &ssh.ClientConfig{
		User:            "root",
		Auth:            []ssh.AuthMethod{ssh.Password("D_Vay!_10&")},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
        // Used for loop to run through command set, "_" throws away the key data while v represents the actual command
	for _, v := range commands {
               //This is who the script will be logging into
               client, err := ssh.Dial("tcp", "ios-xe-mgmt.cisco.com:8181", config)
               if err != nil {
                        log.Fatal("Failed to dial: ", err)
               }
                //This initiates the connection to the remote device
                session, err := client.NewSession()
                if err != nil {
                        log.Fatal("Failed to create session: ", err)
                }
                defer session.Close()
                // Defines a variable to capture the output of the commands being run
                var cmd_output bytes.Buffer
                session.Stdout = &cmd_output
                fmt.Printf("**** %s *** ", v)
                //This is where the commands actually being sent to the remote device
                session.Run(v)
                // Prints output to screen
                fmt.Println(cmd_output.String())
	fmt.Printf("###########\n")
}
}
