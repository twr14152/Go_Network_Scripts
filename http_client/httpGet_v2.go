// This script will use restconf to pull router interface info off the devices in the host_file001.txt
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

var hostList []string

func main() {
	hf, _ := os.Open("host_file001.txt")
	scanner := bufio.NewScanner(hf)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		hostList = append(hostList, scanner.Text())
	} 
	certInfo := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	UN := "developer"
	PW := "C1sco12345"
	fmt.Println(hostList)
	client := &http.Client{Transport: certInfo}
	
	for i, host := range hostList {
		fmt.Println("connected to host: ", i)
		url := ("https://"+host+"/restconf/data/ietf-interfaces:interfaces/interface=Loopback71")

		req, err := http.NewRequest("GET", url, nil)
		req.SetBasicAuth(UN, PW)
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		deviceOutput, err := ioutil.ReadAll(resp.Body)
		fmt.Println(string(deviceOutput))
	}
}


