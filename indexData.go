package tushare

import (
	"encoding/json"
	"net/http"
)

type IndexBasicRequest struct {
	TsCode    string `json:"ts_code,omitempty"`   // str	N	指数代码
	Name      string `json:"name,omitempty"`      // str	N	指数简称
	Market    string `json:"market,omitempty"`    // str	N	交易所或服务商(默认SSE)
	Publisher string `json:"publisher,omitempty"` // str	N	发布商
	Category  string `json:"category,omitempty"`  // str	N	指数类别
}

type IndexBasicItems struct {
	TsCode     bool `json:"ts_code,omitempty"`     // str TS代码
	Name       bool `json:"name,omitempty"`        // str	简称
	Fullname   bool `json:"fullname,omitempty"`    // str	指数全称
	Market     bool `json:"market,omitempty"`      // str	市场
	Publisher  bool `json:"publisher,omitempty"`   // str	发布方
	IndexType  bool `json:"index_type,omitempty"`  // str	指数风格
	Category   bool `json:"category,omitempty"`    // str	指数类别
	BaseDate   bool `json:"base_date,omitempty"`   // str	基期
	BasePoint  bool `json:"base_point,omitempty"`  // float	基点
	ListDate   bool `json:"list_date,omitempty"`   // str	发布日期
	WeightRule bool `json:"weight_rule,omitempty"` // str	加权方式
	Desc       bool `json:"desc,omitempty"`        // str	描述
	ExpDate    bool `json:"exp_date,omitempty"`    // str	终止日期
}

func (item IndexBasicItems) All() IndexBasicItems {
	item.TsCode = true
	item.Name = true
	item.Fullname = true
	item.Market = true
	item.Publisher = true
	item.IndexType = true
	item.Category = true
	item.BaseDate = true
	item.BasePoint = true
	item.ListDate = true
	item.WeightRule = true
	item.Desc = true
	item.ExpDate = true
	return item
}

type IndexBasicData struct {
	TsCode     string  `json:"ts_code,omitempty"`     // str TS代码
	Name       string  `json:"name,omitempty"`        // str	简称
	Fullname   string  `json:"fullname,omitempty"`    // str	指数全称
	Market     string  `json:"market,omitempty"`      // str	市场
	Publisher  string  `json:"publisher,omitempty"`   // str	发布方
	IndexType  string  `json:"index_type,omitempty"`  // str	指数风格
	Category   string  `json:"category,omitempty"`    // str	指数类别
	BaseDate   string  `json:"base_date,omitempty"`   // str	基期
	BasePoint  float64 `json:"base_point,omitempty"`  // float	基点
	ListDate   string  `json:"list_date,omitempty"`   // str	发布日期
	WeightRule string  `json:"weight_rule,omitempty"` // str	加权方式
	Desc       string  `json:"desc,omitempty"`        // str	描述
	ExpDate    string  `json:"exp_date,omitempty"`    // str	终止日期
}

func AssembleIndexBasicData(tsRsp *TushareResponse) []*IndexBasicData {
	tsData := []*IndexBasicData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(IndexBasicData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 获取指数基础信息
func (ts *TuShare) IndexBasic(params IndexBasicRequest, items IndexBasicItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "index_basic",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

// 获取指数每日行情,还可以通过bar接口获取.由于服务器压力,目前规则是单次调取最多取8000行记录,可以设置start和end日期补全.指数行情也可以通过通用行情接口获取数据．常规指数需累积200积分可低频调取,5000积分以上频次相对较高,具体请参阅积分获取办法 https://tushare.pro/document/1?doc_id=13
func (ts *TuShare) IndexDaily(params QuotationRequest, items QuotationItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "index_daily",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

// 获取指数周线行情,单次最大1000行记录,可分批获取,总量不限制,用户需要至少600积分才可以调取,积分越多频次越高,具体请参阅积分获取办法 https://tushare.pro/document/1?doc_id=13
func (ts *TuShare) IndexWeekly(params QuotationRequest, items QuotationItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "index_weekly",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

// 获取指数月线行情,单次最大1000行记录,可分批获取,总量不限制,用户需要至少600积分才可以调取,积分越多频次越高,具体请参阅积分获取办法 https://tushare.pro/document/1?doc_id=13
func (ts *TuShare) IndexMonthly(params QuotationRequest, items QuotationItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "index_monthly",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type IndexWeightRequest struct {
	IndexCode string `json:"index_code,omitempty"` // Y	指数代码 (二选一)
	TradeDate string `json:"trade_date,omitempty"` // str	Y	交易日期 (二选一)
	StartDate string `json:"start_date,omitempty"` // str	N	开始日期(YYYYMMDD)
	EndDate   string `json:"end_date,omitempty"`   // str	N	结束日期(YYYYMMDD)
}

type IndexWeightItems struct {
	IndexCode bool `json:"index_code,omitempty"` // str	指数代码
	ConCode   bool `json:"con_code,omitempty"`   // str	成分代码
	TradeDate bool `json:"trade_date,omitempty"` // str	交易日期
	Weight    bool `json:"weight,omitempty"`     // float	权重
}

func (item IndexWeightItems) All() IndexWeightItems {
	item.IndexCode = true
	item.ConCode = true
	item.TradeDate = true
	item.Weight = true
	return item
}

type IndexWeightData struct {
	IndexCode string  `json:"index_code,omitempty"` // str	指数代码
	ConCode   string  `json:"con_code,omitempty"`   // str	成分代码
	TradeDate string  `json:"trade_date,omitempty"` // str	交易日期
	Weight    float64 `json:"weight,omitempty"`     // float	权重
}

func AssembleIndexWeightData(tsRsp *TushareResponse) []*IndexWeightData {
	tsData := []*IndexWeightData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(IndexWeightData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 来源于指数公司网站公开数据,获取各类指数成分和权重,月度数据 ,如需日度指数成分和权重,请联系 waditu@163.com,用户需要至少400积分才可以调取,具体请参阅积分获取办法 https://tushare.pro/document/1?doc_id=13
func (ts *TuShare) IndexWeight(params IndexWeightRequest, items IndexWeightItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "index_weight",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type IndexDailyBasicItems struct {
	TsCode        bool `json:"ts_code,omitempty"`         // str TS代码
	TradeDate     bool `json:"trade_date,omitempty"`      // str	Y	交易日期
	TotalMV       bool `json:"total_mv,omitempty"`        // float	Y	当日总市值(元)
	FloatMV       bool `json:"float_mv,omitempty"`        // float	Y	当日流通市值(元)
	TotalShare    bool `json:"total_share,omitempty"`     // float	Y	当日总股本(股)
	FloatShare    bool `json:"float_share,omitempty"`     // float	Y	当日流通股本(股)
	FreeShare     bool `json:"free_share,omitempty"`      // float	Y	当日自由流通股本(股)
	TurnoverRate  bool `json:"turnover_rate,omitempty"`   // float	Y	换手率
	TurnoverRateF bool `json:"turnover_rate_f,omitempty"` // float	Y	换手率(基于自由流通股本)
	Pe            bool `json:"pe,omitempty"`              // float	Y	市盈率
	PeTTM         bool `json:"pe_ttm,omitempty"`          // float	Y	市盈率TTM
	Pb            bool `json:"pb,omitempty"`              // float	Y	市净率
}

func (item IndexDailyBasicItems) All() IndexDailyBasicItems {
	item.TsCode = true
	item.TradeDate = true
	item.TotalMV = true
	item.FloatMV = true
	item.TotalShare = true
	item.FloatShare = true
	item.FreeShare = true
	item.TurnoverRate = true
	item.TurnoverRateF = true
	item.Pe = true
	item.PeTTM = true
	item.Pb = true
	return item
}

type IndexDailyBasicData struct {
	TsCode        string  `json:"ts_code,omitempty"`         // str TS代码
	TradeDate     string  `json:"trade_date,omitempty"`      // str	Y	交易日期
	TotalMV       float64 `json:"total_mv,omitempty"`        // float	Y	当日总市值(元)
	FloatMV       float64 `json:"float_mv,omitempty"`        // float	Y	当日流通市值(元)
	TotalShare    float64 `json:"total_share,omitempty"`     // float	Y	当日总股本(股)
	FloatShare    float64 `json:"float_share,omitempty"`     // float	Y	当日流通股本(股)
	FreeShare     float64 `json:"free_share,omitempty"`      // float	Y	当日自由流通股本(股)
	TurnoverRate  float64 `json:"turnover_rate,omitempty"`   // float	Y	换手率
	TurnoverRateF float64 `json:"turnover_rate_f,omitempty"` // float	Y	换手率(基于自由流通股本)
	Pe            float64 `json:"pe,omitempty"`              // float	Y	市盈率
	PeTTM         float64 `json:"pe_ttm,omitempty"`          // float	Y	市盈率TTM
	Pb            float64 `json:"pb,omitempty"`              // float	Y	市净率
}

func AssembleIndexDailyBasicData(tsRsp *TushareResponse) []*IndexDailyBasicData {
	tsData := []*IndexDailyBasicData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(IndexDailyBasicData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 大盘指数每日指标,从2004年1月开始只提供上证综指,深证成指,上证50,中证500,中小板指,创业板指的每日指标数据,用户需要至少400积分才可以调取,具体请参阅积分获取办法 https://tushare.pro/document/1?doc_id=13
func (ts *TuShare) IndexDailyBasic(params QuotationRequest, items IndexDailyBasicItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "index_dailybasic",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type DailyInfoRequest struct {
	TsCode    string `json:"ts_code,omitempty"`    // str	N	板块代码(请参阅下方列表)
	TradeDate string `json:"trade_date,omitempty"` // str	N	交易日期(YYYYMMDD格式,下同)
	Exchange  string `json:"exchange,omitempty"`   // str	N	股票市场(SH上交所 SZ深交所)
	StartDate string `json:"start_date,omitempty"` // str	N	开始日期
	EndDate   string `json:"end_date,omitempty"`   // str	N	结束日期
	Fields    string `json:"fields,omitempty"`     // str	N	指定提取字段
}

type DailyInfoItems struct {
	TradeDate  bool `json:"trade_date,omitempty"`  // str	Y	交易日期
	TsCode     bool `json:"ts_code,omitempty"`     // str	Y	市场代码
	TsName     bool `json:"ts_name,omitempty"`     // str	Y	市场名称
	ComCount   bool `json:"com_count,omitempty"`   // int	Y	挂牌数
	TotalShare bool `json:"total_share,omitempty"` // float	Y	总股本(亿股)
	FloatShare bool `json:"float_share,omitempty"` // float	Y	流通股本(亿股)
	TotalMV    bool `json:"total_mv,omitempty"`    // float	Y	总市值(亿元)
	FloatMV    bool `json:"float_mv,omitempty"`    // float	Y	流通市值(亿元)
	Amount     bool `json:"amount,omitempty"`      // float	Y	交易金额(亿元)
	Vol        bool `json:"vol,omitempty"`         // float	Y	成交量(亿股)
	TransCount bool `json:"trans_count,omitempty"` // int	Y	成交笔数(万笔)
	Pe         bool `json:"pe,omitempty"`          // float	Y	平均市盈率
	Tr         bool `json:"tr,omitempty"`          // float	Y	换手率(％),注：深交所暂无此列
	Exchange   bool `json:"exchange,omitempty"`    // str	Y	交易所(SH上交所 SZ深交所)
}

func (item DailyInfoItems) All() DailyInfoItems {
	item.TradeDate = true
	item.TsCode = true
	item.TsName = true
	item.ComCount = true
	item.TotalShare = true
	item.FloatShare = true
	item.TotalMV = true
	item.FloatMV = true
	item.Amount = true
	item.Vol = true
	item.TransCount = true
	item.Pe = true
	item.Tr = true
	item.Exchange = true
	return item
}

// 板块代码(TS_CODE)	板块名称(TS_NAME)	数据开始日期
// SZ_MARKET	深圳市场	20041231
// SZ_MAIN	深圳主板	20081231
// SZ_A	深圳A股	20080103
// SZ_B	深圳B股	20080103
// SZ_GEM	创业板	20091030
// SZ_SME	中小企业板	20040602
// SZ_FUND	深圳基金市场	20080103
// SZ_FUND_ETF	深圳基金ETF	20080103
// SZ_FUND_LOF	深圳基金LOF	20080103
// SZ_FUND_CEF	深圳封闭基金	20080103
// SZ_FUND_SF	深圳分级基金	20080103
// SZ_BOND	深圳债券	20080103
// SZ_BOND_CN	深圳债券现券	20080103
// SZ_BOND_REP	深圳债券回购	20080103
// SZ_BOND_ABS	深圳债券ABS	20080103
// SZ_BOND_GOV	深圳国债	20080103
// SZ_BOND_ENT	深圳企业债	20080103
// SZ_BOND_COR	深圳公司债	20080103
// SZ_BOND_CB	深圳可转债	20080103
// SZ_WR	深圳权证	20080103
// ----	----	---
// SH_MARKET	上海市场	20190102
// SH_A	上海A股	19910102
// SH_B	上海B股	19920221
// SH_STAR	科创板	20190722
// SH_REP	股票回购	20190102
// SH_FUND	上海基金市场	19901219
// SH_FUND_ETF	上海基金ETF	19901219
// SH_FUND_LOF	上海基金LOF	19901219
// SH_FUND_REP	上海基金回购	19901219
// SH_FUND_CEF	上海封闭式基金	19901219
// SH_FUND_METF	上海交易型货币基金	19901219
type DailyInfoData struct {
	TradeDate  string  `json:"trade_date,omitempty"`  // str	Y	交易日期
	TsCode     string  `json:"ts_code,omitempty"`     // str	Y	市场代码
	TsName     string  `json:"ts_name,omitempty"`     // str	Y	市场名称
	ComCount   int64   `json:"com_count,omitempty"`   // int	Y	挂牌数
	TotalShare float64 `json:"total_share,omitempty"` // float	Y	总股本(亿股)
	FloatShare float64 `json:"float_share,omitempty"` // float	Y	流通股本(亿股)
	TotalMV    float64 `json:"total_mv,omitempty"`    // float	Y	总市值(亿元)
	FloatMV    float64 `json:"float_mv,omitempty"`    // float	Y	流通市值(亿元)
	Amount     float64 `json:"amount,omitempty"`      // float	Y	交易金额(亿元)
	Vol        float64 `json:"vol,omitempty"`         // float	Y	成交量(亿股)
	TransCount int64   `json:"trans_count,omitempty"` // int	Y	成交笔数(万笔)
	Pe         float64 `json:"pe,omitempty"`          // float	Y	平均市盈率
	Tr         float64 `json:"tr,omitempty"`          // float	Y	换手率(％),注：深交所暂无此列
	Exchange   string  `json:"exchange,omitempty"`    // str	Y	交易所(SH上交所 SZ深交所)
}

func AssembleDailyInfoData(tsRsp *TushareResponse) []*DailyInfoData {
	tsData := []*DailyInfoData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(DailyInfoData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 获取交易所股票交易统计,包括各板块明细,单次最大4000条,可循环获取,总量不限制,用户需要至少600积分才可以调取,频次有限制,积分越高每分钟调取频次越高,5000积分以上频次相对较高,具体请参阅积分获取办法 https://tushare.pro/document/1?doc_id=13
func (ts *TuShare) DailyInfo(params DailyInfoRequest, items DailyInfoItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "daily_info",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type THXIndexRequest struct {
	TsCode   string `json:"ts_code,omitempty"`  // str	N	指数代码
	Exchange string `json:"exchange,omitempty"` // str	N	市场类型A-a股 HK-港股 US-美股
	Type     string `json:"type,omitempty"`     // str	N	指数类型 N-板块指数 S-同花顺特色指数
}

type THXIndexItems struct {
	TsCode   bool `json:"ts_code,omitempty"`   // str	Y	代码
	Name     bool `json:"name,omitempty"`      // str	Y	名称
	Count    bool `json:"count,omitempty"`     // int	Y	成分个数
	Exchange bool `json:"exchange,omitempty"`  // str	Y	交易所
	ListDate bool `json:"list_date,omitempty"` // str	Y	上市日期
	Type     bool `json:"type,omitempty"`      // str	Y	N概念指数S特色指数
}

func (item THXIndexItems) All() THXIndexItems {
	item.TsCode = true
	item.Name = true
	item.Count = true
	item.Exchange = true
	item.ListDate = true
	item.Type = true
	return item
}

type THXIndexData struct {
	TsCode   string `json:"ts_code,omitempty"`   // str	Y	代码
	Name     string `json:"name,omitempty"`      // str	Y	名称
	Count    int64  `json:"count,omitempty"`     // int	Y	成分个数
	Exchange string `json:"exchange,omitempty"`  // str	Y	交易所
	ListDate string `json:"list_date,omitempty"` // str	Y	上市日期
	Type     string `json:"type,omitempty"`      // str	Y	N概念指数S特色指数
}

func AssembleTHXIndexData(tsRsp *TushareResponse) []*THXIndexData {
	tsData := []*THXIndexData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(THXIndexData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 获取同花顺板块指数,数据版权归属同花顺,如做商业用途,请主动联系同花顺,如需帮助请联系微信migedata,用户需要至少600积分才可以调取,单次最大5000,一次可提取全部数据,请勿循环提取,具体请参阅积分获取办法 https://tushare.pro/document/1?doc_id=13
func (ts *TuShare) THXIndex(params THXIndexRequest, items THXIndexItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "ths_index",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type THSDailyItems struct {
	TsCode       bool `json:"ts_code,omitempty"`       // str	Y	TS指数代码
	TradeDate    bool `json:"trade_date,omitempty"`    // str	Y	交易日
	Close        bool `json:"close,omitempty"`         // float	Y	收盘点位
	Open         bool `json:"open,omitempty"`          // float	Y	开盘点位
	High         bool `json:"high,omitempty"`          // float	Y	最高点位
	Low          bool `json:"low,omitempty"`           // float	Y	最低点位
	PreClose     bool `json:"pre_close,omitempty"`     // float	Y	昨日收盘点
	AvgPrice     bool `json:"avg_price,omitempty"`     // float	Y	平均价
	Change       bool `json:"change,omitempty"`        // float	Y	涨跌点位
	PctChange    bool `json:"pct_change,omitempty"`    // float	Y	涨跌幅
	Vol          bool `json:"vol,omitempty"`           // float	Y	成交量
	TurnoverRate bool `json:"turnover_rate,omitempty"` // float	Y	换手率
	TotalMV      bool `json:"total_mv,omitempty"`      // float	N	总市值
	FloatMV      bool `json:"float_mv,omitempty"`      // float	N	流通市值
}

func (item THSDailyItems) All() THSDailyItems {
	item.TsCode = true
	item.TradeDate = true
	item.Close = true
	item.Open = true
	item.High = true
	item.Low = true
	item.PreClose = true
	item.AvgPrice = true
	item.Change = true
	item.PctChange = true
	item.Vol = true
	item.TurnoverRate = true
	item.TotalMV = true
	item.FloatMV = true
	return item
}

type THSDailyData struct {
	TsCode       string  `json:"ts_code,omitempty"`       // str	Y	TS指数代码
	TradeDate    string  `json:"trade_date,omitempty"`    // str	Y	交易日
	Close        float64 `json:"close,omitempty"`         // float	Y	收盘点位
	Open         float64 `json:"open,omitempty"`          // float	Y	开盘点位
	High         float64 `json:"high,omitempty"`          // float	Y	最高点位
	Low          float64 `json:"low,omitempty"`           // float	Y	最低点位
	PreClose     float64 `json:"pre_close,omitempty"`     // float	Y	昨日收盘点
	AvgPrice     float64 `json:"avg_price,omitempty"`     // float	Y	平均价
	Change       float64 `json:"change,omitempty"`        // float	Y	涨跌点位
	PctChange    float64 `json:"pct_change,omitempty"`    // float	Y	涨跌幅
	Vol          float64 `json:"vol,omitempty"`           // float	Y	成交量
	TurnoverRate float64 `json:"turnover_rate,omitempty"` // float	Y	换手率
	TotalMV      float64 `json:"total_mv,omitempty"`      // float	N	总市值
	FloatMV      float64 `json:"float_mv,omitempty"`      // float	N	流通市值
}

func AssembleTHSDailyData(tsRsp *TushareResponse) []*THSDailyData {
	tsData := []*THSDailyData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(THSDailyData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 获取同花顺板块指数行情,数据版权归属同花顺,如做商业用途,请主动联系同花顺,如需帮助请联系微信migedata,单次最大3000行数据,可根据指数代码、日期参数循环提取,
func (ts *TuShare) THSDaily(params QuotationRequest, items THSDailyItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "ths_daily",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type THSMemberRequest struct {
	TsCode bool `json:"ts_code,omitempty"` // str	Y	TS股票代码
}

type THSMemberItems struct {
	TsCode  bool `json:"ts_code,omitempty"`  // str	Y	指数代码
	Code    bool `json:"code,omitempty"`     // str	Y	股票代码
	Name    bool `json:"name,omitempty"`     // str	Y	股票名称
	Weight  bool `json:"weight,omitempty"`   // float	N	权重
	InDate  bool `json:"in_date,omitempty"`  // str	N	纳入日期
	OutDate bool `json:"out_date,omitempty"` // str	N	剔除日期
	IsNew   bool `json:"is_new,omitempty"`   // str	N	是否最新Y是N否
}

func (item THSMemberItems) All() THSMemberItems {
	item.TsCode = true
	item.Code = true
	item.Name = true
	item.Weight = true
	item.InDate = true
	item.OutDate = true
	item.IsNew = true
	return item
}

type THSMemberData struct {
	TsCode  string  `json:"ts_code,omitempty"`  // str	Y	指数代码
	Code    string  `json:"code,omitempty"`     // str	Y	股票代码
	Name    string  `json:"name,omitempty"`     // str	Y	股票名称
	Weight  float64 `json:"weight,omitempty"`   // float	N	权重
	InDate  string  `json:"in_date,omitempty"`  // str	N	纳入日期
	OutDate string  `json:"out_date,omitempty"` // str	N	剔除日期
	IsNew   string  `json:"is_new,omitempty"`   // str	N	是否最新Y是N否
}

func AssembleTHSMemberData(tsRsp *TushareResponse) []*THSMemberData {
	tsData := []*THSMemberData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(THSMemberData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 获取同花顺概念板块成分列表,数据版权归属同花顺,如做商业用途,请主动联系同花顺,总量不限制,用户需要至少5000积分才可以调取,可按概念板块代码循环提取所有成分,具体请参阅积分获取办法 https://tushare.pro/document/1?doc_id=13
func (ts *TuShare) THSMember(params THSMemberRequest, items THSMemberItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "ths_member",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type IndexGlobalItems struct {
	TsCode    bool `json:"ts_code,omitempty"`    // str	Y	TS指数代码
	TradeDate bool `json:"trade_date,omitempty"` // str	Y	交易日
	Open      bool `json:"open,omitempty"`       // float	Y	开盘点位
	Close     bool `json:"close,omitempty"`      // float	Y	收盘点位
	High      bool `json:"high,omitempty"`       // float	Y	最高点位
	Low       bool `json:"low,omitempty"`        // float	Y	最低点位
	PreClose  bool `json:"pre_close,omitempty"`  // float	Y	昨日收盘点
	Change    bool `json:"change,omitempty"`     // float	Y	涨跌点位
	PctChg    bool `json:"pct_chg,omitempty"`    // float	Y	涨跌幅
	Swing     bool `json:"swing,omitempty"`      // float	Y	振幅
	Vol       bool `json:"vol,omitempty"`        // float	Y	成交量 (大部分无此项数据)
	Amount    bool `json:"amount,omitempty"`     // float	N	成交额 (大部分无此项数据)
}

func (item IndexGlobalItems) All() IndexGlobalItems {
	item.TsCode = true
	item.TradeDate = true
	item.Open = true
	item.Close = true
	item.High = true
	item.Low = true
	item.PreClose = true
	item.Change = true
	item.PctChg = true
	item.Swing = true
	item.Vol = true
	item.Amount = true
	return item
}

type IndexGlobalData struct {
	TsCode    string  `json:"ts_code,omitempty"`    // str	Y	TS指数代码
	TradeDate string  `json:"trade_date,omitempty"` // str	Y	交易日
	Open      float64 `json:"open,omitempty"`       // float	Y	开盘点位
	Close     float64 `json:"close,omitempty"`      // float	Y	收盘点位
	High      float64 `json:"high,omitempty"`       // float	Y	最高点位
	Low       float64 `json:"low,omitempty"`        // float	Y	最低点位
	PreClose  float64 `json:"pre_close,omitempty"`  // float	Y	昨日收盘点
	Change    float64 `json:"change,omitempty"`     // float	Y	涨跌点位
	PctChg    float64 `json:"pct_chg,omitempty"`    // float	Y	涨跌幅
	Swing     float64 `json:"swing,omitempty"`      // float	Y	振幅
	Vol       float64 `json:"vol,omitempty"`        // float	Y	成交量 (大部分无此项数据)
	Amount    float64 `json:"amount,omitempty"`     // float	N	成交额 (大部分无此项数据)
}

func AssembleIndexGlobalData(tsRsp *TushareResponse) []*IndexGlobalData {
	tsData := []*IndexGlobalData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(IndexGlobalData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// TS指数代码	指数名称
// XIN9	富时中国A50指数 (富时A50)
// HSI	恒生指数
// DJI	道琼斯工业指数
// SPX	标普500指数
// IXIC	纳斯达克指数
// FTSE	富时100指数
// FCHI	法国CAC40指数
// GDAXI	德国DAX指数
// N225	日经225指数
// KS11	韩国综合指数
// AS51	澳大利亚标普200指数
// SENSEX	印度孟买SENSEX指数
// IBOVESPA	巴西IBOVESPA指数
// RTS	俄罗斯RTS指数
// TWII	台湾加权指数
// CKLSE	马来西亚指数
// SPTSX	加拿大S&P/TSX指数
// CSX5P	STOXX欧洲50指数
// 获取国际主要指数日线行情,单次最大2000条,可循环获取,总量不限制,用户需要至少2000积分才可以调取,具体请参阅积分获取办法 https://tushare.pro/document/1?doc_id=13
func (ts *TuShare) IndexGlobal(params QuotationRequest, items IndexGlobalItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "index_global",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type IndexClassifyRequest struct {
	IndexCode string `json:"index_code,omitempty"` // str	N	指数代码
	Level     string `json:"level,omitempty"`      // str	N	行业分级(L1/L2/L3)
	Src       string `json:"src,omitempty"`        // str	N	指数来源(SW申万)
}

type IndexClassifyItems struct {
	IndexCode    bool `json:"index_code,omitempty"`    // str	Y	指数代码
	IndustryName bool `json:"industry_name,omitempty"` // str	Y	行业名称
	Level        bool `json:"level,omitempty"`         // str	Y	行业名称
	IndustryCode bool `json:"industry_code,omitempty"` // str	N	行业代码
	Src          bool `json:"src,omitempty"`           // str	N	行业分类(SW申万)
}

func (item IndexClassifyItems) All() IndexClassifyItems {
	item.IndexCode = true
	item.IndustryName = true
	item.Level = true
	item.IndustryCode = true
	item.Src = true
	return item
}

type IndexClassifyData struct {
	IndexCode    string `json:"index_code,omitempty"`    // str	Y	指数代码
	IndustryName string `json:"industry_name,omitempty"` // str	Y	行业名称
	Level        string `json:"level,omitempty"`         // str	Y	行业名称
	IndustryCode string `json:"industry_code,omitempty"` // str	N	行业代码
	Src          string `json:"src,omitempty"`           // str	N	行业分类(SW申万)
}

func AssembleIndexClassifyData(tsRsp *TushareResponse) []*IndexClassifyData {
	tsData := []*IndexClassifyData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(IndexClassifyData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 获取申万行业分类,包括申万28个一级分类,104个二级分类,227个三级分类的列表信息,用户需要至少2000积分才可以调取,具体请参阅积分获取办法 https://tushare.pro/document/1?doc_id=13
func (ts *TuShare) IndexClassify(params IndexClassifyRequest, items IndexClassifyItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "index_classify",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type IndexMemberRequest struct {
	IndexCode string `json:"index_code,omitempty"` // str	N	指数代码
	TsCode    string `json:"ts_code,omitempty"`    // str	N	股票代码
	IsNew     string `json:"is_new,omitempty"`     // str	N	是否最新(默认为“Y是”)
}

type IndexMemberItems struct {
	IndexCode bool `json:"index_code,omitempty"` // str	Y	指数代码
	IndexName bool `json:"index_name,omitempty"` // str	N	指数名称
	ConCode   bool `json:"con_code,omitempty"`   // str	Y	成分股票代码
	ConName   bool `json:"con_name,omitempty"`   // str	Y	成分股票名称
	InDate    bool `json:"in_date,omitempty"`    // str	Y	纳入日期
	OutDate   bool `json:"out_date,omitempty"`   // str	Y	剔除日期
	IsNew     bool `json:"is_new,omitempty"`     // str	N	是否最新Y是N否
}

func (item IndexMemberItems) All() IndexMemberItems {
	item.IndexCode = true
	item.IndexName = true
	item.ConCode = true
	item.ConName = true
	item.InDate = true
	item.OutDate = true
	item.IsNew = true
	return item
}

type IndexMemberData struct {
	IndexCode string `json:"index_code,omitempty"` // str	Y	指数代码
	IndexName string `json:"index_name,omitempty"` // str	N	指数名称
	ConCode   string `json:"con_code,omitempty"`   // str	Y	成分股票代码
	ConName   string `json:"con_name,omitempty"`   // str	Y	成分股票名称
	InDate    string `json:"in_date,omitempty"`    // str	Y	纳入日期
	OutDate   string `json:"out_date,omitempty"`   // str	Y	剔除日期
	IsNew     string `json:"is_new,omitempty"`     // str	N	是否最新Y是N否
}

func AssembleIndexMemberData(tsRsp *TushareResponse) []*IndexMemberData {
	tsData := []*IndexMemberData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(IndexMemberData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 申万行业成分,单次最大2000行,总量不限制,用户需要至少2000积分才可以调取,具体请参阅积分获取办法 https://tushare.pro/document/1?doc_id=13
func (ts *TuShare) IndexMember(params IndexMemberRequest, items IndexMemberItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "index_member",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}
