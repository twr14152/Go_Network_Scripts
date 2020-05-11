// Using the device type struct you can configure mutiple device with unique configurations
// This script is using ceos-lab platform and http Ports specify different hosts
// (c) Todd Riemenschneider 2020

package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Parameters struct {
	Version int      `json:"version"`
	Cmds    []string `json:"cmds"`
	Format  string   `json:"format"`
}

type Request struct {
	Jsonrpc string     `json:"jsonrpc"`
	Method  string     `json:"method"`
	Params  Parameters `json:"params"`
	Id      string     `json:"id"`
}

type device struct {
	Host string
	Cmds []string
}

func connect(url string, cmds []string, format string) *http.Response {
	p := Parameters{1, cmds, format}
	req := Request{"2.0", "runCmds", p, "1"}
	buf, err := json.Marshal(req)
	resp := new(http.Response)
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	if strings.HasPrefix(url, "https") {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client = &http.Client{Transport: tr}
	}
	resp, err = client.Post(url, "application/json", bytes.NewReader(buf))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return resp
}

func main() {
	ceos1 := &device{
		Host: "8443",
		Cmds: []string{"enable", "configure", "interface loopback3", "description ceos1 Lo3", "ip address 3.3.3.3/32", "router bgp 1", "network 3.3.3.3/32"},
	}
	ceos2 := &device{
		Host: "9443",
		Cmds: []string{"enable", "configure", "interface loopback4", "description ceos2 lo4", "ip address 4.4.4.4/32", "router bgp 2", "network 4.4.4.4/32"},
	}
	ceos_lab := []*device{ceos1, ceos2}
	for _, ceos := range ceos_lab {
		url := "https://arista:arista@192.168.1.117:" + ceos.Host + "/command-api/"
		resp := connect(url, ceos.Cmds, "json")
		fmt.Println(resp)
		if resp.StatusCode == http.StatusOK {
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err)
			}
			bodyString := string(bodyBytes)
			fmt.Println(bodyString)
		}

	}
}
