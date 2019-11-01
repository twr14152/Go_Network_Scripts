// This script will use https to pull router config info off the devices in the host_file001.txt
// (c) 2019 Todd Riemenschneider

package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// slice that will host the hosts
var hostList []string

func main() {
	// Initiate function to create slice that will contain the target hosts
	loginHosts()
	// Initiate function that will connect to target hosts
	connXinfo()
}
// This func will open host_file001 and read contents into hostList slice
func loginHosts() {
	hf, _ := os.Open("host_file001.txt")
	scanner := bufio.NewScanner(hf)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		hostList = append(hostList, scanner.Text())
	}
}
// This func contains info to connect to target hosts and run restconf command
func connXinfo() {
	// certInfo is the variable used to bypass cert validation process 
	certInfo := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	username := "developer"
	passwd := "C1sco12345"
	fmt.Println(hostList)
	// Apply cert bypass variable to https Transport process 
	client := &http.Client{Transport: certInfo}
	// Loop through hosts in hostList slice
	for i, host := range hostList {
		// i is being used to identify which host is being connected to
		fmt.Println("connected to host: ", i)
		// create variable that includes the host and the restconf command to pull the configs off the device
		addr := ("https://"+host+"/restconf/data/Cisco-IOS-XE-native:native?content=config&depth=65535")
		// Issues the GET Request with the variable created above
		req, err := http.NewRequest("GET", addr, nil)
		req.SetBasicAuth(username, passwd)
		resp, err := client.Do(req)
		fmt.Println(resp)
		if err != nil {
			log.Fatal(err)
		}
		bodyText, err := ioutil.ReadAll(resp.Body)
		// Print output of commands to screen
		fmt.Println(string(bodyText))
	}
}
