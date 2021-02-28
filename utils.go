package tushare

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
)

const (
	layoutDay = "20060102"
	layout    = "20060102 15:04:05"
)

// 不能接指针,用以获取返回参数设置
func reflectFields(i interface{}) (fieldParam string) {
	fields := []string{}
	tp := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	for i := 0; i < tp.NumField(); i++ {
		f := tp.Field(i)
		val := v.Field(i).Interface()
		if f.Type.String() == "bool" {
			if val.(bool) {

				fields = append(fields, removeOmitEmpty(f.Tag.Get("json")))
			}
		}
	}
	fieldParam = strings.Join(fields, ",")
	return
}

func removeOmitEmpty(tag string) string {
	// remove omitEmpty
	if strings.HasSuffix(tag, "omitempty") {
		idx := strings.Index(tag, ",")
		if idx > 0 {
			tag = tag[:idx]
		} else {
			tag = ""
		}
	}
	return tag
}

// 压缩请求参数
func buildParams(data interface{}) (params map[string]interface{}) {
	bt, _ := json.Marshal(data)
	params = make(map[string]interface{})
	json.Unmarshal(bt, &params)
	return
}

// 重组数据为[]byte以方便映射struct
func ReflectResponseData(fields []string, data []interface{}) (body []byte, err error) {
	m := make(map[string]interface{})
	if len(fields) != len(data) {
		err = fmt.Errorf("fields(len %d) not fit on data(len %d)", len(fields), len(data))
		return
	}
	if len(fields) == 0 {
		err = errors.New("empty data and fields")
		return
	}
	for n, f := range fields {
		m[f] = data[n]
	}
	body, err = json.Marshal(m)
	return
}

// 转换时间为tushare参数
func Time2TushareTime(t time.Time) string {
	return t.Format(layout)
}

func Time2TushareDayTime(t time.Time) string {
	return t.Format(layoutDay)
}
