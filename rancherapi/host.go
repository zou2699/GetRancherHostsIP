package rancherapi

import (
	"encoding/json"
	"github.com/zou2699/rancherApi/defs"
	"io/ioutil"
	"net/http"
)

// 返回当前环境的IP列表
func GetIps(rancherUrl, accessKey, secretkey string) (ips defs.Ips) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", rancherUrl, nil)
	if err != nil {
		panic(err)
	}

	// 添加认证头信息
	request.SetBasicAuth(accessKey, secretkey)

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	bodyInfo, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	// 使用map解析json,
	m := make(map[string]interface{})

	err = json.Unmarshal([]byte(bodyInfo), &m)
	if err != nil {
		panic(err)
	}

	//fmt.Println(string(bodyInfo))
	//获取json的data数据,存入hosts
	hosts := m["data"].([]interface{})

	var ip string
	//遍历hosts列表中的数据
	for _, host := range hosts {
		//fmt.Println(host)
		//再次使用map解析json中的ip
		ip = host.(map[string]interface{})["agentIpAddress"].(string)
		ips = append(ips, ip)
	}
	return ips
}
