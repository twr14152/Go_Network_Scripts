//This script will use hostfiles to determine what to login to.
//Command files will be used to determine what cmds to run on the devices.
//(c) 2020 Todd Riemenschneider

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
var hostfile string

func loginHosts(hostfile string) {
	hf, _ := os.Open(hostfile)
	scanner := bufio.NewScanner(hf)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		hostList = append(hostList, scanner.Text())
	}
	fmt.Println(hostList)
}

func connect(user, pass, hostfile string) {
	loginHosts(hostfile)
	for _, host := range hostList {

		config := &ssh.ClientConfig{
			User: user,
			Auth: []ssh.AuthMethod{
				ssh.Password(pass),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}

		conn, err := ssh.Dial("tcp", host, config)
		time.Sleep(1)
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
		fmt.Println("\n\nThis is the config file named:" + "file_" + host + ".cfg")
		fmt.Printf("\n\n\n\n")
		cmds, _ := os.Open("file_" + host + ".cfg")
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
	hostList = nil
}
func main() {
	fmt.Println("Connecting to Group1 hosts:")
	connect("user1", "password1", "group1.txt")
	fmt.Println("Connecting to Group2 hosts:")
	connect("user2", "password2", "group2.txt")
}
