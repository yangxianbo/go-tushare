package tushare

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
)

const (
	// tushareAPI = "https://api.waditu.com/"
	tushareAPI = "http://api.tushare.pro/"
)

func init() {
	os.Setenv("LANG", "zh_CN.UTF8")

}

func requestTushare(client *http.Client, method string, sendbody interface{}) (tsRsp *TushareResponse, err error) {
	var req *http.Request
	var resp *http.Response
	var bodyJSON []byte
	bodyJSON, err = json.Marshal(sendbody)
	if err != nil {
		return
	}

	// Build send data
	senddata := io.NopCloser(bytes.NewReader(bodyJSON))
	req, err = http.NewRequest(method, tushareAPI, senddata)
	if err != nil {
		return
	}

	// Set http content type
	req.Header.Set("Content-Type", "application/json")
	resp, err = client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	var body []byte
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	// Check mime type of response
	var mimeType string
	mimeType, _, err = mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if err != nil {
		return
	}
	if mimeType != "application/json" {
		err = fmt.Errorf("Could not execute request (%s)", fmt.Sprintf("Response Content-Type is '%s', but should be 'application/json'.", mimeType))
		return
	}

	if resp.StatusCode != 200 {
		err = fmt.Errorf("code:%d return:%s", resp.StatusCode, string(body))
		return
	}
	tsRsp = new(TushareResponse)

	err = json.Unmarshal(body, &tsRsp)
	if err != nil {
		return
	}
	err = tsRsp.CheckValid()
	return
}
