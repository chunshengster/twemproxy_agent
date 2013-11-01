package agent

//import (
////"encoding/json"
////"launchpad.net/goyaml"
//)

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
	PidFile      string
	TwInitScript string
	TwStatsPort  uint32
}

func GetRemoteCfgData() int {
	return 0
}

func (twI *TwemproxyInstance) ReloadTw() int {
	return 0

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
