# go-proxy
用go搭建的简易代理管理后端


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
    "loglevel": "warning"                                                         
  },                                                                              
  "inbounds": [                                                                   
    {                                                                             
      "listen": "0.0.0.0",                                                        
      "port":60080,                                                               
      "protocol": "dokodemo-door",                                                
      "settings": {                                                               
          "network": "tcp,udp",                                                   
          "followRedirect": true                                                  
        },                                                                        
      "streamSettings": {                                                         
          "sockopt": {                                                            
            "tproxy": "redirect"                                                  
          }                                                                       
        }                                                                         
    }                                                                             
],                                                                                
  "outbounds": [                                                                  
    {                                                                             
      "protocol": "",                                                        
      "settings": {                                                               
        "vnext": [                                                                
          {                                                                       
            "address": "",                               
            "port": 0,                                                        
            "users": [                                                            
              {                                                                   
                "id": "",                     
                "alterId": -1,                                                     
                "security": "auto"                                                
              }                                                                   
            ]                                                                     
          }                                                                       
        ]                                                                         
      },                                                                          
      "streamSettings": {                                                         
        "network": "",                                                          
        "security": "",                                                        
        "wsSettings": {                                                           
          "connectionReuse": true,                                                
          "path": ""                                                    
        }                                                                         
      }                                                                           
    }                                                                             
  ]                                                                               
}                                                                                 
```