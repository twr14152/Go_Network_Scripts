package main

import (
	"fmt"
	//"io/ioutil"
	"log"
	"net"
	"os"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"

	"github.com/tmc/scp"
)

func getAgent() (agent.Agent, error) {
	agentConn, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK"))
	return agent.NewClient(agentConn), err
}

func main() {
	src, err := os.Open("hello.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(src.Name())
	fmt.Printf("the src file is of type %T\n", src)
	src.Close()
	//f, _ := ioutil.TempFile(".", "hello")
	//fmt.Fprintln(f, "scp test file")
	//f.Close()
	//fmt.Printf("f is of type %T\n", f)
	//fmt.Println(f.Name())
	//defer os.Remove(f.Name())
	//defer os.Remove(f.Name() + "-copy")

	//agent, err := getAgent()
	//if err != nil {
	//	log.Fatalln("Failed to connect to SSH_AUTH_SOCK:", err)

	user := "developer"
	pass := "C1sco12345"
	targethost := "ios-xe-mgmt-latest.cisco.com:8181"

	client, err := ssh.Dial("tcp", targethost, &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(pass),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // FIXME: please be more secure in checking host keys
	})
	if err != nil {
		log.Fatalln("Failed to dial:", err)
	}

	session, err := client.NewSession()
	if err != nil {
		log.Fatalln("Failed to create session: " + err.Error())
	}
	//source_file := string(src)
	//fmt.Println(source_file)
	dest := src.Name() + "-copy"
	//fmt.Println("File dest is :", dest)
	//dest := f.Name() + "-copy"
	//err = scp.CopyPath(f.Name(), dest, session)
	err = scp.CopyPath(src.Name(), dest, session)
	if _, err := os.Stat(dest); os.IsNotExist(err) {
		fmt.Printf("no such file or directory: %s", dest)
	} else {
		fmt.Println("success")
	}
}
