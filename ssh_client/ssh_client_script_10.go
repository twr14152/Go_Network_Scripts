// This ssh script will use host_file.txt to determine who to
// login to, and the cmd_file1.txt to determine which commands
// to run.
// (c) 2020 Todd Riemenschneider

package main

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
	"time"
)

var hostList []string
var user string
var pass string

func loginHosts() {
	hf, _ := os.Open("host_file001.txt")
	scanner := bufio.NewScanner(hf)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		hostList = append(hostList, scanner.Text())
	}
	fmt.Println(hostList)
}

func main() {
	loginHosts()
	for _, host := range hostList {
		if host == "ios-xe-mgmt-latest.cisco.com:8181" || host == "ios-xe-mgmt.cisco.com:8181" {
			user = "developer"
			pass = "C1sco12345"
		} else if host == "sbx-nxos-mgmt.cisco.com:8181" {
			user = "admin"
			pass = "Admin_1234!"
		}
		targethost := host

		config := &ssh.ClientConfig{
			User: user,
			Auth: []ssh.AuthMethod{
				ssh.Password(pass),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}

		conn, err := ssh.Dial("tcp", targethost, config)
		time.Sleep(1)
		//Used for troubleshooting
		//fmt.Println(conn)
		if err != nil {
			log.Fatal("Failed to dial: ", err)
		}
		sess, err := conn.NewSession()
		if err != nil {
			log.Fatal("Failed to create session: ", err)
		}
		stdin, err := sess.StdinPipe()
		sess.Stdout = os.Stdout
		sess.Stderr = os.Stderr
		sess.Shell()
		// cmds file should use host.cfg name standard
		fmt.Println("This is the config file named:" + host + ".cfg")
		fmt.Printf("\n\n\n\n")
		cmds, _ := os.Open(host + ".cfg")
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
		fmt.Fprintf(stdin, "exit\n")
		fmt.Fprintf(stdin, "exit\n")
		sess.Wait()
		sess.Close()
	}
}
