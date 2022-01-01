
#### 关于v2ray的订阅配置
```json
{
	
	"path": "outbounds -> settings -> streamSettings -> wsSettings -> path",
	"tls": "outbounds -> settings -> streamSettings -> security",
	"add": "outbounds -> settings -> vnext -> address",
	"port": "outbounds -> settings -> vnext -> port",
	"aid": "outbounds -> settings -> vnext -> users -> alterId",
	"net": "outbounds -> settings -> streamSettings -> network",
	"id": "outbounds -> settings -> vnext -> users -> id",

	"host": "outbounds -> settings -> streamSettings -> tlsSettings -> serverName  and outbounds -> settings -> streamSettings -> wsSettings -> headers -> Host",

	"ps": "备注",
	"v": "2",
	"type": "none",
}
```

#### 关于v2ray的json配置
```json
{
    "log": {
        "access": "",
        "error": "",
        "loglevel": "debug"
    },
    "inbounds": [
        {
            "tag": "transparent",
            "port": 33333,
            "protocol": "dokodemo-door",
            "settings": {
                "network": "tcp,udp",
                "followRedirect": true
            },
            "sniffing": {
                "enabled": true,
                "destOverride": [
                    "http",
                    "tls"
                ]
            },
            "streamSettings": {
                "sockopt": {
                    "tproxy": "tproxy"
                }
            }
        },
        {
            "tag": "test",
            "port": 1080,
            "protocol": "socks",
            "sniffing": {
                "enabled": true,
                "destOverride": [
                    "http",
                    "tls"
                ]
            },
            "settings": {
                "auth": "noauth",
                "udp": true,
                "ip": "127.0.0.1"
            }
        }
    ],
    "outbounds": [
        {
            "tag": "direct",
            "protocol": "freedom"
        },
        {
            "tag": "proxy",
        },
        {
            "tag": "block",
            "protocol": "blackhole",
            "settings": {
                "response": {
                    "type": "http"
                }
            }
        },
        {
            "tag": "dns-out",
            "protocol": "dns"
        }
    ],
    "dns": {
        "servers": [
            {
                "address": "223.5.5.5",
                "port": 53,
                "domains": [
                    "geosite:cn"
                ]
            },
            {
                "address": "114.114.114.114",
                "port": 53,
                "domains": [
                    "geosite:cn"
                ]
            },
            {
                "address": "8.8.8.8",
                "port": 53,
                "domains": [
                    "geosite:geolocation-!cn"
                ]
            },
            {
                "address": "1.1.1.1",
                "port": 53,
                "domains": [
                    "geosite:geolocation-!cn"
                ]
            }
        ]
    },
    "routing": {
        "domainStrategy": "IPIfNonMatch",
        "domainMatcher": "mph",
        "rules": [
            {
                "type": "field",
                "port": 53,
                "network": "udp",
                "outboundTag": "dns-out"
            },
            {
                "type": "field",
                "protocol": [
                    "bittorrent"
                ],
                "outboundTag": "direct"
            },
            {
                "type": "field",
                "outboundTag": "direct",
                "domains": [
                    "agefans"
                ]
            },
            {
                "type": "field",
                "outboundTag": "proxy",
                "domains": [
                    "reimu"
                ]
            },
            {
                "type": "field",
                "outboundTag": "proxy",
                "domains": [
                    "geosite:geolocation-!cn",
                    "geosite:tld-!cn"
                ]
            },
            {
                "type": "field",
                "outboundTag": "proxy",
                "ip": [
                    "geoip:!cn"
                ]
            }
        ]
    }
}                                                                      
```