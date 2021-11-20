package proxy

type Log struct {
	Access   string `json:"access"`
	Error    string `json:"error"`
	Loglevel string `json:"loglevel"`
}

type Inbounds struct {
	Listen         string           `json:"listen"`
	Port           int              `json:"port"`
	Protocol       string           `json:"protocol"`
	Settings       IbSettings       `json:"settings"`
	StreamSettings IbStreamSettings `json:"streamSettings"`
}

type IbSettings struct {
	Network        string `json:"network"`
	FollowRedirect bool   `json:"followRedirect"`
}

type IbStreamSettings struct {
	Sockopt Sockopt `json:"sockopt"`
}

type Sockopt struct {
	Tproxy string `json:"tproxy"`
}

type Outbounds struct {
	Protocol       string           `json:"protocol"`
	Settings       ObSettings       `json:"settings"`
	StreamSettings ObStreamSettings `json:"streamSettings"`
}

type ObSettings struct {
	Vnext []Vnext `json:"vnext"`
}
type Vnext struct {
	Address string  `json:"address"`
	Port    int     `json:"port"`
	Users   []Users `json:"users"`
}
type Users struct {
	Id       string `json:"id"`
	AlterId  int    `json:"alterId"`
	Security string `json:"security"`
}

type ObStreamSettings struct {
	Network    string     `json:"network"`
	Security   string     `json:"security"`
	WsSettings WsSettings `json:"wsSettings"`
}
type WsSettings struct {
	ConnectionReuse bool   `json:"connectionReuse"`
	Path            string `json:"path"`
}

type SimpleV2ray struct {
	Log       Log         `json:"log"`
	Inbounds  []Inbounds  `json:"inbounds"`
	Outbounds []Outbounds `json:"outbounds"`
}
