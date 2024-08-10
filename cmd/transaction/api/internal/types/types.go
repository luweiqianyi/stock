// Code generated by goctl. DO NOT EDIT.
package types

type AddOneTransactionRecordReq struct {
	TransactionDetailReq
}

type AddOneTransactionRecordResp struct {
	CommonResp
}

type CalExpectedReturnsData struct {
	Market           string  `json:"market"`            // 证券交易所
	BuyPrice         float64 `json:"buy_price"`         // 买入价格
	SellPrice        float64 `json:"sell_price"`        // 卖出价格
	Rate             float64 `json:"rate"`              // 股价变化率
	Number           float64 `json:"number"`            // 可以买入的股数
	InvestedCaptical float64 `json:"invested_captical"` // 投入本金
	BuyCost          float64 `json:"buy_cost"`          // 买入费用
	SellCost         float64 `json:"sell_cost"`         // 卖出费用
	TotalCost        float64 `json:"total_cost"`        // 总费用
	Profit           float64 `json:"profit"`            // 最终收益
}

type CalExpectedReturnsReq struct {
	Market    string  `form:"market"`     // 证券交易所
	BuyPrice  float64 `form:"buy_price"`  // 买入价格
	SellPrice float64 `form:"sell_price"` // 卖出价格
	Balance   float64 `form:"balance"`    // 账户结余，用来计算可买入股数
}

type CalExpectedReturnsResp struct {
	CommonResp
	Data CalExpectedReturnsData `json:"data"`
}

type CalculateTransactionProfitReq struct {
	TransactionDetailReq
}

type CalculateTransactionProfitResp struct {
	CommonResp
	TransactionRecordResult TransactionRecordResult `json:"data"`
}

type CommonResp struct {
	Result  int    `json:"result"`
	Message string `json:"message,omitempty"` // omitempty: allow to omit
}

type ListAllTransactionRecordsReq struct {
}

type ListAllTransactionRecordsResp struct {
	CommonResp
	TransactionResults []TransactionRecordResult `json:"data"`
}

type ListOneTransactionRecordReq struct {
	StockCode string `form:"stock_code"` // 股票代码
	BuyDate   string `form:"buy_date"`   // 买入日期
}

type ListOneTransactionRecordResp struct {
	CommonResp
	TransactionResults []TransactionRecordResult `json:"data"`
}

type TransactionDetailReq struct {
	StockCode string  `form:"stock_code"` // 股票代码
	StockName string  `form:"stock_name"` // 股票名称
	Market    string  `form:"market"`     // 交易所类型
	BuyPrice  float64 `form:"buy_price"`  // 买入价格
	Number    float64 `form:"number"`     // 买入数量
	SellPrice float64 `form:"sell_price"` // 卖出价格
	BuyDate   string  `form:"buy_date"`   // 买入日期
	SellDate  string  `form:"sell_date"`  // 卖出日期
}

type TransactionDetailResp struct {
	StockCode string  `json:"stock_code"` // 股票代码
	StockName string  `json:"stock_name"` // 股票名称
	Market    string  `json:"market"`     // 交易所类型
	BuyPrice  float64 `json:"buy_price"`  // 买入价格
	Number    float64 `json:"number"`     // 买入数量
	SellPrice float64 `json:"sell_price"` // 卖出价格
	BuyDate   string  `json:"buy_date"`   // 买入日期
	SellDate  string  `json:"sell_date"`  // 卖出日期
}

type TransactionRecordResult struct {
	TransactionDetailResp
	TransactionResultResp
}

type TransactionResultResp struct {
	BuyCost     float64 `json:"buy_cost"`     // 买入费用
	SellCost    float64 `json:"sell_cost"`    // 卖出费用
	TotalCost   float64 `json:"total_cost"`   // 买卖总费用
	Rate        float64 `json:"rate"`         // 涨跌比例
	GainLoss    float64 `json:"gain_loss"`    // 持仓收益
	FinalProfit float64 `json:"final_profit"` // 最终收益
}
