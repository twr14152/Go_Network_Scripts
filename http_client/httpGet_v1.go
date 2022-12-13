// This script will use https to pull router info
// https://host:port/restconf/data/etc:etc will determine what info is gathered
// (c) 2019 Todd Riemenschneider

package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
)

func main() {
	certInfo := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: certInfo}
	url := "https://sandbox-iosxe-latest-1.cisco.com/restconf/data/Cisco-IOS-XE-native:native?content=config&depth=65535"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/yang-data+json")
	req.Header.Add("Accept", "application/yang-data+json")
	req.SetBasicAuth("developer", "C1sco12345")

	res, _ := client.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

