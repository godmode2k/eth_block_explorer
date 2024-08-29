/* --------------------------------------------------------------
Project:    Ethereum auto-transfer (accounts to specific address(hotwallet))
Purpose:
Author:     Ho-Jung Kim (godmode2k@hotmail.com)
Date:       Since Dec 4, 2020
Filename:   rpc_server_db.go

Last modified:  February 23, 2023
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

    //"runtime"
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
    "eth_block_explorer/types"
)



//! Definition
// --------------------------------------------------------------------
var DEFAULT_LIMIT_PER_PAGE = uint(10000)



//! Implementation
// --------------------------------------------------------------------



func (t *LocalDB) _query_raw(query_str string) *sql.Rows {
    fmt.Println( "query = ", query_str )

    result, err := t.Db.Query( query_str )

    if err != nil {
        panic( err.Error() )
    }

    //fmt.Println( "result = ", result )

    return result
}

func (t *LocalDB) _query(_type uint8, query_str string) []types.Fetch_transactions_st {
    fmt.Println( "query = ", query_str )

    var result []types.Fetch_transactions_st

    //rows, err := t.Db.Query("select id, name from foo limit 1")
    //rows, err := t.Db.Query( query )
    //if err != nil {
    //    log.Fatal(err)
    //}


    //rows *sql.Rows
    rows, err := t.Db.Query( query_str )

    if err != nil {
        //log.Fatal( err.Error() )
        panic( err.Error() )
    }
    defer rows.Close()


    /*
    if ( _type == _E_TYPE__ETH ) {
        for rows.Next() {
            var Index sql.NullInt64
            //var Symbol string
            var From_address string
            var To_address string
            var Is_send string
            var Amount_wei string
            var Amount_eth string
            var Token_type string
            var Token_symbol string
            var Token_decimals string
            var Token_total_supply string
            var Token_contract_address string
            var Token_amount_wei string
            var Token_amount_eth string
            // for SQL: txid_erc1155 table
            //var Token_amount string
            //var Token_uri_ascii string
            //var Token_uri_hexadecimal string
            //var Token_data_length string
            //var Token_data string
            var Timestamp string
            var Datetime string
            var Block_number string
            var Txid string

            err := rows.Scan(
                &Index,
                //&Symbol,
                &From_address,
                &To_address,
                &Is_send,
                &Amount_wei,
                &Amount_eth,
                &Token_type,
                &Token_symbol,
                &Token_decimals,
                &Token_total_supply,
                &Token_contract_address,
                &Token_amount_wei,
                &Token_amount_eth,
                // for SQL: txid_erc1155 table
                //&Token_amount,
                //&Token_uri_ascii,
                //&Token_uri_hexadecimal,
                //&Token_data_length,
                //&Token_data,
                &Timestamp,
                &Datetime,
                &Block_number,
                &Txid,
            )

            if err != nil {
                //log.Fatal(err)
                panic( err )
            }

            _data := types.Fetch_transactions_st {
                //Symbol: Symbol,
                From_address: From_address,
                To_address: To_address,
                Is_send: Is_send,
                Amount_wei: Amount_wei,
                Amount_eth: Amount_eth,
                Token_type: Token_type,
                Token_symbol: Token_symbol,
                Token_decimals: Token_decimals,
                Token_total_supply: Token_total_supply,
                Token_contract_address: Token_contract_address,
                Token_amount_wei: Token_amount_wei,
                Token_amount_eth: Token_amount_eth,
                // for SQL: txid_erc1155 table
                //Token_amount: Token_amount,
                //Token_uri_ascii: Token_uri_ascii,
                //Token_uri_hexadecimal: Token_uri_hexadecimal,
                //Token_data_length: Token_data_length,
                //Token_data: Token_data,
                Timestamp: Timestamp,
                Datetime: Datetime,
                Block_number: Block_number,
                Txid: Txid,
            }

            result = append( result, _data )
        } // for ()
    } else if ( _type == _E_TYPE__ERC20 ) {
        for rows.Next() {
            var Index sql.NullInt64
            //var Symbol string
            var From_address string
            var To_address string
            var Is_send string
            var Amount_wei string
            var Amount_eth string
            var Token_type string
            var Token_symbol string
            var Token_decimals string
            var Token_total_supply string
            var Token_contract_address string
            var Token_amount_wei string
            var Token_amount_eth string
            // for SQL: txid_erc1155 table
            //var Token_amount string
            //var Token_uri_ascii string
            //var Token_uri_hexadecimal string
            //var Token_data_length string
            //var Token_data string
            var Timestamp string
            var Datetime string
            var Block_number string
            var Txid string

            err := rows.Scan(
                &Index,
                //&Symbol,
                &From_address,
                &To_address,
                &Is_send,
                &Amount_wei,
                &Amount_eth,
                &Token_type,
                &Token_symbol,
                &Token_decimals,
                &Token_total_supply,
                &Token_contract_address,
                &Token_amount_wei,
                &Token_amount_eth,
                // for SQL: txid_erc1155 table
                //&Token_amount,
                //&Token_uri_ascii,
                //&Token_uri_hexadecimal,
                //&Token_data_length,
                //&Token_data,
                &Timestamp,
                &Datetime,
                &Block_number,
                &Txid,
            )

            if err != nil {
                //log.Fatal(err)
                panic( err )
            }

            _data := types.Fetch_transactions_st {
                //Symbol: Symbol,
                From_address: From_address,
                To_address: To_address,
                Is_send: Is_send,
                Amount_wei: Amount_wei,
                Amount_eth: Amount_eth,
                Token_type: Token_type,
                Token_symbol: Token_symbol,
                Token_decimals: Token_decimals,
                Token_total_supply: Token_total_supply,
                Token_contract_address: Token_contract_address,
                Token_amount_wei: Token_amount_wei,
                Token_amount_eth: Token_amount_eth,
                // for SQL: txid table
                //Token_amount: Token_amount,
                //Token_uri_ascii: Token_uri_ascii,
                //Token_uri_hexadecimal: Token_uri_hexadecimal,
                //Token_data_length: Token_data_length,
                //Token_data: Token_data,
                Timestamp: Timestamp,
                Datetime: Datetime,
                Block_number: Block_number,
                Txid: Txid,
            }

            result = append( result, _data )
        } // for ()
    } else if ( _type == _E_TYPE__ERC1155 ) {
        for rows.Next() {
            var Index sql.NullInt64
            //var Symbol string
            var From_address string
            var To_address string
            var Is_send string
            // for SQL: txid table
            //var Amount_wei string
            //var Amount_eth string
            var Token_type string
            var Token_symbol string
            var Token_decimals string
            // for SQL: txid table
            //var Token_total_supply string
            var Token_contract_address string
            // for SQL: txid table
            //var Token_amount_wei string
            //var Token_amount_eth string
            var Token_amount string
            var Token_uri_ascii string
            var Token_uri_hexadecimal string
            var Token_data_length string
            var Token_data string
            var Timestamp string
            var Datetime string
            var Block_number string
            var Txid string

            err := rows.Scan(
                &Index,
                //&Symbol,
                &From_address,
                &To_address,
                &Is_send,
                // for SQL: txid table
                //&Amount_wei,
                //&Amount_eth,
                &Token_type,
                &Token_symbol,
                &Token_decimals,
                // for SQL: txid table
                //&Token_total_supply,
                &Token_contract_address,
                // for SQL: txid table
                //&Token_amount_wei,
                //&Token_amount_eth,
                &Token_amount,
                &Token_uri_ascii,
                &Token_uri_hexadecimal,
                &Token_data_length,
                &Token_data,
                &Timestamp,
                &Datetime,
                &Block_number,
                &Txid,
            )

            if err != nil {
                //log.Fatal(err)
                panic( err )
            }

            _data := types.Fetch_transactions_st {
                //Symbol: Symbol,
                From_address: From_address,
                To_address: To_address,
                Is_send: Is_send,
                // for SQL: txid_erc1155 table
                //Amount_wei: Amount_wei,
                //Amount_eth: Amount_eth,
                Token_type: Token_type,
                Token_symbol: Token_symbol,
                // for SQL: txid table
                //Token_decimals: Token_decimals,
                //Token_total_supply: Token_total_supply,
                Token_contract_address: Token_contract_address,
                // for SQL: txid table
                //Token_amount_wei: Token_amount_wei,
                //Token_amount_eth: Token_amount_eth,
                Token_amount: Token_amount,
                Token_uri_ascii: Token_uri_ascii,
                Token_uri_hexadecimal: Token_uri_hexadecimal,
                Token_data_length: Token_data_length,
                Token_data: Token_data,
                Timestamp: Timestamp,
                Datetime: Datetime,
                Block_number: Block_number,
                Txid: Txid,
            }

            result = append( result, _data )
        } // for ()
    } else {
        //result = types.Fetch_transactions_st {}

        for rows.Next() {
            var Index sql.NullInt64
            //var Symbol string
            var From_address string
            var To_address string
            var Is_send string
            var Amount_wei string
            var Amount_eth string
            var Token_type string
            var Token_symbol string
            var Token_decimals string
            var Token_total_supply string
            var Token_contract_address string
            var Token_amount_wei string
            var Token_amount_eth string
            var Token_amount string
            var Token_uri_ascii string
            var Token_uri_hexadecimal string
            var Token_data_length string
            var Token_data string
            var Timestamp string
            var Datetime string
            var Block_number string
            var Txid string

            err := rows.Scan(
                &Index,
                //&Symbol,
                &From_address,
                &To_address,
                &Is_send,
                &Amount_wei,
                &Amount_eth,
                &Token_type,
                &Token_symbol,
                &Token_decimals,
                &Token_total_supply,
                &Token_contract_address,
                &Token_amount_wei,
                &Token_amount_eth,
                &Token_amount,
                &Token_uri_ascii,
                &Token_uri_hexadecimal,
                &Token_data_length,
                &Token_data,
                &Timestamp,
                &Datetime,
                &Block_number,
                &Txid,
            )

            if err != nil {
                //log.Fatal(err)
                panic( err )
            }

            _data := types.Fetch_transactions_st {
                //Symbol: Symbol,
                From_address: From_address,
                To_address: To_address,
                Is_send: Is_send,
                Amount_wei: Amount_wei,
                Amount_eth: Amount_eth,
                Token_type: Token_type,
                Token_symbol: Token_symbol,
                Token_decimals: Token_decimals,
                Token_total_supply: Token_total_supply,
                Token_contract_address: Token_contract_address,
                Token_amount_wei: Token_amount_wei,
                Token_amount_eth: Token_amount_eth,
                Token_amount: Token_amount,
                Token_uri_ascii: Token_uri_ascii,
                Token_uri_hexadecimal: Token_uri_hexadecimal,
                Token_data_length: Token_data_length,
                Token_data: Token_data,
                Timestamp: Timestamp,
                Datetime: Datetime,
                Block_number: Block_number,
                Txid: Txid,
            }

            result = append( result, _data )
        } // for ()
    }
    */


    if ( _type == _E_TYPE__ALL ) {
        //result = types.Fetch_transactions_st {}

        for rows.Next() {
            var Index sql.NullInt64
            //var Symbol string
            var From_address string
            var To_address string
            //var Is_send string
            var Amount_wei string
            var Amount_eth string
            var Token_type string
            var Token_symbol string
            var Token_decimals string
            var Token_total_supply string
            var Token_contract_address string
            var Token_amount_wei string
            var Token_amount_eth string
            var Token_amount string
            var Token_id_ascii string
            var Token_id_hexadecimal string
            var Token_uri_ascii string
            var Token_uri_hexadecimal string
            var Token_data_length string
            var Token_data string
            var Timestamp string
            var Datetime string
            var Block_number string
            var Txid string

            err := rows.Scan(
                &Index,
                //&Symbol,
                &From_address,
                &To_address,
                //&Is_send,
                &Amount_wei,
                &Amount_eth,
                &Token_type,
                &Token_symbol,
                &Token_decimals,
                &Token_total_supply,
                &Token_contract_address,
                &Token_amount_wei,
                &Token_amount_eth,
                &Token_amount,
                &Token_id_ascii,
                &Token_id_hexadecimal,
                &Token_uri_ascii,
                &Token_uri_hexadecimal,
                &Token_data_length,
                &Token_data,
                &Timestamp,
                &Datetime,
                &Block_number,
                &Txid,
            )

            if err != nil {
                //log.Fatal(err)
                panic( err )
            }

            _data := types.Fetch_transactions_st {
                //Symbol: Symbol,
                From_address: From_address,
                To_address: To_address,
                //Is_send: Is_send,
                Amount_wei: Amount_wei,
                Amount_eth: Amount_eth,
                Token_type: Token_type,
                Token_symbol: Token_symbol,
                Token_decimals: Token_decimals,
                Token_total_supply: Token_total_supply,
                Token_contract_address: Token_contract_address,
                Token_amount_wei: Token_amount_wei,
                Token_amount_eth: Token_amount_eth,
                Token_amount: Token_amount,
                Token_id_ascii: Token_id_ascii,
                Token_id_hexadecimal: Token_id_hexadecimal,
                Token_uri_ascii: Token_uri_ascii,
                Token_uri_hexadecimal: Token_uri_hexadecimal,
                Token_data_length: Token_data_length,
                Token_data: Token_data,
                Timestamp: Timestamp,
                Datetime: Datetime,
                Block_number: Block_number,
                Txid: Txid,
            }

            result = append( result, _data )
        } // for ()
    }

    return result
}

func (t *LocalDB) _query_blocks(query_str string) []types.Fetch_blocks_info_st {
    fmt.Println( "query = ", query_str )

    var result []types.Fetch_blocks_info_st

    //rows *sql.Rows
    rows, err := t.Db.Query( query_str )

    if err != nil {
        //log.Fatal( err.Error() )
        panic( err.Error() )
    }
    defer rows.Close()


    for rows.Next() {
        var Index sql.NullInt64
        var Blocks string
        var Blocks_hash string
        var Info string
        var Transactions uint

        err := rows.Scan(
            &Index,
            &Blocks,
            &Blocks_hash,
            &Info,
            &Transactions,
        )

        if err != nil {
            //log.Fatal(err)
            panic( err )
        }

        _data := types.Fetch_blocks_info_st {
            Block_number: Blocks,
            Block_hash: Blocks_hash,
            Info: Info,
            Transactions: fmt.Sprintf("%d", Transactions),
        }

        result = append( result, _data )
    } // for ()

    return result
}

//! TODO
func (t *LocalDB) _query_balances(query_str string) []types.Fetch_balances_st {
    fmt.Println( "query = ", query_str )

    var result []types.Fetch_balances_st

/*
    //rows *sql.Rows
    rows, err := t.Db.Query( query_str )

    if err != nil {
        //log.Fatal( err.Error() )
        panic( err.Error() )
    }
    defer rows.Close()


    for rows.Next() {
        var Index sql.NullInt64
        var Last_blocks string
        var Blocks_hash string
        var Json_data string

        err := rows.Scan(
            &Index,
            &Last_blocks,
            &Blocks_hash,
            &Json_data,
        )

        if err != nil {
            //log.Fatal(err)
            panic( err )
        }

        _data := types.Fetch_balances_st {
            Last_blocks: Last_blocks,
            Blocks_hash: Blocks_hash,
            Json_data: Json_data,
        }

        result = append( result, _data )
    } // for ()
*/
    return result
}

func (t *LocalDB) _query_balances_by_address(query_str string) []types.Fetch_balances_by_address_st {
    fmt.Println( "query = ", query_str )

    var result []types.Fetch_balances_by_address_st

    //rows *sql.Rows
    rows, err := t.Db.Query( query_str )

    if err != nil {
        //log.Fatal( err.Error() )
        panic( err.Error() )
    }
    defer rows.Close()


    for rows.Next() {
        var Index sql.NullInt64
        var Owner_address string
        var Txid string
        var Blocks string
        var Token_type string
        var Token_symbol string
        var Token_contract_address string
        var Token_amount string
        var Is_send string
        var Timestamp string
        var Query_hash string

        err := rows.Scan(
            &Index,
            &Owner_address,
            &Txid,
            &Blocks,
            &Token_type,
            &Token_symbol,
            &Token_contract_address,
            &Token_amount,
            &Is_send,
            &Timestamp,
            &Query_hash,
        )

        if err != nil {
            //log.Fatal(err)
            panic( err )
        }

        _data := types.Fetch_balances_by_address_st {
            Owner_address: Owner_address,
            Txid: Txid,
            Block_number: Blocks,
            Token_type: Token_type,
            Token_symbol: Token_symbol,
            Amount: Token_amount,
        }

        result = append( result, _data )
    } // for ()

    return result
}



/*
func (t *LocalDB) _query_memory(query_str string) []types.Fetch_transactions_st {
    fmt.Println( "query = ", query_str )

    var result []types.Fetch_transactions_st

    //rows, err := t.Db.Query("select id, name from foo limit 1")
    //rows, err := t.Db.Query( query )
    //if err != nil {
    //    log.Fatal(err)
    //}

    if t.DbMemory == nil {
        fmt.Println( "SQLite: connection lost" )
        return nil
    }

    if len(query_str) <= 0 {
        fmt.Println( "SQLite: query is empty" )
        return nil
    }


    //rows *sql.Rows
    //rows, err := t.DbMemory.Query("select id, name from foo limit 1")
    rows, err := t.DbMemory.Query( query_str )
    if err != nil {
        //log.Fatal(err)
        panic( err.Error() )
    }
    defer rows.Close()



//    for rows.Next() {
//        var id int
//        var name string
//        err = rows.Scan(&id, &name)
//
//        if err != nil {
//            log.Fatal(err)
//        }
//        fmt.Println(id, name)
//
//        result = fmt.Sprintf( "id = %d, name = %s", id, name )
//    }
//
//    err = rows.Err()
//    if err != nil {
//        log.Fatal(err)
//    }


    return result
}
*/


/*
func (t *LocalDB) Test_localdb(query string, response *string) error {
    if t.Db == nil {
        fmt.Println( "LocalDB: connection lost" )
        return nil
    }

    if len(query) <= 0 {
        fmt.Println( "LocalDB: query is empty" )
        return nil
    }

    var result string

    //rows, err := t.Db.Query("select id, name from foo limit 1")
    rows, err := t.Db.Query( query )
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()
    for rows.Next() {
        var id int
        var name string
        err = rows.Scan(&id, &name)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(id, name)

        result = fmt.Sprintf( "id = %d, name = %s", id, name )
    }
    err = rows.Err()
    if err != nil {
        log.Fatal(err)
    }


    *response = result

    return nil
}
*/



// ---------------------------------------------------------------



func (t *LocalDB) _Db_select_txns_all_mixed(LIMIT_PER_PAGE uint, request_page uint, response *[]types.Fetch_transactions_st) error {
    //if ( request_page < 1 ) {
    //    fmt.Println( "request page < 1" )
    //    return nil
    //}

    // limit: transactions per page
    //LIMIT_PER_PAGE := uint(30)
    //REQ_ROWS := (LIMIT_PER_PAGE * request_page)
    OFFSET := (LIMIT_PER_PAGE * request_page)

    //query_str := fmt.Sprintf(
    //    "SELECT t.* FROM ((SELECT * FROM %s LIMIT %d,%d) UNION (SELECT * FROM %s LIMIT %d,%d))t ORDER BY timestamp DESC",
    //    TABLE_NAME__ETH_ERC20, OFFSET, LIMIT_PER_PAGE,
    //    TABLE_NAME__ERC1155, OFFSET, LIMIT_PER_PAGE,
    //)

    query_str := fmt.Sprintf(
        "SELECT * FROM %s ORDER BY timestamp DESC LIMIT %d,%d",
        TABLE_NAME__TXID_ALL, OFFSET, LIMIT_PER_PAGE,
    )


    *response = t._query( _E_TYPE__ALL, query_str )

    return nil
}

func (t *LocalDB) Db_select_txns_all_mixed(request_page uint, response *[]types.Fetch_transactions_st) error {
    LIMIT_PER_PAGE := uint(30)
    t._Db_select_txns_all_mixed( LIMIT_PER_PAGE, request_page, response )
    return nil
}

func (t *LocalDB) _Db_select_txns_all_erc(LIMIT_PER_PAGE uint, request_page uint, response *[]types.Fetch_transactions_st) error {
    // limit: 30 transactions per page
    //LIMIT_PER_PAGE := uint(30)
    //REQ_ROWS := (LIMIT_PER_PAGE * request_page)
    OFFSET := (LIMIT_PER_PAGE * request_page)

    //query_str := fmt.Sprintf(
    //    "SELECT t.* FROM ((SELECT * FROM %s WHERE token_type='%s' LIMIT %d,%d) UNION (SELECT * FROM %s LIMIT %d,%d))t ORDER BY timestamp DESC",
    //    TABLE_NAME__ETH_ERC20, TOKEN_TYPE_ERC20, OFFSET, LIMIT_PER_PAGE,
    //    TABLE_NAME__ERC1155, OFFSET, LIMIT_PER_PAGE,
    //)

    query_str := fmt.Sprintf(
        "SELECT * FROM %s WHERE token_type='%s' or token_type='%s' ORDER BY timestamp DESC LIMIT %d,%d",
        TABLE_NAME__TXID_ALL, TOKEN_TYPE_ERC20, TOKEN_TYPE_ERC1155, OFFSET, LIMIT_PER_PAGE,
    )


    *response = t._query( _E_TYPE__ALL, query_str )

    return nil
}

func (t *LocalDB) Db_select_txns_all_erc(request_page uint, response *[]types.Fetch_transactions_st) error {
    LIMIT_PER_PAGE := uint(30)
    t._Db_select_txns_all_erc( LIMIT_PER_PAGE, request_page, response )
    return nil
}

func (t *LocalDB) _Db_select_txns_erc20(LIMIT_PER_PAGE uint, request_page uint, response *[]types.Fetch_transactions_st) error {
    // limit: 30 transactions per page
    //LIMIT_PER_PAGE := uint(30)
    //REQ_ROWS := (LIMIT_PER_PAGE * request_page)
    OFFSET := (LIMIT_PER_PAGE * request_page)

    //query_str := fmt.Sprintf(
    //    "SELECT * FROM %s WHERE token_type = '%s' ORDER BY timestamp DESC LIMIT %d,%d",
    //    TABLE_NAME__ETH_ERC20, OFFSET, LIMIT_PER_PAGE, TOKEN_TYPE_ERC20,
    //)
    //*response = t._query( _E_TYPE__ERC20, query_str )

    query_str := fmt.Sprintf(
        "SELECT * FROM %s WHERE token_type = '%s' ORDER BY timestamp DESC LIMIT %d,%d",
        TABLE_NAME__TXID_ALL, TOKEN_TYPE_ERC20, OFFSET, LIMIT_PER_PAGE,
    )


    *response = t._query( _E_TYPE__ALL, query_str )

    return nil
}

func (t *LocalDB) Db_select_txns_erc20(request_page uint, response *[]types.Fetch_transactions_st) error {
    LIMIT_PER_PAGE := uint(30)
    t._Db_select_txns_erc20( LIMIT_PER_PAGE, request_page, response )
    return nil
}

func (t *LocalDB) _Db_select_txns_erc1155(LIMIT_PER_PAGE uint, request_page uint, response *[]types.Fetch_transactions_st) error {
    // limit: 30 transactions per page
    //LIMIT_PER_PAGE := uint(30)
    //REQ_ROWS := (LIMIT_PER_PAGE * request_page)
    OFFSET := (LIMIT_PER_PAGE * request_page)

    //query_str := fmt.Sprintf(
    //    "SELECT * FROM %s ORDER BY timestamp DESC LIMIT %d,%d",
    //    TABLE_NAME__ERC1155, OFFSET, LIMIT_PER_PAGE,
    //)
    //*response = t._query( _E_TYPE__ERC1155, query_str )

    query_str := fmt.Sprintf(
        "SELECT * FROM %s WHERE token_type = '%s' ORDER BY timestamp DESC LIMIT %d,%d",
        TABLE_NAME__TXID_ALL, TOKEN_TYPE_ERC1155, OFFSET, LIMIT_PER_PAGE,
    )


    *response = t._query( _E_TYPE__ALL, query_str )

    return nil
}

func (t *LocalDB) Db_select_txns_erc1155(request_page uint, response *[]types.Fetch_transactions_st) error {
    LIMIT_PER_PAGE := uint(30)
    t._Db_select_txns_erc1155( LIMIT_PER_PAGE, request_page, response )
    return nil
}

func (t *LocalDB) _Db_select_blocks_info(LIMIT_PER_PAGE uint, request_page uint, response *[]types.Fetch_blocks_info_st) error {
    //if ( request_page < 1 ) {
    //    fmt.Println( "request page < 1" )
    //    return nil
    //}

    // limit: 30 transactions per page
    //LIMIT_PER_PAGE := uint(30)
    //REQ_ROWS := (LIMIT_PER_PAGE * request_page)
    OFFSET := (LIMIT_PER_PAGE * request_page)

    query_str := fmt.Sprintf(
        //"SELECT * FROM %s ORDER BY blocks DESC LIMIT %d,%d",

        //! FIXME
        //"SELECT * FROM %s ORDER BY CAST(blocks as unsigned) DESC LIMIT %d,%d",
        "SELECT * FROM %s ORDER BY blocks+0 DESC LIMIT %d,%d",
        TABLE_NAME__BLOCKS, OFFSET, LIMIT_PER_PAGE,
    )


    *response = t._query_blocks( query_str )

    return nil
}

func (t *LocalDB) Db_select_blocks_info(request_page uint, response *[]types.Fetch_blocks_info_st) error {
    LIMIT_PER_PAGE := uint(30)
    t._Db_select_blocks_info( LIMIT_PER_PAGE, request_page, response )
    return nil
}

func (t *LocalDB) _Db_select_block_by_number(block string, response *[]types.Fetch_blocks_info_st) error {
    //if ( request_page < 1 ) {
    //    fmt.Println( "request page < 1" )
    //    return nil
    //}

    // limit: 30 transactions per page
    //LIMIT_PER_PAGE := uint(30)
    //REQ_ROWS := (LIMIT_PER_PAGE * request_page)
    //OFFSET := (LIMIT_PER_PAGE * request_page)

    query_str := fmt.Sprintf(
        //"SELECT * FROM %s ORDER BY blocks DESC LIMIT %d,%d",

        "SELECT * FROM %s WHERE blocks = '%s'",
        TABLE_NAME__BLOCKS, block,
    )


    *response = t._query_blocks( query_str )

    return nil
}

func (t *LocalDB) Db_select_block_by_number(block string, response *[]types.Fetch_blocks_info_st) error {
    //LIMIT_PER_PAGE := uint(30)
    //t._Db_select_block_by_number( LIMIT_PER_PAGE, block, response )
    t._Db_select_block_by_number( block, response )
    return nil
}

func (t *LocalDB) _Db_select_block_number_by_block_hash(block_hash string, response *[]types.Fetch_blocks_info_st) error {
    //if ( request_page < 1 ) {
    //    fmt.Println( "request page < 1" )
    //    return nil
    //}

    // limit: 30 transactions per page
    //LIMIT_PER_PAGE := uint(30)
    //REQ_ROWS := (LIMIT_PER_PAGE * request_page)
    //OFFSET := (LIMIT_PER_PAGE * request_page)

    query_str := fmt.Sprintf(
        //"SELECT * FROM %s ORDER BY blocks DESC LIMIT %d,%d",

        "SELECT blocks FROM %s WHERE blocks_hash = '%s'",
        TABLE_NAME__BLOCKS, block_hash,
    )


    *response = t._query_blocks( query_str )


    return nil
}

func (t *LocalDB) Db_select_block_number_by_block_hash(block_hash string, response *[]types.Fetch_blocks_info_st) error {
    //LIMIT_PER_PAGE := uint(30)
    //t._Db_select_block_by_block_hash( LIMIT_PER_PAGE, block_hash, response )
    t._Db_select_block_number_by_block_hash( block_hash, response )
    return nil
}

func (t *LocalDB) _Db_select_txns_all_mixed_by_block_number(LIMIT_PER_PAGE uint, request_page uint, block string, response *[]types.Fetch_transactions_st) error {
    //if ( request_page < 1 ) {
    //    fmt.Println( "request page < 1" )
    //    return nil
    //}

    // limit: 30 transactions per page
    //LIMIT_PER_PAGE := uint(30)
    //REQ_ROWS := (LIMIT_PER_PAGE * request_page)
    OFFSET := (LIMIT_PER_PAGE * request_page)

    query_str := fmt.Sprintf(
        "SELECT * FROM %s WHERE blocks = '%s' ORDER BY timestamp DESC LIMIT %d,%d",
        TABLE_NAME__TXID_ALL, block, OFFSET, LIMIT_PER_PAGE,
    )


    *response = t._query( _E_TYPE__ALL, query_str )

    return nil
}

func (t *LocalDB) Db_select_txns_all_mixed_by_block_number(request_page uint, block string, response *[]types.Fetch_transactions_st) error {
    LIMIT_PER_PAGE := uint(30)
    t._Db_select_txns_all_mixed_by_block_number( LIMIT_PER_PAGE, request_page, block, response )
    return nil
}

func (t *LocalDB) _Db_select_txns_all_mixed_by_block_hash(LIMIT_PER_PAGE uint, request_page uint, block_hash string, response *[]types.Fetch_transactions_st) error {
    //if ( request_page < 1 ) {
    //    fmt.Println( "request page < 1" )
    //    return nil
    //}


    query_str := fmt.Sprintf(
        //"SELECT * FROM %s ORDER BY blocks DESC LIMIT %d,%d",

        "SELECT * FROM %s WHERE blocks_hash = '%s'",
        TABLE_NAME__BLOCKS, block_hash,
    )


    response_block_number := t._query_blocks( query_str )
    //result, err_marshal := json.Marshal( response_block_number )
    //if err_marshal != nil {
    //    panic( err_marshal.Error() )
    //}

    block := response_block_number[0].Block_number


    // limit: 30 transactions per page
    //LIMIT_PER_PAGE := uint(30)
    //REQ_ROWS := (LIMIT_PER_PAGE * request_page)
    OFFSET := (LIMIT_PER_PAGE * request_page)

    query_str = fmt.Sprintf(
        "SELECT * FROM %s WHERE blocks = '%s' ORDER BY timestamp DESC LIMIT %d,%d",
        TABLE_NAME__TXID_ALL, block, OFFSET, LIMIT_PER_PAGE,
    )


    *response = t._query( _E_TYPE__ALL, query_str )

    return nil
}

func (t *LocalDB) Db_select_txns_all_mixed_by_block_hash(request_page uint, block_hash string, response *[]types.Fetch_transactions_st) error {
    LIMIT_PER_PAGE := uint(30)
    t._Db_select_txns_all_mixed_by_block_hash( LIMIT_PER_PAGE, request_page, block_hash, response )
    return nil
}

func (t *LocalDB) _Db_select_txns_all_mixed_by_txid(LIMIT_PER_PAGE uint, request_page uint, txid string, response *[]types.Fetch_transactions_st) error {
    //if ( request_page < 1 ) {
    //    fmt.Println( "request page < 1" )
    //    return nil
    //}

    // limit: 30 transactions per page
    //LIMIT_PER_PAGE := uint(30)
    //REQ_ROWS := (LIMIT_PER_PAGE * request_page)
    OFFSET := (LIMIT_PER_PAGE * request_page)

    query_str := fmt.Sprintf(
        "SELECT * FROM %s WHERE txid = '%s' ORDER BY timestamp DESC LIMIT %d,%d",
        TABLE_NAME__TXID_ALL, txid, OFFSET, LIMIT_PER_PAGE,
    )


    *response = t._query( _E_TYPE__ALL, query_str )

    return nil
}

func (t *LocalDB) Db_select_txns_all_mixed_by_txid(request_page uint, txid string, response *[]types.Fetch_transactions_st) error {
    LIMIT_PER_PAGE := uint(30)
    t._Db_select_txns_all_mixed_by_txid( LIMIT_PER_PAGE, request_page, txid, response )
    return nil
}

func (t *LocalDB) _Db_select_txns_all_mixed_by_address(LIMIT_PER_PAGE uint, request_page uint, address string, response *[]types.Fetch_transactions_st) error {
    //if ( request_page < 1 ) {
    //    fmt.Println( "request page < 1" )
    //    return nil
    //}

    // limit: 30 transactions per page
    //LIMIT_PER_PAGE := uint(30)
    //REQ_ROWS := (LIMIT_PER_PAGE * request_page)
    OFFSET := (LIMIT_PER_PAGE * request_page)

    query_str := fmt.Sprintf(
        "SELECT * FROM %s WHERE (from_address = '%s' or to_address = '%s') ORDER BY timestamp DESC LIMIT %d,%d",
        TABLE_NAME__TXID_ALL, address, address, OFFSET, LIMIT_PER_PAGE,
    )


    *response = t._query( _E_TYPE__ALL, query_str )

    return nil
}

func (t *LocalDB) Db_select_txns_all_mixed_by_address(request_page uint, address string, response *[]types.Fetch_transactions_st) error {
    LIMIT_PER_PAGE := uint(30)
    t._Db_select_txns_all_mixed_by_address( LIMIT_PER_PAGE, request_page, address, response )
    return nil
}

//! TODO
func (t *LocalDB) _Db_select_balances_all_mixed(LIMIT_PER_PAGE uint, request_page uint, response *[]types.Fetch_balances_st) error {
    fmt.Println( "TODO: _Db_select_balances_all_mixed()" )

    //if ( request_page < 1 ) {
    //    fmt.Println( "request page < 1" )
    //    return nil
    //}

    // limit: 30 balances per page
    //LIMIT_PER_PAGE := uint(30)
    //REQ_ROWS := (LIMIT_PER_PAGE * request_page)
    OFFSET := (LIMIT_PER_PAGE * request_page)


    //! ALL balances
    query_str := fmt.Sprintf(
        "SELECT * FROM %s LIMIT %d,%d",
        TABLE_NAME__BALANCES_ALL, OFFSET, LIMIT_PER_PAGE,
    )


    *response = t._query_balances( query_str )

    return nil
}

func (t *LocalDB) Db_select_balances_all_mixed(request_page uint, response *[]types.Fetch_balances_st) error {
    LIMIT_PER_PAGE := uint(30)
    t._Db_select_balances_all_mixed( LIMIT_PER_PAGE, request_page, response )
    return nil
}

//! TODO
func (t *LocalDB) _Db_select_balances_by_address(request_page uint, address string, response *[]types.Fetch_balances_by_address_st) error {
    fmt.Println( "TODO: _Db_select_balances_by_address()" )

    //if ( request_page < 1 ) {
    //    fmt.Println( "request page < 1" )
    //    return nil
    //}

    // limit: 30 balances per page
    //LIMIT_PER_PAGE := uint(30)
    //REQ_ROWS := (LIMIT_PER_PAGE * request_page)
    //OFFSET := (LIMIT_PER_PAGE * request_page)



    //! FIXME: balance table data from txid (search performance issues, double space, ...)
    // balance table (balance_0 ~ balance_f): idx, address, txid, token_type, token_symbol, value,... ???
    // NOTE: save every query into log table


    /*
    //! ALL balances (sum of all) from 'txid' table
    query_str := fmt.Sprintf(
        //"SELECT * FROM %s WHERE (from_address = '%s' or to_address = '%s') ORDER BY timestamp DESC LIMIT %d,%d",
        //TABLE_NAME__BALANCES_ALL, address, address, OFFSET, LIMIT_PER_PAGE,


        //! TEST: from 'txid' table
        "SELECT a.token_type, a.token_contract_address, a.token_symbol, a.count_symbol, a.value_ether, a.value_token, a.value_erc1155 FROM",
        " (SELECT token_type, token_contract_address, token_symbol, count(token_symbol) count_symbol, SUM(ether_amount_eth) as value_ether,",
        " SUM(token_amount_eth) as value_token, SUM(erc1155_token_amount) as value_erc1155 FROM %s",
        " WHERE (from_address='%s' or to_address='%s')",
        " GROUP BY token_type, token_contract_address, token_symbol ) AS a",

        TABLE_NAME__BALANCES_ALL, address, address,
    )
    */

    /*
    mysql> select count(owner_address), sum(cast(amount as unsigned)), sum(case when issend = "y" then (-1 * cast(amount as signed)) else cast(amount as unsigned) end) from balances_address_suffix_3 where (owner_address="0xe6e55eed00218faef27eed24def9208f3878b333" and token_type="erc1155") group by token_contract_address;
    +----------------------+-------------------------------+--------------------------------------------------------------------------------------------------+
    | count(owner_address) | sum(cast(amount as unsigned)) | sum(case when issend = "y" then (-1 * cast(amount as signed)) else cast(amount as unsigned) end) |
    +----------------------+-------------------------------+--------------------------------------------------------------------------------------------------+
    |                    4 |           1000000000000000003 |                                                                               999999999999999997 |
    +----------------------+-------------------------------+--------------------------------------------------------------------------------------------------+
    1 row in set (0.00 sec)
    */

    balance_table_name := fmt.Sprintf( "balances_address_suffix_%s", address[len(address)-1:] )

    //! ALL balances
    query_str := fmt.Sprintf(
        "SELECT * FROM %s WHERE owner_address = '%s'",
        //TABLE_NAME__BALANCES_ALL, address,
        balance_table_name, address,
    )


    *response = t._query_balances_by_address( query_str )

    return nil
}

func (t *LocalDB) Db_select_balances_by_address(request_page uint, address string, response *[]types.Fetch_balances_by_address_st) error {
    //LIMIT_PER_PAGE := uint(30)
    t._Db_select_balances_by_address( request_page, address, response )
    return nil
}


