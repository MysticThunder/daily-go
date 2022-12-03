package entity

// GaodeBodyVo 高德地图返回体
type GaodeBodyVo struct {
	Status   string           `json:"status,omitempty"`
	Count    string           `json:"count,omitempty"`
	Info     string           `json:"info,omitempty"`
	Infocode string           `json:"infocode,omitempty"`
	Lives    []GaodeWeatherVo `json:"lives,omitempty"`
}

// GaodeWeatherVo 高德地图数据
type GaodeWeatherVo struct {
	Weather     string
	Temperature string
}

// TianXinVo 天行数据返回体
type TianXinVo struct {
	Code   int
	Msg    string
	Result TianXinContent
}

// TianXinContent 天行数据result
type TianXinContent struct {
	Content string
}

type OptVo struct {
	Value string `json:"value,omitempty"`
	Color string `json:"color,omitempty"`
}

type WxMsgDTO struct {
	Touser     string           `json:"touser,omitempty"`
	TemplateId string           `json:"template_id,omitempty"`
	Topcolor   string           `json:"topcolor,omitempty"`
	Data       map[string]OptVo `json:"data,omitempty"`
}
