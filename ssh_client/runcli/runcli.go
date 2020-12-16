package runcli

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
