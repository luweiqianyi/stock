CREATE DATABASE IF NOT EXISTS stock;

CREATE TABLE IF NOT EXISTS transaction(
    id bigint AUTO_INCREMENT,
    stock_code varchar(255) NOT NULL DEFAULT '' COMMENT 'stock_code',
    stock_name varchar(255) NOT NULL DEFAULT '' COMMENT 'stock_name',
    buy_price FLOAT NOT NULL DEFAULT 0.0 COMMENT 'buy_price',
    sell_price FLOAT NOT NULL DEFAULT 0.0 COMMENT 'sell_price',
    number FLOAT NOT NULL DEFAULT 0.0 COMMENT 'number',
    buy_date DATE NOT NULL COMMENT 'buy_date',
    sell_date DATE NOT NULL COMMENT 'sell_date',
    PRIMARY KEY (id)    # goctl的主键写法不支持在id列的定义后面写
)ENGINE = InnoDB COLLATE utf8mb4_general_ci COMMENT 'transaction table';

CREATE TABLE IF NOT EXISTS transaction_result(
    id bigint AUTO_INCREMENT,
    stock_code varchar(255) NOT NULL DEFAULT '' COMMENT 'stock_code',
    stock_name varchar(255) NOT NULL DEFAULT '' COMMENT 'stock_name',
    buy_price FLOAT NOT NULL DEFAULT 0.0 COMMENT 'buy_price',
    sell_price FLOAT NOT NULL DEFAULT 0.0 COMMENT 'sell_price',
    number FLOAT NOT NULL DEFAULT 0.0 COMMENT 'number',
    buy_date DATE NOT NULL COMMENT 'buy_date',
    sell_date DATE NOT NULL COMMENT 'sell_date',

    buy_cost FLOAT NOT NULL DEFAULT 0.0 COMMENT 'buy_cost',
    sell_cost FLOAT NOT NULL DEFAULT 0.0 COMMENT 'sell_cost',
    total_cost FLOAT NOT NULL DEFAULT 0.0 COMMENT 'total_cost',
    rate FLOAT NOT NULL DEFAULT 0.0 COMMENT 'rate',
    gain_loss FLOAT NOT NULL DEFAULT 0.0 COMMENT 'gain_loss',
    final_profit FLOAT NOT NULL DEFAULT 0.0 COMMENT 'final_profit',
    PRIMARY KEY (id)    # goctl的主键写法不支持在id列的定义后面写
)ENGINE = InnoDB COLLATE utf8mb4_general_ci COMMENT 'transaction_result table';