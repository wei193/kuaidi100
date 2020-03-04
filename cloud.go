package kuaidi100

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
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

//ReqPrintOrder 云打印自定义请求
type ReqPrintOrder struct {
	Kuaidicom       string `json:"kuaidicom"`
	Kuaidinum       string `json:"kuaidinum"`
	RecManName      string `json:"recManName"`
	RecManMobile    string `json:"recManMobile,omitempty"`
	RecManTel       string `json:"recManTel,omitempty"`
	RecManZipCode   string `json:"recManZipCode,omitempty"`
	RecManProvince  string `json:"recManProvince,omitempty"`
	RecManCity      string `json:"recManCity,omitempty"`
	RecManDistrict  string `json:"recManDistrict,omitempty"`
	RecManAddr      string `json:"recManAddr,omitempty"`
	RecManPrintAddr string `json:"recManPrintAddr,omitempty"`
	RecManCompany   string `json:"recManCompany,omitempty"`

	SendManName      string `json:"sendManName"`
	SendManMobile    string `json:"sendManMobile,omitempty"`
	SendManTel       string `json:"sendManTel,omitempty"`
	SendManZipCode   string `json:"sendManZipCode,omitempty"`
	SendManProvince  string `json:"sendManProvince,omitempty"`
	SendManCity      string `json:"sendManCity,omitempty"`
	SendManDistrict  string `json:"sendManDistrict,omitempty"`
	SendManAddr      string `json:"sendManAddr,omitempty"`
	SendManPrintAddr string `json:"sendManPrintAddr,omitempty"`
	SendManCompany   string `json:"sendManCompany,omitempty"`

	Cargo       string  `json:"cargo,omitempty"`
	Remark      string  `json:"remark,omitempty"`
	Siid        string  `json:"siid"`
	Tempid      string  `json:"tempid"`
	Count       int     `json:"count"`
	Weight      float32 `json:"weight"`
	Volumn      float32 `json:"volumn"`
	Height      int     `json:"height,omitempty"`
	Width       int     `json:"width,omitempty"`
	CallBackURL string  `json:"callBackUrl"`
	Salt        string  `json:"salt,omitempty"`
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
	param += "&param=" + url.QueryEscape(string(buf))
	fmt.Println(string(buf))
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

//PrintOrder 快递100云打印自定义内容接口 http://poll.kuaidi100.com/printapi/printtask.do?method=printOrder
func (k *Kuaidi100) PrintOrder(data ReqPrintOrder) (res *ResError, err error) {

	buf, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error() + "0")
		return nil, err
	}
	t := fmt.Sprint(time.Now().Unix())
	param := "method=printOrder"
	param += "&key=" + k.Key
	param += "&sign=" + Md5(string(buf)+t+k.Key+k.Secret)
	param += "&t=" + fmt.Sprint(time.Now().Unix())
	param += "&param=" + url.QueryEscape(string(buf))
	fmt.Println(param)
	req, err := http.NewRequest("POST", "http://poll.kuaidi100.com/printapi/printtask.do?"+param, nil)
	if err != nil {
		fmt.Println(err.Error() + "1")
		return nil, err
	}
	resBody, err := requset(req)
	if err != nil {
		fmt.Println(err.Error() + "2")
		return nil, err
	}
	res = new(ResError)
	err = json.Unmarshal(resBody, res)
	return

}
