/* --------------------------------------------------------------
Project:    Ethereum auto-transfer (accounts to specific address(hotwallet))
Purpose:
Author:     Ho-Jung Kim (godmode2k@hotmail.com)
Date:       Since Dec 4, 2020
Filename:   rpc_server.go

Last modified:  May 22, 2022
License:

*
* Copyright (C) 2020 Ho-Jung Kim (godmode2k@hotmail.com)
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*      http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*
-----------------------------------------------------------------
Note:
-----------------------------------------------------------------
Reference:
 - https://golang.org/pkg/net/rpc/
 - https://pkg.go.dev/database/sql
 - https://pkg.go.dev/github.com/mattn/go-sqlite3

Dependencies:
 - $ go get -u github.com/go-sql-driver/mysql
 - $ go get github.com/mattn/go-sqlite3
-------------------------------------------------------------- */
package rpc_server



//! Header
// ---------------------------------------------------------------

import (
    "fmt"
    //"log"
    //"math/big"
    //"encoding/json"

    "runtime"
    //"regexp"

    // HTTP JSON-RPC Server
    //"net/http"
    //"github.com/gorilla/mux"
    //"github.com/gorilla/rpc"
    //"github.com/gorilla/rpc/json"

    // $ go get -u github.com/go-sql-driver/mysql
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    _ "github.com/mattn/go-sqlite3"

    // eth_txns_fetcher
    //"eth_block_explorer/types"
)



//! Definition
// --------------------------------------------------------------------

//var DB_SERVER_ADDRESS = "127.0.0.1:3306"
//var DB_NAME = "ethereum_block_explorer"
//var DB_LOGIN_USERNAME = "root"
//var DB_LOGIN_PASSWORD = "mysql"
var _E_TYPE__ALL = uint8(0)
var _E_TYPE__ETH = uint8(1)
var _E_TYPE__ERC20 = uint8(2)
var _E_TYPE__ERC1155 = uint8(3)
var _E_TYPE__ERC_TOKEN_CREATION = uint8(4)

var TABLE_NAME__TXID_ALL = "txid"
var TABLE_NAME__ETH_ERC20 = "txid_eth_erc20"
var TABLE_NAME__ERC1155 = "txid_erc1155"
var TABLE_NAME__BLOCKS = "blocks"
var TABLE_NAME__BALANCES_ALL = "balances"
var TABLE_NAME__BALANCES_ADDRESS_SUFFIX_0 = "balances_address_suffix_0"
var TABLE_NAME__BALANCES_ADDRESS_SUFFIX_1 = "balances_address_suffix_1"
var TABLE_NAME__BALANCES_ADDRESS_SUFFIX_2 = "balances_address_suffix_2"
var TABLE_NAME__BALANCES_ADDRESS_SUFFIX_3 = "balances_address_suffix_3"
var TABLE_NAME__BALANCES_ADDRESS_SUFFIX_4 = "balances_address_suffix_4"
var TABLE_NAME__BALANCES_ADDRESS_SUFFIX_5 = "balances_address_suffix_5"
var TABLE_NAME__BALANCES_ADDRESS_SUFFIX_6 = "balances_address_suffix_6"
var TABLE_NAME__BALANCES_ADDRESS_SUFFIX_7 = "balances_address_suffix_7"
var TABLE_NAME__BALANCES_ADDRESS_SUFFIX_8 = "balances_address_suffix_8"
var TABLE_NAME__BALANCES_ADDRESS_SUFFIX_9 = "balances_address_suffix_9"
var TABLE_NAME__BALANCES_ADDRESS_SUFFIX_a = "balances_address_suffix_a"
var TABLE_NAME__BALANCES_ADDRESS_SUFFIX_b = "balances_address_suffix_b"
var TABLE_NAME__BALANCES_ADDRESS_SUFFIX_c = "balances_address_suffix_c"
var TABLE_NAME__BALANCES_ADDRESS_SUFFIX_d = "balances_address_suffix_d"
var TABLE_NAME__BALANCES_ADDRESS_SUFFIX_e = "balances_address_suffix_e"
var TABLE_NAME__BALANCES_ADDRESS_SUFFIX_f = "balances_address_suffix_f"

var TOKEN_TYPE_ETHER = "ether"
var TOKEN_TYPE_ERC20 = "erc20"
var TOKEN_TYPE_ERC1155 = "erc1155"

var DB_MEMORY_TABLE_NAME__TXID_ALL = "memory_txid"
var DB_MEMORY_TABLE_NAME__BLOCKS_INFO = "memory_blocks"
var DB_MEMORY_TABLE_NAME__BALANCES_ALL = "memory_balances"



//! Implementation
// --------------------------------------------------------------------

type LocalDB struct {
    Db *sql.DB
    DbMemory *sql.DB
}



// ---------------------------------------------------------------



// Source: https://stackoverflow.com/questions/25927660/how-to-get-the-current-function-name
func __trace() {
    pc := make([]uintptr, 10)  // at least 1 entry needed
    runtime.Callers(2, pc)
    f := runtime.FuncForPC(pc[0])
    file, line := f.FileLine(pc[0])
    fmt.Printf("%s:%d %s\n", file, line, f.Name())
}



