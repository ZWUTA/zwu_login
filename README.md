# zwu_login
A command-line tool to login ZWU(浙江万里学院) net services, especially for headless devices. Written in Golang

# Useage
1. Download ``zwu`` from GitHub release
2. (Linux only) ``chmod a+x ./zwu``
3. Login ``./zwu login YOUR_USERNAME YOUR_PASSWORD``
4. Logout ``./zwu logout``

# Run as systemd-timer
- /etc/systemd/system/zwu.service
````
# /etc/systemd/system/zwu.service
[Unit]
Description=zwu login
[Service]
ExecStart=/root/zwu login YOUR_USERNAME YOUR_PASSWORD
````
- /etc/systemd/system/zwu.timer
````
# /etc/systemd/system/zwu.timer
[Unit]
Description=zwu Timer

[Timer]
OnCalendar=*-*-* *:30
OnBootSec=20sec
Unit=zwu.service

[Install]
WantedBy=timers.target
````
- ``systemdtl enable --now zwu.timer``
