# zwu_login

[English](./README.md) | [简体中文](./README.i18n/ZH_CN.md)

A command-line tool to login ZWU(浙江万里学院) net services, especially for headless devices. Written in Golang

## Useage

1. Download ``zwu_$PLATFORM_$ARCH`` from GitHub release
2. Rename it to ``zwu``
3. (*Unix only) ``chmod a+x ./zwu``
4. **Login:** ``./zwu -u <username> -p <password>``
5. **Logout:** ``./zwu -L``
6. **Status:** ``./zwu -S``
7. **Help:** ``./zwu -h``

## Run as Systemd-timer

Create file /etc/systemd/system/zwu.service

````text
# /etc/systemd/system/zwu.service
[Unit]
Description=zwu_login
After=network.target

[Service]
ExecStart=/root/zwu -u <username> -p <password>
````

Create file /etc/systemd/system/zwu.timer

````text
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

Create file ``systemdtl enable --now zwu.timer``
