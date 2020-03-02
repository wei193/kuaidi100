package kuaidi100

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

//ReqeOrder 云打印请求
type ReqeOrder struct {
	Type            int     `json:"type"`
	PartnerID       string  `json:"partnerId,omitempty"`
	PartnerKey      string  `json:"partnerKey,omitempty"`
	Net             string  `json:"net,omitempty"`
	Kuaidicom       string  `json:"kuaidicom"`
	Kuaidinum       string  `json:"kuaidinum"`
	RecMan          ManInfo `json:"recMan"`
	SendMan         ManInfo `json:"sendMan"`
	Code            string  `json:"code,omitempty"`
	Cargo           string  `json:"cargo,omitempty"`
	Count           int     `json:"count"`
	Weight          float32 `json:"weight"`
	Volumn          float32 `json:"volumn"`
	PayType         string  `json:"payType,omitempty"`
	ExpType         string  `json:"expType,omitempty"`
	Remark          string  `json:"remark,omitempty"`
	ValinsPay       float32 `json:"valinsPay,omitempty"`
	Collection      float32 `json:"collection,omitempty"`
	NeedChild       int     `json:"needChild,omitempty"`
	NeedBack        int     `json:"needBack,omitempty"`
	OrderID         string  `json:"orderId,omitempty"`
	Height          int     `json:"height,omitempty"`
	Width           int     `json:"width,omitempty"`
	Siid            string  `json:"siid"`
	Tempid          string  `json:"tempid"`
	Salt            string  `json:"salt,omitempty"`
	Op              string  `json:"op,omitempty"`
	CallBackURL     string  `json:"callBackUrl"`
	PollCallBackURL string  `json:"pollCallBackUrl,omitempty"`
	Resultv2        string  `json:"resultv2,omitempty"`
}

//ManInfo 收件人信息
type ManInfo struct {
	Name      string `json:"name"`
	Mobile    string `json:"mobile,omitempty"`
	Tel       string `json:"tel,omitempty"`
	ZipCode   string `json:"zipCode,omitempty"`
	Province  string `json:"province,omitempty"`
	City      string `json:"city,omitempty"`
	District  string `json:"district,omitempty"`
	Addr      string `json:"addr,omitempty"`
	PrintAddr string `json:"printAddr,omitempty"`
	Company   string `json:"company,omitempty"`
}

//EOrder 快递100云打印电子面单接口 http://poll.kuaidi100.com/printapi/printtask.do?method=eOrder
func (k *Kuaidi100) EOrder(data ReqeOrder) (res *ResError, err error) {

	buf, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	t := fmt.Sprint(time.Now().Unix())
	param := "method=eOrder"
	param += "&key=" + k.Key
	param += "&sign=" + Md5(string(buf)+t+k.Key+k.Secret)
	param += "&t=" + fmt.Sprint(time.Now().Unix())
	param += "&param=" + string(buf)

	req, err := http.NewRequest("POST", "http://poll.kuaidi100.com/printapi/printtask.do?"+param, nil)
	if err != nil {
		return nil, err
	}
	resBody, err := requset(req)
	if err != nil {
		return nil, err
	}
	res = new(ResError)
	err = json.Unmarshal(resBody, res)
	return

}

// //GetPrintImg 快递100云打印生成图片接口 http://poll.kuaidi100.com/printapi/printtask.do?method=getPrintImg
// func (k *Kuaidi100) GetPrintImg(data ReqeOrder) (res *ResError, err error) {

// 	buf, err := json.Marshal(data)
// 	if err != nil {
// 		return nil, err
// 	}
// 	t := fmt.Sprint(time.Now().Unix())
// 	param := "method=getPrintImg"
// 	param += "&key=" + k.Key
// 	param += "&sign=" + Md5(string(buf)+t+k.Key+k.Secret)
// 	param += "&t=" + fmt.Sprint(time.Now().Unix())
// 	param += "&param=" + string(buf)

// 	req, err := http.NewRequest("POST", "http://poll.kuaidi100.com/printapi/printtask.do?"+param, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	resBody, err := requset(req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	res = new(ResError)
// 	err = json.Unmarshal(resBody, res)
// 	return

// }
