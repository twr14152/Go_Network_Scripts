/*This is a work in progress!!

*/

package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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

func call(url string, cmds []string, format string) *http.Response {
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
	// Use this set up because its easy for testing in lab
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	cmds := strings.Split(s, ",")
	fmt.Println(cmds)
	target := "https://arista:arista@192.168.1.117:8443/command-api/"
	resp := call(target, cmds, "json")
	fmt.Println(resp)

