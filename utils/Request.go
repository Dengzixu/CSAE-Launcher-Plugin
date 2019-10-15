package utils

import (
	"CSAELauncherPlugin/entity"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func GetPassword(host string, token string) (*entity.ApiPasswordResp, error) {

	tIp := strings.Split(host, ":")[0]
	tPort := strings.Split(host, ":")[1]

	formData := url.Values{}

	apiPasswordResp := entity.ApiPasswordResp{}

	formData.Add("ip", tIp)
	formData.Add("port", tPort)

	payload := strings.NewReader(formData.Encode())

	request, _ := http.NewRequest("POST", "https://hlds.zixutech.cn/index/Server/GetServerPass", payload)

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("token", token)
	request.Header.Add("Content-Length", strconv.Itoa(len(formData.Encode())))

	//tUrl := url.URL{}
	//proxyUrl, _ := tUrl.Parse("http://127.0.0.1:8888")
	//client := &http.Client{Transport: &http.Transport{
	//    Proxy: http.ProxyURL(proxyUrl),
	//},}

	client := &http.Client{}

	resp, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	_ = json.Unmarshal(bodyBytes, &apiPasswordResp)

	if apiPasswordResp.Code != 0 {
		return nil, fmt.Errorf("服务器返回值错误")
	}

	return &apiPasswordResp, nil
}

func GetUserInfo(token string) (*entity.ApiUserInfoResp, error) {
	apiUserInfoResp := entity.ApiUserInfoResp{}

	request, err := http.NewRequest("GET", "https://hlds.zixutech.cn/index/user/getinfo", nil)

	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("token", token)

	client := &http.Client{}

	resp, err := client.Do(request)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	_ = json.Unmarshal(bodyBytes, &apiUserInfoResp)

	return &apiUserInfoResp, nil
}
