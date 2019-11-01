// This script will use https to pull router info 
// https://host:port/restconf/data/etc:etc will determine what info is gathered
// (c) 2019 Todd Riemenschneider

package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	connXinfo()
}

func connXinfo() {
	certInfo := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	username := "developer"
	passwd := "C1sco12345"
	client := &http.Client{Transport: certInfo}
	req, err := http.NewRequest("GET", "https://ios-xe-mgmt-latest.cisco.com:9443/restconf/data/ietf-interfaces:interfaces", nil)
	req.SetBasicAuth(username, passwd)
	resp, err := client.Do(req)
	fmt.Println("Connection info:\n", resp)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	fmt.Println("Interfaces:\n", string(bodyText))
}
