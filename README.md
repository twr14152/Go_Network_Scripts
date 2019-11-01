# Go_Network_Scripts
This is an attempt to try and learn Golang by creating some useful network scripts.

# SSH CLIENT SCRIPTS
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
Third script involved using a slice to create the command set and then using a for loop to run through the commands.
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

```
The 6th Script uses standard input to put the commands into the script to issue commands on the remote device.
The benefit of doing this is you do not need to overwrite a cmds_file to make changes. You simply create a new file with what ever commands you want and input it into the go script using the format below.
How to use script: go run ssh_use_stdin.go < cmd_file.txt
```
- ssh_client/ssh_use_stdin.go
- ssh_client/cmd_file.txt
- ssh_client/cmd_file2.txt

```
The 7th script uses host_file to login to target hosts and a cmd_file to issue commands on them. At this point this script would be useful for running a common set of commands on multiple devices. The script will simply loop through the hosts and the commands.
```
- ssh_client/ssh_using_host_cmd_files.go
- ssh_client/host_file001.txt
- ssh_client/cmd_file001.txt
- ssh_using_host_cmd_files_output.txt - shows output of script
```

# HTTP CLIENT SCRIPTS
The first script uses the https transport protocol to connect to Cisco Devnet router and GET some information off it. As it stands its more of a learning tool than a useful production script. 
```
- http_client/httpGet_ex1.go
- http_client/httpGet_ex1_output.txt
