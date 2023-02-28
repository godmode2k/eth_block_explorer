/* --------------------------------------------------------------
Project:    Ethereum fetch all transactions
Purpose:
Author:     Ho-Jung Kim (godmode2k@hotmail.com)
Date:       Since Dec 4, 2020
Filename:   rpc_server_main.go

Last modified:  May 12, 2022
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
 - https://pkg.go.dev/database/sql

Dependencies:
$ go get -u github.com/go-sql-driver/mysql


1. Build:
	$ go build rpc_server_main.go
    or
	$ go run rpc_server_main.go
-------------------------------------------------------------- */
package main



//! Header
// ---------------------------------------------------------------

import (
    "fmt"
    "log"
    "time"

    // sync.WaitGroup
    // sync.Mutex
    "sync"
    // context.Context: context.WithCancel(context.Background())
    //"context"

    // test
    //"math/rand"

    // HTTP RPC
    "net"
    "net/http"
    "net/rpc"

    // HTTP JSON-RPC
    gorilla_mux "github.com/gorilla/mux"
    gorilla_rpc "github.com/gorilla/rpc"
    gorilla_json "github.com/gorilla/rpc/json"

    // $ go get -u github.com/go-sql-driver/mysql
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    _ "github.com/mattn/go-sqlite3"

    // eth_txns_fetcher
    //"eth_block_explorer/types"
    "eth_block_explorer/server"

    //"reflect"
)



//! Definition
// --------------------------------------------------------------------

var SERVER_ADDRESS = "127.0.0.1"
var SERVER_PORT = "8544"
var SERVER = SERVER_ADDRESS + ":" + SERVER_PORT
var URL = "http://" + SERVER_ADDRESS + ":" + SERVER_PORT
var DB_SERVER_ADDRESS = "127.0.0.1:3306"
var DB_NAME = "ethereum_block_explorer"
var DB_LOGIN_USERNAME = "root"
var DB_LOGIN_PASSWORD = "mysql"
var gDB *sql.DB
var DB_MEMORY_NAME = "memory_ethereum_block_explorer"
var gDBMemory *sql.DB
//var _E_TYPE__ETH = uint8(0)
//var _E_TYPE__ERC20 = uint8(1)
//var _E_TYPE__ERC1155 = uint8(2)


// for HTTP RPC Server
var HTTP_RPC_SERVER_HOST_PORT = ":1234" // Internal
var HTTP_JSONRPC_SERVER_HOST_PORT = ":1235" // External
var g_localdb = new( rpc_server.LocalDB )
var UPDATES_INTERVAL = int(30) // 30 seconds
var gWG sync.WaitGroup
var gChan = make( chan uint8, 1 )
var _E_CHAN__DONE = uint8(0)
var _E_CHAN__CANCEL = uint8(1)


//var gMutex sync.Mutex
// mutex.Lock()
// defer mutex.Unlock()




//! Implementation
// --------------------------------------------------------------------

func __init_db() bool {
    fmt.Println( "main: __init_db(): initialize..." )

    db, err := sql.Open( "mysql", DB_LOGIN_USERNAME + ":" + DB_LOGIN_PASSWORD + "@tcp(" + DB_SERVER_ADDRESS + ")/" + DB_NAME )

    if err != nil {
        panic( err.Error() )
    }

    gDB = db



    // In-memory DB (SQLite)
    //db_sqlite, err := sql.Open("sqlite3", "./localdb_sqlite.db")
    //db_sqlite, err := sql.Open("sqlite3", "file::memory:?mode=memory&cache=shared")
    //db_sqlite_filename := randomString(16) // func creates random string
    db_sqlite_filename := "localdb_sqlite3.db"
    db_sqlite, err := sql.Open( "sqlite3", fmt.Sprintf("file:%s?mode=memory&cache=shared", db_sqlite_filename) )
    if err != nil {
        //log.Fatal(err)
        panic( err.Error() )
    }


    // memory_txid
    query := "CREATE TABLE " + rpc_server.DB_MEMORY_TABLE_NAME__TXID_ALL + " (idx integer not null primary key autoincrement,"
    query += " txid_json text);"
    _, err = db_sqlite.Exec( query )
    if err != nil {
        //log.Printf( "Error: %q: %s\n", err, query )
        fmt.Printf( "Error: %q: %s\n", err, query )
        panic( err.Error() )
    }

    //! FIXME: types
    // memory_blocks
    query = "CREATE TABLE " + rpc_server.DB_MEMORY_TABLE_NAME__BLOCKS_INFO + " (idx integer not null primary key autoincrement,"
    query += " blocks text, blocks_hash text, info text, transactions text);"
    _, err = db_sqlite.Exec( query )
    if err != nil {
        //log.Printf( "Error: %q: %s\n", err, query )
        fmt.Printf( "Error: %q: %s\n", err, query )
        panic( err.Error() )
    }

    //! FIXME: types
    // memory_balances
    query = "CREATE TABLE " + rpc_server.DB_MEMORY_TABLE_NAME__BALANCES_ALL  + " (idx integer not null primary key autoincrement,"
    query += " owner_address text, txid text, blocks text, token_type text, token_symbol text, amount text);"
    _, err = db_sqlite.Exec( query )
    if err != nil {
        //log.Printf( "Error: %q: %s\n", err, query )
        fmt.Printf( "Error: %q: %s\n", err, query )
        panic( err.Error() )
    }


    gDBMemory = db_sqlite



    g_localdb.Db = gDB
    g_localdb.DbMemory = gDBMemory

    return true
}

/*
func __test_db() {
    var TABLE_NAME = "txid"
    //var TABLE_NAME_ERC1155 = "txid_erc1155"

    var query_str = fmt.Sprintf(
        "INSERT INTO %s VALUES (0," +
        "'%s', '%s', '%s', '%s', '%s'," +
        "'%s', '%s', '%s', '%s', '%s'," +
        //"'%s', '%s', '%s', '%s', '%s'," +
        "'%s', '%s'" +
        ")",
        TABLE_NAME,
        //"eth", "from address", "to address", "is send", "amount wei", "amount eth",
        "eth", "from_address", "to_address", "y", "0", "0",
        //"token type", "token symbol", "token decimals", "token contract address",
        "token_type", "token_symbol", "0", "token_contract_address",
        //"token amount wei", "token amount eth", "timestamp", "datetime", "block number", "txid"
        "0", "0", "timestamp", "datetime", "1000", "txid1" )

    fmt.Println( "query = ", query_str )

    //result, err := gDB.Query( "INSERT INTO test VALUES ( 2, 'TEST' )" )
    result, err := gDB.Query( query_str )

    if err != nil {
        panic( err.Error() )
    }

    fmt.Println( "result = ", result )
}
*/

/*
// type: (0) ETH, (1) ERC-20, (2) ERC-1155
func db_insert_txns(_type uint8, _data *types.Fetch_transactions_st) {
    var TABLE_NAME = ""
    //var TABLE_NAME_ERC1155 = "txid_erc1155"

    var query_str = ""

    //! test: random 40 bytes hex
    ////fmt.Println( "test: ", fmt.Sprintf("0x%s", strconv.FormatUint(rand.Uint64(), 16)) )
    //var test_txid = "0x"
    //for i := 0; i < 4; i++ {
    //    test_txid += fmt.Sprintf("%s", strconv.FormatUint(rand.Uint64(), 16) )
    //}
    //fmt.Println( "test: txid = ", test_txid )
    //_data.Txid = test_txid

    // ETH
    if _type == 0 {
        TABLE_NAME = "txid"
        query_str = fmt.Sprintf(
            "INSERT INTO %s VALUES (0," +
            "'%s', '%s', '%s', '%s'," +
            "'%s', '%s', '%s', '%s', '%s'," +
            //"'%s', '%s', '%s', '%s', '%s'," +
            "'%s', '%s'" +
            ")",

            TABLE_NAME,

            //_data.Symbol,
            _data.From_address,
            _data.To_address,
            //_data.Is_send,
            _data.Amount_wei,
            _data.Amount_eth,
            _data.Token_type,
            _data.Token_symbol,
            _data.Token_decimals,
            _data.Token_total_supply,
            _data.Token_contract_address,
            _data.Token_amount_wei,
            _data.Token_amount_eth,
            // for SQL: txid_erc1155 table
            //_data.Token_amount,
            //_data.Token_uri_ascii,
            //_data.Token_uri_hexadecimal,
            //_data.Token_data_length,
            //_data.Token_data,
            _data.Timestamp,
            _data.Datetime,
            _data.Block_number,
            _data.Txid )
    } else {
        // ERC-20
        if _type == 1 {
            TABLE_NAME = "txid"
            query_str = fmt.Sprintf(
                "INSERT INTO %s VALUES (0," +
                "'%s', '%s', '%s', '%s'," +
                "'%s', '%s', '%s', '%s', '%s'," +
                "'%s', '%s', '%s', '%s', '%s'," +
                "'%s', '%s'" +
                ")",

                TABLE_NAME,

                //_data.Symbol,
                _data.From_address,
                _data.To_address,
                //_data.Is_send,
                _data.Amount_wei,
                _data.Amount_eth,
                _data.Token_type,
                _data.Token_symbol,
                _data.Token_decimals,
                _data.Token_total_supply,
                _data.Token_contract_address,
                _data.Token_amount_wei,
                _data.Token_amount_eth,
                // for SQL: txid_erc1155 table
                //_data.Token_amount,
                //_data.Token_uri_ascii,
                //_data.Token_uri_hexadecimal,
                //_data.Token_data_length,
                //_data.Token_data,
                _data.Timestamp,
                _data.Datetime,
                _data.Block_number,
                _data.Txid )
        } else if _type == 2 {
            // ERC-1155
            TABLE_NAME = "txid_erc1155"
            query_str = fmt.Sprintf(
                "INSERT INTO %s VALUES (0," +
                "'%s', '%s', '%s', '%s'," +
                "'%s', '%s', '%s', '%s', '%s'," +
                "'%s', '%s', '%s', '%s', '%s'," +
                //"'%s', '%s', '%s', '%s', '%s'," +
                "'%s', '%s'" +
                ")",

                TABLE_NAME,

                //_data.Symbol,
                _data.From_address,
                _data.To_address,
                //_data.Is_send,
                // for SQL: txid table
                //_data.Amount_wei,
                //_data.Amount_eth,
                _data.Token_type,
                _data.Token_symbol,
                // for SQL: txid table
                _data.Token_decimals,
                //_data.Token_total_supply,
                _data.Token_contract_address,
                // for SQL: txid table
                //_data.Token_amount_wei,
                //_data.Token_amount_eth,
                _data.Token_amount,
                _data.Token_uri_ascii,
                _data.Token_uri_hexadecimal,
                _data.Token_data_length,
                _data.Token_data,
                _data.Timestamp,
                _data.Datetime,
                _data.Block_number,
                _data.Txid )

        } else {
            panic( "Cannot found coin type" )
        }
    }

    fmt.Println( "query = ", query_str )

    //result, err := gDB.Query( "INSERT INTO test VALUES ( 2, 'TEST' )" )
    result, err := gDB.Query( query_str )

    if err != nil {
        panic( err.Error() )
    }

    fmt.Println( "result = ", result )
}
*/



// --------------------------------------------------------------------



//func run_worker_cache(ctx context.Context, ch chan int) {
//    for {
//        select {
//        case <-ctx.Done():
//            fmt.Println( "run_worker_cache()", "context: Done" )
//            close( ch )
//            gWG.Done()
//            break
//        case <-ch:
//            fmt.Println( "run_worker_cache()", "chan: ", <-ch )
//            // ch: 0, 1, 2, ...
//        }
//    }
//
//    fmt.Println( "run_worker_cache()", "finished..." )
//}

// Goroutine
func run_worker_cache() {
    if g_localdb == nil {
        panic( "DB object == NULL" )
    }

    fmt.Println( "run_worker_cache()", "Starting caching..." )


    // SEE: var g_localdb = new( rpc_server.LocalDB )
    //g_localdb.Db_memory_update_txns_all_mixed()

    for {
        fmt.Println( "run_worker_cache()", "Updating txid..." )
        g_localdb.Db_memory_update_txns_all_mixed()

        fmt.Println( "run_worker_cache()", "Updating blocks..." )
        g_localdb.Db_memory_update_blocks_info()

        fmt.Println( "run_worker_cache()", "Updating balances..." )
        g_localdb.Db_memory_update_balances()

        fmt.Println()

        //time.Sleep( time.Second * time.Duration(UPDATES_INTERVAL) )
        time.Sleep( time.Millisecond * 1000 * time.Duration(UPDATES_INTERVAL) )
    } // for ()



    /*
    var is_done = false
    for {
        select {
        case <-gChan:
            switch <-gChan {
            case _E_CHAN__CANCEL:
                fmt.Println( "CHAN: CANCEL:" )
                is_done = true
                break
            case _E_CHAN__DONE:
                fmt.Println( "CHAN: DONE:" )
                is_done = true
                break
            }

            if  is_done == true {
                break
            }
        default:
            fmt.Println( "CHAN: Waiting..." )
        }

        if  is_done == true {
            break
        }

    }
    */

    fmt.Println( "run_worker_cache()", "finished..." )

    gWG.Done()
}

// Goroutine
func run_http_rpc_server() {
    if g_localdb == nil {
        panic( "DB object == NULL" )
    }

    fmt.Println( "Starting HTTP RPC Server..." )


    // SEE: var g_localdb = new( rpc_server.LocalDB )

    rpc.Register( g_localdb )
    rpc.HandleHTTP()

    l, e := net.Listen( "tcp", HTTP_RPC_SERVER_HOST_PORT )
    if e != nil {
        log.Fatal( "listen error:", e )
    }

    //go http.Serve( l, nil )
    http.Serve( l, nil )
}

// Goroutine
func run_http_jsonrpc_server() {
    if g_localdb == nil {
        panic( "DB object == NULL" )
    }

    fmt.Println( "Starting HTTP JSON-RPC Server..." )


    // SEE: var g_localdb = new( rpc_server.LocalDB )

    /*
    rpc.Register( g_localdb )
    rpc.HandleHTTP()

    l, e := net.Listen( "tcp", ":1234" )
    if e != nil {
        log.Fatal( "listen error:", e )
    }

    //go http.Serve( l, nil )
    http.Serve( l, nil )
    */

    _rpc := gorilla_rpc.NewServer()
    _rpc.RegisterCodec( gorilla_json.NewCodec(), "application/json" )
    _rpc.RegisterCodec( gorilla_json.NewCodec(), "application/json;charset==UTF-8" )
    _rpc.RegisterService( g_localdb, "" )
    _router := gorilla_mux.NewRouter()
    _router.Handle( "/rpc",  _rpc )

    http.ListenAndServe( HTTP_JSONRPC_SERVER_HOST_PORT,  _router )
}

// Goroutine
func run_txns_fetcher_main() {
    rpc_server.Txns_fetcher_main_func()
}

// Goroutine
func run_updates_balances_all_to_queue__main() {
    rpc_server.Updates_balances_all_to_queue__main_func()
}

// Goroutine
func run_updates_balances_all_to_queue_worker__main() {
    rpc_server.Updates_balances_all_to_queue_worker__main_func()
}

func run_init_db() {
    rpc_server.Init_db_func()
}

func run_release_db() {
    rpc_server.Release_db_func()
}



// --------------------------------------------------------------------



func main() {
    fmt.Println( "HOST: " + URL )

    // Initialize Database
    if __init_db() != true {
        fmt.Println( "main(): DB init...", "false" )
        return
    }

    // Release Database
    if gDB != nil {
        defer gDB.Close()
    }
    if gDBMemory != nil {
        defer gDBMemory.Close()
    }

    //__test_db()


    //! FIXME: logs
    // -------------------------------------
    // Initialize Database instance for {transactions fetcher, updates balances, ...}
    run_init_db()

    var GOROUTINE_TOTAL = 4
    gWG.Add( GOROUTINE_TOTAL )
    //ctx, cancel := context.WithCancel( context.Background() )
    //ctx = context.WithValue( ctx, "key", "val" )

    // Start Data caching
    go run_worker_cache()

    // Start HTTP RPC Server
    go run_http_rpc_server()

    // Start HTTP JSON-RPC Server
    go run_http_jsonrpc_server()

    // Start Transactions fetcher
    go run_txns_fetcher_main()

    // Start updates balances
    go run_updates_balances_all_to_queue__main()
    go run_updates_balances_all_to_queue_worker__main()



    gWG.Wait()



    // Release Database instance for {transactions fetcher, updates balances, ...}
    run_release_db()
}
