package proxy

import (
	"encoding/json"
	"io/ioutil"
	"mmm3w/go-proxy/subscribe"
	"mmm3w/go-proxy/support"
	"os"
	"path"
)

func startUpV2ray() error {
	_, err := support.ExecCommand("nohup", "v2ray", "--config="+support.V2rayConfigJsonPath(), "</dev/null", "$>>/dev/null", "2>&1", "&")

	return err
}

func startUpSsr() error {
	return nil
}

func applySsrConfig(key string) error {
	return nil
}

func applyV2rayConfig(key string) error {
	cacheFolder := support.JsonCacheFolder()
	configContent, err := ioutil.ReadFile(path.Join(cacheFolder, key+".json"))
	if err != nil {
		return err
	}
	var v2ray subscribe.V2ray
	err = json.Unmarshal([]byte(configContent), &v2ray)
	if err != nil {
		return err
	}

	routingNode := make(map[string]interface{})
	routingNode["domainStrategy"] = "IPIfNonMatch"
	routingNode["domainMatcher"] = "mph"
	routingNode["rules"] = []map[string]interface{}{v2rayRoutingDirectBTNode(),
		v2rayRoutingDirectSiteNode(),
		v2rayRoutingDirectIpNode(),
		v2rayRoutingProxySiteNode(),
		v2rayRoutingProxyIpNode()}

	v2rayConfig := make(map[string]interface{})
	v2rayConfig["log"] = v2rayLogNode()
	v2rayConfig["inbounds"] = []map[string]interface{}{v2rayInboundsNode()}
	v2rayConfig["outbounds"] = []map[string]interface{}{v2rayOutboundsDirectNode(), v2rayOutboundsProxyNode(v2ray)} //注意顺序
	v2rayConfig["routing"] = []map[string]interface{}{}

	jsonData, _ := json.Marshal(v2rayConfig)
	err = ioutil.WriteFile(support.V2rayConfigJsonPath(), jsonData, 0644)
	if err != nil {
		return err
	}

	newConfigInfo := make(map[string]string)
	newConfigInfo[support.KeyCurrentV2rayConfig] = key
	currentFolder, _ := os.Getwd()
	err = support.AppendConfig(path.Join(currentFolder, support.ConfigFileName), newConfigInfo)

	return err
}

func v2rayLogNode() map[string]interface{} {
	log := make(map[string]interface{})
	log["access"] = support.AccessLogFile()
	log["error"] = support.ErrorLogFile()
	log["loglevel"] = support.V2rayLogLevel()
	return log
}

func v2rayInboundsNode() map[string]interface{} {
	ibSettings := make(map[string]interface{})
	ibSettings["network"] = "tcp,udp"
	ibSettings["followRedirect"] = true

	ibSniffing := make(map[string]interface{})
	ibSniffing["enabled"] = true
	ibSniffing["destOverride"] = []string{"http", "tls"}

	sockopt := make(map[string]string)
	sockopt["tproxy"] = "tproxy"

	ibStreamSettings := make(map[string]interface{})
	ibStreamSettings["sockopt"] = sockopt

	inbounds := make(map[string]interface{})
	inbounds["port"] = 333333
	inbounds["protocol"] = "dokodemo-door"
	inbounds["settings"] = ibSettings
	inbounds["sniffing"] = ibSniffing
	inbounds["streamSettings"] = ibStreamSettings

	return inbounds
}

func v2rayOutboundsDirectNode() map[string]interface{} {
	node := make(map[string]interface{})
	node["tag"] = "direct"
	node["protocol"] = "freedom"
	return node
}

func v2rayOutboundsProxyNode(v2ray subscribe.V2ray) map[string]interface{} {
	users := make(map[string]interface{})
	users["id"] = v2ray.Id
	users["alterId"] = v2ray.Aid
	users["security"] = "auto"

	vnext := make(map[string]interface{})
	vnext["address"] = v2ray.Add
	vnext["port"] = v2ray.Port
	vnext["users"] = users

	wsSettings := make(map[string]interface{})
	wsSettings["connectionReuse"] = true
	wsSettings["path"] = v2ray.Path

	obStreamSettings := make(map[string]interface{})
	obStreamSettings["network"] = v2ray.Net
	obStreamSettings["security"] = v2ray.Tls
	obStreamSettings["wsSettings"] = wsSettings

	node := make(map[string]interface{})
	node["tag"] = "proxy"
	node["protocol"] = "vmess"
	node["settings"] = []map[string]interface{}{vnext}
	node["streamSettings"] = obStreamSettings
	return node
}

func v2rayRoutingProxyIpNode() map[string]interface{} {
	node := make(map[string]interface{})
	node["type"] = "field"
	node["outboundTag"] = "proxy"
	node["ip"] = []string{"geoip:!cn"} //后续自定义的在这里添加
	return node
}

func v2rayRoutingProxySiteNode() map[string]interface{} {
	node := make(map[string]interface{})
	node["type"] = "field"
	node["outboundTag"] = "proxy"
	node["domains"] = []string{"geosite:geolocation-!cn", "geosite:tld-!cn"} //后续自定义的在这里添加
	return node
}

func v2rayRoutingDirectBTNode() map[string]interface{} {
	node := make(map[string]interface{})
	node["type"] = "field"
	node["outboundTag"] = "direct"
	node["protocol"] = []string{"bittorrent"}
	return node
}

func v2rayRoutingDirectSiteNode() map[string]interface{} {
	node := make(map[string]interface{})
	node["type"] = "field"
	node["outboundTag"] = "direct"
	node["domains"] = []string{"geosite:cn"} //后续自定义的在这里添加，先添加一个防止空规则
	return node
}

func v2rayRoutingDirectIpNode() map[string]interface{} {
	node := make(map[string]interface{})
	node["type"] = "field"
	node["outboundTag"] = "direct"
	node["ip"] = []string{"geoip:cn"} //后续自定义的在这里添加，先添加一个防止空规则
	return node
}
