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
	//Apply the tls cert bypass to the client variable
	client := &http.Client{Transport: certInfo}
	for _, host := range hostList {
		//Just some code to see what the configs look like in byte code and str fmt
		fmt.Println("The config being loaded:\n", /*config, "\n",*/ string(config))
		// creating variable URL make the http.NewRequest statement easier to read
		url := "https://"+host+"/restconf/data/ietf-interfaces:interfaces"
		// The POST method is used to apply the (config) variable
		req, _ := http.NewRequest("POST", url, bytes.NewBuffer(config))
		// Authentication Parameters
		req.SetBasicAuth("developer", "C1sco12345")
		// Header additions are required for the POST to work with restconf
		req.Header.Add("Content-Type", "application/yang-data+json")
		req.Header.Add("Accept", "application/yang-data+json")
		// client.Do actually kicks off the process by sending http request 
		res, _ := client.Do(req)
		fmt.Println("Connecting to host: ", host)
        	// client must close response body when finished
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
        	// This will print either created (success 201) or conflict (failure 409)  
		fmt.Println("Device Response:\n ",(res))
		// Print the body of the response as a string
		fmt.Println("This is the output of body: ", string(body))
	}	
}
