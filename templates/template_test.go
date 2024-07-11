package templates

import (
	"fmt"
	"os"
	"testing"

	"github.com/chainreactors/logs"
	"github.com/chainreactors/neutron/common"
	"github.com/chainreactors/neutron/protocols"
	"gopkg.in/yaml.v3"
)

var ExecuterOptions *protocols.ExecuterOptions

func TestTemplate_Compile(t1 *testing.T) {
	content, _ := os.ReadFile("tmp.yaml")
	t := &Template{}
	err := yaml.Unmarshal(content, t)
	if err != nil {
		println(err.Error())
		return
	}
	if t != nil {
		err := t.Compile(ExecuterOptions)
		if err != nil {
			println(err.Error())
			return
		}
	}
	println("success")
}

func TestTemplate_Execute(t1 *testing.T) {
	common.NeutronLog.SetLevel(logs.Debug)
	content, _ := os.ReadFile("tmp.yaml")
	t := &Template{}
	err := yaml.Unmarshal(content, t)
	if err != nil {
		println(err.Error())
		return
	}
	if t != nil {
		err := t.Compile(ExecuterOptions)
		if err != nil {
			println(err.Error())
			return
		}
	}

	println("load success")

	res, err := t.Execute("http://81.70.40.138:8000", nil)
	if err == nil {
		fmt.Println("ok", res)
	} else {
		fmt.Println(res)
	}
}
