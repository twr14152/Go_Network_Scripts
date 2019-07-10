// This ssh script will use standard input to push commands/configs to a Cisco IOS device
// To run 'cat cmds_file.txt | go run ssh_use_stdin.go'
// This functionality would allow you to use any commands file without having to rewrite commands file or script. 
// (c) 2019 Todd Riemenschneider

package main

import (
	"bufio"
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
	// Below are the lines of code that will allow stdin to push cmds
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	for _, line := range lines {
		fmt.Fprintf(stdin, "%s\n", line)
	}
	// End of code for this functionality 
	sess.Wait()
	sess.Close()

}
