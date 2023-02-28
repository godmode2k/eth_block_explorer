/* --------------------------------------------------------------
Project:    Ethereum auto-transfer (accounts to specific address(hotwallet))
Purpose:
Author:     Ho-Jung Kim (godmode2k@hotmail.com)
Date:       Since Dec 4, 2020
Filename:   rpc_server_db_memory.go

Last modified:  April 20, 2022
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
    "log"
    //"math/big"
    "encoding/json"

    //"runtime"
    //"regexp"

    // HTTP JSON-RPC Server
    //"net/http"
    //"github.com/gorilla/mux"
    //"github.com/gorilla/rpc"
    //"github.com/gorilla/rpc/json"

    // $ go get -u github.com/go-sql-driver/mysql
    //"database/sql"
    _ "github.com/go-sql-driver/mysql"
    _ "github.com/mattn/go-sqlite3"

    // eth_txns_fetcher
    "eth_block_explorer/types"
)



//! Definition
// --------------------------------------------------------------------



//! Implementation
// --------------------------------------------------------------------



/*
func (t *LocalDB) _query_memory(query_str string) []types.Fetch_transactions_st {
    var result []types.Fetch_transactions_st

    //rows, err := t.Db.Query("select id, name from foo limit 1")
    //rows, err := t.Db.Query( query )
    //if err != nil {
    //    log.Fatal(err)
    //}

    fmt.Println( "Query = ", query_str )

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



// ---------------------------------------------------------------



func (t *LocalDB) Db_memory_update_txns_all_mixed() error {
    var request_page = uint(0)
    var result []types.Fetch_transactions_st


    t._Db_select_txns_all_mixed( DEFAULT_LIMIT_PER_PAGE, request_page, &result )


    //query := "DROP TABLE memory_txid;"
    query := "DROP TABLE " + DB_MEMORY_TABLE_NAME__TXID_ALL + ";"
    _, _err := t.DbMemory.Exec( query )
    if _err != nil {
        //log.Printf( "Error: %q: %s\n", _err, query )
        fmt.Printf( "Error: %q: %s\n", _err, query )
        panic( _err.Error() )
    }
    //query = "CREATE TABLE memory_txid (idx integer not null primary key autoincrement, txid_json text);"
    query = "CREATE TABLE " + DB_MEMORY_TABLE_NAME__TXID_ALL + " (idx integer not null primary key autoincrement,"
    query += " txid_json text);"
    _, _err = t.DbMemory.Exec( query )
    if _err != nil {
        //log.Printf( "Error: %q: %s\n", _err, query )
        fmt.Printf( "Error: %q: %s\n", _err, query )
        panic( _err.Error() )
    }


    //fmt.Printf( "Db_memory_update_txns_all_mixed(): = %s, \n", result[0].From_address )
    for i := 0; i < len(result); i++ {
        //fmt.Printf( "Db_memory_update_txns_all_mixed(): [%d] = %s, \n", i, result[i] )
        res_json := types.Fetch_transactions_json( &result[i] )
        //fmt.Printf( "Db_memory_update_txns_all_mixed(): [%d] = %s, \n", i, res_json )

        tx, err := t.DbMemory.Begin()
        if err != nil {
            //log.Fatal( err )
            panic( err.Error() )
        }
        //stmt, err := tx.Prepare( "INSERT INTO " + DB_MEMORY_TABLE_NAME__TXID_ALL + "(idx, txid_json) VALUES(?, ?)" )
        stmt, err := tx.Prepare( "INSERT INTO " + DB_MEMORY_TABLE_NAME__TXID_ALL + "(txid_json) VALUES(?)" )
        if err != nil {
            //log.Fatal( err )
            panic( err.Error() )
        }
        defer stmt.Close()

        //_, err = stmt.Exec( i, res_json )
        _, err = stmt.Exec( res_json )
        if err != nil {
            //log.Fatal( err )
            panic( err.Error() )
        }
        tx.Commit()
    }


    //t.Db_memory_select_txns_all_mixed()
    /*
    query := "SELECT idx, txid_json FROM " + DB_MEMORY_TABLE_NAME__TXID_ALL

    rows, err := t.DbMemory.Query( query )
    if err != nil {
        //log.Printf( "Error: %q: %s\n", err, query )
        fmt.Printf( "Error: %q: %s\n", err, query )
        panic( err.Error() )
    }

    for rows.Next() {
        var idx int
        var txid_json string
        err = rows.Scan( &idx, &txid_json )

        if err != nil {
            //log.Fatal(err)
            panic( err.Error() )
        }

        fmt.Println( fmt.Sprintf( "idx = %d, txid_json = %s\n", idx, txid_json ) )
    }

    err = rows.Err()
    if err != nil {
        //log.Fatal(err)
        panic( err.Error() )
    }
    */

    return nil
}

func (t *LocalDB) Db_memory_update_blocks_info() error {
    var request_page = uint(0)
    var result []types.Fetch_blocks_info_st


    t._Db_select_blocks_info( DEFAULT_LIMIT_PER_PAGE, request_page, &result )


    //query := "DROP TABLE memory_blocks;"
    query := "DROP TABLE " + DB_MEMORY_TABLE_NAME__BLOCKS_INFO + ";"
    _, _err := t.DbMemory.Exec( query )
    if _err != nil {
        //log.Printf( "Error: %q: %s\n", _err, query )
        fmt.Printf( "Error: %q: %s\n", _err, query )
        panic( _err.Error() )
    }
    //query = "CREATE TABLE memory_blocks (idx integer not null primary key autoincrement, blocks text, blocks_hash text, info text, transactions text);"
    query = "CREATE TABLE " + DB_MEMORY_TABLE_NAME__BLOCKS_INFO + " (idx integer not null primary key autoincrement,"
    query += " blocks text, blocks_hash text, info text, transactions text);"
    _, _err = t.DbMemory.Exec( query )
    if _err != nil {
        //log.Printf( "Error: %q: %s\n", _err, query )
        fmt.Printf( "Error: %q: %s\n", _err, query )
        panic( _err.Error() )
    }


    //fmt.Printf( "Db_memory_update_blocks_info(): = %s, \n", result[0].Block_number )
    for i := 0; i < len(result); i++ {
        //fmt.Printf( "Db_memory_update_blocks_info(): [%d] = %s, \n", i, result[i] )
        //res_json := types.Fetch_blocks_json( &result[i] )

        tx, err := t.DbMemory.Begin()
        if err != nil {
            //log.Fatal( err )
            panic( err.Error() )
        }
        //stmt, err := tx.Prepare( "INSERT INTO " + DB_MEMORY_TABLE_NAME__BLOCKS_INFO + "(idx, blocks, blocks_hash, info, transactions) VALUES(?, ?, ?, ?, ?)" )
        stmt, err := tx.Prepare( "INSERT INTO " + DB_MEMORY_TABLE_NAME__BLOCKS_INFO + "(blocks, blocks_hash, info, transactions) VALUES(?, ?, ?, ?)" )
        if err != nil {
            //log.Fatal( err )
            panic( err.Error() )
        }
        defer stmt.Close()

        _, err = stmt.Exec( //i,
                            &result[i].Block_number,
                            &result[i].Block_hash,
                            &result[i].Info,
                            &result[i].Transactions )
        if err != nil {
            //log.Fatal( err )
            panic( err.Error() )
        }
        tx.Commit()
    }


    //t.Db_memory_select_txns_all_mixed()
    /*
    query := "SELECT * FROM " + DB_MEMORY_TABLE_NAME__BLOCKS_INFO

    rows, err := t.DbMemory.Query( query )
    if err != nil {
        //log.Printf( "Error: %q: %s\n", err, query )
        fmt.Printf( "Error: %q: %s\n", err, query )
        panic( err.Error() )
    }

    for rows.Next() {
        var idx int
        var blocks string
        var blocks_hash string
        var info string
        var transactions string

        err = rows.Scan( &idx, &blocks, &blocks_hash, &info, &transactions )

        if err != nil {
            //log.Fatal(err)
            panic( err.Error() )
        }

        fmt.Println( fmt.Sprintf( "idx = %d, blocks = %s, blocks_hash = %s, info = %s, transactions = %s\n", id, blocks, blocks_hash, info, transactions ) )
    }

    err = rows.Err()
    if err != nil {
        //log.Fatal(err)
        panic( err.Error() )
    }
    */

    return nil
}

//! TODO
func (t *LocalDB) Db_memory_update_balances() error {
    log.Println( "Db_memory_update_balances()", "===== FIXME =====" )
    return nil
}

//! TODO
func (t *LocalDB) Db_memory_update_balances_by_address() error {
    log.Println( "Db_memory_update_balances_by_address()", "===== FIXME =====" )

    /*
    var request_page = uint(0)
    var result []types.Fetch_balances_by_address_st


    //t._Db_select_balances_all_mixed( DEFAULT_LIMIT_PER_PAGE, request_page, &result )
    t._Db_select_balances_by_address( request_page, &result )

    //query := "DROP TABLE memory_balances;"
    query := "DROP TABLE " + DB_MEMORY_TABLE_NAME__BALANCES_ALL + ";"
    _, _err := t.DbMemory.Exec( query )
    if _err != nil {
        //log.Printf( "Error: %q: %s\n", _err, query )
        fmt.Printf( "Error: %q: %s\n", _err, query )
        panic( _err.Error() )
    }
    query = "CREATE TABLE " + DB_MEMORY_TABLE_NAME__BALANCES_ALL  + " (idx integer not null primary key autoincrement,"
    query += " owner_address text, txid text, blocks text, token_type text, token_symbol text, token_contract_address text, amount text);"
    _, _err = t.DbMemory.Exec( query )
    if _err != nil {
        //log.Printf( "Error: %q: %s\n", _err, query )
        fmt.Printf( "Error: %q: %s\n", _err, query )
        panic( _err.Error() )
    }


    //fmt.Printf( "Db_memory_update_balances(): = %s, \n", result[0].owner_address )
    for i := 0; i < len(result); i++ {
        //fmt.Printf( "Db_memory_update_balances(): [%d] = %s, \n", i, result[i] )
        //res_json := types.Fetch_balances_json( &result[i] )

        tx, err := t.DbMemory.Begin()
        if err != nil {
            //log.Fatal( err )
            panic( err.Error() )
        }
        //stmt, err := tx.Prepare( "INSERT INTO " + DB_MEMORY_TABLE_NAME__BALANCES_ALL +
        //    "(idx, owner_address, txid, blocks, token_type, token_symbol, token_contract_address, amount) VALUES(?, ?, ?, ?, ?, ?, ?, ?)" )
        stmt, err := tx.Prepare( "INSERT INTO " + DB_MEMORY_TABLE_NAME__BALANCES_ALL +
            "(owner_address, txid, blocks, token_type, token_symbol, token_contract_address, amount) VALUES(?, ?, ?, ?, ?, ?, ?)" )
        if err != nil {
            //log.Fatal( err )
            panic( err.Error() )
        }
        defer stmt.Close()

        _, err = stmt.Exec( //i,
                            &result[i].Owner_address,
                            &result[i].Txid,
                            &result[i].Block_number,
                            &result[i].Token_type,
                            &result[i].Token_symbol,
                            &result[i].Token_contract_address,
                            &result[i].Amount )
        if err != nil {
            //log.Fatal( err )
            panic( err.Error() )
        }
        tx.Commit()
    }
    */


    //t.Db_memory_select_txns_all_mixed()
    /*
    query := "SELECT * FROM " + DB_MEMORY_TABLE_NAME__BALANCES_ALL

    rows, err := t.DbMemory.Query( query )
    if err != nil {
        //log.Printf( "Error: %q: %s\n", err, query )
        fmt.Printf( "Error: %q: %s\n", err, query )
        panic( err.Error() )
    }

    for rows.Next() {
        var idx int
        var owner_address string
        var txid string
        var blocks string
        var token_type string
        var token_symbol string
        var amount string

        err = rows.Scan( &idx, &owner_address, &txid, &blocks, &token_type, &token_symbol, &amount )

        if err != nil {
            //log.Fatal(err)
            panic( err.Error() )
        }

        fmt.Println( fmt.Sprintf( "owner_address = %d, txid = %s, blocks %s, token_type %s, token_symbol %s, amount = %s\n",
                                    owner_address, txid , blocks , token_type , token_symbol , amount ) )
    }

    err = rows.Err()
    if err != nil {
        //log.Fatal(err)
        panic( err.Error() )
    }
    */

    return nil

}



// ---------------------------------------------------------------



//func (t *LocalDB) Db_memory_select_txns_all_mixed(request_page uint, response *string) error {
func (t *LocalDB) Db_memory_select_txns_all_mixed(request_page uint, response *[]types.Fetch_transactions_st) error {
    // limit: 30 transactions per page
    LIMIT_PER_PAGE := uint(30)
    ////REQ_ROWS := (LIMIT_PER_PAGE * request_page)
    OFFSET := (LIMIT_PER_PAGE * request_page)

    //var json_map_arr = make( map[string][]interface{} )
    //var json_arr = make( []interface{} )
    var json_arr []types.Fetch_transactions_st
    ////var json_map = make( map[string]interface{} )


    query := fmt.Sprintf(
        //"SELECT idx, txid_json FROM %s limit %d,%d",
        //DB_MEMORY_TABLE_NAME__TXID_ALL, OFFSET, LIMIT_PER_PAGE,

        //! DO NOT USE "ORDER BY idx DESC"
        // SEE: Db_select_txns_all_mixed()
        // - "ORDER BY timestamp DESC" already
        "SELECT * FROM %s LIMIT %d,%d",
        DB_MEMORY_TABLE_NAME__TXID_ALL, OFFSET, LIMIT_PER_PAGE,
    )

    rows, err := t.DbMemory.Query( query )
    if err != nil {
        //log.Printf( "Error: %q: %s\n", err, query )
        fmt.Printf( "Error: %q: %s\n", err, query )
        panic( err.Error() )
    }
    defer rows.Close()

    for rows.Next() {
        var idx int
        var txid_json string

        err = rows.Scan( &idx, &txid_json )

        if err != nil {
            //log.Fatal(err)
            panic( err.Error() )
        }

        //fmt.Println( fmt.Sprintf( "idx = %d, txid_json = %s\n", idx, txid_json ) )

        //var _data = make( map[string]interface{} )
        var _data types.Fetch_transactions_st
        err = json.Unmarshal( []byte(txid_json), &_data )
        if err != nil {
            //log.Fatal(err)
            panic( err.Error() )
        }
        ////json_map["txid"] = _data
        //json_map_arr["txid"] = append( json_map_arr["txid"], _data )
        json_arr = append( json_arr, _data )
    }

    err = rows.Err()
    if err != nil {
        //log.Fatal(err)
        panic( err.Error() )
    }

    ////result, err_marshal := json.Marshal( json_map )
    //result, err_marshal := json.Marshal( json_map_arr )
    //if err_marshal != nil {
    //    panic( err_marshal.Error() )
    //}


    //*response = string(result)
    *response = json_arr

    return nil
}

func (t *LocalDB) Db_memory_select_blocks_info(request_page uint, response *[]types.Fetch_blocks_info_st) error {
    // limit: 30 transactions per page
    LIMIT_PER_PAGE := uint(30)
    ////REQ_ROWS := (LIMIT_PER_PAGE * request_page)
    OFFSET := (LIMIT_PER_PAGE * request_page)

    //var json_map_arr = make( map[string][]interface{} )
    //var json_arr = make( []interface{} )
    var json_arr []types.Fetch_blocks_info_st
    ////var json_map = make( map[string]interface{} )


    query := fmt.Sprintf(
        //"SELECT idx, blocks, blocks_hash, info, transactions FROM %s limit %d,%d",
        //DB_MEMORY_TABLE_NAME__BLOCKS_INFO, OFFSET, LIMIT_PER_PAGE,

        //! DO NOT USE "ORDER BY blocks DESC"
        // SEE: Db_select_blocks_info()
        // - "ORDER BY blocks DESC" already
        "SELECT * FROM %s LIMIT %d,%d",
        DB_MEMORY_TABLE_NAME__BLOCKS_INFO, OFFSET, LIMIT_PER_PAGE,
    )

    rows, err := t.DbMemory.Query( query )
    if err != nil {
        //log.Printf( "Error: %q: %s\n", err, query )
        fmt.Printf( "Error: %q: %s\n", err, query )
        panic( err.Error() )
    }
    defer rows.Close()

    for rows.Next() {
        var idx int
        var blocks string
        var blocks_hash string
        var info string
        var transactions string

        err = rows.Scan( &idx, &blocks, &blocks_hash, &info, &transactions )

        if err != nil {
            //log.Fatal(err)
            panic( err.Error() )
        }

        //fmt.Println( fmt.Sprintf( "idx = %d, blocks = %s, blocks_hash = %s, info = %s, transactions = %s\n", idx, blocks, blocks_hash, info, transactions ) )

        //var _data = make( map[string]interface{} )
        var _data = types.Fetch_blocks_info_st { Block_number: blocks, Block_hash: blocks_hash, Info: info, Transactions: transactions }
        /*
        err = json.Unmarshal( []byte(info), &_data )
        if err != nil {
            //log.Fatal(err)
            panic( err.Error() )
        }
        */
        ////json_map["blocks_info"] = _data
        //json_map_arr["blocks_info"] = append( json_map_arr["blocks_info"], _data )
        json_arr = append( json_arr, _data )
    }

    err = rows.Err()
    if err != nil {
        //log.Fatal(err)
        panic( err.Error() )
    }

    ////result, err_marshal := json.Marshal( json_map )
    //result, err_marshal := json.Marshal( json_map_arr )
    //if err_marshal != nil {
    //    panic( err_marshal.Error() )
    //}


    //*response = string(result)
    *response = json_arr

    return nil
}



