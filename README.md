# go-proxy 翻新中
用go搭建的简易代理管理后端


#### 说明
`enable.sh`,`disable.sh` 基础规则，变动的机会不大




#### 使用
1.`go build`

2.`cp -af go-proxy /etc/go-proxy/main`

3.`chmod +x /etc/go-proxy/main`

4.`install proxygo /usr/local/bin`

5.`install -m 644 proxygo.service /etc/systemd/system`

6.`systemctl start proxygo.service`


配合[go-proxy-cline](https://github.com/MitsukiNIBAN/go-proxy-client)在目录`/etc/go-proxy`下部署静态界面