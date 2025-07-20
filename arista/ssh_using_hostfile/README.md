# Script Details
- This was a major pain in the ass to get working.
- This script will allow you configure multiple devices using a host file and host.cfg files
- First you need to create your host file in this example we call it group1.txt
- Just add your "host:port" you are connecting on. You need the port for the script to work 
```
toddriemenschneider@clab:~/clabs/labs/ceos_labs/lab3/scripts/go_folder/run_script$ cat group1.txt 
clab-lab3-ceos1:22
clab-lab3-ceos2:22

```
- Then create your config/command file detailing what you want to do
- The naming format is important it should be file_ + host:22 + .cfg
- This format will allows the host and the command file to be paired when the script runs
```
toddriemenschneider@clab:~/clabs/labs/ceos_labs/lab3/scripts/go_folder/run_script$ cat file_clab-lab3-ceos1\:22.cfg 
enable
config t
interface loopback 69
description test_script
ip address 69.69.69.69/32
end
show hostname
show ip interface brief
show run interface loopback 69
exit
toddriemenschneider@clab:~/clabs/labs/ceos_labs/lab3/scripts/go_folder/run_script$ cat file_clab-lab3-ceos2\:22.cfg 
enable 
config 
interface loopback69
description test_script
ip address 69.69.69.69/32
end
show hostname
show ip int brief
show run int loopback69
exit
```
- Those details just discussed get added to main()
- Run the script
```
toddriemenschneider@clab:~/clabs/labs/ceos_labs/lab3/scripts/go_folder/run_script$ go run main.go

--- Connecting to clab-lab3-ceos1:22 ---
[Waiting for prompt: >]
[PROMPT FOUND]
[Waiting for prompt: #]
[PROMPT FOUND]
[Waiting for prompt: #]
[PROMPT FOUND]
Sending commands from: file_clab-lab3-ceos1:22.cfg
[SENDING] enable
[OUTPUT]
enable
ceos1#
[SENDING] config t
[OUTPUT]
config t
ceos1(config)#
[SENDING] interface loopback 69
[OUTPUT]
interface loopback 69
ceos1(config-if-Lo69)#
[SENDING] description test_script
[OUTPUT]
description test_script
ceos1(config-if-Lo69)#
[SENDING] ip address 69.69.69.69/32
[OUTPUT]
ip address 69.69.69.69/32
ceos1(config-if-Lo69)#
[SENDING] end
[OUTPUT]
end
ceos1#
[SENDING] show hostname
[OUTPUT]
show hostname
Hostname: ceos1
FQDN:     ceos1
ceos1#
[SENDING] show ip interface brief
[OUTPUT]
show ip interface brief
                                       
Interface     IP Address       Status  
------------- ---------------- --------
Ethernet1     192.168.0.0/31   up      
Ethernet2     192.168.0.2/31   up      
Loopback0     1.1.1.1/32       up      
Loopback64    unassigned       up      
Loopback69    69.69.69.69/32   up      
Management0   172.20.20.3/24   up      

                                Address
Interface    Protocol      MTU  Owner  
------------ ---------- ------- -------
Ethernet1    up           1500         
Ethernet2    up           1500         
Loopback0    up          65535         
Loopback64   up          65535         
Loopback69   up          65535         
Management0  up           1500         

ceos1#
[SENDING] show run interface loopback 69
[OUTPUT]
show run interface loopback 69
interface Loopback69
   description test_script
   ip address 69.69.69.69/32
ceos1#
[SENDING] exit
[OUTPUT]
exit


--- Connecting to clab-lab3-ceos2:22 ---
[Waiting for prompt: >]
[PROMPT FOUND]
[Waiting for prompt: #]
[PROMPT FOUND]
[Waiting for prompt: #]
[PROMPT FOUND]
Sending commands from: file_clab-lab3-ceos2:22.cfg
[SENDING] enable 
[OUTPUT]
enable 
ceos2#
[SENDING] config 
[OUTPUT]
config 
ceos2(config)#
[SENDING] interface loopback69
[OUTPUT]
interface loopback69
ceos2(config-if-Lo69)#
[SENDING] description test_script
[OUTPUT]
description test_script
ceos2(config-if-Lo69)#
[SENDING] ip address 69.69.69.69/32
[OUTPUT]
ip address 69.69.69.69/32
ceos2(config-if-Lo69)#
[SENDING] end
[OUTPUT]
end
ceos2#
[SENDING] show hostname
[OUTPUT]
show hostname
Hostname: ceos2
FQDN:     ceos2
ceos2#
[SENDING] show ip int brief
[OUTPUT]
show ip int brief
                                       
Interface     IP Address       Status  
------------- ---------------- --------
Ethernet1     unassigned       up      
Ethernet2     unassigned       up      
Loopback69    69.69.69.69/32   up      
Management0   172.20.20.4/24   up      

                                Address
Interface    Protocol      MTU  Owner  
------------ ---------- ------- -------
Ethernet1    up           1500         
Ethernet2    up           1500         
Loopback69   up          65535         
Management0  up           1500         

ceos2#
[SENDING] show run int loopback69
[OUTPUT]
show run int loopback69
interface Loopback69
   description test_script
   ip address 69.69.69.69/32
ceos2#
[SENDING] exit
[OUTPUT]
exit
```
