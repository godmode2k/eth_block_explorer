/* --------------------------------------------------------------
Project:    Ethereum auto-transfer (accounts to specific address(hotwallet))
Purpose:
Author:     Ho-Jung Kim (godmode2k@hotmail.com)
Date:       Since Dec 4, 2020
Filename:   rpc_server_json_apis.go

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
    "log"
    //"math/big"
    "encoding/json"

    //"runtime"
    "regexp"

    // HTTP JSON-RPC Server
    "net/http"
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



// HTTP JSON-RPC Server: for frontend

func (t *LocalDB) JSONRPC_test(request *http.Request, rpc_args *types.RPC_DummyArgs_st, response *string) error {
    *response = fmt.Sprintf( "JSONRPC_test(): %d", rpc_args.Dummy )

    return nil
}

func (t *LocalDB) JSONRPC_get_txns_all_mixed(request *http.Request, rpc_args *types.RPC_DummyArgs_st, response *string) error {
    log.Println( "JSONRPC_get_txns_all_mixed()" )

    var _result []types.Fetch_transactions_st
    //var result_str string
    OFFSET := uint(0)

    //err := t.Db_memory_select_txns_all_mixed( OFFSET, &result_str )
    err := t.Db_memory_select_txns_all_mixed( OFFSET, &_result )
    if err != nil {
        log.Fatal( "JSONRPC_get_txns_all_mixed(): Error: ", err )
    }

    //*response = result_str


    result, err_marshal := json.Marshal( _result )
    if err_marshal != nil {
        panic( err_marshal.Error() )
    }


    //var json_map_arr = make( map[string]interface{} )
    //json_map_arr["txid"] = _result
    //
    //result, err_marshal := json.Marshal( json_map_arr )
    //if err_marshal != nil {
    //    panic( err_marshal.Error() )
    //}


    *response = string(result)


    return nil
}

func (t *LocalDB) JSONRPC_get_latest_block_number(request *http.Request, rpc_args *types.RPC_DummyArgs_st, response *string) error {
    log.Println( "JSONRPC_get_latest_block_number()" )

    var _result []types.Fetch_blocks_info_st
    //var result_str string
    OFFSET := uint(0)

    //err := t.Db_memory_select_blocks_info( OFFSET, &result_str )
    err := t.Db_memory_select_blocks_info( OFFSET, &_result )
    if err != nil {
        log.Fatal( "JSONRPC_get_latest_block_number(): Error: ", err )
    }

    //*response = result_str


    if len(_result) <= 0 {
        return nil
    }


    // first row (latest block; DESC already)
    result, err_marshal := json.Marshal( _result[0] )
    if err_marshal != nil {
        panic( err_marshal.Error() )
    }


    //var json_map_arr = make( map[string]interface{} )
    //json_map_arr["blocks_info"] = _result
    //
    //result, err_marshal := json.Marshal( json_map_arr )
    //if err_marshal != nil {
    //    panic( err_marshal.Error() )
    //}


    *response = string(result)


    return nil
}

func (t *LocalDB) JSONRPC_get_blocks_info(request *http.Request, rpc_args *types.RPC_DummyArgs_st, response *string) error {
    log.Println( "JSONRPC_get_blocks_info()" )

    var _result []types.Fetch_blocks_info_st
    //var result_str string
    OFFSET := uint(0)

    //err := t.Db_memory_select_blocks_info( OFFSET, &result_str )
    err := t.Db_memory_select_blocks_info( OFFSET, &_result )
    if err != nil {
        log.Fatal( "JSONRPC_get_blocks_info(): Error: ", err )
    }

    //*response = result_str


    result, err_marshal := json.Marshal( _result )
    if err_marshal != nil {
        panic( err_marshal.Error() )
    }


    //var json_map_arr = make( map[string]interface{} )
    //json_map_arr["blocks_info"] = _result
    //
    //result, err_marshal := json.Marshal( json_map_arr )
    //if err_marshal != nil {
    //    panic( err_marshal.Error() )
    //}


    *response = string(result)


    return nil
}

func (t *LocalDB) JSONRPC_get_block_by_number(request *http.Request, rpc_args *types.RPC_Args_st, response *string) error {
    log.Println( "JSONRPC_get_block_by_number()" )

    var _result []types.Fetch_blocks_info_st
    //var result_str string
    //OFFSET := uint( rpc_args.Req_page )
    BLOCK := rpc_args.Req_block

    err := t.Db_select_block_by_number( BLOCK, &_result )
    if err != nil {
        log.Fatal( "JSONRPC_get_block_by_number(): Error: ", err )
    }

    //*response = result_str


    result, err_marshal := json.Marshal( _result )
    if err_marshal != nil {
        panic( err_marshal.Error() )
    }


    //var json_map_arr = make( map[string]interface{} )
    //json_map_arr["blocks_info"] = _result
    //
    //result, err_marshal := json.Marshal( json_map_arr )
    //if err_marshal != nil {
    //    panic( err_marshal.Error() )
    //}


    *response = string(result)


    return nil
}

func (t *LocalDB) JSONRPC_get_block_number_by_block_hash(request *http.Request, rpc_args *types.RPC_Args_st, response *string) error {
    log.Println( "JSONRPC_get_block_number_by_block_hash()" )

    var _result []types.Fetch_blocks_info_st
    //var result_str string
    //OFFSET := uint( rpc_args.Req_page )
    BLOCK_HASH := rpc_args.Req_block_hash

    err := t.Db_select_block_number_by_block_hash( BLOCK_HASH, &_result )
    if err != nil {
        log.Fatal( "JSONRPC_get_block_number__by_block_hash(): Error: ", err )
    }

    //*response = result_str


    result, err_marshal := json.Marshal( _result )
    if err_marshal != nil {
        panic( err_marshal.Error() )
    }


    //var json_map_arr = make( map[string]interface{} )
    //json_map_arr["blocks_info"] = _result
    //
    //result, err_marshal := json.Marshal( json_map_arr )
    //if err_marshal != nil {
    //    panic( err_marshal.Error() )
    //}


    *response = string(result)


    return nil
}

func (t *LocalDB) JSONRPC_get_txns_all_mixed_by_block_number(request *http.Request, rpc_args *types.RPC_Args_st, response *string) error {
    log.Println( "JSONRPC_get_txns_all_mixed_by_block_number()" )

    var _result []types.Fetch_transactions_st
    //var result_str string
    OFFSET := uint( rpc_args.Req_page )
    BLOCK := rpc_args.Req_block

    err := t.Db_select_txns_all_mixed_by_block_number( OFFSET, BLOCK, &_result )
    if err != nil {
        log.Fatal( "JSONRPC_get_txns_all_mixed_by_block_number(): Error: ", err )
    }


    result, err_marshal := json.Marshal( _result )
    if err_marshal != nil {
        panic( err_marshal.Error() )
    }


    //var json_map_arr = make( map[string]interface{} )
    //json_map_arr["txid"] = _result
    //
    //result, err_marshal := json.Marshal( json_map_arr )
    //if err_marshal != nil {
    //    panic( err_marshal.Error() )
    //}


    *response = string(result)

    return nil
}

func (t *LocalDB) JSONRPC_get_txns_all_mixed_by_block_hash(request *http.Request, rpc_args *types.RPC_Args_st, response *string) error {
    log.Println( "JSONRPC_get_txns_all_mixed_by_block_hash()" )

    var _result []types.Fetch_transactions_st
    //var result_str string
    OFFSET := uint( rpc_args.Req_page )
    BLOCK_HASH := rpc_args.Req_block_hash

    err := t.Db_select_txns_all_mixed_by_block_hash( OFFSET, BLOCK_HASH, &_result )
    if err != nil {
        log.Fatal( "JSONRPC_get_txns_all_mixed_by_block_hash(): Error: ", err )
    }


    result, err_marshal := json.Marshal( _result )
    if err_marshal != nil {
        panic( err_marshal.Error() )
    }


    //var json_map_arr = make( map[string]interface{} )
    //json_map_arr["txid"] = _result
    //
    //result, err_marshal := json.Marshal( json_map_arr )
    //if err_marshal != nil {
    //    panic( err_marshal.Error() )
    //}


    *response = string(result)

    return nil
}

func (t *LocalDB) JSONRPC_get_txns_all_mixed_by_txid(request *http.Request, rpc_args *types.RPC_Args_st, response *string) error {
    log.Println( "JSONRPC_get_txns_all_mixed_by_txid()" )

    var _result []types.Fetch_transactions_st
    //var result_str string
    OFFSET := uint( rpc_args.Req_page )
    TXID := rpc_args.Req_txid

    err := t.Db_select_txns_all_mixed_by_txid( OFFSET, TXID, &_result )
    if err != nil {
        log.Fatal( "JSONRPC_get_txns_all_mixed_by_txid(): Error: ", err )
    }


    result, err_marshal := json.Marshal( _result )
    if err_marshal != nil {
        panic( err_marshal.Error() )
    }


    //var json_map_arr = make( map[string]interface{} )
    //json_map_arr["txid"] = _result
    //
    //result, err_marshal := json.Marshal( json_map_arr )
    //if err_marshal != nil {
    //    panic( err_marshal.Error() )
    //}


    *response = string(result)

    return nil
}

func (t *LocalDB) JSONRPC_get_txns_all_mixed_by_address(request *http.Request, rpc_args *types.RPC_Args_st, response *string) error {
    log.Println( "JSONRPC_get_txns_all_mixed_by_address()" )

    var _result []types.Fetch_transactions_st
    //var result_str string
    OFFSET := uint( rpc_args.Req_page )
    ADDRESS := rpc_args.Req_address

    err := t.Db_select_txns_all_mixed_by_address( OFFSET, ADDRESS, &_result )
    if err != nil {
        log.Fatal( "JSONRPC_get_txns_all_mixed_by_address(): Error: ", err )
    }


    result, err_marshal := json.Marshal( _result )
    if err_marshal != nil {
        panic( err_marshal.Error() )
    }


    //var json_map_arr = make( map[string]interface{} )
    //json_map_arr["txid"] = _result
    //  
    //result, err_marshal := json.Marshal( json_map_arr )
    //if err_marshal != nil {
    //    panic( err_marshal.Error() )
    //}


    *response = string(result)

    return nil
}

//! TODO
func (t *LocalDB) JSONRPC_get_balances_by_address(request *http.Request, rpc_args *types.RPC_Args_st, response *string) error {
    log.Println( "JSONRPC_get_balances_by_address()" )

    var _result []types.Fetch_balances_by_address_st
    //var result_str string
    OFFSET := uint( rpc_args.Req_page )
    ADDRESS := rpc_args.Req_address

    err := t.Db_select_balances_by_address( OFFSET, ADDRESS, &_result )
    if err != nil {
        log.Fatal( "JSONRPC_get_balances_by_address(): Error: ", err )
    }


    result, err_marshal := json.Marshal( _result )
    if err_marshal != nil {
        panic( err_marshal.Error() )
    }


    //var json_map_arr = make( map[string]interface{} )
    //json_map_arr["txid"] = _result
    //
    //result, err_marshal := json.Marshal( json_map_arr )
    //if err_marshal != nil {
    //    panic( err_marshal.Error() )
    //}


    *response = string(result)

    return nil
}

func (t *LocalDB) JSONRPC_search_all(request *http.Request, rpc_args *types.RPC_Args_st, response *string) error {
    log.Println( "JSONRPC_search_all()" )

    var _result []types.Fetch_transactions_st
    var err error
    //var result_str string
    OFFSET := uint( rpc_args.Req_page )
    SEARCH := rpc_args.Req_search
    IS_DIGIT := regexp.MustCompile(`^[0-9]+$`)

    if SEARCH[0] == '0' && SEARCH[1] == 'x' {
        if len(SEARCH) == (2+40) {
            // address, contract address
            fmt.Println( "JSONRPC_search_all(): Search: address, contract address" )

            err = t.Db_select_txns_all_mixed_by_address( OFFSET, SEARCH, &_result )
        } else if len(SEARCH) == (2+64) {
            // txid
            fmt.Println( "JSONRPC_search_all(): Search: txid" )

            err = t.Db_select_txns_all_mixed_by_txid( OFFSET, SEARCH, &_result )
        } else {
            // error
            fmt.Println( "JSONRPC_search_all(): Search: unknown" )
            return nil
        }
    } else if IS_DIGIT.MatchString(SEARCH) {
        // block number
        fmt.Println( "JSONRPC_search_all(): Search: block number" )

        err = t.Db_select_txns_all_mixed_by_block_number( OFFSET, SEARCH, &_result )
    } else {
        // error
        fmt.Println( "JSONRPC_search_all(): Search: unknown" )
        return nil
    }



    if err != nil {
        fmt.Println( "JSONRPC_search_all(): Error: ", err )
        return nil
    }

    result, err_marshal := json.Marshal( _result )
    if err_marshal != nil {
        //panic( err_marshal.Error() )
        log.Fatal( "JSONRPC_search_all(): Error: ", err_marshal.Error() )
        return nil
    }


    *response = string(result)

    return nil
}



