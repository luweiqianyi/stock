package fee

import "math"

const (
	// CommissionRatio 佣金费率
	CommissionRatio = 2.354 / 10000
	// StampDutyRatio 印花税费率
	StampDutyRatio = 0.0005
	// TransferFeeRatio 过户费费率
	TransferFeeRatio = 0.00001
)

type Transaction struct {
	Entry     StockEntry
	BuyPrice  float64
	SellPrice float64
	Number    float64
}

// CommissionFee 佣金费用计算
func (tr Transaction) CommissionFee(buy bool) (fee float64) {

	if buy {
		fee = tr.BuyPrice * tr.Number * CommissionRatio
	} else {
		fee = tr.SellPrice * tr.Number * CommissionRatio
	}

	if math.Abs(fee-5.0) < 0.000001 {
		return 5.0
	}
	return fee
}

// StampDutyFee 印花税计算
func (tr Transaction) StampDutyFee(buy bool) (fee float64) {
	if buy {
		return 0.0
	}
	fee = tr.SellPrice * tr.Number * StampDutyRatio

	return fee
}

// TransferFee 过户费计算
func (tr Transaction) TransferFee(buy bool) (fee float64) {
	if tr.Entry.Market == SZ {
		return 0.0
	}
	if buy {
		fee = tr.BuyPrice * tr.Number * TransferFeeRatio
	} else {
		fee = tr.SellPrice * tr.Number * TransferFeeRatio
	}

	return fee
}

// Cost 投入资本
func (tr Transaction) Cost() float64 {
	return tr.BuyPrice * tr.Number
}

// Ratio 涨跌比例
func (tr Transaction) Ratio() float64 {
	return (tr.SellPrice - tr.BuyPrice) / tr.BuyPrice
}

// BuyFee 买入费用
func (tr Transaction) BuyFee() (fee float64) {
	fee = tr.CommissionFee(true) + tr.StampDutyFee(true) + tr.TransferFee(true)
	return fee
}

// SellFee 卖出费用
func (tr Transaction) SellFee() (fee float64) {
	fee = tr.CommissionFee(false) + tr.StampDutyFee(false) + tr.TransferFee(false)
	return fee
}

// TotalFee 买卖所花费的费用
func (tr Transaction) TotalFee() (fee float64) {
	fee = tr.BuyFee() + tr.SellFee()
	return fee
}

// ProfitAndLoss 持仓盈亏
func (tr Transaction) ProfitAndLoss() float64 {
	return (tr.SellPrice-tr.BuyPrice)*tr.Number - tr.BuyFee()
}

// FinalProfit 最终收益，也叫清仓收益
func (tr Transaction) FinalProfit() float64 {
	return (tr.SellPrice-tr.BuyPrice)*tr.Number - tr.TotalFee()
}
