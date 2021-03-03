package tushare

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type TuShare struct {
	token  string
	client *http.Client
}

func NewTuShare(token string) (ts *TuShare) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	ts = new(TuShare)
	ts.client = client
	ts.token = token
	return
}

type TushareResponse struct {
	RequestID string      `json:"request_id"`
	Code      int         `json:"code"` // 不等于0都为错误,-2001 参数错误,-2002 积分不足
	Msg       interface{} `json:"msg"`
	Data      struct {
		Fields []string        `json:"fields"`
		Items  [][]interface{} `json:"items"`
	} `json:"data"`
}

func (ts *TushareResponse) CheckValid() error {
	switch ts.Code {
	case -2001:
		return fmt.Errorf("[TushareAPI Response] code:%d,argument error:%s", ts.Code, ts.Msg)
	case -2002:
		return fmt.Errorf("[TushareAPI Response] code:%d,privilege error:%s", ts.Code, ts.Msg)
	case 0:
		if len(ts.Data.Items) == 0 {
			return fmt.Errorf("[TushareAPI Response] code:%d,msg:%s but empty data", ts.Code, ts.Msg)
		}
		return nil
	default:
		return fmt.Errorf("[TushareAPI Response] code:%d,msg:%s", ts.Code, ts.Msg)
	}
}

type TushareRequest struct {
	APIName string      `json:"api_name"`
	Fields  string      `json:"fields"` // 需要展示的列
	Params  interface{} `json:"params"` // 传参
	Token   string      `json:"token"`
}

func (ts *TushareRequest) CheckValid() error {
	if ts.Token == "" {
		return errors.New("[TushareAPI Request] must set user token")
	}
	if ts.APIName == "" {
		return errors.New("[TushareAPI Request] must set api_name")
	}
	return nil
}
