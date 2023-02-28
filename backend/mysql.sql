#
# DB scheme for Ethereum Block Explorer (ETH, ERC-20, ERC-1155)
#
# hjkim, 2018.12.05
# updated: 2022.05.19
#


\! echo '===== EDIT first... =====';
#exit



#MySQL utf-8
#
#[mysqld]
#collation-server = utf8_unicode_ci
#character-set-server = utf8
#skip-character-set-client-handshake
#
#
# ERROR: "; this is incompatible with sql_mode=only_full_group_by"
# Source
#sql_mode=ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION
# Fix an error
#sql_mode=STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION
#
#
#[client]
#default-character-set=utf8
#
#
#[mysql]
#default-character-set=utf8
#
#
#mysql> show variables like 'char%';

#CREATE USER 'ethereum_block_explorer'@'localhost' IDENTIFIED BY 'test';
#REVOKE ALL ON ethereum_block_explorer.* FROM 'ethereum_block_explorer'@'localhost';
#GRANT ALL ON ethereum_block_explorer.* TO 'ethereum_block_explorer'@'localhost';
#mysql> show grants for ethereum_block_explorer@localhost



CREATE DATABASE IF NOT EXISTS ethereum_block_explorer;
USE ethereum_block_explorer;



#
# ETHER, ERC-20
#
# FIXME: types
#CREATE TABLE IF NOT EXISTS txid_eth_erc20 (
#idx BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
##symbol VARCHAR(20) NOT NULL,
#from_address VARCHAR(100) NOT NULL,
#to_address VARCHAR(100) NOT NULL,
##send VARCHAR(1) NOT NULL,
#amount_wei TEXT NOT NULL,
#amount_eth TEXT NOT NULL,
#token_type VARCHAR(32) NOT NULL,
#token_symbol VARCHAR(32) NOT NULL,
#token_decimals VARCHAR(4) NOT NULL,
#token_total_supply TEXT NOT NULL,
#token_contract_address VARCHAR(100) NOT NULL,
#token_amount_wei TEXT NOT NULL,
#token_amount_eth TEXT NOT NULL,
#timestamp VARCHAR(20) NOT NULL,
#datetime TEXT NOT NULL,
#blocks TEXT NOT NULL,
#txid VARCHAR(100) NOT NULL,
#PRIMARY KEY (idx),
#UNIQUE KEY unique_col (txid)
#);



#
# ERC-1155
#
# FIXME: types
#CREATE TABLE IF NOT EXISTS txid_erc1155 (
#idx BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
##symbol VARCHAR(20) NOT NULL,
#from_address VARCHAR(100) NOT NULL,
#to_address VARCHAR(100) NOT NULL,
##send VARCHAR(1) NOT NULL,
#token_type VARCHAR(32) NOT NULL,
#token_symbol VARCHAR(32) NOT NULL,
#token_decimals VARCHAR(4) NOT NULL,
#token_contract_address VARCHAR(100) NOT NULL,
#token_amount VARCHAR(100) NOT NULL,
#token_uri_ascii TEXT NOT NULL,
#token_uri_hexadecimal TEXT NOT NULL,
#token_data_length TEXT NOT NULL,
#token_data TEXT NOT NULL,
#timestamp VARCHAR(20) NOT NULL,
#datetime TEXT NOT NULL,
#blocks TEXT NOT NULL,
#txid VARCHAR(100) NOT NULL,
#PRIMARY KEY (idx),
#UNIQUE KEY unique_col (txid)
#);



#
# All: Ether, ERC-20, ERC-1155
#
# FIXME: types
CREATE TABLE IF NOT EXISTS txid (
idx BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
#symbol VARCHAR(20) NOT NULL,
from_address VARCHAR(100) NOT NULL COMMENT "From address",
to_address VARCHAR(100) NOT NULL COMMENT "To address",
#send VARCHAR(1) NOT NULL,
ether_amount_wei TEXT NOT NULL COMMENT "Ether amount (wei)",
ether_amount_eth TEXT NOT NULL COMMENT "Ether amount (eth)",
token_type VARCHAR(32) NOT NULL COMMENT "Token type: ETH, ERC20, ERC1155",
token_symbol VARCHAR(32) NOT NULL COMMENT "Token symbol name",
token_decimals VARCHAR(4) NOT NULL COMMENT "ERC-20 Token Decimal",
token_total_supply TEXT NOT NULL COMMENT "ERC-20 Token Total Supply",
token_contract_address VARCHAR(100) NOT NULL COMMENT "ERC Contract address",
token_amount_wei TEXT NOT NULL COMMENT "ERC-20 Token amount (wei)",
token_amount_eth TEXT NOT NULL COMMENT "ERC-20 Token amoiunt (eth)",
erc1155_token_amount VARCHAR(100) NOT NULL COMMENT "ERC-1155 Token amount",
erc1155_token_uri_ascii TEXT NOT NULL COMMENT "ERC-1155 Token URI (ASCII)",
erc1155_token_uri_hexadecimal TEXT NOT NULL COMMENT "ERC-1155 Token URI (Hex)",
erc1155_token_data_length TEXT NOT NULL COMMENT "ERC-1155 Token Data length",
erc1155_token_data TEXT NOT NULL COMMENT "ERC-1155 Token Data",
timestamp VARCHAR(20) NOT NULL COMMENT "Timestamp: 1643798246",
datetime TEXT NOT NULL COMMENT "Datetime: 2022-02-02 19:37:26 +0900 KST",
blocks TEXT NOT NULL COMMENT "Block number",
txid VARCHAR(100) NOT NULL COMMENT "TXID: Transaction Id",
PRIMARY KEY (idx),
UNIQUE KEY unique_col (txid)
);



#
# Blocks
#
# FIXME: types
CREATE TABLE IF NOT EXISTS blocks (
idx BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
blocks TEXT NOT NULL COMMENT "Block number",
blocks_hash VARCHAR(100) NOT NULL COMMENT "Block Hash",
info TEXT NOT NULL COMMENT "Block information",
transactions BIGINT UNSIGNED NOT NULL COMMENT "Number of transactions",
PRIMARY KEY (idx),
UNIQUE KEY unique_col (blocks_hash)
);



#
# Balances: inserts if not exist, updates otherwise
#
# FIXME: types
# json_data JSON NOT NULL COMMENT 'JSON Array'
#
CREATE TABLE IF NOT EXISTS balances (
idx BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
last_blocks TEXT NOT NULL COMMENT "Block number",
blocks_hash VARCHAR(100) NOT NULL COMMENT "Block Hash",
json_data JSON NOT NULL COMMENT 'JSON Array',
# DELETE [{ "block": "", "txid": "", "address": "", "token_type": "", "token_symbol": "", "token_contract_address": "",
# "ether_amount_eth": "", "token_amount_eth": "", "erc1155_token_amount": "", "issend": "" }, ...]
PRIMARY KEY (idx),
UNIQUE KEY unique_col (blocks_hash)
);

CREATE TABLE IF NOT EXISTS balances_query_queue (
idx BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
blocks TEXT NOT NULL COMMENT "Block number",
txid VARCHAR(100) NOT NULL COMMENT "TXID: Transaction Id",
query_update TEXT NOT NULL COMMENT "Query string for UPDATE",
query_insert TEXT NOT NULL COMMENT "Query string for INSERT",
query_hash VARCHAR(100) NOT NULL COMMENT "Query hash: SHA1(query_update + query_insert)",
PRIMARY KEY (idx),
UNIQUE KEY unique_col (query_hash)
);

CREATE TABLE IF NOT EXISTS balances_address_suffix_0 (
idx BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
owner_address VARCHAR(100) NOT NULL COMMENT "Owner address",
txid VARCHAR(100) NOT NULL COMMENT "Block Hash",
blocks TEXT NOT NULL COMMENT "Block number",
token_type VARCHAR(32) NOT NULL COMMENT "Token type: ETH, ERC20, ERC1155",
token_symbol VARCHAR(32) NOT NULL COMMENT "Token symbol name",
token_contract_address VARCHAR(100) NOT NULL COMMENT "ERC Contract address",
amount TEXT NOT NULL COMMENT "amount",
issend VARCHAR(1) NOT NULL DEFAULT "n" COMMENT "is send: y or n",
timestamp VARCHAR(20) NOT NULL COMMENT "Timestamp: 1643798246",
query_hash VARCHAR(100) NOT NULL COMMENT "Query hash: SHA1(All, without idx)",
PRIMARY KEY (idx),
UNIQUE KEY unique_col (query_hash)
);
CREATE TABLE IF NOT EXISTS balances_address_suffix_1 (
idx BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
owner_address VARCHAR(100) NOT NULL COMMENT "Owner address",
txid VARCHAR(100) NOT NULL COMMENT "Block Hash",
blocks TEXT NOT NULL COMMENT "Block number",
token_type VARCHAR(32) NOT NULL COMMENT "Token type: ETH, ERC20, ERC1155",
token_symbol VARCHAR(32) NOT NULL COMMENT "Token symbol name",
token_contract_address VARCHAR(100) NOT NULL COMMENT "ERC Contract address",
amount TEXT NOT NULL COMMENT "amount",
issend VARCHAR(1) NOT NULL DEFAULT "n" COMMENT "is send: y or n",
timestamp VARCHAR(20) NOT NULL COMMENT "Timestamp: 1643798246",
query_hash VARCHAR(100) NOT NULL COMMENT "Query hash: SHA1(All, without idx)",
PRIMARY KEY (idx),
UNIQUE KEY unique_col (query_hash)
);
CREATE TABLE IF NOT EXISTS balances_address_suffix_2 (
idx BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
owner_address VARCHAR(100) NOT NULL COMMENT "Owner address",
txid VARCHAR(100) NOT NULL COMMENT "Block Hash",
blocks TEXT NOT NULL COMMENT "Block number",
token_type VARCHAR(32) NOT NULL COMMENT "Token type: ETH, ERC20, ERC1155",
token_symbol VARCHAR(32) NOT NULL COMMENT "Token symbol name",
token_contract_address VARCHAR(100) NOT NULL COMMENT "ERC Contract address",
amount TEXT NOT NULL COMMENT "amount",
issend VARCHAR(1) NOT NULL DEFAULT "n" COMMENT "is send: y or n",
timestamp VARCHAR(20) NOT NULL COMMENT "Timestamp: 1643798246",
query_hash VARCHAR(100) NOT NULL COMMENT "Query hash: SHA1(All, without idx)",
PRIMARY KEY (idx),
UNIQUE KEY unique_col (query_hash)
);
CREATE TABLE IF NOT EXISTS balances_address_suffix_3 (
idx BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
owner_address VARCHAR(100) NOT NULL COMMENT "Owner address",
txid VARCHAR(100) NOT NULL COMMENT "Block Hash",
blocks TEXT NOT NULL COMMENT "Block number",
token_type VARCHAR(32) NOT NULL COMMENT "Token type: ETH, ERC20, ERC1155",
token_symbol VARCHAR(32) NOT NULL COMMENT "Token symbol name",
token_contract_address VARCHAR(100) NOT NULL COMMENT "ERC Contract address",
amount TEXT NOT NULL COMMENT "amount",
issend VARCHAR(1) NOT NULL DEFAULT "n" COMMENT "is send: y or n",
timestamp VARCHAR(20) NOT NULL COMMENT "Timestamp: 1643798246",
query_hash VARCHAR(100) NOT NULL COMMENT "Query hash: SHA1(All, without idx)",
PRIMARY KEY (idx),
UNIQUE KEY unique_col (query_hash)
);
CREATE TABLE IF NOT EXISTS balances_address_suffix_4 (
idx BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
owner_address VARCHAR(100) NOT NULL COMMENT "Owner address",
txid VARCHAR(100) NOT NULL COMMENT "Block Hash",
blocks TEXT NOT NULL COMMENT "Block number",
token_type VARCHAR(32) NOT NULL COMMENT "Token type: ETH, ERC20, ERC1155",
token_symbol VARCHAR(32) NOT NULL COMMENT "Token symbol name",
token_contract_address VARCHAR(100) NOT NULL COMMENT "ERC Contract address",
amount TEXT NOT NULL COMMENT "amount",
issend VARCHAR(1) NOT NULL DEFAULT "n" COMMENT "is send: y or n",
timestamp VARCHAR(20) NOT NULL COMMENT "Timestamp: 1643798246",
query_hash VARCHAR(100) NOT NULL COMMENT "Query hash: SHA1(All, without idx)",
PRIMARY KEY (idx),
UNIQUE KEY unique_col (query_hash)
);
CREATE TABLE IF NOT EXISTS balances_address_suffix_5 (
idx BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
owner_address VARCHAR(100) NOT NULL COMMENT "Owner address",
txid VARCHAR(100) NOT NULL COMMENT "Block Hash",
blocks TEXT NOT NULL COMMENT "Block number",
token_type VARCHAR(32) NOT NULL COMMENT "Token type: ETH, ERC20, ERC1155",
token_symbol VARCHAR(32) NOT NULL COMMENT "Token symbol name",
token_contract_address VARCHAR(100) NOT NULL COMMENT "ERC Contract address",
amount TEXT NOT NULL COMMENT "amount",
issend VARCHAR(1) NOT NULL DEFAULT "n" COMMENT "is send: y or n",
timestamp VARCHAR(20) NOT NULL COMMENT "Timestamp: 1643798246",
query_hash VARCHAR(100) NOT NULL COMMENT "Query hash: SHA1(All, without idx)",
PRIMARY KEY (idx),
UNIQUE KEY unique_col (query_hash)
);
CREATE TABLE IF NOT EXISTS balances_address_suffix_6 (
idx BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
owner_address VARCHAR(100) NOT NULL COMMENT "Owner address",
txid VARCHAR(100) NOT NULL COMMENT "Block Hash",
blocks TEXT NOT NULL COMMENT "Block number",
token_type VARCHAR(32) NOT NULL COMMENT "Token type: ETH, ERC20, ERC1155",
token_symbol VARCHAR(32) NOT NULL COMMENT "Token symbol name",
token_contract_address VARCHAR(100) NOT NULL COMMENT "ERC Contract address",
amount TEXT NOT NULL COMMENT "amount",
issend VARCHAR(1) NOT NULL DEFAULT "n" COMMENT "is send: y or n",
timestamp VARCHAR(20) NOT NULL COMMENT "Timestamp: 1643798246",
query_hash VARCHAR(100) NOT NULL COMMENT "Query hash: SHA1(All, without idx)",
PRIMARY KEY (idx),
UNIQUE KEY unique_col (query_hash)
);
CREATE TABLE IF NOT EXISTS balances_address_suffix_7 (
idx BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
owner_address VARCHAR(100) NOT NULL COMMENT "Owner address",
txid VARCHAR(100) NOT NULL COMMENT "Block Hash",
blocks TEXT NOT NULL COMMENT "Block number",
token_type VARCHAR(32) NOT NULL COMMENT "Token type: ETH, ERC20, ERC1155",
token_symbol VARCHAR(32) NOT NULL COMMENT "Token symbol name",
token_contract_address VARCHAR(100) NOT NULL COMMENT "ERC Contract address",
amount TEXT NOT NULL COMMENT "amount",
issend VARCHAR(1) NOT NULL DEFAULT "n" COMMENT "is send: y or n",
timestamp VARCHAR(20) NOT NULL COMMENT "Timestamp: 1643798246",
query_hash VARCHAR(100) NOT NULL COMMENT "Query hash: SHA1(All, without idx)",
PRIMARY KEY (idx),
UNIQUE KEY unique_col (query_hash)
);
CREATE TABLE IF NOT EXISTS balances_address_suffix_8 (
idx BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
owner_address VARCHAR(100) NOT NULL COMMENT "Owner address",
txid VARCHAR(100) NOT NULL COMMENT "Block Hash",
blocks TEXT NOT NULL COMMENT "Block number",
token_type VARCHAR(32) NOT NULL COMMENT "Token type: ETH, ERC20, ERC1155",
token_symbol VARCHAR(32) NOT NULL COMMENT "Token symbol name",
token_contract_address VARCHAR(100) NOT NULL COMMENT "ERC Contract address",
amount TEXT NOT NULL COMMENT "amount",
issend VARCHAR(1) NOT NULL DEFAULT "n" COMMENT "is send: y or n",
timestamp VARCHAR(20) NOT NULL COMMENT "Timestamp: 1643798246",
query_hash VARCHAR(100) NOT NULL COMMENT "Query hash: SHA1(All, without idx)",
PRIMARY KEY (idx),
UNIQUE KEY unique_col (query_hash)
);
CREATE TABLE IF NOT EXISTS balances_address_suffix_9 (
idx BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
owner_address VARCHAR(100) NOT NULL COMMENT "Owner address",
txid VARCHAR(100) NOT NULL COMMENT "Block Hash",
blocks TEXT NOT NULL COMMENT "Block number",
token_type VARCHAR(32) NOT NULL COMMENT "Token type: ETH, ERC20, ERC1155",
token_symbol VARCHAR(32) NOT NULL COMMENT "Token symbol name",
token_contract_address VARCHAR(100) NOT NULL COMMENT "ERC Contract address",
amount TEXT NOT NULL COMMENT "amount",
issend VARCHAR(1) NOT NULL DEFAULT "n" COMMENT "is send: y or n",
timestamp VARCHAR(20) NOT NULL COMMENT "Timestamp: 1643798246",
query_hash VARCHAR(100) NOT NULL COMMENT "Query hash: SHA1(All, without idx)",
PRIMARY KEY (idx),
UNIQUE KEY unique_col (query_hash)
);
CREATE TABLE IF NOT EXISTS balances_address_suffix_a (
idx BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
owner_address VARCHAR(100) NOT NULL COMMENT "Owner address",
txid VARCHAR(100) NOT NULL COMMENT "Block Hash",
blocks TEXT NOT NULL COMMENT "Block number",
token_type VARCHAR(32) NOT NULL COMMENT "Token type: ETH, ERC20, ERC1155",
token_symbol VARCHAR(32) NOT NULL COMMENT "Token symbol name",
token_contract_address VARCHAR(100) NOT NULL COMMENT "ERC Contract address",
amount TEXT NOT NULL COMMENT "amount",
issend VARCHAR(1) NOT NULL DEFAULT "n" COMMENT "is send: y or n",
timestamp VARCHAR(20) NOT NULL COMMENT "Timestamp: 1643798246",
query_hash VARCHAR(100) NOT NULL COMMENT "Query hash: SHA1(All, without idx)",
PRIMARY KEY (idx),
UNIQUE KEY unique_col (query_hash)
);
CREATE TABLE IF NOT EXISTS balances_address_suffix_b (
idx BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
owner_address VARCHAR(100) NOT NULL COMMENT "Owner address",
txid VARCHAR(100) NOT NULL COMMENT "Block Hash",
blocks TEXT NOT NULL COMMENT "Block number",
token_type VARCHAR(32) NOT NULL COMMENT "Token type: ETH, ERC20, ERC1155",
token_symbol VARCHAR(32) NOT NULL COMMENT "Token symbol name",
token_contract_address VARCHAR(100) NOT NULL COMMENT "ERC Contract address",
amount TEXT NOT NULL COMMENT "amount",
issend VARCHAR(1) NOT NULL DEFAULT "n" COMMENT "is send: y or n",
timestamp VARCHAR(20) NOT NULL COMMENT "Timestamp: 1643798246",
query_hash VARCHAR(100) NOT NULL COMMENT "Query hash: SHA1(All, without idx)",
PRIMARY KEY (idx),
UNIQUE KEY unique_col (query_hash)
);
CREATE TABLE IF NOT EXISTS balances_address_suffix_c (
idx BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
owner_address VARCHAR(100) NOT NULL COMMENT "Owner address",
txid VARCHAR(100) NOT NULL COMMENT "Block Hash",
blocks TEXT NOT NULL COMMENT "Block number",
token_type VARCHAR(32) NOT NULL COMMENT "Token type: ETH, ERC20, ERC1155",
token_symbol VARCHAR(32) NOT NULL COMMENT "Token symbol name",
token_contract_address VARCHAR(100) NOT NULL COMMENT "ERC Contract address",
amount TEXT NOT NULL COMMENT "amount",
issend VARCHAR(1) NOT NULL DEFAULT "n" COMMENT "is send: y or n",
timestamp VARCHAR(20) NOT NULL COMMENT "Timestamp: 1643798246",
query_hash VARCHAR(100) NOT NULL COMMENT "Query hash: SHA1(All, without idx)",
PRIMARY KEY (idx),
UNIQUE KEY unique_col (query_hash)
);
CREATE TABLE IF NOT EXISTS balances_address_suffix_d (
idx BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
owner_address VARCHAR(100) NOT NULL COMMENT "Owner address",
txid VARCHAR(100) NOT NULL COMMENT "Block Hash",
blocks TEXT NOT NULL COMMENT "Block number",
token_type VARCHAR(32) NOT NULL COMMENT "Token type: ETH, ERC20, ERC1155",
token_symbol VARCHAR(32) NOT NULL COMMENT "Token symbol name",
token_contract_address VARCHAR(100) NOT NULL COMMENT "ERC Contract address",
amount TEXT NOT NULL COMMENT "amount",
issend VARCHAR(1) NOT NULL DEFAULT "n" COMMENT "is send: y or n",
timestamp VARCHAR(20) NOT NULL COMMENT "Timestamp: 1643798246",
query_hash VARCHAR(100) NOT NULL COMMENT "Query hash: SHA1(All, without idx)",
PRIMARY KEY (idx),
UNIQUE KEY unique_col (query_hash)
);
CREATE TABLE IF NOT EXISTS balances_address_suffix_e (
idx BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
owner_address VARCHAR(100) NOT NULL COMMENT "Owner address",
txid VARCHAR(100) NOT NULL COMMENT "Block Hash",
blocks TEXT NOT NULL COMMENT "Block number",
token_type VARCHAR(32) NOT NULL COMMENT "Token type: ETH, ERC20, ERC1155",
token_symbol VARCHAR(32) NOT NULL COMMENT "Token symbol name",
token_contract_address VARCHAR(100) NOT NULL COMMENT "ERC Contract address",
amount TEXT NOT NULL COMMENT "amount",
issend VARCHAR(1) NOT NULL DEFAULT "n" COMMENT "is send: y or n",
timestamp VARCHAR(20) NOT NULL COMMENT "Timestamp: 1643798246",
query_hash VARCHAR(100) NOT NULL COMMENT "Query hash: SHA1(All, without idx)",
PRIMARY KEY (idx),
UNIQUE KEY unique_col (query_hash)
);
CREATE TABLE IF NOT EXISTS balances_address_suffix_f (
idx BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
owner_address VARCHAR(100) NOT NULL COMMENT "Owner address",
txid VARCHAR(100) NOT NULL COMMENT "Block Hash",
blocks TEXT NOT NULL COMMENT "Block number",
token_type VARCHAR(32) NOT NULL COMMENT "Token type: ETH, ERC20, ERC1155",
token_symbol VARCHAR(32) NOT NULL COMMENT "Token symbol name",
token_contract_address VARCHAR(100) NOT NULL COMMENT "ERC Contract address",
amount TEXT NOT NULL COMMENT "amount",
issend VARCHAR(1) NOT NULL DEFAULT "n" COMMENT "is send: y or n",
timestamp VARCHAR(20) NOT NULL COMMENT "Timestamp: 1643798246",
query_hash VARCHAR(100) NOT NULL COMMENT "Query hash: SHA1(All, without idx)",
PRIMARY KEY (idx),
UNIQUE KEY unique_col (query_hash)
);




# Ethereum (ETHER, ERC-20)
#INSERT INTO txid VALUES (0, "from_address", "to_address", "amount_wei", "amount_eth", "token_type", "token_symbol", "token_decimals", "token_contract_address", "token_amount_wei", "token_amount_eth", "timestamp", "datetime", "1000", "txid");


# Ethereum (ERC-1155)
#INSERT INTO txid_erc1155 VALUES (0, "from_address", "to_address", "token_type", "token_symbol", "token_decimals", "token_contract_address", "token_amount", "token_uri_ascii", "token_uri_hex", "token_data_length", "token_data", "timestamp", "datetime", "1000", "txid");


# Ethereum (ETHER, ERC-20, ERC-1155)
#INSERT INTO txid_erc1155 VALUES (0, "from_address", "to_address", "ether_amount_wei", "ether_amount_eth", "token_type", "token_symbol", "token_decimals", "token_total_supply", "token_contract_address", "token_amount_wei", "token_amount_eth", "erc1155_token_amount", "erc1155_token_uri_ascii", "erc1155_token_uri_hex", "erc1155_token_data_length", "erc1155_token_data", "timestamp", "datetime", "1000", "txid");





#
# procedure test
#
#delimiter //
#create procedure test()
#begin
#select * from txid;
#end
#//
#delimiter ;
#



