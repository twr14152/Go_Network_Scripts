### GoEAPI
I found that the goeapi documentation was rather sparse. It was decent for show commands but configuration documentation was lacking. After playing around with it and looking at go.dev package documentation I got a feel for what it was capable of. I did sort out how to use it for configuration. I definately feel like this framework could use some attention. Doesn't appear to be actively updated.

Between pyeapi and goeapi I would probably stick with pyeapi. It is useable though.



- Sample config change
```
/*
twr14152@DESKTOP-S55FNN9:~/code_folder/go_folder/misc/goeapi/change_config$ go run main.go
Connected to ceos1



Pre-Change state:


[map[command:show ip interface brief result:Address
Interface       IP Address         Status      Protocol          MTU    Owner
--------------- ------------------ ----------- ------------- ---------- -------
Loopback0       1.1.1.1/32         up          up              65535
Management0     172.17.0.2/16      up          up               1500]] <nil>
true



Post change state:


[map[command:show ip interface brief result:Address
Interface       IP Address         Status      Protocol          MTU    Owner
--------------- ------------------ ----------- ------------- ---------- -------
Loopback0       1.1.1.1/32         up          up              65535
Loopback55      55.55.55.55/32     up          up              65535
Management0     172.17.0.2/16      up          up               1500]] <nil>


-------------------------------

Connected to ceos2



Pre-Change state:


[map[command:show ip interface brief result:Address
Interface       IP Address         Status      Protocol          MTU    Owner
--------------- ------------------ ----------- ------------- ---------- -------
Loopback0       2.2.2.2/32         up          up              65535
Loopback69      unassigned         up          up              65535
Management0     172.17.0.3/16      up          up               1500]] <nil>
true



Post change state:


[map[command:show ip interface brief result:Address
Interface       IP Address         Status      Protocol          MTU    Owner
--------------- ------------------ ----------- ------------- ---------- -------
Loopback0       2.2.2.2/32         up          up              65535
Loopback55      55.55.55.55/32     up          up              65535
Loopback69      unassigned         up          up              65535
Management0     172.17.0.3/16      up          up               1500]] <nil>


-------------------------------

Connected to ceos3



Pre-Change state:


[map[command:show ip interface brief result:Address
Interface       IP Address         Status      Protocol          MTU    Owner
--------------- ------------------ ----------- ------------- ---------- -------
Loopback0       3.3.3.3/32         up          up              65535
Management0     172.17.0.4/16      up          up               1500]] <nil>
true



Post change state:


[map[command:show ip interface brief result:Address
Interface       IP Address         Status      Protocol          MTU    Owner
--------------- ------------------ ----------- ------------- ---------- -------
Loopback0       3.3.3.3/32         up          up              65535
Loopback55      55.55.55.55/32     up          up              65535
Management0     172.17.0.4/16      up          up               1500]] <nil>


-------------------------------

Connected to ceos4



Pre-Change state:


[map[command:show ip interface brief result:Address
Interface       IP Address         Status      Protocol          MTU    Owner
--------------- ------------------ ----------- ------------- ---------- -------
Loopback0       4.4.4.4/32         up          up              65535
Management0     172.17.0.5/16      up          up               1500]] <nil>
true



Post change state:


[map[command:show ip interface brief result:Address
Interface       IP Address         Status      Protocol          MTU    Owner
--------------- ------------------ ----------- ------------- ---------- -------
Loopback0       4.4.4.4/32         up          up              65535
Loopback55      55.55.55.55/32     up          up              65535
Management0     172.17.0.5/16      up          up               1500]] <nil>


-------------------------------

Connected to ceos5



Pre-Change state:


[map[command:show ip interface brief result:Address
Interface       IP Address         Status      Protocol          MTU    Owner
--------------- ------------------ ----------- ------------- ---------- -------
Loopback0       5.5.5.5/32         up          up              65535
Management0     172.17.0.6/16      up          up               1500]] <nil>
true



Post change state:


[map[command:show ip interface brief result:Address
Interface       IP Address         Status      Protocol          MTU    Owner
--------------- ------------------ ----------- ------------- ---------- -------
Loopback0       5.5.5.5/32         up          up              65535
Loopback55      55.55.55.55/32     up          up              65535
Management0     172.17.0.6/16      up          up               1500]] <nil>


-------------------------------

Connected to ceos6



Pre-Change state:


[map[command:show ip interface brief result:Address
Interface       IP Address         Status      Protocol          MTU    Owner
--------------- ------------------ ----------- ------------- ---------- -------
Loopback0       6.6.6.6/32         up          up              65535
Management0     172.17.0.7/16      up          up               1500]] <nil>
true



Post change state:


[map[command:show ip interface brief result:Address
Interface       IP Address         Status      Protocol          MTU    Owner
--------------- ------------------ ----------- ------------- ---------- -------
Loopback0       6.6.6.6/32         up          up              65535
Loopback55      55.55.55.55/32     up          up              65535
Management0     172.17.0.7/16      up          up               1500]] <nil>


-------------------------------

twr14152@DESKTOP-S55FNN9:~/code_folder/go_folder/misc/goeapi/change_config$
*/
```

- sample of the show commands
```
twr14152@DESKTOP-S55FNN9:~/code_folder/go_folder/misc/goeapi/show_commands$ go run main.go show version, show ip interface brief
[show version  show ip interface brief]
Connected to ceos1
[map[command:show version result:Arista cEOSLab
Hardware version:
Serial number:
Hardware MAC address: 0242.ac20.996c
System MAC address: 0242.ac20.996c

Software image version: 4.30.5M-35156751.4305M (engineering build)
Architecture: x86_64
Internal build version: 4.30.5M-35156751.4305M
Internal build ID: 29383dc6-2c4f-445b-8162-1209cd1b57df
Image format version: 1.0
Image optimization: None

cEOS tools version: (unknown)
Kernel version: 5.15.146.1-microsoft-standard-WSL2

Uptime: 4 days, 4 hours and 56 minutes
Total memory: 8108944 kB
Free memory: 916656 kB] map[command: show ip interface brief result:Address
Interface       IP Address         Status      Protocol          MTU    Owner
--------------- ------------------ ----------- ------------- ---------- -------
Loopback0       1.1.1.1/32         up          up              65535
Loopback55      55.55.55.55/32     up          up              65535
Management0     172.17.0.2/16      up          up               1500]] <nil>




Connected to ceos2
[map[command:show version result:Arista cEOSLab
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

Uptime: 4 days, 4 hours and 56 minutes
Total memory: 8108944 kB
Free memory: 916656 kB] map[command: show ip interface brief result:Address
Interface       IP Address         Status      Protocol          MTU    Owner
--------------- ------------------ ----------- ------------- ---------- -------
Loopback0       2.2.2.2/32         up          up              65535
Loopback55      55.55.55.55/32     up          up              65535
Loopback69      unassigned         up          up              65535
Management0     172.17.0.3/16      up          up               1500]] <nil>




Connected to ceos3
[map[command:show version result:Arista cEOSLab
Hardware version:
Serial number:
Hardware MAC address: 0242.ac53.f242
System MAC address: 0242.ac53.f242

Software image version: 4.30.5M-35156751.4305M (engineering build)
Architecture: x86_64
Internal build version: 4.30.5M-35156751.4305M
Internal build ID: 29383dc6-2c4f-445b-8162-1209cd1b57df
Image format version: 1.0
Image optimization: None

cEOS tools version: (unknown)
Kernel version: 5.15.146.1-microsoft-standard-WSL2

Uptime: 4 days, 4 hours and 56 minutes
Total memory: 8108944 kB
Free memory: 916656 kB] map[command: show ip interface brief result:Address
Interface       IP Address         Status      Protocol          MTU    Owner
--------------- ------------------ ----------- ------------- ---------- -------
Loopback0       3.3.3.3/32         up          up              65535
Loopback55      55.55.55.55/32     up          up              65535
Management0     172.17.0.4/16      up          up               1500]] <nil>




Connected to ceos4
[map[command:show version result:Arista cEOSLab
Hardware version:
Serial number:
Hardware MAC address: 0242.ac6b.0386
System MAC address: 0242.ac6b.0386

Software image version: 4.30.5M-35156751.4305M (engineering build)
Architecture: x86_64
Internal build version: 4.30.5M-35156751.4305M
Internal build ID: 29383dc6-2c4f-445b-8162-1209cd1b57df
Image format version: 1.0
Image optimization: None

cEOS tools version: (unknown)
Kernel version: 5.15.146.1-microsoft-standard-WSL2

Uptime: 4 days, 4 hours and 56 minutes
Total memory: 8108944 kB
Free memory: 916656 kB] map[command: show ip interface brief result:Address
Interface       IP Address         Status      Protocol          MTU    Owner
--------------- ------------------ ----------- ------------- ---------- -------
Loopback0       4.4.4.4/32         up          up              65535
Loopback55      55.55.55.55/32     up          up              65535
Management0     172.17.0.5/16      up          up               1500]] <nil>




Connected to ceos5
[map[command:show version result:Arista cEOSLab
Hardware version:
Serial number:
Hardware MAC address: 0242.ac92.2c1a
System MAC address: 0242.ac92.2c1a

Software image version: 4.30.5M-35156751.4305M (engineering build)
Architecture: x86_64
Internal build version: 4.30.5M-35156751.4305M
Internal build ID: 29383dc6-2c4f-445b-8162-1209cd1b57df
Image format version: 1.0
Image optimization: None

cEOS tools version: (unknown)
Kernel version: 5.15.146.1-microsoft-standard-WSL2

Uptime: 1 day, 16 hours and 25 minutes
Total memory: 8108944 kB
Free memory: 916560 kB] map[command: show ip interface brief result:Address
Interface       IP Address         Status      Protocol          MTU    Owner
--------------- ------------------ ----------- ------------- ---------- -------
Loopback0       5.5.5.5/32         up          up              65535
Loopback55      55.55.55.55/32     up          up              65535
Management0     172.17.0.6/16      up          up               1500]] <nil>




Connected to ceos6
[map[command:show version result:Arista cEOSLab
Hardware version:
Serial number:
Hardware MAC address: 0242.ac07.09bd
System MAC address: 0242.ac07.09bd

Software image version: 4.30.5M-35156751.4305M (engineering build)
Architecture: x86_64
Internal build version: 4.30.5M-35156751.4305M
Internal build ID: 29383dc6-2c4f-445b-8162-1209cd1b57df
Image format version: 1.0
Image optimization: None

cEOS tools version: (unknown)
Kernel version: 5.15.146.1-microsoft-standard-WSL2

Uptime: 1 day, 16 hours and 25 minutes
Total memory: 8108944 kB
Free memory: 916560 kB] map[command: show ip interface brief result:Address
Interface       IP Address         Status      Protocol          MTU    Owner
--------------- ------------------ ----------- ------------- ---------- -------
Loopback0       6.6.6.6/32         up          up              65535
Loopback55      55.55.55.55/32     up          up              65535
Management0     172.17.0.7/16      up          up               1500]] <nil>




twr14152@DESKTOP-S55FNN9:~/code_folder/go_folder/misc/goeapi/show_commands$
```
