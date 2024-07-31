package fee

// StockReplenishTransaction 补仓交易(补仓交易发生在发生亏损的前提下)
type StockReplenishTransaction struct {
	Entry              StockEntry // 股票详细信息
	OriginBuyPrice     float64    // 初始买入价格
	OriginBuyNumber    float64    // 初始买入数量
	ReplenishPrice     float64    // 补仓价格(即补仓时的买入价格)
	ReplenishNumber    float64    // 补仓数量
	ReplenishSellPrice float64    // 补仓后的进行预卖出的价格(也是初始买入时的卖出价格：我们的目的是当股票价格下跌后，在低位买入股票进行补仓后，以相同的价格卖出来弥补第一次买入造成的亏损)

	OriginProfitRatio    float64 // 原始收益率(一般来说，原始收益率为负值，因为股票正处于下跌趋势，给投资者带来相应的亏损，需要进行补仓来弥补这部分的亏损)
	ReplenishProfitRatio float64 // 补仓收益率(补仓必须收益为正，不然补仓没有任何意义)
	OriginProfit         float64 // 初始收益(一般来说这个值为负值，如果为正值就不需要进行补仓了)
	ReplenishProfit      float64 // 补仓收益
	FinalProfit          float64 // 最终收益
	FinalProfitRatio     float64 // 最终收益率
}

func (tr *StockReplenishTransaction) Calculate() {
	tr1 := Transaction{
		Entry:     tr.Entry,
		BuyPrice:  tr.OriginBuyPrice,
		SellPrice: tr.ReplenishSellPrice,
		Number:    tr.OriginBuyNumber,
	}
	// 初次投入收益计算
	tr.OriginProfit = tr1.FinalProfit()

	tr2 := Transaction{
		Entry:     tr.Entry,
		BuyPrice:  tr.ReplenishPrice,
		SellPrice: tr.ReplenishSellPrice,
		Number:    tr.ReplenishNumber,
	}
	// 补仓收益计算
	tr.ReplenishProfit = tr2.FinalProfit()

	// 最终收益计算
	tr.FinalProfit = tr.OriginProfit + tr.ReplenishProfit
	// 原始收益率计算
	tr.OriginProfitRatio = (tr.ReplenishSellPrice - tr.OriginBuyPrice) / tr.OriginBuyPrice
	// 补仓收益率计算
	tr.ReplenishProfitRatio = (tr.ReplenishSellPrice - tr.ReplenishPrice) / tr.ReplenishPrice
	// 最终收益率计算
	tr.FinalProfitRatio = tr.FinalProfit / (tr.OriginBuyPrice * tr.OriginBuyNumber)
}
