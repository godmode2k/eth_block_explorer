/* --------------------------------------------------------------
Project:    Ethereum fetch all transactions
Purpose:
Author:     Ho-Jung Kim (godmode2k@hotmail.com)
Date:       Since Dec 4, 2020
Filename:   rpc_client_main.go

Last modified:  June 7, 2022
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
 -

1. Build:
	$ go build rpc_client_main.go
    or
	$ go run rpc_client_main.go
-------------------------------------------------------------- */
package main



//! Header
// ---------------------------------------------------------------

import (
    "fmt"
    "log"

    "net/rpc"
    "net/http"

    // HTTP JSON-RPC
    "bytes"
    gorilla_json "github.com/gorilla/rpc/json"

    // eth_txns_fetcher
    "eth_block_explorer/types"
)



//! Definition
// --------------------------------------------------------------------

//var SERVER_ADDRESS = "127.0.0.1"
//var SERVER_PORT = "8544"
//var SERVER = SERVER_ADDRESS + ":" + SERVER_PORT
//var URL = "http://" + SERVER_ADDRESS + ":" + SERVER_PORT
var HTTP_RPC_SERVER_HOST_PORT = ":1234" // Internal
var HTTP_JSONRPC_SERVER_HOST_PORT = ":1235" // External



//! Implementation
// --------------------------------------------------------------------

func json_rpc_request(api string, args *types.RPC_Args_st, url string) string {
    var result_str string
    message, err := gorilla_json.EncodeClientRequest( api, args )

    if err != nil {
        log.Fatalf("%s", err)
    }
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(message))
    if err != nil {
        log.Fatalf("%s", err)
    }
    req.Header.Set("Content-Type", "application/json")
    client_jsonrpc := new(http.Client)
    resp, err := client_jsonrpc.Do(req)
    if err != nil {
        log.Fatalf("http.Client.Do(): Error: URL = %s, %s", url, err)
    }
    defer resp.Body.Close()

    err = gorilla_json.DecodeClientResponse(resp.Body, &result_str)
    if err != nil {
        log.Fatalf("DecodeClientResponse(): Error: %s", err)
    }

    return result_str
}

func main() {
    // HTTP RPC Server
    client, err := rpc.DialHTTP("tcp", HTTP_RPC_SERVER_HOST_PORT)
    if err != nil {
        log.Fatal("dialing:", err)
    }

    /*
    // Synchronous call
    args := &server.Args{7,8}
    var reply int
    err = client.Call("Arith.Multiply", args, &reply)
    if err != nil {
        log.Fatal("arith error:", err)
    }
    fmt.Printf("Arith: %d*%d=%d", args.A, args.B, reply)


    // Asynchronous call
    quotient := new(Quotient)
    divCall := client.Go("Arith.Divide", args, quotient, nil)
    replyCall := <-divCall.Done	// will be equal to divCall
    // check errors, print, etc.
    */



    // Synchronous call
    var result []types.Fetch_transactions_st
    var result_str string

    //query := "select id, name from foo limit 1"
    OFFSET := uint(0)

    //err = client.Call( "LocalDB.Db_select_txns_erc1155", OFFSET, &result )

    //! FIXME: Merge DB table (txid and txid_erc1155)
    err = client.Call( "LocalDB.Db_select_txns_all_mixed", OFFSET, &result )
    //err = client.Call( "LocalDB.Db_select_txns_all_erc", OFFSET, &result )
    //err = client.Call( "LocalDB.Db_select_txns_erc20", OFFSET, &result )
    //err = client.Call( "LocalDB.Db_select_txns_erc1155", OFFSET, &result )

    if err != nil {
        log.Fatal( "RPC: error: ", err )
    }
    //fmt.Println( "RPC: result = \n", result )

    for i := 0; i < len(result); i++ {
        fmt.Printf( "[%d] = %s, \n", i, result[i] )
    }
    fmt.Println( "\n\n" )



    //err = client.Call( "LocalDB.Db_memory_select_txns_all_mixed", OFFSET, &result_str )
    err = client.Call( "LocalDB.Db_memory_select_txns_all_mixed", OFFSET, &result )
    if err != nil {
        log.Fatal( "RPC: error: ", err )
    }
    //fmt.Println( "result = \n", result_str )
    for i := 0; i < len(result); i++ {
        fmt.Printf( "[%d] = %s, \n", i, result[i] )
    }
    fmt.Println( "\n\n" )


    //err = client.Call( "LocalDB.Rpc_test", uint(0), &result_str )
    //if err != nil {
    //    log.Fatal( "RPC: error: ", err )
    //}
    //fmt.Println( "result = \n", result_str )



    // ------------------------------------------



    // HTTP JSON-RPC Server
    url := "http://localhost:1235/rpc"


    api := "LocalDB.JSONRPC_get_txns_all_mixed"
    fmt.Println( api )
    args := &types.RPC_Args_st { Req_page: 0 }
    result_str = json_rpc_request( api, args, url )
    fmt.Println( "result = \n", result_str )
    fmt.Println( "\n\n" )



    //args := &types.RPC_DummyArgs_st { Dummy: 10, }
    //message, err := gorilla_json.EncodeClientRequest( "LocalDB.JSONRPC_get_txns_all_mixed", args )

    /*
    args := &types.RPC_Args_st { Req_page: 0, Req_block: "1324" }
    message, err := gorilla_json.EncodeClientRequest( "LocalDB.JSONRPC_get_txns_all_mixed_by_block_number", args )

    if err != nil {
        log.Fatalf("%s", err)
    }
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(message))
    if err != nil {
        log.Fatalf("%s", err)
    }
    req.Header.Set("Content-Type", "application/json")
    client_jsonrpc := new(http.Client)
    resp, err := client_jsonrpc.Do(req)
    if err != nil {
        log.Fatalf("http.Client.Do(): Error: URL = %s, %s", url, err)
    }
    defer resp.Body.Close()

    err = gorilla_json.DecodeClientResponse(resp.Body, &result_str)
    if err != nil {
        log.Fatalf("DecodeClientResponse(): Error: %s", err)
    }
    fmt.Println( "result = \n", result_str )
    */

    api = "LocalDB.JSONRPC_get_txns_all_mixed"
    fmt.Println( api )
    args = &types.RPC_Args_st { Req_page: 0, Req_block: "1324" }
    result_str = json_rpc_request( api, args, url )
    fmt.Println( "result = \n", result_str )
    fmt.Println( "\n\n" )


    api = "LocalDB.JSONRPC_get_blocks_info"
    fmt.Println( api )
    args = &types.RPC_Args_st { Req_page: 0 }
    result_str = json_rpc_request( api, args, url )
    fmt.Println( "result = \n", result_str )
    fmt.Println( "\n\n" )

    api = "LocalDB.JSONRPC_get_block_by_number"
    fmt.Println( api )
    args = &types.RPC_Args_st { Req_block: "1345" }
    result_str = json_rpc_request( api, args, url )
    fmt.Println( "result = \n", result_str )
    fmt.Println( "\n\n" )

    api = "LocalDB.JSONRPC_get_latest_block_number"
    fmt.Println( api )
    args = &types.RPC_Args_st { Req_page: 0 }
    result_str = json_rpc_request( api, args, url )
    fmt.Println( "result = \n", result_str )
    fmt.Println( "\n\n" )

    api = "LocalDB.JSONRPC_get_txns_all_mixed_by_txid"
    fmt.Println( api )
    args = &types.RPC_Args_st { Req_txid: "0x516ef91be8d560fcb6d2bab8a0f1eab8efdb2a8d7ccfb0159a47b3985d4f13e6" }
    result_str = json_rpc_request( api, args, url )
    fmt.Println( "result = \n", result_str )
    fmt.Println( "\n\n" )

    api = "LocalDB.JSONRPC_get_txns_all_mixed_by_address"
    fmt.Println( api )
    args = &types.RPC_Args_st { Req_address: "0xe6e55eed00218faef27eed24def9208f3878b333" }
    result_str = json_rpc_request( api, args, url )
    fmt.Println( "result = \n", result_str )
    fmt.Println( "\n\n" )

    api = "LocalDB.JSONRPC_get_balances_by_address"
    fmt.Println( api )
    args = &types.RPC_Args_st { Req_address: "0xe6e55eed00218faef27eed24def9208f3878b333" }
    result_str = json_rpc_request( api, args, url )
    fmt.Println( "result = \n", result_str )
    fmt.Println( "\n\n" )




    api = "LocalDB.JSONRPC_search_all"
    fmt.Println( api )
    // address, contract address
    //args = &types.RPC_Args_st { Req_page: 0, Req_search: "0xe6e55eed00218faef27eed24def9208f3878b333" }
    // txid
    args = &types.RPC_Args_st { Req_page: 0, Req_search: "0xed7b9b146f3c9a28c2cb882097c8dd5177754216da2a92cf0793a984e60105b7" }
    // block number
    //args = &types.RPC_Args_st { Req_page: 0, Req_search: "1345" }
    result_str = json_rpc_request( api, args, url )
    fmt.Println( "result = \n", result_str )
    fmt.Println( "\n\n" )

}



