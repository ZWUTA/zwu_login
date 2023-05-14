# zwu_login

[English](./README.md) | [简体中文](./README.i18n/ZH_CN.md)

A command-line tool to login ZWU(浙江万里学院) net services, especially for headless devices. Written in Golang

## Useage

1. Download ``zwu_$PLATFORM_$ARCH`` from GitHub release
2. Rename it to ``zwu``
3. (*Unix only) ``chmod a+x ./zwu``
4. **Parsing zwu.toml file and login:** ``./zwu [-f]``
5. **Login:** ``./zwu [-f] -u <username> -p <password>``
6. **Logout:** ``./zwu -L``
7. **Status:** ``./zwu -S``
8. **Create zwu.toml template file:** ``./zwu -C``
9. **Parsing zwu.toml and run in daemon mode:** ``./zwu -D``
10. **Help:** ``./zwu -h``

## Template of zwu.toml

**It is better to generate it by ``./zwu -C`` rather then write it by you own!!**
````toml
[config]
Randomize = false
IntervalMinutes = 5
ChangeOnHours = 0
ChangeOnGB = 0.0

[[user]]
Enabled = false
Username = 'username 1'
Password = 'password 1'

[[user]]
Enabled = false
Username = 'username 2'
Password = 'password 2'
````

## Run in Daemon mode

1. ``./zwu -C`` to create zwu.toml template file
2. Modify ``./zwu.toml`` file
3. ``./zwu -D`` to run Daemon mode

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
