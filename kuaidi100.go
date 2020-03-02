package kuaidi100

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

//Kuaidi100 快递100配置
type Kuaidi100 struct {
	Key      string
	Customer string
	Secret   string
	Userid   string
	Smart    string
}

//NewKuaidi100 创建快递100对象
func NewKuaidi100() {

}

//ReqPoll 订阅发送数据结构
type ReqPoll struct {
	Company    string         `json:"company"`
	Number     string         `json:"number"`
	From       string         `json:"from,omitempty"`
	To         string         `json:"to,omitempty"`
	Key        string         `json:"key"`
	Parameters PollParameters `json:"parameters"`
}

//ReqPollQuery 实时查询
type ReqPollQuery struct {
	Com      string `json:"com"`
	Num      string `json:"num"`
	Phone    string `json:"phone,omitempty"`
	From     string `json:"from,omitempty"`
	To       string `json:"to,omitempty"`
	Resultv2 int    `json:"resultv2,omitempty"`
	Show     string `json:"show,omitempty"`
	Order    string `json:"order,omitempty"`
}

//ResError 返回错误信息
type ResError struct {
	Result     bool   `json:"result"`
	ReturnCode string `json:"returnCode"`
	Message    string `json:"message"`
}

//PollParameters 订阅Parameters
type PollParameters struct {
	Callbackurl        string `json:"callbackurl"`
	Salt               string `json:"salt,omitempty"`
	Resultv2           string `json:"resultv2,omitempty"`
	AutoCom            string `json:"autoCom,omitempty"`
	InterCom           string `json:"interCom,omitempty"`
	DepartureCountry   string `json:"departureCountry,omitempty"`
	DepartureCom       string `json:"departureCom,omitempty"`
	DestinationCountry string `json:"destinationCountry,omitempty"`
	DestinationCom     string `json:"destinationCom,omitempty"`
	Phone              string `json:"phone,omitempty"`
}

//ResPollCallBack 订阅返回
type ResPollCallBack struct {
	Status     string `json:"status"`
	Message    string `json:"message"`
	AutoCheck  string `json:"autoCheck"`
	ComOld     string `json:"comOld"`
	ComNew     string `json:"comNew"`
	LastResult Result `json:"lastResult"`
	DestResult Result `json:"destResult"`
}

//Result 物流结果
type Result struct {
	State   string       `json:"state"`
	Ischeck string       `json:"ischeck"`
	Com     string       `json:"com"`
	Nu      string       `json:"nu"`
	Data    []ResultData `json:"data"`
}

//ResultData 物流记录
type ResultData struct {
	Context  string `json:"context"`
	Time     string `json:"time"`
	Ftime    string `json:"ftime"`
	Status   string `json:"status"`
	AreaCode string `json:"areaCode"`
	AreaName string `json:"areaName"`
}

//ResAuto 自动查询快递公司
type ResAuto struct {
	ComCode string `json:"comCode"`
}

//Poll 快递100订阅请 https://poll.kuaidi100.com/poll
func (k *Kuaidi100) Poll(data ReqPoll) (res *ResError, err error) {
	data.Key = k.Key
	buf, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	param := "schema=json"
	param += "&param=" + string(buf)

	req, err := http.NewRequest("POST", "https://poll.kuaidi100.com/poll?"+param, nil)
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

//PollQuery 快递100实时查询 https://poll.kuaidi100.com/poll/query.do
func (k *Kuaidi100) PollQuery(data ReqPollQuery) (res *Result, err error) {

	buf, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	param := "customer=" + k.Customer
	param += "&sign=" + Md5(string(buf)+k.Key+k.Customer)
	param += "&param=" + string(buf)
	req, err := http.NewRequest("POST", "https://poll.kuaidi100.com/poll/query.do?"+param, nil)
	if err != nil {
		return nil, err
	}

	resBody, err := requset(req)
	if err != nil {
		return nil, err
	}
	res = new(Result)
	err = json.Unmarshal(resBody, res)
	return
}

//Autonumber 快递100智能判断 http://www.kuaidi100.com/autonumber/auto
func (k *Kuaidi100) Autonumber(num string) (res []ResAuto, err error) {
	param := "num=" + num
	param += "&key=" + k.Key
	req, err := http.NewRequest("POST", "http://www.kuaidi100.com/autonumber/auto?"+param, nil)
	if err != nil {
		return nil, err
	}

	resBody, err := requset(req)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resBody, &res)
	return
}

func requset(req *http.Request) ([]byte, error) {
	client := &http.Client{Timeout: 60 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

//Md5 计算MD5
func Md5(str string) string {
	t := md5.New()
	io.WriteString(t, str)
	return fmt.Sprintf("%X", t.Sum(nil))
}
