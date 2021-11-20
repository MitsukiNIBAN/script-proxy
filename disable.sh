#!/bin/bash
#关闭转发功能
sysctl -w net.ipv4.ip_forward=0 &>/dev/null
#删除路由规则
ip route del local default dev lo table 233 &>/dev/null
ip rule del fwmark 1 table 233 &>/dev/null
#删除对应链
iptables -t mangle -D PREROUTING -j V2RAY_PREROUTING &>/dev/null
iptables -t mangle -D OUTPUT -j V2RAY_OUTPUT &>/dev/null
#清空相应链中内容
iptables -t mangle -F V2RAY_PREROUTING  &>/dev/null
iptables -t mangle -X V2RAY_PREROUTING  &>/dev/null
iptables -t mangle -F V2RAY_OUTPUT  &>/dev/null
iptables -t mangle -X V2RAY_OUTPUT  &>/dev/null