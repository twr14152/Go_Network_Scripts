# ssh_cli_client

The idea behind this code is that you could quickly gather info on the fly as well as make minor changes if so desired. When the code is run, it will ask you how many devices you want to connect to, as well as what commands you want to run.
You will need to add login credential to main.go.


If you have multiple logins you simply run a new instance of RunCli() for each group of devices.

### Update main.go with the groups of devices you want to run the script against.

```
func main() {
    fmt.Println("This is the login to Group1 hosts")
    RunCli("username1", "password1")
    //fmt.Println("This is the login to Group2 hosts")
    //This will connect to another group of hosts using different auth
    //RunCli("username2", "password2")
```
### Running code with show commands:
- For host you need to include ssh port you connecting to. In the host:port format.
```
pi@raspberrypi:~/Code_folder/go_folder/go2run/ssh_cli_client $ go run main.go 
Connecting to IOS-XE hosts:
Number of hosts: 1
Hostname: fastxe:22  

cmds: show ip int brief | inc up


Welcome to the DevNet Sandbox for CSR1000v and IOS XE
 
The following programmability features are already enabled:
  - NETCONF
  - RESTCONF
 
Thanks for stopping by.


csr1000v-1#term len 0
csr1000v-1#show ip int brief | inc up
GigabitEthernet1       10.10.20.48     YES NVRAM  up                    up      
Loopback201            10.99.98.1      YES manual up                    up      
csr1000v-1#
csr1000v-1#exit
Connecting to NX-OS hosts:
Number of hosts: 1       
Hostname: nxos:8181

cmds: show ip int brief | inc up


stty: standard input: Inappropriate ioctl for device
gl_set_term_size: NULL arguments(s).
Lo1                  172.16.0.1      protocol-up/link-up/admin-up       
Lo98                 10.98.98.1      protocol-up/link-up/admin-up       
Lo99                 10.99.99.1      protocol-up/link-up/admin-up       
pi@raspberrypi:~/Code_folder/go_folder/go2run/ssh_cli_client $ 

```



### Running code with configuration commands + validation commands.

```
pi@raspberrypi:~/Code_folder/go_folder/go2run/ssh_cli_client $ go run main.go 
Connecting to IOS-XE hosts:
Number of hosts: 1
Hostname: fastxe:22

cmds: config t, interface loopback 75, description testing_go_script, ip address 75.75.75.75 255.255.255.255, exit, exit, sh interface description, show ip int brief | inc up


Welcome to the DevNet Sandbox for CSR1000v and IOS XE
 
The following programmability features are already enabled:
  - NETCONF
  - RESTCONF
 
Thanks for stopping by.



csr1000v-1#term len 0
csr1000v-1#config t
Enter configuration commands, one per line.  End with CNTL/Z.
csr1000v-1(config)# interface loopback 75
csr1000v-1(config-if)# description testing_go_script
csr1000v-1(config-if)# ip address 75.75.75.75 255.255.255.255
csr1000v-1(config-if)# exit
csr1000v-1(config)# exit
csr1000v-1# sh interface description
Interface                      Status         Protocol Description
Gi1                            up             up       MANAGEMENT INTERFACE - DON'T TOUCH ME
Gi2                            admin down     down     achelove interface
Gi3                            admin down     down     Network Interface
Lo75                           up             up       testing_go_script
Lo201                          up             up       
csr1000v-1# show ip int brief | inc up
GigabitEthernet1       10.10.20.48     YES NVRAM  up                    up      
Loopback75             75.75.75.75     YES manual up                    up      
Loopback201            10.99.98.1      YES manual up                    up      
csr1000v-1#
csr1000v-1#exit
Connecting to NX-OS hosts:
Number of hosts: 1    
Hostname: nxos:8181

cmds: config t, interface loopback 76, ip address 76.76.76.76/32, description testing_go_script, exit, exit, show ip int brief | inc up, show interface description


stty: standard input: Inappropriate ioctl for device
gl_set_term_size: NULL arguments(s).
Lo1                  172.16.0.1      protocol-up/link-up/admin-up       
Lo76                 76.76.76.76     protocol-up/link-up/admin-up       
Lo98                 10.98.98.1      protocol-up/link-up/admin-up       
Lo99                 10.99.99.1      protocol-up/link-up/admin-up       

<cropped for brevity>
-------------------------------------------------------------------------------
Interface                Description                                            
-------------------------------------------------------------------------------
Lo1                      --
Lo30                     My Learning Lab Loopback
Lo76                     testing_go_script
Lo98                     Configured using OpenConfig Model
Lo99                     Full intf config via NETCONF

-------------------------------------------------------------------------------
pi@raspberrypi:~/Code_folder/go_folder/go2run/ssh_cli_client $ 


```

