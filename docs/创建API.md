# 创建API
## 准备api接口文件
在指定目录下创建`api`文件，比如：`transaction.api`
目录结构如下所示
```sh
D:.
├─.idea
├─cmd
│  └─transaction
│      ├─api
├─docs
├─pkg
│  └─fee
└─scripts
```
> 这里是在`cmd/transaction/api`目录下创建该文件。

`transaction.api`文件的内容如下所示：
```api
syntax = "v1"

info (
title:   "transaction"
desc:    "transaction crud "
author:  "luweiqianyi"
email:   "runningriven@gmail.com"
version: "1.0.0"
)

// 通用响应
type CommonResp {
Result  int    `json:"result"`
Message string `json:"message,omitempty"` // omitempty: allow to omit
}

type (
// 交易详情
TransactionDetailReq {
StockCode string  `form:"stock_code"` // 股票代码
StockName string  `form:"stock_name"` // 股票名称
BuyPrice  float64 `form:"buy_price"` // 买入价格
Number    float64 `form:"number"` // 买入数量
SellPrice float64 `form:"sell_price"` // 卖出价格
BuyDate   string  `form:"buyDate"` // 买入日期
SellDate  string  `form:"sellDate"` // 卖出日期
}
TransactionDetailResp {
StockCode string  `json:"stock_code"` // 股票代码
StockName string  `json:"stock_name"` // 股票名称
BuyPrice  float64 `json:"buy_price"` // 买入价格
Number    float64 `json:"number"` // 买入数量
SellPrice float64 `json:"sell_price"` // 卖出价格
BuyDate   string  `json:"buyDate"` // 买入日期
SellDate  string  `json:"sellDate"` // 卖出日期
}
// 针对某次交易详情的计算结果
TransactionResultResp {
BuyCost     float64 `json:"buy_cost"` // 买入费用
SellCost    float64 `json:"sell_cost"` // 卖出费用
TotalCost   float64 `json:"total_cost"` // 买卖总费用
Rate        float64 `json:"rate"` // 涨跌比例
GainLoss    float64 `json:"gain_loss"` // 持仓收益
FinalProfit float64 `json:"final_profit"` // 最终收益
}
)

// 增加一条交易记录(计不计算由后台逻辑决定，保不保存计算结果也由后台逻辑决定)
type (
AddOneTransactionRecordReq {
TransactionDetailReq
}
AddOneTransactionRecordResp {
CommonResp
}
)

// 列出所有已成交交易的详细情况
type (
ListAllTransactionRecordsReq  {}
TransactionRecordResult {
TransactionDetailResp
TransactionResultResp
}
ListAllTransactionRecordsResp {
CommonResp
TransactionResults []TransactionRecordResult `json:"transaction_results"`
}
)

// 列出某一次成交交易的详细情况(同一支股票在当天可能有多次买入)
type (
ListOneTransactionRecordReq {
StockCode string `form:"stock_code"` // 股票代码
BuyDate   string `form:"buyDate"` // 买入日期
}
ListOneTransactionRecordResp {
CommonResp
TransactionResults []TransactionRecordResult `json:"transaction_results"`
}
)

service TransactionApi {
@handler AddOneTransactionRecordHandler
post /addOneTransactionRecord (AddOneTransactionRecordReq) returns (AddOneTransactionRecordResp)

@handler ListAllTransactionRecordsHandler
post /listAllTransactionRecords (ListAllTransactionRecordsReq) returns (ListAllTransactionRecordsResp)
}
```
## 使用goctl执行命令自动生成模板代码
```sh
goctl api go --api transaction.api --dir .
```