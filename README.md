# Vindicta
CCDC Blue Team Tool For Linux

# Run
```
go mod tidy
go run main.go
```

# Structure
```
1. The Main TUI App is TabbedPanels Layout and the library used is https://code.rocketnine.space/tslocum/cview
3. Code for each Tab's UI and function is placed under ui directory
```

# ToDo List:

Work to be done are listed below:

## Todo

For gathering information in real time and rendering it in the app, concurrency might be involved.
Once this phase is solved, we can implement same idea for other items in the Todo list below.
- [ ] Find a way to display ssh logs on real time to the app
- [x] Integrate output of `Fsnotify` golang library to the app for filesystem changes notfications
- [ ] Read Firewall configuration and display it on the app
- [ ] Read Web Server Logs and display it on the app
- [x] Keep a track of network connections with something like `lsof` and display it on the app
- [ ] Gather suspicious processes and monitor their path in real time
- [ ] Detect Scanning and Monitor suspicious IP addresses and Processes
- [ ] List of Important Services and Their Status (Up / Down)
- [ ] Detect user os and switch commands depending on that



## Resource
- https://www.loggly.com/ultimate-guide/linux-logging-basics/
- https://www.digitalocean.com/community/tutorials/how-to-view-and-configure-linux-logs-on-ubuntu-debian-and-centos
- https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/6/html/security_guide/chap-system_auditing
- https://documentation.suse.com/sles/12-SP4/html/SLES-all/cha-audit-comp.html
- https://www.digitalocean.com/community/tutorials/how-to-use-journalctl-to-view-and-manipulate-systemd-logs
- https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/7/html/security_guide/sec-defining_audit_rules_and_controls
- https://geekflare.com/tcpdump-examples/
- https://sysdig.com/blog/file-integrity-monitoring/
- https://www.tutorialspoint.com/linux-process-monitoring#:~:text=In%20Linux%2C%20Top%20command%20is,regularly%20by%20this%20Top%20command.
- https://www.opensourceforu.com/2018/10/how-to-monitor-and-manage-linux-processes/
