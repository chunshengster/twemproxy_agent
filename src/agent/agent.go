package agent

import (
	//"fmt"
	"io/ioutil"
	"time"
)

type AgentConf struct {
	CheckIntelval    uint32 `yaml:"check_interval"`
	TwemproxyPID     string `yaml:"twemproxy_pid"`
	TwemproxyCFGFile string `yaml:"twemproxy_conf"`
	PolicyUrl        string `yaml:"policy_url"`
	LogFile          string `yaml:""`
}

type AgentInstance struct {
	ConfFile   string
	Conf       AgentConf
	AgentPid   string
	TwInstance TwemproxyInstance
}

func NewAgent(cf string) *AgentInstance {
	return &AgentInstance{
		ConfFile: cf,
	}
}

func (ai *AgentInstance) Run() bool {

	re := ai.ParseConf(ai.ConfFile)
	if re == false {
		return false
	}

	DebugPrint(ai.Conf)
	for {
		ai.AgentCheckRun()
	}
	return true
}

func (ai *AgentInstance) AgentCheckRun() {
	time.Sleep(time.Duration(ai.Conf.CheckIntelval) * time.Second)
	ai.TwInstance.CheckAndRun()
}
func (ai *AgentInstance) ParseConf(pConf string) bool {

	f, err := ioutil.ReadFile(pConf)
	if err != nil {
		DebugPrint("Read file pConf error : " + pConf + err.Error())
		return false
	}

	var acf map[string]AgentConf
	err = ParseYaml(f, &acf)
	if err != nil {
		DebugPrint("Parse Yaml file faild : " + pConf)
		return false
	}
	ai.Conf = acf["agent"]
	if len(ai.Conf.PolicyUrl) == 0 || ai.Conf.CheckIntelval <= 0 || len(ai.Conf.TwemproxyCFGFile) < 2 || len(ai.Conf.TwemproxyPID) < 2 {
		DebugPrint("Error ,not found 'agent' configuration in " + pConf + " file")
		return false
	}

	DebugPrint("Reading and parse conf file " + pConf + " successed ")
	ai.TwInstance.TWCFGPolicyURI = ai.Conf.PolicyUrl
	ai.TwInstance.TwCFGFile = ai.Conf.TwemproxyCFGFile

	return true
}

/*
func (ai *AgentInstance) setTwinstance() bool {
	ai.TwInstance.TWCFGPolicyURI = ai.Conf.PolicyUrl
	ai.TwInstance.TwCFGFile = ai.Conf.TwemproxyCFGFile
	fmt.Println(ai.Conf)
	fmt.Println("ai.TwInstance")
	fmt.Println(ai.TwInstance)
	ai.TwInstance.GetTWCFG()
	return true
}
*/
