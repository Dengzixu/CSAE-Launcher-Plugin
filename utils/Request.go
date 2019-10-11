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

func GetPassword(host string, token string) string {

	formData := url.Values{}

	apiPasswordResp := entity.ApiPasswordResp{}

	tIp := string(strings.Split(host, ":")[0])
	tPort := string(strings.Split(host, ":")[1])

	formData.Add("ip", tIp)
	formData.Add("port", tPort)

	payload := strings.NewReader(formData.Encode())

	request, err := http.NewRequest("POST", "https://hlds.zixutech.cn/index/Server/GetServerPass", payload)

	if err != nil {
		return ""
	}

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
	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	tempd, _ := ioutil.ReadAll(resp.Body)

	_ = json.Unmarshal(tempd, &apiPasswordResp)

	return apiPasswordResp.Data.ServerPass
}
