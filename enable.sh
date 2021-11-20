#!/bin/bash
if [ $(id -u) != 0 ]
then
    echo 'not root'
    exit 0
fi

if [ $(sysctl -n net.ipv4.ip_forward) != 0 ]
then
    echo 'is enabled'
    exit 0 
fi

#启用转发功能
sysctl -w net.ipv4.ip_forward=1 &>/dev/null
#添加名为233的路由表，路由为将所有流量转发至本地lo
ip route add local default dev lo table 233
#添加路由规则，将所有标记为1的流量应用至233路由表
ip rule add fwmark 1 table 233

#在mangle表中新建V2RAY_PREROUTING链
iptables -t mangle -N V2RAY_PREROUTING
#为V2RAY_PREROUTING链添加常规直连规则，包含一些内网地址和常规地址
iptables -t mangle -A V2RAY_PREROUTING -d 0.0.0.0/8 -j RETURN
iptables -t mangle -A V2RAY_PREROUTING -d 100.64.0.0/10 -j RETURN
iptables -t mangle -A V2RAY_PREROUTING -d 127.0.0.0/8 -j RETURN
iptables -t mangle -A V2RAY_PREROUTING -d 169.254.0.0/16 -j RETURN
iptables -t mangle -A V2RAY_PREROUTING -d 172.16.0.0/12 -j RETURN
iptables -t mangle -A V2RAY_PREROUTING -d 192.0.0.0/24 -j RETURN
iptables -t mangle -A V2RAY_PREROUTING -d 192.0.2.0/24 -j RETURN
iptables -t mangle -A V2RAY_PREROUTING -d 192.88.99.0/24 -j RETURN
iptables -t mangle -A V2RAY_PREROUTING -d 192.168.0.0/16 -j RETURN
iptables -t mangle -A V2RAY_PREROUTING -d 198.18.0.0/15 -j RETURN
iptables -t mangle -A V2RAY_PREROUTING -d 198.51.100.0/24 -j RETURN
iptables -t mangle -A V2RAY_PREROUTING -d 203.0.113.0/24 -j RETURN
iptables -t mangle -A V2RAY_PREROUTING -d 224.0.0.0/4 -j RETURN
iptables -t mangle -A V2RAY_PREROUTING -d 240.0.0.0/4 -j RETURN
iptables -t mangle -A V2RAY_PREROUTING -d 255.255.255.255/32 -j RETURN
#这一行包含自己的内网地址，所以需要额外处理
iptables -t mangle -A V2RAY_PREROUTING -d 10.0.0.0/8 -p tcp -j RETURN
iptables -t mangle -A V2RAY_PREROUTING -d 10.0.0.0/8 -p udp ! --dport 53 -j RETURN
#之后再添加额外的规则转发53端口udp流量以及指定的游戏udp流量
#现在先将所有流量通过tproxy转发至v2ray，打上标记 1 走路由表233
iptables -t mangle -A V2RAY_PREROUTING -p tcp -j TPROXY --on-port 33333 --tproxy-mark 1
iptables -t mangle -A V2RAY_PREROUTING -p udp -j TPROXY --on-port 33333 --tproxy-mark 1
#将V2RAY_PREROUTING应用到mangle表的PREROUTING链
iptables -t mangle -A PREROUTING -j V2RAY_PREROUTING

#为了处理udp流量，不考虑做本机代理了，然后output链就不用处理了
#在mangle表中新建V2RAY_OUTPUT链
#iptables -t mangle -N V2RAY_OUTPUT
#为V2RAY_OUTPUT链添加常规直连规则，包含一些内网地址和常规地址
#iptables -t mangle -A V2RAY_OUTPUT -d 0.0.0.0/8 -j RETURN
#iptables -t mangle -A V2RAY_OUTPUT -d 10.0.0.0/8 -j RETURN
#iptables -t mangle -A V2RAY_OUTPUT -d 100.64.0.0/10 -j RETURN
#iptables -t mangle -A V2RAY_OUTPUT -d 169.254.0.0/16 -j RETURN
#iptables -t mangle -A V2RAY_OUTPUT -d 172.16.0.0/12 -j RETURN
#iptables -t mangle -A V2RAY_OUTPUT -d 192.0.0.0/24 -j RETURN
#iptables -t mangle -A V2RAY_OUTPUT -d 192.0.2.0/24 -j RETURN
#iptables -t mangle -A V2RAY_OUTPUT -d 192.88.99.0/24 -j RETURN
#iptables -t mangle -A V2RAY_OUTPUT -d 192.168.0.0/16 -j RETURN
#iptables -t mangle -A V2RAY_OUTPUT -d 198.18.0.0/15 -j RETURN
#iptables -t mangle -A V2RAY_OUTPUT -d 198.51.100.0/24 -j RETURN
#iptables -t mangle -A V2RAY_OUTPUT -d 203.0.113.0/24 -j RETURN
#iptables -t mangle -A V2RAY_OUTPUT -d 224.0.0.0/4 -j RETURN
#iptables -t mangle -A V2RAY_OUTPUT -d 240.0.0.0/4 -j RETURN
#iptables -t mangle -A V2RAY_OUTPUT -d 255.255.255.255/32 -j RETURN
#作用与V2RAY_PREROUTING链中的相同
#iptables -t mangle -A V2RAY_OUTPUT -p udp --dport 53 -j RETURN
#匹配到 MARK 2 的流量直接放出去
#iptables -t mangle -A V2RAY_OUTPUT -j RETURN -m mark --mark 2
#剩余的流量 标记上 mark 1
#iptables -t mangle -A V2RAY_OUTPUT -p tcp -j MARK --set-mark 1
#iptables -t mangle -A V2RAY_OUTPUT -p udp -j MARK --set-mark 1
#将V2RAY_OUTPUT应用到mangle表的OUTPUT链
#iptables -t mangle -A OUTPUT -j V2RAY_OUTPUT

echo 'enable'