# runscript
The goal of this package is to allow the user use host and cmd files to change and validate device configuration.
You will simply need to import runscript into your main.go file and call the runscript.Connect() to connect to you devices.
All you will need is the username password and the name of the hostfile you create with those common login parameters.
If you have multiple login parameters create multiple hostfiles grouping those with common parameters.

To install you need to go get github.com/twr14152/Go_Network_Scripts/ssh_client/runscript 
Unfortunately I have had to delete go.mod file from the local repository to make the install work.
There's either a bug with go.mod or I dont know what the hell I'm doing. Later is probably the case.

Will update once resolved.

sample code:
main.go
```
package main

import (
    "fmt"
    "github.com/twr14152/Go_Network_Scripts/ssh_client/runscript"
)

func main() {
    fmt.Println("Connecting to Group1 hosts")
    runscript.Connect("user1", "password1", "group1.txt")
    fmt.Println("Connecting to Group2 hosts")
    runscript.Connect("user2", "password2", "group2.txt")
}

```

You can call the hostfile what ever you want. It is neccesary to include the connecting port in the file.

hostfile.txt
- hostname:port
- ip_address:port

# This is key

runscript.go calls unique cmd files for each device using the following naming standard to determine the commands to apply to each device.

Commands files:

File name format:
- "file_"+"hostname"+":ssh_Port"+".cfg" 
- or -
- "file_"+"ip address"+":ssh_Port"+".cfg"

Eg.
- file_core_r1:22.cfg
- file_10.0.1.100:8181.cfg
 
By using the "file_" in front of the name your able to use IP address as well has hostnames.

# In summary to use this package you will need to do the following:
```
- Import runscript into main.go
- Then use runscript.Connect("username", "password", "hostsfile.txt") to your devices 
- Create commands files for each device USING THE FORMAT PROVIDED.
- Create hostfile grouping those hosts that share common login parameters

```

Example provided:
--------------------

#Package files

```
 ls -l
total 20
-rw-r--r-- 1 runner runner 194 Dec 20 15:52 file_131.226.217.143:22.cfg
-rw-r--r-- 1 runner runner 185 Dec 20 16:51 file_64.103.37.14:8181.cfg
-rw-r--r-- 1 runner runner  18 Dec 20 15:53 group1.txt
-rw-r--r-- 1 runner runner  18 Dec 20 16:44 group2.txt
-rw-r--r-- 1 runner runner 318 Dec 20 16:55 main.go

```

# host_files
```
 cat group1.txt 
131.226.217.143:22

 cat group2.txt 
64.103.37.14:8181
 
```
# cmds for host1
```
 cat file_131.226.217.143\:22.cfg 
sh ip int brief
config t
interface loopback74
 description script_test_iosxe_dev
 ip address 74.74.74.74 255.255.255.255
 exit
exit
show ip int brief
config t
no interface loopback 74
exit
exit

```

# cmds for host2
```
cat file_64.103.37.14\:8181.cfg 
show ip int brief
config t
interface loopback 75
 description script_test_nxos_dev
 ip address 75.75.75.75/32
exit
exit
show ip int brief
config t
 no interface loopback 75
 exit
exit
```

-------------------
# Results of the program

```
 go run main.go 
Connecting to Group1 hosts
[131.226.217.143:22]


This is the config file named:file_131.226.217.143:22.cfg





Welcome to the DevNet Sandbox for CSR1000v and IOS XE
 
The following programmability features are already enabled:
  - NETCONF
  - RESTCONF
 
Thanks for stopping by.


csr1000v-1#sh ip int brief
Interface              IP-Address      OK? Method Status                Protocol
GigabitEthernet1       10.10.20.48     YES NVRAM  up                    up      
GigabitEthernet2       unassigned      YES NVRAM  administratively down down    
GigabitEthernet3       unassigned      YES NVRAM  administratively down down    
Loopback105            192.168.1.1     YES manual up                    up      
Loopback106            192.168.1.2     YES manual up                    up      
csr1000v-1#config t
Enter configuration commands, one per line.  End with CNTL/Z.
csr1000v-1(config)#interface loopback74
csr1000v-1(config-if)# description script_test_iosxe_dev
csr1000v-1(config-if)# ip address 74.74.74.74 255.255.255.255
csr1000v-1(config-if)# exit
csr1000v-1(config)#exit
csr1000v-1#show ip int brief
Interface              IP-Address      OK? Method Status                Protocol
GigabitEthernet1       10.10.20.48     YES NVRAM  up                    up      
GigabitEthernet2       unassigned      YES NVRAM  administratively down down    
GigabitEthernet3       unassigned      YES NVRAM  administratively down down    
Loopback74             74.74.74.74     YES manual up                    up      
Loopback105            192.168.1.1     YES manual up                    up      
Loopback106            192.168.1.2     YES manual up                    up      
csr1000v-1#config t
Enter configuration commands, one per line.  End with CNTL/Z.
csr1000v-1(config)#no interface loopback 74
csr1000v-1(config)#exit
csr1000v-1#exit
Connecting to Group2 hosts
[64.103.37.14:8181]


This is the config file named:file_64.103.37.14:8181.cfg





stty: standard input: Inappropriate ioctl for device

IP Interface Status for VRF "default"(1)
Interface            IP Address      Interface Status
Vlan100              172.16.100.1    protocol-down/link-down/admin-down 
Vlan101              172.16.101.1    protocol-down/link-down/admin-down 
Vlan102              172.16.102.1    protocol-down/link-down/admin-down 
Vlan103              172.16.103.1    protocol-down/link-down/admin-down 
Vlan104              172.16.104.1    protocol-down/link-down/admin-down 
Vlan105              172.16.105.1    protocol-down/link-down/admin-down 
Lo1                  172.16.0.1      protocol-up/link-up/admin-up       
Lo98                 10.98.98.1      protocol-up/link-up/admin-up       
Lo99                 10.99.99.1      protocol-up/link-up/admin-up       
Eth1/5               172.16.1.1      protocol-down/link-down/admin-down 

IP Interface Status for VRF "default"(1)
Interface            IP Address      Interface Status
Vlan100              172.16.100.1    protocol-down/link-down/admin-down 
Vlan101              172.16.101.1    protocol-down/link-down/admin-down 
Vlan102              172.16.102.1    protocol-down/link-down/admin-down 
Vlan103              172.16.103.1    protocol-down/link-down/admin-down 
Vlan104              172.16.104.1    protocol-down/link-down/admin-down 
Vlan105              172.16.105.1    protocol-down/link-down/admin-down 
Lo1                  172.16.0.1      protocol-up/link-up/admin-up       
Lo75                 75.75.75.75     protocol-up/link-up/admin-up       
Lo98                 10.98.98.1      protocol-up/link-up/admin-up       
Lo99                 10.99.99.1      protocol-up/link-up/admin-up       
Eth1/5               172.16.1.1      protocol-down/link-down/admin-down 
 
```
