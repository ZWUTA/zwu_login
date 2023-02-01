# zwu_login

[English](../README.md) | [简体中文](./README.i18n/ZH_CN.md)

一个用于登录 ZWU(浙江万里学院) 网络服务的命令行工具，对无显示设备尤其有用。使用Golang编写。

## 使用方法

1. 从 Github release 下载 ``zwu_$目标平台_$目标架构``
2. 将其重命名为 ``zwu``
3. (仅 *Unix 系统需要) ``chmod a+x ./zwu``
4. **登录:** ``./zwu -u <username> -p <password>``
5. **登出:** ``./zwu -L``
6. **帮助信息:** ``./zwu -h``

## 作为 Systemd 计时器运行

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
