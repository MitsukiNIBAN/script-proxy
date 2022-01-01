package proxy

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"mmm3w/go-proxy/subscribe"
	"mmm3w/go-proxy/support"
	"os"
	"os/exec"
	"strings"
)

func getV2rayPid() (string, error) {
	return support.ExecCommand("pidof", "v2ray")
}

func getSSRPid() (string, error) {
	return support.ExecCommand("pidof", "ssr-redir")
}

func getScriptInfo() (string, error) {
	data, err := support.ExecCommand("sysctl", "-n", "net.ipv4.ip_forward")
	if err != nil {
		return "", err
	}

	data = strings.TrimSpace(data)

	if data == "1" {
		return "enable", nil
	} else {
		return "", nil
	}
}

func forwardSwitch(enable bool) error {
	var tag string
	if enable {
		tag = "1"
	} else {
		tag = "0"
	}
	_, err := support.ExecCommand("sysctl", "-w", "net.ipv4.ip_forward="+tag)
	return err
}

func startUpV2ray() error {
	vcf := support.V2rayConfigFile()
	if !support.Exists(vcf) {
		return errors.New("未找到V2ray配置文件")
	}
	cmd := exec.Command("v2ray", "--config="+vcf, "</dev/null", "$>/dev/null", "2>&1", "&")
	err := cmd.Start()
	return err
}

func startUpSsr() error {
	scf := support.SsrConfigFile()
	if !support.Exists(scf) {
		return errors.New("未找到SSR配置文件")
	}
	cmd := exec.Command("ssr-redir", "-c", scf, "-u", "</dev/null", ">/dev/null", "&")
	err := cmd.Start()
	return err
}

func startScript() error {
	ssf := support.StartScriptFile()
	if !support.Exists(ssf) {
		return errors.New("未找到启动脚本")
	}
	_, err := support.ExecCommand("bash", ssf)
	return err
}

func stopScript() error {
	ssf := support.StopScriptFile()
	if !support.Exists(ssf) {
		return errors.New("未找到停用脚本")
	}
	_, err := support.ExecCommand("bash", ssf)
	return err
}

func killProcess(pid string) error {
	pidData := strings.Split(pid, ",")
	for _, item := range pidData {
		if len(item) > 0 {
			_, err := support.ExecCommand("kill", "-9", item)
			if err != nil {
				return err
			}
		}

	}
	return nil
}
func applySsrConfig(cpath string) error {
	support.PutC("ssr_config_file", cpath)
	return nil
}

func applyV2rayConfig(cpath string) error {

	vcf := support.V2rayConfigFile()
	if len(vcf) <= 0 {
		return errors.New("未配置V2ray配置文件路径")
	}

	fiii := strings.LastIndex(vcf, "/")
	p := vcf[:fiii]
	if !support.Exists(p) {
		err := os.MkdirAll(p, os.ModePerm)
		if err != nil {
			return err
		}
	}

	configContent, err := ioutil.ReadFile(cpath)
	if err != nil {
		return err
	}
	var v2ray subscribe.V2ray
	err = json.Unmarshal([]byte(configContent), &v2ray)
	if err != nil {
		return err
	}

	v2rayConfig := make(map[string]interface{})
	v2rayConfig["log"] = v2rayLogNode()
	v2rayConfig["inbounds"] = []map[string]interface{}{v2rayInboundsNode()}
	v2rayConfig["outbounds"] = []map[string]interface{}{
		v2rayOutboundsDirectNode(),
		v2rayOutboundsProxyNode(v2ray),
		v2rayOutboundsBlockNode(),
		v2rayOutboundsDNSNode()}
	v2rayConfig["dns"] = v2rayDNSNode()
	v2rayConfig["routing"] = v2rayRoutingNode()

	jsonData, _ := json.Marshal(v2rayConfig)
	return ioutil.WriteFile(vcf, jsonData, 0644)
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
	inbounds["tag"] = "transparent"
	inbounds["port"] = 33333
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

	vnextItem := make(map[string]interface{})
	vnextItem["address"] = v2ray.Add
	vnextItem["port"] = v2ray.Port
	vnextItem["users"] = []map[string]interface{}{users}

	vnext := make(map[string]interface{})
	vnext["vnext"] = []map[string]interface{}{vnextItem}

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
	node["settings"] = vnext
	node["streamSettings"] = obStreamSettings
	return node
}

func v2rayOutboundsBlockNode() map[string]interface{} {
	response := make(map[string]interface{})
	response["type"] = "http"

	settings := make(map[string]interface{})
	settings["response"] = response

	node := make(map[string]interface{})
	node["tag"] = "block"
	node["protocol"] = "blackhole"
	node["settings"] = settings
	return node
}

func v2rayOutboundsDNSNode() map[string]interface{} {
	node := make(map[string]interface{})
	node["tag"] = "dns-out"
	node["protocol"] = "dns"
	return node
}

func v2rayDNSNode() map[string]interface{} {
	server := make(map[string]interface{})
	server["address"] = "114.114.114.114"
	server["port"] = 53
	server["domains"] = []string{"geosite:cn"}

	var servers []interface{} = make([]interface{}, 2)
	servers[0] = server
	servers[1] = "8.8.8.8"

	node := make(map[string]interface{})
	node["servers"] = servers
	return node
}

func v2rayRoutingNode() map[string]interface{} {
	node := make(map[string]interface{})
	node["domainStrategy"] = "IPIfNonMatch"
	node["domainMatcher"] = "mph"
	node["rules"] = []map[string]interface{}{
		v2rayRoutingDNSNode(),
		v2rayRoutingDirectBTNode(),
		v2rayRoutingDirectSiteNode(),
		v2rayRoutingProxySiteNode(),
		v2rayRoutingProxyIpNode()}
	return node
}

func v2rayRoutingDNSNode() map[string]interface{} {
	node := make(map[string]interface{})
	node["type"] = "field"
	node["port"] = 53
	node["network"] = "udp"
	node["outboundTag"] = "dns-out"
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
	node["domains"] = []string{"agefans"} //后续再改，先写死一个防空
	return node
}

func v2rayRoutingProxySiteNode() map[string]interface{} {
	node := make(map[string]interface{})
	node["type"] = "field"
	node["outboundTag"] = "proxy"
	node["domains"] = []string{"geosite:geolocation-!cn", "geosite:tld-!cn"}
	return node
}

func v2rayRoutingProxyIpNode() map[string]interface{} {
	node := make(map[string]interface{})
	node["type"] = "field"
	node["outboundTag"] = "proxy"
	node["ip"] = []string{"geoip:!cn"} //后续自定义的在这里添加
	return node
}
