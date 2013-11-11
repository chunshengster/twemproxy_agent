package agent

import (
	"encoding/json"
	"launchpad.net/goyaml"
	"os"
	"os/exec"
)

type CFGGetter interface {
	GetCFG(interface{}) []byte
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
	TWCFGGetter  CFGGetter
	TwCFG        map[string]TWPool
	TwCFGFile    string
	PidFile      string
	TwInitScript string
	TwStatsPort  uint
}

func (Twi *TwemproxyInstance) checkFileExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return true, nil
}

//use nutcracker -t -f file to check new configuration is OK
func testTWCfg(cfgPath string) bool {
	isExist, _ := checkFileExist(cfgPath)
	if isExist != true {
		//FIXME:do some log and print something
		return false
	}

	//FIXME:do nutcracker -t -f TwCFGFile,and return the error
	return true
}

// check if nutcracker process is running
func (TWI *TwemproxyInstance) checkTWRunning() bool {
	//FIXME: add more code
	return true
}

func doReloadTW(cmdString string) (bool, error) {

	_, err := os.Stat(cmdString)
	if err != nil {
		//FIXME:do some logs
		return false, err
	}

	var reloadArg = "start"
	if checkTWRunning() {
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

func (TWI *TwemproxyInstance) ReloadTw() bool {

	cfgOk := testTWCfg(TWI.TwCFG.TwCFGFile)
	if cfgOk != true {
		return false
	}

	reloadRe, _ := doReloadTW(TWI.TwInitScript)
	if reloadRe != true {
		//FIXME: do some log and print something
	}

	return true

}

func (TWI *TwemproxyInstance) CheckRunning() int {
	return 0
}

func (TWI *TwemproxyInstance) ReportUpdateResult() int {
	return 0
}
