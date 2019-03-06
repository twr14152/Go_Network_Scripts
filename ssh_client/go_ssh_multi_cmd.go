// This is a multi-cmd script
// (c) 2019 Todd Riemenschneider

package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
)

func main() {
	user := "root"
	pass := "D_Vay!_10&"
	targethost := "ios-xe-mgmt.cisco.com:8181"

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(pass),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	conn, err := ssh.Dial("tcp", targethost, config)
	if err != nil {
		log.Fatal("Failed to dial: ", err)
	}
	sess, err := conn.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}
	stdin, _ := sess.StdinPipe()
	sess.Stdout = os.Stdout
	sess.Stderr = os.Stderr
	sess.Shell()
	cmds := []string{
		"config t",
		"interface loopback 71",
		"description golang_test",
		"exit",
		"exit",
		"show run int loopback 71",
		"exit"}
	for _, cmd := range cmds {
		fmt.Fprintf(stdin, "%s\n", cmd)
	}
	sess.Wait()
	sess.Close()
}
