# ssh_client
The goal of this SSH script is to use host and cmd files to change and validate device configuration.

In the code you will add your login credentials and hostfile as parameters of the connect function in main.go.
 
You can call the hostfile what ever you want. It is neccesary to include the connecting port in the file.

hostfile.txt
- hostname:port
- ip_address:port

Main.go calls unique cmd files for each device using the following naming standard to determine the commands to apply to each device.


Commands files:

File name format:
- "file_"+"hostname"+":ssh_Port"+".cfg" 
- or -
- "file_"+"ip address"+":ssh_Port"+".cfg"

Eg.
- file_core_r1:22.cfg
- file_10.0.1.100:8181.cfg
 
By using the "file_" in front of the name your able to use IP address as well has hostnames.

In sum to use this package you will need to do the following:
```
- Create commands files for each device using naming format provided.
- Create hostfile and add it and your login parameters to the body of func main() in main.go
- If you have multiple logins you can create groups to put the devices in and use those for host files.

```

Example provided:
--------------------

#Package files

```
$ ls -l
total 36
-rw-r--r-- 1 pi pi  185 Dec 15 18:15 file_fastxe:22.cfg
-rw-r--r-- 1 pi pi  176 Dec 15 18:15 file_nxos:8181.cfg
-rw-r--r-- 1 pi pi  103 Dec 15 18:24 go.mod
-rw-r--r-- 1 pi pi 1045 Dec 15 18:24 go.sum
-rw-r--r-- 1 pi pi   10 Dec 15 18:17 group1.txt
-rw-r--r-- 1 pi pi   10 Dec 15 18:25 group2.txt
-rw-r--r-- 1 pi pi 1891 Dec 15 18:27 main.go
-rw-r--r-- 1 pi pi 7211 Dec 15 18:39 README.md
 $ 


# host_files
```
$ cat group1.txt 
fastxe:22
$ cat group2.txt 
nxos:8181

```
# cmds for host1
```
$ cat file_fastxe\:22.cfg
sh ip int brief
config t
interface loopback74
 description Script_test
 ip address 74.74.74.74 255.255.255.255
 exit
exit
show ip int brief
config t
no interface loopback 74
exit
exit


```

# cmds for host2

```
$ cat file_nxos\:8181.cfg         
show ip int brief
config t
interface loopback 75
 description Script_test
 ip address 75.75.75.75/32
exit
exit
show ip int brief
config t
 no interface loopback 75
 exit
exit
$ 
````
-------------------
# Results of the program

```
$ go run main.go 
Connecting to Group1 hosts:
[fastxe:22]


This is the config file named:file_fastxe:22.cfg





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
Loopback1              10.10.1.1       YES other  up                    up      
Loopback2              10.20.1.1       YES other  up                    up      
Loopback10             100.100.100.10  YES manual up                    up      
Loopback20             22.22.22.22     YES manual up                    up      
Loopback30             33.33.33.33     YES manual up                    up      
Loopback100            172.16.1.101    YES manual up                    up      
Loopback101            unassigned      YES unset  up                    up      
Loopback140            10.12.12.14     YES manual up                    up      
Loopback141            10.12.12.15     YES manual up                    up      
Loopback200            172.31.1.200    YES manual up                    up      
csr1000v-1#config t
Enter configuration commands, one per line.  End with CNTL/Z.
csr1000v-1(config)#interface loopback74
csr1000v-1(config-if)# description Script_test
csr1000v-1(config-if)# ip address 74.74.74.74 255.255.255.255
csr1000v-1(config-if)# exit
csr1000v-1(config)#exit
csr1000v-1#show ip int brief
Interface              IP-Address      OK? Method Status                Protocol
GigabitEthernet1       10.10.20.48     YES NVRAM  up                    up      
GigabitEthernet2       unassigned      YES NVRAM  administratively down down    
GigabitEthernet3       unassigned      YES NVRAM  administratively down down    
Loopback1              10.10.1.1       YES other  up                    up      
Loopback2              10.20.1.1       YES other  up                    up      
Loopback10             100.100.100.10  YES manual up                    up      
Loopback20             22.22.22.22     YES manual up                    up      
Loopback30             33.33.33.33     YES manual up                    up      
Loopback74             74.74.74.74     YES manual up                    up      
Loopback100            172.16.1.101    YES manual up                    up      
Loopback101            unassigned      YES unset  up                    up      
Loopback140            10.12.12.14     YES manual up                    up      
Loopback141            10.12.12.15     YES manual up                    up      
Loopback200            172.31.1.200    YES manual up                    up      
csr1000v-1#config t
Enter configuration commands, one per line.  End with CNTL/Z.
csr1000v-1(config)#no interface loopback 74
csr1000v-1(config)#exit
csr1000v-1#exit
Connecting to Group2 hosts:
[nxos:8181]


This is the config file named:file_nxos:8181.cfg





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
$
```
