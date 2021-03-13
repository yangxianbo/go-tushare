package tushare

import (
	"encoding/json"
	"net/http"
)

type GGTTop10Request struct {
	TsCode     string `json:"ts_code,omitempty"`     // str	N	股票代码（二选一）
	TradeDate  string `json:"trade_date,omitempty"`  // str	N	交易日期（二选一）
	StartDate  string `json:"start_date,omitempty"`  // str	N	开始日期
	EndDate    string `json:"end_date,omitempty"`    // str	N	结束日期
	MarketType string `json:"market_type,omitempty"` // str	N	市场类型 2：港股通（沪） 4：港股通（深）
}

type GGTTop10Items struct {
	TradeDate   bool `json:"trade_date,omitempty"`    // str	交易日期
	TsCode      bool `json:"ts_code,omitempty"`       // str	股票代码
	Name        bool `json:"name,omitempty"`          // str	股票名称
	Close       bool `json:"close,omitempty"`         // float	收盘价
	PChange     bool `json:"p_change,omitempty"`      // float	涨跌幅
	Rank        bool `json:"rank,omitempty"`          // str	资金排名
	MarketType  bool `json:"market_type,omitempty"`   // str	市场类型 2：港股通（沪） 4：港股通（深）
	Amount      bool `json:"amount,omitempty"`        // float	累计成交金额（元）
	NetAmount   bool `json:"net_amount,omitempty"`    // float	净买入金额（元）
	ShAmount    bool `json:"sh_amount,omitempty"`     // float	沪市成交金额（元）
	ShNetAmount bool `json:"sh_net_amount,omitempty"` // float	沪市净买入金额（元）
	ShBuy       bool `json:"sh_buy,omitempty"`        // float	沪市买入金额（元）
	ShSell      bool `json:"sh_sell,omitempty"`       // float	沪市卖出金额
	SzAmount    bool `json:"sz_amount,omitempty"`     // float	深市成交金额（元）
	SzNetAmount bool `json:"sz_net_amount,omitempty"` // float	深市净买入金额（元）
	SzBuy       bool `json:"sz_buy,omitempty"`        // float	深市买入金额（元）
	SzSell      bool `json:"sz_sell,omitempty"`       // float	深市卖出金额（元）
}

func (item GGTTop10Items) All() GGTTop10Items {
	item.TradeDate = true
	item.TsCode = true
	item.Name = true
	item.Close = true
	item.PChange = true
	item.Rank = true
	item.MarketType = true
	item.Amount = true
	item.NetAmount = true
	item.ShAmount = true
	item.ShNetAmount = true
	item.ShBuy = true
	item.ShSell = true
	item.SzAmount = true
	item.SzNetAmount = true
	item.SzBuy = true
	item.SzSell = true
	return item
}

type GGTTop10Data struct {
	TradeDate   string  `json:"trade_date,omitempty"`    // str	交易日期
	TsCode      string  `json:"ts_code,omitempty"`       // str	股票代码
	Name        string  `json:"name,omitempty"`          // str	股票名称
	Close       float64 `json:"close,omitempty"`         // float	收盘价
	PChange     float64 `json:"p_change,omitempty"`      // float	涨跌幅
	Rank        string  `json:"rank,omitempty"`          // str	资金排名
	MarketType  string  `json:"market_type,omitempty"`   // str	市场类型 2：港股通（沪） 4：港股通（深）
	Amount      float64 `json:"amount,omitempty"`        // float	累计成交金额（元）
	NetAmount   float64 `json:"net_amount,omitempty"`    // float	净买入金额（元）
	ShAmount    float64 `json:"sh_amount,omitempty"`     // float	沪市成交金额（元）
	ShNetAmount float64 `json:"sh_net_amount,omitempty"` // float	沪市净买入金额（元）
	ShBuy       float64 `json:"sh_buy,omitempty"`        // float	沪市买入金额（元）
	ShSell      float64 `json:"sh_sell,omitempty"`       // float	沪市卖出金额
	SzAmount    float64 `json:"sz_amount,omitempty"`     // float	深市成交金额（元）
	SzNetAmount float64 `json:"sz_net_amount,omitempty"` // float	深市净买入金额（元）
	SzBuy       float64 `json:"sz_buy,omitempty"`        // float	深市买入金额（元）
	SzSell      float64 `json:"sz_sell,omitempty"`       // float	深市卖出金额（元）
}

func AssembleGGTTop10Data(tsRsp *TushareResponse) []*GGTTop10Data {
	tsData := []*GGTTop10Data{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(GGTTop10Data)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 获取港股通每日成交数据，其中包括沪市、深市详细数据
func (ts *TuShare) GGTTop10(params GGTTop10Request, items GGTTop10Items) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "ggt_top10",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type MarginRequest struct {
	TradeDate  string `json:"trade_date,omitempty"`  // str	N	交易日期
	ExchangeID string `json:"exchange_id,omitempty"` // str	N	交易所代码
	StartDate  string `json:"start_date,omitempty"`  // str	N	开始日期
	EndDate    string `json:"end_date,omitempty"`    // str	N	结束日期
}

type MarginItems struct {
	TradeDate  bool `json:"trade_date,omitempty"`  // str	交易日期
	ExchangeID bool `json:"exchange_id,omitempty"` // str	交易所代码（SSE上交所SZSE深交所）
	Rzye       bool `json:"rzye,omitempty"`        // float	融资余额(元)
	Rzmre      bool `json:"rzmre,omitempty"`       // float	融资买入额(元)
	Rzche      bool `json:"rzche,omitempty"`       // float	融资偿还额(元)
	Rqye       bool `json:"rqye,omitempty"`        // float	融券余额(元)
	Rqmcl      bool `json:"rqmcl,omitempty"`       // float	融券卖出量(股,份,手)
	Rzrqye     bool `json:"rzrqye,omitempty"`      // float	融资融券余额(元)
	Rqyl       bool `json:"rqyl,omitempty"`        // float	融券余量(股,份,手)
}

func (item MarginItems) All() MarginItems {
	item.TradeDate = true
	item.ExchangeID = true
	item.Rzye = true
	item.Rzmre = true
	item.Rzche = true
	item.Rqye = true
	item.Rqmcl = true
	item.Rzrqye = true
	item.Rqyl = true
	return item
}

type MarginData struct {
	TradeDate  string  `json:"trade_date,omitempty"`  // str	交易日期
	ExchangeID string  `json:"exchange_id,omitempty"` // str	交易所代码（SSE上交所SZSE深交所）
	Rzye       float64 `json:"rzye,omitempty"`        // float	融资余额(元)
	Rzmre      float64 `json:"rzmre,omitempty"`       // float	融资买入额(元)
	Rzche      float64 `json:"rzche,omitempty"`       // float	融资偿还额(元)
	Rqye       float64 `json:"rqye,omitempty"`        // float	融券余额(元)
	Rqmcl      float64 `json:"rqmcl,omitempty"`       // float	融券卖出量(股,份,手)
	Rzrqye     float64 `json:"rzrqye,omitempty"`      // float	融资融券余额(元)
	Rqyl       float64 `json:"rqyl,omitempty"`        // float	融券余量(股,份,手)
}

func AssembleMarginData(tsRsp *TushareResponse) []*MarginData {
	tsData := []*MarginData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(MarginData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 获取融资融券每日交易汇总数据
func (ts *TuShare) Margin(params MarginRequest, items MarginItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "margin",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type MarginDetailRequest struct {
	TradeDate string `json:"trade_date,omitempty"` // str	N	交易日期
	TsCode    string `json:"ts_code,omitempty"`    // str	N	TS代码
	StartDate string `json:"start_date,omitempty"` // str	N	开始日期
	EndDate   string `json:"end_date,omitempty"`   // str	N	结束日期
}

type MarginDetailItems struct {
	TradeDate bool `json:"trade_date,omitempty"` // str	交易日期
	TsCode    bool `json:"ts_code,omitempty"`    // str	TS股票代码
	Name      bool `json:"name,omitempty"`       // str	股票名称 （20190910后有数据）
	Rzye      bool `json:"rzye,omitempty"`       // float	融资余额(元)
	Rqye      bool `json:"rqye,omitempty"`       // float	融券余额(元)
	Rzmre     bool `json:"rzmre,omitempty"`      // float	融资买入额(元)
	Rqyl      bool `json:"rqyl,omitempty"`       // float	融券余量（手）
	Rzche     bool `json:"rzche,omitempty"`      // float	融资偿还额(元)
	Rqchl     bool `json:"rqchl,omitempty"`      // float	融券偿还量(手)
	Rqmcl     bool `json:"rqmcl,omitempty"`      // float	融券卖出量(股,份,手)
	Rzrqye    bool `json:"rzrqye,omitempty"`     // float	融资融券余额(元)
}

func (item MarginDetailItems) All() MarginDetailItems {
	item.TradeDate = true
	item.TsCode = true
	item.Name = true
	item.Rzye = true
	item.Rqye = true
	item.Rzmre = true
	item.Rqyl = true
	item.Rzche = true
	item.Rqchl = true
	item.Rqmcl = true
	item.Rzrqye = true
	return item
}

type MarginDetailData struct {
	TradeDate string  `json:"trade_date,omitempty"` // str	交易日期
	TsCode    string  `json:"ts_code,omitempty"`    // str	TS股票代码
	Name      string  `json:"name,omitempty"`       // str	股票名称 （20190910后有数据）
	Rzye      float64 `json:"rzye,omitempty"`       // float	融资余额(元)
	Rqye      float64 `json:"rqye,omitempty"`       // float	融券余额(元)
	Rzmre     float64 `json:"rzmre,omitempty"`      // float	融资买入额(元)
	Rqyl      float64 `json:"rqyl,omitempty"`       // float	融券余量（手）
	Rzche     float64 `json:"rzche,omitempty"`      // float	融资偿还额(元)
	Rqchl     float64 `json:"rqchl,omitempty"`      // float	融券偿还量(手)
	Rqmcl     float64 `json:"rqmcl,omitempty"`      // float	融券卖出量(股,份,手)
	Rzrqye    float64 `json:"rzrqye,omitempty"`     // float	融资融券余额(元)
}

func AssembleMarginDetailData(tsRsp *TushareResponse) []*MarginDetailData {
	tsData := []*MarginDetailData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(MarginDetailData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 获取沪深两市每日融资融券明细
func (ts *TuShare) MarginDetail(params MarginDetailRequest, items MarginDetailItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "margin_detail",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type HoldersRequest struct {
	TsCode    string `json:"ts_code,omitempty"`    // str	Y	TS代码
	Period    string `json:"period,omitempty"`     // str	N	报告期
	AnnDate   string `json:"ann_date,omitempty"`   // str	N	公告日期
	StartDate string `json:"start_date,omitempty"` // str	N	报告期开始日期
	EndDate   string `json:"end_date,omitempty"`   // str	N	报告期结束日期
}

type Top10HoldersItems struct {
	TsCode     bool `json:"ts_code,omitempty"`     // str	TS股票代码
	AnnDate    bool `json:"ann_date,omitempty"`    // str	公告日期
	EndDate    bool `json:"end_date,omitempty"`    // str	报告期
	HolderName bool `json:"holder_name,omitempty"` // str	股东名称
	HoldAmount bool `json:"hold_amount,omitempty"` // float	持有数量（股）
	HoldRatio  bool `json:"hold_ratio,omitempty"`  // float	持有比例
}

func (item Top10HoldersItems) All() Top10HoldersItems {
	item.TsCode = true
	item.AnnDate = true
	item.EndDate = true
	item.HolderName = true
	item.HoldAmount = true
	item.HoldRatio = true
	return item
}

type Top10HoldersData struct {
	TsCode     string  `json:"ts_code,omitempty"`     // str	TS股票代码
	AnnDate    string  `json:"ann_date,omitempty"`    // str	公告日期
	EndDate    string  `json:"end_date,omitempty"`    // str	报告期
	HolderName string  `json:"holder_name,omitempty"` // str	股东名称
	HoldAmount float64 `json:"hold_amount,omitempty"` // float	持有数量（股）
	HoldRatio  float64 `json:"hold_ratio,omitempty"`  // float	持有比例
}

func AssembleTop10HoldersData(tsRsp *TushareResponse) []*Top10HoldersData {
	tsData := []*Top10HoldersData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(Top10HoldersData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 获取上市公司前十大股东数据，包括持有数量和比例等信息
func (ts *TuShare) Top10Holders(params HoldersRequest, items Top10HoldersItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "top10_holders",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type Top10FloatHoldersItems struct {
	TsCode     bool `json:"ts_code,omitempty"`     // str	TS股票代码
	AnnDate    bool `json:"ann_date,omitempty"`    // str	公告日期
	EndDate    bool `json:"end_date,omitempty"`    // str	报告期
	HolderName bool `json:"holder_name,omitempty"` // str	股东名称
	HoldAmount bool `json:"hold_amount,omitempty"` // float	持有数量（股）
}

func (item Top10FloatHoldersItems) All() Top10FloatHoldersItems {
	item.TsCode = true
	item.AnnDate = true
	item.EndDate = true
	item.HolderName = true
	item.HoldAmount = true
	return item
}

type Top10FloatHoldersData struct {
	TsCode     string  `json:"ts_code,omitempty"`     // str	TS股票代码
	AnnDate    string  `json:"ann_date,omitempty"`    // str	公告日期
	EndDate    string  `json:"end_date,omitempty"`    // str	报告期
	HolderName string  `json:"holder_name,omitempty"` // str	股东名称
	HoldAmount float64 `json:"hold_amount,omitempty"` // float	持有数量（股）
}

func AssembleTop10FloatHoldersData(tsRsp *TushareResponse) []*Top10FloatHoldersData {
	tsData := []*Top10FloatHoldersData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(Top10FloatHoldersData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 获取上市公司前十大流通股东数据
func (ts *TuShare) Top10FloatHolders(params HoldersRequest, items Top10FloatHoldersItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "top10_floatholders",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type TopRequest struct {
	TradeDate string `json:"trade_date,omitempty"` // str	Y	交易日期
	TsCode    string `json:"ts_code,omitempty"`    // str	N	股票代码
}

type TopListItems struct {
	TradeDate    bool `json:"trade_date,omitempty"`    // str	Y	交易日期
	TsCode       bool `json:"ts_code,omitempty"`       // str	Y	TS代码
	Name         bool `json:"name,omitempty"`          // str	Y	名称
	Close        bool `json:"close,omitempty"`         // float	Y	收盘价
	PctChange    bool `json:"pct_change,omitempty"`    // float	Y	涨跌幅
	TurnoverRate bool `json:"turnover_rate,omitempty"` // float	Y	换手率
	Amount       bool `json:"amount,omitempty"`        // float	Y	总成交额
	LSell        bool `json:"l_sell,omitempty"`        // float	Y	龙虎榜卖出额
	LBuy         bool `json:"l_buy,omitempty"`         // float	Y	龙虎榜买入额
	LAmount      bool `json:"l_amount,omitempty"`      // float	Y	龙虎榜成交额
	NetAmount    bool `json:"net_amount,omitempty"`    // float	Y	龙虎榜净买入额
	NetRate      bool `json:"net_rate,omitempty"`      // float	Y	龙虎榜净买额占比
	AmountRate   bool `json:"amount_rate,omitempty"`   // float	Y	龙虎榜成交额占比
	FloatValues  bool `json:"float_values,omitempty"`  // float	Y	当日流通市值
	Reason       bool `json:"reason,omitempty"`        // str	Y	上榜理由
}

func (item TopListItems) All() TopListItems {
	item.TradeDate = true
	item.TsCode = true
	item.Name = true
	item.Close = true
	item.PctChange = true
	item.TurnoverRate = true
	item.Amount = true
	item.LSell = true
	item.LBuy = true
	item.LAmount = true
	item.NetAmount = true
	item.NetRate = true
	item.AmountRate = true
	item.FloatValues = true
	item.Reason = true
	return item
}

type TopListData struct {
	TradeDate    string  `json:"trade_date,omitempty"`    // str	Y	交易日期
	TsCode       string  `json:"ts_code,omitempty"`       // str	Y	TS代码
	Name         string  `json:"name,omitempty"`          // str	Y	名称
	Close        float64 `json:"close,omitempty"`         // float	Y	收盘价
	PctChange    float64 `json:"pct_change,omitempty"`    // float	Y	涨跌幅
	TurnoverRate float64 `json:"turnover_rate,omitempty"` // float	Y	换手率
	Amount       float64 `json:"amount,omitempty"`        // float	Y	总成交额
	LSell        float64 `json:"l_sell,omitempty"`        // float	Y	龙虎榜卖出额
	LBuy         float64 `json:"l_buy,omitempty"`         // float	Y	龙虎榜买入额
	LAmount      float64 `json:"l_amount,omitempty"`      // float	Y	龙虎榜成交额
	NetAmount    float64 `json:"net_amount,omitempty"`    // float	Y	龙虎榜净买入额
	NetRate      float64 `json:"net_rate,omitempty"`      // float	Y	龙虎榜净买额占比
	AmountRate   float64 `json:"amount_rate,omitempty"`   // float	Y	龙虎榜成交额占比
	FloatValues  float64 `json:"float_values,omitempty"`  // float	Y	当日流通市值
	Reason       string  `json:"reason,omitempty"`        // str	Y	上榜理由
}

func AssembleTopListData(tsRsp *TushareResponse) []*TopListData {
	tsData := []*TopListData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(TopListData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 龙虎榜每日交易明细,数据为2005年至今,单次提取10000条,用户需要至少300积分才可以调取,具体请参阅积分获取办法 https://tushare.pro/document/1?doc_id=13
func (ts *TuShare) TopList(params TopRequest, items TopListItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "top_list",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type TopInstItems struct {
	TradeDate bool `json:"trade_date,omitempty"` // str	Y	交易日期
	TsCode    bool `json:"ts_code,omitempty"`    // str	Y	TS代码
	Exalter   bool `json:"exalter,omitempty"`    // str	Y	营业部名称
	Buy       bool `json:"buy,omitempty"`        // float	Y	买入额（万）
	BuyRate   bool `json:"buy_rate,omitempty"`   // float	Y	买入占总成交比例
	Sell      bool `json:"sell,omitempty"`       // float	Y	卖出额（万）
	SellRate  bool `json:"sell_rate,omitempty"`  // float	Y	卖出占总成交比例
	NetBuy    bool `json:"net_buy,omitempty"`    // float	Y	净成交额（万）
}

func (item TopInstItems) All() TopInstItems {
	item.TradeDate = true
	item.TsCode = true
	item.Exalter = true
	item.Buy = true
	item.BuyRate = true
	item.Sell = true
	item.SellRate = true
	item.NetBuy = true
	return item
}

type TopInstData struct {
	TradeDate string  `json:"trade_date,omitempty"` // str	Y	交易日期
	TsCode    string  `json:"ts_code,omitempty"`    // str	Y	TS代码
	Exalter   string  `json:"exalter,omitempty"`    // str	Y	营业部名称
	Buy       float64 `json:"buy,omitempty"`        // float	Y	买入额（万）
	BuyRate   float64 `json:"buy_rate,omitempty"`   // float	Y	买入占总成交比例
	Sell      float64 `json:"sell,omitempty"`       // float	Y	卖出额（万）
	SellRate  float64 `json:"sell_rate,omitempty"`  // float	Y	卖出占总成交比例
	NetBuy    float64 `json:"net_buy,omitempty"`    // float	Y	净成交额（万）
}

func AssembleTopInstData(tsRsp *TushareResponse) []*TopInstData {
	tsData := []*TopInstData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(TopInstData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 龙虎榜机构成交明细,数据为2005年至今,单次提取10000条,用户需要至少300积分才可以调取,具体请参阅积分获取办法 https://tushare.pro/document/1?doc_id=13
func (ts *TuShare) TopInst(params TopRequest, items TopInstItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "top_inst",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type PledgeStatRequest struct {
	TsCode  string `json:"ts_code,omitempty"`  // str	N	股票代码
	EndDate string `json:"end_date,omitempty"` // str	N	截止日期
}

type PledgeStatItems struct {
	TsCode       bool `json:"ts_code,omitempty"`       // str	Y	TS代码
	EndDate      bool `json:"end_date,omitempty"`      // str	Y	截止日期
	PledgeCount  bool `json:"pledge_count,omitempty"`  // int	Y	质押次数
	UnrestPledge bool `json:"unrest_pledge,omitempty"` // float	Y	无限售股质押数量（万）
	RestPledge   bool `json:"rest_pledge,omitempty"`   // float	Y	限售股份质押数量（万）
	TotalShare   bool `json:"total_share,omitempty"`   // float	Y	总股本
	PledgeRatio  bool `json:"pledge_ratio,omitempty"`  // float	Y	质押比例
}

func (item PledgeStatItems) All() PledgeStatItems {
	item.TsCode = true
	item.EndDate = true
	item.PledgeCount = true
	item.UnrestPledge = true
	item.RestPledge = true
	item.TotalShare = true
	item.PledgeRatio = true
	return item
}

type PledgeStatData struct {
	TsCode       string  `json:"ts_code,omitempty"`       // str	Y	TS代码
	EndDate      string  `json:"end_date,omitempty"`      // str	Y	截止日期
	PledgeCount  int64   `json:"pledge_count,omitempty"`  // int	Y	质押次数
	UnrestPledge float64 `json:"unrest_pledge,omitempty"` // float	Y	无限售股质押数量（万）
	RestPledge   float64 `json:"rest_pledge,omitempty"`   // float	Y	限售股份质押数量（万）
	TotalShare   float64 `json:"total_share,omitempty"`   // float	Y	总股本
	PledgeRatio  float64 `json:"pledge_ratio,omitempty"`  // float	Y	质押比例
}

func AssemblePledgeStatData(tsRsp *TushareResponse) []*PledgeStatData {
	tsData := []*PledgeStatData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(PledgeStatData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 获取股票质押统计数据,单次提取1000条,用户需要至少500积分才可以调取,具体请参阅积分获取办法 https://tushare.pro/document/1?doc_id=13
func (ts *TuShare) PledgeStat(params PledgeStatRequest, items PledgeStatItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "pledge_stat",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type PledgeDetailRequest struct {
	TsCode string `json:"ts_code,omitempty"` // N	股票代码
}

type PledgeDetailItems struct {
	TsCode        bool `json:"ts_code,omitempty"`        // str	Y	TS股票代码
	AnnDate       bool `json:"ann_date,omitempty"`       // str	Y	公告日期
	HolderName    bool `json:"holder_name,omitempty"`    // str	Y	股东名称
	PledgeAmount  bool `json:"pledge_amount,omitempty"`  // float	Y	质押数量（万股）
	StartDate     bool `json:"start_date,omitempty"`     // str	Y	质押开始日期
	EndDate       bool `json:"end_date,omitempty"`       // str	Y	质押结束日期
	IsRelease     bool `json:"is_release,omitempty"`     // str	Y	是否已解押
	ReleaseDate   bool `json:"release_date,omitempty"`   // str	Y	解押日期
	Pledgor       bool `json:"pledgor,omitempty"`        // str	Y	质押方
	HoldingAmount bool `json:"holding_amount,omitempty"` // float	Y	持股总数（万股）
	PledgedAmount bool `json:"pledged_amount,omitempty"` // float	Y	质押总数（万股）
	PTotalRatio   bool `json:"p_total_ratio,omitempty"`  // float	Y	本次质押占总股本比例
	HTotalRatio   bool `json:"h_total_ratio,omitempty"`  // float	Y	持股总数占总股本比例
	IsBuyback     bool `json:"is_buyback,omitempty"`     // str	Y	是否回购
}

func (item PledgeDetailItems) All() PledgeDetailItems {
	item.TsCode = true
	item.AnnDate = true
	item.HolderName = true
	item.PledgeAmount = true
	item.StartDate = true
	item.EndDate = true
	item.IsRelease = true
	item.ReleaseDate = true
	item.Pledgor = true
	item.HoldingAmount = true
	item.PledgedAmount = true
	item.PTotalRatio = true
	item.HTotalRatio = true
	item.IsBuyback = true
	return item
}

type PledgeDetailData struct {
	TsCode        string  `json:"ts_code,omitempty"`        // str	Y	TS股票代码
	AnnDate       string  `json:"ann_date,omitempty"`       // str	Y	公告日期
	HolderName    string  `json:"holder_name,omitempty"`    // str	Y	股东名称
	PledgeAmount  float64 `json:"pledge_amount,omitempty"`  // float	Y	质押数量（万股）
	StartDate     string  `json:"start_date,omitempty"`     // str	Y	质押开始日期
	EndDate       string  `json:"end_date,omitempty"`       // str	Y	质押结束日期
	IsRelease     string  `json:"is_release,omitempty"`     // str	Y	是否已解押
	ReleaseDate   string  `json:"release_date,omitempty"`   // str	Y	解押日期
	Pledgor       string  `json:"pledgor,omitempty"`        // str	Y	质押方
	HoldingAmount float64 `json:"holding_amount,omitempty"` // float	Y	持股总数（万股）
	PledgedAmount float64 `json:"pledged_amount,omitempty"` // float	Y	质押总数（万股）
	PTotalRatio   float64 `json:"p_total_ratio,omitempty"`  // float	Y	本次质押占总股本比例
	HTotalRatio   float64 `json:"h_total_ratio,omitempty"`  // float	Y	持股总数占总股本比例
	IsBuyback     string  `json:"is_buyback,omitempty"`     // str	Y	是否回购
}

func AssemblePledgeDetailData(tsRsp *TushareResponse) []*PledgeDetailData {
	tsData := []*PledgeDetailData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(PledgeDetailData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 获取股票质押明细数据,单次提取1000条用户需要至少500积分才可以调取,具体请参阅积分获取办法 https://tushare.pro/document/1?doc_id=13
func (ts *TuShare) PledgeDetail(params PledgeDetailRequest, items PledgeDetailItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "pledge_detail",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type RepurchaseRequest struct {
	AnnDate   string `json:"ann_date,omitempty"`   // str	N	公告日期（任意填参数，如果都不填，单次默认返回2000条）
	StartDate string `json:"start_date,omitempty"` // str	N	公告开始日期
	EndDate   string `json:"end_date,omitempty"`   // str	N	公告结束日期
}

type RepurchaseItems struct {
	TsCode    bool `json:"ts_code,omitempty"`    // str	Y	TS代码
	AnnDate   bool `json:"ann_date,omitempty"`   // str	Y	公告日期
	EndDate   bool `json:"end_date,omitempty"`   // str	Y	截止日期
	Proc      bool `json:"proc,omitempty"`       // str	Y	进度
	ExpDate   bool `json:"exp_date,omitempty"`   // str	Y	过期日期
	Vol       bool `json:"vol,omitempty"`        // float	Y	回购数量
	Amount    bool `json:"amount,omitempty"`     // float	Y	回购金额
	HighLimit bool `json:"high_limit,omitempty"` // float	Y	回购最高价
	LowLimit  bool `json:"low_limit,omitempty"`  // float	Y	回购最低价
}

func (item RepurchaseItems) All() RepurchaseItems {
	item.TsCode = true
	item.AnnDate = true
	item.EndDate = true
	item.Proc = true
	item.ExpDate = true
	item.Vol = true
	item.Amount = true
	item.HighLimit = true
	item.LowLimit = true
	return item
}

type RepurchaseData struct {
	TsCode    string  `json:"ts_code,omitempty"`    // str	Y	TS代码
	AnnDate   string  `json:"ann_date,omitempty"`   // str	Y	公告日期
	EndDate   string  `json:"end_date,omitempty"`   // str	Y	截止日期
	Proc      string  `json:"proc,omitempty"`       // str	Y	进度
	ExpDate   string  `json:"exp_date,omitempty"`   // str	Y	过期日期
	Vol       float64 `json:"vol,omitempty"`        // float	Y	回购数量
	Amount    float64 `json:"amount,omitempty"`     // float	Y	回购金额
	HighLimit float64 `json:"high_limit,omitempty"` // float	Y	回购最高价
	LowLimit  float64 `json:"low_limit,omitempty"`  // float	Y	回购最低价
}

func AssembleRepurchaseData(tsRsp *TushareResponse) []*RepurchaseData {
	tsData := []*RepurchaseData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(RepurchaseData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 获取上市公司回购股票数据,用户需要至少600积分才可以调取,具体请参阅积分获取办法 https://tushare.pro/document/1?doc_id=13
func (ts *TuShare) Repurchase(params RepurchaseRequest, items RepurchaseItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "repurchase",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type ConceptRequest struct {
	Src string `json:"src,omitempty"` // str	N	来源，默认为ts
}

type ConceptItems struct {
	Code bool `json:"code,omitempty"` // str	Y	概念分类ID
	Name bool `json:"name,omitempty"` // str	Y	概念分类名称
	Src  bool `json:"src,omitempty"`  // str	Y	来源
}

func (item ConceptItems) All() ConceptItems {
	item.Code = true
	item.Name = true
	item.Src = true
	return item
}

type ConceptData struct {
	Code string `json:"code,omitempty"` // str	Y	概念分类ID
	Name string `json:"name,omitempty"` // str	Y	概念分类名称
	Src  string `json:"src,omitempty"`  // str	Y	来源
}

func AssembleConceptData(tsRsp *TushareResponse) []*ConceptData {
	tsData := []*ConceptData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(ConceptData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

//获取概念股分类，目前只有ts一个来源，未来将逐步增加来源,用户需要至少300积分才可以调取,具体请参阅积分获取办法 https://tushare.pro/document/1?doc_id=13
func (ts *TuShare) Concept(params ConceptRequest, items ConceptItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "concept",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type ConceptDetailRequest struct {
	ID     string `json:"id,omitempty"`      // str	N	概念分类ID （id来自概念股分类接口:TS2）
	TsCode string `json:"ts_code,omitempty"` // str	N	股票代码 （以上参数二选一）
}

type ConceptDetailItems struct {
	ID          bool `json:"id,omitempty"`           // str	Y	概念代码
	ConceptName bool `json:"concept_name,omitempty"` // str	Y	概念名称
	TsCode      bool `json:"ts_code,omitempty"`      // str	Y	股票代码
	Name        bool `json:"name,omitempty"`         // str	Y	股票名称
	InDate      bool `json:"in_date,omitempty"`      // str	N	纳入日期
	OutDate     bool `json:"out_date,omitempty"`     // str	N	剔除日期
}

func (item ConceptDetailItems) All() ConceptDetailItems {
	item.ID = true
	item.ConceptName = true
	item.TsCode = true
	item.Name = true
	item.InDate = true
	item.OutDate = true
	return item
}

type ConceptDetailData struct {
	ID          string `json:"id,omitempty"`           // str	Y	概念代码
	ConceptName string `json:"concept_name,omitempty"` // str	Y	概念名称
	TsCode      string `json:"ts_code,omitempty"`      // str	Y	股票代码
	Name        string `json:"name,omitempty"`         // str	Y	股票名称
	InDate      string `json:"in_date,omitempty"`      // str	N	纳入日期
	OutDate     string `json:"out_date,omitempty"`     // str	N	剔除日期
}

func AssembleConceptDetailData(tsRsp *TushareResponse) []*ConceptDetailData {
	tsData := []*ConceptDetailData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(ConceptDetailData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

//获取概念股分类明细数据,用户需要至少300积分才可以调取,具体请参阅积分获取办法 https://tushare.pro/document/1?doc_id=13
func (ts *TuShare) ConceptDetail(params ConceptDetailRequest, items ConceptDetailItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "concept_detail",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type ShareFloatRequest struct {
	TsCode    string `json:"ts_code,omitempty"`    // str	N	TS股票代码（至少输入一个参数）
	AnnDate   string `json:"ann_date,omitempty"`   // str	N	公告日期（日期格式：YYYYMMDD，下同）
	FloatDate string `json:"float_date,omitempty"` // str	N	解禁日期
	StartDate string `json:"start_date,omitempty"` // str	N	解禁开始日期
	EndDate   string `json:"end_date,omitempty"`   // str	N	解禁结束日期
}

type ShareFloatItems struct {
	TsCode     bool `json:"ts_code,omitempty"`     // str	Y	TS代码
	AnnDate    bool `json:"ann_date,omitempty"`    // str	Y	公告日期
	FloatDate  bool `json:"float_date,omitempty"`  // str	Y	解禁日期
	FloatShare bool `json:"float_share,omitempty"` // float	Y	流通股份
	FloatRatio bool `json:"float_ratio,omitempty"` // float	Y	流通股份占总股本比率
	HolderName bool `json:"holder_name,omitempty"` // str	Y	股东名称
	ShareType  bool `json:"share_type,omitempty"`  // str	Y	股份类型
}

func (item ShareFloatItems) All() ShareFloatItems {
	item.TsCode = true
	item.AnnDate = true
	item.FloatDate = true
	item.FloatShare = true
	item.FloatRatio = true
	item.HolderName = true
	item.ShareType = true
	return item
}

type ShareFloatData struct {
	TsCode     string  `json:"ts_code,omitempty"`     // str	Y	TS代码
	AnnDate    string  `json:"ann_date,omitempty"`    // str	Y	公告日期
	FloatDate  string  `json:"float_date,omitempty"`  // str	Y	解禁日期
	FloatShare float64 `json:"float_share,omitempty"` // float	Y	流通股份
	FloatRatio float64 `json:"float_ratio,omitempty"` // float	Y	流通股份占总股本比率
	HolderName string  `json:"holder_name,omitempty"` // str	Y	股东名称
	ShareType  string  `json:"share_type,omitempty"`  // str	Y	股份类型
}

func AssembleShareFloatData(tsRsp *TushareResponse) []*ShareFloatData {
	tsData := []*ShareFloatData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(ShareFloatData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 获取限售股解禁,单次提取5000条,用户需要至少120积分才可以调取,每分钟内限制次数，超过5000积分频次相对较高,具体请参阅积分获取办法 https://tushare.pro/document/1?doc_id=13
func (ts *TuShare) ShareFloat(params ShareFloatRequest, items ShareFloatItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "share_float",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type BlockTradeRequest struct {
	TsCode    string `json:"ts_code,omitempty"`    // str	N	TS代码（股票代码和日期至少输入一个参数）
	TradeDate string `json:"trade_date,omitempty"` // str	N	交易日期（格式：YYYYMMDD，下同）
	StartDate string `json:"start_date,omitempty"` // str	N	开始日期
	EndDate   string `json:"end_date,omitempty"`   // str	N	结束日期
}

type BlockTradeItems struct {
	TsCode    bool `json:"ts_code,omitempty"`    // str	Y	TS代码
	TradeDate bool `json:"trade_date,omitempty"` // str	Y	交易日历
	Price     bool `json:"price,omitempty"`      // float	Y	成交价
	Vol       bool `json:"vol,omitempty"`        // float	Y	成交量（万股）
	Amount    bool `json:"amount,omitempty"`     // float	Y	成交金额
	Buyer     bool `json:"buyer,omitempty"`      // str	Y	买方营业部
	Seller    bool `json:"seller,omitempty"`     // str	Y	卖方营业部
}

func (item BlockTradeItems) All() BlockTradeItems {
	item.TsCode = true
	item.TradeDate = true
	item.Price = true
	item.Vol = true
	item.Amount = true
	item.Buyer = true
	item.Seller = true
	return item
}

type BlockTradeData struct {
	TsCode    string  `json:"ts_code,omitempty"`    // str	Y	TS代码
	TradeDate string  `json:"trade_date,omitempty"` // str	Y	交易日历
	Price     float64 `json:"price,omitempty"`      // float	Y	成交价
	Vol       float64 `json:"vol,omitempty"`        // float	Y	成交量（万股）
	Amount    float64 `json:"amount,omitempty"`     // float	Y	成交金额
	Buyer     string  `json:"buyer,omitempty"`      // str	Y	买方营业部
	Seller    string  `json:"seller,omitempty"`     // str	Y	卖方营业部
}

func AssembleBlockTradeData(tsRsp *TushareResponse) []*BlockTradeData {
	tsData := []*BlockTradeData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(BlockTradeData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 大宗交易,单次提取1000条,总量不限制,用户需要至少300积分才可以调取,每分钟内限制次数，超过5000积分频次相对较高,具体请参阅积分获取办法 https://tushare.pro/document/1?doc_id=13
func (ts *TuShare) BlockTrade(params BlockTradeRequest, items BlockTradeItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "block_trade",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type StkHoldernumberRequest struct {
	TsCode    string `json:"ts_code,omitempty"`    // str	N	TS股票代码
	Enddate   string `json:"enddate,omitempty"`    // str	N	截止日期
	StartDate string `json:"start_date,omitempty"` // str	N	公告开始日期
	EndDate   string `json:"end_date,omitempty"`   // str	N	公告结束日期
}

type StkHoldernumberItems struct {
	TsCode    bool `json:"ts_code,omitempty"`    // str	Y	TS股票代码
	AnnDate   bool `json:"ann_date,omitempty"`   // str	Y	公告日期
	EndDate   bool `json:"end_date,omitempty"`   // str	Y	截止日期
	HolderNum bool `json:"holder_num,omitempty"` // int	Y	股东户数
}

func (item StkHoldernumberItems) All() StkHoldernumberItems {
	item.TsCode = true
	item.AnnDate = true
	item.EndDate = true
	item.HolderNum = true
	return item
}

type StkHoldernumberData struct {
	TsCode    string `json:"ts_code,omitempty"`    // str	Y	TS股票代码
	AnnDate   string `json:"ann_date,omitempty"`   // str	Y	公告日期
	EndDate   string `json:"end_date,omitempty"`   // str	Y	截止日期
	HolderNum int64  `json:"holder_num,omitempty"` // int	Y	股东户数
}

func AssembleStkHoldernumberData(tsRsp *TushareResponse) []*StkHoldernumberData {
	tsData := []*StkHoldernumberData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(StkHoldernumberData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 获取上市公司股东户数数据，数据不定期公布,单次提取3000,总量不限制,用户需要至少600积分才可以调取,基础积分每分钟调取100次，5000积分以上无限制,具体请参阅积分获取办法 https://tushare.pro/document/1?doc_id=13
func (ts *TuShare) StkHoldernumber(params StkHoldernumberRequest, items StkHoldernumberItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "stk_holdernumber",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type STKHoldertradeRequest struct {
	TsCode     string `json:"ts_code,omitempty"`     // str	N	TS股票代码
	AnnDate    string `json:"ann_date,omitempty"`    // str	N	公告日期
	StartDate  string `json:"start_date,omitempty"`  // str	N	公告开始日期
	EndDate    string `json:"end_date,omitempty"`    // str	N	公告结束日期
	TradeType  string `json:"trade_type,omitempty"`  // str	N	交易类型IN增持DE减持
	HolderType string `json:"holder_type,omitempty"` // str	N	股东类型C公司P个人G高管
}

type STKHoldertradeItems struct {
	TsCode      bool `json:"ts_code,omitempty"`      // str	Y	TS代码
	AnnDate     bool `json:"ann_date,omitempty"`     // str	Y	公告日期
	HolderName  bool `json:"holder_name,omitempty"`  // str	Y	股东名称
	HolderType  bool `json:"holder_type,omitempty"`  // str	Y	股东类型G高管P个人C公司
	InDe        bool `json:"in_de,omitempty"`        // str	Y	类型IN增持DE减持
	ChangeVol   bool `json:"change_vol,omitempty"`   // float	Y	变动数量
	ChangeRatio bool `json:"change_ratio,omitempty"` // float	Y	占流通比例（%）
	AfterShare  bool `json:"after_share,omitempty"`  // float	Y	变动后持股
	AfterRatio  bool `json:"after_ratio,omitempty"`  // float	Y	变动后占流通比例（%）
	AvgPrice    bool `json:"avg_price,omitempty"`    // float	Y	平均价格
	TotalShare  bool `json:"total_share,omitempty"`  // float	Y	持股总数
	BeginDate   bool `json:"begin_date,omitempty"`   // str	N	增减持开始日期
	CloseDate   bool `json:"close_date,omitempty"`   // str	N	增减持结束日期
}

func (item STKHoldertradeItems) All() STKHoldertradeItems {
	item.TsCode = true
	item.AnnDate = true
	item.HolderName = true
	item.HolderType = true
	item.InDe = true
	item.ChangeVol = true
	item.ChangeRatio = true
	item.AfterShare = true
	item.AfterRatio = true
	item.AvgPrice = true
	item.TotalShare = true
	item.BeginDate = true
	item.CloseDate = true
	return item
}

type STKHoldertradeData struct {
	TsCode      string  `json:"ts_code,omitempty"`      // str	Y	TS代码
	AnnDate     string  `json:"ann_date,omitempty"`     // str	Y	公告日期
	HolderName  string  `json:"holder_name,omitempty"`  // str	Y	股东名称
	HolderType  string  `json:"holder_type,omitempty"`  // str	Y	股东类型G高管P个人C公司
	InDe        string  `json:"in_de,omitempty"`        // str	Y	类型IN增持DE减持
	ChangeVol   float64 `json:"change_vol,omitempty"`   // float	Y	变动数量
	ChangeRatio float64 `json:"change_ratio,omitempty"` // float	Y	占流通比例（%）
	AfterShare  float64 `json:"after_share,omitempty"`  // float	Y	变动后持股
	AfterRatio  float64 `json:"after_ratio,omitempty"`  // float	Y	变动后占流通比例（%）
	AvgPrice    float64 `json:"avg_price,omitempty"`    // float	Y	平均价格
	TotalShare  float64 `json:"total_share,omitempty"`  // float	Y	持股总数
	BeginDate   string  `json:"begin_date,omitempty"`   // str	N	增减持开始日期
	CloseDate   string  `json:"close_date,omitempty"`   // str	N	增减持结束日期
}

func AssembleSTKHoldertradeData(tsRsp *TushareResponse) []*STKHoldertradeData {
	tsData := []*STKHoldertradeData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(STKHoldertradeData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 获取上市公司管理层薪酬和持股,用户需要至少2000积分才可以调取,具体请参阅积分获取办法 https://tushare.pro/document/1?doc_id=13
func (ts *TuShare) STKHoldertrade(params STKHoldertradeRequest, items STKHoldertradeItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "stk_holdertrade",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}
