package agent

import (
	//"log"
	"fmt"
	"io/ioutil"
	"os"
)

func WriteTMPConf(b []byte) (string, error) {
	f, err := ioutil.TempFile("/tmp", "twemproxy_agent")
	if err != nil {
		//FIXME:do some log
		return "", err
	}
	defer f.Close()
	length, err := f.Write(b)
	if err != nil {
		return "", err
	}
	
	return f.Name(), err
}

func WriteConf(b []byte, p string) bool {
	if fn,err := WriteTMPConf(b); err == nil; fn != ""{
		err = os.Rename(fn,p)
		if err != nil{
			//FIXME:do some log
			return false
		}
		return true
	}
	return false
}

func GenFileHash() (string,error) {
	return 
}

func Debug() bool {
	return true
}
