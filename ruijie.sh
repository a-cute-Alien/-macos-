#!/bin/bash

#mac 管理员密码
net=""
#需要替换mac地址的网卡
passwd=""
#要替换的mac地址
mac=""
#锐捷登陆用户名
username=""
#锐捷登陆密码
userpass=""
echo $passwd | sudo -S ifconfig $net down
echo $passwd | sudo -S ifconfig $net up
echo $passwd | sudo -S ifconfig $net ether $mac
echo $passwd | sudo -S /Users/wuzichun/Desktop/mentohust-macos/mentohust -u$username -p$userpass -n$net -dhcp-type0