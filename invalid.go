package tushare

/*
type DailyRequest struct {
	TsCode    string `json:"ts_code,omitempty"`    // str	N	股票代码（支持多个股票同时提取，逗号分隔）
	TradeDate string `json:"trade_date,omitempty"` // str	N	交易日期（YYYYMMDD）
	StartDate string `json:"start_date,omitempty"` // str	N	开始日期(YYYYMMDD)
	EndDate   string `json:"end_date,omitempty"`   // str	N	结束日期(YYYYMMDD)
}

type DailyItems struct {
	TsCode    bool `json:"ts_code,omitempty"`    // str	股票代码
	TradeDate bool `json:"trade_date,omitempty"` // str	交易日期
	Open      bool `json:"open,omitempty"`       // float	开盘价
	High      bool `json:"high,omitempty"`       // float	最高价
	Low       bool `json:"low,omitempty"`        // float	最低价
	Close     bool `json:"close,omitempty"`      // float	收盘价
	PreClose  bool `json:"pre_close,omitempty"`  // float	昨收价
	Change    bool `json:"change,omitempty"`     // float	涨跌额
	PctChg    bool `json:"pct_chg,omitempty"`    // float	涨跌幅 （未复权，如果是复权请用 通用行情接口 ）
	Vol       bool `json:"vol,omitempty"`        // float	成交量 （手）
	Amount    bool `json:"amount,omitempty"`     // float	成交额 （千元）
}

func (item DailyItems) All() DailyItems {
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

type DailyData struct {
	TsCode    string  `json:"ts_code,omitempty"`    // str	股票代码
	TradeDate string  `json:"trade_date,omitempty"` // str	交易日期
	Open      float64 `json:"open,omitempty"`       // float	开盘价
	High      float64 `json:"high,omitempty"`       // float	最高价
	Low       float64 `json:"low,omitempty"`        // float	最低价
	Close     float64 `json:"close,omitempty"`      // float	收盘价
	PreClose  float64 `json:"pre_close,omitempty"`  // float	昨收价
	Change    float64 `json:"change,omitempty"`     // float	涨跌额
	PctChg    float64 `json:"pct_chg,omitempty"`    // float	涨跌幅 （未复权，如果是复权请用 通用行情接口 ）
	Vol       float64 `json:"vol,omitempty"`        // float	成交量 （手）
	Amount    float64 `json:"amount,omitempty"`     // float	成交额 （千元）
}

func AssembleDailyData(tsRsp *TushareResponse) []*DailyData {
	tsData := []*DailyData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(DailyData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

type WeeklyRequest struct {
	TsCode    string `json:"ts_code,omitempty"`    // str	N	TS代码 （ts_code,trade_date两个参数任选一）
	TradeDate string `json:"trade_date,omitempty"` // str	N （每周最后一个交易日期，YYYYMMDD格式）
	StartDate string `json:"start_date,omitempty"` // str	N	开始日期(YYYYMMDD)
	EndDate   string `json:"end_date,omitempty"`   // str	N	结束日期(YYYYMMDD)
}

type WeeklyItems struct {
	TsCode    bool `json:"ts_code,omitempty"`    // str	股票代码
	TradeDate bool `json:"trade_date,omitempty"` // str	交易日期
	Open      bool `json:"open,omitempty"`       // float	周开盘价
	High      bool `json:"high,omitempty"`       // float	周最高价
	Low       bool `json:"low,omitempty"`        // float	周最低价
	Close     bool `json:"close,omitempty"`      // float	周收盘价
	PreClose  bool `json:"pre_close,omitempty"`  // float	上一周收盘价
	Change    bool `json:"change,omitempty"`     // float	周涨跌额
	PctChg    bool `json:"pct_chg,omitempty"`    // float	周涨跌幅 （未复权，如果是复权请用 通用行情接口 ）
	Vol       bool `json:"vol,omitempty"`        // float	周成交量 （手）
	Amount    bool `json:"amount,omitempty"`     // float	周成交额 （千元）
}

func (item WeeklyItems) All() WeeklyItems {
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

type WeeklyData struct {
	TsCode    string  `json:"ts_code,omitempty"`    // str	股票代码
	TradeDate string  `json:"trade_date,omitempty"` // str	交易日期
	Open      float64 `json:"open,omitempty"`       // float	周开盘价
	High      float64 `json:"high,omitempty"`       // float	周最高价
	Low       float64 `json:"low,omitempty"`        // float	周最低价
	Close     float64 `json:"close,omitempty"`      // float	周收盘价
	PreClose  float64 `json:"pre_close,omitempty"`  // float	上一周收盘价
	Change    float64 `json:"change,omitempty"`     // float	周涨跌额
	PctChg    float64 `json:"pct_chg,omitempty"`    // float	周涨跌幅 （未复权，如果是复权请用 通用行情接口 ）
	Vol       float64 `json:"vol,omitempty"`        // float	周成交量 （手）
	Amount    float64 `json:"amount,omitempty"`     // float	周成交额 （千元）
}

func AssembleWeeklyData(tsRsp *TushareResponse) []*WeeklyData {
	tsData := []*WeeklyData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(WeeklyData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}



type MonthlyRequest struct {
	TsCode    string `json:"ts_code,omitempty"`    // str	N	TS代码 （ts_code,trade_date两个参数任选一）
	TradeDate string `json:"trade_date,omitempty"` // str	N （每周最后一个交易日期，YYYYMMDD格式）
	StartDate string `json:"start_date,omitempty"` // str	N	开始日期(YYYYMMDD)
	EndDate   string `json:"end_date,omitempty"`   // str	N	结束日期(YYYYMMDD)
}

type MonthlyItems struct {
	TsCode    bool `json:"ts_code,omitempty"`    // str	股票代码
	TradeDate bool `json:"trade_date,omitempty"` // str	交易日期
	Open      bool `json:"open,omitempty"`       // float	月开盘价
	High      bool `json:"high,omitempty"`       // float	月最高价
	Low       bool `json:"low,omitempty"`        // float	月最低价
	Close     bool `json:"close,omitempty"`      // float	月收盘价
	PreClose  bool `json:"pre_close,omitempty"`  // float	上一月收盘价
	Change    bool `json:"change,omitempty"`     // float	月涨跌额
	PctChg    bool `json:"pct_chg,omitempty"`    // float	月涨跌幅 （未复权，如果是复权请用 通用行情接口 ）
	Vol       bool `json:"vol,omitempty"`        // float	月成交量 （手）
	Amount    bool `json:"amount,omitempty"`     // float	月成交额 （千元）
}

func (item MonthlyItems) All() MonthlyItems {
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

type MonthlyData struct {
	TsCode    string  `json:"ts_code,omitempty"`    // str	股票代码
	TradeDate string  `json:"trade_date,omitempty"` // str	交易日期
	Open      float64 `json:"open,omitempty"`       // float	月开盘价
	High      float64 `json:"high,omitempty"`       // float	月最高价
	Low       float64 `json:"low,omitempty"`        // float	月最低价
	Close     float64 `json:"close,omitempty"`      // float	月收盘价
	PreClose  float64 `json:"pre_close,omitempty"`  // float	上一月收盘价
	Change    float64 `json:"change,omitempty"`     // float	月涨跌额
	PctChg    float64 `json:"pct_chg,omitempty"`    // float	月涨跌幅 （未复权，如果是复权请用 通用行情接口 ）
	Vol       float64 `json:"vol,omitempty"`        // float	月成交量 （手）
	Amount    float64 `json:"amount,omitempty"`     // float	月成交额 （千元）
}

func AssembleMonthlyData(tsRsp *TushareResponse) []*MonthlyData {
	tsData := []*MonthlyData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(MonthlyData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

type AdjFactorRequest struct {
	TsCode    string `json:"ts_code,omitempty"`    // str	Y	股票代码
	TradeDate string `json:"trade_date,omitempty"` // str	N	交易日期(YYYYMMDD，下同)
	StartDate string `json:"start_date,omitempty"` // str	N	开始日期
	EndDate   string `json:"end_date,omitempty"`   // str	N	结束日期
}

type ProBarItems struct {
	QuotationItems
	MA bool `json:"ma,omitempty"` // list	N	均线，支持任意周期的均价和均量，输入任意合理int数值
}

func (item ProBarItems) All() ProBarItems {
	item.QuotationItems = item.QuotationItems.All()
	item.MA = true
	return item
}
*/
