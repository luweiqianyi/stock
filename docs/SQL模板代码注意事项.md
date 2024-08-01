# SQL模板代码注意事项
## 数据类型`varchar(255) NOT NULL DEFAULT ''`与`varchar(255) NULL`的差别
```sql
CREATE TABLE IF NOT EXISTS transaction(
    id bigint AUTO_INCREMENT,
    stock_code varchar(255) NOT NULL DEFAULT '' COMMENT 'stock_code',
    stock_name varchar(255) NULL COMMENT 'stock_name',
    buy_price FLOAT COMMENT 'buy_price',
    sell_price FLOAT COMMENT 'sell_price',
    number FLOAT COMMENT 'number',
    buy_date DATE COMMENT 'buy_date',
    sell_date DATE COMMENT 'sell_date',
    PRIMARY KEY (id)    # goctl的主键写法不支持在id列的定义后面写
)ENGINE = InnoDB COLLATE utf8mb4_general_ci COMMENT 'transaction table';
```

以上面的SQL定义语句为例，用`varchar(255) NOT NULL DEFAULT ''`和`varchar(255) NULL`来定义表数据成员在生成模版代码时会有相应的差别，差别如下所示：
```go
type(
    Transaction struct {
        Id        int64           `db:"id"`
        StockCode string          `db:"stock_code"` // stock_code
        StockName sql.NullString  `db:"stock_name"` // stock_name
        BuyPrice  sql.NullFloat64 `db:"buy_price"`  // buy_price
        SellPrice sql.NullFloat64 `db:"sell_price"` // sell_price
        Number    sql.NullFloat64 `db:"number"`     // number
        BuyDate   sql.NullTime    `db:"buy_date"`   // buy_date
        SellDate  sql.NullTime    `db:"sell_date"`  // sell_date
    }
)
```
可以看到：表数据成员`stock_code`在生成模板代码时的数据类型为`string`，而表数据成员`stock_name`在生成模板代码时的数据类型为`sql.NullString`。`sql.NullString`是一种结构体数据成员，它的定义如下所示：
```go
type NullString struct {
    String string
    Valid  bool // Valid is true if String is not NULL
}
```