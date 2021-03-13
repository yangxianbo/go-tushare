GO-TuShare
=======

官方网站：https://tushare.pro/

本项目作为一个普通爱好者做的GO版本SDK，目前实现了沪深股票,指数,公募基金 2000积分(含)的所有接口调用，并且增加了Boll线数据的获取，后续会逐步增加其他接口和更多技术指标数据

已经支持了的接口
=======
```golang
type TuShare
    func NewTuShare(token string) (ts *TuShare)
    func (ts *TuShare) AdjFactor(params QuotationRequest, items AdjFactorItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) BakDaily(params BakDailyRequest, items BakDailyItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) BalanceSheet(params SheetRequest, items BalanceSheetItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) BlockTrade(params BlockTradeRequest, items BlockTradeItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) CashFlow(params SheetRequest, items CashFlowItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) Concept(params ConceptRequest, items ConceptItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) ConceptDetail(params ConceptDetailRequest, items ConceptDetailItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) Daily(params QuotationRequest, items QuotationItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) DailyBasic(params QuotationRequest, items DailyBasicItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) DailyInfo(params DailyInfoRequest, items DailyInfoItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) DisclosureDate(params DisclosureDateRequest, items DisclosureDateItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) Dividend(params DividendRequest, items DividendItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) Express(params FinanceRequest, items ExpressItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) FinaAudit(params FinanceRequest, items FinaAuditItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) FinaIndicator(params FinanceRequest, items FinaIndicatorItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) FinaMainBZ(params FinaMainBZRequest, items FinaMainBZItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) Forecast(params ForecastRequest, items ForecastItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) FundAdj(params FundAdjRequest, items FundAdjItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) FundBasic(params FundBasicRequest, items FundBasicItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) FundCompany(items FundCompanyItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) FundDaily(params QuotationRequest, items FundDailyItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) FundDiv(params FundDivRequest, items FundDivItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) FundManager(params FundManagerRequest, items FundManagerItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) FundNav(params FundNavRequest, items FundNavItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) FundPortfolio(params FundPortfolioRequest, items FundPortfolioItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) FundShare(params FundShareRequest, items FundShareItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) GGTDaily(params GGTDailyRequest, items GGTDailyItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) GGTTop10(params GGTTop10Request, items GGTTop10Items) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) HKHold(params HKHoldRequest, items HKHoldItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) HSGTTop10(params HSGTTop10Request, items HSGTTop10Items) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) Health() (err error)
    func (ts *TuShare) HsConst(params HsConstRequest, items HsConstItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) Income(params SheetRequest, items IncomeItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) IndexBasic(params IndexBasicRequest, items IndexBasicItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) IndexClassify(params IndexClassifyRequest, items IndexClassifyItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) IndexDaily(params QuotationRequest, items QuotationItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) IndexDailyBasic(params QuotationRequest, items IndexDailyBasicItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) IndexGlobal(params QuotationRequest, items IndexGlobalItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) IndexMember(params IndexMemberRequest, items IndexMemberItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) IndexMonthly(params QuotationRequest, items QuotationItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) IndexWeekly(params QuotationRequest, items QuotationItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) IndexWeight(params IndexWeightRequest, items IndexWeightItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) LimitList(params LimitListRequest, items LimitListItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) Margin(params MarginRequest, items MarginItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) MarginDetail(params MarginDetailRequest, items MarginDetailItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) Moneyflow(params QuotationRequest, items MoneyflowItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) MoneyflowHSGT(params MoneyflowHSGTRequest, items MoneyflowHSGTItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) Monthly(params QuotationRequest, items QuotationItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) NameChange(params NameChangeRequest, items NameChangeItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) NewShare(params NewShareCompanyRequest, items NewShareCompanyItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) PledgeDetail(params PledgeDetailRequest, items PledgeDetailItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) PledgeStat(params PledgeStatRequest, items PledgeStatItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) ProBar(params ProBarRequest) (data []*ProBarData, err error)
    func (ts *TuShare) Repurchase(params RepurchaseRequest, items RepurchaseItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) STKHoldertrade(params STKHoldertradeRequest, items STKHoldertradeItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) STKManagers(params STKManagersRequest, items STKManagersItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) STKRewards(params STKRewardsRequest, items STKRewardsItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) ShareFloat(params ShareFloatRequest, items ShareFloatItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) StkHoldernumber(params StkHoldernumberRequest, items StkHoldernumberItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) StkLimit(params QuotationRequest, items StkLimitItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) StockBasic(params StockBasicRequest, items StockBasicItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) StockCompany(params StockCompanyRequest, items StockCompanyItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) Suspend(params SuspendRequest, items SuspendItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) SuspendD(params SuspendDRequest, items SuspendDItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) THSDaily(params QuotationRequest, items THSDailyItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) THSMember(params THSMemberRequest, items THSMemberItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) THXIndex(params THXIndexRequest, items THXIndexItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) Top10FloatHolders(params HoldersRequest, items Top10FloatHoldersItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) Top10Holders(params HoldersRequest, items Top10HoldersItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) TopInst(params TopRequest, items TopInstItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) TopList(params TopRequest, items TopListItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) TradeCal(params TradeCalRequest, items TradeCalItems) (tsRsp *TushareResponse, err error)
    func (ts *TuShare) Weekly(params QuotationRequest, items QuotationItems) (tsRsp *TushareResponse, err error)
```

Usage
=======
```golang
package main

import (
	"github.com/davecgh/go-spew/spew"
	tushare "github.com/yangxianbo/go-tushare"
)

func main() {
	// 具体接口请从go doc中查看
	ts := tushare.NewTuShare("")
	// 标准用法示例，由请求参数结构体和返回值参数结构体构成，返回值参数结构体通常都有一个all方法来默认获取全部参数
	rsp, err := ts.StockBasic(tushare.StockBasicRequest{}, tushare.StockBasicItems{}.All())
	if err != nil {
		panic(err)
	}
	// 标准接口一定会配备一个数据组装接口，用来解析tushare-http接口的标准回复
	resp := tushare.AssembleStockBasicData(rsp)
	spew.Dump(resp[0])
	/*
		(*tushare.StockBasicData)(0xc000528000)({
			TsCode: (string) (len=9) "000001.SZ",
			Symbol: (string) (len=6) "000001",
			Name: (string) (len=12) "平安银行",
			Area: (string) (len=6) "深圳",
			Industry: (string) (len=6) "银行",
			Fullname: (string) (len=30) "平安银行股份有限公司",
			Enname: (string) (len=22) "Ping An Bank Co., Ltd.",
			Market: (string) (len=6) "主板",
			Exchange: (string) (len=4) "SZSE",
			CurrType: (string) (len=3) "CNY",
			ListStatus: (string) (len=1) "L",
			ListDate: (string) (len=8) "19910403",
			DelistDate: (string) "",
			IsHs: (string) (len=1) "S"
		   })
	*/
	// 唯一一个特殊接口，此接口是由多个Http接口组合而成
	data, err := ts.ProBar(tushare.ProBarRequest{TsCode: "600167.SH", StartDate: "20210122", EndDate: "20212026", Adj: "qfq", Freq: "D", MA: []int{5}, BOLL: tushare.BOLLRequest{SD: 20, Width: 2}})
	if err != nil {
		panic(err)
	}
	// 数据时间顺序对比原版相反，数据结构也稍有不同
	spew.Dump(data[len(data)-1])
	/*
		(*tushare.ProBarData)(0xc000282d80)({
			QuotationData: (tushare.QuotationData) {
			 TsCode: (string) (len=9) "600167.SH",
			 TradeDate: (string) (len=8) "20210226",
			 Open: (float64) 10.1,
			 High: (float64) 10.14,
			 Low: (float64) 9.879999999999999,
			 Close: (float64) 9.91,
			 PreClose: (float64) 10.18,
			 Change: (float64) -0.27,
			 PctChg: (float64) -2.6523,
			 Vol: (float64) 105070.34,
			 Amount: (float64) 104776.305
			},
			DailyBasic: (*tushare.DailyBasicData)(0xc0001cc320)({
			 TsCode: (string) (len=9) "600167.SH",
			 TradeDate: (string) (len=8) "20210226",
			 Close: (float64) 9.91,
			 TurnoverRate: (float64) 0.4592,
			 TurnoverRateF: (float64) 1.4267,
			 VolumeRatio: (float64) 0.88,
			 Pe: (float64) 14.2474,
			 PeTTM: (float64) 12.3698,
			 Pb: (float64) 2.7512,
			 Ps: (float64) 6.6767,
			 PsTTM: (float64) 5.9302,
			 DvRatio: (float64) 2.1191,
			 DvTtm: (float64) 2.1191,
			 TotalShare: (float64) 228811.9475,
			 FloatShare: (float64) 228811.9475,
			 FreeShare: (float64) 73643.8514,
			 TotalMV: (float64) 2.2675263997e+06,
			 CircMV: (float64) 2.2675263997e+06
			}),
			MA: (map[int]float64) (len=1) {
			 (int) 5: (float64) 10.179999999999998
			},
			BOLL: (tushare.BOLLData) {
			 UP: (float64) 10.451752574361263,
			 MID: (float64) 10.101,
			 LOW: (float64) 9.750247425638738
			}
		   })
	*/
}
```