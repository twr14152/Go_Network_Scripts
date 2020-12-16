// This ssh script will use host_file.txt to determine who to 
// login to, and the cmd_file1.txt to determine which commands
// to run.
// (c) 2019 Todd Riemenschneider

package main

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
)

var hostList []string

func loginHosts() {
	hf, _ := os.Open("host_file001.txt")
	scanner := bufio.NewScanner(hf)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		hostList = append(hostList, scanner.Text())
		// scaffolding for troubleshooting
		//       fmt.Println(hostList)
	}
}
func main() {
	loginHosts()
	//scaffolding for troubleshooting
	//for _, host := range hostList {
	//		fmt.Println(host)
	for _, host := range hostList {
		user := "developer"
		pass := "C1sco12345"
		targethost := host

		config := &ssh.ClientConfig{
			User: user,
			Auth: []ssh.AuthMethod{
				ssh.Password(pass),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}
		conn, err := ssh.Dial("tcp", targethost, config)
		//Scaffolding for troubleshooting
		//fmt.Println(conn)
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
		cmds, _ := os.Open("cmd_file001.txt")
		scanner := bufio.NewScanner(cmds)
		scanner.Split(bufio.ScanLines)
		var lines []string
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		cmds.Close()
		for _, line := range lines {
			fmt.Fprintf(stdin, "%s\n", line)
		}
		sess.Wait()
		sess.Close()
	}
}



















































