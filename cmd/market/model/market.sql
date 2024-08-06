CREATE DATABASE IF NOT EXISTS stock;

CREATE TABLE IF NOT EXISTS market(
    id bigint AUTO_INCREMENT,
    market_type varchar(255) NOT NULL DEFAULT '' COMMENT 'market_type',
    PRIMARY KEY (id),    # goctl的主键写法不支持在id列的定义后面写
    UNIQUE KEY uc_market_type (market_type)
)ENGINE = InnoDB COLLATE utf8mb4_general_ci COMMENT 'transaction table';