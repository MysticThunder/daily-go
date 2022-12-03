package task

import (
	"bytes"
	"daily/entity"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const appId = "wx8b13609a44fdd432"
const appSecret = "dd9534bf81d215c1d2838e4b96a8c9f3"
const templateId = "Jn2lKrSv10O4Gf33JAjQlr-QRRbhC_nnB-9kA6GPwh0"
const tokenUrl = "https://api.weixin.qq.com/cgi-bin/token"
const pushUrl = "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token="
const loveDay = "2018-06-15"

// GetAccessToken 获取access_token
func GetAccessToken() string {
	params := url.Values{}
	params.Add("grant_type", "client_credential")
	params.Add("appid", appId)
	params.Add("secret", appSecret)

	var temp map[string]interface{}
	err := json.Unmarshal(LaunchGet(params, tokenUrl), &temp)
	if err != nil {
		fmt.Printf("%v\n", err.Error())
	}
	accessToken := temp["access_token"].(string)
	return accessToken
}

// Push 推送消息
func Push() {
	wxMsgDTO := entity.WxMsgDTO{
		Touser:     "ouP8y53goQ1gRyIC0SLO8-0w_RNs",
		TemplateId: templateId,
		Topcolor:   "#FF1493",
		Data:       assemble(),
	}
	//post请求提交json数据
	ba, _ := json.Marshal(wxMsgDTO)

	resp, err := http.Post(pushUrl+GetAccessToken(), "application/json", bytes.NewBuffer(ba))
	if err != nil {
		log.Printf("========================> %v\n", err.Error())
	}
	log.Printf("------------------------> %v\n", "推送消息成功")
	resp.Body.Close()
}

// 组装消息
func assemble() map[string]entity.OptVo {
	date := map[string]entity.OptVo{}
	date["now"] = entity.OptVo{
		Value: time.Now().Format("2006年01🈷02日") + "\n不管是哪一天都是想锤你的一天",
		Color: "#800080",
	}

	date["weather"] = entity.OptVo{
		Value: "南京市🌤" + QueryWeather().Lives[0].Weather,
		Color: "#FF0000",
	}

	date["now_temperature"] = entity.OptVo{
		Value: QueryWeather().Lives[0].Temperature + "℃",
		Color: "#008000",
	}

	date["love_day"] = entity.OptVo{
		Value: strconv.Itoa(calLoveDay()) + "	天💖",
		Color: "#6A5ACD",
	}

	date["rainbowPi"] = entity.OptVo{
		Value: QueryRainbow().Result.Content,
		Color: "#A52A2A",
	}

	date["zaoan"] = entity.OptVo{
		Value: queryZaoan().Result.Content,
		Color: "#191970",
	}
	return date
}

// CalLoveDay 计算相恋多少天
func calLoveDay() int {
	loveDay, _ := time.Parse("2006-01-02", loveDay)
	duringDay := int(time.Now().Sub(loveDay).Hours() / 24)
	return duringDay
}
