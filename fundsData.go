package tushare

import (
	"encoding/json"
	"net/http"
)

type FundBasicRequest struct {
	Market string `json:"market,omitempty"` // 	str	N	交易市场: E场内 O场外(默认E)
	Status string `json:"status,omitempty"` // 	str	N	存续状态 D摘牌 I发行 L上市中
}

type FundBasicItems struct {
	TsCode        bool `json:"ts_code,omitempty"`        // 	str	Y	基金代码
	Name          bool `json:"name,omitempty"`           // 	str	Y	简称
	Management    bool `json:"management,omitempty"`     // 	str	Y	管理人
	Custodian     bool `json:"custodian,omitempty"`      // 	str	Y	托管人
	FundType      bool `json:"fund_type,omitempty"`      // 	str	Y	投资类型
	FoundDate     bool `json:"found_date,omitempty"`     // 	str	Y	成立日期
	DueDate       bool `json:"due_date,omitempty"`       // 	str	Y	到期日期
	ListDate      bool `json:"list_date,omitempty"`      // 	str	Y	上市时间
	IssueDate     bool `json:"issue_date,omitempty"`     // 	str	Y	发行日期
	DelistDate    bool `json:"delist_date,omitempty"`    // 	str	Y	退市日期
	IssueAmount   bool `json:"issue_amount,omitempty"`   // 	float	Y	发行份额(亿)
	MFee          bool `json:"m_fee,omitempty"`          // 	float	Y	管理费
	CFee          bool `json:"c_fee,omitempty"`          // 	float	Y	托管费
	DurationYear  bool `json:"duration_year,omitempty"`  // 	float	Y	存续期
	PValue        bool `json:"p_value,omitempty"`        // 	float	Y	面值
	MinAmount     bool `json:"min_amount,omitempty"`     // 	float	Y	起点金额(万元)
	ExpReturn     bool `json:"exp_return,omitempty"`     // 	float	Y	预期收益率
	Benchmark     bool `json:"benchmark,omitempty"`      // 	str	Y	业绩比较基准
	Status        bool `json:"status,omitempty"`         // 	str	Y	存续状态D摘牌 I发行 L已上市
	InvestType    bool `json:"invest_type,omitempty"`    // 	str	Y	投资风格
	Type          bool `json:"type,omitempty"`           // 	str	Y	基金类型
	Trustee       bool `json:"trustee,omitempty"`        // 	str	Y	受托人
	PurcStartdate bool `json:"purc_startdate,omitempty"` // 	str	Y	日常申购起始日
	RedmStartdate bool `json:"redm_startdate,omitempty"` // 	str	Y	日常赎回起始日
	Market        bool `json:"market,omitempty"`         // 	str	Y	E场内O场外
}

func (item FundBasicItems) All() FundBasicItems {
	item.TsCode = true
	item.Name = true
	item.Management = true
	item.Custodian = true
	item.FundType = true
	item.FoundDate = true
	item.DueDate = true
	item.ListDate = true
	item.IssueDate = true
	item.DelistDate = true
	item.IssueAmount = true
	item.MFee = true
	item.CFee = true
	item.DurationYear = true
	item.PValue = true
	item.MinAmount = true
	item.ExpReturn = true
	item.Benchmark = true
	item.Status = true
	item.InvestType = true
	item.Type = true
	item.Trustee = true
	item.PurcStartdate = true
	item.RedmStartdate = true
	item.Market = true
	return item
}

type FundBasicData struct {
	TsCode        string  `json:"ts_code,omitempty" gorm:"uniqueIndex"` // 	str	Y	基金代码
	Name          string  `json:"name,omitempty"`                       // 	str	Y	简称
	Management    string  `json:"management,omitempty"`                 // 	str	Y	管理人
	Custodian     string  `json:"custodian,omitempty"`                  // 	str	Y	托管人
	FundType      string  `json:"fund_type,omitempty"`                  // 	str	Y	投资类型
	FoundDate     string  `json:"found_date,omitempty"`                 // 	str	Y	成立日期
	DueDate       string  `json:"due_date,omitempty"`                   // 	str	Y	到期日期
	ListDate      string  `json:"list_date,omitempty"`                  // 	str	Y	上市时间
	IssueDate     string  `json:"issue_date,omitempty"`                 // 	str	Y	发行日期
	DelistDate    string  `json:"delist_date,omitempty"`                // 	str	Y	退市日期
	IssueAmount   float64 `json:"issue_amount,omitempty"`               // 	float	Y	发行份额(亿)
	MFee          float64 `json:"m_fee,omitempty"`                      // 	float	Y	管理费
	CFee          float64 `json:"c_fee,omitempty"`                      // 	float	Y	托管费
	DurationYear  float64 `json:"duration_year,omitempty"`              // 	float	Y	存续期
	PValue        float64 `json:"p_value,omitempty"`                    // 	float	Y	面值
	MinAmount     float64 `json:"min_amount,omitempty"`                 // 	float	Y	起点金额(万元)
	ExpReturn     float64 `json:"exp_return,omitempty"`                 // 	float	Y	预期收益率
	Benchmark     string  `json:"benchmark,omitempty"`                  // 	str	Y	业绩比较基准
	Status        string  `json:"status,omitempty" gorm:"index"`        // 	str	Y	存续状态D摘牌 I发行 L已上市
	InvestType    string  `json:"invest_type,omitempty"`                // 	str	Y	投资风格
	Type          string  `json:"type,omitempty"`                       // 	str	Y	基金类型
	Trustee       string  `json:"trustee,omitempty"`                    // 	str	Y	受托人
	PurcStartdate string  `json:"purc_startdate,omitempty"`             // 	str	Y	日常申购起始日
	RedmStartdate string  `json:"redm_startdate,omitempty"`             // 	str	Y	日常赎回起始日
	Market        string  `json:"market,omitempty"`                     // 	str	Y	E场内O场外
}

func AssembleFundBasicData(tsRsp *TushareResponse) []*FundBasicData {
	tsData := []*FundBasicData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(FundBasicData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 获取公募基金数据列表,包括场内和场外基金,用户需要至少1500积分才可以调取,具体请参阅积分获取办法 https://tushare.pro/document/1?doc_id=13
func (ts *TuShare) FundBasic(params FundBasicRequest, items FundBasicItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "fund_basic",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type FundCompanyItems struct {
	Name         bool `json:"name,omitempty"`          // 	str	Y	基金公司名称
	Shortname    bool `json:"shortname,omitempty"`     // 	str	Y	简称
	ShortEnname  bool `json:"short_enname,omitempty"`  // 	str	N	英文缩写
	Province     bool `json:"province,omitempty"`      // 	str	Y	省份
	City         bool `json:"city,omitempty"`          // 	str	Y	城市
	Address      bool `json:"address,omitempty"`       // 	str	Y	注册地址
	Phone        bool `json:"phone,omitempty"`         // 	str	Y	电话
	Office       bool `json:"office,omitempty"`        // 	str	Y	办公地址
	Website      bool `json:"website,omitempty"`       // 	str	Y	公司网址
	Chairman     bool `json:"chairman,omitempty"`      // 	str	Y	法人代表
	Manager      bool `json:"manager,omitempty"`       // 	str	Y	总经理
	RegCapital   bool `json:"reg_capital,omitempty"`   // 	float	Y	注册资本
	SetupDate    bool `json:"setup_date,omitempty"`    // 	str	Y	成立日期
	EndDate      bool `json:"end_date,omitempty"`      // 	str	Y	公司终止日期
	Employees    bool `json:"employees,omitempty"`     // 	float	Y	员工总数
	MainBusiness bool `json:"main_business,omitempty"` // 	str	Y	主要产品及业务
	OrgCode      bool `json:"org_code,omitempty"`      // 	str	Y	组织机构代码
	CreditCode   bool `json:"credit_code,omitempty"`   // 	str	Y	统一社会信用代码
}

func (item FundCompanyItems) All() FundCompanyItems {
	item.Name = true
	item.Shortname = true
	item.ShortEnname = true
	item.Province = true
	item.City = true
	item.Address = true
	item.Phone = true
	item.Office = true
	item.Website = true
	item.Chairman = true
	item.Manager = true
	item.RegCapital = true
	item.SetupDate = true
	item.EndDate = true
	item.Employees = true
	item.MainBusiness = true
	item.OrgCode = true
	item.CreditCode = true
	return item
}

type FundCompanyData struct {
	Name         string  `json:"name,omitempty"`          // 	str	Y	基金公司名称
	Shortname    string  `json:"shortname,omitempty"`     // 	str	Y	简称
	ShortEnname  string  `json:"short_enname,omitempty"`  // 	str	N	英文缩写
	Province     string  `json:"province,omitempty"`      // 	str	Y	省份
	City         string  `json:"city,omitempty"`          // 	str	Y	城市
	Address      string  `json:"address,omitempty"`       // 	str	Y	注册地址
	Phone        string  `json:"phone,omitempty"`         // 	str	Y	电话
	Office       string  `json:"office,omitempty"`        // 	str	Y	办公地址
	Website      string  `json:"website,omitempty"`       // 	str	Y	公司网址
	Chairman     string  `json:"chairman,omitempty"`      // 	str	Y	法人代表
	Manager      string  `json:"manager,omitempty"`       // 	str	Y	总经理
	RegCapital   float64 `json:"reg_capital,omitempty"`   // 	float	Y	注册资本
	SetupDate    string  `json:"setup_date,omitempty"`    // 	str	Y	成立日期
	EndDate      string  `json:"end_date,omitempty"`      // 	str	Y	公司终止日期
	Employees    float64 `json:"employees,omitempty"`     // 	float	Y	员工总数
	MainBusiness string  `json:"main_business,omitempty"` // 	str	Y	主要产品及业务
	OrgCode      string  `json:"org_code,omitempty"`      // 	str	Y	组织机构代码
	CreditCode   string  `json:"credit_code,omitempty"`   // 	str	Y	统一社会信用代码
}

func AssembleFundCompanyData(tsRsp *TushareResponse) []*FundCompanyData {
	tsData := []*FundCompanyData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(FundCompanyData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 获取公募基金管理人列表,用户需要至少1500积分才可以调取,具体请参阅积分获取办法 https://tushare.pro/document/1?doc_id=13
func (ts *TuShare) FundCompany(items FundCompanyItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "fund_company",
		Token:   ts.token,
		Params:  make(map[string]interface{}),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type FundManagerRequest struct {
	TsCode  string `json:"ts_code,omitempty"`  // 	str	N	基金代码,支持多只基金,逗号分隔
	AnnDate string `json:"ann_date,omitempty"` // 	str	N	公告日期,格式：YYYYMMDD
	Name    string `json:"name,omitempty"`     // 	str	N	基金经理姓名
	Offset  int64  `json:"offset,omitempty"`   // 	int	N	开始行数
	Limit   int64  `json:"limit,omitempty"`    // 	int	N	每页行数
}

type FundManagerItems struct {
	TsCode      bool `json:"ts_code,omitempty"`     // 	str	Y	基金代码
	AnnDate     bool `json:"ann_date,omitempty"`    // 	str	Y	公告日期
	Name        bool `json:"name,omitempty"`        // 	str	Y	基金经理姓名
	Gender      bool `json:"gender,omitempty"`      // 	str	Y	性别
	BirthYear   bool `json:"birth_year,omitempty"`  // 	str	Y	出生年份
	Edu         bool `json:"edu,omitempty"`         // 	str	Y	学历
	Nationality bool `json:"nationality,omitempty"` // 	str	Y	国籍
	BeginDate   bool `json:"begin_date,omitempty"`  // 	str	Y	任职日期
	EndDate     bool `json:"end_date,omitempty"`    // 	str	Y	离任日期
	Resume      bool `json:"resume,omitempty"`      // 	str	Y	简历
}

func (item FundManagerItems) All() FundManagerItems {
	item.TsCode = true
	item.AnnDate = true
	item.Name = true
	item.Gender = true
	item.BirthYear = true
	item.Edu = true
	item.Nationality = true
	item.BeginDate = true
	item.EndDate = true
	item.Resume = true
	return item
}

type FundManagerData struct {
	TsCode      string `json:"ts_code,omitempty"`     // 	str	Y	基金代码
	AnnDate     string `json:"ann_date,omitempty"`    // 	str	Y	公告日期
	Name        string `json:"name,omitempty"`        // 	str	Y	基金经理姓名
	Gender      string `json:"gender,omitempty"`      // 	str	Y	性别
	BirthYear   string `json:"birth_year,omitempty"`  // 	str	Y	出生年份
	Edu         string `json:"edu,omitempty"`         // 	str	Y	学历
	Nationality string `json:"nationality,omitempty"` // 	str	Y	国籍
	BeginDate   string `json:"begin_date,omitempty"`  // 	str	Y	任职日期
	EndDate     string `json:"end_date,omitempty"`    // 	str	Y	离任日期
	Resume      string `json:"resume,omitempty"`      // 	str	Y	简历
}

func AssembleFundManagerData(tsRsp *TushareResponse) []*FundManagerData {
	tsData := []*FundManagerData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(FundManagerData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 获取公募基金经理数据,包括基金经理简历等数据,单次最大5000,支持分页提取数据,用户需要至少500积分才可以调取,具体请参阅积分获取办法 https://tushare.pro/document/1?doc_id=13
func (ts *TuShare) FundManager(params FundManagerRequest, items FundManagerItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "fund_manager",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type FundShareRequest struct {
	TsCode    string `json:"ts_code,omitempty"`    // 	str	N	TS基金代码
	TradeDate string `json:"trade_date,omitempty"` // 	str	N	交易日期
	StartDate string `json:"start_date,omitempty"` // 	str	N	开始日期
	EndDate   string `json:"end_date,omitempty"`   // 	str	N	结束日期
	FundType  string `json:"fund_type,omitempty"`  // 	str	N	基金类型,见下表
	Market    string `json:"market,omitempty"`     // 	str	N	市场：SH/SZ
}

type FundShareItems struct {
	TsCode    bool `json:"ts_code,omitempty"`    // 	str	Y	基金代码,支持多只基金同时提取,用逗号分隔
	TradeDate bool `json:"trade_date,omitempty"` // 	str	Y	交易(变动)日期,格式YYYYMMDD
	FdShare   bool `json:"fd_share,omitempty"`   // 	float	Y	基金份额(万)
}

func (item FundShareItems) All() FundShareItems {
	item.TsCode = true
	item.TradeDate = true
	item.FdShare = true
	return item
}

// 标识	含义
// ETF	ETF基金
// LOF	LOF基金
// SF	分级基金
// CEF	封闭基金
type FundShareData struct {
	TsCode    string  `json:"ts_code,omitempty"`    // 	str	Y	基金代码,支持多只基金同时提取,用逗号分隔
	TradeDate string  `json:"trade_date,omitempty"` // 	str	Y	交易(变动)日期,格式YYYYMMDD
	FdShare   float64 `json:"fd_share,omitempty"`   // 	float	Y	基金份额(万)
}

func AssembleFundShareData(tsRsp *TushareResponse) []*FundShareData {
	tsData := []*FundShareData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(FundShareData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 获取基金规模数据,包含上海和深圳ETF基金,单次最大提取2000行数据,用户需要至少2000积分才可以调取,具体请参阅积分获取办法 https://tushare.pro/document/1?doc_id=13
func (ts *TuShare) FundShare(params FundShareRequest, items FundShareItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "fund_share",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type FundNavRequest struct {
	TsCode  string `json:"ts_code,omitempty"`  // 	str	N	TS基金代码 (二选一)
	EndDate string `json:"end_date,omitempty"` // 	str	N	净值日期 (二选一)
	Market  string `json:"market,omitempty"`   // 	str	N	E场内 O场外
}

type FundNavItems struct {
	TsCode        bool `json:"ts_code,omitempty"`        // 	str	Y	TS代码
	AnnDate       bool `json:"ann_date,omitempty"`       // 	str	Y	公告日期
	EndDate       bool `json:"end_date,omitempty"`       // 	str	Y	截止日期
	UnitNav       bool `json:"unit_nav,omitempty"`       // 	float	Y	单位净值
	AccumNav      bool `json:"accum_nav,omitempty"`      // 	float	Y	累计净值
	AccumDiv      bool `json:"accum_div,omitempty"`      // 	float	Y	累计分红
	NetAsset      bool `json:"net_asset,omitempty"`      // 	float	Y	资产净值
	TotalNetasset bool `json:"total_netasset,omitempty"` // 	float	Y	合计资产净值
	AdjNav        bool `json:"adj_nav,omitempty"`        // 	float	Y	复权单位净值
}

func (item FundNavItems) All() FundNavItems {
	item.TsCode = true
	item.AnnDate = true
	item.EndDate = true
	item.UnitNav = true
	item.AccumNav = true
	item.AccumDiv = true
	item.NetAsset = true
	item.TotalNetasset = true
	item.AdjNav = true
	return item
}

type FundNavData struct {
	TsCode        string  `json:"ts_code,omitempty"`        // 	str	Y	TS代码
	AnnDate       string  `json:"ann_date,omitempty"`       // 	str	Y	公告日期
	EndDate       string  `json:"end_date,omitempty"`       // 	str	Y	截止日期
	UnitNav       float64 `json:"unit_nav,omitempty"`       // 	float	Y	单位净值
	AccumNav      float64 `json:"accum_nav,omitempty"`      // 	float	Y	累计净值
	AccumDiv      float64 `json:"accum_div,omitempty"`      // 	float	Y	累计分红
	NetAsset      float64 `json:"net_asset,omitempty"`      // 	float	Y	资产净值
	TotalNetasset float64 `json:"total_netasset,omitempty"` // 	float	Y	合计资产净值
	AdjNav        float64 `json:"adj_nav,omitempty"`        // 	float	Y	复权单位净值
}

func AssembleFundNavData(tsRsp *TushareResponse) []*FundNavData {
	tsData := []*FundNavData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(FundNavData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 获取公募基金净值数据,用户需要至少2000积分才可以调取,具体请参阅积分获取办法 https://tushare.pro/document/1?doc_id=13
func (ts *TuShare) FundNav(params FundNavRequest, items FundNavItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "fund_nav",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type FundDivRequest struct {
	AnnDate string `json:"ann_date,omitempty"` // 	str	N	公告日(以下参数四选一)
	ExDate  string `json:"ex_date,omitempty"`  // 	str	N	除息日
	PayDate string `json:"pay_date,omitempty"` // 	str	N	派息日
	TsCode  string `json:"ts_code,omitempty"`  // 	str	N	基金代码
}

type FundDivItems struct {
	TsCode      bool `json:"ts_code,omitempty"`      // 	str	Y	TS代码
	AnnDate     bool `json:"ann_date,omitempty"`     // 	str	Y	公告日期
	ImpAnndate  bool `json:"imp_anndate,omitempty"`  // 	str	Y	分红实施公告日
	BaseDate    bool `json:"base_date,omitempty"`    // 	str	Y	分配收益基准日
	DivProc     bool `json:"div_proc,omitempty"`     // 	str	Y	方案进度
	RecordDate  bool `json:"record_date,omitempty"`  // 	str	Y	权益登记日
	ExDate      bool `json:"ex_date,omitempty"`      // 	str	Y	除息日
	PayDate     bool `json:"pay_date,omitempty"`     // 	str	Y	派息日
	EarpayDate  bool `json:"earpay_date,omitempty"`  // 	str	Y	收益支付日
	NetExDate   bool `json:"net_ex_date,omitempty"`  // 	str	Y	净值除权日
	DivCash     bool `json:"div_cash,omitempty"`     // 	float	Y	每股派息(元)
	BaseUnit    bool `json:"base_unit,omitempty"`    // 	float	Y	基准基金份额(万份)
	EarDistr    bool `json:"ear_distr,omitempty"`    // 	float	Y	可分配收益(元)
	EarAmount   bool `json:"ear_amount,omitempty"`   // 	float	Y	收益分配金额(元)
	AccountDate bool `json:"account_date,omitempty"` // 	str	Y	红利再投资到账日
	BaseYear    bool `json:"base_year,omitempty"`    // 	str	Y	份额基准年度
}

func (item FundDivItems) All() FundDivItems {
	item.TsCode = true
	item.AnnDate = true
	item.ImpAnndate = true
	item.BaseDate = true
	item.DivProc = true
	item.RecordDate = true
	item.ExDate = true
	item.PayDate = true
	item.EarpayDate = true
	item.NetExDate = true
	item.DivCash = true
	item.BaseUnit = true
	item.EarDistr = true
	item.EarAmount = true
	item.AccountDate = true
	item.BaseYear = true
	return item
}

type FundDivData struct {
	TsCode      string  `json:"ts_code,omitempty"`      // 	str	Y	TS代码
	AnnDate     string  `json:"ann_date,omitempty"`     // 	str	Y	公告日期
	ImpAnndate  string  `json:"imp_anndate,omitempty"`  // 	str	Y	分红实施公告日
	BaseDate    string  `json:"base_date,omitempty"`    // 	str	Y	分配收益基准日
	DivProc     string  `json:"div_proc,omitempty"`     // 	str	Y	方案进度
	RecordDate  string  `json:"record_date,omitempty"`  // 	str	Y	权益登记日
	ExDate      string  `json:"ex_date,omitempty"`      // 	str	Y	除息日
	PayDate     string  `json:"pay_date,omitempty"`     // 	str	Y	派息日
	EarpayDate  string  `json:"earpay_date,omitempty"`  // 	str	Y	收益支付日
	NetExDate   string  `json:"net_ex_date,omitempty"`  // 	str	Y	净值除权日
	DivCash     float64 `json:"div_cash,omitempty"`     // 	float	Y	每股派息(元)
	BaseUnit    float64 `json:"base_unit,omitempty"`    // 	float	Y	基准基金份额(万份)
	EarDistr    float64 `json:"ear_distr,omitempty"`    // 	float	Y	可分配收益(元)
	EarAmount   float64 `json:"ear_amount,omitempty"`   // 	float	Y	收益分配金额(元)
	AccountDate string  `json:"account_date,omitempty"` // 	str	Y	红利再投资到账日
	BaseYear    string  `json:"base_year,omitempty"`    // 	str	Y	份额基准年度
}

func AssembleFundDivData(tsRsp *TushareResponse) []*FundDivData {
	tsData := []*FundDivData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(FundDivData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 获取公募基金分红数据,用户需要至少400积分才可以调取,具体请参阅积分获取办法 https://tushare.pro/document/1?doc_id=13
func (ts *TuShare) FundDiv(params FundDivRequest, items FundDivItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "fund_div",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type FundPortfolioRequest struct {
	TsCode    string `json:"ts_code,omitempty"`    // 	str	Y	基金代码
	AnnDate   string `json:"ann_date,omitempty"`   // 	str	N	公告日期(YYYYMMDD格式)
	StartDate string `json:"start_date,omitempty"` // 	str	N	报告期开始日期(YYYYMMDD格式)
	EndDate   string `json:"end_date,omitempty"`   // 	str	N	报告期结束日期(YYYYMMDD格式)
}

type FundPortfolioItems struct {
	TsCode        bool `json:"ts_code,omitempty"`         // 	str	Y	TS基金代码
	AnnDate       bool `json:"ann_date,omitempty"`        // 	str	Y	公告日期
	EndDate       bool `json:"end_date,omitempty"`        // 	str	Y	截止日期
	Symbol        bool `json:"symbol,omitempty"`          // 	str	Y	股票代码
	Mkv           bool `json:"mkv,omitempty"`             // 	float	Y	持有股票市值(元)
	Amount        bool `json:"amount,omitempty"`          // 	float	Y	持有股票数量(股)
	StkMkvRatio   bool `json:"stk_mkv_ratio,omitempty"`   // 	float	Y	占股票市值比
	StkFloatRatio bool `json:"stk_float_ratio,omitempty"` // 	float	Y	占流通股本比例
}

func (item FundPortfolioItems) All() FundPortfolioItems {
	item.TsCode = true
	item.AnnDate = true
	item.EndDate = true
	item.Symbol = true
	item.Mkv = true
	item.Amount = true
	item.StkMkvRatio = true
	item.StkFloatRatio = true
	return item
}

type FundPortfolioData struct {
	TsCode        string  `json:"ts_code,omitempty" gorm:"uniqueIndex:fcode"` // 	str	Y	TS基金代码
	AnnDate       string  `json:"ann_date,omitempty"`                         // 	str	Y	公告日期
	EndDate       string  `json:"end_date,omitempty"`                         // 	str	Y	截止日期
	Symbol        string  `json:"symbol,omitempty" gorm:"uniqueIndex:fcode"`  // 	str	Y	股票代码
	Mkv           float64 `json:"mkv,omitempty"`                              // 	float	Y	持有股票市值(元)
	Amount        float64 `json:"amount,omitempty"`                           // 	float	Y	持有股票数量(股)
	StkMkvRatio   float64 `json:"stk_mkv_ratio,omitempty"`                    // 	float	Y	占股票市值比
	StkFloatRatio float64 `json:"stk_float_ratio,omitempty"`                  // 	float	Y	占流通股本比例
}

func AssembleFundPortfolioData(tsRsp *TushareResponse) []*FundPortfolioData {
	tsData := []*FundPortfolioData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(FundPortfolioData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 获取公募基金持仓数据,季度更新,用户需要至少2000积分才可以调取,具体请参阅积分获取办法 https://tushare.pro/document/1?doc_id=13
func (ts *TuShare) FundPortfolio(params FundPortfolioRequest, items FundPortfolioItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "fund_portfolio",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type FundDailyItems struct {
	TsCode    bool `json:"ts_code,omitempty"`    // 	str	Y	TS代码
	TradeDate bool `json:"trade_date,omitempty"` // 	str	Y	交易日期
	Open      bool `json:"open,omitempty"`       // 	float	Y	开盘价(元)
	High      bool `json:"high,omitempty"`       // 	float	Y	最高价(元)
	Low       bool `json:"low,omitempty"`        // 	float	Y	最低价(元)
	Close     bool `json:"close,omitempty"`      // 	float	Y	收盘价(元)
	PreClose  bool `json:"pre_close,omitempty"`  // 	float	Y	昨收盘价(元)
	Change    bool `json:"change,omitempty"`     // 	float	Y	涨跌额(元)
	PctChg    bool `json:"pct_chg,omitempty"`    // 	float	Y	涨跌幅(%)
	Vol       bool `json:"vol,omitempty"`        // 	float	Y	成交量(手)
	Amount    bool `json:"amount,omitempty"`     // 	float	Y	成交额(千元)
}

func (item FundDailyItems) All() FundDailyItems {
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

type FundDailyData struct {
	TsCode    string  `json:"ts_code,omitempty"`    // 	str	Y	TS代码
	TradeDate string  `json:"trade_date,omitempty"` // 	str	Y	交易日期
	Open      float64 `json:"open,omitempty"`       // 	float	Y	开盘价(元)
	High      float64 `json:"high,omitempty"`       // 	float	Y	最高价(元)
	Low       float64 `json:"low,omitempty"`        // 	float	Y	最低价(元)
	Close     float64 `json:"close,omitempty"`      // 	float	Y	收盘价(元)
	PreClose  float64 `json:"pre_close,omitempty"`  // 	float	Y	昨收盘价(元)
	Change    float64 `json:"change,omitempty"`     // 	float	Y	涨跌额(元)
	PctChg    float64 `json:"pct_chg,omitempty"`    // 	float	Y	涨跌幅(%)
	Vol       float64 `json:"vol,omitempty"`        // 	float	Y	成交量(手)
	Amount    float64 `json:"amount,omitempty"`     // 	float	Y	成交额(千元)
}

func AssembleFundDailyData(tsRsp *TushareResponse) []*FundDailyData {
	tsData := []*FundDailyData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(FundDailyData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 获取场内基金日线行情,类似股票日行情,每日收盘后2小时内更新,单次最大800行记录,总量不限制,用户需要至少2000积分才可以调取,具体请参阅积分获取办法 https://tushare.pro/document/1?doc_id=13
func (ts *TuShare) FundDaily(params QuotationRequest, items FundDailyItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "fund_daily",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type FundAdjRequest struct {
	TsCode    string `json:"ts_code,omitempty"`    // 	str	N	TS基金代码(支持多只基金输入)
	TradeDate string `json:"trade_date,omitempty"` // 	str	N	交易日期(格式：yyyymmdd,下同)
	StartDate string `json:"start_date,omitempty"` // 	str	N	开始日期
	EndDate   string `json:"end_date,omitempty"`   // 	str	N	结束日期
	Offset    string `json:"offset,omitempty"`     // 	str	N	开始行数
	Limit     string `json:"limit,omitempty"`      // 	str	N	最大行数
}

type FundAdjItems struct {
	TsCode    bool `json:"ts_code,omitempty"`    // 	str	Y	ts基金代码
	TradeDate bool `json:"trade_date,omitempty"` // 	str	Y	交易日期
	AdjFactor bool `json:"adj_factor,omitempty"` // 	float	Y	复权因子
}

func (item FundAdjItems) All() FundAdjItems {
	item.TsCode = true
	item.TradeDate = true
	item.AdjFactor = true
	return item
}

type FundAdjData struct {
	TsCode    string  `json:"ts_code,omitempty"`    // 	str	Y	ts基金代码
	TradeDate string  `json:"trade_date,omitempty"` // 	str	Y	交易日期
	AdjFactor float64 `json:"adj_factor,omitempty"` // 	float	Y	复权因子
}

func AssembleFundAdjData(tsRsp *TushareResponse) []*FundAdjData {
	tsData := []*FundAdjData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(FundAdjData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 获取基金复权因子,用于计算基金复权行情,单次最大提取2000行记录,可循环提取,数据总量不限制,用户需要至少600积分才可以调取,具体请参阅积分获取办法 https://tushare.pro/document/1?doc_id=13
func (ts *TuShare) FundAdj(params FundAdjRequest, items FundAdjItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "fund_adj",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}
