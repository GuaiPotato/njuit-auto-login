package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	// "os/exec"
	// "runtime"

	"github.com/spf13/viper"
)

type LoginResponse struct {
	ReplyCode int     `json:"reply_code"`
	Results   Results `json:"results"`
	ServeTime int     `json:"serve_time"`
}

type Results struct {
	AcctSessionID string `json:"acctsessionid"`
	AcctStartTime int    `json:"acctstarttime"`
	AreaName      string `json:"area_name"`
	Balance       int    `json:"balance"`
	Domain        string `json:"domain"`
	FullName      string `json:"fullname"`
	MAC           string `json:"mac"`
	ServiceName   string `json:"service_name"`
	UserIPv4      int    `json:"user_ipv4"`
	UserIPv6      string `json:"user_ipv6"`
	Username      string `json:"username"`
}

func main() {
	// 获取当前工作目录
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return
	}

	// 设置配置文件的名称和路径
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(dir)

	// 读取配置文件
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return
	}

	// 从配置文件中获取参数值
	username := viper.GetString("username")
	domain := viper.GetString("domain")
	password := viper.GetString("password")
	url := viper.GetString("url")

	// 构造请求体数据
	requestBody := map[string]interface{}{
		"domain":   domain,
		"username": username,
		"password": password,
		"url":      url,
	}

	// 将请求体数据转换为JSON格式
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("JSON marshal error:", err)
		return
	}

	// 发送HTTP POST请求
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("HTTP request error:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应数据
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read response error:", err)
		return
	}

	// 解析响应数据
	var loginResponse LoginResponse
	err = json.Unmarshal(responseData, &loginResponse)
	if err != nil {
		fmt.Println("JSON unmarshal error:", err)
		return
	}

	// 根据响应判断登录是否成功
	if loginResponse.ReplyCode == 0 {
		fmt.Println("校园网登录成功")
	} else {
		fmt.Println("校园网登录失败")
	}

	// 保留画面2s
	time.Sleep(2 * time.Second)
}
