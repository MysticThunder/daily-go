package task

import (
	"daily/entity"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const KEY = "2dc9faa7abb06c0d2cd3b40599b191e6"
const CITY = "南京市"
const WeatherUrl = "https://restapi.amap.com/v3/weather/weatherInfo"
const TianKey = "11d9587c7c357a81c283c657d3e1b4f3"
const Rainbow = "https://apis.tianapi.com/caihongpi/index"
const ZaoAn = "https://apis.tianapi.com/zaoan/index"

// QueryWeather 高德地图查询天气
func QueryWeather() entity.GaodeBodyVo {
	params := url.Values{}
	params.Add("key", KEY)
	params.Add("city", CITY)
	//请求地址
	stringUrl := WeatherUrl
	body := LaunchGet(params, stringUrl)
	var BodyVo = entity.GaodeBodyVo{}
	//转换为自定义结构体
	_ = json.Unmarshal(body, &BodyVo)
	return BodyVo
}

// QueryRainbow  获取彩虹屁
func QueryRainbow() entity.TianXinVo {
	params := url.Values{}
	params.Add("key", TianKey)
	//请求地址
	stringUrl := Rainbow
	body := LaunchGet(params, stringUrl)
	var BodyVo = entity.TianXinVo{}
	//转换为自定义结构体
	_ = json.Unmarshal(body, &BodyVo)
	return BodyVo
}

// queryZaoan  获取早安心语
func queryZaoan() entity.TianXinVo {
	params := url.Values{}
	params.Add("key", TianKey)
	//请求地址
	stringUrl := ZaoAn
	body := LaunchGet(params, stringUrl)
	var BodyVo = entity.TianXinVo{}
	//转换为自定义结构体
	_ = json.Unmarshal(body, &BodyVo)
	return BodyVo
}

// LaunchGet 发起get请求
func LaunchGet(params url.Values, stringUrl string) []byte {
	reqUrl, _ := url.Parse(stringUrl)
	//组合url
	reqUrl.RawQuery = params.Encode()
	//发起get请求
	resp, _ := http.Get(reqUrl.String())
	defer resp.Body.Close()
	//解析请求信息
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
	}

	return body
}
