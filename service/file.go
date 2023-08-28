package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Result struct {
	TotalCount int `json:"totalCount"`
	Blocks     []struct {
		Cid          string `json:"cid"`
		Height       int    `json:"height"`
		Timestamp    int    `json:"timestamp"`
		Size         int    `json:"size"`
		WinCount     int    `json:"winCount"`
		Reward       string `json:"reward"`
		Penalty      string `json:"penalty"`
		MessageCount int    `json:"messageCount"`
	} `json:"blocks"`
}

func Getdata(addr string) {
	urladd := "https://filfox.info/api/v1/address/"
	url := fmt.Sprintf("%s/%s/blocks?pageSize=20&page=0", urladd, addr)

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("HTTP GET请求失败:", err)
		return
	}
	defer response.Body.Close()

	// 检查响应状态码
	if response.StatusCode != http.StatusOK {
		fmt.Println("HTTP请求返回状态码不是200 OK:", response.Status)
		return
	}

	// 读取响应主体
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("读取响应主体失败:", err)
		return
	}
	var result Result
	// 将响应主体内容打印到控制台

	err = json.Unmarshal(body, &result)
	if err != nil {
		return
	}
	for k, v := range result.Blocks {
		fmt.Println(k, " : ", v.Height)
		fmt.Println(k, " : ", v.Cid)

	}

}
