package agent

import (
	"errors"
	"github.com/mreiferson/go-httpclient"
	"io/ioutil"
	"launchpad.net/goyaml"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func ParseYaml(i []byte, o interface{}) error {
	DebugPrint(string(i))
	err := goyaml.Unmarshal(i, o)
	if err != nil {
		DebugPrint(err.Error())
		return err
	}
	return nil
}

func EncodeYaml(i interface{}) ([]byte, error) {
	return goyaml.Marshal(i)
}

func WriteTMPConf(b []byte) (interface{}, error) {
	f, err := ioutil.TempFile("/tmp", "twemproxy_agent_")
	if err != nil {
		//FIXME:do some log
		return nil, err
	}
	defer f.Close()
	length, err := f.Write(b)
	if err != nil && length > 0 {
		return nil, err
	}

	return f.Name(), err
}

func WriteConf(b []byte, p string) bool {

	fn, err := WriteTMPConf(b)
	if err != nil {
		//FIXME:do some log
		DebugPrint(err)
		return false
	}

	switch fn.(type) {
	case string:
		DebugPrint("................Write 	Conf")
		err = MoveFile(fn.(string), p)
		if err != nil {
			DebugPrint(err)
			return false
		}
		return true
	default:
		return false
	}

	return false
}

func MoveFile(s, d string) error {
	DebugPrint(s, d)
	if CheckFileExist(s) {
		DebugPrint(s + " exist!")
		err := os.Rename(s, d)
		if err != nil {
			c, err := ioutil.ReadFile(s)
			if err != nil {
				DebugPrint(err)
				return err
			}
			err = ioutil.WriteFile(d, c, 0644)
			if err != nil {
				DebugPrint(err)
				return err
			}
			err = os.Remove(s)
			if err != nil {
				DebugPrint(err)
				// Do not return false,but Warning
			}
			return nil
		} else {
			DebugPrint(err.Error())
		}
	} else {
		return errors.New("MoveFile faild,source file does not exits")
	}
	return errors.New("Move faild")
}

func RenameFile(s, d string) error {
	err := os.Rename(s, d)
	if err != nil {
		return err
	}
	return nil
}

func HttpGet(url string) ([]byte, error) {

	if len(url) <= 0 {
		return nil, errors.New("url param can not be null")
	}
	transport := &httpclient.Transport{
		ConnectTimeout:        1 * time.Second,
		ResponseHeaderTimeout: 2 * time.Second,
		RequestTimeout:        3 * time.Second,
	}
	defer transport.Close()

	c := &http.Client{Transport: transport}
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := c.Do(req)

	if err != nil {
		//FIXME: do some log
		DebugPrint(err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	robots, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//FIXME: do some log
		DebugPrint(err.Error())
	}
	return robots, nil
}
func CheckFileExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	} else if err != nil && os.IsNotExist(err) {
		return false
	} else {
		//FIXME: do some log
		return false
	}

}
func ExecCommand(cmdString string, arg ...string) (out []byte, err error) {
	if len(cmdString) <= 0 {
		return nil, errors.New("cmdString param is nil")
	}
	_, err = exec.LookPath(cmdString)
	if err != nil {
		return nil, err
	}
	cmd := exec.Command(cmdString, arg...)
	out, err = cmd.CombinedOutput()

	if err != nil {
		return nil, err
	}
	return out, nil
}

func FileCompire(f1, f2 string) bool {
	return false
}

func GenerateFileHash() {
	return
}

func DebugPrint(i ...interface{}) {
	L := log.New(os.Stderr, "Twemproxy_agent ", log.Ldate|log.Lmicroseconds|log.Llongfile)
	L.Println(i...)
}
