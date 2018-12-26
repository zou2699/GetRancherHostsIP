package conf

import (
	"github.com/zou2699/rancherApi/defs"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func Getconf() (envs defs.RancherEnv) {
	yamlfile, err := ioutil.ReadFile("./conf/conf.yaml")
	if err != nil {
		panic(err)
	}
	var conf defs.RancherEnv
	err = yaml.Unmarshal(yamlfile, &conf)
	if err != nil {
		panic(err)
	}

	return conf
}
