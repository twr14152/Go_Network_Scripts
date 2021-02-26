package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
	"log"
	"net"
	"os"

	"github.com/tmc/scp"
)

func getAgent() (agent.Agent, error) {
	agentConn, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK"))
	return agent.NewClient(agentConn), err
}

func scpFunc(dir, file string) {
	f := dir + "/" + file
	src, err := os.Open(f)
	if err != nil {
		fmt.Println(err)
	}
	src.Close()
	user := "developer"
	pass := "C1sco12345"
	host := "sandbox-iosxe-latest-1.cisco.com:22"

	client, err := ssh.Dial("tcp", host, &ssh.ClientConfig{
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
	dest := src.Name()
	err = scp.CopyPath(src.Name(), dest, session)
	if _, err := os.Stat(dest); os.IsNotExist(err) {
		fmt.Printf("no such file or directory: %s", dest)
	} else {
		fmt.Println("success")
	}
}

func createDir(dir string) {
	user := "developer"
	pass := "C1sco12345"
	host := "sandbox-iosxe-latest-1.cisco.com:22"

	client, err := ssh.Dial("tcp", host, &ssh.ClientConfig{
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
	stdin, _ := session.StdinPipe()
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Shell()
	fmt.Fprintf(stdin, "mkdir "+dir+"\n")
	fmt.Fprintf(stdin, "\n\n\n")
	fmt.Fprintf(stdin, "exit\n")
	session.Wait()
	session.Close()
}

func main() {
	dir := "test_dir"
	createDir(dir)
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
		scpFunc(dir, file.Name())
	}

}



/*
pi@raspberrypi:~/Code_folder/go_folder/netOps/scp_client $ ls -l
total 3776
-rw-r--r-- 1 pi pi     257 Feb 25 22:40 go.mod
-rw-r--r-- 1 pi pi    1384 Feb 25 22:40 go.sum
-rwxr-xr-x 1 pi pi 3842756 Feb 25 22:55 scp_client
-rw-r--r-- 1 pi pi    2095 Feb 25 23:57 scp_clientv3.go
drwxr-xr-x 2 pi pi    4096 Feb 25 23:12 test_dir
pi@raspberrypi:~/Code_folder/go_folder/netOps/scp_client $ cd test_dir/
pi@raspberrypi:~/Code_folder/go_folder/netOps/scp_client/test_dir $ ls -l
total 12
-rw-r--r-- 1 pi pi 8 Feb 25 23:12 text1.txt
-rw-r--r-- 1 pi pi 7 Feb 25 22:38 text2.txt
-rw-r--r-- 1 pi pi 7 Feb 25 22:39 text3.txt
pi@raspberrypi:~/Code_folder/go_folder/netOps/scp_client/test_dir $ 

pi@raspberrypi:~/Code_folder/go_folder/netOps/scp_client $ go run scp_clientv3.go 

Welcome to the DevNet Sandbox for CSR1000v and IOS XE
 
The following programmability features are already enabled:
  - NETCONF
  - RESTCONF
 
Thanks for stopping by.


csr1000v-1#mkdir test_dir
Create directory filename [test_dir]? 
Created dir bootflash:/test_dir
csr1000v-1#
csr1000v-1#
csr1000v-1#exit
text1.txt
success
text2.txt
success
text3.txt
success
pi@raspberrypi:~/Code_folder/go_folder/netOps/scp_client $ ls -l
total 3776
-rw-r--r-- 1 pi pi     257 Feb 25 22:40 go.mod
-rw-r--r-- 1 pi pi    1384 Feb 25 22:40 go.sum
-rwxr-xr-x 1 pi pi 3842756 Feb 25 22:55 scp_client
-rw-r--r-- 1 pi pi    2095 Feb 25 23:57 scp_clientv3.go
drwxr-xr-x 2 pi pi    4096 Feb 25 23:12 test_dir
pi@raspberrypi:~/Code_folder/go_folder/netOps/scp_client $
*/
