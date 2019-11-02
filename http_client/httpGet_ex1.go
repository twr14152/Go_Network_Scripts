// This script will use https to pull router info 
// https://host:port/restconf/data/etc:etc will determine what info is gathered
// (c) 2019 Todd Riemenschneider


package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"crypto/tls"
)

func main() {
	certInfo := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: certInfo}
	url := "https://ios-xe-mgmt.cisco.com:9443/restconf//data/Cisco-IOS-XE-native:native?content=config&depth=65535"

	req, _ := http.NewRequest("GET", url, nil)
	req.SetBasicAuth("developer", "C1sco12345")
	res, _ := client.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
