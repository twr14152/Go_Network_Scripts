// Restconf configuring an interface
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
			"name": "Loopback72",
			"description": "Testing Golang RestConf",
			"type": "iana-if-type:softwareLoopback",
			"enabled": true,
			"ietf-ip:ipv4": {
				"address": [
					{
						"ip": "72.72.72.1",
						"netmask": "255.255.255.255"
					}
				]
			}
		}
	}`)

	client := &http.Client{Transport: certInfo}
	url := "https://ios-xe-mgmt.cisco.com:9443/restconf/data/ietf-interfaces:interfaces"

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(config))
	req.SetBasicAuth("developer", "C1sco12345")
	req.Header.Add("Content-Type", "application/yang-data+json")
	req.Header.Add("Accept", "application/yang-data+json")
	res, _ := client.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
