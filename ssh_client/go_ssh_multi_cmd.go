// This is a multi-cmd script
// (c) 2019 Todd Riemenschneider


package main

import (
        "fmt"
        "os"
        "golang.org/x/crypto/ssh"
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
        conn, _ := ssh.Dial("tcp", targethost, config)
        sess, _ := conn.NewSession()
        stdin, _ := sess.StdinPipe()
        sess.Stdout = os.Stdout
        sess.Stderr = os.Stderr
        sess.Shell()
        cmds := []string{
                "config t",
                "interface loopback 71",
                "description golang_test",
                "exit",
                "exit",
                "show run int loopback 71",
                "exit",}
        for _, cmd := range cmds {
                fmt.Fprintf(stdin, "%s\n", cmd)
        }
        sess.Wait()
        sess.Close()
}
