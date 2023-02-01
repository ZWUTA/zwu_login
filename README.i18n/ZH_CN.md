# zwu_login

[English](../README.md) | [简体中文](./README.i18n/ZH_CN.md)

zwu_login 是一款适用于登录ZWU（浙江万里学院）网络服务的命令行工具，特别适合在没有显示设备的情况下使用。该工具是用Golang编写的。

## 使用方法

1. 从 Github releases 下载符合 ``zwu_<目标平台>_<目标架构>`` 的文件。
2. 将其重命名为 ``zwu``
3. (仅适用于Unix系统) 运行``chmod a+x ./zwu`` 以授予执行权限
4. **登录:** ``./zwu -u <用户名> -p <密码>``
5. **登出:** ``./zwu -L``
6. **获取帮助信息:** ``./zwu -h``

## 使用 Systemd 定时器运行

创建文件 /etc/systemd/system/zwu.service

````text
# /etc/systemd/system/zwu.service
[Unit]
Description=zwu_login
After=network.target

[Service]
ExecStart=/root/zwu -u <username> -p <password>
````

创建文件 /etc/systemd/system/zwu.timer

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

启动并立即运行 ``systemctl enable --now zwu.timer``
