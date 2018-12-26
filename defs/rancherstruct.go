package defs

type RancherEnv struct {
	Envs []struct {
		Env       string `yaml:"env"`
		HostsURL  string `yaml:"hostsUrl"`
		AccessKey string `yaml:"accessKey"`
		Secretkey string `yaml:"secretkey"`
	} `json:"envs"`
}

type Ips []string
