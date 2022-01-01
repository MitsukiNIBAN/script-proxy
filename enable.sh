#!/bin/bash
if [ $(id -u) != 0 ]
then
    exit 0
fi

if [ $(sysctl -n net.ipv4.ip_forward) != 0 ]
then
    exit 0 
fi

#启用转发功能
sysctl -w net.ipv4.ip_forward=1 &>/dev/null
#添加名为233的路由表，路由为将所有流量转发至本地lo
ip route add local default dev lo table 233
#添加路由规则，将所有标记为1的流量应用至233路由表
ip rule add fwmark 1 table 233
ip rule add fwmark 3 table 233

#ipset集合
ipset -X gameset &>/dev/null
ipset create gameset hash:net family inet &>/dev/null
ipset -x udplist &>/dev/null
ipset create udplist hash:net family inet &>/dev/null
#然后添加ip相关的规则域名相关的规则由dnsmasq添加

#SSR流量白名单匹配 这里只有udp tcp流量将走nat表
iptables -t mangle -N PROXY_SSR_RULE
iptables -t mangle -A PROXY_SSR_RULE -j CONNMARK --restore-mark
iptables -t mangle -A PROXY_SSR_RULE -m mark --mark 1/1 -j RETURN
iptables -t mangle -A PROXY_SSR_RULE -p udp -m set --match-set gameset dst -m conntrack --ctstate NEW -j MARK --set-mark 3
iptables -t mangle -A PROXY_SSR_RULE -j CONNMARK --save-mark

#V2RAY流量TCP全进入V2RAY分流,udp流量开启白名单时进行白名单匹配,关闭白名单时全部进入v2ray分流
iptables -t mangle -N PROXY_V2RAY_RULE
iptables -t mangle -A PROXY_V2RAY_RULE -j CONNMARK --restore-mark
iptables -t mangle -A PROXY_V2RAY_RULE -m mark --mark 1/1 -j RETURN
iptables -t mangle -A PROXY_V2RAY_RULE -p tcp -m set --match-set gameset dst -j RETURN
iptables -t mangle -A PROXY_V2RAY_RULE -p tcp -j MARK --set-mark 1
#dns流量送进v2ray
iptables -t mangle -A PROXY_V2RAY_RULE -p udp --dport 53 -m conntrack --ctstate NEW -j MARK --set-mark 1
#关闭udp白名单
# iptables -t mangle -A PROXY_V2RAY_RULE -p udp -m multiport --dports 1:65535 -m conntrack --ctstate NEW -j MARK --set-mark 1
#开启udp白名单
# iptables -t mangle -A PROXY_V2RAY_RULE -p udp -m set --match-set udplist dst -m multiport --dports 1:65535 -m conntrack --ctstate NEW -j MARK --set-mark 3
iptables -t mangle -A PROXY_V2RAY_RULE -j CONNMARK --save-mark

iptables -t mangle -N PROXY_PREROUTING
iptables -t mangle -A PROXY_PREROUTING -d 10.0.0.0/8 -j RETURN
iptables -t mangle -A PROXY_PREROUTING -d 172.16.0.0/12 -j RETURN
iptables -t mangle -A PROXY_PREROUTING -d 192.168.0.0/16 -j RETURN
#这里mark用掩码匹配,1和3都会被命中
iptables -t mangle -A PROXY_PREROUTING -i lo -m mark ! --mark 1/1 -j RETURN
#SSR udp流量处理
iptables -t mangle -A PROXY_PREROUTING -m addrtype ! --src-type LOCAL ! --dst-type LOCAL -p udp -j PROXY_SSR_RULE
iptables -t mangle -A PROXY_PREROUTING -p udp -m mark --mark 3 -j TPROXY --on-port 33334 --tproxy-mark 3
#V2RAY 全部流量处理 放行ssr白名单中的tcp流量
iptables -t mangle -A PROXY_PREROUTING -m addrtype ! --src-type LOCAL ! --dst-type LOCAL -p tcp -j PROXY_V2RAY_RULE
iptables -t mangle -A PROXY_PREROUTING -m addrtype ! --src-type LOCAL ! --dst-type LOCAL -p udp -j PROXY_V2RAY_RULE
iptables -t mangle -A PROXY_PREROUTING -p tcp -m mark --mark 1 -j TPROXY --on-port 33333 --tproxy-mark 1
iptables -t mangle -A PROXY_PREROUTING -p udp -m mark --mark 1 -j TPROXY --on-port 33333 --tproxy-mark 1
#应用链
iptables -t mangle -A PREROUTING -j PROXY_PREROUTING

#nat表中处理白名单中的需要走ssr的tcp流量
iptables -t nat -N PROXY_SSR_RULE
iptables -t nat -A PROXY_SSR_RULE -j CONNMARK --restore-mark
iptables -t nat -A PROXY_SSR_RULE -p tcp -m set --match-set gameset dst -j MARK --set-mark 3
iptables -t nat -A PROXY_SSR_RULE -j CONNMARK --save-mark

iptables -t nat -N PROXY_PREROUTING
iptables -t nat -A PROXY_PREROUTING -i lo -m mark ! --mark 1/1 -j RETURN
iptables -t nat -A PROXY_PREROUTING -m addrtype ! --src-type LOCAL ! --dst-type LOCAL -p tcp -j PROXY_SSR_RULE
iptables -t nat -A PROXY_PREROUTING -p tcp -m mark --mark 3 -j REDIRECT --to-ports 33334
iptables -t nat -A PREROUTING -j PROXY_PREROUTING


# iptables -t mangle -N PROXY_OUTPUT
# # 将mproxy用户的数据直接return掉，因为是output链，所以就直接出去了
# iptables -t mangle -A PROXY_OUTPUT -m owner --uid-owner mproxy -j RETURN
# # 将本机的tcp流量过一遍PROXY_RULE链，需要被转发的流量会被打上标记
# iptables -t mangle -A PROXY_OUTPUT -m addrtype --src-type LOCAL ! --dst-type LOCAL -p tcp -j PROXY_V2RAY_RULE
# # 将本机的udp流量过一遍PROXY_RULE链，需要被转发的流量会被打上标记
# iptables -t mangle -A PROXY_OUTPUT -m addrtype --src-type LOCAL ! --dst-type LOCAL -p udp -j PROXY_V2RAY_RULE
# iptables -t mangle -A OUTPUT -j PROXY_OUTPUT