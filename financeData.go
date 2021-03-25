package tushare

import (
	"encoding/json"
	"net/http"
)

// 主要报表类型说明
// 代码	类型	说明
// 1	合并报表	上市公司最新报表(默认)
// 2	单季合并	单一季度的合并报表
// 3	调整单季合并表	调整后的单季合并报表(如果有)
// 4	调整合并报表	本年度公布上年同期的财务报表数据,报告期为上年度
// 5	调整前合并报表	数据发生变更,将原数据进行保留,即调整前的原数据
// 6	母公司报表	该公司母公司的财务报表数据
// 7	母公司单季表	母公司的单季度表
// 8	母公司调整单季表	母公司调整后的单季表
// 9	母公司调整表	该公司母公司的本年度公布上年同期的财务报表数据
// 10	母公司调整前报表	母公司调整之前的原始财务报表数据
// 11	调整前合并报表	调整之前合并报表原数据
// 12	母公司调整前报表	母公司报表发生变更前保留的原数据
type SheetRequest struct {
	TsCode     string `json:"ts_code,omitempty"`     // str	Y	股票代码
	AnnDate    string `json:"ann_date,omitempty"`    // str	N	公告日期
	StartDate  string `json:"start_date,omitempty"`  // str	N	公告开始日期
	EndDate    string `json:"end_date,omitempty"`    // str	N	公告结束日期
	Period     string `json:"period,omitempty"`      // str	N	报告期(每个季度最后一天的日期,比如20171231表示年报)
	ReportType string `json:"report_type,omitempty"` // str	N	报告类型： 参考下表说明
	CompType   string `json:"comp_type,omitempty"`   // str	N	公司类型：1一般工商业 2银行 3保险 4证券
}

type IncomeItems struct {
	TsCode            bool `json:"ts_code,omitempty"`             // str	Y	TS代码
	AnnDate           bool `json:"ann_date,omitempty"`            // str	Y	公告日期
	FAnnDate          bool `json:"f_ann_date,omitempty"`          // str	Y	实际公告日期
	EndDate           bool `json:"end_date,omitempty"`            // str	Y	报告期
	ReportType        bool `json:"report_type,omitempty"`         // str	Y	报告类型 1合并报表 2单季合并 3调整单季合并表 4调整合并报表 5调整前合并报表 6母公司报表 7母公司单季表 8 母公司调整单季表 9母公司调整表 10母公司调整前报表 11调整前合并报表 12母公司调整前报表
	CompType          bool `json:"comp_type,omitempty"`           // str	Y	公司类型(1一般工商业2银行3保险4证券)
	BasicEps          bool `json:"basic_eps,omitempty"`           // float	Y	基本每股收益
	DilutedEps        bool `json:"diluted_eps,omitempty"`         // float	Y	稀释每股收益
	TotalRevenue      bool `json:"total_revenue,omitempty"`       // float	Y	营业总收入
	Revenue           bool `json:"revenue,omitempty"`             // float	Y	营业收入
	IntIncome         bool `json:"int_income,omitempty"`          // float	Y	利息收入
	PremEarned        bool `json:"prem_earned,omitempty"`         // float	Y	已赚保费
	CommIncome        bool `json:"comm_income,omitempty"`         // float	Y	手续费及佣金收入
	NCommisIncome     bool `json:"n_commis_income,omitempty"`     // float	Y	手续费及佣金净收入
	NOthIncome        bool `json:"n_oth_income,omitempty"`        // float	Y	其他经营净收益
	NOthBIncome       bool `json:"n_oth_b_income,omitempty"`      // float	Y	加:其他业务净收益
	PremIncome        bool `json:"prem_income,omitempty"`         // float	Y	保险业务收入
	OutPrem           bool `json:"out_prem,omitempty"`            // float	Y	减:分出保费
	UnePremReser      bool `json:"une_prem_reser,omitempty"`      // float	Y	提取未到期责任准备金
	ReinsIncome       bool `json:"reins_income,omitempty"`        // float	Y	其中:分保费收入
	NSecTbIncome      bool `json:"n_sec_tb_income,omitempty"`     // float	Y	代理买卖证券业务净收入
	NSecUwIncome      bool `json:"n_sec_uw_income,omitempty"`     // float	Y	证券承销业务净收入
	NAssetMgIncome    bool `json:"n_asset_mg_income,omitempty"`   // float	Y	受托客户资产管理业务净收入
	OthBIncome        bool `json:"oth_b_income,omitempty"`        // float	Y	其他业务收入
	FvValueChgGain    bool `json:"fv_value_chg_gain,omitempty"`   // float	Y	加:公允价值变动净收益
	InvestIncome      bool `json:"invest_income,omitempty"`       // float	Y	加:投资净收益
	AssInvestIncome   bool `json:"ass_invest_income,omitempty"`   // float	Y	其中:对联营企业和合营企业的投资收益
	ForexGain         bool `json:"forex_gain,omitempty"`          // float	Y	加:汇兑净收益
	TotalCogs         bool `json:"total_cogs,omitempty"`          // float	Y	营业总成本
	OperCost          bool `json:"oper_cost,omitempty"`           // float	Y	减:营业成本
	IntExp            bool `json:"int_exp,omitempty"`             // float	Y	减:利息支出
	CommExp           bool `json:"comm_exp,omitempty"`            // float	Y	减:手续费及佣金支出
	BizTaxSurchg      bool `json:"biz_tax_surchg,omitempty"`      // float	Y	减:营业税金及附加
	SellExp           bool `json:"sell_exp,omitempty"`            // float	Y	减:销售费用
	AdminExp          bool `json:"admin_exp,omitempty"`           // float	Y	减:管理费用
	FinExp            bool `json:"fin_exp,omitempty"`             // float	Y	减:财务费用
	AssetsImpairLoss  bool `json:"assets_impair_loss,omitempty"`  // float	Y	减:资产减值损失
	PremRefund        bool `json:"prem_refund,omitempty"`         // float	Y	退保金
	CompensPayout     bool `json:"compens_payout,omitempty"`      // float	Y	赔付总支出
	ReserInsurLiab    bool `json:"reser_insur_liab,omitempty"`    // float	Y	提取保险责任准备金
	DivPayt           bool `json:"div_payt,omitempty"`            // float	Y	保户红利支出
	ReinsExp          bool `json:"reins_exp,omitempty"`           // float	Y	分保费用
	OperExp           bool `json:"oper_exp,omitempty"`            // float	Y	营业支出
	CompensPayoutRefu bool `json:"compens_payout_refu,omitempty"` // float	Y	减:摊回赔付支出
	InsurReserRefu    bool `json:"insur_reser_refu,omitempty"`    // float	Y	减:摊回保险责任准备金
	ReinsCostRefund   bool `json:"reins_cost_refund,omitempty"`   // float	Y	减:摊回分保费用
	OtherBusCost      bool `json:"other_bus_cost,omitempty"`      // float	Y	其他业务成本
	OperateProfit     bool `json:"operate_profit,omitempty"`      // float	Y	营业利润
	NonOperIncome     bool `json:"non_oper_income,omitempty"`     // float	Y	加:营业外收入
	NonOperExp        bool `json:"non_oper_exp,omitempty"`        // float	Y	减:营业外支出
	NcaDisploss       bool `json:"nca_disploss,omitempty"`        // float	Y	其中:减:非流动资产处置净损失
	TotalProfit       bool `json:"total_profit,omitempty"`        // float	Y	利润总额
	IncomeTax         bool `json:"income_tax,omitempty"`          // float	Y	所得税费用
	NIncome           bool `json:"n_income,omitempty"`            // float	Y	净利润(含少数股东损益)
	NIncomeAttrP      bool `json:"n_income_attr_p,omitempty"`     // float	Y	净利润(不含少数股东损益)
	MinorityGain      bool `json:"minority_gain,omitempty"`       // float	Y	少数股东损益
	OthComprIncome    bool `json:"oth_compr_income,omitempty"`    // float	Y	其他综合收益
	TComprIncome      bool `json:"t_compr_income,omitempty"`      // float	Y	综合收益总额
	ComprIncAttrP     bool `json:"compr_inc_attr_p,omitempty"`    // float	Y	归属于母公司(或股东)的综合收益总额
	ComprIncAttrMS    bool `json:"compr_inc_attr_m_s,omitempty"`  // float	Y	归属于少数股东的综合收益总额
	Ebit              bool `json:"ebit,omitempty"`                // float	Y	息税前利润
	Ebitda            bool `json:"ebitda,omitempty"`              // float	Y	息税折旧摊销前利润
	InsuranceExp      bool `json:"insurance_exp,omitempty"`       // float	Y	保险业务支出
	UndistProfit      bool `json:"undist_profit,omitempty"`       // float	Y	年初未分配利润
	DistableProfit    bool `json:"distable_profit,omitempty"`     // float	Y	可分配利润
	UpdateFlag        bool `json:"update_flag,omitempty"`         // str	N	更新标识,0未修改1更正过
}

func (item IncomeItems) All() IncomeItems {
	item.TsCode = true
	item.AnnDate = true
	item.FAnnDate = true
	item.EndDate = true
	item.ReportType = true
	item.CompType = true
	item.BasicEps = true
	item.DilutedEps = true
	item.TotalRevenue = true
	item.Revenue = true
	item.IntIncome = true
	item.PremEarned = true
	item.CommIncome = true
	item.NCommisIncome = true
	item.NOthIncome = true
	item.NOthBIncome = true
	item.PremIncome = true
	item.OutPrem = true
	item.UnePremReser = true
	item.ReinsIncome = true
	item.NSecTbIncome = true
	item.NSecUwIncome = true
	item.NAssetMgIncome = true
	item.OthBIncome = true
	item.FvValueChgGain = true
	item.InvestIncome = true
	item.AssInvestIncome = true
	item.ForexGain = true
	item.TotalCogs = true
	item.OperCost = true
	item.IntExp = true
	item.CommExp = true
	item.BizTaxSurchg = true
	item.SellExp = true
	item.AdminExp = true
	item.FinExp = true
	item.AssetsImpairLoss = true
	item.PremRefund = true
	item.CompensPayout = true
	item.ReserInsurLiab = true
	item.DivPayt = true
	item.ReinsExp = true
	item.OperExp = true
	item.CompensPayoutRefu = true
	item.InsurReserRefu = true
	item.ReinsCostRefund = true
	item.OtherBusCost = true
	item.OperateProfit = true
	item.NonOperIncome = true
	item.NonOperExp = true
	item.NcaDisploss = true
	item.TotalProfit = true
	item.IncomeTax = true
	item.NIncome = true
	item.NIncomeAttrP = true
	item.MinorityGain = true
	item.OthComprIncome = true
	item.TComprIncome = true
	item.ComprIncAttrP = true
	item.ComprIncAttrMS = true
	item.Ebit = true
	item.Ebitda = true
	item.InsuranceExp = true
	item.UndistProfit = true
	item.DistableProfit = true
	item.UpdateFlag = true
	return item
}

type IncomeData struct {
	TsCode            string  `json:"ts_code,omitempty"`             // str	Y	TS代码
	AnnDate           string  `json:"ann_date,omitempty"`            // str	Y	公告日期
	FAnnDate          string  `json:"f_ann_date,omitempty"`          // str	Y	实际公告日期
	EndDate           string  `json:"end_date,omitempty"`            // str	Y	报告期
	ReportType        string  `json:"report_type,omitempty"`         // str	Y	报告类型 1合并报表 2单季合并 3调整单季合并表 4调整合并报表 5调整前合并报表 6母公司报表 7母公司单季表 8 母公司调整单季表 9母公司调整表 10母公司调整前报表 11调整前合并报表 12母公司调整前报表
	CompType          string  `json:"comp_type,omitempty"`           // str	Y	公司类型(1一般工商业2银行3保险4证券)
	BasicEps          float64 `json:"basic_eps,omitempty"`           // float	Y	基本每股收益
	DilutedEps        float64 `json:"diluted_eps,omitempty"`         // float	Y	稀释每股收益
	TotalRevenue      float64 `json:"total_revenue,omitempty"`       // float	Y	营业总收入
	Revenue           float64 `json:"revenue,omitempty"`             // float	Y	营业收入
	IntIncome         float64 `json:"int_income,omitempty"`          // float	Y	利息收入
	PremEarned        float64 `json:"prem_earned,omitempty"`         // float	Y	已赚保费
	CommIncome        float64 `json:"comm_income,omitempty"`         // float	Y	手续费及佣金收入
	NCommisIncome     float64 `json:"n_commis_income,omitempty"`     // float	Y	手续费及佣金净收入
	NOthIncome        float64 `json:"n_oth_income,omitempty"`        // float	Y	其他经营净收益
	NOthBIncome       float64 `json:"n_oth_b_income,omitempty"`      // float	Y	加:其他业务净收益
	PremIncome        float64 `json:"prem_income,omitempty"`         // float	Y	保险业务收入
	OutPrem           float64 `json:"out_prem,omitempty"`            // float	Y	减:分出保费
	UnePremReser      float64 `json:"une_prem_reser,omitempty"`      // float	Y	提取未到期责任准备金
	ReinsIncome       float64 `json:"reins_income,omitempty"`        // float	Y	其中:分保费收入
	NSecTbIncome      float64 `json:"n_sec_tb_income,omitempty"`     // float	Y	代理买卖证券业务净收入
	NSecUwIncome      float64 `json:"n_sec_uw_income,omitempty"`     // float	Y	证券承销业务净收入
	NAssetMgIncome    float64 `json:"n_asset_mg_income,omitempty"`   // float	Y	受托客户资产管理业务净收入
	OthBIncome        float64 `json:"oth_b_income,omitempty"`        // float	Y	其他业务收入
	FvValueChgGain    float64 `json:"fv_value_chg_gain,omitempty"`   // float	Y	加:公允价值变动净收益
	InvestIncome      float64 `json:"invest_income,omitempty"`       // float	Y	加:投资净收益
	AssInvestIncome   float64 `json:"ass_invest_income,omitempty"`   // float	Y	其中:对联营企业和合营企业的投资收益
	ForexGain         float64 `json:"forex_gain,omitempty"`          // float	Y	加:汇兑净收益
	TotalCogs         float64 `json:"total_cogs,omitempty"`          // float	Y	营业总成本
	OperCost          float64 `json:"oper_cost,omitempty"`           // float	Y	减:营业成本
	IntExp            float64 `json:"int_exp,omitempty"`             // float	Y	减:利息支出
	CommExp           float64 `json:"comm_exp,omitempty"`            // float	Y	减:手续费及佣金支出
	BizTaxSurchg      float64 `json:"biz_tax_surchg,omitempty"`      // float	Y	减:营业税金及附加
	SellExp           float64 `json:"sell_exp,omitempty"`            // float	Y	减:销售费用
	AdminExp          float64 `json:"admin_exp,omitempty"`           // float	Y	减:管理费用
	FinExp            float64 `json:"fin_exp,omitempty"`             // float	Y	减:财务费用
	AssetsImpairLoss  float64 `json:"assets_impair_loss,omitempty"`  // float	Y	减:资产减值损失
	PremRefund        float64 `json:"prem_refund,omitempty"`         // float	Y	退保金
	CompensPayout     float64 `json:"compens_payout,omitempty"`      // float	Y	赔付总支出
	ReserInsurLiab    float64 `json:"reser_insur_liab,omitempty"`    // float	Y	提取保险责任准备金
	DivPayt           float64 `json:"div_payt,omitempty"`            // float	Y	保户红利支出
	ReinsExp          float64 `json:"reins_exp,omitempty"`           // float	Y	分保费用
	OperExp           float64 `json:"oper_exp,omitempty"`            // float	Y	营业支出
	CompensPayoutRefu float64 `json:"compens_payout_refu,omitempty"` // float	Y	减:摊回赔付支出
	InsurReserRefu    float64 `json:"insur_reser_refu,omitempty"`    // float	Y	减:摊回保险责任准备金
	ReinsCostRefund   float64 `json:"reins_cost_refund,omitempty"`   // float	Y	减:摊回分保费用
	OtherBusCost      float64 `json:"other_bus_cost,omitempty"`      // float	Y	其他业务成本
	OperateProfit     float64 `json:"operate_profit,omitempty"`      // float	Y	营业利润
	NonOperIncome     float64 `json:"non_oper_income,omitempty"`     // float	Y	加:营业外收入
	NonOperExp        float64 `json:"non_oper_exp,omitempty"`        // float	Y	减:营业外支出
	NcaDisploss       float64 `json:"nca_disploss,omitempty"`        // float	Y	其中:减:非流动资产处置净损失
	TotalProfit       float64 `json:"total_profit,omitempty"`        // float	Y	利润总额
	IncomeTax         float64 `json:"income_tax,omitempty"`          // float	Y	所得税费用
	NIncome           float64 `json:"n_income,omitempty"`            // float	Y	净利润(含少数股东损益)
	NIncomeAttrP      float64 `json:"n_income_attr_p,omitempty"`     // float	Y	净利润(不含少数股东损益)
	MinorityGain      float64 `json:"minority_gain,omitempty"`       // float	Y	少数股东损益
	OthComprIncome    float64 `json:"oth_compr_income,omitempty"`    // float	Y	其他综合收益
	TComprIncome      float64 `json:"t_compr_income,omitempty"`      // float	Y	综合收益总额
	ComprIncAttrP     float64 `json:"compr_inc_attr_p,omitempty"`    // float	Y	归属于母公司(或股东)的综合收益总额
	ComprIncAttrMS    float64 `json:"compr_inc_attr_m_s,omitempty"`  // float	Y	归属于少数股东的综合收益总额
	Ebit              float64 `json:"ebit,omitempty"`                // float	Y	息税前利润
	Ebitda            float64 `json:"ebitda,omitempty"`              // float	Y	息税折旧摊销前利润
	InsuranceExp      float64 `json:"insurance_exp,omitempty"`       // float	Y	保险业务支出
	UndistProfit      float64 `json:"undist_profit,omitempty"`       // float	Y	年初未分配利润
	DistableProfit    float64 `json:"distable_profit,omitempty"`     // float	Y	可分配利润
	UpdateFlag        string  `json:"update_flag,omitempty"`         // str	N	更新标识,0未修改1更正过
}

func AssembleIncomeData(tsRsp *TushareResponse) []*IncomeData {
	tsData := []*IncomeData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(IncomeData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 获取上市公司财务利润表数据,用户需要至少800积分才可以调取,具体请参阅积分获取办法 https://tushare.pro/document/1?doc_id=13
func (ts *TuShare) Income(params SheetRequest, items IncomeItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "income",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type BalanceSheetItems struct {
	TsCode                bool `json:"ts_code,omitempty"`                    // str Y TS股票代码
	AnnDate               bool `json:"ann_date,omitempty"`                   // str Y 公告日期
	FAnnDate              bool `json:"f_ann_date,omitempty"`                 // str Y 实际公告日期
	EndDate               bool `json:"end_date,omitempty"`                   // str Y 报告期
	ReportType            bool `json:"report_type,omitempty"`                // str Y 报表类型
	CompType              bool `json:"comp_type,omitempty"`                  // str Y 公司类型
	TotalShare            bool `json:"total_share,omitempty"`                // float Y 期末总股本
	CapRese               bool `json:"cap_rese,omitempty"`                   // float Y 资本公积金
	UndistrPorfit         bool `json:"undistr_porfit,omitempty"`             // float Y 未分配利润
	SurplusRese           bool `json:"surplus_rese,omitempty"`               // float Y 盈余公积金
	SpecialRese           bool `json:"special_rese,omitempty"`               // float Y 专项储备
	MoneyCap              bool `json:"money_cap,omitempty"`                  // float Y 货币资金
	TradAsset             bool `json:"trad_asset,omitempty"`                 // float Y 交易性金融资产
	NotesReceiv           bool `json:"notes_receiv,omitempty"`               // float Y 应收票据
	AccountsReceiv        bool `json:"accounts_receiv,omitempty"`            // float Y 应收账款
	OthReceiv             bool `json:"oth_receiv,omitempty"`                 // float Y 其他应收款
	Prepayment            bool `json:"prepayment,omitempty"`                 // float Y 预付款项
	DivReceiv             bool `json:"div_receiv,omitempty"`                 // float Y 应收股利
	IntReceiv             bool `json:"int_receiv,omitempty"`                 // float Y 应收利息
	Inventories           bool `json:"inventories,omitempty"`                // float Y 存货
	AmorExp               bool `json:"amor_exp,omitempty"`                   // float Y 待摊费用
	NcaWithin1y           bool `json:"nca_within_1y,omitempty"`              // float Y 一年内到期的非流动资产
	SettRsrv              bool `json:"sett_rsrv,omitempty"`                  // float Y 结算备付金
	LoantoOthBankFi       bool `json:"loanto_oth_bank_fi,omitempty"`         // float Y 拆出资金
	PremiumReceiv         bool `json:"premium_receiv,omitempty"`             // float Y 应收保费
	ReinsurReceiv         bool `json:"reinsur_receiv,omitempty"`             // float Y 应收分保账款
	ReinsurResReceiv      bool `json:"reinsur_res_receiv,omitempty"`         // float Y 应收分保合同准备金
	PurResaleFa           bool `json:"pur_resale_fa,omitempty"`              // float Y 买入返售金融资产
	OthCurAssets          bool `json:"oth_cur_assets,omitempty"`             // float Y 其他流动资产
	TotalCurAssets        bool `json:"total_cur_assets,omitempty"`           // float Y 流动资产合计
	FaAvailForSale        bool `json:"fa_avail_for_sale,omitempty"`          // float Y 可供出售金融资产
	HtmInvest             bool `json:"htm_invest,omitempty"`                 // float Y 持有至到期投资
	LtEqtInvest           bool `json:"lt_eqt_invest,omitempty"`              // float Y 长期股权投资
	InvestRealEstate      bool `json:"invest_real_estate,omitempty"`         // float Y 投资性房地产
	TimeDeposits          bool `json:"time_deposits,omitempty"`              // float Y 定期存款
	OthAssets             bool `json:"oth_assets,omitempty"`                 // float Y 其他资产
	LtRec                 bool `json:"lt_rec,omitempty"`                     // float Y 长期应收款
	FixAssets             bool `json:"fix_assets,omitempty"`                 // float Y 固定资产
	Cip                   bool `json:"cip,omitempty"`                        // float Y 在建工程
	ConstMaterials        bool `json:"const_materials,omitempty"`            // float Y 工程物资
	FixedAssetsDisp       bool `json:"fixed_assets_disp,omitempty"`          // float Y 固定资产清理
	ProducBioAssets       bool `json:"produc_bio_assets,omitempty"`          // float Y 生产性生物资产
	OilAndGasAssets       bool `json:"oil_and_gas_assets,omitempty"`         // float Y 油气资产
	IntanAssets           bool `json:"intan_assets,omitempty"`               // float Y 无形资产
	RAndD                 bool `json:"r_and_d,omitempty"`                    // float Y 研发支出
	Goodwill              bool `json:"goodwill,omitempty"`                   // float Y 商誉
	LtAmorExp             bool `json:"lt_amor_exp,omitempty"`                // float Y 长期待摊费用
	DeferTaxAssets        bool `json:"defer_tax_assets,omitempty"`           // float Y 递延所得税资产
	DecrInDisbur          bool `json:"decr_in_disbur,omitempty"`             // float Y 发放贷款及垫款
	OthNca                bool `json:"oth_nca,omitempty"`                    // float Y 其他非流动资产
	TotalNca              bool `json:"total_nca,omitempty"`                  // float Y 非流动资产合计
	CashReserCb           bool `json:"cash_reser_cb,omitempty"`              // float Y 现金及存放中央银行款项
	DeposInOthBfi         bool `json:"depos_in_oth_bfi,omitempty"`           // float Y 存放同业和其它金融机构款项
	PrecMetals            bool `json:"prec_metals,omitempty"`                // float Y 贵金属
	DerivAssets           bool `json:"deriv_assets,omitempty"`               // float Y 衍生金融资产
	RrReinsUnePrem        bool `json:"rr_reins_une_prem,omitempty"`          // float Y 应收分保未到期责任准备金
	RrReinsOutstdCla      bool `json:"rr_reins_outstd_cla,omitempty"`        // float Y 应收分保未决赔款准备金
	RrReinsLinsLiab       bool `json:"rr_reins_lins_liab,omitempty"`         // float Y 应收分保寿险责任准备金
	RrReinsLthinsLiab     bool `json:"rr_reins_lthins_liab,omitempty"`       // float Y 应收分保长期健康险责任准备金
	RefundDepos           bool `json:"refund_depos,omitempty"`               // float Y 存出保证金
	PhPledgeLoans         bool `json:"ph_pledge_loans,omitempty"`            // float Y 保户质押贷款
	RefundCapDepos        bool `json:"refund_cap_depos,omitempty"`           // float Y 存出资本保证金
	IndepAcctAssets       bool `json:"indep_acct_assets,omitempty"`          // float Y 独立账户资产
	ClientDepos           bool `json:"client_depos,omitempty"`               // float Y 其中：客户资金存款
	ClientProv            bool `json:"client_prov,omitempty"`                // float Y 其中：客户备付金
	TransacSeatFee        bool `json:"transac_seat_fee,omitempty"`           // float Y 其中:交易席位费
	InvestAsReceiv        bool `json:"invest_as_receiv,omitempty"`           // float Y 应收款项类投资
	TotalAssets           bool `json:"total_assets,omitempty"`               // float Y 资产总计
	LtBorr                bool `json:"lt_borr,omitempty"`                    // float Y 长期借款
	StBorr                bool `json:"st_borr,omitempty"`                    // float Y 短期借款
	CbBorr                bool `json:"cb_borr,omitempty"`                    // float Y 向中央银行借款
	DeposIbDeposits       bool `json:"depos_ib_deposits,omitempty"`          // float Y 吸收存款及同业存放
	LoanOthBank           bool `json:"loan_oth_bank,omitempty"`              // float Y 拆入资金
	TradingFl             bool `json:"trading_fl,omitempty"`                 // float Y 交易性金融负债
	NotesPayable          bool `json:"notes_payable,omitempty"`              // float Y 应付票据
	AcctPayable           bool `json:"acct_payable,omitempty"`               // float Y 应付账款
	AdvReceipts           bool `json:"adv_receipts,omitempty"`               // float Y 预收款项
	SoldForRepurFa        bool `json:"sold_for_repur_fa,omitempty"`          // float Y 卖出回购金融资产款
	CommPayable           bool `json:"comm_payable,omitempty"`               // float Y 应付手续费及佣金
	PayrollPayable        bool `json:"payroll_payable,omitempty"`            // float Y 应付职工薪酬
	TaxesPayable          bool `json:"taxes_payable,omitempty"`              // float Y 应交税费
	IntPayable            bool `json:"int_payable,omitempty"`                // float Y 应付利息
	DivPayable            bool `json:"div_payable,omitempty"`                // float Y 应付股利
	OthPayable            bool `json:"oth_payable,omitempty"`                // float Y 其他应付款
	AccExp                bool `json:"acc_exp,omitempty"`                    // float Y 预提费用
	DeferredInc           bool `json:"deferred_inc,omitempty"`               // float Y 递延收益
	StBondsPayable        bool `json:"st_bonds_payable,omitempty"`           // float Y 应付短期债券
	PayableToReinsurer    bool `json:"payable_to_reinsurer,omitempty"`       // float Y 应付分保账款
	RsrvInsurCont         bool `json:"rsrv_insur_cont,omitempty"`            // float Y 保险合同准备金
	ActingTradingSec      bool `json:"acting_trading_sec,omitempty"`         // float Y 代理买卖证券款
	ActingUwSec           bool `json:"acting_uw_sec,omitempty"`              // float Y 代理承销证券款
	NonCurLiabDue1y       bool `json:"non_cur_liab_due_1y,omitempty"`        // float Y 一年内到期的非流动负债
	OthCurLiab            bool `json:"oth_cur_liab,omitempty"`               // float Y 其他流动负债
	TotalCurLiab          bool `json:"total_cur_liab,omitempty"`             // float Y 流动负债合计
	BondPayable           bool `json:"bond_payable,omitempty"`               // float Y 应付债券
	LtPayable             bool `json:"lt_payable,omitempty"`                 // float Y 长期应付款
	SpecificPayables      bool `json:"specific_payables,omitempty"`          // float Y 专项应付款
	EstimatedLiab         bool `json:"estimated_liab,omitempty"`             // float Y 预计负债
	DeferTaxLiab          bool `json:"defer_tax_liab,omitempty"`             // float Y 递延所得税负债
	DeferIncNonCurLiab    bool `json:"defer_inc_non_cur_liab,omitempty"`     // float Y 递延收益-非流动负债
	OthNcl                bool `json:"oth_ncl,omitempty"`                    // float Y 其他非流动负债
	TotalNcl              bool `json:"total_ncl,omitempty"`                  // float Y 非流动负债合计
	DeposOthBfi           bool `json:"depos_oth_bfi,omitempty"`              // float Y 同业和其它金融机构存放款项
	DerivLiab             bool `json:"deriv_liab,omitempty"`                 // float Y 衍生金融负债
	Depos                 bool `json:"depos,omitempty"`                      // float Y 吸收存款
	AgencyBusLiab         bool `json:"agency_bus_liab,omitempty"`            // float Y 代理业务负债
	OthLiab               bool `json:"oth_liab,omitempty"`                   // float Y 其他负债
	PremReceivAdva        bool `json:"prem_receiv_adva,omitempty"`           // float Y 预收保费
	DeposReceived         bool `json:"depos_received,omitempty"`             // float Y 存入保证金
	PhInvest              bool `json:"ph_invest,omitempty"`                  // float Y 保户储金及投资款
	ReserUnePrem          bool `json:"reser_une_prem,omitempty"`             // float Y 未到期责任准备金
	ReserOutstdClaims     bool `json:"reser_outstd_claims,omitempty"`        // float Y 未决赔款准备金
	ReserLinsLiab         bool `json:"reser_lins_liab,omitempty"`            // float Y 寿险责任准备金
	ReserLthinsLiab       bool `json:"reser_lthins_liab,omitempty"`          // float Y 长期健康险责任准备金
	IndeptAccLiab         bool `json:"indept_acc_liab,omitempty"`            // float Y 独立账户负债
	PledgeBorr            bool `json:"pledge_borr,omitempty"`                // float Y 其中:质押借款
	IndemPayable          bool `json:"indem_payable,omitempty"`              // float Y 应付赔付款
	PolicyDivPayable      bool `json:"policy_div_payable,omitempty"`         // float Y 应付保单红利
	TotalLiab             bool `json:"total_liab,omitempty"`                 // float Y 负债合计
	TreasuryShare         bool `json:"treasury_share,omitempty"`             // float Y 减:库存股
	OrdinRiskReser        bool `json:"ordin_risk_reser,omitempty"`           // float Y 一般风险准备
	ForexDiffer           bool `json:"forex_differ,omitempty"`               // float Y 外币报表折算差额
	InvestLossUnconf      bool `json:"invest_loss_unconf,omitempty"`         // float Y 未确认的投资损失
	MinorityInt           bool `json:"minority_int,omitempty"`               // float Y 少数股东权益
	TotalHldrEqyExcMinInt bool `json:"total_hldr_eqy_exc_min_int,omitempty"` // float Y 股东权益合计(不含少数股东权益)
	TotalHldrEqyIncMinInt bool `json:"total_hldr_eqy_inc_min_int,omitempty"` // float Y 股东权益合计(含少数股东权益)
	TotalLiabHldrEqy      bool `json:"total_liab_hldr_eqy,omitempty"`        // float Y 负债及股东权益总计
	LtPayrollPayable      bool `json:"lt_payroll_payable,omitempty"`         // float Y 长期应付职工薪酬
	OthCompIncome         bool `json:"oth_comp_income,omitempty"`            // float Y 其他综合收益
	OthEqtTools           bool `json:"oth_eqt_tools,omitempty"`              // float Y 其他权益工具
	OthEqtToolsPShr       bool `json:"oth_eqt_tools_p_shr,omitempty"`        // float Y 其他权益工具(优先股)
	LendingFunds          bool `json:"lending_funds,omitempty"`              // float Y 融出资金
	AccReceivable         bool `json:"acc_receivable,omitempty"`             // float Y 应收款项
	StFinPayable          bool `json:"st_fin_payable,omitempty"`             // float Y 应付短期融资款
	Payables              bool `json:"payables,omitempty"`                   // float Y 应付款项
	HfsAssets             bool `json:"hfs_assets,omitempty"`                 // float Y 持有待售的资产
	HfsSales              bool `json:"hfs_sales,omitempty"`                  // float Y 持有待售的负债
	UpdateFlag            bool `json:"update_flag,omitempty"`                // str N 更新标识
}

func (item BalanceSheetItems) All() BalanceSheetItems {
	item.TsCode = true
	item.AnnDate = true
	item.FAnnDate = true
	item.EndDate = true
	item.ReportType = true
	item.CompType = true
	item.TotalShare = true
	item.CapRese = true
	item.UndistrPorfit = true
	item.SurplusRese = true
	item.SpecialRese = true
	item.MoneyCap = true
	item.TradAsset = true
	item.NotesReceiv = true
	item.AccountsReceiv = true
	item.OthReceiv = true
	item.Prepayment = true
	item.DivReceiv = true
	item.IntReceiv = true
	item.Inventories = true
	item.AmorExp = true
	item.NcaWithin1y = true
	item.SettRsrv = true
	item.LoantoOthBankFi = true
	item.PremiumReceiv = true
	item.ReinsurReceiv = true
	item.ReinsurResReceiv = true
	item.PurResaleFa = true
	item.OthCurAssets = true
	item.TotalCurAssets = true
	item.FaAvailForSale = true
	item.HtmInvest = true
	item.LtEqtInvest = true
	item.InvestRealEstate = true
	item.TimeDeposits = true
	item.OthAssets = true
	item.LtRec = true
	item.FixAssets = true
	item.Cip = true
	item.ConstMaterials = true
	item.FixedAssetsDisp = true
	item.ProducBioAssets = true
	item.OilAndGasAssets = true
	item.IntanAssets = true
	item.RAndD = true
	item.Goodwill = true
	item.LtAmorExp = true
	item.DeferTaxAssets = true
	item.DecrInDisbur = true
	item.OthNca = true
	item.TotalNca = true
	item.CashReserCb = true
	item.DeposInOthBfi = true
	item.PrecMetals = true
	item.DerivAssets = true
	item.RrReinsUnePrem = true
	item.RrReinsOutstdCla = true
	item.RrReinsLinsLiab = true
	item.RrReinsLthinsLiab = true
	item.RefundDepos = true
	item.PhPledgeLoans = true
	item.RefundCapDepos = true
	item.IndepAcctAssets = true
	item.ClientDepos = true
	item.ClientProv = true
	item.TransacSeatFee = true
	item.InvestAsReceiv = true
	item.TotalAssets = true
	item.LtBorr = true
	item.StBorr = true
	item.CbBorr = true
	item.DeposIbDeposits = true
	item.LoanOthBank = true
	item.TradingFl = true
	item.NotesPayable = true
	item.AcctPayable = true
	item.AdvReceipts = true
	item.SoldForRepurFa = true
	item.CommPayable = true
	item.PayrollPayable = true
	item.TaxesPayable = true
	item.IntPayable = true
	item.DivPayable = true
	item.OthPayable = true
	item.AccExp = true
	item.DeferredInc = true
	item.StBondsPayable = true
	item.PayableToReinsurer = true
	item.RsrvInsurCont = true
	item.ActingTradingSec = true
	item.ActingUwSec = true
	item.NonCurLiabDue1y = true
	item.OthCurLiab = true
	item.TotalCurLiab = true
	item.BondPayable = true
	item.LtPayable = true
	item.SpecificPayables = true
	item.EstimatedLiab = true
	item.DeferTaxLiab = true
	item.DeferIncNonCurLiab = true
	item.OthNcl = true
	item.TotalNcl = true
	item.DeposOthBfi = true
	item.DerivLiab = true
	item.Depos = true
	item.AgencyBusLiab = true
	item.OthLiab = true
	item.PremReceivAdva = true
	item.DeposReceived = true
	item.PhInvest = true
	item.ReserUnePrem = true
	item.ReserOutstdClaims = true
	item.ReserLinsLiab = true
	item.ReserLthinsLiab = true
	item.IndeptAccLiab = true
	item.PledgeBorr = true
	item.IndemPayable = true
	item.PolicyDivPayable = true
	item.TotalLiab = true
	item.TreasuryShare = true
	item.OrdinRiskReser = true
	item.ForexDiffer = true
	item.InvestLossUnconf = true
	item.MinorityInt = true
	item.TotalHldrEqyExcMinInt = true
	item.TotalHldrEqyIncMinInt = true
	item.TotalLiabHldrEqy = true
	item.LtPayrollPayable = true
	item.OthCompIncome = true
	item.OthEqtTools = true
	item.OthEqtToolsPShr = true
	item.LendingFunds = true
	item.AccReceivable = true
	item.StFinPayable = true
	item.Payables = true
	item.HfsAssets = true
	item.HfsSales = true
	item.UpdateFlag = true
	return item
}

type BalanceSheetData struct {
	TsCode                string  `json:"ts_code,omitempty"`                    // str Y TS股票代码
	AnnDate               string  `json:"ann_date,omitempty"`                   // str Y 公告日期
	FAnnDate              string  `json:"f_ann_date,omitempty"`                 // str Y 实际公告日期
	EndDate               string  `json:"end_date,omitempty"`                   // str Y 报告期
	ReportType            string  `json:"report_type,omitempty"`                // str Y 报表类型
	CompType              string  `json:"comp_type,omitempty"`                  // str Y 公司类型
	TotalShare            float64 `json:"total_share,omitempty"`                // float Y 期末总股本
	CapRese               float64 `json:"cap_rese,omitempty"`                   // float Y 资本公积金
	UndistrPorfit         float64 `json:"undistr_porfit,omitempty"`             // float Y 未分配利润
	SurplusRese           float64 `json:"surplus_rese,omitempty"`               // float Y 盈余公积金
	SpecialRese           float64 `json:"special_rese,omitempty"`               // float Y 专项储备
	MoneyCap              float64 `json:"money_cap,omitempty"`                  // float Y 货币资金
	TradAsset             float64 `json:"trad_asset,omitempty"`                 // float Y 交易性金融资产
	NotesReceiv           float64 `json:"notes_receiv,omitempty"`               // float Y 应收票据
	AccountsReceiv        float64 `json:"accounts_receiv,omitempty"`            // float Y 应收账款
	OthReceiv             float64 `json:"oth_receiv,omitempty"`                 // float Y 其他应收款
	Prepayment            float64 `json:"prepayment,omitempty"`                 // float Y 预付款项
	DivReceiv             float64 `json:"div_receiv,omitempty"`                 // float Y 应收股利
	IntReceiv             float64 `json:"int_receiv,omitempty"`                 // float Y 应收利息
	Inventories           float64 `json:"inventories,omitempty"`                // float Y 存货
	AmorExp               float64 `json:"amor_exp,omitempty"`                   // float Y 待摊费用
	NcaWithin1y           float64 `json:"nca_within_1y,omitempty"`              // float Y 一年内到期的非流动资产
	SettRsrv              float64 `json:"sett_rsrv,omitempty"`                  // float Y 结算备付金
	LoantoOthBankFi       float64 `json:"loanto_oth_bank_fi,omitempty"`         // float Y 拆出资金
	PremiumReceiv         float64 `json:"premium_receiv,omitempty"`             // float Y 应收保费
	ReinsurReceiv         float64 `json:"reinsur_receiv,omitempty"`             // float Y 应收分保账款
	ReinsurResReceiv      float64 `json:"reinsur_res_receiv,omitempty"`         // float Y 应收分保合同准备金
	PurResaleFa           float64 `json:"pur_resale_fa,omitempty"`              // float Y 买入返售金融资产
	OthCurAssets          float64 `json:"oth_cur_assets,omitempty"`             // float Y 其他流动资产
	TotalCurAssets        float64 `json:"total_cur_assets,omitempty"`           // float Y 流动资产合计
	FaAvailForSale        float64 `json:"fa_avail_for_sale,omitempty"`          // float Y 可供出售金融资产
	HtmInvest             float64 `json:"htm_invest,omitempty"`                 // float Y 持有至到期投资
	LtEqtInvest           float64 `json:"lt_eqt_invest,omitempty"`              // float Y 长期股权投资
	InvestRealEstate      float64 `json:"invest_real_estate,omitempty"`         // float Y 投资性房地产
	TimeDeposits          float64 `json:"time_deposits,omitempty"`              // float Y 定期存款
	OthAssets             float64 `json:"oth_assets,omitempty"`                 // float Y 其他资产
	LtRec                 float64 `json:"lt_rec,omitempty"`                     // float Y 长期应收款
	FixAssets             float64 `json:"fix_assets,omitempty"`                 // float Y 固定资产
	Cip                   float64 `json:"cip,omitempty"`                        // float Y 在建工程
	ConstMaterials        float64 `json:"const_materials,omitempty"`            // float Y 工程物资
	FixedAssetsDisp       float64 `json:"fixed_assets_disp,omitempty"`          // float Y 固定资产清理
	ProducBioAssets       float64 `json:"produc_bio_assets,omitempty"`          // float Y 生产性生物资产
	OilAndGasAssets       float64 `json:"oil_and_gas_assets,omitempty"`         // float Y 油气资产
	IntanAssets           float64 `json:"intan_assets,omitempty"`               // float Y 无形资产
	RAndD                 float64 `json:"r_and_d,omitempty"`                    // float Y 研发支出
	Goodwill              float64 `json:"goodwill,omitempty"`                   // float Y 商誉
	LtAmorExp             float64 `json:"lt_amor_exp,omitempty"`                // float Y 长期待摊费用
	DeferTaxAssets        float64 `json:"defer_tax_assets,omitempty"`           // float Y 递延所得税资产
	DecrInDisbur          float64 `json:"decr_in_disbur,omitempty"`             // float Y 发放贷款及垫款
	OthNca                float64 `json:"oth_nca,omitempty"`                    // float Y 其他非流动资产
	TotalNca              float64 `json:"total_nca,omitempty"`                  // float Y 非流动资产合计
	CashReserCb           float64 `json:"cash_reser_cb,omitempty"`              // float Y 现金及存放中央银行款项
	DeposInOthBfi         float64 `json:"depos_in_oth_bfi,omitempty"`           // float Y 存放同业和其它金融机构款项
	PrecMetals            float64 `json:"prec_metals,omitempty"`                // float Y 贵金属
	DerivAssets           float64 `json:"deriv_assets,omitempty"`               // float Y 衍生金融资产
	RrReinsUnePrem        float64 `json:"rr_reins_une_prem,omitempty"`          // float Y 应收分保未到期责任准备金
	RrReinsOutstdCla      float64 `json:"rr_reins_outstd_cla,omitempty"`        // float Y 应收分保未决赔款准备金
	RrReinsLinsLiab       float64 `json:"rr_reins_lins_liab,omitempty"`         // float Y 应收分保寿险责任准备金
	RrReinsLthinsLiab     float64 `json:"rr_reins_lthins_liab,omitempty"`       // float Y 应收分保长期健康险责任准备金
	RefundDepos           float64 `json:"refund_depos,omitempty"`               // float Y 存出保证金
	PhPledgeLoans         float64 `json:"ph_pledge_loans,omitempty"`            // float Y 保户质押贷款
	RefundCapDepos        float64 `json:"refund_cap_depos,omitempty"`           // float Y 存出资本保证金
	IndepAcctAssets       float64 `json:"indep_acct_assets,omitempty"`          // float Y 独立账户资产
	ClientDepos           float64 `json:"client_depos,omitempty"`               // float Y 其中：客户资金存款
	ClientProv            float64 `json:"client_prov,omitempty"`                // float Y 其中：客户备付金
	TransacSeatFee        float64 `json:"transac_seat_fee,omitempty"`           // float Y 其中:交易席位费
	InvestAsReceiv        float64 `json:"invest_as_receiv,omitempty"`           // float Y 应收款项类投资
	TotalAssets           float64 `json:"total_assets,omitempty"`               // float Y 资产总计
	LtBorr                float64 `json:"lt_borr,omitempty"`                    // float Y 长期借款
	StBorr                float64 `json:"st_borr,omitempty"`                    // float Y 短期借款
	CbBorr                float64 `json:"cb_borr,omitempty"`                    // float Y 向中央银行借款
	DeposIbDeposits       float64 `json:"depos_ib_deposits,omitempty"`          // float Y 吸收存款及同业存放
	LoanOthBank           float64 `json:"loan_oth_bank,omitempty"`              // float Y 拆入资金
	TradingFl             float64 `json:"trading_fl,omitempty"`                 // float Y 交易性金融负债
	NotesPayable          float64 `json:"notes_payable,omitempty"`              // float Y 应付票据
	AcctPayable           float64 `json:"acct_payable,omitempty"`               // float Y 应付账款
	AdvReceipts           float64 `json:"adv_receipts,omitempty"`               // float Y 预收款项
	SoldForRepurFa        float64 `json:"sold_for_repur_fa,omitempty"`          // float Y 卖出回购金融资产款
	CommPayable           float64 `json:"comm_payable,omitempty"`               // float Y 应付手续费及佣金
	PayrollPayable        float64 `json:"payroll_payable,omitempty"`            // float Y 应付职工薪酬
	TaxesPayable          float64 `json:"taxes_payable,omitempty"`              // float Y 应交税费
	IntPayable            float64 `json:"int_payable,omitempty"`                // float Y 应付利息
	DivPayable            float64 `json:"div_payable,omitempty"`                // float Y 应付股利
	OthPayable            float64 `json:"oth_payable,omitempty"`                // float Y 其他应付款
	AccExp                float64 `json:"acc_exp,omitempty"`                    // float Y 预提费用
	DeferredInc           float64 `json:"deferred_inc,omitempty"`               // float Y 递延收益
	StBondsPayable        float64 `json:"st_bonds_payable,omitempty"`           // float Y 应付短期债券
	PayableToReinsurer    float64 `json:"payable_to_reinsurer,omitempty"`       // float Y 应付分保账款
	RsrvInsurCont         float64 `json:"rsrv_insur_cont,omitempty"`            // float Y 保险合同准备金
	ActingTradingSec      float64 `json:"acting_trading_sec,omitempty"`         // float Y 代理买卖证券款
	ActingUwSec           float64 `json:"acting_uw_sec,omitempty"`              // float Y 代理承销证券款
	NonCurLiabDue1y       float64 `json:"non_cur_liab_due_1y,omitempty"`        // float Y 一年内到期的非流动负债
	OthCurLiab            float64 `json:"oth_cur_liab,omitempty"`               // float Y 其他流动负债
	TotalCurLiab          float64 `json:"total_cur_liab,omitempty"`             // float Y 流动负债合计
	BondPayable           float64 `json:"bond_payable,omitempty"`               // float Y 应付债券
	LtPayable             float64 `json:"lt_payable,omitempty"`                 // float Y 长期应付款
	SpecificPayables      float64 `json:"specific_payables,omitempty"`          // float Y 专项应付款
	EstimatedLiab         float64 `json:"estimated_liab,omitempty"`             // float Y 预计负债
	DeferTaxLiab          float64 `json:"defer_tax_liab,omitempty"`             // float Y 递延所得税负债
	DeferIncNonCurLiab    float64 `json:"defer_inc_non_cur_liab,omitempty"`     // float Y 递延收益-非流动负债
	OthNcl                float64 `json:"oth_ncl,omitempty"`                    // float Y 其他非流动负债
	TotalNcl              float64 `json:"total_ncl,omitempty"`                  // float Y 非流动负债合计
	DeposOthBfi           float64 `json:"depos_oth_bfi,omitempty"`              // float Y 同业和其它金融机构存放款项
	DerivLiab             float64 `json:"deriv_liab,omitempty"`                 // float Y 衍生金融负债
	Depos                 float64 `json:"depos,omitempty"`                      // float Y 吸收存款
	AgencyBusLiab         float64 `json:"agency_bus_liab,omitempty"`            // float Y 代理业务负债
	OthLiab               float64 `json:"oth_liab,omitempty"`                   // float Y 其他负债
	PremReceivAdva        float64 `json:"prem_receiv_adva,omitempty"`           // float Y 预收保费
	DeposReceived         float64 `json:"depos_received,omitempty"`             // float Y 存入保证金
	PhInvest              float64 `json:"ph_invest,omitempty"`                  // float Y 保户储金及投资款
	ReserUnePrem          float64 `json:"reser_une_prem,omitempty"`             // float Y 未到期责任准备金
	ReserOutstdClaims     float64 `json:"reser_outstd_claims,omitempty"`        // float Y 未决赔款准备金
	ReserLinsLiab         float64 `json:"reser_lins_liab,omitempty"`            // float Y 寿险责任准备金
	ReserLthinsLiab       float64 `json:"reser_lthins_liab,omitempty"`          // float Y 长期健康险责任准备金
	IndeptAccLiab         float64 `json:"indept_acc_liab,omitempty"`            // float Y 独立账户负债
	PledgeBorr            float64 `json:"pledge_borr,omitempty"`                // float Y 其中:质押借款
	IndemPayable          float64 `json:"indem_payable,omitempty"`              // float Y 应付赔付款
	PolicyDivPayable      float64 `json:"policy_div_payable,omitempty"`         // float Y 应付保单红利
	TotalLiab             float64 `json:"total_liab,omitempty"`                 // float Y 负债合计
	TreasuryShare         float64 `json:"treasury_share,omitempty"`             // float Y 减:库存股
	OrdinRiskReser        float64 `json:"ordin_risk_reser,omitempty"`           // float Y 一般风险准备
	ForexDiffer           float64 `json:"forex_differ,omitempty"`               // float Y 外币报表折算差额
	InvestLossUnconf      float64 `json:"invest_loss_unconf,omitempty"`         // float Y 未确认的投资损失
	MinorityInt           float64 `json:"minority_int,omitempty"`               // float Y 少数股东权益
	TotalHldrEqyExcMinInt float64 `json:"total_hldr_eqy_exc_min_int,omitempty"` // float Y 股东权益合计(不含少数股东权益)
	TotalHldrEqyIncMinInt float64 `json:"total_hldr_eqy_inc_min_int,omitempty"` // float Y 股东权益合计(含少数股东权益)
	TotalLiabHldrEqy      float64 `json:"total_liab_hldr_eqy,omitempty"`        // float Y 负债及股东权益总计
	LtPayrollPayable      float64 `json:"lt_payroll_payable,omitempty"`         // float Y 长期应付职工薪酬
	OthCompIncome         float64 `json:"oth_comp_income,omitempty"`            // float Y 其他综合收益
	OthEqtTools           float64 `json:"oth_eqt_tools,omitempty"`              // float Y 其他权益工具
	OthEqtToolsPShr       float64 `json:"oth_eqt_tools_p_shr,omitempty"`        // float Y 其他权益工具(优先股)
	LendingFunds          float64 `json:"lending_funds,omitempty"`              // float Y 融出资金
	AccReceivable         float64 `json:"acc_receivable,omitempty"`             // float Y 应收款项
	StFinPayable          float64 `json:"st_fin_payable,omitempty"`             // float Y 应付短期融资款
	Payables              float64 `json:"payables,omitempty"`                   // float Y 应付款项
	HfsAssets             float64 `json:"hfs_assets,omitempty"`                 // float Y 持有待售的资产
	HfsSales              float64 `json:"hfs_sales,omitempty"`                  // float Y 持有待售的负债
	UpdateFlag            string  `json:"update_flag,omitempty"`                // str N 更新标识
}

func AssembleBalanceSheetData(tsRsp *TushareResponse) []*BalanceSheetData {
	tsData := []*BalanceSheetData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(BalanceSheetData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 获取上市公司资产负债表,用户需要至少800积分才可以调取,具体请参阅积分获取办法 https://tushare.pro/document/1?doc_id=13
func (ts *TuShare) BalanceSheet(params SheetRequest, items BalanceSheetItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "balancesheet",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type CashFlowItems struct {
	TsCode                  bool `json:"ts_code,omitempty"`                     // str	Y	TS股票代码
	AnnDate                 bool `json:"ann_date,omitempty"`                    // str	Y	公告日期
	FAnnDate                bool `json:"f_ann_date,omitempty"`                  // str	Y	实际公告日期
	EndDate                 bool `json:"end_date,omitempty"`                    // str	Y	报告期
	CompType                bool `json:"comp_type,omitempty"`                   // str	Y	公司类型
	ReportType              bool `json:"report_type,omitempty"`                 // str	Y	报表类型
	NetProfit               bool `json:"net_profit,omitempty"`                  // float	Y	净利润
	FinanExp                bool `json:"finan_exp,omitempty"`                   // float	Y	财务费用
	CFrSaleSg               bool `json:"c_fr_sale_sg,omitempty"`                // float	Y	销售商品、提供劳务收到的现金
	RecpTaxRends            bool `json:"recp_tax_rends,omitempty"`              // float	Y	收到的税费返还
	NDeposIncrFi            bool `json:"n_depos_incr_fi,omitempty"`             // float	Y	客户存款和同业存放款项净增加额
	NIncrLoansCb            bool `json:"n_incr_loans_cb,omitempty"`             // float	Y	向中央银行借款净增加额
	NIncBorrOthFi           bool `json:"n_inc_borr_oth_fi,omitempty"`           // float	Y	向其他金融机构拆入资金净增加额
	PremFrOrigContr         bool `json:"prem_fr_orig_contr,omitempty"`          // float	Y	收到原保险合同保费取得的现金
	NIncrInsuredDep         bool `json:"n_incr_insured_dep,omitempty"`          // float	Y	保户储金净增加额
	NReinsurPrem            bool `json:"n_reinsur_prem,omitempty"`              // float	Y	收到再保业务现金净额
	NIncrDispTfa            bool `json:"n_incr_disp_tfa,omitempty"`             // float	Y	处置交易性金融资产净增加额
	IfcCashIncr             bool `json:"ifc_cash_incr,omitempty"`               // float	Y	收取利息和手续费净增加额
	NIncrDispFaas           bool `json:"n_incr_disp_faas,omitempty"`            // float	Y	处置可供出售金融资产净增加额
	NIncrLoansOthBank       bool `json:"n_incr_loans_oth_bank,omitempty"`       // float	Y	拆入资金净增加额
	NCapIncrRepur           bool `json:"n_cap_incr_repur,omitempty"`            // float	Y	回购业务资金净增加额
	CFrOthOperateA          bool `json:"c_fr_oth_operate_a,omitempty"`          // float	Y	收到其他与经营活动有关的现金
	CInfFrOperateA          bool `json:"c_inf_fr_operate_a,omitempty"`          // float	Y	经营活动现金流入小计
	CPaidGoodsS             bool `json:"c_paid_goods_s,omitempty"`              // float	Y	购买商品、接受劳务支付的现金
	CPaidToForEmpl          bool `json:"c_paid_to_for_empl,omitempty"`          // float	Y	支付给职工以及为职工支付的现金
	CPaidForTaxes           bool `json:"c_paid_for_taxes,omitempty"`            // float	Y	支付的各项税费
	NIncrCltLoanAdv         bool `json:"n_incr_clt_loan_adv,omitempty"`         // float	Y	客户贷款及垫款净增加额
	NIncrDepCbob            bool `json:"n_incr_dep_cbob,omitempty"`             // float	Y	存放央行和同业款项净增加额
	CPayClaimsOrigInco      bool `json:"c_pay_claims_orig_inco,omitempty"`      // float	Y	支付原保险合同赔付款项的现金
	PayHandlingChrg         bool `json:"pay_handling_chrg,omitempty"`           // float	Y	支付手续费的现金
	PayCommInsurPlcy        bool `json:"pay_comm_insur_plcy,omitempty"`         // float	Y	支付保单红利的现金
	OthCashPayOperAct       bool `json:"oth_cash_pay_oper_act,omitempty"`       // float	Y	支付其他与经营活动有关的现金
	StCashOutAct            bool `json:"st_cash_out_act,omitempty"`             // float	Y	经营活动现金流出小计
	NCashflowAct            bool `json:"n_cashflow_act,omitempty"`              // float	Y	经营活动产生的现金流量净额
	OthRecpRalInvAct        bool `json:"oth_recp_ral_inv_act,omitempty"`        // float	Y	收到其他与投资活动有关的现金
	CDispWithdrwlInvest     bool `json:"c_disp_withdrwl_invest,omitempty"`      // float	Y	收回投资收到的现金
	CRecpReturnInvest       bool `json:"c_recp_return_invest,omitempty"`        // float	Y	取得投资收益收到的现金
	NRecpDispFiolta         bool `json:"n_recp_disp_fiolta,omitempty"`          // float	Y	处置固定资产、无形资产和其他长期资产收回的现金净额
	NRecpDispSobu           bool `json:"n_recp_disp_sobu,omitempty"`            // float	Y	处置子公司及其他营业单位收到的现金净额
	StotInflowsInvAct       bool `json:"stot_inflows_inv_act,omitempty"`        // float	Y	投资活动现金流入小计
	CPayAcqConstFiolta      bool `json:"c_pay_acq_const_fiolta,omitempty"`      // float	Y	购建固定资产、无形资产和其他长期资产支付的现金
	CPaidInvest             bool `json:"c_paid_invest,omitempty"`               // float	Y	投资支付的现金
	NDispSubsOthBiz         bool `json:"n_disp_subs_oth_biz,omitempty"`         // float	Y	取得子公司及其他营业单位支付的现金净额
	OthPayRalInvAct         bool `json:"oth_pay_ral_inv_act,omitempty"`         // float	Y	支付其他与投资活动有关的现金
	NIncrPledgeLoan         bool `json:"n_incr_pledge_loan,omitempty"`          // float	Y	质押贷款净增加额
	StotOutInvAct           bool `json:"stot_out_inv_act,omitempty"`            // float	Y	投资活动现金流出小计
	NCashflowInvAct         bool `json:"n_cashflow_inv_act,omitempty"`          // float	Y	投资活动产生的现金流量净额
	CRecpBorrow             bool `json:"c_recp_borrow,omitempty"`               // float	Y	取得借款收到的现金
	ProcIssueBonds          bool `json:"proc_issue_bonds,omitempty"`            // float	Y	发行债券收到的现金
	OthCashRecpRalFncAct    bool `json:"oth_cash_recp_ral_fnc_act,omitempty"`   // float	Y	收到其他与筹资活动有关的现金
	StotCashInFncAct        bool `json:"stot_cash_in_fnc_act,omitempty"`        // float	Y	筹资活动现金流入小计
	FreeCashflow            bool `json:"free_cashflow,omitempty"`               // float	Y	企业自由现金流量
	CPrepayAmtBorr          bool `json:"c_prepay_amt_borr,omitempty"`           // float	Y	偿还债务支付的现金
	CPayDistDpcpIntExp      bool `json:"c_pay_dist_dpcp_int_exp,omitempty"`     // float	Y	分配股利、利润或偿付利息支付的现金
	InclDvdProfitPaidScMs   bool `json:"incl_dvd_profit_paid_sc_ms,omitempty"`  // float	Y	其中:子公司支付给少数股东的股利、利润
	OthCashpayRalFncAct     bool `json:"oth_cashpay_ral_fnc_act,omitempty"`     // float	Y	支付其他与筹资活动有关的现金
	StotCashoutFncAct       bool `json:"stot_cashout_fnc_act,omitempty"`        // float	Y	筹资活动现金流出小计
	NCashFlowsFncAct        bool `json:"n_cash_flows_fnc_act,omitempty"`        // float	Y	筹资活动产生的现金流量净额
	EffFxFluCash            bool `json:"eff_fx_flu_cash,omitempty"`             // float	Y	汇率变动对现金的影响
	NIncrCashCashEqu        bool `json:"n_incr_cash_cash_equ,omitempty"`        // float	Y	现金及现金等价物净增加额
	CCashEquBegPeriod       bool `json:"c_cash_equ_beg_period,omitempty"`       // float	Y	期初现金及现金等价物余额
	CCashEquEndPeriod       bool `json:"c_cash_equ_end_period,omitempty"`       // float	Y	期末现金及现金等价物余额
	CRecpCapContrib         bool `json:"c_recp_cap_contrib,omitempty"`          // float	Y	吸收投资收到的现金
	InclCashRecSaims        bool `json:"incl_cash_rec_saims,omitempty"`         // float	Y	其中:子公司吸收少数股东投资收到的现金
	UnconInvestLoss         bool `json:"uncon_invest_loss,omitempty"`           // float	Y	未确认投资损失
	ProvDeprAssets          bool `json:"prov_depr_assets,omitempty"`            // float	Y	加:资产减值准备
	DeprFaCogaDpba          bool `json:"depr_fa_coga_dpba,omitempty"`           // float	Y	固定资产折旧、油气资产折耗、生产性生物资产折旧
	AmortIntangAssets       bool `json:"amort_intang_assets,omitempty"`         // float	Y	无形资产摊销
	LtAmortDeferredExp      bool `json:"lt_amort_deferred_exp,omitempty"`       // float	Y	长期待摊费用摊销
	DecrDeferredExp         bool `json:"decr_deferred_exp,omitempty"`           // float	Y	待摊费用减少
	IncrAccExp              bool `json:"incr_acc_exp,omitempty"`                // float	Y	预提费用增加
	LossDispFiolta          bool `json:"loss_disp_fiolta,omitempty"`            // float	Y	处置固定、无形资产和其他长期资产的损失
	LossScrFa               bool `json:"loss_scr_fa,omitempty"`                 // float	Y	固定资产报废损失
	LossFvChg               bool `json:"loss_fv_chg,omitempty"`                 // float	Y	公允价值变动损失
	InvestLoss              bool `json:"invest_loss,omitempty"`                 // float	Y	投资损失
	DecrDefIncTaxAssets     bool `json:"decr_def_inc_tax_assets,omitempty"`     // float	Y	递延所得税资产减少
	IncrDefIncTaxLiab       bool `json:"incr_def_inc_tax_liab,omitempty"`       // float	Y	递延所得税负债增加
	DecrInventories         bool `json:"decr_inventories,omitempty"`            // float	Y	存货的减少
	DecrOperPayable         bool `json:"decr_oper_payable,omitempty"`           // float	Y	经营性应收项目的减少
	IncrOperPayable         bool `json:"incr_oper_payable,omitempty"`           // float	Y	经营性应付项目的增加
	Others                  bool `json:"others,omitempty"`                      // float	Y	其他
	ImNetCashflowOperAct    bool `json:"im_net_cashflow_oper_act,omitempty"`    // float	Y	经营活动产生的现金流量净额(间接法)
	ConvDebtIntoCap         bool `json:"conv_debt_into_cap,omitempty"`          // float	Y	债务转为资本
	ConvCopbondsDueWithin1y bool `json:"conv_copbonds_due_within_1y,omitempty"` // float	Y	一年内到期的可转换公司债券
	FaFncLeases             bool `json:"fa_fnc_leases,omitempty"`               // float	Y	融资租入固定资产
	EndBalCash              bool `json:"end_bal_cash,omitempty"`                // float	Y	现金的期末余额
	BegBalCash              bool `json:"beg_bal_cash,omitempty"`                // float	Y	减:现金的期初余额
	EndBalCashEqu           bool `json:"end_bal_cash_equ,omitempty"`            // float	Y	加:现金等价物的期末余额
	BegBalCashEqu           bool `json:"beg_bal_cash_equ,omitempty"`            // float	Y	减:现金等价物的期初余额
	ImNIncrCashEqu          bool `json:"im_n_incr_cash_equ,omitempty"`          // float	Y	现金及现金等价物净增加额(间接法)
	UpdateFlag              bool `json:"update_flag,omitempty"`                 // str	N	更新标识
}

func (item CashFlowItems) All() CashFlowItems {
	item.TsCode = true
	item.AnnDate = true
	item.FAnnDate = true
	item.EndDate = true
	item.CompType = true
	item.ReportType = true
	item.NetProfit = true
	item.FinanExp = true
	item.CFrSaleSg = true
	item.RecpTaxRends = true
	item.NDeposIncrFi = true
	item.NIncrLoansCb = true
	item.NIncBorrOthFi = true
	item.PremFrOrigContr = true
	item.NIncrInsuredDep = true
	item.NReinsurPrem = true
	item.NIncrDispTfa = true
	item.IfcCashIncr = true
	item.NIncrDispFaas = true
	item.NIncrLoansOthBank = true
	item.NCapIncrRepur = true
	item.CFrOthOperateA = true
	item.CInfFrOperateA = true
	item.CPaidGoodsS = true
	item.CPaidToForEmpl = true
	item.CPaidForTaxes = true
	item.NIncrCltLoanAdv = true
	item.NIncrDepCbob = true
	item.CPayClaimsOrigInco = true
	item.PayHandlingChrg = true
	item.PayCommInsurPlcy = true
	item.OthCashPayOperAct = true
	item.StCashOutAct = true
	item.NCashflowAct = true
	item.OthRecpRalInvAct = true
	item.CDispWithdrwlInvest = true
	item.CRecpReturnInvest = true
	item.NRecpDispFiolta = true
	item.NRecpDispSobu = true
	item.StotInflowsInvAct = true
	item.CPayAcqConstFiolta = true
	item.CPaidInvest = true
	item.NDispSubsOthBiz = true
	item.OthPayRalInvAct = true
	item.NIncrPledgeLoan = true
	item.StotOutInvAct = true
	item.NCashflowInvAct = true
	item.CRecpBorrow = true
	item.ProcIssueBonds = true
	item.OthCashRecpRalFncAct = true
	item.StotCashInFncAct = true
	item.FreeCashflow = true
	item.CPrepayAmtBorr = true
	item.CPayDistDpcpIntExp = true
	item.InclDvdProfitPaidScMs = true
	item.OthCashpayRalFncAct = true
	item.StotCashoutFncAct = true
	item.NCashFlowsFncAct = true
	item.EffFxFluCash = true
	item.NIncrCashCashEqu = true
	item.CCashEquBegPeriod = true
	item.CCashEquEndPeriod = true
	item.CRecpCapContrib = true
	item.InclCashRecSaims = true
	item.UnconInvestLoss = true
	item.ProvDeprAssets = true
	item.DeprFaCogaDpba = true
	item.AmortIntangAssets = true
	item.LtAmortDeferredExp = true
	item.DecrDeferredExp = true
	item.IncrAccExp = true
	item.LossDispFiolta = true
	item.LossScrFa = true
	item.LossFvChg = true
	item.InvestLoss = true
	item.DecrDefIncTaxAssets = true
	item.IncrDefIncTaxLiab = true
	item.DecrInventories = true
	item.DecrOperPayable = true
	item.IncrOperPayable = true
	item.Others = true
	item.ImNetCashflowOperAct = true
	item.ConvDebtIntoCap = true
	item.ConvCopbondsDueWithin1y = true
	item.FaFncLeases = true
	item.EndBalCash = true
	item.BegBalCash = true
	item.EndBalCashEqu = true
	item.BegBalCashEqu = true
	item.ImNIncrCashEqu = true
	item.UpdateFlag = true
	return item
}

type CashFlowData struct {
	TsCode                  string  `json:"ts_code,omitempty"`                     // str	Y	TS股票代码
	AnnDate                 string  `json:"ann_date,omitempty"`                    // str	Y	公告日期
	FAnnDate                string  `json:"f_ann_date,omitempty"`                  // str	Y	实际公告日期
	EndDate                 string  `json:"end_date,omitempty"`                    // str	Y	报告期
	CompType                string  `json:"comp_type,omitempty"`                   // str	Y	公司类型
	ReportType              string  `json:"report_type,omitempty"`                 // str	Y	报表类型
	NetProfit               float64 `json:"net_profit,omitempty"`                  // float	Y	净利润
	FinanExp                float64 `json:"finan_exp,omitempty"`                   // float	Y	财务费用
	CFrSaleSg               float64 `json:"c_fr_sale_sg,omitempty"`                // float	Y	销售商品、提供劳务收到的现金
	RecpTaxRends            float64 `json:"recp_tax_rends,omitempty"`              // float	Y	收到的税费返还
	NDeposIncrFi            float64 `json:"n_depos_incr_fi,omitempty"`             // float	Y	客户存款和同业存放款项净增加额
	NIncrLoansCb            float64 `json:"n_incr_loans_cb,omitempty"`             // float	Y	向中央银行借款净增加额
	NIncBorrOthFi           float64 `json:"n_inc_borr_oth_fi,omitempty"`           // float	Y	向其他金融机构拆入资金净增加额
	PremFrOrigContr         float64 `json:"prem_fr_orig_contr,omitempty"`          // float	Y	收到原保险合同保费取得的现金
	NIncrInsuredDep         float64 `json:"n_incr_insured_dep,omitempty"`          // float	Y	保户储金净增加额
	NReinsurPrem            float64 `json:"n_reinsur_prem,omitempty"`              // float	Y	收到再保业务现金净额
	NIncrDispTfa            float64 `json:"n_incr_disp_tfa,omitempty"`             // float	Y	处置交易性金融资产净增加额
	IfcCashIncr             float64 `json:"ifc_cash_incr,omitempty"`               // float	Y	收取利息和手续费净增加额
	NIncrDispFaas           float64 `json:"n_incr_disp_faas,omitempty"`            // float	Y	处置可供出售金融资产净增加额
	NIncrLoansOthBank       float64 `json:"n_incr_loans_oth_bank,omitempty"`       // float	Y	拆入资金净增加额
	NCapIncrRepur           float64 `json:"n_cap_incr_repur,omitempty"`            // float	Y	回购业务资金净增加额
	CFrOthOperateA          float64 `json:"c_fr_oth_operate_a,omitempty"`          // float	Y	收到其他与经营活动有关的现金
	CInfFrOperateA          float64 `json:"c_inf_fr_operate_a,omitempty"`          // float	Y	经营活动现金流入小计
	CPaidGoodsS             float64 `json:"c_paid_goods_s,omitempty"`              // float	Y	购买商品、接受劳务支付的现金
	CPaidToForEmpl          float64 `json:"c_paid_to_for_empl,omitempty"`          // float	Y	支付给职工以及为职工支付的现金
	CPaidForTaxes           float64 `json:"c_paid_for_taxes,omitempty"`            // float	Y	支付的各项税费
	NIncrCltLoanAdv         float64 `json:"n_incr_clt_loan_adv,omitempty"`         // float	Y	客户贷款及垫款净增加额
	NIncrDepCbob            float64 `json:"n_incr_dep_cbob,omitempty"`             // float	Y	存放央行和同业款项净增加额
	CPayClaimsOrigInco      float64 `json:"c_pay_claims_orig_inco,omitempty"`      // float	Y	支付原保险合同赔付款项的现金
	PayHandlingChrg         float64 `json:"pay_handling_chrg,omitempty"`           // float	Y	支付手续费的现金
	PayCommInsurPlcy        float64 `json:"pay_comm_insur_plcy,omitempty"`         // float	Y	支付保单红利的现金
	OthCashPayOperAct       float64 `json:"oth_cash_pay_oper_act,omitempty"`       // float	Y	支付其他与经营活动有关的现金
	StCashOutAct            float64 `json:"st_cash_out_act,omitempty"`             // float	Y	经营活动现金流出小计
	NCashflowAct            float64 `json:"n_cashflow_act,omitempty"`              // float	Y	经营活动产生的现金流量净额
	OthRecpRalInvAct        float64 `json:"oth_recp_ral_inv_act,omitempty"`        // float	Y	收到其他与投资活动有关的现金
	CDispWithdrwlInvest     float64 `json:"c_disp_withdrwl_invest,omitempty"`      // float	Y	收回投资收到的现金
	CRecpReturnInvest       float64 `json:"c_recp_return_invest,omitempty"`        // float	Y	取得投资收益收到的现金
	NRecpDispFiolta         float64 `json:"n_recp_disp_fiolta,omitempty"`          // float	Y	处置固定资产、无形资产和其他长期资产收回的现金净额
	NRecpDispSobu           float64 `json:"n_recp_disp_sobu,omitempty"`            // float	Y	处置子公司及其他营业单位收到的现金净额
	StotInflowsInvAct       float64 `json:"stot_inflows_inv_act,omitempty"`        // float	Y	投资活动现金流入小计
	CPayAcqConstFiolta      float64 `json:"c_pay_acq_const_fiolta,omitempty"`      // float	Y	购建固定资产、无形资产和其他长期资产支付的现金
	CPaidInvest             float64 `json:"c_paid_invest,omitempty"`               // float	Y	投资支付的现金
	NDispSubsOthBiz         float64 `json:"n_disp_subs_oth_biz,omitempty"`         // float	Y	取得子公司及其他营业单位支付的现金净额
	OthPayRalInvAct         float64 `json:"oth_pay_ral_inv_act,omitempty"`         // float	Y	支付其他与投资活动有关的现金
	NIncrPledgeLoan         float64 `json:"n_incr_pledge_loan,omitempty"`          // float	Y	质押贷款净增加额
	StotOutInvAct           float64 `json:"stot_out_inv_act,omitempty"`            // float	Y	投资活动现金流出小计
	NCashflowInvAct         float64 `json:"n_cashflow_inv_act,omitempty"`          // float	Y	投资活动产生的现金流量净额
	CRecpBorrow             float64 `json:"c_recp_borrow,omitempty"`               // float	Y	取得借款收到的现金
	ProcIssueBonds          float64 `json:"proc_issue_bonds,omitempty"`            // float	Y	发行债券收到的现金
	OthCashRecpRalFncAct    float64 `json:"oth_cash_recp_ral_fnc_act,omitempty"`   // float	Y	收到其他与筹资活动有关的现金
	StotCashInFncAct        float64 `json:"stot_cash_in_fnc_act,omitempty"`        // float	Y	筹资活动现金流入小计
	FreeCashflow            float64 `json:"free_cashflow,omitempty"`               // float	Y	企业自由现金流量
	CPrepayAmtBorr          float64 `json:"c_prepay_amt_borr,omitempty"`           // float	Y	偿还债务支付的现金
	CPayDistDpcpIntExp      float64 `json:"c_pay_dist_dpcp_int_exp,omitempty"`     // float	Y	分配股利、利润或偿付利息支付的现金
	InclDvdProfitPaidScMs   float64 `json:"incl_dvd_profit_paid_sc_ms,omitempty"`  // float	Y	其中:子公司支付给少数股东的股利、利润
	OthCashpayRalFncAct     float64 `json:"oth_cashpay_ral_fnc_act,omitempty"`     // float	Y	支付其他与筹资活动有关的现金
	StotCashoutFncAct       float64 `json:"stot_cashout_fnc_act,omitempty"`        // float	Y	筹资活动现金流出小计
	NCashFlowsFncAct        float64 `json:"n_cash_flows_fnc_act,omitempty"`        // float	Y	筹资活动产生的现金流量净额
	EffFxFluCash            float64 `json:"eff_fx_flu_cash,omitempty"`             // float	Y	汇率变动对现金的影响
	NIncrCashCashEqu        float64 `json:"n_incr_cash_cash_equ,omitempty"`        // float	Y	现金及现金等价物净增加额
	CCashEquBegPeriod       float64 `json:"c_cash_equ_beg_period,omitempty"`       // float	Y	期初现金及现金等价物余额
	CCashEquEndPeriod       float64 `json:"c_cash_equ_end_period,omitempty"`       // float	Y	期末现金及现金等价物余额
	CRecpCapContrib         float64 `json:"c_recp_cap_contrib,omitempty"`          // float	Y	吸收投资收到的现金
	InclCashRecSaims        float64 `json:"incl_cash_rec_saims,omitempty"`         // float	Y	其中:子公司吸收少数股东投资收到的现金
	UnconInvestLoss         float64 `json:"uncon_invest_loss,omitempty"`           // float	Y	未确认投资损失
	ProvDeprAssets          float64 `json:"prov_depr_assets,omitempty"`            // float	Y	加:资产减值准备
	DeprFaCogaDpba          float64 `json:"depr_fa_coga_dpba,omitempty"`           // float	Y	固定资产折旧、油气资产折耗、生产性生物资产折旧
	AmortIntangAssets       float64 `json:"amort_intang_assets,omitempty"`         // float	Y	无形资产摊销
	LtAmortDeferredExp      float64 `json:"lt_amort_deferred_exp,omitempty"`       // float	Y	长期待摊费用摊销
	DecrDeferredExp         float64 `json:"decr_deferred_exp,omitempty"`           // float	Y	待摊费用减少
	IncrAccExp              float64 `json:"incr_acc_exp,omitempty"`                // float	Y	预提费用增加
	LossDispFiolta          float64 `json:"loss_disp_fiolta,omitempty"`            // float	Y	处置固定、无形资产和其他长期资产的损失
	LossScrFa               float64 `json:"loss_scr_fa,omitempty"`                 // float	Y	固定资产报废损失
	LossFvChg               float64 `json:"loss_fv_chg,omitempty"`                 // float	Y	公允价值变动损失
	InvestLoss              float64 `json:"invest_loss,omitempty"`                 // float	Y	投资损失
	DecrDefIncTaxAssets     float64 `json:"decr_def_inc_tax_assets,omitempty"`     // float	Y	递延所得税资产减少
	IncrDefIncTaxLiab       float64 `json:"incr_def_inc_tax_liab,omitempty"`       // float	Y	递延所得税负债增加
	DecrInventories         float64 `json:"decr_inventories,omitempty"`            // float	Y	存货的减少
	DecrOperPayable         float64 `json:"decr_oper_payable,omitempty"`           // float	Y	经营性应收项目的减少
	IncrOperPayable         float64 `json:"incr_oper_payable,omitempty"`           // float	Y	经营性应付项目的增加
	Others                  float64 `json:"others,omitempty"`                      // float	Y	其他
	ImNetCashflowOperAct    float64 `json:"im_net_cashflow_oper_act,omitempty"`    // float	Y	经营活动产生的现金流量净额(间接法)
	ConvDebtIntoCap         float64 `json:"conv_debt_into_cap,omitempty"`          // float	Y	债务转为资本
	ConvCopbondsDueWithin1y float64 `json:"conv_copbonds_due_within_1y,omitempty"` // float	Y	一年内到期的可转换公司债券
	FaFncLeases             float64 `json:"fa_fnc_leases,omitempty"`               // float	Y	融资租入固定资产
	EndBalCash              float64 `json:"end_bal_cash,omitempty"`                // float	Y	现金的期末余额
	BegBalCash              float64 `json:"beg_bal_cash,omitempty"`                // float	Y	减:现金的期初余额
	EndBalCashEqu           float64 `json:"end_bal_cash_equ,omitempty"`            // float	Y	加:现金等价物的期末余额
	BegBalCashEqu           float64 `json:"beg_bal_cash_equ,omitempty"`            // float	Y	减:现金等价物的期初余额
	ImNIncrCashEqu          float64 `json:"im_n_incr_cash_equ,omitempty"`          // float	Y	现金及现金等价物净增加额(间接法)
	UpdateFlag              string  `json:"update_flag,omitempty"`                 // str	N	更新标识
}

func AssembleCashFlowData(tsRsp *TushareResponse) []*CashFlowData {
	tsData := []*CashFlowData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(CashFlowData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 获取新股上市列表数据,单次最大2000条,总量不限制,用户需要至少120积分才可以调取,具体请参阅积分获取办法 https://tushare.pro/document/1?doc_id=13
func (ts *TuShare) CashFlow(params SheetRequest, items CashFlowItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "cashflow",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type ForecastRequest struct {
	TsCode    string `json:"ts_code,omitempty"`    // str	Y	股票代码
	AnnDate   string `json:"ann_date,omitempty"`   // str	N	公告日期
	StartDate string `json:"start_date,omitempty"` // str	N	公告开始日期
	EndDate   string `json:"end_date,omitempty"`   // str	N	公告结束日期
	Period    string `json:"period,omitempty"`     // str	N	报告期(每个季度最后一天的日期,比如20171231表示年报)
	Type      string `json:"type,omitempty"`       // str	N	预告类型(预增/预减/扭亏/首亏/续亏/续盈/略增/略减)
}

type ForecastItems struct {
	TsCode        bool `json:"ts_code,omitempty"`         // str	TS股票代码
	AnnDate       bool `json:"ann_date,omitempty"`        // str	公告日期
	EndDate       bool `json:"end_date,omitempty"`        // str	报告期
	Type          bool `json:"type,omitempty"`            // str	业绩预告类型(预增/预减/扭亏/首亏/续亏/续盈/略增/略减)
	PChangeMin    bool `json:"p_change_min,omitempty"`    // float	预告净利润变动幅度下限(%)
	PChangeMax    bool `json:"p_change_max,omitempty"`    // float	预告净利润变动幅度上限(%)
	NetProfitMin  bool `json:"net_profit_min,omitempty"`  // float	预告净利润下限(万元)
	NetProfitMax  bool `json:"net_profit_max,omitempty"`  // float	预告净利润上限(万元)
	LastParentNet bool `json:"last_parent_net,omitempty"` // float	上年同期归属母公司净利润
	FirstAnnDate  bool `json:"first_ann_date,omitempty"`  // str	首次公告日
	Summary       bool `json:"summary,omitempty"`         // str	业绩预告摘要
	ChangeReason  bool `json:"change_reason,omitempty"`   // str	业绩变动原因
}

func (item ForecastItems) All() ForecastItems {
	item.TsCode = true
	item.AnnDate = true
	item.EndDate = true
	item.Type = true
	item.PChangeMin = true
	item.PChangeMax = true
	item.NetProfitMin = true
	item.NetProfitMax = true
	item.LastParentNet = true
	item.FirstAnnDate = true
	item.Summary = true
	item.ChangeReason = true
	return item
}

type ForecastData struct {
	TsCode        string  `json:"ts_code,omitempty"`         // str	TS股票代码
	AnnDate       string  `json:"ann_date,omitempty"`        // str	公告日期
	EndDate       string  `json:"end_date,omitempty"`        // str	报告期
	Type          string  `json:"type,omitempty"`            // str	业绩预告类型(预增/预减/扭亏/首亏/续亏/续盈/略增/略减)
	PChangeMin    float64 `json:"p_change_min,omitempty"`    // float	预告净利润变动幅度下限(%)
	PChangeMax    float64 `json:"p_change_max,omitempty"`    // float	预告净利润变动幅度上限(%)
	NetProfitMin  float64 `json:"net_profit_min,omitempty"`  // float	预告净利润下限(万元)
	NetProfitMax  float64 `json:"net_profit_max,omitempty"`  // float	预告净利润上限(万元)
	LastParentNet float64 `json:"last_parent_net,omitempty"` // float	上年同期归属母公司净利润
	FirstAnnDate  string  `json:"first_ann_date,omitempty"`  // str	首次公告日
	Summary       string  `json:"summary,omitempty"`         // str	业绩预告摘要
	ChangeReason  string  `json:"change_reason,omitempty"`   // str	业绩变动原因
}

func AssembleForecastData(tsRsp *TushareResponse) []*ForecastData {
	tsData := []*ForecastData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(ForecastData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 获取业绩预告数据,用户需要至少800积分才可以调取,具体请参阅积分获取办法 https://tushare.pro/document/1?doc_id=13
func (ts *TuShare) Forecast(params ForecastRequest, items ForecastItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "forecast",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type FinanceRequest struct {
	TsCode    string `json:"ts_code,omitempty"`    // str	Y	股票代码
	AnnDate   string `json:"ann_date,omitempty" `  // str	N	公告日期
	StartDate string `json:"start_date,omitempty"` // str	N	公告开始日期
	EndDate   string `json:"end_date,omitempty"`   // str	N	公告结束日期
	Period    string `json:"period,omitempty"`     // str	N	报告期(每个季度最后一天的日期,比如20171231表示年报)
}

type ExpressItems struct {
	TsCode                bool `json:"ts_code,omitempty"`                    //	str	TS股票代码
	AnnDate               bool `json:"ann_date,omitempty"`                   //	str	公告日期
	EndDate               bool `json:"end_date,omitempty"`                   //	str	报告期
	Revenue               bool `json:"revenue,omitempty"`                    //	float	营业收入(元)
	OperateProfit         bool `json:"operate_profit,omitempty"`             //	float	营业利润(元)
	TotalProfit           bool `json:"total_profit,omitempty"`               //	float	利润总额(元)
	NIncome               bool `json:"n_income,omitempty"`                   //	float	净利润(元)
	TotalAssets           bool `json:"total_assets,omitempty"`               //	float	总资产(元)
	TotalHldrEqyExcMinInt bool `json:"total_hldr_eqy_exc_min_int,omitempty"` //	float	股东权益合计(不含少数股东权益)(元)
	DilutedEps            bool `json:"diluted_eps,omitempty"`                //	float	每股收益(摊薄)(元)
	DilutedRoe            bool `json:"diluted_roe,omitempty"`                //	float	净资产收益率(摊薄)(%)
	YoyNetProfit          bool `json:"yoy_net_profit,omitempty"`             //	float	去年同期修正后净利润
	Bps                   bool `json:"bps,omitempty"`                        //	float	每股净资产
	YoySales              bool `json:"yoy_sales,omitempty"`                  //	float	同比增长率:营业收入
	YoyOp                 bool `json:"yoy_op,omitempty"`                     //	float	同比增长率:营业利润
	YoyTp                 bool `json:"yoy_tp,omitempty"`                     //	float	同比增长率:利润总额
	YoyDeduNp             bool `json:"yoy_dedu_np,omitempty"`                //	float	同比增长率:归属母公司股东的净利润
	YoyEps                bool `json:"yoy_eps,omitempty"`                    //	float	同比增长率:基本每股收益
	YoyRoe                bool `json:"yoy_roe,omitempty"`                    //	float	同比增减:加权平均净资产收益率
	GrowthAssets          bool `json:"growth_assets,omitempty"`              //	float	比年初增长率:总资产
	YoyEquity             bool `json:"yoy_equity,omitempty"`                 //	float	比年初增长率:归属母公司的股东权益
	GrowthBps             bool `json:"growth_bps,omitempty"`                 //	float	比年初增长率:归属于母公司股东的每股净资产
	OrLastYear            bool `json:"or_last_year,omitempty"`               //	float	去年同期营业收入
	OpLastYear            bool `json:"op_last_year,omitempty"`               //	float	去年同期营业利润
	TpLastYear            bool `json:"tp_last_year,omitempty"`               //	float	去年同期利润总额
	NpLastYear            bool `json:"np_last_year,omitempty"`               //	float	去年同期净利润
	EpsLastYear           bool `json:"eps_last_year,omitempty"`              //	float	去年同期每股收益
	OpenNetAssets         bool `json:"open_net_assets,omitempty"`            //	float	期初净资产
	OpenBps               bool `json:"open_bps,omitempty"`                   //	float	期初每股净资产
	PerfSummary           bool `json:"perf_summary,omitempty"`               //	str	业绩简要说明
	IsAudit               bool `json:"is_audit,omitempty"`                   //	int	是否审计： 1是 0否
	Remark                bool `json:"remark,omitempty"`                     //	str	备注
}

func (item ExpressItems) All() ExpressItems {
	item.TsCode = true
	item.AnnDate = true
	item.EndDate = true
	item.Revenue = true
	item.OperateProfit = true
	item.TotalProfit = true
	item.NIncome = true
	item.TotalAssets = true
	item.TotalHldrEqyExcMinInt = true
	item.DilutedEps = true
	item.DilutedRoe = true
	item.YoyNetProfit = true
	item.Bps = true
	item.YoySales = true
	item.YoyOp = true
	item.YoyTp = true
	item.YoyDeduNp = true
	item.YoyEps = true
	item.YoyRoe = true
	item.GrowthAssets = true
	item.YoyEquity = true
	item.GrowthBps = true
	item.OrLastYear = true
	item.OpLastYear = true
	item.TpLastYear = true
	item.NpLastYear = true
	item.EpsLastYear = true
	item.OpenNetAssets = true
	item.OpenBps = true
	item.PerfSummary = true
	item.IsAudit = true
	item.Remark = true
	return item
}

type ExpressData struct {
	TsCode                string  `json:"ts_code,omitempty"`                    //	str	TS股票代码
	AnnDate               string  `json:"ann_date,omitempty"`                   //	str	公告日期
	EndDate               string  `json:"end_date,omitempty"`                   //	str	报告期
	Revenue               float64 `json:"revenue,omitempty"`                    //	float	营业收入(元)
	OperateProfit         float64 `json:"operate_profit,omitempty"`             //	float	营业利润(元)
	TotalProfit           float64 `json:"total_profit,omitempty"`               //	float	利润总额(元)
	NIncome               float64 `json:"n_income,omitempty"`                   //	float	净利润(元)
	TotalAssets           float64 `json:"total_assets,omitempty"`               //	float	总资产(元)
	TotalHldrEqyExcMinInt float64 `json:"total_hldr_eqy_exc_min_int,omitempty"` //	float	股东权益合计(不含少数股东权益)(元)
	DilutedEps            float64 `json:"diluted_eps,omitempty"`                //	float	每股收益(摊薄)(元)
	DilutedRoe            float64 `json:"diluted_roe,omitempty"`                //	float	净资产收益率(摊薄)(%)
	YoyNetProfit          float64 `json:"yoy_net_profit,omitempty"`             //	float	去年同期修正后净利润
	Bps                   float64 `json:"bps,omitempty"`                        //	float	每股净资产
	YoySales              float64 `json:"yoy_sales,omitempty"`                  //	float	同比增长率:营业收入
	YoyOp                 float64 `json:"yoy_op,omitempty"`                     //	float	同比增长率:营业利润
	YoyTp                 float64 `json:"yoy_tp,omitempty"`                     //	float	同比增长率:利润总额
	YoyDeduNp             float64 `json:"yoy_dedu_np,omitempty"`                //	float	同比增长率:归属母公司股东的净利润
	YoyEps                float64 `json:"yoy_eps,omitempty"`                    //	float	同比增长率:基本每股收益
	YoyRoe                float64 `json:"yoy_roe,omitempty"`                    //	float	同比增减:加权平均净资产收益率
	GrowthAssets          float64 `json:"growth_assets,omitempty"`              //	float	比年初增长率:总资产
	YoyEquity             float64 `json:"yoy_equity,omitempty"`                 //	float	比年初增长率:归属母公司的股东权益
	GrowthBps             float64 `json:"growth_bps,omitempty"`                 //	float	比年初增长率:归属于母公司股东的每股净资产
	OrLastYear            float64 `json:"or_last_year,omitempty"`               //	float	去年同期营业收入
	OpLastYear            float64 `json:"op_last_year,omitempty"`               //	float	去年同期营业利润
	TpLastYear            float64 `json:"tp_last_year,omitempty"`               //	float	去年同期利润总额
	NpLastYear            float64 `json:"np_last_year,omitempty"`               //	float	去年同期净利润
	EpsLastYear           float64 `json:"eps_last_year,omitempty"`              //	float	去年同期每股收益
	OpenNetAssets         float64 `json:"open_net_assets,omitempty"`            //	float	期初净资产
	OpenBps               float64 `json:"open_bps,omitempty"`                   //	float	期初每股净资产
	PerfSummary           string  `json:"perf_summary,omitempty"`               //	str	业绩简要说明
	IsAudit               int64   `json:"is_audit,omitempty"`                   //	int	是否审计： 1是 0否
	Remark                string  `json:"remark,omitempty"`                     //	str	备注
}

func AssembleExpressData(tsRsp *TushareResponse) []*ExpressData {
	tsData := []*ExpressData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(ExpressData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 获取上市公司业绩快报,用户需要至少800积分才可以调取,具体请参阅积分获取办法 https://tushare.pro/document/1?doc_id=13
func (ts *TuShare) Express(params FinanceRequest, items ExpressItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "express",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type DividendRequest struct {
	TsCode     string `json:"ts_code,omitempty"`      // str	Y	股票代码
	AnnDate    string `json:"ann_date,omitempty"`     // str	N	公告日期
	RecordDate string `json:"record_date,omitempty"`  // str	N	股权登记日期
	ExDate     string `json:"ex_date,omitempty"`      // str	N	除权除息日
	ImpAnnDate string `json:"imp_ann_date,omitempty"` // str	N	实施公告日
}

type DividendItems struct {
	TsCode      bool `json:"ts_code,omitempty"`      // str	Y	TS代码
	EndDate     bool `json:"end_date,omitempty"`     // str	Y	分红年度
	AnnDate     bool `json:"ann_date,omitempty"`     // str	Y	预案公告日
	DivProc     bool `json:"div_proc,omitempty"`     // str	Y	实施进度
	StkDiv      bool `json:"stk_div,omitempty"`      // float	Y	每股送转
	StkBoRate   bool `json:"stk_bo_rate,omitempty"`  // float	Y	每股送股比例
	StkCoRate   bool `json:"stk_co_rate,omitempty"`  // float	Y	每股转增比例
	CashDiv     bool `json:"cash_div,omitempty"`     // float	Y	每股分红(税后)
	CashDivTax  bool `json:"cash_div_tax,omitempty"` // float	Y	每股分红(税前)
	RecordDate  bool `json:"record_date,omitempty"`  // str	Y	股权登记日
	ExDate      bool `json:"ex_date,omitempty"`      // str	Y	除权除息日
	PayDate     bool `json:"pay_date,omitempty"`     // str	Y	派息日
	DivListdate bool `json:"div_listdate,omitempty"` // str	Y	红股上市日
	ImpAnnDate  bool `json:"imp_ann_date,omitempty"` // str	Y	实施公告日
	BaseDate    bool `json:"base_date,omitempty"`    // str	N	基准日
	BaseShare   bool `json:"base_share,omitempty"`   // float	N	基准股本(万)
}

func (item DividendItems) All() DividendItems {
	item.TsCode = true
	item.EndDate = true
	item.AnnDate = true
	item.DivProc = true
	item.StkDiv = true
	item.StkBoRate = true
	item.StkCoRate = true
	item.CashDiv = true
	item.CashDivTax = true
	item.RecordDate = true
	item.ExDate = true
	item.PayDate = true
	item.DivListdate = true
	item.ImpAnnDate = true
	item.BaseDate = true
	item.BaseShare = true
	return item
}

type DividendData struct {
	TsCode      string  `json:"ts_code,omitempty"`      // str	Y	TS代码
	EndDate     string  `json:"end_date,omitempty"`     // str	Y	分红年度
	AnnDate     string  `json:"ann_date,omitempty"`     // str	Y	预案公告日
	DivProc     string  `json:"div_proc,omitempty"`     // str	Y	实施进度
	StkDiv      float64 `json:"stk_div,omitempty"`      // float	Y	每股送转
	StkBoRate   float64 `json:"stk_bo_rate,omitempty"`  // float	Y	每股送股比例
	StkCoRate   float64 `json:"stk_co_rate,omitempty"`  // float	Y	每股转增比例
	CashDiv     float64 `json:"cash_div,omitempty"`     // float	Y	每股分红(税后)
	CashDivTax  float64 `json:"cash_div_tax,omitempty"` // float	Y	每股分红(税前)
	RecordDate  string  `json:"record_date,omitempty"`  // str	Y	股权登记日
	ExDate      string  `json:"ex_date,omitempty"`      // str	Y	除权除息日
	PayDate     string  `json:"pay_date,omitempty"`     // str	Y	派息日
	DivListdate string  `json:"div_listdate,omitempty"` // str	Y	红股上市日
	ImpAnnDate  string  `json:"imp_ann_date,omitempty"` // str	Y	实施公告日
	BaseDate    string  `json:"base_date,omitempty"`    // str	N	基准日
	BaseShare   float64 `json:"base_share,omitempty"`   // float	N	基准股本(万)
}

func AssembleDividendData(tsRsp *TushareResponse) []*DividendData {
	tsData := []*DividendData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(DividendData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 分红送股数据,用户需要至少900积分才可以调取,具体请参阅积分获取办法 https://tushare.pro/document/1?doc_id=13
func (ts *TuShare) Dividend(params DividendRequest, items DividendItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "dividend",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type FinaIndicatorItems struct {
	TsCode                    bool `json:"ts_code,omitempty"`                      //	str	Y	TS代码
	AnnDate                   bool `json:"ann_date,omitempty"`                     //	str	Y	公告日期
	EndDate                   bool `json:"end_date,omitempty"`                     //	str	Y	报告期
	Eps                       bool `json:"eps,omitempty"`                          //	float	Y	基本每股收益
	DtEps                     bool `json:"dt_eps,omitempty"`                       //	float	Y	稀释每股收益
	TotalRevenuePs            bool `json:"total_revenue_ps,omitempty"`             //	float	Y	每股营业总收入
	RevenuePs                 bool `json:"revenue_ps,omitempty"`                   //	float	Y	每股营业收入
	CapitalResePs             bool `json:"capital_rese_ps,omitempty"`              //	float	Y	每股资本公积
	SurplusResePs             bool `json:"surplus_rese_ps,omitempty"`              //	float	Y	每股盈余公积
	UndistProfitPs            bool `json:"undist_profit_ps,omitempty"`             //	float	Y	每股未分配利润
	ExtraItem                 bool `json:"extra_item,omitempty"`                   //	float	Y	非经常性损益
	ProfitDedt                bool `json:"profit_dedt,omitempty"`                  //	float	Y	扣除非经常性损益后的净利润
	GrossMargin               bool `json:"gross_margin,omitempty"`                 //	float	Y	毛利
	CurrentRatio              bool `json:"current_ratio,omitempty"`                //	float	Y	流动比率
	QuickRatio                bool `json:"quick_ratio,omitempty"`                  //	float	Y	速动比率
	CashRatio                 bool `json:"cash_ratio,omitempty"`                   //	float	Y	保守速动比率
	InvturnDays               bool `json:"invturn_days,omitempty"`                 //	float	N	存货周转天数
	ArturnDays                bool `json:"arturn_days,omitempty"`                  //	float	N	应收账款周转天数
	InvTurn                   bool `json:"inv_turn,omitempty"`                     //	float	N	存货周转率
	ArTurn                    bool `json:"ar_turn,omitempty"`                      //	float	Y	应收账款周转率
	CaTurn                    bool `json:"ca_turn,omitempty"`                      //	float	Y	流动资产周转率
	FaTurn                    bool `json:"fa_turn,omitempty"`                      //	float	Y	固定资产周转率
	AssetsTurn                bool `json:"assets_turn,omitempty"`                  //	float	Y	总资产周转率
	OpIncome                  bool `json:"op_income,omitempty"`                    //	float	Y	经营活动净收益
	ValuechangeIncome         bool `json:"valuechange_income,omitempty"`           //	float	N	价值变动净收益
	InterstIncome             bool `json:"interst_income,omitempty"`               //	float	N	利息费用
	Daa                       bool `json:"daa,omitempty"`                          //	float	N	折旧与摊销
	Ebit                      bool `json:"ebit,omitempty"`                         //	float	Y	息税前利润
	Ebitda                    bool `json:"ebitda,omitempty"`                       //	float	Y	息税折旧摊销前利润
	Fcff                      bool `json:"fcff,omitempty"`                         //	float	Y	企业自由现金流量
	Fcfe                      bool `json:"fcfe,omitempty"`                         //	float	Y	股权自由现金流量
	CurrentExint              bool `json:"current_exint,omitempty"`                //	float	Y	无息流动负债
	NoncurrentExint           bool `json:"noncurrent_exint,omitempty"`             //	float	Y	无息非流动负债
	Interestdebt              bool `json:"interestdebt,omitempty"`                 //	float	Y	带息债务
	Netdebt                   bool `json:"netdebt,omitempty"`                      //	float	Y	净债务
	TangibleAsset             bool `json:"tangible_asset,omitempty"`               //	float	Y	有形资产
	WorkingCapital            bool `json:"working_capital,omitempty"`              //	float	Y	营运资金
	NetworkingCapital         bool `json:"networking_capital,omitempty"`           //	float	Y	营运流动资本
	InvestCapital             bool `json:"invest_capital,omitempty"`               //	float	Y	全部投入资本
	RetainedEarnings          bool `json:"retained_earnings,omitempty"`            //	float	Y	留存收益
	Diluted2Eps               bool `json:"diluted2_eps,omitempty"`                 //	float	Y	期末摊薄每股收益
	Bps                       bool `json:"bps,omitempty"`                          //	float	Y	每股净资产
	Ocfps                     bool `json:"ocfps,omitempty"`                        //	float	Y	每股经营活动产生的现金流量净额
	Retainedps                bool `json:"retainedps,omitempty"`                   //	float	Y	每股留存收益
	Cfps                      bool `json:"cfps,omitempty"`                         //	float	Y	每股现金流量净额
	EbitPs                    bool `json:"ebit_ps,omitempty"`                      //	float	Y	每股息税前利润
	FcffPs                    bool `json:"fcff_ps,omitempty"`                      //	float	Y	每股企业自由现金流量
	FcfePs                    bool `json:"fcfe_ps,omitempty"`                      //	float	Y	每股股东自由现金流量
	NetprofitMargin           bool `json:"netprofit_margin,omitempty"`             //	float	Y	销售净利率
	GrossprofitMargin         bool `json:"grossprofit_margin,omitempty"`           //	float	Y	销售毛利率
	CogsOfSales               bool `json:"cogs_of_sales,omitempty"`                //	float	Y	销售成本率
	ExpenseOfSales            bool `json:"expense_of_sales,omitempty"`             //	float	Y	销售期间费用率
	ProfitToGr                bool `json:"profit_to_gr,omitempty"`                 //	float	Y	净利润/营业总收入
	SaleexpToGr               bool `json:"saleexp_to_gr,omitempty"`                //	float	Y	销售费用/营业总收入
	AdminexpOfGr              bool `json:"adminexp_of_gr,omitempty"`               //	float	Y	管理费用/营业总收入
	FinaexpOfGr               bool `json:"finaexp_of_gr,omitempty"`                //	float	Y	财务费用/营业总收入
	ImpaiTtm                  bool `json:"impai_ttm,omitempty"`                    //	float	Y	资产减值损失/营业总收入
	GcOfGr                    bool `json:"gc_of_gr,omitempty"`                     //	float	Y	营业总成本/营业总收入
	OpOfGr                    bool `json:"op_of_gr,omitempty"`                     //	float	Y	营业利润/营业总收入
	EbitOfGr                  bool `json:"ebit_of_gr,omitempty"`                   //	float	Y	息税前利润/营业总收入
	Roe                       bool `json:"roe,omitempty"`                          //	float	Y	净资产收益率
	RoeWaa                    bool `json:"roe_waa,omitempty"`                      //	float	Y	加权平均净资产收益率
	RoeDt                     bool `json:"roe_dt,omitempty"`                       //	float	Y	净资产收益率(扣除非经常损益)
	Roa                       bool `json:"roa,omitempty"`                          //	float	Y	总资产报酬率
	Npta                      bool `json:"npta,omitempty"`                         //	float	Y	总资产净利润
	Roic                      bool `json:"roic,omitempty"`                         //	float	Y	投入资本回报率
	RoeYearly                 bool `json:"roe_yearly,omitempty"`                   //	float	Y	年化净资产收益率
	Roa2Yearly                bool `json:"roa2_yearly,omitempty"`                  //	float	Y	年化总资产报酬率
	RoeAvg                    bool `json:"roe_avg,omitempty"`                      //	float	N	平均净资产收益率(增发条件)
	OpincomeOfEbt             bool `json:"opincome_of_ebt,omitempty"`              //	float	N	经营活动净收益/利润总额
	InvestincomeOfEbt         bool `json:"investincome_of_ebt,omitempty"`          //	float	N	价值变动净收益/利润总额
	NOpProfitOfEbt            bool `json:"n_op_profit_of_ebt,omitempty"`           //	float	N	营业外收支净额/利润总额
	TaxToEbt                  bool `json:"tax_to_ebt,omitempty"`                   //	float	N	所得税/利润总额
	DtprofitToProfit          bool `json:"dtprofit_to_profit,omitempty"`           //	float	N	扣除非经常损益后的净利润/净利润
	SalescashToOr             bool `json:"salescash_to_or,omitempty"`              //	float	N	销售商品提供劳务收到的现金/营业收入
	OcfToOr                   bool `json:"ocf_to_or,omitempty"`                    //	float	N	经营活动产生的现金流量净额/营业收入
	OcfToOpincome             bool `json:"ocf_to_opincome,omitempty"`              //	float	N	经营活动产生的现金流量净额/经营活动净收益
	CapitalizedToDa           bool `json:"capitalized_to_da,omitempty"`            //	float	N	资本支出/折旧和摊销
	DebtToAssets              bool `json:"debt_to_assets,omitempty"`               //	float	Y	资产负债率
	AssetsToEqt               bool `json:"assets_to_eqt,omitempty"`                //	float	Y	权益乘数
	DpAssetsToEqt             bool `json:"dp_assets_to_eqt,omitempty"`             //	float	Y	权益乘数(杜邦分析)
	CaToAssets                bool `json:"ca_to_assets,omitempty"`                 //	float	Y	流动资产/总资产
	NcaToAssets               bool `json:"nca_to_assets,omitempty"`                //	float	Y	非流动资产/总资产
	TbassetsToTotalassets     bool `json:"tbassets_to_totalassets,omitempty"`      //	float	Y	有形资产/总资产
	IntToTalcap               bool `json:"int_to_talcap,omitempty"`                //	float	Y	带息债务/全部投入资本
	EqtToTalcapital           bool `json:"eqt_to_talcapital,omitempty"`            //	float	Y	归属于母公司的股东权益/全部投入资本
	CurrentdebtToDebt         bool `json:"currentdebt_to_debt,omitempty"`          //	float	Y	流动负债/负债合计
	LongdebToDebt             bool `json:"longdeb_to_debt,omitempty"`              //	float	Y	非流动负债/负债合计
	OcfToShortdebt            bool `json:"ocf_to_shortdebt,omitempty"`             //	float	Y	经营活动产生的现金流量净额/流动负债
	DebtToEqt                 bool `json:"debt_to_eqt,omitempty"`                  //	float	Y	产权比率
	EqtToDebt                 bool `json:"eqt_to_debt,omitempty"`                  //	float	Y	归属于母公司的股东权益/负债合计
	EqtToInterestdebt         bool `json:"eqt_to_interestdebt,omitempty"`          //	float	Y	归属于母公司的股东权益/带息债务
	TangibleassetToDebt       bool `json:"tangibleasset_to_debt,omitempty"`        //	float	Y	有形资产/负债合计
	TangassetToIntdebt        bool `json:"tangasset_to_intdebt,omitempty"`         //	float	Y	有形资产/带息债务
	TangibleassetToNetdebt    bool `json:"tangibleasset_to_netdebt,omitempty"`     //	float	Y	有形资产/净债务
	OcfToDebt                 bool `json:"ocf_to_debt,omitempty"`                  //	float	Y	经营活动产生的现金流量净额/负债合计
	OcfToInterestdebt         bool `json:"ocf_to_interestdebt,omitempty"`          //	float	N	经营活动产生的现金流量净额/带息债务
	OcfToNetdebt              bool `json:"ocf_to_netdebt,omitempty"`               //	float	N	经营活动产生的现金流量净额/净债务
	EbitToInterest            bool `json:"ebit_to_interest,omitempty"`             //	float	N	已获利息倍数(EBIT/利息费用)
	LongdebtToWorkingcapital  bool `json:"longdebt_to_workingcapital,omitempty"`   //	float	N	长期债务与营运资金比率
	EbitdaToDebt              bool `json:"ebitda_to_debt,omitempty"`               //	float	N	息税折旧摊销前利润/负债合计
	TurnDays                  bool `json:"turn_days,omitempty"`                    //	float	Y	营业周期
	RoaYearly                 bool `json:"roa_yearly,omitempty"`                   //	float	Y	年化总资产净利率
	RoaDp                     bool `json:"roa_dp,omitempty"`                       //	float	Y	总资产净利率(杜邦分析)
	FixedAssets               bool `json:"fixed_assets,omitempty"`                 //	float	Y	固定资产合计
	ProfitPrefinExp           bool `json:"profit_prefin_exp,omitempty"`            //	float	N	扣除财务费用前营业利润
	NonOpProfit               bool `json:"non_op_profit,omitempty"`                //	float	N	非营业利润
	OpToEbt                   bool `json:"op_to_ebt,omitempty"`                    //	float	N	营业利润／利润总额
	NopToEbt                  bool `json:"nop_to_ebt,omitempty"`                   //	float	N	非营业利润／利润总额
	OcfToProfit               bool `json:"ocf_to_profit,omitempty"`                //	float	N	经营活动产生的现金流量净额／营业利润
	CashToLiqdebt             bool `json:"cash_to_liqdebt,omitempty"`              //	float	N	货币资金／流动负债
	CashToLiqdebtWithinterest bool `json:"cash_to_liqdebt_withinterest,omitempty"` //	float	N	货币资金／带息流动负债
	OpToLiqdebt               bool `json:"op_to_liqdebt,omitempty"`                //	float	N	营业利润／流动负债
	OpToDebt                  bool `json:"op_to_debt,omitempty"`                   //	float	N	营业利润／负债合计
	RoicYearly                bool `json:"roic_yearly,omitempty"`                  //	float	N	年化投入资本回报率
	TotalFaTrun               bool `json:"total_fa_trun,omitempty"`                //	float	N	固定资产合计周转率
	ProfitToOp                bool `json:"profit_to_op,omitempty"`                 //	float	Y	利润总额／营业收入
	QOpincome                 bool `json:"q_opincome,omitempty"`                   //	float	N	经营活动单季度净收益
	QInvestincome             bool `json:"q_investincome,omitempty"`               //	float	N	价值变动单季度净收益
	QDtprofit                 bool `json:"q_dtprofit,omitempty"`                   //	float	N	扣除非经常损益后的单季度净利润
	QEps                      bool `json:"q_eps,omitempty"`                        //	float	N	每股收益(单季度)
	QNetprofitMargin          bool `json:"q_netprofit_margin,omitempty"`           //	float	N	销售净利率(单季度)
	QGsprofitMargin           bool `json:"q_gsprofit_margin,omitempty"`            //	float	N	销售毛利率(单季度)
	QExpToSales               bool `json:"q_exp_to_sales,omitempty"`               //	float	N	销售期间费用率(单季度)
	QProfitToGr               bool `json:"q_profit_to_gr,omitempty"`               //	float	N	净利润／营业总收入(单季度)
	QSaleexpToGr              bool `json:"q_saleexp_to_gr,omitempty"`              //	float	Y	销售费用／营业总收入 (单季度)
	QAdminexpToGr             bool `json:"q_adminexp_to_gr,omitempty"`             //	float	N	管理费用／营业总收入 (单季度)
	QFinaexpToGr              bool `json:"q_finaexp_to_gr,omitempty"`              //	float	N	财务费用／营业总收入 (单季度)
	QImpairToGrTtm            bool `json:"q_impair_to_gr_ttm,omitempty"`           //	float	N	资产减值损失／营业总收入(单季度)
	QGcToGr                   bool `json:"q_gc_to_gr,omitempty"`                   //	float	Y	营业总成本／营业总收入 (单季度)
	QOpToGr                   bool `json:"q_op_to_gr,omitempty"`                   //	float	N	营业利润／营业总收入(单季度)
	QRoe                      bool `json:"q_roe,omitempty"`                        //	float	Y	净资产收益率(单季度)
	QDtRoe                    bool `json:"q_dt_roe,omitempty"`                     //	float	Y	净资产单季度收益率(扣除非经常损益)
	QNpta                     bool `json:"q_npta,omitempty"`                       //	float	Y	总资产净利润(单季度)
	QOpincomeToEbt            bool `json:"q_opincome_to_ebt,omitempty"`            //	float	N	经营活动净收益／利润总额(单季度)
	QInvestincomeToEbt        bool `json:"q_investincome_to_ebt,omitempty"`        //	float	N	价值变动净收益／利润总额(单季度)
	QDtprofitToProfit         bool `json:"q_dtprofit_to_profit,omitempty"`         //	float	N	扣除非经常损益后的净利润／净利润(单季度)
	QSalescashToOr            bool `json:"q_salescash_to_or,omitempty"`            //	float	N	销售商品提供劳务收到的现金／营业收入(单季度)
	QOcfToSales               bool `json:"q_ocf_to_sales,omitempty"`               //	float	Y	经营活动产生的现金流量净额／营业收入(单季度)
	QOcfToOr                  bool `json:"q_ocf_to_or,omitempty"`                  //	float	N	经营活动产生的现金流量净额／经营活动净收益(单季度)
	BasicEpsYoy               bool `json:"basic_eps_yoy,omitempty"`                //	float	Y	基本每股收益同比增长率(%)
	DtEpsYoy                  bool `json:"dt_eps_yoy,omitempty"`                   //	float	Y	稀释每股收益同比增长率(%)
	CfpsYoy                   bool `json:"cfps_yoy,omitempty"`                     //	float	Y	每股经营活动产生的现金流量净额同比增长率(%)
	OpYoy                     bool `json:"op_yoy,omitempty"`                       //	float	Y	营业利润同比增长率(%)
	EbtYoy                    bool `json:"ebt_yoy,omitempty"`                      //	float	Y	利润总额同比增长率(%)
	NetprofitYoy              bool `json:"netprofit_yoy,omitempty"`                //	float	Y	归属母公司股东的净利润同比增长率(%)
	DtNetprofitYoy            bool `json:"dt_netprofit_yoy,omitempty"`             //	float	Y	归属母公司股东的净利润-扣除非经常损益同比增长率(%)
	OcfYoy                    bool `json:"ocf_yoy,omitempty"`                      //	float	Y	经营活动产生的现金流量净额同比增长率(%)
	RoeYoy                    bool `json:"roe_yoy,omitempty"`                      //	float	Y	净资产收益率(摊薄)同比增长率(%)
	BpsYoy                    bool `json:"bps_yoy,omitempty"`                      //	float	Y	每股净资产相对年初增长率(%)
	AssetsYoy                 bool `json:"assets_yoy,omitempty"`                   //	float	Y	资产总计相对年初增长率(%)
	EqtYoy                    bool `json:"eqt_yoy,omitempty"`                      //	float	Y	归属母公司的股东权益相对年初增长率(%)
	TrYoy                     bool `json:"tr_yoy,omitempty"`                       //	float	Y	营业总收入同比增长率(%)
	OrYoy                     bool `json:"or_yoy,omitempty"`                       //	float	Y	营业收入同比增长率(%)
	QGrYoy                    bool `json:"q_gr_yoy,omitempty"`                     //	float	N	营业总收入同比增长率(%)(单季度)
	QGrQoq                    bool `json:"q_gr_qoq,omitempty"`                     //	float	N	营业总收入环比增长率(%)(单季度)
	QSalesYoy                 bool `json:"q_sales_yoy,omitempty"`                  //	float	Y	营业收入同比增长率(%)(单季度)
	QSalesQoq                 bool `json:"q_sales_qoq,omitempty"`                  //	float	N	营业收入环比增长率(%)(单季度)
	QOpYoy                    bool `json:"q_op_yoy,omitempty"`                     //	float	N	营业利润同比增长率(%)(单季度)
	QOpQoq                    bool `json:"q_op_qoq,omitempty"`                     //	float	Y	营业利润环比增长率(%)(单季度)
	QProfitYoy                bool `json:"q_profit_yoy,omitempty"`                 //	float	N	净利润同比增长率(%)(单季度)
	QProfitQoq                bool `json:"q_profit_qoq,omitempty"`                 //	float	N	净利润环比增长率(%)(单季度)
	QNetprofitYoy             bool `json:"q_netprofit_yoy,omitempty"`              //	float	N	归属母公司股东的净利润同比增长率(%)(单季度)
	QNetprofitQoq             bool `json:"q_netprofit_qoq,omitempty"`              //	float	N	归属母公司股东的净利润环比增长率(%)(单季度)
	EquityYoy                 bool `json:"equity_yoy,omitempty"`                   //	float	Y	净资产同比增长率
	RdExp                     bool `json:"rd_exp,omitempty"`                       //	float	N	研发费用
	UpdateFlag                bool `json:"update_flag,omitempty"`                  //	str	N	更新标识
}

func (item FinaIndicatorItems) All() FinaIndicatorItems {
	item.TsCode = true
	item.AnnDate = true
	item.EndDate = true
	item.Eps = true
	item.DtEps = true
	item.TotalRevenuePs = true
	item.RevenuePs = true
	item.CapitalResePs = true
	item.SurplusResePs = true
	item.UndistProfitPs = true
	item.ExtraItem = true
	item.ProfitDedt = true
	item.GrossMargin = true
	item.CurrentRatio = true
	item.QuickRatio = true
	item.CashRatio = true
	item.InvturnDays = true
	item.ArturnDays = true
	item.InvTurn = true
	item.ArTurn = true
	item.CaTurn = true
	item.FaTurn = true
	item.AssetsTurn = true
	item.OpIncome = true
	item.ValuechangeIncome = true
	item.InterstIncome = true
	item.Daa = true
	item.Ebit = true
	item.Ebitda = true
	item.Fcff = true
	item.Fcfe = true
	item.CurrentExint = true
	item.NoncurrentExint = true
	item.Interestdebt = true
	item.Netdebt = true
	item.TangibleAsset = true
	item.WorkingCapital = true
	item.NetworkingCapital = true
	item.InvestCapital = true
	item.RetainedEarnings = true
	item.Diluted2Eps = true
	item.Bps = true
	item.Ocfps = true
	item.Retainedps = true
	item.Cfps = true
	item.EbitPs = true
	item.FcffPs = true
	item.FcfePs = true
	item.NetprofitMargin = true
	item.GrossprofitMargin = true
	item.CogsOfSales = true
	item.ExpenseOfSales = true
	item.ProfitToGr = true
	item.SaleexpToGr = true
	item.AdminexpOfGr = true
	item.FinaexpOfGr = true
	item.ImpaiTtm = true
	item.GcOfGr = true
	item.OpOfGr = true
	item.EbitOfGr = true
	item.Roe = true
	item.RoeWaa = true
	item.RoeDt = true
	item.Roa = true
	item.Npta = true
	item.Roic = true
	item.RoeYearly = true
	item.Roa2Yearly = true
	item.RoeAvg = true
	item.OpincomeOfEbt = true
	item.InvestincomeOfEbt = true
	item.NOpProfitOfEbt = true
	item.TaxToEbt = true
	item.DtprofitToProfit = true
	item.SalescashToOr = true
	item.OcfToOr = true
	item.OcfToOpincome = true
	item.CapitalizedToDa = true
	item.DebtToAssets = true
	item.AssetsToEqt = true
	item.DpAssetsToEqt = true
	item.CaToAssets = true
	item.NcaToAssets = true
	item.TbassetsToTotalassets = true
	item.IntToTalcap = true
	item.EqtToTalcapital = true
	item.CurrentdebtToDebt = true
	item.LongdebToDebt = true
	item.OcfToShortdebt = true
	item.DebtToEqt = true
	item.EqtToDebt = true
	item.EqtToInterestdebt = true
	item.TangibleassetToDebt = true
	item.TangassetToIntdebt = true
	item.TangibleassetToNetdebt = true
	item.OcfToDebt = true
	item.OcfToInterestdebt = true
	item.OcfToNetdebt = true
	item.EbitToInterest = true
	item.LongdebtToWorkingcapital = true
	item.EbitdaToDebt = true
	item.TurnDays = true
	item.RoaYearly = true
	item.RoaDp = true
	item.FixedAssets = true
	item.ProfitPrefinExp = true
	item.NonOpProfit = true
	item.OpToEbt = true
	item.NopToEbt = true
	item.OcfToProfit = true
	item.CashToLiqdebt = true
	item.CashToLiqdebtWithinterest = true
	item.OpToLiqdebt = true
	item.OpToDebt = true
	item.RoicYearly = true
	item.TotalFaTrun = true
	item.ProfitToOp = true
	item.QOpincome = true
	item.QInvestincome = true
	item.QDtprofit = true
	item.QEps = true
	item.QNetprofitMargin = true
	item.QGsprofitMargin = true
	item.QExpToSales = true
	item.QProfitToGr = true
	item.QSaleexpToGr = true
	item.QAdminexpToGr = true
	item.QFinaexpToGr = true
	item.QImpairToGrTtm = true
	item.QGcToGr = true
	item.QOpToGr = true
	item.QRoe = true
	item.QDtRoe = true
	item.QNpta = true
	item.QOpincomeToEbt = true
	item.QInvestincomeToEbt = true
	item.QDtprofitToProfit = true
	item.QSalescashToOr = true
	item.QOcfToSales = true
	item.QOcfToOr = true
	item.BasicEpsYoy = true
	item.DtEpsYoy = true
	item.CfpsYoy = true
	item.OpYoy = true
	item.EbtYoy = true
	item.NetprofitYoy = true
	item.DtNetprofitYoy = true
	item.OcfYoy = true
	item.RoeYoy = true
	item.BpsYoy = true
	item.AssetsYoy = true
	item.EqtYoy = true
	item.TrYoy = true
	item.OrYoy = true
	item.QGrYoy = true
	item.QGrQoq = true
	item.QSalesYoy = true
	item.QSalesQoq = true
	item.QOpYoy = true
	item.QOpQoq = true
	item.QProfitYoy = true
	item.QProfitQoq = true
	item.QNetprofitYoy = true
	item.QNetprofitQoq = true
	item.EquityYoy = true
	item.RdExp = true
	item.UpdateFlag = true
	return item
}

type FinaIndicatorData struct {
	TsCode                    string  `json:"ts_code,omitempty"  gorm:"index"`        //	str	Y	TS代码
	AnnDate                   string  `json:"ann_date,omitempty"`                     //	str	Y	公告日期
	EndDate                   string  `json:"end_date,omitempty"`                     //	str	Y	报告期
	Eps                       float64 `json:"eps,omitempty"`                          //	float	Y	基本每股收益
	DtEps                     float64 `json:"dt_eps,omitempty"`                       //	float	Y	稀释每股收益
	TotalRevenuePs            float64 `json:"total_revenue_ps,omitempty"`             //	float	Y	每股营业总收入
	RevenuePs                 float64 `json:"revenue_ps,omitempty"`                   //	float	Y	每股营业收入
	CapitalResePs             float64 `json:"capital_rese_ps,omitempty"`              //	float	Y	每股资本公积
	SurplusResePs             float64 `json:"surplus_rese_ps,omitempty"`              //	float	Y	每股盈余公积
	UndistProfitPs            float64 `json:"undist_profit_ps,omitempty"`             //	float	Y	每股未分配利润
	ExtraItem                 float64 `json:"extra_item,omitempty"`                   //	float	Y	非经常性损益
	ProfitDedt                float64 `json:"profit_dedt,omitempty"`                  //	float	Y	扣除非经常性损益后的净利润
	GrossMargin               float64 `json:"gross_margin,omitempty"`                 //	float	Y	毛利
	CurrentRatio              float64 `json:"current_ratio,omitempty"`                //	float	Y	流动比率
	QuickRatio                float64 `json:"quick_ratio,omitempty"`                  //	float	Y	速动比率
	CashRatio                 float64 `json:"cash_ratio,omitempty"`                   //	float	Y	保守速动比率
	InvturnDays               float64 `json:"invturn_days,omitempty"`                 //	float	N	存货周转天数
	ArturnDays                float64 `json:"arturn_days,omitempty"`                  //	float	N	应收账款周转天数
	InvTurn                   float64 `json:"inv_turn,omitempty"`                     //	float	N	存货周转率
	ArTurn                    float64 `json:"ar_turn,omitempty"`                      //	float	Y	应收账款周转率
	CaTurn                    float64 `json:"ca_turn,omitempty"`                      //	float	Y	流动资产周转率
	FaTurn                    float64 `json:"fa_turn,omitempty"`                      //	float	Y	固定资产周转率
	AssetsTurn                float64 `json:"assets_turn,omitempty"`                  //	float	Y	总资产周转率
	OpIncome                  float64 `json:"op_income,omitempty"`                    //	float	Y	经营活动净收益
	ValuechangeIncome         float64 `json:"valuechange_income,omitempty"`           //	float	N	价值变动净收益
	InterstIncome             float64 `json:"interst_income,omitempty"`               //	float	N	利息费用
	Daa                       float64 `json:"daa,omitempty"`                          //	float	N	折旧与摊销
	Ebit                      float64 `json:"ebit,omitempty"`                         //	float	Y	息税前利润
	Ebitda                    float64 `json:"ebitda,omitempty"`                       //	float	Y	息税折旧摊销前利润
	Fcff                      float64 `json:"fcff,omitempty"`                         //	float	Y	企业自由现金流量
	Fcfe                      float64 `json:"fcfe,omitempty"`                         //	float	Y	股权自由现金流量
	CurrentExint              float64 `json:"current_exint,omitempty"`                //	float	Y	无息流动负债
	NoncurrentExint           float64 `json:"noncurrent_exint,omitempty"`             //	float	Y	无息非流动负债
	Interestdebt              float64 `json:"interestdebt,omitempty"`                 //	float	Y	带息债务
	Netdebt                   float64 `json:"netdebt,omitempty"`                      //	float	Y	净债务
	TangibleAsset             float64 `json:"tangible_asset,omitempty"`               //	float	Y	有形资产
	WorkingCapital            float64 `json:"working_capital,omitempty"`              //	float	Y	营运资金
	NetworkingCapital         float64 `json:"networking_capital,omitempty"`           //	float	Y	营运流动资本
	InvestCapital             float64 `json:"invest_capital,omitempty"`               //	float	Y	全部投入资本
	RetainedEarnings          float64 `json:"retained_earnings,omitempty"`            //	float	Y	留存收益
	Diluted2Eps               float64 `json:"diluted2_eps,omitempty"`                 //	float	Y	期末摊薄每股收益
	Bps                       float64 `json:"bps,omitempty"`                          //	float	Y	每股净资产
	Ocfps                     float64 `json:"ocfps,omitempty"`                        //	float	Y	每股经营活动产生的现金流量净额
	Retainedps                float64 `json:"retainedps,omitempty"`                   //	float	Y	每股留存收益
	Cfps                      float64 `json:"cfps,omitempty"`                         //	float	Y	每股现金流量净额
	EbitPs                    float64 `json:"ebit_ps,omitempty"`                      //	float	Y	每股息税前利润
	FcffPs                    float64 `json:"fcff_ps,omitempty"`                      //	float	Y	每股企业自由现金流量
	FcfePs                    float64 `json:"fcfe_ps,omitempty"`                      //	float	Y	每股股东自由现金流量
	NetprofitMargin           float64 `json:"netprofit_margin,omitempty"`             //	float	Y	销售净利率
	GrossprofitMargin         float64 `json:"grossprofit_margin,omitempty"`           //	float	Y	销售毛利率
	CogsOfSales               float64 `json:"cogs_of_sales,omitempty"`                //	float	Y	销售成本率
	ExpenseOfSales            float64 `json:"expense_of_sales,omitempty"`             //	float	Y	销售期间费用率
	ProfitToGr                float64 `json:"profit_to_gr,omitempty"`                 //	float	Y	净利润/营业总收入
	SaleexpToGr               float64 `json:"saleexp_to_gr,omitempty"`                //	float	Y	销售费用/营业总收入
	AdminexpOfGr              float64 `json:"adminexp_of_gr,omitempty"`               //	float	Y	管理费用/营业总收入
	FinaexpOfGr               float64 `json:"finaexp_of_gr,omitempty"`                //	float	Y	财务费用/营业总收入
	ImpaiTtm                  float64 `json:"impai_ttm,omitempty"`                    //	float	Y	资产减值损失/营业总收入
	GcOfGr                    float64 `json:"gc_of_gr,omitempty"`                     //	float	Y	营业总成本/营业总收入
	OpOfGr                    float64 `json:"op_of_gr,omitempty"`                     //	float	Y	营业利润/营业总收入
	EbitOfGr                  float64 `json:"ebit_of_gr,omitempty"`                   //	float	Y	息税前利润/营业总收入
	Roe                       float64 `json:"roe,omitempty"`                          //	float	Y	净资产收益率
	RoeWaa                    float64 `json:"roe_waa,omitempty"`                      //	float	Y	加权平均净资产收益率
	RoeDt                     float64 `json:"roe_dt,omitempty"`                       //	float	Y	净资产收益率(扣除非经常损益)
	Roa                       float64 `json:"roa,omitempty"`                          //	float	Y	总资产报酬率
	Npta                      float64 `json:"npta,omitempty"`                         //	float	Y	总资产净利润
	Roic                      float64 `json:"roic,omitempty"`                         //	float	Y	投入资本回报率
	RoeYearly                 float64 `json:"roe_yearly,omitempty"`                   //	float	Y	年化净资产收益率
	Roa2Yearly                float64 `json:"roa2_yearly,omitempty"`                  //	float	Y	年化总资产报酬率
	RoeAvg                    float64 `json:"roe_avg,omitempty"`                      //	float	N	平均净资产收益率(增发条件)
	OpincomeOfEbt             float64 `json:"opincome_of_ebt,omitempty"`              //	float	N	经营活动净收益/利润总额
	InvestincomeOfEbt         float64 `json:"investincome_of_ebt,omitempty"`          //	float	N	价值变动净收益/利润总额
	NOpProfitOfEbt            float64 `json:"n_op_profit_of_ebt,omitempty"`           //	float	N	营业外收支净额/利润总额
	TaxToEbt                  float64 `json:"tax_to_ebt,omitempty"`                   //	float	N	所得税/利润总额
	DtprofitToProfit          float64 `json:"dtprofit_to_profit,omitempty"`           //	float	N	扣除非经常损益后的净利润/净利润
	SalescashToOr             float64 `json:"salescash_to_or,omitempty"`              //	float	N	销售商品提供劳务收到的现金/营业收入
	OcfToOr                   float64 `json:"ocf_to_or,omitempty"`                    //	float	N	经营活动产生的现金流量净额/营业收入
	OcfToOpincome             float64 `json:"ocf_to_opincome,omitempty"`              //	float	N	经营活动产生的现金流量净额/经营活动净收益
	CapitalizedToDa           float64 `json:"capitalized_to_da,omitempty"`            //	float	N	资本支出/折旧和摊销
	DebtToAssets              float64 `json:"debt_to_assets,omitempty"`               //	float	Y	资产负债率
	AssetsToEqt               float64 `json:"assets_to_eqt,omitempty"`                //	float	Y	权益乘数
	DpAssetsToEqt             float64 `json:"dp_assets_to_eqt,omitempty"`             //	float	Y	权益乘数(杜邦分析)
	CaToAssets                float64 `json:"ca_to_assets,omitempty"`                 //	float	Y	流动资产/总资产
	NcaToAssets               float64 `json:"nca_to_assets,omitempty"`                //	float	Y	非流动资产/总资产
	TbassetsToTotalassets     float64 `json:"tbassets_to_totalassets,omitempty"`      //	float	Y	有形资产/总资产
	IntToTalcap               float64 `json:"int_to_talcap,omitempty"`                //	float	Y	带息债务/全部投入资本
	EqtToTalcapital           float64 `json:"eqt_to_talcapital,omitempty"`            //	float	Y	归属于母公司的股东权益/全部投入资本
	CurrentdebtToDebt         float64 `json:"currentdebt_to_debt,omitempty"`          //	float	Y	流动负债/负债合计
	LongdebToDebt             float64 `json:"longdeb_to_debt,omitempty"`              //	float	Y	非流动负债/负债合计
	OcfToShortdebt            float64 `json:"ocf_to_shortdebt,omitempty"`             //	float	Y	经营活动产生的现金流量净额/流动负债
	DebtToEqt                 float64 `json:"debt_to_eqt,omitempty"`                  //	float	Y	产权比率
	EqtToDebt                 float64 `json:"eqt_to_debt,omitempty"`                  //	float	Y	归属于母公司的股东权益/负债合计
	EqtToInterestdebt         float64 `json:"eqt_to_interestdebt,omitempty"`          //	float	Y	归属于母公司的股东权益/带息债务
	TangibleassetToDebt       float64 `json:"tangibleasset_to_debt,omitempty"`        //	float	Y	有形资产/负债合计
	TangassetToIntdebt        float64 `json:"tangasset_to_intdebt,omitempty"`         //	float	Y	有形资产/带息债务
	TangibleassetToNetdebt    float64 `json:"tangibleasset_to_netdebt,omitempty"`     //	float	Y	有形资产/净债务
	OcfToDebt                 float64 `json:"ocf_to_debt,omitempty"`                  //	float	Y	经营活动产生的现金流量净额/负债合计
	OcfToInterestdebt         float64 `json:"ocf_to_interestdebt,omitempty"`          //	float	N	经营活动产生的现金流量净额/带息债务
	OcfToNetdebt              float64 `json:"ocf_to_netdebt,omitempty"`               //	float	N	经营活动产生的现金流量净额/净债务
	EbitToInterest            float64 `json:"ebit_to_interest,omitempty"`             //	float	N	已获利息倍数(EBIT/利息费用)
	LongdebtToWorkingcapital  float64 `json:"longdebt_to_workingcapital,omitempty"`   //	float	N	长期债务与营运资金比率
	EbitdaToDebt              float64 `json:"ebitda_to_debt,omitempty"`               //	float	N	息税折旧摊销前利润/负债合计
	TurnDays                  float64 `json:"turn_days,omitempty"`                    //	float	Y	营业周期
	RoaYearly                 float64 `json:"roa_yearly,omitempty"`                   //	float	Y	年化总资产净利率
	RoaDp                     float64 `json:"roa_dp,omitempty"`                       //	float	Y	总资产净利率(杜邦分析)
	FixedAssets               float64 `json:"fixed_assets,omitempty"`                 //	float	Y	固定资产合计
	ProfitPrefinExp           float64 `json:"profit_prefin_exp,omitempty"`            //	float	N	扣除财务费用前营业利润
	NonOpProfit               float64 `json:"non_op_profit,omitempty"`                //	float	N	非营业利润
	OpToEbt                   float64 `json:"op_to_ebt,omitempty"`                    //	float	N	营业利润／利润总额
	NopToEbt                  float64 `json:"nop_to_ebt,omitempty"`                   //	float	N	非营业利润／利润总额
	OcfToProfit               float64 `json:"ocf_to_profit,omitempty"`                //	float	N	经营活动产生的现金流量净额／营业利润
	CashToLiqdebt             float64 `json:"cash_to_liqdebt,omitempty"`              //	float	N	货币资金／流动负债
	CashToLiqdebtWithinterest float64 `json:"cash_to_liqdebt_withinterest,omitempty"` //	float	N	货币资金／带息流动负债
	OpToLiqdebt               float64 `json:"op_to_liqdebt,omitempty"`                //	float	N	营业利润／流动负债
	OpToDebt                  float64 `json:"op_to_debt,omitempty"`                   //	float	N	营业利润／负债合计
	RoicYearly                float64 `json:"roic_yearly,omitempty"`                  //	float	N	年化投入资本回报率
	TotalFaTrun               float64 `json:"total_fa_trun,omitempty"`                //	float	N	固定资产合计周转率
	ProfitToOp                float64 `json:"profit_to_op,omitempty"`                 //	float	Y	利润总额／营业收入
	QOpincome                 float64 `json:"q_opincome,omitempty"`                   //	float	N	经营活动单季度净收益
	QInvestincome             float64 `json:"q_investincome,omitempty"`               //	float	N	价值变动单季度净收益
	QDtprofit                 float64 `json:"q_dtprofit,omitempty"`                   //	float	N	扣除非经常损益后的单季度净利润
	QEps                      float64 `json:"q_eps,omitempty"`                        //	float	N	每股收益(单季度)
	QNetprofitMargin          float64 `json:"q_netprofit_margin,omitempty"`           //	float	N	销售净利率(单季度)
	QGsprofitMargin           float64 `json:"q_gsprofit_margin,omitempty"`            //	float	N	销售毛利率(单季度)
	QExpToSales               float64 `json:"q_exp_to_sales,omitempty"`               //	float	N	销售期间费用率(单季度)
	QProfitToGr               float64 `json:"q_profit_to_gr,omitempty"`               //	float	N	净利润／营业总收入(单季度)
	QSaleexpToGr              float64 `json:"q_saleexp_to_gr,omitempty"`              //	float	Y	销售费用／营业总收入 (单季度)
	QAdminexpToGr             float64 `json:"q_adminexp_to_gr,omitempty"`             //	float	N	管理费用／营业总收入 (单季度)
	QFinaexpToGr              float64 `json:"q_finaexp_to_gr,omitempty"`              //	float	N	财务费用／营业总收入 (单季度)
	QImpairToGrTtm            float64 `json:"q_impair_to_gr_ttm,omitempty"`           //	float	N	资产减值损失／营业总收入(单季度)
	QGcToGr                   float64 `json:"q_gc_to_gr,omitempty"`                   //	float	Y	营业总成本／营业总收入 (单季度)
	QOpToGr                   float64 `json:"q_op_to_gr,omitempty"`                   //	float	N	营业利润／营业总收入(单季度)
	QRoe                      float64 `json:"q_roe,omitempty"`                        //	float	Y	净资产收益率(单季度)
	QDtRoe                    float64 `json:"q_dt_roe,omitempty"`                     //	float	Y	净资产单季度收益率(扣除非经常损益)
	QNpta                     float64 `json:"q_npta,omitempty"`                       //	float	Y	总资产净利润(单季度)
	QOpincomeToEbt            float64 `json:"q_opincome_to_ebt,omitempty"`            //	float	N	经营活动净收益／利润总额(单季度)
	QInvestincomeToEbt        float64 `json:"q_investincome_to_ebt,omitempty"`        //	float	N	价值变动净收益／利润总额(单季度)
	QDtprofitToProfit         float64 `json:"q_dtprofit_to_profit,omitempty"`         //	float	N	扣除非经常损益后的净利润／净利润(单季度)
	QSalescashToOr            float64 `json:"q_salescash_to_or,omitempty"`            //	float	N	销售商品提供劳务收到的现金／营业收入(单季度)
	QOcfToSales               float64 `json:"q_ocf_to_sales,omitempty"`               //	float	Y	经营活动产生的现金流量净额／营业收入(单季度)
	QOcfToOr                  float64 `json:"q_ocf_to_or,omitempty"`                  //	float	N	经营活动产生的现金流量净额／经营活动净收益(单季度)
	BasicEpsYoy               float64 `json:"basic_eps_yoy,omitempty"`                //	float	Y	基本每股收益同比增长率(%)
	DtEpsYoy                  float64 `json:"dt_eps_yoy,omitempty"`                   //	float	Y	稀释每股收益同比增长率(%)
	CfpsYoy                   float64 `json:"cfps_yoy,omitempty"`                     //	float	Y	每股经营活动产生的现金流量净额同比增长率(%)
	OpYoy                     float64 `json:"op_yoy,omitempty"`                       //	float	Y	营业利润同比增长率(%)
	EbtYoy                    float64 `json:"ebt_yoy,omitempty"`                      //	float	Y	利润总额同比增长率(%)
	NetprofitYoy              float64 `json:"netprofit_yoy,omitempty"`                //	float	Y	归属母公司股东的净利润同比增长率(%)
	DtNetprofitYoy            float64 `json:"dt_netprofit_yoy,omitempty"`             //	float	Y	归属母公司股东的净利润-扣除非经常损益同比增长率(%)
	OcfYoy                    float64 `json:"ocf_yoy,omitempty"`                      //	float	Y	经营活动产生的现金流量净额同比增长率(%)
	RoeYoy                    float64 `json:"roe_yoy,omitempty"`                      //	float	Y	净资产收益率(摊薄)同比增长率(%)
	BpsYoy                    float64 `json:"bps_yoy,omitempty"`                      //	float	Y	每股净资产相对年初增长率(%)
	AssetsYoy                 float64 `json:"assets_yoy,omitempty"`                   //	float	Y	资产总计相对年初增长率(%)
	EqtYoy                    float64 `json:"eqt_yoy,omitempty"`                      //	float	Y	归属母公司的股东权益相对年初增长率(%)
	TrYoy                     float64 `json:"tr_yoy,omitempty"`                       //	float	Y	营业总收入同比增长率(%)
	OrYoy                     float64 `json:"or_yoy,omitempty"`                       //	float	Y	营业收入同比增长率(%)
	QGrYoy                    float64 `json:"q_gr_yoy,omitempty"`                     //	float	N	营业总收入同比增长率(%)(单季度)
	QGrQoq                    float64 `json:"q_gr_qoq,omitempty"`                     //	float	N	营业总收入环比增长率(%)(单季度)
	QSalesYoy                 float64 `json:"q_sales_yoy,omitempty"`                  //	float	Y	营业收入同比增长率(%)(单季度)
	QSalesQoq                 float64 `json:"q_sales_qoq,omitempty"`                  //	float	N	营业收入环比增长率(%)(单季度)
	QOpYoy                    float64 `json:"q_op_yoy,omitempty"`                     //	float	N	营业利润同比增长率(%)(单季度)
	QOpQoq                    float64 `json:"q_op_qoq,omitempty"`                     //	float	Y	营业利润环比增长率(%)(单季度)
	QProfitYoy                float64 `json:"q_profit_yoy,omitempty"`                 //	float	N	净利润同比增长率(%)(单季度)
	QProfitQoq                float64 `json:"q_profit_qoq,omitempty"`                 //	float	N	净利润环比增长率(%)(单季度)
	QNetprofitYoy             float64 `json:"q_netprofit_yoy,omitempty"`              //	float	N	归属母公司股东的净利润同比增长率(%)(单季度)
	QNetprofitQoq             float64 `json:"q_netprofit_qoq,omitempty"`              //	float	N	归属母公司股东的净利润环比增长率(%)(单季度)
	EquityYoy                 float64 `json:"equity_yoy,omitempty"`                   //	float	Y	净资产同比增长率
	RdExp                     float64 `json:"rd_exp,omitempty"`                       //	float	N	研发费用
	UpdateFlag                string  `json:"update_flag,omitempty"`                  //	str	N	更新标识
}

func AssembleFinaIndicatorData(tsRsp *TushareResponse) []*FinaIndicatorData {
	tsData := []*FinaIndicatorData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(FinaIndicatorData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 获取上市公司财务指标数据,为避免服务器压力,现阶段每次请求最多返回60条记录,可通过设置日期多次请求获取更多数据,用户需要至少800积分才可以调取,具体请参阅积分获取办法 https://tushare.pro/document/1?doc_id=13
func (ts *TuShare) FinaIndicator(params FinanceRequest, items FinaIndicatorItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "fina_indicator",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type FinaAuditItems struct {
	TsCode      bool `json:"ts_code,omitempty"`      // str	Y	TS股票代码
	AnnDate     bool `json:"ann_date,omitempty"`     // str	公告日期
	EndDate     bool `json:"end_date,omitempty"`     // str	报告期
	AuditResult bool `json:"audit_result,omitempty"` // str	审计结果
	AuditFees   bool `json:"audit_fees,omitempty"`   // float	审计总费用(元)
	AuditAgency bool `json:"audit_agency,omitempty"` // str	会计事务所
	AuditSign   bool `json:"audit_sign,omitempty"`   // str	签字会计师
}

func (item FinaAuditItems) All() FinaAuditItems {
	item.TsCode = true
	item.AnnDate = true
	item.EndDate = true
	item.AuditResult = true
	item.AuditFees = true
	item.AuditAgency = true
	item.AuditSign = true
	return item
}

type FinaAuditData struct {
	TsCode      string  `json:"ts_code,omitempty"`      // str	Y	TS股票代码
	AnnDate     string  `json:"ann_date,omitempty"`     // str	公告日期
	EndDate     string  `json:"end_date,omitempty"`     // str	报告期
	AuditResult string  `json:"audit_result,omitempty"` // str	审计结果
	AuditFees   float64 `json:"audit_fees,omitempty"`   // float	审计总费用(元)
	AuditAgency string  `json:"audit_agency,omitempty"` // str	会计事务所
	AuditSign   string  `json:"audit_sign,omitempty"`   // str	签字会计师
}

func AssembleFinaAuditData(tsRsp *TushareResponse) []*FinaAuditData {
	tsData := []*FinaAuditData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(FinaAuditData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 获取上市公司定期财务审计意见数据,用户需要至少500积分才可以调取,具体请参阅积分获取办法 https://tushare.pro/document/1?doc_id=13
func (ts *TuShare) FinaAudit(params FinanceRequest, items FinaAuditItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "fina_audit",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type FinaMainBZRequest struct {
	TsCode    string `json:"ts_code,omitempty"`    // str	Y	股票代码
	Period    string `json:"period,omitempty"`     // str	N	报告期(每个季度最后一天的日期,比如20171231表示年报)
	Type      string `json:"type,omitempty"`       // str	N	类型：P按产品 D按地区(请输入大写字母P或者D)
	StartDate string `json:"start_date,omitempty"` // str	N	报告期开始日期
	EndDate   string `json:"end_date,omitempty"`   // str	N	报告期结束日期
}

type FinaMainBZItems struct {
	TsCode     bool `json:"ts_code,omitempty"`     // str	TS代码
	EndDate    bool `json:"end_date,omitempty"`    // str	报告期
	BzItem     bool `json:"bz_item,omitempty"`     // str	主营业务来源
	BzSales    bool `json:"bz_sales,omitempty"`    // float	主营业务收入(元)
	BzProfit   bool `json:"bz_profit,omitempty"`   // float	主营业务利润(元)
	BzCost     bool `json:"bz_cost,omitempty"`     // float	主营业务成本(元)
	CurrType   bool `json:"curr_type,omitempty"`   // str	货币代码
	UpdateFlag bool `json:"update_flag,omitempty"` // str	是否更新
}

func (item FinaMainBZItems) All() FinaMainBZItems {
	item.TsCode = true
	item.EndDate = true
	item.BzItem = true
	item.BzSales = true
	item.BzProfit = true
	item.BzCost = true
	item.CurrType = true
	item.UpdateFlag = true
	return item
}

type FinaMainBZData struct {
	TsCode     string  `json:"ts_code,omitempty"`     // str	TS代码
	EndDate    string  `json:"end_date,omitempty"`    // str	报告期
	BzItem     string  `json:"bz_item,omitempty"`     // str	主营业务来源
	BzSales    float64 `json:"bz_sales,omitempty"`    // float	主营业务收入(元)
	BzProfit   float64 `json:"bz_profit,omitempty"`   // float	主营业务利润(元)
	BzCost     float64 `json:"bz_cost,omitempty"`     // float	主营业务成本(元)
	CurrType   string  `json:"curr_type,omitempty"`   // str	货币代码
	UpdateFlag string  `json:"update_flag,omitempty"` // str	是否更新
}

func AssembleFinaMainBZData(tsRsp *TushareResponse) []*FinaMainBZData {
	tsData := []*FinaMainBZData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(FinaMainBZData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 获得上市公司主营业务构成,分地区和产品两种方式,用户需要至少900积分才可以调取,具体请参阅积分获取办法 https://tushare.pro/document/1?doc_id=13
func (ts *TuShare) FinaMainBZ(params FinaMainBZRequest, items FinaMainBZItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "fina_mainbz",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}

type DisclosureDateRequest struct {
	TsCode     string `json:"ts_code,omitempty"`     // str	N	TS股票代码
	EndDate    string `json:"end_date,omitempty"`    // str	N	财报周期(比如20181231表示2018年年报,20180630表示中报)
	PreDate    string `json:"pre_date,omitempty"`    // str	N	计划披露日期
	ActualDate string `json:"actual_date,omitempty"` // str	N	实际披露日期
}

type DisclosureDateItems struct {
	TsCode     bool `json:"ts_code,omitempty"`     // str	Y	TS股票代码
	AnnDate    bool `json:"ann_date,omitempty"`    // str	Y	最新披露公告日
	EndDate    bool `json:"end_date,omitempty"`    // str	Y	报告期
	PreDate    bool `json:"pre_date,omitempty"`    // str	Y	预计披露日期
	ActualDate bool `json:"actual_date,omitempty"` // str	Y	实际披露日期
	ModifyDate bool `json:"modify_date,omitempty"` // str	N	披露日期修正记录
}

func (item DisclosureDateItems) All() DisclosureDateItems {
	item.TsCode = true
	item.AnnDate = true
	item.EndDate = true
	item.PreDate = true
	item.ActualDate = true
	item.ModifyDate = true
	return item
}

type DisclosureDateData struct {
	TsCode     string `json:"ts_code,omitempty"`     // str	Y	TS股票代码
	AnnDate    string `json:"ann_date,omitempty"`    // str	Y	最新披露公告日
	EndDate    string `json:"end_date,omitempty"`    // str	Y	报告期
	PreDate    string `json:"pre_date,omitempty"`    // str	Y	预计披露日期
	ActualDate string `json:"actual_date,omitempty"` // str	Y	实际披露日期
	ModifyDate string `json:"modify_date,omitempty"` // str	N	披露日期修正记录
}

func AssembleDisclosureDateData(tsRsp *TushareResponse) []*DisclosureDateData {
	tsData := []*DisclosureDateData{}
	for _, data := range tsRsp.Data.Items {
		body, err := ReflectResponseData(tsRsp.Data.Fields, data)
		if err == nil {
			n := new(DisclosureDateData)
			err = json.Unmarshal(body, &n)
			if err == nil {
				tsData = append(tsData, n)
			}
		}
	}
	return tsData
}

// 获取财报披露计划日期,单次最大3000,总量不限制,用户需要至少500积分才可以调取,具体请参阅积分获取办法 https://tushare.pro/document/1?doc_id=13
func (ts *TuShare) DisclosureDate(params DisclosureDateRequest, items DisclosureDateItems) (tsRsp *TushareResponse, err error) {
	req := &TushareRequest{
		APIName: "disclosure_date",
		Token:   ts.token,
		Params:  buildParams(params),
		Fields:  reflectFields(items),
	}
	return requestTushare(ts.client, http.MethodPost, req)
}
