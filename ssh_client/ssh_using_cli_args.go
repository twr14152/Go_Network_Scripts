// This is a user interactive ssh script
// You will enter the commands to run at the time you run the script
// (c) 2020 Todd Riemenschneider

package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
	"strings"
)

func main() {
	user := "developer"
	pass := "C1sco12345"
	targethost := "ios-xe-mgmt-latest.cisco.com:8181"

	cmdString := strings.Join(os.Args[1:], " ")
	fmt.Println("Commands being run :", cmdString)
	cmds := strings.Split(cmdString, ",")

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
	fmt.Fprintf(stdin, "term len 0\n")
	for _, cmd := range cmds {
		fmt.Fprintf(stdin, "%s\n", cmd)
	}
	fmt.Fprintf(stdin, "exit\n")
	fmt.Fprintf(stdin, "exit\n")
	sess.Wait()
	sess.Close()
}
