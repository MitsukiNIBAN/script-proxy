#!/bin/bash
#关闭转发功能
sysctl -w net.ipv4.ip_forward=0 &>/dev/null
#删除路由规则
ip route del local default dev lo table 233 &>/dev/null
ip rule del fwmark 1 table 233 &>/dev/null
ip rule del fwmark 3 table 233 &>/dev/null
#删除对应链
iptables -t mangle -D PREROUTING -j PROXY_PREROUTING &>/dev/null
iptables -t mangle -D OUTPUT -j PROXY_OUTPUT &>/dev/null
iptables -t nat -D PREROUTING -j PROXY_PREROUTING &>/dev/null
#清空相应链中内容
iptables -t mangle -F PROXY_PREROUTING  &>/dev/null
iptables -t mangle -X PROXY_PREROUTING  &>/dev/null
iptables -t mangle -F PROXY_OUTPUT  &>/dev/null
iptables -t mangle -X PROXY_OUTPUT  &>/dev/null
iptables -t mangle -F PROXY_SSR_RULE  &>/dev/null
iptables -t mangle -X PROXY_SSR_RULE  &>/dev/null
iptables -t mangle -F PROXY_V2RAY_RULE  &>/dev/null
iptables -t mangle -X PROXY_V2RAY_RULE  &>/dev/null
iptables -t nat -F PROXY_PREROUTING  &>/dev/null
iptables -t nat -X PROXY_PREROUTING  &>/dev/null
iptables -t nat -F PROXY_SSR_RULE  &>/dev/null
iptables -t nat -X PROXY_SSR_RULE  &>/dev/null
#清除ipset合集
ipset -X gameset &>/dev/null
ipset -X udplist &>/dev/null