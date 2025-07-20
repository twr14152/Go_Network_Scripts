package main

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

var hostList []string

func loginHosts(hostfile string) {
	hostList = nil
	hf, err := os.Open(hostfile)
	if err != nil {
		log.Fatal("Failed to open file: ", err)
	}
	defer hf.Close()

	scanner := bufio.NewScanner(hf)
	for scanner.Scan() {
		hostList = append(hostList, scanner.Text())
	}
}

func Connect(user, pass, hostfile string) {
	loginHosts(hostfile)

	interactiveAuth := ssh.KeyboardInteractive(
		func(user, instruction string, questions []string, echos []bool) ([]string, error) {
			answers := make([]string, len(questions))
			for i := range questions {
				answers[i] = pass
			}
			return answers, nil
		},
	)

	for _, host := range hostList {
		fmt.Printf("\n--- Connecting to %s ---\n", host)

		config := &ssh.ClientConfig{
			User:            user,
			Auth:            []ssh.AuthMethod{interactiveAuth},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			Timeout:         5 * time.Second,
		}

		conn, err := ssh.Dial("tcp", host, config)
		if err != nil {
			log.Printf("Failed to dial %s: %v\n", host, err)
			continue
		}

		sess, err := conn.NewSession()
		if err != nil {
			log.Printf("Failed to create session for %s: %v\n", host, err)
			continue
		}

		modes := ssh.TerminalModes{
			ssh.ECHO:          1,
			ssh.TTY_OP_ISPEED: 14400,
			ssh.TTY_OP_OSPEED: 14400,
		}

		err = sess.RequestPty("xterm", 80, 40, modes)
		if err != nil {
			log.Fatalf("PTY request failed: %v", err)
		}

		stdin, _ := sess.StdinPipe()
		stdout, _ := sess.StdoutPipe()
		sess.Stderr = os.Stderr

		err = sess.Shell()
		if err != nil {
			log.Fatalf("failed to start shell: %v", err)
		}

		reader := stdout

		// Wait for initial prompt
		waitForPrompt(reader, ">")

		// Enter enable mode
		fmt.Fprintf(stdin, "enable\n")
		waitForPrompt(reader, "#")

		// Disable terminal paging
		fmt.Fprintf(stdin, "term len 0\n")
		waitForPrompt(reader, "#")

		// Read commands from file
		cfgFile := "file_" + host + ".cfg"
		fmt.Printf("Sending commands from: %s\n", cfgFile)

		cmds, err := os.Open(cfgFile)
		if err != nil {
			log.Printf("Could not open config file %s: %v\n", cfgFile, err)
			continue
		}
		scanner := bufio.NewScanner(cmds)
		var lines []string
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		cmds.Close()

		for _, line := range lines {
			if strings.TrimSpace(line) == "" {
				continue
			}
			fmt.Printf("[SENDING] %s\n", line)
			fmt.Fprintf(stdin, "%s\n", line)
			time.Sleep(200 * time.Millisecond)
			output := readUntilPrompt(reader, "#")
			fmt.Println("[OUTPUT]")
			fmt.Println(output)
		}

		fmt.Fprintf(stdin, "exit\n")
		sess.Wait()
		sess.Close()
		conn.Close()
	}
}

// Wait for a specific CLI prompt like > or #
func waitForPrompt(reader io.Reader, prompt string) {
	fmt.Printf("[Waiting for prompt: %s]\n", prompt)
	var buffer strings.Builder
	buf := make([]byte, 1)

	for {
		n, err := reader.Read(buf)
		if err != nil {
			log.Fatalf("Error reading from SSH session: %v", err)
		}
		if n > 0 {
			buffer.WriteString(string(buf[:n]))
			if strings.Contains(buffer.String(), prompt) {
				fmt.Println("[PROMPT FOUND]")
				break
			}
		}
	}
}

// Read output until we get back to prompt
func readUntilPrompt(reader io.Reader, prompt string) string {
	var buffer strings.Builder
	buf := make([]byte, 1)

	for {
		n, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatalf("Error reading from SSH session: %v", err)
		}
		if n > 0 {
			char := string(buf[:n])
			buffer.WriteString(char)
			if strings.Contains(buffer.String(), prompt) {
				break
			}
		} else if err == io.EOF {
			break
		}
	}

	return buffer.String()
}

func main() {
	Connect("username", "password", "group1.txt")
}
