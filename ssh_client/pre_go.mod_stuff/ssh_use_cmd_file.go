// This ssh script will use a command file to push commands/configs to a Cisco IOS device
// (c) 2019 Todd Riemenschneider

package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
        "bufio"
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
        fh, _ := os.Open("cmd_file.txt")
        scanner := bufio.NewScanner(fh)
        scanner.Split(bufio.ScanLines)
        var lines []string
        for scanner.Scan() {
             lines = append(lines, scanner.Text())
         }
         fh.Close()
	for _, line := range lines {
		fmt.Fprintf(stdin, "%s\n", line)
	}
	sess.Wait()
	sess.Close()
}
