package fee

import (
	"fmt"
	"os"
	"testing"
	"text/tabwriter"
)

func TestStockReplenishTransaction(t *testing.T) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n",
		"初始买入价",
		"买入数量",
		"补仓价格",
		"补仓数量",
		"补仓后预卖出价格",
		"原始价格收益率",
		"补仓价格收益率",
		"初始收益",
		"补仓收益",
		"最终收益",
		"最终收益率",
	)

	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n",
		"----",
		"----",
		"----",
		"----",
		"----",
		"----",
		"----",
		"----",
		"----",
		"----",
		"----",
	)

	tr := StockReplenishTransaction{
		Entry: StockEntry{
			Code:   "600250",
			Name:   "南京商旅",
			Market: SH,
		},
		OriginBuyPrice:     19.91,
		OriginBuyNumber:    1100,
		ReplenishPrice:     19.20,
		ReplenishNumber:    1100,
		ReplenishSellPrice: 20.19,
	}

	tr.Calculate()

	fmt.Fprintf(w, "%f\t%f\t%f\t%f\t%f\t%.2f%%\t%.2f%%\t%f\t%f\t%f\t%.2f%%\n",
		tr.OriginBuyPrice,
		tr.OriginBuyNumber,
		tr.ReplenishPrice,
		tr.ReplenishNumber,
		tr.ReplenishSellPrice,
		tr.OriginProfitRatio*100,
		tr.ReplenishProfitRatio*100,
		tr.OriginProfit,
		tr.ReplenishProfit,
		tr.FinalProfit,
		tr.FinalProfitRatio*100,
	)
	w.Flush()
}
