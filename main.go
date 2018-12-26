//GOARCH=amd64 GOOS=linux go build -o rancherHosts .
package main

import (
	"fmt"
	"github.com/zou2699/rancherApi/pkgs/conf"
	"github.com/zou2699/rancherApi/rancherapi"
	"io"
	"os"
)

func main() {
	// 获取配置信息
	envs := conf.Getconf()

	// 将信息写入content，在写入到文件中
	var content string
	for _, env := range envs.Envs {
		fmt.Println(env.Env)
		// 调用方法获取ip
		ips := rancherapi.GetIps(env.HostsURL, env.AccessKey, env.Secretkey)

		fmt.Println(ips)

		content += "[" + env.Env + "]" + "\n"
		for _, ip := range ips {
			content += ip + "\n"
		}
		content += "\n"
	}

	f, err := os.OpenFile("./rancherHost", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	content += `[all:vars]
ansible_ssh_user=root
ansible_ssh_pass=passw0rd
ansible_python_interpreter=/usr/bin/python3
`

	_, err = io.WriteString(f, content)
	if err != nil {
		panic(err)
	}

}
