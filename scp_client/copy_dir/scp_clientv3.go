/* 
This script will copy a directory and its contents over to a remote device. 
One of the challenges is that the directory needed to be created on the remote device prior to copying files over.
*/

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
pi@raspberrypi:~/Code_folder/go_folder/netOps/scp_client $ 


//remote device//

csr1000v-1#dir
Directory of bootflash:/

201601  drwx            20480  Feb 26 2021 05:08:57 +00:00  tracelogs
153217  drwx             4096  Feb 26 2021 04:58:12 +00:00  test_dir
241921  drwx             4096  Feb 26 2021 04:47:23 +00:00  exit
80641   drwx             4096  Feb 25 2021 23:49:45 +00:00  .installer
22      -rw-              157  Feb 25 2021 23:49:04 +00:00  csrlxc-cfg.log
137089  drwx             4096  Feb 25 2021 23:49:04 +00:00  license_evlog
19      -rw-             2288  Feb 25 2021 23:49:02 +00:00  cvac.log
18      -rw-               30  Feb 25 2021 23:49:00 +00:00  throughput_monitor_params
15      -rw-             1216  Feb 25 2021 23:48:06 +00:00  mode_event_log
64513   drwx             4096   Sep 1 2020 14:51:38 +00:00  .dbpersist
274177  drwx             4096   Sep 1 2020 14:51:34 +00:00  onep
21      -rw-               16   Sep 1 2020 14:51:32 +00:00  ovf-env.xml.md5
20      -rw-                1   Sep 1 2020 14:51:32 +00:00  .cvac_version
104833  drwx             4096   Sep 1 2020 14:51:29 +00:00  pnp-info
145153  drwx             4096   Sep 1 2020 14:50:48 +00:00  virtual-instance
17      -rwx             1314   Sep 1 2020 14:50:21 +00:00  trustidrootx3_ca.ca
16      -rw-            20109   Sep 1 2020 14:50:21 +00:00  ios_core.p7b
193537  drwx             4096   Sep 1 2020 14:50:18 +00:00  gs_script
40321   drwx             4096   Sep 1 2020 14:50:16 +00:00  core
169345  drwx             4096   Sep 1 2020 14:50:12 +00:00  bootlog_history
161281  drwx             4096   Sep 1 2020 14:50:07 +00:00  .prst_sync
14      -rw-             1105   Sep 1 2020 14:49:08 +00:00  packages.conf
13      -rw-         48321761   Sep 1 2020 14:49:08 +00:00  csr1000v-rpboot.17.03.01a.SPA.pkg
12      -rw-        470611036   Sep 1 2020 14:49:08 +00:00  csr1000v-mono-universalk9.17.03.01a.SPA.pkg
8065    drwx             4096   Sep 1 2020 14:49:03 +00:00  .rollback_timer
11      drwx            16384   Sep 1 2020 14:48:15 +00:00  lost+found

6286540800 bytes total (5434560512 bytes free)
csr1000v-1#cd test
csr1000v-1#cd test_dir
csr1000v-1#dir
Directory of bootflash:/test_dir/

153220  -rw-                7  Feb 26 2021 04:58:12 +00:00  text3.txt
153219  -rw-                7  Feb 26 2021 04:58:11 +00:00  text2.txt
153218  -rw-                8  Feb 26 2021 04:58:10 +00:00  text1.txt

6286540800 bytes total (5434560512 bytes free)
csr1000v-1#
*/
