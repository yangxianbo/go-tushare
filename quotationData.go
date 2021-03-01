package tushare

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/montanaflynn/stats"
)

type QuotationRequest struct {
	TsCode    string `json:"ts_code,omitempty"`    // str	N	股票代码(支持多个股票同时提取,逗号分隔)
	TradeDate string `json:"trade_date,omitempty"` // str	N	交易日期(YYYYMMDD)
	StartDate string `json:"start_date,omitempty"` // str	N	开始日期(YYYYMMDD,提取分钟数据请用2019-09-01 09:00:00这种格式)
	EndDate   string `json:"end_date,omitempty"`   // str	N	结束日期(YYYYMMDD,提取分钟数据请用2019-09-01 09:00:00这种格式)
}

type QuotationItems struct {
	TsCode    bool `json:"ts_code,omitempty"`    // str	股票代码
	TradeDate bool `json:"trade_date,omitempty"` // str	交易日期
	Open      bool `json:"open,omitempty"`       // float	(1/5/15/30/60分钟) D日 W周 M月 开盘价
	High      bool `json:"high,omitempty"`       // float	(1/5/15/30/60分钟) D日 W周 M月 最高价
	Low       bool `json:"low,omitempty"`        // float	(1/5/15/30/60分钟) D日 W周 M月 最低价
	Close     bool `json:"close,omitempty"`      // float	(1/5/15/30/60分钟) D日 W周 M月 收盘价
	PreClose  bool `json:"pre_close,omitempty"`  // float	上一(1/5/15/30/60分钟) D日 W周 M月 收盘价
	Change    bool `json:"change,omitempty"`     // float	(1/5/15/30/60分钟) D日 W周 M月 涨跌额
	PctChg    bool `json:"pct_chg,omitempty"`    // float	(1/5/15/30/60分钟) D日 W周 M月 涨跌幅
	Vol       bool `json:"vol,omitempty"`        // float	(1/5/15/30/60分钟) D日 W周 M月 成交量 (手)
	Amount    bool `json:"amount,omitempty"`     // float	(1/5/15/30/60分钟) D日 W周 M月 成交额 (千元)
}

func (item QuotationItems) All() QuotationItems {
	item.TsCode = true
	item.TradeDate = true
	item.Open = true
	item.High = true
	item.Low = true
	item.Close = true
	item.PreClose = true
	item.Change = true
	item.PctChg = true
	item.Vol = true
	item.Amount = true
	return item
}

type QuotationData struct {
	TsCode    string  `json:"ts_code,omitempty"`    // str	股票代码
	TradeDate string  `json:"trade_date,omitempty"` // str	交易日期
	Open      float64 `json:"open,omitempty"`       // float	(1/5/15/30/60分钟) D日 W周 M月 开盘价
	High      float64 `json:"high,omitempty"`       // float	(1/5/15/30/60分钟) D日 W周 M月 最高价
	Low       float64 `json:"low,omitempty"`        // float	(1/5/15/30/60分钟) D日 W周 M月 最低价
	Close     float64 `json:"close,omitempty"`      // float	(1/5/15/30/60分钟) D日 W周 M月 收盘价
	PreClose  float64 `json:"pre_close,omitempty"`  // float	上一(1/5/15/30/60分钟) D日 W周 M月 收盘价
	Change    float64 `json:"change,omitempty"`     // float	(1/5/15/30/60分钟) D日 W周 M月 涨跌额
	PctChg    float64 `json:"pct_chg,omitempty"`    // float	(1/5/15/30/60分钟) D日 W周 M月 涨跌幅
	Vol       float64 `json:"vol,omitempty"`        // float	(1/5/15/30/60分钟) D日 W周 M月 成交量 (手)
	Amount    float64 `json:"amount,omitempty"`     // float	(1/5/15/30/60分钟) D日 W周 M月 成交额 (千元)
}

func AssembleQuotationData(tsRsp *TushareResponse) []*QuotationData {
	tsData := []*QuotationData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(QuotationData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 日线行情,交易日每天15点～16点之间.本接口是未复权行情,停牌期间不提供数据,基础积分每分钟内最多调取500次,每次5000条数据,相当于23年历史,用户获得超过5000积分正常调取无频次限制
func (ts *TuShare) Daily(params QuotationRequest, items QuotationItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "daily",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

// 周线行情,单次最大4500行,总量不限制,用户需要至少300积分才可以调取,具体请参阅积分获取办法 https://tushare.pro/document/1?doc_id=13
func (ts *TuShare) Weekly(params QuotationRequest, items QuotationItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "weekly",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

// 月线行情,单次最大4500行,总量不限制,用户需要至少300积分才可以调取,具体请参阅积分获取办法 https://tushare.pro/document/1?doc_id=13
func (ts *TuShare) Monthly(params QuotationRequest, items QuotationItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "monthly",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type AdjFactorItems struct {
	TsCode    bool `json:"ts_code,omitempty"`    // str	股票代码
	TradeDate bool `json:"trade_date,omitempty"` // str	交易日期
	AdjFactor bool `json:"adj_factor,omitempty"` // float	复权因子
}

func (item AdjFactorItems) All() AdjFactorItems {
	item.TsCode = true
	item.TradeDate = true
	item.AdjFactor = true
	return item
}

type AdjFactorData struct {
	TsCode    string  `json:"ts_code,omitempty"`    // str	股票代码
	TradeDate string  `json:"trade_date,omitempty"` // str	交易日期
	AdjFactor float64 `json:"adj_factor,omitempty"` // float	复权因子
}

func AssembleAdjFactorData(tsRsp *TushareResponse) []*AdjFactorData {
	tsData := []*AdjFactorData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(AdjFactorData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 复权因子,获取股票复权因子,早上9点30分更新,可提取单只股票全部历史复权因子,也可以提取单日全部股票的复权因子
func (ts *TuShare) AdjFactor(params QuotationRequest, items AdjFactorItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "adj_factor",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

// 分钟级别的请求,1000分以下每日5次
type ProBarRequest struct {
	TsCode    string      `json:"ts_code,omitempty"`    // str	Y	股票代码
	StartDate string      `json:"start_date,omitempty"` // str	N	开始日期 不填写默认前365天 (提取分钟数据请用2019-09-01 09:00:00这种格式)
	EndDate   string      `json:"end_date,omitempty"`   // str	N	结束日期 不填写默认当天 (提取分钟数据请用2019-09-01 09:00:00这种格式)
	Adj       string      `json:"adj,omitempty"`        // str	N	复权类型(只针对股票):为空时未复权 qfq前复权 hfq后复权 , 默认空
	Freq      string      `json:"freq,omitempty"`       // str	Y	数据频度 :1MIN表示1分钟(1/5/15/30/60分钟) D日线 W周线 M月线 ,默认D
	MA        []int       `json:"ma,omitempty"`         // list	N	均线,支持任意周期的均价和均量,输入任意合理int数值(5,20,120)
	BOLL      BOLLRequest `json:"boll,omitempty"`       // BOLLRequest	N	布林线,输入标准差和宽度(20,2)
}

type BOLLRequest struct {
	SD    int `json:"sd,omitempty"`    // int	N	标准差
	Width int `json:"width,omitempty"` // int	N	宽度
}

func (params ProBarRequest) init() (newParams ProBarRequest, err error) {
	if params.TsCode == "" {
		err = errors.New("must set tscode")
		return
	}
	newParams = params
	if params.StartDate == "" || params.EndDate == "" {
		newParams.StartDate = Time2TushareTime(time.Unix(time.Now().Unix()-31536000, 0))
		newParams.StartDate = Time2TushareTime(time.Now())
	}
	switch strings.ToLower(params.Adj) {
	case "qfq", "hfq":
		newParams.Adj = strings.ToLower(params.Adj)
	default:
		newParams.Adj = ""
	}
	switch strings.ToUpper(params.Freq) {
	case "1MIN", "5MIN", "15MIN", "30MIN", "60MIN", "D", "W", "M":
		newParams.Freq = strings.ToUpper(params.Freq)
	default:
		newParams.Freq = "D"
	}
	return
}

type ProBarData struct {
	QuotationData
	DailyBasic *DailyBasicData `json:"daily_basic,omitempty"` // 每日指标数据
	MA         map[int]float64 `json:"ma,omitempty"`          // list	N	均线,支持任意周期的均价和均量,输入任意合理int数值
	BOLL       BOLLData        `json:"boll,omitempty"`        // BOLLData	N	均线,支持任意周期的均价和均量,输入任意合理int数值
}

type BOLLData struct {
	UP  float64 `json:"up,omitempty"`  // int	N	上线
	MID float64 `json:"mid,omitempty"` // int	N	中线
	LOW float64 `json:"low,omitempty"` // int	N	下线
}

// 复权行情,获取股票复权行情,只支持股票,月复权会少最近一个月的数据,为了方便计算,数据的时间排序与pysdk相反,时间排序为过去->现在
func (ts *TuShare) ProBar(params ProBarRequest) (data []*ProBarData, err error) {
	params, err = params.init()
	if err != nil {
		return
	}
	// 获取行情数据
	var dataRsp, dbRsq, adjRsp *TushareResponse
	req := QuotationRequest{
		TsCode:    params.TsCode,
		StartDate: params.StartDate,
		EndDate:   params.EndDate,
	}
	dataItems := QuotationItems{}.All()
	switch params.Freq {
	case "D":
		dataRsp, err = ts.Daily(req, dataItems)
		if err != nil {
			return
		}
	case "W":
		dataRsp, err = ts.Weekly(req, dataItems)
		if err != nil {
			return
		}
	case "M":
		dataRsp, err = ts.Monthly(req, dataItems)
		if err != nil {
			return
		}
	}
	// 组装行情数据
	qData := AssembleQuotationData(dataRsp)
	// 获取每日指标数据
	dbRsq, err = ts.DailyBasic(req, DailyBasicItems{}.All())
	if err != nil {
		return
	}
	// 组装每日指标数据
	dData := AssembleDailyBasicData(dbRsq)
	// 组装每日指标索引
	dMap := make(map[string]*DailyBasicData)
	for _, d := range dData {
		dMap[d.TradeDate] = d
	}
	// 获取复权因子
	adjRsp, err = ts.AdjFactor(req, AdjFactorItems{}.All())
	if err != nil {
		return
	}
	// 组装复权因子数据
	aData := AssembleAdjFactorData(adjRsp)
	// 组装复权因子索引
	aMap := make(map[string]float64)
	for _, v := range aData {
		aMap[v.TradeDate] = v.AdjFactor
	}
	switch params.Adj {
	case "qfq":
		data = qfq(qData, dMap, aMap, aData[0].AdjFactor)
	case "hfq":
		data = hfq(qData, dMap, aMap)
	default:
		data = nfq(qData, dMap)
	}
	// MA
	closeData := []float64{}
	for _, p := range data {
		closeData = append(closeData, p.Close)
	}
	for _, m := range params.MA {
		maResult := ma(m, closeData)
		for n, maNum := range maResult {
			if data[n].MA == nil {
				data[n].MA = make(map[int]float64)
			}
			data[n].MA[m] = maNum
		}
	}
	// BOLL
	if params.BOLL.SD > 0 && params.BOLL.Width > 0 {
		newBoll(params.BOLL.SD, params.BOLL.Width).boll(data)
	}
	return
}

func qfq(qData []*QuotationData, dMap map[string]*DailyBasicData, aMap map[string]float64, adjNew float64) (pData []*ProBarData) {
	pData = []*ProBarData{}
	for i := len(qData) - 1; i >= 0; i-- {
		q := qData[i]
		p := new(ProBarData)
		p.TsCode = q.TsCode
		p.TradeDate = q.TradeDate
		p.Open = q.Open * aMap[q.TradeDate] / adjNew
		p.High = q.High * aMap[q.TradeDate] / adjNew
		p.Low = q.Low * aMap[q.TradeDate] / adjNew
		p.Close = q.Close * aMap[q.TradeDate] / adjNew
		p.PreClose = q.PreClose * aMap[q.TradeDate] / adjNew
		p.Change = q.Change * aMap[q.TradeDate] / adjNew
		p.PctChg = q.PctChg * aMap[q.TradeDate] / adjNew
		p.Vol = q.Vol
		p.Amount = q.Amount
		p.DailyBasic = dMap[q.TradeDate]
		pData = append(pData, p)
	}
	return
}

func hfq(qData []*QuotationData, dMap map[string]*DailyBasicData, aMap map[string]float64) (pData []*ProBarData) {
	pData = []*ProBarData{}
	for i := len(qData) - 1; i >= 0; i-- {
		q := qData[i]
		p := new(ProBarData)
		p.TsCode = q.TsCode
		p.TradeDate = q.TradeDate
		p.Open = q.Open * aMap[q.TradeDate]
		p.High = q.High * aMap[q.TradeDate]
		p.Low = q.Low * aMap[q.TradeDate]
		p.Close = q.Close * aMap[q.TradeDate]
		p.PreClose = q.PreClose * aMap[q.TradeDate]
		p.Change = q.Change * aMap[q.TradeDate]
		p.PctChg = q.PctChg * aMap[q.TradeDate]
		p.Vol = q.Vol
		p.Amount = q.Amount
		p.DailyBasic = dMap[q.TradeDate]
		pData = append(pData, p)
	}
	return
}

func nfq(qData []*QuotationData, dMap map[string]*DailyBasicData) (pData []*ProBarData) {
	pData = []*ProBarData{}
	for i := len(qData) - 1; i >= 0; i-- {
		q := qData[i]
		p := new(ProBarData)
		p.TsCode = q.TsCode
		p.TradeDate = q.TradeDate
		p.Open = q.Open
		p.High = q.High
		p.Low = q.Low
		p.Close = q.Close
		p.PreClose = q.PreClose
		p.Change = q.Change
		p.PctChg = q.PctChg
		p.Vol = q.Vol
		p.Amount = q.Amount
		p.DailyBasic = dMap[q.TradeDate]
		pData = append(pData, p)
	}
	return
}

func ma(nday int, data []float64) []float64 {
	if nday == 0 {
		return data
	}
	ma := []float64{}
	var p int
	for n := range data {
		if n < nday-1 {
			ma = append(ma, 0)
		}
		if n == nday-1 {
			m, _ := stats.Mean(data[p:nday])
			ma = append(ma, m)
			p++
		}
		if n > nday-1 {
			m, _ := stats.Mean(data[p : nday+p])
			ma = append(ma, m)
			p++
		}
	}
	return ma
}

type SuspendRequest struct {
	TsCode      string `json:"ts_code,omitempty"`      // str	N	股票代码(三选一)
	SuspendDate string `json:"suspend_date,omitempty"` // str	N	停牌日期(三选一)
	ResumeDate  string `json:"resume_date,omitempty"`  // str	N	复牌日期(三选一)
}

type SuspendItems struct {
	TsCode        bool `json:"ts_code,omitempty"`        // str TS代码
	SuspendDate   bool `json:"suspend_date,omitempty"`   // str	停牌日期
	ResumeDate    bool `json:"resume_date,omitempty"`    // str	复牌日期
	AnnDate       bool `json:"ann_date,omitempty"`       // str	公告日期
	SuspendReason bool `json:"suspend_reason,omitempty"` // str	停牌原因
	ReasonType    bool `json:"reason_type,omitempty"`    // str	停牌原因类别
}

func (item SuspendItems) All() SuspendItems {
	item.TsCode = true
	item.SuspendDate = true
	item.ResumeDate = true
	item.AnnDate = true
	item.SuspendReason = true
	item.ReasonType = true
	return item
}

type SuspendData struct {
	TsCode        string `json:"ts_code,omitempty"`        // str TS代码
	SuspendDate   string `json:"suspend_date,omitempty"`   // str	停牌日期
	ResumeDate    string `json:"resume_date,omitempty"`    // str	复牌日期
	AnnDate       string `json:"ann_date,omitempty"`       // str	公告日期
	SuspendReason string `json:"suspend_reason,omitempty"` // str	停牌原因
	ReasonType    string `json:"reason_type,omitempty"`    // str	停牌原因类别
}

func AssembleSuspendData(tsRsp *TushareResponse) []*SuspendData {
	tsData := []*SuspendData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(SuspendData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 获取股票每日停复牌信息,不定期更新
func (ts *TuShare) Suspend(params SuspendRequest, items SuspendItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "suspend",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type SuspendDRequest struct {
	TsCode      string `json:"ts_code,omitempty"`      // str	N	股票代码(可输入多值)
	TradeDate   string `json:"trade_date,omitempty"`   // str	N	交易日日期
	StartDate   string `json:"start_date,omitempty"`   // str	N	停复牌查询开始日期
	EndDate     string `json:"end_date,omitempty"`     // str	N	停复牌查询结束日期
	SuspendType string `json:"suspend_type,omitempty"` // str	N	停复牌类型:S-停牌,R-复牌一)
}

type SuspendDItems struct {
	TsCode        bool `json:"ts_code,omitempty"`        // str TS代码
	TradeDate     bool `json:"trade_date,omitempty"`     // str	Y	停复牌日期
	SuspendTiming bool `json:"suspend_timing,omitempty"` // str	Y	日内停牌时间段
	SuspendType   bool `json:"suspend_type,omitempty"`   // str	Y	停复牌类型:S-停牌,R-复牌
}

func (item SuspendDItems) All() SuspendDItems {
	item.TsCode = true
	item.TradeDate = true
	item.SuspendTiming = true
	item.SuspendType = true
	return item
}

type SuspendDData struct {
	TsCode        string `json:"ts_code,omitempty"`        // str TS代码
	TradeDate     string `json:"trade_date,omitempty"`     // str	Y	停复牌日期
	SuspendTiming string `json:"suspend_timing,omitempty"` // str	Y	日内停牌时间段
	SuspendType   string `json:"suspend_type,omitempty"`   // str	Y	停复牌类型:S-停牌,R-复牌
}

func AssembleSuspendDData(tsRsp *TushareResponse) []*SuspendDData {
	tsData := []*SuspendDData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(SuspendDData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 按日期方式获取股票每日停复牌信息,不定期更新
func (ts *TuShare) SuspendD(params SuspendDRequest, items SuspendDItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "suspend_d",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type DailyBasicItems struct {
	TsCode        bool `json:"ts_code,omitempty"`         // str TS代码
	TradeDate     bool `json:"trade_date,omitempty"`      // str	交易日期
	Close         bool `json:"close,omitempty"`           // float	当日收盘价
	TurnoverRate  bool `json:"turnover_rate,omitempty"`   // float	换手率(%)
	TurnoverRateF bool `json:"turnover_rate_f,omitempty"` // float	换手率(自由流通股)
	VolumeRatio   bool `json:"volume_ratio,omitempty"`    // float	量比
	Pe            bool `json:"pe,omitempty"`              // float	市盈率(总市值/净利润, 亏损的PE为空)
	PeTTM         bool `json:"pe_ttm,omitempty"`          // float	市盈率(TTM,亏损的PE为空)
	Pb            bool `json:"pb,omitempty"`              // float	市净率(总市值/净资产)
	Ps            bool `json:"ps,omitempty"`              // float	市销率
	PsTTM         bool `json:"ps_ttm,omitempty"`          // float	市销率(TTM)
	DvRatio       bool `json:"dv_ratio,omitempty"`        // float	股息率 (%)
	DvTtm         bool `json:"dv_ttm,omitempty"`          // float	股息率(TTM)(%)
	TotalShare    bool `json:"total_share,omitempty"`     // float	总股本 (万股)
	FloatShare    bool `json:"float_share,omitempty"`     // float	流通股本 (万股)
	FreeShare     bool `json:"free_share,omitempty"`      // float	自由流通股本 (万)
	TotalMV       bool `json:"total_mv,omitempty"`        // float	总市值 (万元)
	CircMV        bool `json:"circ_mv,omitempty"`         // float	流通市值(万元)
}

func (item DailyBasicItems) All() DailyBasicItems {
	item.TsCode = true
	item.TradeDate = true
	item.Close = true
	item.TurnoverRate = true
	item.TurnoverRateF = true
	item.VolumeRatio = true
	item.Pe = true
	item.PeTTM = true
	item.Pb = true
	item.Ps = true
	item.PsTTM = true
	item.DvRatio = true
	item.DvTtm = true
	item.TotalShare = true
	item.FloatShare = true
	item.FreeShare = true
	item.TotalMV = true
	item.CircMV = true
	return item
}

type DailyBasicData struct {
	TsCode        string  `json:"ts_code,omitempty"`         // str TS代码
	TradeDate     string  `json:"trade_date,omitempty"`      // str	交易日期
	Close         float64 `json:"close,omitempty"`           // float	当日收盘价
	TurnoverRate  float64 `json:"turnover_rate,omitempty"`   // float	换手率(%)
	TurnoverRateF float64 `json:"turnover_rate_f,omitempty"` // float	换手率(自由流通股)
	VolumeRatio   float64 `json:"volume_ratio,omitempty"`    // float	量比
	Pe            float64 `json:"pe,omitempty"`              // float	市盈率(总市值/净利润, 亏损的PE为空)
	PeTTM         float64 `json:"pe_ttm,omitempty"`          // float	市盈率(TTM,亏损的PE为空)
	Pb            float64 `json:"pb,omitempty"`              // float	市净率(总市值/净资产)
	Ps            float64 `json:"ps,omitempty"`              // float	市销率
	PsTTM         float64 `json:"ps_ttm,omitempty"`          // float	市销率(TTM)
	DvRatio       float64 `json:"dv_ratio,omitempty"`        // float	股息率 (%)
	DvTtm         float64 `json:"dv_ttm,omitempty"`          // float	股息率(TTM)(%)
	TotalShare    float64 `json:"total_share,omitempty"`     // float	总股本 (万股)
	FloatShare    float64 `json:"float_share,omitempty"`     // float	流通股本 (万股)
	FreeShare     float64 `json:"free_share,omitempty"`      // float	自由流通股本 (万)
	TotalMV       float64 `json:"total_mv,omitempty"`        // float	总市值 (万元)
	CircMV        float64 `json:"circ_mv,omitempty"`         // float	流通市值(万元)
}

func AssembleDailyBasicData(tsRsp *TushareResponse) []*DailyBasicData {
	tsData := []*DailyBasicData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(DailyBasicData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 获取全部股票每日重要的基本面指标,可用于选股分析,报表展示等.交易日每日15点～17点之间更新数据,用户需要至少600积分才可以调取,具体请参阅积分获取办法 https://tushare.pro/document/1?doc_id=13
func (ts *TuShare) DailyBasic(params QuotationRequest, items DailyBasicItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "daily_basic",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type StkLimitItems struct {
	TsCode    bool `json:"ts_code,omitempty"`    // str	Y	TS股票代码
	TradeDate bool `json:"trade_date,omitempty"` // str	Y	交易日期
	PreClose  bool `json:"pre_close,omitempty"`  // float	N	昨日收盘价
	UpLimit   bool `json:"up_limit,omitempty"`   // float	Y	涨停价
	DownLimit bool `json:"down_limit,omitempty"` // float	Y	跌停价
}

func (item StkLimitItems) All() StkLimitItems {
	item.TsCode = true
	item.TradeDate = true
	item.PreClose = true
	item.UpLimit = true
	item.DownLimit = true
	return item
}

type StkLimitData struct {
	TsCode    string  `json:"ts_code,omitempty"`    // str	Y	TS股票代码
	TradeDate string  `json:"trade_date,omitempty"` // str	Y	交易日期
	PreClose  float64 `json:"pre_close,omitempty"`  // float	N	昨日收盘价
	UpLimit   float64 `json:"up_limit,omitempty"`   // float	Y	涨停价
	DownLimit float64 `json:"down_limit,omitempty"` // float	Y	跌停价
}

func AssembleStkLimitData(tsRsp *TushareResponse) []*StkLimitData {
	tsData := []*StkLimitData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(StkLimitData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 获取全市场(包含A/B股和基金)每日涨跌停价格,包括涨停价格,跌停价格等,每个交易日8点40左右更新当日股票涨跌停价格,单次最多提取4800条记录,可循环调取,总量不限制,用户需要至少120积分才可以调取,具体请参阅积分获取办法 https://tushare.pro/document/1?doc_id=13
func (ts *TuShare) StkLimit(params QuotationRequest, items StkLimitItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "stk_limit",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type MoneyflowHSGTRequest struct {
	TradeDate string `json:"trade_date,omitempty"` // str	N	交易日期(YYYYMMDD,二选一)
	StartDate string `json:"start_date,omitempty"` // str	N	开始日期(YYYYMMDD,二选一)
	EndDate   string `json:"end_date,omitempty"`   // str	N	结束日期(YYYYMMDD,提取分钟数据请用2019-09-01 09:00:00这种格式)
}

type MoneyflowHSGTItems struct {
	TradeDate  bool `json:"trade_date,omitempty"`  // str	交易日期
	GGTSS      bool `json:"ggt_ss,omitempty"`      // float	港股通(上海)
	GGTSZ      bool `json:"ggt_sz,omitempty"`      // float	港股通(深圳)
	HGT        bool `json:"hgt,omitempty"`         // float	沪股通(百万元)
	SGT        bool `json:"sgt,omitempty"`         // float	深股通(百万元)
	NorthMoney bool `json:"north_money,omitempty"` // float	北向资金(百万元)
	SouthMoney bool `json:"south_money,omitempty"` // float	南向资金(百万元)
}

func (item MoneyflowHSGTItems) All() MoneyflowHSGTItems {
	item.TradeDate = true
	item.GGTSS = true
	item.GGTSZ = true
	item.HGT = true
	item.SGT = true
	item.NorthMoney = true
	item.SouthMoney = true
	return item
}

type MoneyflowHSGTData struct {
	TradeDate  string  `json:"trade_date,omitempty"`  // str	交易日期
	GGTSS      float64 `json:"ggt_ss,omitempty"`      // float	港股通(上海)
	GGTSZ      float64 `json:"ggt_sz,omitempty"`      // float	港股通(深圳)
	HGT        float64 `json:"hgt,omitempty"`         // float	沪股通(百万元)
	SGT        float64 `json:"sgt,omitempty"`         // float	深股通(百万元)
	NorthMoney float64 `json:"north_money,omitempty"` // float	北向资金(百万元)
	SouthMoney float64 `json:"south_money,omitempty"` // float	南向资金(百万元)
}

func AssembleMoneyflowHSGTData(tsRsp *TushareResponse) []*MoneyflowHSGTData {
	tsData := []*MoneyflowHSGTData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(MoneyflowHSGTData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 获取沪股通、深股通、港股通每日资金流向数据,每次最多返回300条记录,总量不限制
func (ts *TuShare) MoneyflowHSGT(params MoneyflowHSGTRequest, items MoneyflowHSGTItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "moneyflow_hsgt",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type HSGTTop10Request struct {
	TsCode     string `json:"ts_code,omitempty"`     // str	N	股票代码(支持多个股票同时提取,逗号分隔)
	TradeDate  string `json:"trade_date,omitempty"`  // str	N	交易日期(YYYYMMDD)
	StartDate  string `json:"start_date,omitempty"`  // str	N	开始日期(YYYYMMDD,提取分钟数据请用2019-09-01 09:00:00这种格式)
	EndDate    string `json:"end_date,omitempty"`    // str	N	结束日期(YYYYMMDD,提取分钟数据请用2019-09-01 09:00:00这种格式)
	MarketType string `json:"market_type,omitempty"` // str	市场类型(1：沪市 3：深市)
}

type HSGTTop10Items struct {
	TsCode     bool `json:"ts_code,omitempty"`     // str	股票代码
	TradeDate  bool `json:"trade_date,omitempty"`  // str	交易日期
	Name       bool `json:"name,omitempty"`        // str	股票名称
	Close      bool `json:"close,omitempty"`       // float	收盘价
	Change     bool `json:"change,omitempty"`      // float	涨跌额
	Rank       bool `json:"rank,omitempty"`        // int	资金排名
	MarketType bool `json:"market_type,omitempty"` // str	市场类型(1：沪市 3：深市)
	Amount     bool `json:"amount,omitempty"`      // float	成交金额(元)
	NetAmount  bool `json:"net_amount,omitempty"`  // float	净成交金额(元)
	Buy        bool `json:"buy,omitempty"`         // float	买入金额(元)
	Sell       bool `json:"sell,omitempty"`        // float	卖出金额(元)
}

func (item HSGTTop10Items) All() HSGTTop10Items {
	item.TsCode = true
	item.TradeDate = true
	item.Name = true
	item.Close = true
	item.Change = true
	item.Rank = true
	item.MarketType = true
	item.Amount = true
	item.NetAmount = true
	item.Buy = true
	item.Sell = true
	return item
}

type HSGTTop10Data struct {
	TsCode     string  `json:"ts_code,omitempty"`     // str	股票代码
	TradeDate  string  `json:"trade_date,omitempty"`  // str	交易日期
	Name       string  `json:"name,omitempty"`        // str	股票名称
	Close      float64 `json:"close,omitempty"`       // float	收盘价
	Change     float64 `json:"change,omitempty"`      // float	涨跌额
	Rank       int64   `json:"rank,omitempty"`        // int	资金排名
	MarketType string  `json:"market_type,omitempty"` // str	市场类型(1：沪市 3：深市)
	Amount     float64 `json:"amount,omitempty"`      // float	成交金额(元)
	NetAmount  float64 `json:"net_amount,omitempty"`  // float	净成交金额(元)
	Buy        float64 `json:"buy,omitempty"`         // float	买入金额(元)
	Sell       float64 `json:"sell,omitempty"`        // float	卖出金额(元)
}

func AssembleHSGTTop10Data(tsRsp *TushareResponse) []*HSGTTop10Data {
	tsData := []*HSGTTop10Data{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(HSGTTop10Data)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 获取沪股通、深股通每日前十大成交详细数据
func (ts *TuShare) HSGTTop10(params HSGTTop10Request, items HSGTTop10Items) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "hsgt_top10",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type HKHoldRequest struct {
	TsCode    string `json:"ts_code,omitempty"`    // str	N	股票代码(支持多个股票同时提取,逗号分隔)
	TradeDate string `json:"trade_date,omitempty"` // str	N	交易日期(YYYYMMDD)
	StartDate string `json:"start_date,omitempty"` // str	N	开始日期(YYYYMMDD,提取分钟数据请用2019-09-01 09:00:00这种格式)
	EndDate   string `json:"end_date,omitempty"`   // str	N	结束日期(YYYYMMDD,提取分钟数据请用2019-09-01 09:00:00这种格式)
	Code      string `json:"code,omitempty"`       // str	N	交易所代码
	Exchange  string `json:"exchange,omitempty"`   // str	N	类型：SH沪股通(北向)SZ深股通(北向)HK港股通(南向持股)
}

type HKHoldItems struct {
	TsCode    bool `json:"ts_code,omitempty"`    // str	Y	TS股票代码
	TradeDate bool `json:"trade_date,omitempty"` // str	Y	交易日期
	Code      bool `json:"code,omitempty"`       // str	Y	原始代码
	Name      bool `json:"name,omitempty"`       // str	Y	股票名称
	Vol       bool `json:"vol,omitempty"`        // Y	持股数量(股)
	Ratio     bool `json:"ratio,omitempty"`      // float	Y	持股占比(%),占已发行股份百分比
	Exchange  bool `json:"exchange,omitempty"`   // str	Y	类型：SH沪股通SZ深股通HK港股通
}

func (item HKHoldItems) All() HKHoldItems {
	item.TsCode = true
	item.TradeDate = true
	item.Code = true
	item.Name = true
	item.Vol = true
	item.Ratio = true
	item.Exchange = true
	return item
}

type HKHoldData struct {
	TsCode    string  `json:"ts_code,omitempty"`    // str	Y	TS股票代码
	TradeDate string  `json:"trade_date,omitempty"` // str	Y	交易日期
	Code      string  `json:"code,omitempty"`       // str	Y	原始代码
	Name      string  `json:"name,omitempty"`       // str	Y	股票名称
	Vol       int64   `json:"vol,omitempty"`        // int	Y	持股数量(股)
	Ratio     float64 `json:"ratio,omitempty"`      // float	Y	持股占比(%),占已发行股份百分比
	Exchange  string  `json:"exchange,omitempty"`   // str	Y	类型：SH沪股通SZ深股通HK港股通
}

func AssembleHKHoldData(tsRsp *TushareResponse) []*HKHoldData {
	tsData := []*HKHoldData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(HKHoldData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 获取新股上市列表数据,单次最大2000条,总量不限制,用户需要至少120积分才可以调取,具体请参阅积分获取办法 https://tushare.pro/document/1?doc_id=13
func (ts *TuShare) HKHold(params HKHoldRequest, items HKHoldItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "hk_hold",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type BakDailyRequest struct {
	TsCode    string `json:"ts_code,omitempty"`    // str	N	股票代码(支持多个股票同时提取,逗号分隔)
	TradeDate string `json:"trade_date,omitempty"` // str	N	交易日期(YYYYMMDD)
	StartDate string `json:"start_date,omitempty"` // str	N	开始日期(YYYYMMDD,提取分钟数据请用2019-09-01 09:00:00这种格式)
	EndDate   string `json:"end_date,omitempty"`   // str	N	结束日期(YYYYMMDD,提取分钟数据请用2019-09-01 09:00:00这种格式)
	Offset    string `json:"offset,omitempty"`     // str	N	开始行数
	Limit     string `json:"limit,omitempty"`      // str	N	最大行数
}

type BakDailyItems struct {
	TsCode      bool `json:"ts_code,omitempty"`      // str	Y	股票代码
	TradeDate   bool `json:"trade_date,omitempty"`   // str	Y	交易日期
	Name        bool `json:"name,omitempty"`         // str	Y	股票名称
	PctChange   bool `json:"pct_change,omitempty"`   // float	Y	涨跌幅
	Close       bool `json:"close,omitempty"`        // float	Y	收盘价
	Change      bool `json:"change,omitempty"`       // float	Y	涨跌额
	Open        bool `json:"open,omitempty"`         // float	Y	开盘价
	High        bool `json:"high,omitempty"`         // float	Y	最高价
	Low         bool `json:"low,omitempty"`          // float	Y	最低价
	PreClose    bool `json:"pre_close,omitempty"`    // float	Y	昨收价
	VolRatio    bool `json:"vol_ratio,omitempty"`    // float	Y	量比
	TurnOver    bool `json:"turn_over,omitempty"`    // float	Y	换手率
	Swing       bool `json:"swing,omitempty"`        // float	Y	振幅
	Vol         bool `json:"vol,omitempty"`          // float	Y	成交量
	Amount      bool `json:"amount,omitempty"`       // float	Y	成交额
	Selling     bool `json:"selling,omitempty"`      // float	Y	内盘(主动卖,手)
	Buying      bool `json:"buying,omitempty"`       // float	Y	外盘(主动买, 手)
	TotalShare  bool `json:"total_share,omitempty"`  // float	Y	总股本(万)
	FloatShare  bool `json:"float_share,omitempty"`  // float	Y	流通股本(万)
	Pe          bool `json:"pe,omitempty"`           // float	Y	市盈(动)
	Industry    bool `json:"industry,omitempty"`     // str	Y	所属行业
	Area        bool `json:"area,omitempty"`         // str	Y	所属地域
	FloatMV     bool `json:"float_mv,omitempty"`     // float	Y	流通市值
	TotalMV     bool `json:"total_mv,omitempty"`     // float	Y	总市值
	AvgPrice    bool `json:"avg_price,omitempty"`    // float	Y	平均价
	Ctrength    bool `json:"strength,omitempty"`     // float	Y	强弱度(%)
	Activity    bool `json:"activity,omitempty"`     // float	Y	活跃度(%)
	AvgTurnover bool `json:"avg_turnover,omitempty"` // float	Y	笔换手
	Attack      bool `json:"attack,omitempty"`       // float	Y	攻击波(%)
	Interval3   bool `json:"interval_3,omitempty"`   // float	Y	近3月涨幅
	Interval6   bool `json:"interval_6,omitempty"`   // float	Y	近6月涨幅
}

func (item BakDailyItems) All() BakDailyItems {
	item.TsCode = true
	item.TradeDate = true
	item.Name = true
	item.PctChange = true
	item.Close = true
	item.Change = true
	item.Open = true
	item.High = true
	item.Low = true
	item.PreClose = true
	item.VolRatio = true
	item.TurnOver = true
	item.Vol = true
	item.Amount = true
	item.Selling = true
	item.Buying = true
	item.TotalShare = true
	item.FloatShare = true
	item.Pe = true
	item.Industry = true
	item.Area = true
	item.FloatMV = true
	item.TotalMV = true
	item.AvgPrice = true
	item.Ctrength = true
	item.Activity = true
	item.AvgTurnover = true
	item.Attack = true
	item.Interval3 = true
	item.Interval6 = true
	return item
}

type BakDailyData struct {
	TsCode      string  `json:"ts_code,omitempty"`      // str	Y	股票代码
	TradeDate   string  `json:"trade_date,omitempty"`   // str	Y	交易日期
	Name        string  `json:"name,omitempty"`         // str	Y	股票名称
	PctChange   float64 `json:"pct_change,omitempty"`   // float	Y	涨跌幅
	Close       float64 `json:"close,omitempty"`        // float	Y	收盘价
	Change      float64 `json:"change,omitempty"`       // float	Y	涨跌额
	Open        float64 `json:"open,omitempty"`         // float	Y	开盘价
	High        float64 `json:"high,omitempty"`         // float	Y	最高价
	Low         float64 `json:"low,omitempty"`          // float	Y	最低价
	PreClose    float64 `json:"pre_close,omitempty"`    // float	Y	昨收价
	VolRatio    float64 `json:"vol_ratio,omitempty"`    // float	Y	量比
	TurnOver    float64 `json:"turn_over,omitempty"`    // float	Y	换手率
	Swing       float64 `json:"swing,omitempty"`        // float	Y	振幅
	Vol         float64 `json:"vol,omitempty"`          // float	Y	成交量
	Amount      float64 `json:"amount,omitempty"`       // float	Y	成交额
	Selling     float64 `json:"selling,omitempty"`      // float	Y	内盘(主动卖,手)
	Buying      float64 `json:"buying,omitempty"`       // float	Y	外盘(主动买, 手)
	TotalShare  float64 `json:"total_share,omitempty"`  // float	Y	总股本(万)
	FloatShare  float64 `json:"float_share,omitempty"`  // float	Y	流通股本(万)
	Pe          float64 `json:"pe,omitempty"`           // float	Y	市盈(动)
	Industry    string  `json:"industry,omitempty"`     // str	Y	所属行业
	Area        string  `json:"area,omitempty"`         // str	Y	所属地域
	FloatMV     float64 `json:"float_mv,omitempty"`     // float	Y	流通市值
	TotalMV     float64 `json:"total_mv,omitempty"`     // float	Y	总市值
	AvgPrice    float64 `json:"avg_price,omitempty"`    // float	Y	平均价
	Ctrength    float64 `json:"strength,omitempty"`     // float	Y	强弱度(%)
	Activity    float64 `json:"activity,omitempty"`     // float	Y	活跃度(%)
	AvgTurnover float64 `json:"avg_turnover,omitempty"` // float	Y	笔换手
	Attack      float64 `json:"attack,omitempty"`       // float	Y	攻击波(%)
	Interval3   float64 `json:"interval_3,omitempty"`   // float	Y	近3月涨幅
	Interval6   float64 `json:"interval_6,omitempty"`   // float	Y	近6月涨幅
}

func AssembleBakDailyData(tsRsp *TushareResponse) []*BakDailyData {
	tsData := []*BakDailyData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(BakDailyData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 获取备用行情,包括特定的行情指标,单次最大5000行数据,可以根据日期参数循环获取
func (ts *TuShare) BakDaily(params BakDailyRequest, items BakDailyItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "bak_daily",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}
