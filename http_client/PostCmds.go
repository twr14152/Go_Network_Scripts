package main

var config = []byte(`{
		"ietf-interfaces:interface": {
			"name": "Loopback69",
			"description": "Testing Golang RestConf",
			"type": "iana-if-type:softwareLoopback",
			"enabled": true,
			"ietf-ip:ipv4": {
				"address": [
					{
						"ip": "69.69.69.1",
						"netmask": "255.255.255.255"
					}
				]
			}
		}
	}`)
