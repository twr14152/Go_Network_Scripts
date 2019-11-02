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
	hf, _ := os.Open("host_file001.txt")
	scanner := bufio.NewScanner(hf)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		hostList = append(hostList, scanner.Text())
	}
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
		url := ("https://"+host+"/restconf/data/Cisco-IOS-XE-native:native?content=config&depth=65535")
		// Issues the GET Request with the variable created above
		req, err := http.NewRequest("GET", url, nil)
		req.SetBasicAuth(username, passwd)
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		deviceOutput, err := ioutil.ReadAll(resp.Body)
		// Print output of commands to screen
		fmt.Println(string(deviceOutput))
	}
}

