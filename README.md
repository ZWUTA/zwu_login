# zwu_login
A command-line tool to login ZWU(浙江万里学院) net services, especially for headless devices. Written in Golang

# Useage
1. Download ``zwu_$ARCH_$PLATFORM`` from GitHub release
2. Rename it to ``zwu``
3. (Linux only) ``chmod a+x ./zwu``
4. **Login:** ``./zwu -u <username> -p <password>``
5. **Logout:** ``./zwu -L``
6. **Help:** ``./zwu -h``

# Run as systemd-timer
- /etc/systemd/system/zwu.service
````
# /etc/systemd/system/zwu.service
[Unit]
Description=zwu_login
After=network.target

[Service]
ExecStart=/root/zwu -u <username> -p <password>
````
- /etc/systemd/system/zwu.timer
````
# /etc/systemd/system/zwu.timer
[Unit]
Description=zwu Timer

[Timer]
OnCalendar=hourly
OnBootSec=10sec
Unit=zwu.service

[Install]
WantedBy=timers.target
````
- ``systemdtl enable --now zwu.timer``
