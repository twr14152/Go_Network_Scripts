# Go_Network_Scripts
This is an attempt to try and learn Golang by creating some useful network scripts.


First script was to ssh to Cisco Devnet device and print the output.
It works but its not a real useful script.
```
- ssh_client/can_I_ssh_script.go
```
Second script had some success getting multiple show commands to run on the device.
I used a map which is analogous to a dictionary in python to create my command set. Then used a for loop to run through those commands. 

However it needs to reconnect after each command. Not a fan of that, especially for config mode. 
```
- ssh_client/go_ssh_script_2.go
```
Third script involed using a slice to create the command set and then using a for loop to run through the commands.
Nothing earth shattering but it was an opportunity to create another iterable data set.
```
- ssh_client/go_ssh_script_3.go
```
The fourth SSH script will allow multi-line commands show or config. Previous script used Run() command which has a one command limit per session. This was resolved by using Shell() rather than Run(). 
```
 - ssh_client/ssh_multi_cmd.go 
 
```
The fifth script opens and reads another file to get the commands and then uses a for loop to execute the commands on the remote device.
The benefit of this method is that to make changes to the target host you no longer need to update the script rather the cmd_file.txt 
```
- ssh_client/ssh_use_cmd_file.go
- ssh_client/cmd_file.txt

pi@raspberrypi:~/Coding_Folder $ go run ssh_use_cmd_file.go 

Welcome to the DevNet Always On Sandbox for IOS XE

This is a shared sandbox available for anyone to use to
test APIs, explore features, and test scripts.  Please
keep this in mind as you use it, and respect others use.

The following programmability features are already enabled:
  - NETCONF
  - RESTCONF

Thanks for stopping by.



csr1000v#enable
csr1000v#config t
Enter configuration commands, one per line.  End with CNTL/Z.
csr1000v(config)#interface loopback 72
csr1000v(config-if)#description golang test script
csr1000v(config-if)#exit
csr1000v(config)#exit
csr1000v#show run interface loopback72
Building configuration...

Current configuration : 75 bytes
!
interface Loopback72
 description golang test script
 no ip address
end

csr1000v#
csr1000v#show ip int brief
Interface              IP-Address      OK? Method Status                Protocol
GigabitEthernet1       10.10.20.48     YES NVRAM  up                    up      
GigabitEthernet2       unassigned      YES NVRAM  up                    up      
GigabitEthernet3       unassigned      YES NVRAM  administratively down down    
Loopback71             unassigned      YES unset  up                    up      
Loopback72             unassigned      YES unset  up                    up      
csr1000v#exit
pi@raspberrypi:~/Coding_Folder $ 

```
The 6th Script uses standard input to put the commands into the script to issue commands on the remote device.
The benefit of doing this is you do not need to overwrite a cmds_file to make changes. You simply create a new file with what ever commands you want(naming it what ever you want) and pipe the output to the script. 
How to use script: cat cmd_file.txt | go run ssh_use_stdin.go
```
- ssh_client/ssh_use_stdin.go
- ssh_client/cmd_file.txt


pi@raspberrypi:~/Coding_Folder $ cat cmd_file.txt 
enable
terminal length 0
config t
interface loopback 71
 description golang_script_test
interface loopback 72
 description golang_script_test
 exit
 exit
show ip int brief
show version
show run all | inc ssh
exit
pi@raspberrypi:~/Coding_Folder $

pi@raspberrypi:~/Coding_Folder $ cat cmd_file.txt| go run ssh_use_stdin.go 

Welcome to the DevNet Always On Sandbox for IOS XE

This is a shared sandbox available for anyone to use to
test APIs, explore features, and test scripts.  Please
keep this in mind as you use it, and respect others use.

The following programmability features are already enabled:
  - NETCONF
  - RESTCONF

Thanks for stopping by.



csr1000v#enable
csr1000v#terminal length 0
csr1000v#config t
Enter configuration commands, one per line.  End with CNTL/Z.
csr1000v(config)#interface loopback 71
csr1000v(config-if)# description golang_script_test
csr1000v(config-if)#interface loopback 72
csr1000v(config-if)# description golang_script_test
csr1000v(config-if)# exit
csr1000v(config)# exit
csr1000v#show ip int brief
Interface              IP-Address      OK? Method Status                Protocol
GigabitEthernet1       10.10.20.48     YES NVRAM  up                    up      
GigabitEthernet2       unassigned      YES NVRAM  up                    up      
GigabitEthernet3       unassigned      YES NVRAM  administratively down down    
Loopback71             unassigned      YES unset  up                    up      
Loopback72             unassigned      YES unset  up                    up      
Loopback100            172.16.100.1    YES other  up                    up      
Loopback103            unassigned      YES unset  up                    up      
csr1000v#show version
Cisco IOS XE Software, Version 16.08.01
Cisco IOS Software [Fuji], Virtual XE Software (X86_64_LINUX_IOSD-UNIVERSALK9-M), Version 16.8.1, RELEASE SOFTWARE (fc3)
Technical Support: http://www.cisco.com/techsupport
Copyright (c) 1986-2018 by Cisco Systems, Inc.
Compiled Tue 27-Mar-18 13:32 by mcpre


Cisco IOS-XE software, Copyright (c) 2005-2018 by cisco Systems, Inc.
All rights reserved.  Certain components of Cisco IOS-XE software are
licensed under the GNU General Public License ("GPL") Version 2.0.  The
software code licensed under GPL Version 2.0 is free software that comes
with ABSOLUTELY NO WARRANTY.  You can redistribute and/or modify such
GPL code under the terms of GPL Version 2.0.  For more details, see the
documentation or "License Notice" file accompanying the IOS-XE software,
or the applicable URL provided on the flyer accompanying the IOS-XE
software.


ROM: IOS-XE ROMMON

csr1000v uptime is 16 hours, 20 minutes
Uptime for this control processor is 16 hours, 21 minutes
System returned to ROM by reload
System image file is "bootflash:packages.conf"
Last reload reason: reload



This product contains cryptographic features and is subject to United
States and local country laws governing import, export, transfer and
use. Delivery of Cisco cryptographic products does not imply
third-party authority to import, export, distribute or use encryption.
Importers, exporters, distributors and users are responsible for
compliance with U.S. and local country laws. By using this product you
agree to comply with applicable laws and regulations. If you are unable
to comply with U.S. and local laws, return this product immediately.

A summary of U.S. laws governing Cisco cryptographic products may be found at:
http://www.cisco.com/wwl/export/crypto/tool/stqrg.html

If you require further assistance please contact us by sending email to
export@cisco.com.

License Level: ax
License Type: Default. No valid license found.
Next reload license Level: ax

cisco CSR1000V (VXE) processor (revision VXE) with 2396264K/3075K bytes of memory.
Processor board ID 9JGOSIUGQVN
3 Gigabit Ethernet interfaces
32768K bytes of non-volatile configuration memory.
16370384K bytes of physical memory.
7774207K bytes of virtual hard disk at bootflash:.
0K bytes of WebUI ODM Files at webui:.

Configuration register is 0x2102

csr1000v#show run all | inc ssh
netconf-yang ssh port 830
ip ssh time-out 120
ip ssh authentication-retries 3
ip ssh window-size 8192
ip ssh rsa keypair-name ssh-key
ip ssh break-string ~break
ip ssh version 2
ip ssh dh min size 2048
no ip ssh rekey time
no ip ssh rekey volume
ip ssh server authenticate user publickey
ip ssh server authenticate user keyboard
ip ssh server authenticate user password
no ip ssh server peruser session limit
ip ssh server certificate profile
ip ssh server algorithm mac hmac-sha2-256 hmac-sha2-512 hmac-sha1 hmac-sha1-96
ip ssh server algorithm encryption aes128-ctr aes192-ctr aes256-ctr
ip ssh server algorithm kex diffie-hellman-group-exchange-sha1 diffie-hellman-group14-sha1
ip ssh server algorithm hostkey x509v3-ssh-rsa ssh-rsa
ip ssh server algorithm authentication publickey keyboard password
ip ssh server algorithm publickey x509v3-ssh-rsa ssh-rsa
ip ssh client algorithm mac hmac-sha2-256 hmac-sha2-512 hmac-sha1 hmac-sha1-96
ip ssh client algorithm encryption aes128-ctr aes192-ctr aes256-ctr
ip ssh client algorithm kex diffie-hellman-group-exchange-sha1 diffie-hellman-group14-sha1
 transport input ssh
no transport type persistent ssh input 
csr1000v#exit

Try another commands file.
pi@raspberrypi:~/Coding_Folder $ cat cmd_file2.txt 
enable
show ip route
show version | inc IOS
exit

pi@raspberrypi:~/Coding_Folder $ 
pi@raspberrypi:~/Coding_Folder $ cat cmd_file2.txt | go run ssh_use_stdin.go 

Welcome to the DevNet Always On Sandbox for IOS XE

This is a shared sandbox available for anyone to use to
test APIs, explore features, and test scripts.  Please
keep this in mind as you use it, and respect others use.

The following programmability features are already enabled:
  - NETCONF
  - RESTCONF

Thanks for stopping by.



csr1000v#enable
csr1000v#show ip route
Codes: L - local, C - connected, S - static, R - RIP, M - mobile, B - BGP
       D - EIGRP, EX - EIGRP external, O - OSPF, IA - OSPF inter area 
       N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2
       E1 - OSPF external type 1, E2 - OSPF external type 2
       i - IS-IS, su - IS-IS summary, L1 - IS-IS level-1, L2 - IS-IS level-2
       ia - IS-IS inter area, * - candidate default, U - per-user static route
       o - ODR, P - periodic downloaded static route, H - NHRP, l - LISP
       a - application route
       + - replicated route, % - next hop override, p - overrides from PfR

Gateway of last resort is 10.10.20.254 to network 0.0.0.0

S*    0.0.0.0/0 [1/0] via 10.10.20.254, GigabitEthernet1
      10.0.0.0/8 is variably subnetted, 2 subnets, 2 masks
C        10.10.20.0/24 is directly connected, GigabitEthernet1
L        10.10.20.48/32 is directly connected, GigabitEthernet1
      172.16.0.0/16 is variably subnetted, 2 subnets, 2 masks
C        172.16.100.0/24 is directly connected, Loopback100
L        172.16.100.1/32 is directly connected, Loopback100
csr1000v#show version | inc IOS
Cisco IOS XE Software, Version 16.08.01
Cisco IOS Software [Fuji], Virtual XE Software (X86_64_LINUX_IOSD-UNIVERSALK9-M), Version 16.8.1, RELEASE SOFTWARE (fc3)
Cisco IOS-XE software, Copyright (c) 2005-2018 by cisco Systems, Inc.
All rights reserved.  Certain components of Cisco IOS-XE software are
documentation or "License Notice" file accompanying the IOS-XE software,
or the applicable URL provided on the flyer accompanying the IOS-XE
ROM: IOS-XE ROMMON
csr1000v#exit
pi@raspberrypi:~/Coding_Folder $


```

