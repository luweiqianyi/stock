package fee

import (
	"fmt"
	"os"
	"testing"
	"text/tabwriter"
)

func TestTransaction(t *testing.T) {
	//os.Stdout 作为第一个参数，表示输出将写入标准输出。
	//第二个参数 0 表示每个单元格的最小宽度。
	//第三个参数 0 表示单元格之间的填充宽度。
	//第四个参数 2 表示制表符填充的额外空间（用于对齐）。
	//第五个参数 ' ' 是填充字符，在这里使用空格作为填充字符。
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n",
		"买入单价",
		"卖出单价",
		"交易数量",
		"投入成本",
		"变化率",
		"买入费用",
		"卖出费用",
		"总费用",
		"持仓盈亏",
		"收益",
	)

	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n",
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

	tr := Transaction{
		Entry: StockEntry{
			Code:   "600250",
			Name:   "南京商旅",
			Market: SH,
		},
		BuyPrice:  7.77,
		SellPrice: 8.31,
		Number:    14400,
	}

	fmt.Fprintf(w, "%f\t%f\t%f\t%f\t%.2f%%\t%f\t%f\t%f\t%f\t%f\n",
		tr.BuyPrice,
		tr.SellPrice,
		tr.Number,
		tr.Cost(),
		tr.Ratio()*100,
		tr.BuyFee(),
		tr.SellFee(),
		tr.TotalFee(),
		tr.ProfitAndLoss(),
		tr.FinalProfit(),
	)

	// 刷新 tabwriter，将缓冲区的内容输出到 os.Stdout
	w.Flush()
}
