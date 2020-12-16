// This is my attempt at building viable ssh script to capture data from show commands on Cisco Devnet devices.
// I used a string array to create the command set to loop through.
// (c) 2019 Todd Riemenschneider

package main

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
)

func main() {

        //Create the command set you want using a string array
        commands :=  []string {"show ip arp", "show version"}

        //Client ssh parameters to log in
	config := &ssh.ClientConfig{
		User:            "root",
		Auth:            []ssh.AuthMethod{ssh.Password("D_Vay!_10&")},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
        //Loop through the command set the "_" is used to ignore the index and the "s" is for string value
	for _, s := range commands {
               //Who you are connecting to
               client, err := ssh.Dial("tcp", "ios-xe-mgmt.cisco.com:8181", config)
               if err != nil {
                        log.Fatal("Failed to dial: ", err)
               }
               // Initiate connection
                session, err := client.NewSession()
                if err != nil {
                        log.Fatal("Failed to create session: ", err)
                }
                defer session.Close()
                //Output variiable
                var cmd_output bytes.Buffer
                session.Stdout = &cmd_output
                //Title the commands being run
                fmt.Printf("%s", s)
                session.Run(s)
                //Print output of commands
                fmt.Println(cmd_output.String())
}
}
