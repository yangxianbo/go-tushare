package tushare

import (
	"encoding/json"
	"net/http"
	"time"
)

type StockBasicRequest struct {
	TsCode     string `json:"ts_code,omitempty"`     // N	股票代码
	ListStatus string `json:"list_status,omitempty"` // N	上市状态: L上市 D退市 P暂停上市,默认L
	Exchange   string `json:"exchange,omitempty"`    // N	交易所 SSE上交所 SZSE深交所 HKEX港交所(未上线)
	IsHs       string `json:"is_hs,omitempty"`       // N  是否沪深港通标的,N否 H沪股通 S深股通
}

type StockBasicItems struct {
	TsCode     bool `json:"ts_code,omitempty"`     // str TS代码
	Symbol     bool `json:"symbol,omitempty"`      // str 股票代码
	Name       bool `json:"name,omitempty"`        // str 股票名称
	Area       bool `json:"area,omitempty"`        // str 所在地域
	Industry   bool `json:"industry,omitempty"`    // str 所属行业
	Fullname   bool `json:"fullname,omitempty"`    // str 股票全称
	Enname     bool `json:"enname,omitempty"`      // str 英文全称
	Market     bool `json:"market,omitempty"`      // str 市场类型 (主板/中小板/创业板/科创板/CDR)
	Exchange   bool `json:"exchange,omitempty"`    // str 交易所代码
	CurrType   bool `json:"curr_type,omitempty"`   // str 交易货币
	ListStatus bool `json:"list_status,omitempty"` // str 上市状态: L上市 D退市 P暂停上市
	ListDate   bool `json:"list_date,omitempty"`   // str 上市日期
	DelistDate bool `json:"delist_date,omitempty"` // str 退市日期
	IsHs       bool `json:"is_hs,omitempty"`       // str 是否沪深港通标的,N否 H沪股通 S深股通
}

func (item StockBasicItems) All() StockBasicItems {
	item.TsCode = true
	item.Symbol = true
	item.Name = true
	item.Area = true
	item.Industry = true
	item.Fullname = true
	item.Enname = true
	item.Market = true
	item.Exchange = true
	item.CurrType = true
	item.ListStatus = true
	item.ListDate = true
	item.DelistDate = true
	item.IsHs = true
	return item
}

type StockBasicData struct {
	TsCode     string `json:"ts_code,omitempty" gorm:"uniqueIndex"` // str TS代码
	Symbol     string `json:"symbol,omitempty"`                     // str 股票代码
	Name       string `json:"name,omitempty"`                       // str 股票名称
	Area       string `json:"area,omitempty"`                       // str 所在地域
	Industry   string `json:"industry,omitempty"`                   // str 所属行业
	Fullname   string `json:"fullname,omitempty"`                   // str 股票全称
	Enname     string `json:"enname,omitempty"`                     // str 英文全称
	Market     string `json:"market,omitempty"`                     // str 市场类型 (主板/中小板/创业板/科创板/CDR)
	Exchange   string `json:"exchange,omitempty"`                   // str 交易所代码
	CurrType   string `json:"curr_type,omitempty"`                  // str 交易货币
	ListStatus string `json:"list_status,omitempty"`                // str 上市状态: L上市 D退市 P暂停上市
	ListDate   string `json:"list_date,omitempty"`                  // str 上市日期
	DelistDate string `json:"delist_date,omitempty"`                // str 退市日期
	IsHs       string `json:"is_hs,omitempty"`                      // str 是否沪深港通标的,N否 H沪股通 S深股通
}

func AssembleStockBasicData(tsRsp *TushareResponse) []*StockBasicData {
	tsData := []*StockBasicData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(StockBasicData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 获取基础信息数据,包括股票代码、名称、上市日期、退市日期等
func (ts *TuShare) StockBasic(params StockBasicRequest, items StockBasicItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "stock_basic",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type TradeCalRequest struct {
	Exchange  string `json:"exchange,omitempty"`   // N	交易所 SSE上交所,SZSE深交所,CFFEX 中金所,SHFE 上期所,CZCE 郑商所,DCE 大商所,INE 上能源
	StartDate string `json:"start_date,omitempty"` // N	开始日期 (格式:YYYYMMDD 下同)
	EndDate   string `json:"end_date,omitempty"`   // N	结束日期
	IsOpen    string `json:"is_open,omitempty"`    // N	是否交易 '0'休市 '1'交易
}

type TradeCalItems struct {
	Exchange     bool `json:"exchange,omitempty"`      // str Y	交易所 SSE上交所 SZSE深交所
	CalDate      bool `json:"cal_date,omitempty"`      // str Y	日历日期
	IsOpen       bool `json:"is_open,omitempty"`       // str Y	是否交易 0休市 1交易
	PretradeDate bool `json:"pretrade_date,omitempty"` // str N	上一个交易日
}

func (item TradeCalItems) All() TradeCalItems {
	item.Exchange = true
	item.CalDate = true
	item.IsOpen = true
	item.PretradeDate = true
	return item
}

type TradeCalData struct {
	Exchange     string `json:"exchange,omitempty"`                    // str Y	交易所 SSE上交所 SZSE深交所
	CalDate      string `json:"cal_date,omitempty" gorm:"uniqueIndex"` // str Y	日历日期
	IsOpen       string `json:"is_open,omitempty"`                     // str Y	是否交易 0休市 1交易
	PretradeDate string `json:"pretrade_date,omitempty"`               // str N	上一个交易日
}

func AssembleTradeCalData(tsRsp *TushareResponse) []*TradeCalData {
	tsData := []*TradeCalData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(TradeCalData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 获取各大交易所交易日历数据,默认提取的是上交所
func (ts *TuShare) TradeCal(params TradeCalRequest, items TradeCalItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "trade_cal",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type NameChangeRequest struct {
	TsCode    string `json:"ts_code,omitempty"`    // N	TS代码
	StartDate string `json:"start_date,omitempty"` // N	公告开始日期
	EndDate   string `json:"end_date,omitempty"`   // N	公告结束日期
}

type NameChangeItems struct {
	TsCode       bool `json:"ts_code,omitempty"`       // str Y	TS代码
	Name         bool `json:"name,omitempty"`          // str Y	证券名称
	StartDate    bool `json:"start_date,omitempty"`    // str Y	开始日期
	EndDate      bool `json:"end_date,omitempty"`      // str Y	结束日期
	AnnDate      bool `json:"STKRewards,omitempty"`    // str Y	公告日期
	ChangeReason bool `json:"change_reason,omitempty"` // str Y	变更原因
}

func (item NameChangeItems) All() NameChangeItems {
	item.TsCode = true
	item.Name = true
	item.StartDate = true
	item.EndDate = true
	item.AnnDate = true
	item.ChangeReason = true
	return item
}

type NameChangeData struct {
	TsCode       string `json:"ts_code,omitempty"`       // str Y	TS代码
	Name         string `json:"name,omitempty"`          // str Y	证券名称
	StartDate    string `json:"start_date,omitempty"`    // str Y	开始日期
	EndDate      string `json:"end_date,omitempty"`      // str Y	结束日期
	AnnDate      string `json:"STKRewards,omitempty"`    // str Y	公告日期
	ChangeReason string `json:"change_reason,omitempty"` // str Y	变更原因
}

func AssembleNameChangeData(tsRsp *TushareResponse) []*NameChangeData {
	tsData := []*NameChangeData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(NameChangeData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 历史名称变更记录
func (ts *TuShare) NameChange(params NameChangeRequest, items NameChangeItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "namechange",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type HsConstRequest struct {
	HsType string `json:"hs_type,omitempty"` // Y	类型SH沪股通SZ深股通
	IsNew  string `json:"is_new,omitempty"`  // N	是否最新 1 是 0 否 (默认1)
}

type HsConstItems struct {
	TsCode  bool `json:"ts_code,omitempty"`  // str Y	TS代码
	HsType  bool `json:"hs_type,omitempty"`  // str Y	沪深港通类型SH沪SZ深
	InDate  bool `json:"in_date,omitempty"`  // str Y	纳入日期
	OutDate bool `json:"out_date,omitempty"` // str Y	剔除日期
	IsNew   bool `json:"is_new,omitempty"`   // str Y	是否最新 1是 0否
}

func (item HsConstItems) All() HsConstItems {
	item.TsCode = true
	item.HsType = true
	item.InDate = true
	item.OutDate = true
	item.IsNew = true
	return item
}

type HsConstData struct {
	TsCode  string `json:"ts_code,omitempty"`  // str Y	TS代码
	HsType  string `json:"hs_type,omitempty"`  // str Y	沪深港通类型SH沪SZ深
	InDate  string `json:"in_date,omitempty"`  // str Y	纳入日期
	OutDate string `json:"out_date,omitempty"` // str Y	剔除日期
	IsNew   string `json:"is_new,omitempty"`   // str Y	是否最新 1是 0否
}

func AssembleHsConstData(tsRsp *TushareResponse) []*HsConstData {
	tsData := []*HsConstData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(HsConstData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 获取沪股通、深股通成分数据
func (ts *TuShare) HsConst(params HsConstRequest, items HsConstItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "hs_const",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type StockCompanyRequest struct {
	TsCode   string `json:"ts_code,omitempty"`  // N	股票代码
	Exchange string `json:"exchange,omitempty"` // N	交易所代码 ,SSE上交所 SZSE深交所
}

type StockCompanyItems struct {
	TsCode        bool `json:"ts_code,omitempty"`        // str Y	股票代码
	Exchange      bool `json:"exchange,omitempty"`       // str Y	交易所代码 ,SSE上交所 SZSE深交所
	Chairman      bool `json:"chairman,omitempty"`       // str Y	法人代表
	Manager       bool `json:"manager,omitempty"`        // str Y	总经理
	Secretary     bool `json:"secretary,omitempty"`      // str Y	董秘
	RegCapital    bool `json:"reg_capital,omitempty"`    // float64 Y	注册资本
	SetupDate     bool `json:"setup_date,omitempty"`     // str Y	注册日期
	Province      bool `json:"province,omitempty"`       // flote64 Y	所在省份
	City          bool `json:"city,omitempty"`           // str Y	所在城市
	Introduction  bool `json:"introduction,omitempty"`   // str N	公司介绍
	Website       bool `json:"website,omitempty"`        // str Y	公司主页
	Email         bool `json:"email,omitempty"`          // str Y	电子邮件
	Office        bool `json:"office,omitempty"`         // str N	办公室
	Employees     bool `json:"employees,omitempty"`      // int64 Y	员工人数
	MainBusiness  bool `json:"main_business,omitempty"`  // str N	主要业务及产品
	BusinessScope bool `json:"business_scope,omitempty"` // str N	经营范围
}

func (item StockCompanyItems) All() StockCompanyItems {
	item.TsCode = true
	item.Exchange = true
	item.Chairman = true
	item.Manager = true
	item.Secretary = true
	item.RegCapital = true
	item.SetupDate = true
	item.Province = true
	item.City = true
	item.Website = true
	item.Email = true
	item.Office = true
	item.Employees = true
	item.MainBusiness = true
	item.BusinessScope = true
	return item
}

type StockCompanyData struct {
	TsCode        string  `json:"ts_code,omitempty"`        // str Y	股票代码
	Exchange      string  `json:"exchange,omitempty"`       // str Y	交易所代码 ,SSE上交所 SZSE深交所
	Chairman      string  `json:"chairman,omitempty"`       // str Y	法人代表
	Manager       string  `json:"manager,omitempty"`        // str Y	总经理
	Secretary     string  `json:"secretary,omitempty"`      // str Y	董秘
	RegCapital    float64 `json:"reg_capital,omitempty"`    // float64 Y	注册资本
	SetupDate     string  `json:"setup_date,omitempty"`     // str Y	注册日期
	Province      float64 `json:"province,omitempty"`       // flote64 Y	所在省份
	City          string  `json:"city,omitempty"`           // str Y	所在城市
	Introduction  string  `json:"introduction,omitempty"`   // str N	公司介绍
	Website       string  `json:"website,omitempty"`        // str Y	公司主页
	Email         string  `json:"email,omitempty"`          // str Y	电子邮件
	Office        string  `json:"office,omitempty"`         // str N	办公室
	Employees     int64   `json:"employees,omitempty"`      // int64 Y	员工人数
	MainBusiness  string  `json:"main_business,omitempty"`  // str N	主要业务及产品
	BusinessScope string  `json:"business_scope,omitempty"` // str N	经营范围
}

func AssembleStockCompanyData(tsRsp *TushareResponse) []*StockCompanyData {
	tsData := []*StockCompanyData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(StockCompanyData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 获取上市公司基础信息,单次提取4500条,可以根据交易所分批提取,用户需要至少120积分才可以调取,具体请参阅积分获取办法 https://tushare.pro/document/1?doc_id=13
func (ts *TuShare) StockCompany(params StockCompanyRequest, items StockCompanyItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "stock_company",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type NewShareCompanyRequest struct {
	StartDate string `json:"start_date,omitempty"` // Y	开始日期
	EndDate   string `json:"end_date,omitempty"`   // Y	结束日期
}

type NewShareCompanyItems struct {
	TsCode       bool `json:"ts_code,omitempty"`       // str	Y	TS股票代码
	SubCode      bool `json:"sub_code,omitempty"`      // str	Y	申购代码
	Name         bool `json:"name,omitempty"`          // str	Y	名称
	IpoDate      bool `json:"ipo_date,omitempty"`      // str	Y	上网发行日期
	IssueDate    bool `json:"issue_date,omitempty"`    // str	Y	上市日期
	Amount       bool `json:"amount,omitempty"`        // float	Y	发行总量(万股)
	MarketAmount bool `json:"market_amount,omitempty"` // float	Y	上网发行总量(万股)
	Price        bool `json:"price,omitempty"`         // float	Y	发行价格
	Pe           bool `json:"pe,omitempty"`            // float	Y	市盈率
	LimitAmount  bool `json:"limit_amount,omitempty"`  // float	Y	个人申购上限(万股)
	Funds        bool `json:"funds,omitempty"`         // float	Y	募集资金(亿元)
	Ballot       bool `json:"ballot,omitempty"`        // float	Y	中签率
}

func (item NewShareCompanyItems) All() NewShareCompanyItems {
	item.TsCode = true
	item.SubCode = true
	item.Name = true
	item.IpoDate = true
	item.IssueDate = true
	item.Amount = true
	item.MarketAmount = true
	item.Price = true
	item.Pe = true
	item.LimitAmount = true
	item.Funds = true
	item.Ballot = true
	return item
}

type NewShareCompanyData struct {
	TsCode       string  `json:"ts_code,omitempty"`       // str	Y	TS股票代码
	SubCode      string  `json:"sub_code,omitempty"`      // str	Y	申购代码
	Name         string  `json:"name,omitempty"`          // str	Y	名称
	IpoDate      string  `json:"ipo_date,omitempty"`      // str	Y	上网发行日期
	IssueDate    string  `json:"issue_date,omitempty"`    // str	Y	上市日期
	Amount       float64 `json:"amount,omitempty"`        // float	Y	发行总量(万股)
	MarketAmount float64 `json:"market_amount,omitempty"` // float	Y	上网发行总量(万股)
	Price        float64 `json:"price,omitempty"`         // float	Y	发行价格
	Pe           float64 `json:"pe,omitempty"`            // float	Y	市盈率
	LimitAmount  float64 `json:"limit_amount,omitempty"`  // float	Y	个人申购上限(万股)
	Funds        float64 `json:"funds,omitempty"`         // float	Y	募集资金(亿元)
	Ballot       float64 `json:"ballot,omitempty"`        // float	Y	中签率
}

func AssembleNewShareCompanyData(tsRsp *TushareResponse) []*NewShareCompanyData {
	tsData := []*NewShareCompanyData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(NewShareCompanyData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 获取新股上市列表数据,单次最大2000条,总量不限制,用户需要至少120积分才可以调取,具体请参阅积分获取办法 https://tushare.pro/document/1?doc_id=13
func (ts *TuShare) NewShare(params NewShareCompanyRequest, items NewShareCompanyItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "new_share",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

// 就是封装了trade_cal,测试tushare服务是否能连接
func (ts *TuShare) Health() (err error) {
	_, err = ts.TradeCal(TradeCalRequest{StartDate: Time2TushareDayTime(time.Now()), EndDate: Time2TushareDayTime(time.Now())}, TradeCalItems{}.All())
	return err
}

type STKManagersRequest struct {
	TsCode    string `json:"ts_code,omitempty"`    // str	N	股票代码,支持单个或多个股票输入 000001.SZ,600000.SH
	AnnDate   string `json:"STKRewards,omitempty"` // str	N	公告日期(YYYYMMDD格式,下同)
	StartDate string `json:"start_date,omitempty"` // str	N	公告开始日期
	EndDate   string `json:"end_date,omitempty"`   // str	N	公告结束日期
}

type STKManagersItems struct {
	TsCode    bool `json:"ts_code,omitempty"`    // str	Y	TS股票代码
	AnnDate   bool `json:"STKRewards,omitempty"` // str	Y	公告日期
	Name      bool `json:"name,omitempty"`       // str	Y	姓名
	Gender    bool `json:"gender,omitempty"`     // str	Y	性别
	Lev       bool `json:"lev,omitempty"`        // str	Y	岗位类别
	Title     bool `json:"title,omitempty"`      // str	Y	岗位
	Edu       bool `json:"edu,omitempty"`        // str	Y	学历
	National  bool `json:"national,omitempty"`   // str	Y	国籍
	Birthday  bool `json:"birthday,omitempty"`   // str	Y	出生年月
	BeginDate bool `json:"begin_date,omitempty"` // str	Y	上任日期
	EndDate   bool `json:"end_date,omitempty"`   // str	Y	离任日期
	Resume    bool `json:"resume,omitempty"`     // str	N	个人简历
}

func (item STKManagersItems) All() STKManagersItems {
	item.TsCode = true
	item.AnnDate = true
	item.Name = true
	item.Gender = true
	item.Lev = true
	item.Title = true
	item.Edu = true
	item.National = true
	item.Birthday = true
	item.BeginDate = true
	item.EndDate = true
	item.Resume = true
	return item
}

type STKManagersData struct {
	TsCode    string `json:"ts_code,omitempty"`    // str	Y	TS股票代码
	AnnDate   string `json:"STKRewards,omitempty"` // str	Y	公告日期
	Name      string `json:"name,omitempty"`       // str	Y	姓名
	Gender    string `json:"gender,omitempty"`     // str	Y	性别
	Lev       string `json:"lev,omitempty"`        // str	Y	岗位类别
	Title     string `json:"title,omitempty"`      // str	Y	岗位
	Edu       string `json:"edu,omitempty"`        // str	Y	学历
	National  string `json:"national,omitempty"`   // str	Y	国籍
	Birthday  string `json:"birthday,omitempty"`   // str	Y	出生年月
	BeginDate string `json:"begin_date,omitempty"` // str	Y	上任日期
	EndDate   string `json:"end_date,omitempty"`   // str	Y	离任日期
	Resume    string `json:"resume,omitempty"`     // str	N	个人简历
}

func AssembleSTKManagersData(tsRsp *TushareResponse) []*STKManagersData {
	tsData := []*STKManagersData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(STKManagersData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 获取上市公司管理层,用户需要至少2000积分才可以调取,具体请参阅积分获取办法 https://tushare.pro/document/1?doc_id=13
func (ts *TuShare) STKManagers(params STKManagersRequest, items STKManagersItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "stk_managers",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type STKRewardsRequest struct {
	TsCode  string `json:"ts_code,omitempty"`  // str	N	股票代码,支持单个或多个股票输入 000001.SZ,600000.SH
	EndDate string `json:"end_date,omitempty"` // str	N	报告期
}

type STKRewardsItems struct {
	TsCode  bool `json:"ts_code,omitempty"`  // str	Y	TS股票代码
	AnnDate bool `json:"ann_date,omitempty"` // str	Y	公告日期
	EndDate bool `json:"end_date,omitempty"` // str	Y	截止日期
	Name    bool `json:"name,omitempty"`     // str	Y	姓名
	Title   bool `json:"title,omitempty"`    // str	Y	职务
	Reward  bool `json:"reward,omitempty"`   // float	Y	报酬
	HoldVol bool `json:"hold_vol,omitempty"` // float	Y	持股数
}

func (item STKRewardsItems) All() STKRewardsItems {
	item.TsCode = true
	item.AnnDate = true
	item.EndDate = true
	item.Name = true
	item.Title = true
	item.Reward = true
	item.HoldVol = true
	return item
}

type STKRewardsData struct {
	TsCode  string  `json:"ts_code,omitempty"`  // str	Y	TS股票代码
	AnnDate string  `json:"ann_date,omitempty"` // str	Y	公告日期
	EndDate string  `json:"end_date,omitempty"` // str	Y	截止日期
	Name    string  `json:"name,omitempty"`     // str	Y	姓名
	Title   string  `json:"title,omitempty"`    // str	Y	职务
	Reward  float64 `json:"reward,omitempty"`   // float	Y	报酬
	HoldVol float64 `json:"hold_vol,omitempty"` // float	Y	持股数
}

func AssembleSTKRewardsData(tsRsp *TushareResponse) []*STKRewardsData {
	tsData := []*STKRewardsData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(STKRewardsData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 获取上市公司管理层薪酬和持股,用户需要至少2000积分才可以调取,具体请参阅积分获取办法 https://tushare.pro/document/1?doc_id=13
func (ts *TuShare) STKRewards(params STKRewardsRequest, items STKRewardsItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "stk_rewards",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}
