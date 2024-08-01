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