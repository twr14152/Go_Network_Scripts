/*
MIT License

Copyright (c) 2020 Todd Riemenschneider

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package main

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
	"strings"
)

func commands() map[string][]string {
	var count int
	var host string
	fmt.Print("Number of hosts: ")
	fmt.Scanf("%d", &count)
	hostCmds := make(map[string][]string)
	for i := 1; i <= count; i++ {
		fmt.Print("Hostname: ")
		fmt.Scanf("%s", &host)
		fmt.Println()
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("cmds: ")
		cmds, _ := reader.ReadString('\n')
		s := strings.Split(cmds, ",")
		fmt.Println()
		hostCmds[host] = s
	}
	return hostCmds
}

func RunCli(user, pass string) {

	hostData := commands()

	for host, commands := range hostData {
		config := &ssh.ClientConfig{
			User: user,
			Auth: []ssh.AuthMethod{
				ssh.Password(pass),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}
		conn, err := ssh.Dial("tcp", host, config)
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
		for _, v := range commands {
			fmt.Fprintf(stdin, "%s\n", v)
		}
		fmt.Fprintf(stdin, "exit\n")
		fmt.Fprintf(stdin, "exit\n")
		sess.Wait()
		sess.Close()
	}
}

func main() {
	fmt.Println("Connecting to Group1 hosts:")
	RunCli("developer", "C1sco12345")
	//fmt.Println("Connecting to Group2  hosts:")
	//RunCli("username2", "password2")
	//fmt.Println("Connecting to Group3 hosts:")
	//RunCli("username3", "password3")

}
