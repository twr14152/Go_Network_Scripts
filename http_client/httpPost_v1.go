// Connecting to Restconf API on Cisco router and configuring an interface
// (c) 2019 Todd Riemenschneider 

package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"crypto/tls"
	"bytes"
)

func main() {
	certInfo := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	config := []byte(`{
		"ietf-interfaces:interface": {
			"name": "Loopback74",
			"description": "Testing Golang RestConf",
			"type": "iana-if-type:softwareLoopback",
			"enabled": true,
			"ietf-ip:ipv4": {
				"address": [
					{
						"ip": "74.74.74.1",
						"netmask": "255.255.255.255"
					}
				]
			}
		}
	}`)
	
	fmt.Println("config in bytes then in str fmt:\n ", config,"\n", string(config))
	client := &http.Client{Transport: certInfo}
	url := "https://ios-xe-mgmt-latest.cisco.com:9443/restconf/data/ietf-interfaces:interfaces"

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(config))
	req.SetBasicAuth("developer", "C1sco12345")
	req.Header.Add("Content-Type", "application/yang-data+json")
	req.Header.Add("Accept", "application/yang-data+json") 

	res, _ := client.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Device Response:\n ",(res))
	fmt.Println("This is the output of body: ", string(body))
}
