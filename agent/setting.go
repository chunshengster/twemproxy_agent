package agent

import (
	"os"

//"encoding/json"
//"launchpad.net/goyaml"
)

//type AgentSeting struct{
//	AgentCfgServerURI string
//	Agent
//}

type CFGGetter interface {
	GetCFG(dst string) TwemproxyCft
}

type TwemproxyCft struct {
	TwCFGGetter  interface{}
	TwCFGFile    string
	TwCFGContent string
}

type TwemproxyInstance struct {
	TwCFG        *TwemproxyCft
	PidFile      string
	TwInitScript string
	TwStatsPort  uint32
}

func checkFile(path string) (bool, err) {
	_, err := os.Stat(path)
	return os.IsExist(err), err
}

//FIXME: use nutcracker -t -f file to check new configuration is OK
func testTWCfg() (bool,err) {
	return true,nil
}

func doReloadTW(cmd string)(bool,err){
	
	return true,nil
}
func (twI *TwemproxyInstance) ReloadTw() (bool, err) {
	isExist, err := checkFile(twI.TwCFG.TwCFGFile)
	if isExist != true {
		//FIXME:do some log and print something
		return isExist, err
	}
	cfgOk,err := testTWCfg()
	if cfgOk !== true{
		return cfgOk,err
	}
	
	reloadRe,err := doReloadTW()
	if reloadRe != true{
		//FIXME: do some log and print something
	}
	
	return true,nil

}

func (twI *TwemproxyInstance) CheckRunning() int {
	return 0
}

func (twI *TwemproxyInstance) ReportUpdateResult() int {
	return 0
}

//func WriteTwemproxyCfg(obj interface{}) {
//	data,err = goyaml.
//}
