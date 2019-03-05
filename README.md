# Go_Network_Scripts
This is an attempt to try and learn Golang by creating some useful network scripts. 


First script was to ssh to Cisco Devnet device and print the output.
It works but its not a real useful script.

- ssh_client/can_I_ssh_script.go

Second script had some success getting multiple show commands to run on the device.
I used a map which is analogous to a dictionary in python to create my command set. Then used a for loop to run through those commands. 

However it needs to reconnect after each command. Not a fan of that, especially for config mode. 

- ssh_client/go_ssh_script_2.go

Third script involed using string array to create the command set and then using a for loop to run through the commands.
Nothing earth shattering but it was an opportunity to create another iterable data set.

- ssh_client/go_ssh_script_3.go

SSH script that will allow multi-line commands show or config. Previous one command limit per session resolved by using shell().

 - go_ssh_multi_cmd.go 
