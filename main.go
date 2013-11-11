// twemproxy_agent project main.go
package main

import (
	"agent"
)

func main() {
	agent.DebugPrint("hello world")
	//re := agent.Run("./agent.conf")
	//re := agent.ParseYamla("./agent.conf")
	ai := agent.NewAgent("./agent.conf")
	//fmt.Println(ai)
	ai.Run()

}
