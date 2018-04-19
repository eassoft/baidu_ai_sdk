package config

import (
	//"encoding/json"
	//"os"
	//"encoding/json"
	"hm/eclibs/config"

	//"hm/eclibs/formatcom"
)

type CompanyConfig struct {
	Address  string `json:"address"`
	TelPhone string `json:"telphone"`
	EMail    string `json:"email"`
}

var (
	CompanyConf   CompanyConfig
	ListenAddr    = ":88"
	TLSListenAddr = ":443"
	//ImageServer   = "127.0.0.1"
	//CarImageServer string
	B_App_Id     = "10975184"
	B_Api_Key    = "auXGm4G1vpR6RfYHtL3zXnLu"
	B_Secret_Key = "mSzPFYNqnzqaZzgQFF9bLVAscKqWuwP7"

	Dev = "false" //调试true 发布 false
)

func init() {

	myConfig := new(config.Config)
	myConfig.InitConfig("conf/app.conf")

	ListenAddr = myConfig.Read("default", "listenaddr")
	TLSListenAddr = myConfig.Read("default", "tlslistenaddr")
	//ImageServer = myConfig.Read("default", "imageserver")
	B_App_Id = myConfig.Read("baiduai", "B_App_Id")
	B_Api_Key = myConfig.Read("baiduai", "B_Api_Key")
	B_Secret_Key = myConfig.Read("baiduai", "B_Secret_Key")
	Dev = myConfig.Read("default", "dev")

}
