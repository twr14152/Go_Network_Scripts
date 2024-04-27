package main

import (
        "fmt"
        "github.com/aristanetworks/goeapi"
)

//hosts to configure
var HostList = []string{"ceos1", "ceos2", "ceos3", "ceos4", "ceos5", "ceos6"}

func main() {

        //cmds := []string{"show running-config"}
        cmds1 := []string{"show ip interface brief"}

        // Configuration commands
        cfg1 := "interface loopback 55"
        cfg2 := "description testing goeapi"
        cfg3 := "ip address 55.55.55.55/32"
        //cfg4 := "no interface loopback 55"

        for _, host := range HostList {
                node, err := goeapi.ConnectTo(host)
                if err != nil {
                        fmt.Println(err)
                }
                fmt.Printf("Connected to %v\n", host)
                fmt.Printf("\n\n\nPre-Change state:\n\n\n")
                fmt.Println(node.Enable(cmds1))
                fmt.Println(node.Config(cfg1, cfg2, cfg3))
                fmt.Printf("\n\n\nPost change state: \n\n\n")
                fmt.Println(node.Enable(cmds1))
                fmt.Printf("\n\n-------------------------------\n\n")
        }
}
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
