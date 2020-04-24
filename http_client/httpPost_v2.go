package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"crypto/tls"
	"bytes"
	"bufio"
	"os"
)


var hostList []string

func main() {
	hf, _ := os.Open("host_file002.txt")
	scanner := bufio.NewScanner(hf)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		hostList = append(hostList, scanner.Text())
	}
	certInfo := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: certInfo}
	for _, host := range hostList {
		fmt.Println("The config being loaded:\n", /*config, "\n",*/ string(config))
		url := "https://"+host+"/restconf/data/ietf-interfaces:interfaces"

                // the config variable comes from PostCmds.go file
		req, _ := http.NewRequest("POST", url, bytes.NewBuffer(config))
		req.SetBasicAuth("developer", "C1sco12345")
		req.Header.Add("Content-Type", "application/yang-data+json")
		req.Header.Add("Accept", "application/yang-data+json")
		
                res, _ := client.Do(req)
		fmt.Println("Connecting to host: ", host)
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)

        	// This will print either created (success 201) or conflict (failure 409)  
		fmt.Println("Device Response:\n ",(res))
		fmt.Println("This is the output of body: ", string(body))
	}	
}
