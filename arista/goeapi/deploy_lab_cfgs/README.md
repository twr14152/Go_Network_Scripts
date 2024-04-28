### Deployed lab BGP routing using main.go script

```
twr14152@DESKTOP-S55FNN9:~/code_folder/go_folder/misc/goeapi/deploy_lab_cfgs$ go run main.go
Connected to ceos1



Pre-Change state:


[map[command:show running-config result:! Command: show running-config
! device: ceos1 (cEOSLab, EOS-4.30.5M-35156751.4305M (engineering build))
!
no aaa root
!
username arista privilege 15 secret sha512 $6$6bnw2XC1I0yqRO1h$DhTS4uj2JKHDitpGcuPTRkZL39ziF8Jcx8YattYV.tEAjN8xUWEwNucOvVJgD./6AerWFT27HsZrvi.sN2faJ/
!
transceiver qsfp default-mode 4x10G
!
service routing protocols model multi-agent
!
agent PowerManager shutdown
agent LedPolicy shutdown
agent Thermostat shutdown
agent PowerFuse shutdown
agent StandbyCpld shutdown
agent LicenseManager shutdown
!
hostname ceos1
!
spanning-tree mode mstp
!
system l1
   unsupported speed action error
   unsupported error-correction action error
!
management api http-commands
   no shutdown
!
interface Ethernet1
   no switchport
   ip address 10.0.0.1/24
!
interface Ethernet2
!
interface Ethernet3
   no switchport
   ip address 157.130.1.1/30
!
interface Loopback0
   description router_id
   ip address 1.1.1.1/32
!
interface Management0
   ip address 172.17.0.2/16
!
ip routing
!
ip route 0.0.0.0/0 172.17.0.1
!
router ospf 1
   router-id 1.1.1.1
   network 1.1.1.1/32 area 0.0.0.0
   network 10.0.0.0/24 area 0.0.0.0
   max-lsa 12000
!
end]] <nil>
<nil>



Post-change state:


[map[command:show running-config result:! Command: show running-config
! device: ceos1 (cEOSLab, EOS-4.30.5M-35156751.4305M (engineering build))
!
no aaa root
!
username arista privilege 15 secret sha512 $6$6bnw2XC1I0yqRO1h$DhTS4uj2JKHDitpGcuPTRkZL39ziF8Jcx8YattYV.tEAjN8xUWEwNucOvVJgD./6AerWFT27HsZrvi.sN2faJ/
!
transceiver qsfp default-mode 4x10G
!
service routing protocols model multi-agent
!
agent PowerManager shutdown
agent LedPolicy shutdown
agent Thermostat shutdown
agent PowerFuse shutdown
agent StandbyCpld shutdown
agent LicenseManager shutdown
!
hostname ceos1
!
spanning-tree mode mstp
!
system l1
   unsupported speed action error
   unsupported error-correction action error
!
management api http-commands
   no shutdown
!
interface Ethernet1
   no switchport
   ip address 10.0.0.1/24
!
interface Ethernet2
!
interface Ethernet3
   no switchport
   ip address 157.130.1.1/30
!
interface Loopback0
   description router_id
   ip address 1.1.1.1/32
!
interface Management0
   ip address 172.17.0.2/16
!
ip routing
!
ip route 0.0.0.0/0 172.17.0.1
!
router bgp 100
   router-id 1.1.1.1
   neighbor 2.2.2.2 remote-as 100
   neighbor 2.2.2.2 update-source Loopback0
   neighbor 157.130.1.2 remote-as 300
   network 1.1.1.1/32
   redistribute connected
!
router ospf 1
   router-id 1.1.1.1
   network 1.1.1.1/32 area 0.0.0.0
   network 10.0.0.0/24 area 0.0.0.0
   max-lsa 12000
!
end]] <nil>
[map[command:show ip interface brief | json result:{
    "interfaces": {
        "Ethernet1": {
            "name": "Ethernet1",
            "lineProtocolStatus": "up",
            "interfaceStatus": "connected",
            "mtu": 1500,
            "ipv4Routable240": false,
            "ipv4Routable0": false,
            "interfaceAddress": {
                "ipAddr": {
                    "address": "10.0.0.1",
                    "maskLen": 24
                }
            },
            "nonRoutableClassEIntf": false
        },
        "Ethernet3": {
            "name": "Ethernet3",
            "lineProtocolStatus": "up",
            "interfaceStatus": "connected",
            "mtu": 1500,
            "ipv4Routable240": false,
            "ipv4Routable0": false,
            "interfaceAddress": {
                "ipAddr": {
                    "address": "157.130.1.1",
                    "maskLen": 30
                }
            },
            "nonRoutableClassEIntf": false
        },
        "Loopback0": {
            "name": "Loopback0",
            "lineProtocolStatus": "up",
            "interfaceStatus": "connected",
            "mtu": 65535,
            "ipv4Routable240": false,
            "ipv4Routable0": false,
            "interfaceAddress": {
                "ipAddr": {
                    "address": "1.1.1.1",
                    "maskLen": 32
                }
            },
            "nonRoutableClassEIntf": false
        },
        "Management0": {
            "name": "Management0",
            "lineProtocolStatus": "up",
            "interfaceStatus": "connected",
            "mtu": 1500,
            "ipv4Routable240": false,
            "ipv4Routable0": false,
            "interfaceAddress": {
                "ipAddr": {
                    "address": "172.17.0.2",
                    "maskLen": 16
                }
            },
            "nonRoutableClassEIntf": false
        }
    }
}] map[command:show ip ospf neighbor | json result:{
    "vrfs": {
        "default": {
            "instList": {
                "1": {
                    "ospfNeighborEntries": [
                        {
                            "routerId": "2.2.2.2",
                            "interfaceAddress": "10.0.0.2",
                            "interfaceName": "Ethernet1",
                            "priority": 1,
                            "adjacencyState": "full",
                            "drState": "DR",
                            "options": {
                                "multitopologyCapability": false,
                                "externalRoutingCapability": true,
                                "multicastCapability": false,
                                "nssaCapability": false,
                                "linkLocalSignaling": false,
                                "demandCircuitsSupport": false,
                                "opaqueLsaSupport": false,
                                "doNotUseInRouteCalc": false
                            },
                            "inactivity": 1714314588.3374653,
                            "details": {
                                "areaId": "0.0.0.0",
                                "designatedRouter": "10.0.0.2",
                                "backupDesignatedRouter": "10.0.0.1",
                                "numberOfStateChanges": 7,
                                "stateTime": 1714311501.337507,
                                "inactivityDefers": 0,
                                "retransmissionCount": 1,
                                "bfdState": "adminDown",
                                "bfdRequestSent": false,
                                "grHelperTimer": null,
                                "grNumAttempts": 0,
                                "grLastRestartTime": null
                            }
                        }
                    ]
                }
            }
        }
    }
}] map[command:show ip bgp summary | json result:{
    "vrfs": {
        "default": {
            "vrf": "default",
            "routerId": "1.1.1.1",
            "asn": "100",
            "peers": {
                "2.2.2.2": {
                    "version": 4,
                    "msgReceived": 0,
                    "msgSent": 0,
                    "inMsgQueue": 0,
                    "outMsgQueue": 0,
                    "asn": "100",
                    "prefixAccepted": 0,
                    "prefixReceived": 0,
                    "upDownTime": 1714314539.193897,
                    "underMaintenance": false,
                    "peerState": "Active"
                },
                "157.130.1.2": {
                    "version": 4,
                    "msgReceived": 0,
                    "msgSent": 0,
                    "inMsgQueue": 0,
                    "outMsgQueue": 0,
                    "asn": "300",
                    "prefixAccepted": 0,
                    "prefixReceived": 0,
                    "upDownTime": 1714314539.192227,
                    "underMaintenance": false,
                    "peerState": "Active"
                }
            }
        }
    }
}]] <nil>




Connected to ceos2



Pre-Change state:


[map[command:show running-config result:! Command: show running-config
! device: ceos2 (cEOSLab, EOS-4.30.5M-35156751.4305M (engineering build))
!
no aaa root
!
username arista privilege 15 secret sha512 $6$BnWvgv1XhHbAmWJr$j8Qfhf5/6NvrLx4R12IT1HMCB8AJYwufxJajG4dOA6uM04OR1gzIuzR/h8vq3ce/KghLPFFS0SuPEGr2ihgxr1
!
transceiver qsfp default-mode 4x10G
!
service routing protocols model multi-agent
!
agent PowerManager shutdown
agent LedPolicy shutdown
agent Thermostat shutdown
agent PowerFuse shutdown
agent StandbyCpld shutdown
agent LicenseManager shutdown
!
hostname ceos2
!
spanning-tree mode mstp
!
system l1
   unsupported speed action error
   unsupported error-correction action error
!
management api http-commands
   no shutdown
!
interface Ethernet1
   no switchport
   ip address 10.0.0.2/24
!
interface Ethernet2
!
interface Ethernet3
   no switchport
   ip address 157.130.2.1/30
!
interface Loopback0
   description router_id
   ip address 2.2.2.2/32
!
interface Management0
   ip address 172.17.0.3/16
!
ip routing
!
ip route 0.0.0.0/0 172.17.0.1
!
router ospf 1
   router-id 2.2.2.2
   network 2.2.2.2/32 area 0.0.0.0
   network 10.0.0.0/24 area 0.0.0.0
   max-lsa 12000
!
end]] <nil>
<nil>



Post-change state:


[map[command:show running-config result:! Command: show running-config
! device: ceos2 (cEOSLab, EOS-4.30.5M-35156751.4305M (engineering build))
!
no aaa root
!
username arista privilege 15 secret sha512 $6$BnWvgv1XhHbAmWJr$j8Qfhf5/6NvrLx4R12IT1HMCB8AJYwufxJajG4dOA6uM04OR1gzIuzR/h8vq3ce/KghLPFFS0SuPEGr2ihgxr1
!
transceiver qsfp default-mode 4x10G
!
service routing protocols model multi-agent
!
agent PowerManager shutdown
agent LedPolicy shutdown
agent Thermostat shutdown
agent PowerFuse shutdown
agent StandbyCpld shutdown
agent LicenseManager shutdown
!
hostname ceos2
!
spanning-tree mode mstp
!
system l1
   unsupported speed action error
   unsupported error-correction action error
!
management api http-commands
   no shutdown
!
interface Ethernet1
   no switchport
   ip address 10.0.0.2/24
!
interface Ethernet2
!
interface Ethernet3
   no switchport
   ip address 157.130.2.1/30
!
interface Loopback0
   description router_id
   ip address 2.2.2.2/32
!
interface Management0
   ip address 172.17.0.3/16
!
ip routing
!
ip route 0.0.0.0/0 172.17.0.1
!
router bgp 100
   router-id 2.2.2.2
   neighbor 1.1.1.1 remote-as 100
   neighbor 1.1.1.1 update-source Loopback0
   neighbor 157.130.2.2 remote-as 400
   network 2.2.2.2/32
   redistribute connected
!
router ospf 1
   router-id 2.2.2.2
   network 2.2.2.2/32 area 0.0.0.0
   network 10.0.0.0/24 area 0.0.0.0
   max-lsa 12000
!
end]] <nil>
[map[command:show ip interface brief | json result:{
    "interfaces": {
        "Ethernet1": {
            "name": "Ethernet1",
            "lineProtocolStatus": "up",
            "interfaceStatus": "connected",
            "mtu": 1500,
            "ipv4Routable240": false,
            "ipv4Routable0": false,
            "interfaceAddress": {
                "ipAddr": {
                    "address": "10.0.0.2",
                    "maskLen": 24
                }
            },
            "nonRoutableClassEIntf": false
        },
        "Ethernet3": {
            "name": "Ethernet3",
            "lineProtocolStatus": "up",
            "interfaceStatus": "connected",
            "mtu": 1500,
            "ipv4Routable240": false,
            "ipv4Routable0": false,
            "interfaceAddress": {
                "ipAddr": {
                    "address": "157.130.2.1",
                    "maskLen": 30
                }
            },
            "nonRoutableClassEIntf": false
        },
        "Loopback0": {
            "name": "Loopback0",
            "lineProtocolStatus": "up",
            "interfaceStatus": "connected",
            "mtu": 65535,
            "ipv4Routable240": false,
            "ipv4Routable0": false,
            "interfaceAddress": {
                "ipAddr": {
                    "address": "2.2.2.2",
                    "maskLen": 32
                }
            },
            "nonRoutableClassEIntf": false
        },
        "Management0": {
            "name": "Management0",
            "lineProtocolStatus": "up",
            "interfaceStatus": "connected",
            "mtu": 1500,
            "ipv4Routable240": false,
            "ipv4Routable0": false,
            "interfaceAddress": {
                "ipAddr": {
                    "address": "172.17.0.3",
                    "maskLen": 16
                }
            },
            "nonRoutableClassEIntf": false
        }
    }
}] map[command:show ip ospf neighbor | json result:{
    "vrfs": {
        "default": {
            "instList": {
                "1": {
                    "ospfNeighborEntries": [
                        {
                            "routerId": "1.1.1.1",
                            "interfaceAddress": "10.0.0.1",
                            "interfaceName": "Ethernet1",
                            "priority": 1,
                            "adjacencyState": "full",
                            "drState": "BDR",
                            "options": {
                                "multitopologyCapability": false,
                                "externalRoutingCapability": true,
                                "multicastCapability": false,
                                "nssaCapability": false,
                                "linkLocalSignaling": false,
                                "demandCircuitsSupport": false,
                                "opaqueLsaSupport": false,
                                "doNotUseInRouteCalc": false
                            },
                            "inactivity": 1714314604.8482008,
                            "details": {
                                "areaId": "0.0.0.0",
                                "designatedRouter": "10.0.0.2",
                                "backupDesignatedRouter": "10.0.0.1",
                                "numberOfStateChanges": 6,
                                "stateTime": 1714311501.8482406,
                                "inactivityDefers": 0,
                                "retransmissionCount": 1,
                                "bfdState": "adminDown",
                                "bfdRequestSent": false,
                                "grHelperTimer": null,
                                "grNumAttempts": 0,
                                "grLastRestartTime": null
                            }
                        }
                    ]
                }
            }
        }
    }
}] map[command:show ip bgp summary | json result:{
    "vrfs": {
        "default": {
            "vrf": "default",
            "routerId": "2.2.2.2",
            "asn": "100",
            "peers": {
                "1.1.1.1": {
                    "version": 4,
                    "msgReceived": 4,
                    "msgSent": 4,
                    "inMsgQueue": 0,
                    "outMsgQueue": 0,
                    "asn": "100",
                    "prefixAccepted": 4,
                    "prefixReceived": 4,
                    "upDownTime": 1714314555.991261,
                    "underMaintenance": false,
                    "peerState": "Established"
                },
                "157.130.2.2": {
                    "version": 4,
                    "msgReceived": 0,
                    "msgSent": 0,
                    "inMsgQueue": 0,
                    "outMsgQueue": 0,
                    "asn": "400",
                    "prefixAccepted": 0,
                    "prefixReceived": 0,
                    "upDownTime": 1714314554.778451,
                    "underMaintenance": false,
                    "peerState": "Active"
                }
            }
        }
    }
}]] <nil>




Connected to ceos3



Pre-Change state:


[map[command:show running-config result:! Command: show running-config
! device: ceos3 (cEOSLab, EOS-4.30.5M-35156751.4305M (engineering build))
!
no aaa root
!
username arista privilege 15 secret sha512 $6$h6zmybe/ny3jqx3T$CJRiRk5zYjn/q/01vgHFe6INjSKjZFJRoNV42O59VA41adxJKfb2MpgFR9Ry4zyYo9JODn/G5RKUxpo7da/pB/
!
transceiver qsfp default-mode 4x10G
!
service routing protocols model multi-agent
!
agent PowerManager shutdown
agent LedPolicy shutdown
agent Thermostat shutdown
agent PowerFuse shutdown
agent StandbyCpld shutdown
agent LicenseManager shutdown
!
hostname ceos3
!
spanning-tree mode mstp
!
system l1
   unsupported speed action error
   unsupported error-correction action error
!
management api http-commands
   no shutdown
!
interface Ethernet1
   no switchport
   ip address 157.130.1.2/30
!
interface Ethernet2
!
interface Loopback0
   description router_id
   ip address 3.3.3.3/32
!
interface Management0
   ip address 172.17.0.4/16
!
ip routing
!
ip route 0.0.0.0/0 172.17.0.1
!
router ospf 1
   router-id 3.3.3.3
   network 3.3.3.3/32 area 0.0.0.0
   max-lsa 12000
!
end]] <nil>
<nil>



Post-change state:


[map[command:show running-config result:! Command: show running-config
! device: ceos3 (cEOSLab, EOS-4.30.5M-35156751.4305M (engineering build))
!
no aaa root
!
username arista privilege 15 secret sha512 $6$h6zmybe/ny3jqx3T$CJRiRk5zYjn/q/01vgHFe6INjSKjZFJRoNV42O59VA41adxJKfb2MpgFR9Ry4zyYo9JODn/G5RKUxpo7da/pB/
!
transceiver qsfp default-mode 4x10G
!
service routing protocols model multi-agent
!
agent PowerManager shutdown
agent LedPolicy shutdown
agent Thermostat shutdown
agent PowerFuse shutdown
agent StandbyCpld shutdown
agent LicenseManager shutdown
!
hostname ceos3
!
spanning-tree mode mstp
!
system l1
   unsupported speed action error
   unsupported error-correction action error
!
management api http-commands
   no shutdown
!
interface Ethernet1
   no switchport
   ip address 157.130.1.2/30
!
interface Ethernet2
!
interface Loopback0
   description router_id
   ip address 3.3.3.3/32
!
interface Management0
   ip address 172.17.0.4/16
!
ip routing
!
ip route 0.0.0.0/0 172.17.0.1
!
router bgp 300
   router-id 3.3.3.3
   neighbor 157.130.1.1 remote-as 100
   network 3.3.3.3/32
   redistribute connected
!
router ospf 1
   router-id 3.3.3.3
   network 3.3.3.3/32 area 0.0.0.0
   max-lsa 12000
!
end]] <nil>
[map[command:show ip interface brief | json result:{
    "interfaces": {
        "Ethernet1": {
            "name": "Ethernet1",
            "lineProtocolStatus": "up",
            "interfaceStatus": "connected",
            "mtu": 1500,
            "ipv4Routable240": false,
            "ipv4Routable0": false,
            "interfaceAddress": {
                "ipAddr": {
                    "address": "157.130.1.2",
                    "maskLen": 30
                }
            },
            "nonRoutableClassEIntf": false
        },
        "Loopback0": {
            "name": "Loopback0",
            "lineProtocolStatus": "up",
            "interfaceStatus": "connected",
            "mtu": 65535,
            "ipv4Routable240": false,
            "ipv4Routable0": false,
            "interfaceAddress": {
                "ipAddr": {
                    "address": "3.3.3.3",
                    "maskLen": 32
                }
            },
            "nonRoutableClassEIntf": false
        },
        "Management0": {
            "name": "Management0",
            "lineProtocolStatus": "up",
            "interfaceStatus": "connected",
            "mtu": 1500,
            "ipv4Routable240": false,
            "ipv4Routable0": false,
            "interfaceAddress": {
                "ipAddr": {
                    "address": "172.17.0.4",
                    "maskLen": 16
                }
            },
            "nonRoutableClassEIntf": false
        }
    }
}] map[command:show ip ospf neighbor | json result:{
    "vrfs": {
        "default": {
            "instList": {
                "1": {
                    "ospfNeighborEntries": []
                }
            }
        }
    }
}] map[command:show ip bgp summary | json result:{
    "vrfs": {
        "default": {
            "vrf": "default",
            "routerId": "3.3.3.3",
            "asn": "300",
            "peers": {
                "157.130.1.1": {
                    "version": 4,
                    "msgReceived": 6,
                    "msgSent": 5,
                    "inMsgQueue": 0,
                    "outMsgQueue": 0,
                    "asn": "100",
                    "prefixAccepted": 6,
                    "prefixReceived": 6,
                    "upDownTime": 1714314571.444802,
                    "underMaintenance": false,
                    "peerState": "Established"
                }
            }
        }
    }
}]] <nil>




Connected to ceos4



Pre-Change state:


[map[command:show running-config result:! Command: show running-config
! device: ceos4 (cEOSLab, EOS-4.30.5M-35156751.4305M (engineering build))
!
no aaa root
!
username arista privilege 15 secret sha512 $6$qa4k4GAjHzuJp9SG$zAorcCrx32KfXiPTcuJexqIIv3Y9MIt.goVV6FHWpwbeD2CdRLudKE/8uO0yJurAVorTZishqMsWeHD5awPp71
!
transceiver qsfp default-mode 4x10G
!
service routing protocols model multi-agent
!
agent PowerManager shutdown
agent LedPolicy shutdown
agent Thermostat shutdown
agent PowerFuse shutdown
agent StandbyCpld shutdown
agent LicenseManager shutdown
!
hostname ceos4
!
spanning-tree mode mstp
!
system l1
   unsupported speed action error
   unsupported error-correction action error
!
management api http-commands
   no shutdown
!
interface Ethernet1
   no switchport
   ip address 157.130.2.2/30
!
interface Ethernet2
!
interface Loopback0
   description router_id
   ip address 4.4.4.4/32
!
interface Management0
   ip address 172.17.0.5/16
!
ip routing
!
ip route 0.0.0.0/0 172.17.0.1
!
router ospf 1
   router-id 4.4.4.4
   network 4.4.4.4/32 area 0.0.0.0
   max-lsa 12000
!
end]] <nil>
<nil>



Post-change state:


[map[command:show running-config result:! Command: show running-config
! device: ceos4 (cEOSLab, EOS-4.30.5M-35156751.4305M (engineering build))
!
no aaa root
!
username arista privilege 15 secret sha512 $6$qa4k4GAjHzuJp9SG$zAorcCrx32KfXiPTcuJexqIIv3Y9MIt.goVV6FHWpwbeD2CdRLudKE/8uO0yJurAVorTZishqMsWeHD5awPp71
!
transceiver qsfp default-mode 4x10G
!
service routing protocols model multi-agent
!
agent PowerManager shutdown
agent LedPolicy shutdown
agent Thermostat shutdown
agent PowerFuse shutdown
agent StandbyCpld shutdown
agent LicenseManager shutdown
!
hostname ceos4
!
spanning-tree mode mstp
!
system l1
   unsupported speed action error
   unsupported error-correction action error
!
management api http-commands
   no shutdown
!
interface Ethernet1
   no switchport
   ip address 157.130.2.2/30
!
interface Ethernet2
!
interface Loopback0
   description router_id
   ip address 4.4.4.4/32
!
interface Management0
   ip address 172.17.0.5/16
!
ip routing
!
ip route 0.0.0.0/0 172.17.0.1
!
router bgp 400
   router-id 4.4.4.4
   neighbor 157.130.2.1 remote-as 100
   network 4.4.4.4/32
   redistribute connected
!
router ospf 1
   router-id 4.4.4.4
   network 4.4.4.4/32 area 0.0.0.0
   max-lsa 12000
!
end]] <nil>
[map[command:show ip interface brief | json result:{
    "interfaces": {
        "Ethernet1": {
            "name": "Ethernet1",
            "lineProtocolStatus": "up",
            "interfaceStatus": "connected",
            "mtu": 1500,
            "ipv4Routable240": false,
            "ipv4Routable0": false,
            "interfaceAddress": {
                "ipAddr": {
                    "address": "157.130.2.2",
                    "maskLen": 30
                }
            },
            "nonRoutableClassEIntf": false
        },
        "Loopback0": {
            "name": "Loopback0",
            "lineProtocolStatus": "up",
            "interfaceStatus": "connected",
            "mtu": 65535,
            "ipv4Routable240": false,
            "ipv4Routable0": false,
            "interfaceAddress": {
                "ipAddr": {
                    "address": "4.4.4.4",
                    "maskLen": 32
                }
            },
            "nonRoutableClassEIntf": false
        },
        "Management0": {
            "name": "Management0",
            "lineProtocolStatus": "up",
            "interfaceStatus": "connected",
            "mtu": 1500,
            "ipv4Routable240": false,
            "ipv4Routable0": false,
            "interfaceAddress": {
                "ipAddr": {
                    "address": "172.17.0.5",
                    "maskLen": 16
                }
            },
            "nonRoutableClassEIntf": false
        }
    }
}] map[command:show ip ospf neighbor | json result:{
    "vrfs": {
        "default": {
            "instList": {
                "1": {
                    "ospfNeighborEntries": []
                }
            }
        }
    }
}] map[command:show ip bgp summary | json result:{
    "vrfs": {
        "default": {
            "vrf": "default",
            "routerId": "4.4.4.4",
            "asn": "400",
            "peers": {
                "157.130.2.1": {
                    "version": 4,
                    "msgReceived": 7,
                    "msgSent": 5,
                    "inMsgQueue": 0,
                    "outMsgQueue": 0,
                    "asn": "100",
                    "prefixAccepted": 7,
                    "prefixReceived": 7,
                    "upDownTime": 1714314587.021549,
                    "underMaintenance": false,
                    "peerState": "Established"
                }
            }
        }
    }
}]] <nil>




twr14152@DESKTOP-S55FNN9:~/code_folder/go_folder/misc/goeapi/deploy_lab_cfgs$
```
