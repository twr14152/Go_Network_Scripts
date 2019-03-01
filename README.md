# Go_Network_Scripts
This is an attempt to try and learn Golang by creating some useful network scripts. 


First script was to ssh to Cisco Devnet device and print the output.
It works but its not a real useful script.

- ssh_client/can_I_ssh_script.go

Second script had some success getting multiple commands to run on the device.
However it needs to reconnect after each command. Not a fan of that, especially for config mode.

- ssh_client/go_ssh_script_2.go
