// Connecting to Restconf API on Cisco router and configuring an interface and enabling ospf for that interface
// (c) 2020 Todd Riemenschneider

package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	certInfo := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	if_config := []byte(`{
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

	rting_config := []byte(`
	{
		  "Cisco-IOS-XE-ospf:ospf": [
    			{
      			"id": 74,
      			"router-id": "74.74.74.1",
      			"network": [
        			{
          				"ip": "74.74.74.1",
          				"mask": "0.0.0.0",
          				"area": 74
        			}
      				]
    			}
 			]
			}
	`)

	fmt.Println("config being applied:\n ", string(if_config))
	client := &http.Client{Transport: certInfo}
	url_intf := "https://ios-xe-mgmt-latest.cisco.com:9443/restconf/data/ietf-interfaces:interfaces"
	req, _ := http.NewRequest("POST", url_intf, bytes.NewBuffer(if_config))
	req.SetBasicAuth("developer", "C1sco12345")
	req.Header.Add("Content-Type", "application/yang-data+json")
	req.Header.Add("Accept", "application/yang-data+json")

	res, _ := client.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Device Response:\n ", (res))
	fmt.Println("This is the output of body: ", string(body))

	fmt.Println("config being applied: \n", string(rting_config))
	url_rting := ("https://ios-xe-mgmt-latest.cisco.com:9443/restconf/data/Cisco-IOS-XE-native:native/router/")
	req1, _ := http.NewRequest("POST", url_rting, bytes.NewBuffer(rting_config))
	req1.SetBasicAuth("developer", "C1sco12345")
	req1.Header.Add("Content-Type", "application/yang-data+json")
	req1.Header.Add("Accept", "application/yang-data+json")

	res1, _ := client.Do(req1)
	defer res1.Body.Close()
	body1, _ := ioutil.ReadAll(res1.Body)
	fmt.Println("Device Response:\n ", (res1))
	fmt.Println("This is the output of body: ", string(body1))

}
