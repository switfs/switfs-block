package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
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
	url := fmt.Sprintf("%s/%s/blocks?pageSize=100&page=0", urladd, addr)

	response, err := http.Get(url)
	if err != nil {
		log.Errorf("HTTP GET请求失败:%v", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Errorf("关闭响应体失败:%v", err)
		}
	}(response.Body)

	// 检查响应状态码
	if response.StatusCode != http.StatusOK {
		log.Errorf("HTTP请求返回状态码不是200 OK:%v", response.Status)
		return
	}

	// 读取响应主体
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Error("读取响应主体失败:", err)
		return
	}
	var result Result
	// 将响应主体内容打印到控制台

	err = json.Unmarshal(body, &result)
	if err != nil {
		return
	}
	for _, v := range result.Blocks {
		height := v.Height - 1
		Reward := NanoOrAttoToFIL(v.Reward, AttoFIL)

		err = MinerUP(v.Cid, IntToString(height), addr, Reward)
		if err != nil {
			log.Error(err.Error())
			return
		}
	}

}

func IntToString(e int) string {
	return strconv.Itoa(e)
}

func GetPostData(addr string) error {
	url := "https://api.filutils.com/api/v2/block"
	//{"miner":"f02942808","height":0,"pageIndex":1,"pageSize":20}
	params := request{
		Miner:     addr,
		Height:    0,
		PageIndex: 1,
		PageSize:  10,
	}

	jsonParams, err := json.Marshal(params)
	if err != nil {
		log.Error("json序列化失败:", err)
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonParams))
	if err != nil {
		log.Error("创建请求失败:", err)
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		log.Error("请求失败:", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	var result data
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	for _, v := range result.Data {
		height := v.Height - 1
		//Reward := NanoOrAttoToFIL(v.Reward, AttoFIL)
		err = MinerUP(v.Cid, IntToString(height), addr, v.ExactReward)
		if err != nil {
			log.Error(err.Error())
			return err
		}
	}

	return nil
}

type request struct {
	Miner     string `json:"miner"`
	Height    int    `json:"height"`
	PageIndex int    `json:"pageIndex"`
	PageSize  int    `json:"pageSize"`
}

type data struct {
	Code      int    `json:"code"`
	Total     int    `json:"total"`
	PageIndex int    `json:"pageIndex"`
	PageSize  int    `json:"pageSize"`
	Message   string `json:"Message"`
	Data      []struct {
		Height       int    `json:"height"`
		Cid          string `json:"cid"`
		MineTime     string `json:"mineTime"`
		MessageCount int    `json:"messageCount"`
		Size         int    `json:"size"`
		Miner        string `json:"miner"`
		MinerTag     string `json:"minerTag"`
		IsVerified   int    `json:"isVerified"`
		Reward       string `json:"reward"`
		ExactReward  string `json:"exactReward"`
	} `json:"data"`
}
