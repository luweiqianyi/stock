# 创建SQL模板代码
## 准备sql文件
在指定目录下创建`sql`文件，比如：`transaction.sql`
目录结构如下所示
```sh
D:.
├─.idea
├─cmd
│  ├─common
│  └─transaction
│      ├─api
│      │  ├─etc
│      │  └─internal
│      │      ├─config
│      │      ├─handler
│      │      ├─logic
│      │      ├─svc
│      │      └─types
│      ├─model
│      └─rpc
├─docs
├─pkg
│  └─fee
└─scripts
```
> 这里主要在`model`目录下放置创建的`transaction.sql`文件。

`transaction.sql`内容如下：
```sql
CREATE DATABASE IF NOT EXISTS stock;

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

CREATE TABLE IF NOT EXISTS transaction_result(
id bigint AUTO_INCREMENT,
stock_code varchar(255) NOT NULL DEFAULT '' COMMENT 'stock_code',
stock_name varchar(255) NULL COMMENT 'stock_name',
buy_price FLOAT COMMENT 'buy_price',
sell_price FLOAT COMMENT 'sell_price',
number FLOAT COMMENT 'number',
buy_date DATE COMMENT 'buy_date',
sell_date DATE COMMENT 'sell_date',

buy_cost FLOAT COMMENT 'buy_cost',
sell_cost FLOAT COMMENT 'sell_cost',
total_cost FLOAT COMMENT 'total_cost',
rate FLOAT COMMENT 'rate',
gain_loss FLOAT COMMENT 'gain_loss',
final_profit FLOAT COMMENT 'final_profit',
PRIMARY KEY (id)    # goctl的主键写法不支持在id列的定义后面写
)ENGINE = InnoDB COLLATE utf8mb4_general_ci COMMENT 'transaction_result table';
```

## 使用goctl创建sql相关的模板代码
```sh
goctl model mysql ddl --src transaction.sql --dir .
```
创建完成后生成与`mysql`持久化存储相关的模板代码。