Uses may vary. Good for lab environment.
```
twr14152@DESKTOP-S55FNN9:~/code_folder/go_folder/misc/go_net_scripts/ssh_client$ go run main.go
Number of hosts: 2

Enter host and ssh port being used. (Ex. x.x.x.x:22 or hostname:22)


Enter host: ceos1:22

cmds:  enable, config, interface loopback 9, description testing_script, sh run int loopback 9, show ip int brief, no interface loopback 9, show ip int brief

Enter host: ceos2:22

cmds: enable, show version, show ip arp, show ip int brief

++++++++++++++++++++++++++++++++
Connected to:  ceos1:22
++++++++++++++++++++++++++++++++
Pagination disabled.
interface Loopback9
   description testing_script
                                                                        Address
Interface       IP Address         Status      Protocol          MTU    Owner
--------------- ------------------ ----------- ------------- ---------- -------
Loopback0       1.1.1.1/32         up          up              65535
Loopback9       unassigned         up          up              65535
Management0     172.17.0.2/16      up          up               1500

                                                                        Address
Interface       IP Address         Status      Protocol          MTU    Owner
--------------- ------------------ ----------- ------------- ---------- -------
Loopback0       1.1.1.1/32         up          up              65535
Management0     172.17.0.2/16      up          up               1500

++++++++++++++++++++++++++++++++
Connected to:  ceos2:22
++++++++++++++++++++++++++++++++
Pagination disabled.
Arista cEOSLab
Hardware version:
Serial number:
Hardware MAC address: 0242.ac97.38c5
System MAC address: 0242.ac97.38c5

Software image version: 4.30.5M-35156751.4305M (engineering build)
Architecture: x86_64
Internal build version: 4.30.5M-35156751.4305M
Internal build ID: 29383dc6-2c4f-445b-8162-1209cd1b57df
Image format version: 1.0
Image optimization: None

cEOS tools version: (unknown)
Kernel version: 5.15.146.1-microsoft-standard-WSL2

Uptime: 2 days, 7 hours and 37 minutes
Total memory: 8108944 kB
Free memory: 1780132 kB

Address         Age (sec)  Hardware Addr   Interface
172.17.0.1        0:00:00  0242.da6b.0d0e  Management0
                                                                        Address
Interface       IP Address         Status      Protocol          MTU    Owner
--------------- ------------------ ----------- ------------- ---------- -------
Loopback0       2.2.2.2/32         up          up              65535
Loopback69      unassigned         up          up              65535
Management0     172.17.0.3/16      up          up               1500

twr14152@DESKTOP-S55FNN9:~/code_folder/go_folder/misc/go_net_scripts/ssh_client$

```
