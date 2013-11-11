package agent

import (
	"encoding/json"
	"os"
	"os/exec"
	//"time"
)

type HTTPCFGGetter struct {
	//Url string
}

func (HG *HTTPCFGGetter) GetCFG(url string) []byte {

	r, err := HttpGet(url)

	if err != nil {
		//FIXME: do some log
		DebugPrint(err.Error())
		return nil
	} else {
		if len(r) <= 1 {
			//FIXME: do some log
			DebugPrint("The response from " + url + " is too short")
			return nil
		}
		DebugPrint(string(r))
		return r
	}
}

type TWPool struct {
	name               string
	Listen             string   `yaml:"listen" json:"listen"`
	Hash               string   `yaml:"hash" json:"hash"`
	HashTag            string   `yaml:"hash_tag" json:"hash_tag"`
	Distribution       string   `yaml:"distribution" json:"distribution"`
	TimeOut            int      `yaml:"timeout" json:"timeout"`
	Backlog            int      `yaml:"backlog" json:"backlog"`
	ClientConnections  int      `yaml:"" json:""`
	Redis              bool     `yaml:"redis" json:"redis"`
	Preconnect         int      `yaml:"preconnect" json:"preconnect"`
	AutoEjectHosts     bool     `yaml:"auto_eject_hosts" json:"auto_eject_hosts"`
	ServerConnections  int      `yaml:"server_connections" json:"server_connections"`
	ServerRetryTimeout int      `yaml:"server_retry_timeout" json:"server_retry_timeout"`
	ServerFailureLimit int      `yaml:"server_failure_limit" json:"server_failure_limit"`
	Servers            []string `yaml:"servers" json:"servers"`
}

type TwemproxyInstance struct {
	TWCFGGetter    HTTPCFGGetter
	TWCFGPolicyURI string
	TwCFG          map[string]TWPool
	TwCFGFile      string
	PidFile        string
	//Daemonize	   bool
	TwInitScript string
	TwStatsPort  uint
}

// Get twemproxy config from remote URI,json format
//
func (twi *TwemproxyInstance) GetTWCFG() bool {
	//FIXME: 从 agent.conf 中获取策略服务器URI，得到json文件，然后进行更新
	if len(twi.TWCFGPolicyURI) <= 0 {
		DebugPrint("twi.TWCFGPolicyURL lenth is 0,error return")
		return false
	}

	r := twi.TWCFGGetter.GetCFG(twi.TWCFGPolicyURI)
	if r != nil {
		var tmpCFG map[string]TWPool
		err := json.Unmarshal(r, &tmpCFG)
		if err != nil {
			//FIXME: do some log
			DebugPrint(err.Error())
			return false
		}
		twi.TwCFG = tmpCFG
		DebugPrint("............................................")
		DebugPrint(twi.TwCFG)
		DebugPrint(twi.TwCFGFile)
		DebugPrint("............................................")
		//return true

		isExist := CheckFileExist(twi.TwCFGFile)
		if isExist != true {
			DebugPrint(twi.TwCFGFile + " is not exist")
			re, err := EncodeYaml(twi.TwCFG)
			if err != nil {
				DebugPrint(err.Error())
			}
			//DebugPrint(re)
			ok := WriteConf(re, twi.TwCFGFile)
			DebugPrint("Write Conf")
			DebugPrint(ok)
			if ok != true {
				return false
			} else {
				return true
			}
			DebugPrint(re)
		} else {
			//FIXME: 如果配置文件已经存在，则需要先进行配置文件内容对比，如果文件内容一致，则跳过；否则，用新的配置对旧配置进行替换，并reload twemproxy
			// 此处逻辑需要清除，对比操作在 WriteConf 方法中进行实现
			isSame := FileCompare(twi.TwCFGFile, "abcddd")
			DebugPrint("Need file compare and reload twemproxy")
		}
	}
	return false
}

//use nutcracker -t -f file to check new configuration is OK
func (twi *TwemproxyInstance) testTWCfg(cfgPath string) bool {
	isExist := CheckFileExist(cfgPath)
	if isExist != true {
		//FIXME:do some log and print something
		return false
	}

	//FIXME:do nutcracker -t -f TwCFGFile,and return the error
	return true
}

// check if nutcracker process is running
func (twi *TwemproxyInstance) checkTWRunning() bool {
	//FIXME: add more code

	re, err := ExecCommand(twi.TwInitScript, "status")
	if err != nil {
		// FIXME: do panic and log
		return false
	} else {
		DebugPrint(re)
	}
	return false
}

func (twi *TwemproxyInstance) doReloadTW(cmdString string) (bool, error) {

	_, err := os.Stat(cmdString)
	if err != nil {
		//FIXME:do some logs
		return false, err
	}

	var reloadArg = "start"
	if twi.checkTWRunning() {
		reloadArg = "reload"
	}

	cmd := exec.Command(cmdString, reloadArg)
	if err != nil {
		//FIXME:do some logs
		return false, err
	}

	err = cmd.Start()
	if err == nil {
		//FIXME:do some logs
		cmd.Wait()
		return true, err
	} else {
		cmd.Wait()
		return false, err
	}
	return false, nil
}

func (twi *TwemproxyInstance) ReloadTw() bool {

	cfgOk := twi.testTWCfg(twi.TwCFGFile)
	if cfgOk != true {
		return false
	}

	reloadRe, _ := twi.doReloadTW(twi.TwInitScript)
	if reloadRe != true {
		//FIXME: do some log and print something
	}

	return true

}

func (twi *TwemproxyInstance) CheckAndRun() bool {
	r := twi.GetTWCFG()
	if r != true {

		//FIXME: do some log
		//return false
	}
	return true
}

func (twi *TwemproxyInstance) ReportUpdateResult() int {
	return 0
}
