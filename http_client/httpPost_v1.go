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
	//certInfo is being used to bypass tls cert validation
	certInfo := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	//creates the config variable
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
	//Apply the tls cert bypass to the client variable
	client := &http.Client{Transport: certInfo}
	// creating variable URL make the http.NewRequest statement easier to read
	url := "https://ios-xe-mgmt-latest.cisco.com:9443/restconf/data/ietf-interfaces:interfaces"
	// The POST method is used to apply the (config) variable
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(config))
	// Authentication Parameters
	req.SetBasicAuth("developer", "C1sco12345")
	// Header additions are required for the POST to work with restconf
	req.Header.Add("Content-Type", "application/yang-data+json")
	req.Header.Add("Accept", "application/yang-data+json")
	// client.Do actually kicks off the process by sending http request 
	res, _ := client.Do(req)
        // client must close response body when finished
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
        
	fmt.Println(res)
	// Print the body of the response as a string
	fmt.Println(string(body))

}
