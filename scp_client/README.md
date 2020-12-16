In its current format the way to run this app is to update the filename os.Open("xxx.txt") in main.go 
with the file you want to transfer to the device. Make sure you add the file to the scp_client folder.
Then issue go run main.go or go build then run the app with ./scp_client 

To customize for your needs you would want update host and login info as well as file you're going to transfer.

```
vi main.go

src, err := os.Open("xxx.txt") //.txt,.cfg,,,etc
user := "xxx"
pass := "xxx"
targethost := "hostname_or_ip:port"
```

Local directory setup for demo:

```
pi@raspberrypi:~/Code_folder/go_folder/go2run/scp_client $ ls -l
total 3844
-rw-r--r-- 1 pi pi    5293 Dec  5 12:06 README.md
-rw-r--r-- 1 pi pi     233 Dec  5 10:56 go.mod
-rw-r--r-- 1 pi pi    1284 Dec  5 10:19 go.sum
-rw-r--r-- 1 pi pi    1106 Dec  5 11:48 main.go
-rwxr-xr-x 1 pi pi 3904571 Dec  5 11:39 scp_client
-rw-r--r-- 1 pi pi     129 Dec  5 10:55 testfile001.txt
pi@raspberrypi:~/Code_folder/go_folder/go2run/scp_client $

```




Script results:
```
testfile001.txt does not exist on destination host...

csr1000v-1#dir
Directory of bootflash:/

201601  drwx            24576   Dec 5 2020 16:20:13 +00:00  tracelogs
193537  drwx             4096   Dec 5 2020 10:52:06 +00:00  gs_script
225793  drwx             4096   Dec 5 2020 10:39:02 +00:00  SHARED-IOX
217729  drwx             4096   Dec 5 2020 10:34:42 +00:00  iox_host_data_share
209665  drwx             4096   Dec 5 2020 10:34:42 +00:00  guest-share
80641   drwx             4096   Dec 3 2020 22:18:39 +00:00  .installer
22      -rw-              157   Dec 3 2020 22:17:57 +00:00  csrlxc-cfg.log
137089  drwx             4096   Dec 3 2020 22:17:57 +00:00  license_evlog
19      -rw-             2288   Dec 3 2020 22:17:55 +00:00  cvac.log
18      -rw-               30   Dec 3 2020 22:17:53 +00:00  throughput_monitor_params
15      -rw-             1216   Dec 3 2020 22:16:40 +00:00  mode_event_log
64513   drwx             4096   Sep 1 2020 14:51:38 +00:00  .dbpersist
274177  drwx             4096   Sep 1 2020 14:51:34 +00:00  onep
21      -rw-               16   Sep 1 2020 14:51:32 +00:00  ovf-env.xml.md5
20      -rw-                1   Sep 1 2020 14:51:32 +00:00  .cvac_version
104833  drwx             4096   Sep 1 2020 14:51:29 +00:00  pnp-info
145153  drwx             4096   Sep 1 2020 14:50:48 +00:00  virtual-instance
17      -rwx             1314   Sep 1 2020 14:50:21 +00:00  trustidrootx3_ca.ca
16      -rw-            20109   Sep 1 2020 14:50:21 +00:00  ios_core.p7b
40321   drwx             4096   Sep 1 2020 14:50:16 +00:00  core
169345  drwx             4096   Sep 1 2020 14:50:12 +00:00  bootlog_history
161281  drwx             4096   Sep 1 2020 14:50:07 +00:00  .prst_sync
14      -rw-             1105   Sep 1 2020 14:49:08 +00:00  packages.conf
13      -rw-         48321761   Sep 1 2020 14:49:08 +00:00  csr1000v-rpboot.17.03.01a.SPA.pkg
12      -rw-        470611036   Sep 1 2020 14:49:08 +00:00  csr1000v-mono-universalk9.17.03.01a.SPA.pkg
8065    drwx             4096   Sep 1 2020 14:49:03 +00:00  .rollback_timer
11      drwx            16384   Sep 1 2020 14:48:15 +00:00  lost+found

6286540800 bytes total (4401459200 bytes free)
csr1000v-1#exit
Connection to sandbox-iosxe-latest-1.cisco.com closed by remote host.
Connection to sandbox-iosxe-latest-1.cisco.com closed.

Run script to put testfile001.txt on remote host

pi@raspberrypi:~/Code_folder/go_folder/go2run/scp_client $ ./scp_client
success


pi@raspberrypi:~/Code_folder/go_folder/go2run/scp_client $ ssh developer@sandbox-iosxe-latest-1.cisco.com
Password:

Welcome to the DevNet Sandbox for CSR1000v and IOS XE

The following programmability features are already enabled:
  - NETCONF
  - RESTCONF

Thanks for stopping by.


csr1000v-1#dir
Directory of bootflash:/

201601  drwx            24576   Dec 5 2020 16:23:20 +00:00  tracelogs
23      -rw-              129   Dec 5 2020 16:23:03 +00:00  testfile001.txt  <-- File was copied over
193537  drwx             4096   Dec 5 2020 10:52:06 +00:00  gs_script
225793  drwx             4096   Dec 5 2020 10:39:02 +00:00  SHARED-IOX
217729  drwx             4096   Dec 5 2020 10:34:42 +00:00  iox_host_data_share
209665  drwx             4096   Dec 5 2020 10:34:42 +00:00  guest-share
80641   drwx             4096   Dec 3 2020 22:18:39 +00:00  .installer
22      -rw-              157   Dec 3 2020 22:17:57 +00:00  csrlxc-cfg.log
137089  drwx             4096   Dec 3 2020 22:17:57 +00:00  license_evlog
19      -rw-             2288   Dec 3 2020 22:17:55 +00:00  cvac.log
18      -rw-               30   Dec 3 2020 22:17:53 +00:00  throughput_monitor_params
15      -rw-             1216   Dec 3 2020 22:16:40 +00:00  mode_event_log
64513   drwx             4096   Sep 1 2020 14:51:38 +00:00  .dbpersist
274177  drwx             4096   Sep 1 2020 14:51:34 +00:00  onep
21      -rw-               16   Sep 1 2020 14:51:32 +00:00  ovf-env.xml.md5
20      -rw-                1   Sep 1 2020 14:51:32 +00:00  .cvac_version
104833  drwx             4096   Sep 1 2020 14:51:29 +00:00  pnp-info
145153  drwx             4096   Sep 1 2020 14:50:48 +00:00  virtual-instance
17      -rwx             1314   Sep 1 2020 14:50:21 +00:00  trustidrootx3_ca.ca
16      -rw-            20109   Sep 1 2020 14:50:21 +00:00  ios_core.p7b

csr1000v-1#more testfile001.txt
Testfile001.txt


This is being used to test scp functionality.

If you are reading this on the remote host test was successful.


csr1000v-1#delete testfile001.txt
Delete filename [testfile001.txt]?
Delete bootflash:/testfile001.txt? [confirm]
csr1000v-1#exit
Connection to sandbox-iosxe-latest-1.cisco.com closed by remote host.
Connection to sandbox-iosxe-latest-1.cisco.com closed.
pi@raspberrypi:~/Code_folder/go_folder/go2run/scp_client $

```

