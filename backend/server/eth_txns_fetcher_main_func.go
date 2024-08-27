/* --------------------------------------------------------------
Project:    Ethereum fetch all transactions
Purpose:
Author:     Ho-Jung Kim (godmode2k@hotmail.com)
Date:       Since Dec 4, 2020
Filename:   eth_txns_fetcher_main_func.go

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
 - https://pkg.go.dev/database/sql

Dependencies:
$ go get -u github.com/go-sql-driver/mysql


1. Build:
	$ go build eth_txns_fetcher_main_func.go
    or
	$ go run eth_txns_fetcher_main_func.go
-------------------------------------------------------------- */
package rpc_server



//! Header
// ---------------------------------------------------------------

import (
    "fmt"
    "log"
    "bytes"
    "strconv"
    "math"
    "math/big"
    "encoding/hex"
    "strings"
    "time"

    // test
    //"math/rand"

    "net/http"
    "io/ioutil"
    "encoding/json"

    // $ go get -u github.com/go-sql-driver/mysql
    "database/sql"
    _ "github.com/go-sql-driver/mysql"

    // eth_txns_fetcher
    "eth_block_explorer/types"

    "reflect"
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
//var _E_TYPE__ETH = uint8(0)
//var _E_TYPE__ERC20 = uint8(1)
//var _E_TYPE__ERC1155 = uint8(2)

var INTERVAL = int(1) // seconds



//! Implementation
// --------------------------------------------------------------------

func __init_db() {
    fmt.Println( "fetcher_main: __init_db(): initialize..." )

    /*
    db, err := sql.Open( "mysql", DB_LOGIN_USERNAME + ":" + DB_LOGIN_PASSWORD + "@tcp(" + DB_SERVER_ADDRESS + ")/" + DB_NAME )

    if err != nil {
        panic( err.Error() )
    }

    gDB = db
    */


    if gDB != nil {
        // ...
    } else {
        db, err := sql.Open( "mysql", DB_LOGIN_USERNAME + ":" + DB_LOGIN_PASSWORD + "@tcp(" + DB_SERVER_ADDRESS + ")/" + DB_NAME )

        if err != nil {
            panic( err.Error() )
        }

        gDB = db


        //if gDB != nil {
        //    defer gDB.Close()
        //}
    }
}

func __release_db() {
    fmt.Println( "fetcher_main: __release_db(): destroy..." )

    if gDB != nil {
        gDB.Close()
    }
}

/*
func __test_db() {
    var TABLE_NAME = "txid"
    //var TABLE_NAME_ERC1155 = "txid_erc1155"

    var query_str = fmt.Sprintf(
        "INSERT INTO %s VALUES (" +
        "'%s', '%s', '%s', '%s', '%s'," +
        "'%s', '%s', '%s', '%s', '%s'," +
        "'%s', '%s', '%s', '%s', '%s'," +
        "'%s', '%s'" +
        ")",
        TABLE_NAME,
        //0, "eth", "from address", "to address", 1, "amount wei", "amount eth",
        "0", "eth", "from_address", "to_address", "1", "0", "0",
        //"token type", "token symbol", "token decimals", "token contract address",
        "token_type", "token_symbol", "0", "token_contract_address",
        //"token amount wei", "token amount eth", "timestamp", "datetime", 1000, "txid"
        "0", "0", "timestamp", "datetime", "1000", "txid1" )

    fmt.Println( "query = ", query_str )

    //result, err := gDB.Query( "INSERT INTO test VALUES ( 2, 'TEST' )" )
    result, err := gDB.Query( query_str )

    if err != nil {
        panic( err.Error() )
    }
    //defer result.Close()
    result.Close()

    fmt.Println( "result = ", result )
}
*/

// type: (0) ETH, (1) ERC-20, (2) ERC-1155
func db_insert_txns(_type uint8, _data *types.Fetch_transactions_st) {
    TAG := "db_insert_txns(): "

    var TABLE_NAME = ""
    //var TABLE_NAME_ERC1155 = "txid_erc1155"

    var query_str = ""
    var query_str_update = ""

    if reflect.ValueOf(_data).IsNil() {
        fmt.Println( TAG, "data == NULL" )
        return
    }

    /*
    //! test: random 40 bytes hex
    //fmt.Println( "test: ", fmt.Sprintf("0x%s", strconv.FormatUint(rand.Uint64(), 16)) )
    var test_txid = "0x"
    for i := 0; i < 4; i++ {
        test_txid += fmt.Sprintf("%s", strconv.FormatUint(rand.Uint64(), 16) )
    }
    fmt.Println( "test: txid = ", test_txid )
    _data.Txid = test_txid
    */


    /*
    // ETH
    if _type == 0 {
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
            _data.Is_send,
            _data.Amount_wei,
            _data.Amount_eth,
            _data.Token_type,
            _data.Token_symbol,
            _data.Token_decimals,
            _data.Token_total_supply,
            _data.Token_contract_address,
            _data.Token_amount_wei,
            _data.Token_amount_eth,
            // for SQL: txid table
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
                _data.Is_send,
                _data.Amount_wei,
                _data.Amount_eth,
                _data.Token_type,
                _data.Token_symbol,
                _data.Token_decimals,
                _data.Token_total_supply,
                _data.Token_contract_address,
                _data.Token_amount_wei,
                _data.Token_amount_eth,
                // for SQL: txid table
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
                _data.Is_send,
                // for SQL: txid_erc1155 table
                //_data.Amount_wei,
                //_data.Amount_eth,
                _data.Token_type,
                _data.Token_symbol,
                // for SQL: txid_erc1155 table
                _data.Token_decimals,
                //_data.Token_total_supply,
                _data.Token_contract_address,
                // for SQL: txid_erc1155 table
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
    */

    {
        // ALL: ETHER, ERC-20, ERC-1155
        TABLE_NAME = "txid"
        query_str = fmt.Sprintf(
            "INSERT INTO %s VALUES (0," +
            "'%s', '%s', '%s'," +
            "'%s', '%s', '%s', '%s', '%s'," +
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
            _data.Token_amount,
            _data.Token_uri_ascii,
            _data.Token_uri_hexadecimal,
            _data.Token_data_length,
            _data.Token_data,
            _data.Timestamp,
            _data.Datetime,
            _data.Block_number,
            _data.Txid )


        query_str_update = fmt.Sprintf(
            "UPDATE %s SET " +
            "from_address='%s', to_address='%s', ether_amount_wei='%s'," +
            "ether_amount_eth='%s', token_type='%s', token_symbol='%s', token_decimals='%s', token_total_supply='%s'," +
            "token_contract_address='%s', token_amount_wei='%s', token_amount_eth='%s', erc1155_token_amount='%s', erc1155_token_uri_ascii='%s'," +
            "erc1155_token_uri_hexadecimal='%s', erc1155_token_data_length='%s', erc1155_token_data='%s', timestamp='%s', datetime='%s'," +
            "blocks='%s', txid='%s'" +
            " WHERE txid='%s'" +
            "",

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
            _data.Token_amount,
            _data.Token_uri_ascii,
            _data.Token_uri_hexadecimal,
            _data.Token_data_length,
            _data.Token_data,
            _data.Timestamp,
            _data.Datetime,
            _data.Block_number,
            _data.Txid,
            _data.Txid )
    }

    fmt.Println( "query = ", query_str )

    //result, err := gDB.Query( "INSERT INTO test VALUES ( 2, 'TEST' )" )
    result, err := gDB.Query( query_str )

    if err != nil {
        //! FIXME: Handling Error codes
        if strings.Contains( err.Error(), "Duplicate entry" ) {
            fmt.Println( "Error: Duplicate entry" )

            // update(replace) it...
            fmt.Println( "update(replace) query = ", query_str_update )
            result, err = gDB.Query( query_str_update )

            if err != nil {
                //panic( err.Error() )
                fmt.Println( err.Error() )
                //return
            }

            //defer result.Close()
            result.Close()
        } else {
            //panic( err.Error() )
            fmt.Println( err.Error() )
            //return
        }

        return
    }

    //defer result.Close()
    result.Close()

    fmt.Println( "result = ", result )
}

func db_insert_blocks_info(block_number string, block_hash string, block_info string, num_of_txns int) {
    var TABLE_NAME = "blocks"

    var query_str = ""
    var query_str_update = ""

    {
        query_str = fmt.Sprintf(
            "INSERT INTO %s VALUES (0," +
            "'%s', '%s', '%s', '%d'" +
            ")",

            TABLE_NAME,

            block_number, block_hash, block_info, num_of_txns )


        query_str_update = fmt.Sprintf(
            "UPDATE %s SET " +
            "blocks='%s', blocks_hash='%s', info='%s', transactions=%d" +
            " WHERE blocks_hash='%s'" +
            "",

            TABLE_NAME,

            block_number, block_hash, block_info, num_of_txns, block_hash )
    }

    fmt.Println( "query = ", query_str )

    result, err := gDB.Query( query_str )

    if err != nil {
        //! FIXME: Handling Error codes
        if strings.Contains( err.Error(), "Duplicate entry" ) {
            fmt.Println( "Error: Duplicate entry" )

            // update(replace) it...
            fmt.Println( "update(replace) query = ", query_str_update )
            result, err = gDB.Query( query_str_update )

            if err != nil {
                //panic( err.Error() )
                fmt.Println( err.Error() )
                return
            }

            //defer result.Close()
            result.Close()
        } else {
            //panic( err.Error() )
            fmt.Println( err.Error() )
            return
        }

        return
    }

    //defer result.Close()
    result.Close()

    fmt.Println( "result = ", result )
}



func eth_get_balance(_address string) {
    // eth: eth_getBalance
    //
    // request:
    // $ curl -X POST --data
    //  '{"jsonrpc":"2.0",
    //  "method":"eth_getBalance",
    //  "params":["0xe6e55eed00218faef27eed24def9208f3878b333", "latest"],"id":0}'
    //  -H "Content-Type: application/json" http://127.0.0.1:8544/

    fmt.Println( "eth_getBalance()" )

    var result types.Result

    var params []interface{}
    //params = append( params, "0xe6e55eed00218faef27eed24def9208f3878b333", "latest" )
    params = append( params, _address, "latest" )
    request_data := types.RequestData { Jsonrpc: "2.0", Method: "eth_getBalance", Params: params, Id: 0 }
    //request_data := types.RequestData { Jsonrpc: "2.0", Method: "eth_getBalance",
    //	Params: []interface{}{"0xe6e55eed00218faef27eed24def9208f3878b333", "latest"}, Id: 0 }

    message, _ := json.Marshal( request_data )
    //fmt.Println( "message: ", request_data )
    response, err := http.Post( URL, "application/json", bytes.NewBuffer(message) )
    defer response.Body.Close()
    if err != nil {
        log.Fatal( "http.Post: ", err )
    }

    //fmt.Println( "response: " )
    responseBody, err := ioutil.ReadAll( response.Body )
    if err != nil {
        log.Fatal( "ioutil.ReadAll: ", err )
    }

    //fmt.Println( string(responseBody) )
    err = json.Unmarshal( responseBody, &result )
    if err != nil {
        log.Fatal( "json.Unmarshal: ", err )
    }
    //fmt.Println( "jsonrpc =" , result.Jsonrpc, ", id =", result.Id, ", result =", result.Result )

    // SEE:
    // - https://golang.org/pkg/math/big/
    // - https://golang.org/pkg/strconv/
    // - https://goethereumbook.org/account-balance/
    balance_wei_int := new(big.Int)
    balance_wei_int.SetString( result.Result[2:], 16 )
    fmt.Println( "ether hex-string to int: ", balance_wei_int, "(wei)" )
    balance_wei_float := new(big.Float)
    balance_wei_float.SetString( balance_wei_int.String() )
    balance_eth := new(big.Float).Quo(balance_wei_float, big.NewFloat(math.Pow10(18)))
    fmt.Println( "ether balance: ", balance_eth, "(ether)" )
}

func eth_send_transaction(_from string, _to string, _amount string, _gas string, _gasprice string) {
    // eth: eth_sendTransaction
    //
    // request:
    // $ curl -X POST --data
    //  '{"jsonrpc":"2.0",
    //  "method":"eth_sendTransaction",
    //  "params":[{
    //      "from": "0xe6e55eed00218faef27eed24def9208f3878b333",
    //      "to": "0x8f5b2b7608e3e3a3dc0426c3396420fbf1849454",
    //      "value": "0x8ca93d72e1ed4000", "gas": "0x11170", "gasPrice": "0x12a05f2000"}],"id":0}'
    //  -H "Content-Type: application/json" http://127.0.0.1:8544/

    fmt.Println( "eth_sendTransaction()" )

    var result types.Result

    from := _from
    to := _to
    gas := _gas
    gasprice := _gasprice
    value := _amount

    //gas := "70000" // 70000
    //gasprice := "100" // 100 * gwei(1e9)
    //value := "11.1357" // 10 * wei(1e18)
    gas_hex := ""
    gasprice_hex := ""
    value_hex := ""
    // ---
    {
        gas_int := new(big.Int)
        gas_float := new(big.Float)
        gasprice_int := new(big.Int)
        gasprice_float := new(big.Float)

        gas_float.SetString( gas )
        gasprice_float.SetString( gasprice )
        gasprice_decimals := big.NewFloat( math.Pow10(9) ) //new(big.Float)( math.Pow10(9) )
        gasprice_float_mul := new(big.Float).Mul( gasprice_float, gasprice_decimals ) // value * decimals(wei: 1e9)

        // float to int for hex
        // SEE: https://stackoverflow.com/questions/47545898/golang-convert-big-float-to-big-int
        gas_float.Int( gas_int )
        gasprice_float_mul.Int( gasprice_int )

        // ---

        value_float := new(big.Float)
        value_float.SetString( value )
        value_decimals := big.NewFloat( math.Pow10(18) ) //new(big.Float)( math.Pow10(18) )
        value_float_mul := new(big.Float).Mul( value_float, value_decimals ) // value * decimals(wei: 1e18)
        // DO NOT USE [
        //value_result := value_float_mul.Text( 'f', 8 ) // precision: 8, no exponent
        //value_result := value_float_mul.Text( 'x', 8 ) // precision: 8, hexadecimal mantissa
        //fmt.Println( "result:", value_result )
        //
        // USE THIS
        // SEE: https://stackoverflow.com/questions/47545898/golang-convert-big-float-to-big-int
        value_int := new(big.Int)
        value_float_mul.Int( value_int ) // float to int for hex
        // ]

        //fmt.Println( "value:" , value, "value_float:", value_float, "value_decimals:", value_decimals )
        //fmt.Printf( "%f\n", value_float_mul )
        //fmt.Printf( "hex = %s\n", hex.EncodeToString([]byte(value_result)) ) // DO NOT USE
        //fmt.Printf( "%s, %s\n", value_int, value_int.Text(16) ) // hex

        // ---

        //gas_hex := "0x" + hex.EncodeToString( []byte(gas) )
        //gasprice_hex := "0x" + hex.EncodeToString( []byte(gasprice) )
        //value_hex := "0x" + hex.EncodeToString( []byte(string(value_result)) )

        gas_hex = "0x" + gas_int.Text( 16 )
        gasprice_hex = "0x" + gasprice_int.Text( 16 )
        value_hex = "0x" + value_int.Text( 16 )
    }
    // ---
    //from := "0xe6e55eed00218faef27eed24def9208f3878b333"
    ////to := "0x1e57f9561600b269a37437f02ce9da31e5b830ce" // erc-20: contract address (abcd token)
    //to := "0x8f5b2b7608e3e3a3dc0426c3396420fbf1849454"
    //holder_address := ""
    method := "eth_sendTransaction"
    //! DO NOT USE [
    // except: cancel pending transaction, ...
    //data := ""
    //nonce := ""
    //data_hex := "0x" + data
    //nonce_hex := "0x" + nonce
    // ]
    var params []interface{}

    request_data_param := types.RequestData_params_transaction {
        From: from, To: to, Value: value_hex, Gas: gas_hex, Gasprice: gasprice_hex,
        //! DO NOT USE [
        // except: cancel pending transaction, ...
        //Data: data_hex, Nonce: nonce_hex
        // ]
    }
    params = append( params, request_data_param )
    request_data := types.RequestData { Jsonrpc: "2.0", Method: method, Params: params, Id: 0 }


    {
        // unlock: personal_unlockAccount
        //
        // request:
        // $ curl -X POST --data
        //  '{"jsonrpc":"2.0",
        //  "method":"personal_unlockAccount",
        //  "params": ["0xe6e55eed00218faef27eed24def9208f3878b333","12345678",5], "id":0}'
        //  -H "Content-Type: application/json" http://127.0.0.1:8544/

        type Result struct {
            Jsonrpc string `json:"jsonrpc"`
            Id int `json:"id"`
            Result bool `json:"result"`
        }
        var result Result


        passphrase := "12345678"
        duration := 5
        var params []interface{}
        params = append( params, from, passphrase, duration )
        request_data := types.RequestData { Jsonrpc: "2.0", Method: "personal_unlockAccount", Params: params, Id: 0 }

        message, _ := json.Marshal( request_data )
        //fmt.Println( "message: ", request_data )

        response, err := http.Post( URL, "application/json", bytes.NewBuffer(message) )
        defer response.Body.Close()
        if err != nil {
            log.Fatal( "http.Post: ", err )
        }

        //fmt.Println( "response: " )
        responseBody, err := ioutil.ReadAll( response.Body )
        if err != nil {
            log.Fatal( "ioutil.ReadAll: ", err )
        }

        //fmt.Println( string(responseBody) )
        err = json.Unmarshal( responseBody, &result )
        if err != nil {
            log.Fatal( "json.Unmarshal: ", err )
        }
        //fmt.Println( "jsonrpc =" , result.Jsonrpc, ", id =", result.Id, ", result =", result.Result )
    }


    message, _ := json.Marshal( request_data )
    //fmt.Println( "message: ", request_data )

    fmt.Printf( "send ether:\nfrom = %s\nto = %s\nvalue = %s, gas = %s, gasPrice = %s\n", from, to, value, gas, gasprice )

    response, err := http.Post( URL, "application/json", bytes.NewBuffer(message) )
    defer response.Body.Close()
    if err != nil {
        log.Fatal( "http.Post: ", err )
    }

    //fmt.Println( "response: " )
    responseBody, err := ioutil.ReadAll( response.Body )
    if err != nil {
        log.Fatal( "ioutil.ReadAll: ", err )
    }

    //fmt.Println( string(responseBody) )
    err = json.Unmarshal( responseBody, &result )
    if err != nil {
        log.Fatal( "json.Unmarshal: ", err )
    }
    //fmt.Println( "jsonrpc =" , result.Jsonrpc, ", id =", result.Id, ", result =", result.Result )
    fmt.Println( "txid: ", result.Result )
}

func erc20_get_balance(_to string, _holder_address string) {
    // eth erc-20: eth_call
    //params := Params_ERC20 { "0xe6e55eed00218faef27eed24def9208f3878b333", "0x70a08231" }


    // eth erc-20: balanceOf(address)
    //
    // request:
    // $ curl -X POST --data
    //  '{"jsonrpc":"2.0",
    //  "method":"eth_call",
    //  "params":[{"to": "0x1e57f9561600b269a37437f02ce9da31e5b830ce", // ABCD token contract address
    //  "data":"0x70a08231000000000000000000000000e6e55eed00218faef27eed24def9208f3878b333"}, "latest"],"id":67}'
    //  -H "Content-Type: application/json" http://127.0.0.1:8544/
    //
    // method name:
    // > web3.sha3('balanceOf(address)')
    // "0x70a08231b98ef4ca268c9cc3f6b4590e4bfec28280db06bb5d45e689f2a360be"
    //
    // data:
    // <method name> + '0 x 24' + <token holder address>
    // <70a08231> 000000000000000000000000 <token holder address>

    fmt.Println( "eth_call(): balanceOf()" )

    var result types.Result

    //gas := "70000"
    //gasprice := "100"
    //value := ""
    //from := ""
    to := _to // erc-20 contract address
    holder_address := _holder_address
    //to := "0x1e57f9561600b269a37437f02ce9da31e5b830ce" // erc-20 contract address
    //holder_address := "0xe6e55eed00218faef27eed24def9208f3878b333"
    method := "0x70a08231"
    data := method + "000000000000000000000000" + holder_address[2:]

    var params []interface{}
    //request_data_param := types.RequestData_params_erc20_transaction { From: from, To: to, Value: value, Gas: gas, Gasprice: gasprice, Data: data }
    request_data_param := types.RequestData_params_erc20 { To: to, Data: data }
    params = append( params, request_data_param, "latest" )
    request_data := types.RequestData { Jsonrpc: "2.0", Method: "eth_call", Params: params, Id: 0 }

    message, _ := json.Marshal( request_data )
    //fmt.Println( "message: ", request_data )
    response, err := http.Post( URL, "application/json", bytes.NewBuffer(message) )
    defer response.Body.Close()
    if err != nil {
        log.Fatal( "http.Post: ", err )
    }

    //fmt.Println( "response: " )
    responseBody, err := ioutil.ReadAll( response.Body )
    if err != nil {
        log.Fatal( "ioutil.ReadAll: ", err )
    }

    //fmt.Println( string(responseBody) )
    err = json.Unmarshal( responseBody, &result )
    if err != nil {
        log.Fatal( "json.Unmarshal: ", err )
    }
    //fmt.Println( "jsonrpc =" , result.Jsonrpc, ", id =", result.Id, ", result =", result.Result )

    // SEE:
    // - https://golang.org/pkg/math/big/
    // - https://golang.org/pkg/strconv/
    // - https://goethereumbook.org/account-balance/
    //balance_wei_int := new(big.Int)
    //balance_wei_int.SetString( result.Result[2:], 16 )
    //fmt.Println( "hex-string to int: ", balance_wei_int, "(wei)" )
    //balance_wei_float := new(big.Float)
    //balance_wei_float.SetString( balance_wei_int.String() )
    //balance_token := new(big.Float).Quo(balance_wei_float, big.NewFloat(math.Pow10(18)))
    //fmt.Printf( "erc-20 token: %f ()\n", balance_token )


    balance_wei := result.Result
    _token_name := ""
    _token_symbol := ""
    _token_decimals := ""
    _token_total_supply := ""

    {
        // Token: name
        method = "0x06fdde03"
        data = method + "000000000000000000000000" + holder_address[2:]

        var params []interface{}
        request_data_param := types.RequestData_params_erc20 { To: to, Data: data }
        params = append( params, request_data_param, "latest" )
        request_data := types.RequestData { Jsonrpc: "2.0", Method: "eth_call", Params: params, Id: 0 }

        message, _ := json.Marshal( request_data )
        //fmt.Println( "message: ", request_data )
        response, err := http.Post( URL, "application/json", bytes.NewBuffer(message) )
        defer response.Body.Close()
        if err != nil {
            log.Fatal( "http.Post: ", err )
        }

        //fmt.Println( "response: " )
        responseBody, err := ioutil.ReadAll( response.Body )
        if err != nil {
            log.Fatal( "ioutil.ReadAll: ", err )
        }

        //fmt.Println( string(responseBody) )
        err = json.Unmarshal( responseBody, &result )
        if err != nil {
            log.Fatal( "json.Unmarshal: ", err )
        }
        //fmt.Println( "jsonrpc =" , result.Jsonrpc, ", id =", result.Id, ", result =", result.Result )
        _token_name = result.Result
    }


    {
        // Token: symbol
        method = "0x95d89b41"
        data = method + "000000000000000000000000" + holder_address[2:]

        var params []interface{}
        request_data_param := types.RequestData_params_erc20 { To: to, Data: data }
        params = append( params, request_data_param, "latest" )
        request_data := types.RequestData { Jsonrpc: "2.0", Method: "eth_call", Params: params, Id: 0 }

        message, _ := json.Marshal( request_data )
        //fmt.Println( "message: ", request_data )
        response, err := http.Post( URL, "application/json", bytes.NewBuffer(message) )
        defer response.Body.Close()
        if err != nil {
            log.Fatal( "http.Post: ", err )
        }

        //fmt.Println( "response: " )
        responseBody, err := ioutil.ReadAll( response.Body )
        if err != nil {
            log.Fatal( "ioutil.ReadAll: ", err )
        }

        //fmt.Println( string(responseBody) )
        err = json.Unmarshal( responseBody, &result )
        if err != nil {
            log.Fatal( "json.Unmarshal: ", err )
        }
        //fmt.Println( "jsonrpc =" , result.Jsonrpc, ", id =", result.Id, ", result =", result.Result )
        _token_symbol = result.Result
    }


    {
        // Token: decimals
        method = "0x313ce567"
        data = method + "000000000000000000000000" + holder_address[2:]

        var params []interface{}
        request_data_param := types.RequestData_params_erc20 { To: to, Data: data }
        params = append( params, request_data_param, "latest" )
        request_data := types.RequestData { Jsonrpc: "2.0", Method: "eth_call", Params: params, Id: 0 }

        message, _ := json.Marshal( request_data )
        //fmt.Println( "message: ", request_data )
        response, err := http.Post( URL, "application/json", bytes.NewBuffer(message) )
        defer response.Body.Close()
        if err != nil {
            log.Fatal( "http.Post: ", err )
        }

        //fmt.Println( "response: " )
        responseBody, err := ioutil.ReadAll( response.Body )
        if err != nil {
            log.Fatal( "ioutil.ReadAll: ", err )
        }

        //fmt.Println( string(responseBody) )
        err = json.Unmarshal( responseBody, &result )
        if err != nil {
            log.Fatal( "json.Unmarshal: ", err )
        }
        //fmt.Println( "jsonrpc =" , result.Jsonrpc, ", id =", result.Id, ", result =", result.Result )
        _token_decimals = result.Result
    }


    {
        // Token: total_supply 
        method = "0x18160ddd"
        data = method + "000000000000000000000000" + holder_address[2:]

        var params []interface{}
        request_data_param := types.RequestData_params_erc20 { To: to, Data: data }
        params = append( params, request_data_param, "latest" )
        request_data := types.RequestData { Jsonrpc: "2.0", Method: "eth_call", Params: params, Id: 0 }

        message, _ := json.Marshal( request_data )
        //fmt.Println( "message: ", request_data )
        response, err := http.Post( URL, "application/json", bytes.NewBuffer(message) )
        defer response.Body.Close()
        if err != nil {
            log.Fatal( "http.Post: ", err )
        }

        //fmt.Println( "response: " )
        responseBody, err := ioutil.ReadAll( response.Body )
        if err != nil {
            log.Fatal( "ioutil.ReadAll: ", err )
        }

        //fmt.Println( string(responseBody) )
        err = json.Unmarshal( responseBody, &result )
        if err != nil {
            log.Fatal( "json.Unmarshal: ", err )
        }
        //fmt.Println( "jsonrpc =" , result.Jsonrpc, ", id =", result.Id, ", result =", result.Result )
        _token_total_supply = result.Result
    }

    //-----{
    // token name: 0x + [60 bytes] + [4 bytes] + [60 bytes] + [4 bytes]:
    token_name, _ := hex.DecodeString( _token_name[2 + 60 + 4 + 60 + 4:] )

    // token symbol: 0x + [60 bytes] + [4 bytes] + [60 bytes] + [4 bytes]:
    token_symbol, _ := hex.DecodeString( _token_symbol[2 + 60 + 4 + 60 + 4:] )

    // token decimals: 0x + [60 bytes] + [4 bytes]
    token_decimals_int := new(big.Int)
    token_decimals_int.SetString( _token_decimals[2:], 16 )
    token_decimals := token_decimals_int.String()
    token_decimals_int32, _ := strconv.Atoi( token_decimals )

    // token total supply:
    token_total_supply_int := new(big.Int)
    token_total_supply_int.SetString( _token_total_supply[2:], 16 )

    //token_total_supply := token_total_supply_int.String()
    token_total_supply_float := new(big.Float)
    token_total_supply_float.SetString( token_total_supply_int.String() )
    token_total_supply := new(big.Float).Quo(token_total_supply_float, big.NewFloat(math.Pow10(token_decimals_int32)))

    fmt.Println( "token name:", string(bytes.Trim(token_name, "\x00")) )
    //fmt.Printf( "token name hex = %s\n", hex.EncodeToString([]byte(token_name)) )
    fmt.Println( "token_symbol:", string(bytes.Trim(token_symbol, "\x00")) )
    fmt.Println( "token_decimals:", token_decimals )
    fmt.Printf( "token_total_supply: %f\n", token_total_supply )
    //-----}


    // SEE:
    // - https://golang.org/pkg/math/big/
    // - https://golang.org/pkg/strconv/
    // - https://goethereumbook.org/account-balance/
    balance_wei_int := new(big.Int)
    balance_wei_int.SetString( balance_wei[2:], 16 )
    fmt.Println( "erc-20 token balance hex-string to int: ", balance_wei_int, "(wei)" )
    balance_wei_float := new(big.Float)
    balance_wei_float.SetString( balance_wei_int.String() )
    balance_token := new(big.Float).Quo(balance_wei_float, big.NewFloat(math.Pow10(18)))
    fmt.Printf( "erc-20 token balance: %.8f (%s)\n", balance_token, string(bytes.Trim(token_symbol, "\x00")) )
}

func erc20_send_transaction(_contract_address string, _from string, _to string, _amount string, _gas string, _gasprice string) {
    // eth erc-20: transfer(address,uint256)
    //
    // request:
    // $ curl -X POST --data
    //  '{"jsonrpc":"2.0",
    //  "method":"eth_sendTransaction",
    //  "params":[{
    //  "from":"0xe6e55eed00218faef27eed24def9208f3878b333",
    //  "to":"0x1e57f9561600b269a37437f02ce9da31e5b830ce",
    //  "gas":"0x11170","gasPrice":"0x174876e800",
    //  "data":"0xa9059cbb0000000000000000000000008f5b2b7608e3e3a3dc0426c3396420fbf18494540000000000000000000000000000000000000000000000000fc2d121ff694000"}],"id":0}'
    //  -H "Content-Type: application/json" http://127.0.0.1:8544/
    //
    // method name:
    // > web3.sha3('transfer(address,uint256)')
    // "0xa9059cbb2ab09eb219583f4a59a5d0623ade346d962bcd4e46b11da047c9049b"
    //
    // data:
    // <method name>           + // 4 bytes
    // '0 x 24' + <to address> + // 32 bytes
    // '0 x X' + <amount>        // 32 bytes
    //
    // <0xa9059cbb> 000000000000000000000000 <to address>
    // <0 x X> + <amount>

    fmt.Println( "eth_call(): transfer()" )

    var result types.Result

    contract_address := _contract_address // contract address
    from := _from
    to := _to
    //contract_address := "0x1e57f9561600b269a37437f02ce9da31e5b830ce" // contract address
    //from := "0xe6e55eed00218faef27eed24def9208f3878b333"
    //to := "0x8f5b2b7608e3e3a3dc0426c3396420fbf1849454"
    //holder_address := ""

    gas := _gas
    gasprice := _gasprice // 100 * gwei(1e9)
    value_amount := _amount // 10 * (erc-20 token decimals)

    value := "0" // for Ether, "0" fixed if ERC-20 transfer()
    //gas := "70000" // 70000
    //gasprice := "100" // 100 * gwei(1e9)
    //value_amount := "1.1" // 10 * (erc-20 token decimals)
    gas_hex := ""
    gasprice_hex := ""
    value_hex := ""
    value_amount_hex := ""
    // ---
    {
        gas_int := new(big.Int)
        gas_float := new(big.Float)
        gasprice_int := new(big.Int)
        gasprice_float := new(big.Float)

        gas_float.SetString( gas )
        gasprice_float.SetString( gasprice )
        gasprice_decimals := big.NewFloat( math.Pow10(9) ) //new(big.Float)( math.Pow10(9) )
        gasprice_float_mul := new(big.Float).Mul( gasprice_float, gasprice_decimals ) // value * decimals(wei: 1e9)

        // float to int for hex
        // SEE: https://stackoverflow.com/questions/47545898/golang-convert-big-float-to-big-int
        gas_float.Int( gas_int )
        gasprice_float_mul.Int( gasprice_int )

        // ---

        var token_decimals_int32 = 0
        {
            // Token: decimals
            method := "0x313ce567"
            data := method + "000000000000000000000000" + from[2:] // holder address

            var params []interface{}
            request_data_param := types.RequestData_params_erc20 { To: contract_address, Data: data }
            params = append( params, request_data_param, "latest" )
            request_data := types.RequestData { Jsonrpc: "2.0", Method: "eth_call", Params: params, Id: 0 }

            message, _ := json.Marshal( request_data )
            //fmt.Println( "message: ", request_data )
            response, err := http.Post( URL, "application/json", bytes.NewBuffer(message) )
            defer response.Body.Close()
            if err != nil {
                log.Fatal( "http.Post: ", err )
            }

            //fmt.Println( "response: " )
            responseBody, err := ioutil.ReadAll( response.Body )
            if err != nil {
                log.Fatal( "ioutil.ReadAll: ", err )
            }

            //fmt.Println( string(responseBody) )
            err = json.Unmarshal( responseBody, &result )
            if err != nil {
                log.Fatal( "json.Unmarshal: ", err )
            }
            //fmt.Println( "jsonrpc =" , result.Jsonrpc, ", id =", result.Id, ", result =", result.Result )
            _token_decimals := result.Result


            // token decimals: 0x + [60 bytes] + [4 bytes]
            token_decimals_int := new(big.Int)
            token_decimals_int.SetString( _token_decimals[2:], 16 )
            token_decimals := token_decimals_int.String()
            token_decimals_int32, _ = strconv.Atoi( token_decimals )
        }

        // ---

        value_amount_float := new(big.Float)
        value_amount_float.SetString( value_amount )
        value_amount_decimals := big.NewFloat( math.Pow10(token_decimals_int32) ) //new(big.Float)( math.Pow10(18) )
        value_amount_float_mul := new(big.Float).Mul( value_amount_float, value_amount_decimals ) // value * decimals(wei: 1e18)
        // DO NOT USE [
        //value_amount_result := value_amount_float_mul.Text( 'f', 8 ) // precision: 8, no exponent
        //value_amount_result := value_amount_float_mul.Text( 'x', 8 ) // precision: 8, hexadecimal mantissa
        //fmt.Println( "result:", value_amount_result )
        //
        // USE THIS
        // SEE: https://stackoverflow.com/questions/47545898/golang-convert-big-float-to-big-int
        value_amount_int := new(big.Int)
        value_amount_float_mul.Int( value_amount_int ) // float to int for hex
        // ]

        //fmt.Println( "value_amount:" , value_amount, "value_amount_float:", value_amount_float, "value_amount_decimals:", value_amount_decimals )
        //fmt.Printf( "%f\n", value_amount_float_mul )
        //fmt.Printf( "hex = %s\n", hex.EncodeToString([]byte(value_amount_result)) ) // DO NOT USE
        //fmt.Printf( "%s, %s\n", value_amount_int, value_amount_int.Text(16) ) // hex

        // ---

        //gas_hex := "0x" + hex.EncodeToString( []byte(gas) )
        //gasprice_hex := "0x" + hex.EncodeToString( []byte(gasprice) )
        //value_hex := "0x" + hex.EncodeToString( []byte(string(value_result)) )

        gas_hex = "0x" + gas_int.Text( 16 )
        gasprice_hex = "0x" + gasprice_int.Text( 16 )
        value_hex = "0x" + value // always '0x0' for erc-20
        value_amount_hex = "0x" + value_amount_int.Text( 16 )
    }
    // ---
    method := "0xa9059cbb"
    data := method + "000000000000000000000000" + to[2:] +
            strings.Repeat("0", 64 - len(value_amount_hex[2:])) + value_amount_hex[2:]

    var params []interface{}
    request_data_param := types.RequestData_params_erc20_transaction {
        From: from, To: contract_address, Value: value_hex, Gas: gas_hex, Gasprice: gasprice_hex,
        Data: data,
    }
    //request_data_param := types.RequestData_params_erc20 { To: to, Data: data }
    params = append( params, request_data_param )
    request_data := types.RequestData { Jsonrpc: "2.0", Method: "eth_sendTransaction", Params: params, Id: 0 }


    {
        // unlock: personal_unlockAccount
        //
        // request:
        // $ curl -X POST --data
        //  '{"jsonrpc":"2.0",
        //  "method":"personal_unlockAccount",
        //  "params": ["0xe6e55eed00218faef27eed24def9208f3878b333","12345678",5], "id":0}'
        //  -H "Content-Type: application/json" http://127.0.0.1:8544/

        type Result struct {
            Jsonrpc string `json:"jsonrpc"`
            Id int `json:"id"`
            Result bool `json:"result"`
        }
        var result Result


        passphrase := "12345678"
        duration := 5
        var params []interface{}
        params = append( params, from, passphrase, duration )
        request_data := types.RequestData { Jsonrpc: "2.0", Method: "personal_unlockAccount", Params: params, Id: 0 }

        message, _ := json.Marshal( request_data )
        //fmt.Println( "message: ", request_data )

        response, err := http.Post( URL, "application/json", bytes.NewBuffer(message) )
        defer response.Body.Close()
        if err != nil {
            log.Fatal( "http.Post: ", err )
        }

        //fmt.Println( "response: " )
        responseBody, err := ioutil.ReadAll( response.Body )
        if err != nil {
            log.Fatal( "ioutil.ReadAll: ", err )
        }

        //fmt.Println( string(responseBody) )
        err = json.Unmarshal( responseBody, &result )
        if err != nil {
            log.Fatal( "json.Unmarshal: ", err )
        }
        //fmt.Println( "jsonrpc =" , result.Jsonrpc, ", id =", result.Id, ", result =", result.Result )
    }


    message, _ := json.Marshal( request_data )
    //fmt.Println( "message: ", request_data )

    response, err := http.Post( URL, "application/json", bytes.NewBuffer(message) )
    defer response.Body.Close()
    if err != nil {
        log.Fatal( "http.Post: ", err )
    }

    //fmt.Println( "response: " )
    responseBody, err := ioutil.ReadAll( response.Body )
    if err != nil {
        log.Fatal( "ioutil.ReadAll: ", err )
    }

    //fmt.Println( string(responseBody) )
    err = json.Unmarshal( responseBody, &result )
    if err != nil {
        log.Fatal( "json.Unmarshal: ", err )
    }
    //fmt.Println( "jsonrpc =" , result.Jsonrpc, ", id =", result.Id, ", result =", result.Result )
    fmt.Println( "txid: ", result.Result )
}

func erc1155_get_balance(_to string, _holder_address string, _token_id string) {
    // eth erc-1155: eth_call

    // https://docs.openzeppelin.com/contracts/3.x/api/token/erc1155#IERC1155MetadataURI
    // https://eth.wiki/json-rpc/API
    // https://docs.soliditylang.org/en/latest/abi-spec.html
    // eth erc-1155: balanceOf(address,uint256)
    //
    // request:
    // $ curl -X POST --data
    //  '{"jsonrpc":"2.0",
    //  "method":"eth_call",
    //  "params":[{"to": "0x1249CDA86774Bc170CAb843437DD37484F173ca8",
    //  "data":"0x00fdd58e000000000000000000000000e6e55eed00218faef27eed24def9208f3878b3330000000000000000000000000000000000000000000000000000000000000000"}, "latest"],"id":0}'
    //  -H "Content-Type: application/json" http://127.0.0.1:8544/
    //
    // method name:
    // > web3.sha3('balanceOf(address,uint256)')
    // "0x00fdd58ea0325fd79f486f8008ad3fad17dcb1cd2ee8474215c114771d87863e"
    //
    // data:
    // <method name>           + // 4 bytes
    // '0 x 24' + <to address> + // 32 bytes (64 chars)
    // '0 x X' + <token id>      // 32 bytes (64 chars)
    //
    // <0xa9059cbb> 000000000000000000000000 <to address>
    // <0 x X> + <token id>

    // result:
    // {"jsonrpc":"2.0","id":0,"result":"0x0000000000000000000000000000000000000000000000000de0b6b3a7640000"}



    fmt.Println( "eth_call(): balanceOf()" )
    fmt.Println( "eth_call(): holder address = ", _holder_address )
    fmt.Println( "eth_call(): token id = ", _token_id )

    var result types.Result

    //gas := "70000"
    //gasprice := "100"
    //value := ""
    //from := ""
    to := _to // erc-1155 contract address
    holder_address := _holder_address
    token_id := _token_id
    //to := "0x1e57f9561600b269a37437f02ce9da31e5b830ce" // erc-1155 contract address
    //holder_address := "0xe6e55eed00218faef27eed24def9208f3878b333"
    method := "0x00fdd58e"
    token_id_int := new(big.Int)
    token_id_int.SetString( token_id, 10 )
    token_id_hex := "0x" + token_id_int.Text( 16 )
    data := method + "000000000000000000000000" + holder_address[2:] +
            strings.Repeat("0", 64 - len(token_id_hex[2:])) + token_id_hex[2:]

    var params []interface{}
    //request_data_param := types.RequestData_params_erc1155_transaction { From: from, To: to, Value: value, Gas: gas, Gasprice: gasprice, Data: data }
    request_data_param := types.RequestData_params_erc1155 { To: to, Data: data }
    params = append( params, request_data_param, "latest" )
    request_data := types.RequestData { Jsonrpc: "2.0", Method: "eth_call", Params: params, Id: 0 }

    message, _ := json.Marshal( request_data )
    //fmt.Println( "message: ", request_data )
    response, err := http.Post( URL, "application/json", bytes.NewBuffer(message) )
    defer response.Body.Close()
    if err != nil {
        log.Fatal( "http.Post: ", err )
    }

    //fmt.Println( "response: " )
    responseBody, err := ioutil.ReadAll( response.Body )
    if err != nil {
        log.Fatal( "ioutil.ReadAll: ", err )
    }

    //fmt.Println( string(responseBody) )
    err = json.Unmarshal( responseBody, &result )
    if err != nil {
        log.Fatal( "json.Unmarshal: ", err )
    }
    fmt.Println( "jsonrpc =" , result.Jsonrpc, ", id =", result.Id, ", result =", result.Result )


    balance_wei := result.Result

    // SEE:
    // - https://golang.org/pkg/math/big/
    // - https://golang.org/pkg/strconv/
    // - https://goethereumbook.org/account-balance/
    balance_wei_int := new(big.Int)
    balance_wei_int.SetString( balance_wei[2:], 16 )
    fmt.Println( "erc-1155 token balance hex-string to int: ", balance_wei_int )
    //balance_wei_float := new(big.Float)
    //balance_wei_float.SetString( balance_wei_int.String() )
    //balance_token := new(big.Float).Quo(balance_wei_float, big.NewFloat(math.Pow10(18)))
    //fmt.Printf( "erc-1155 token balance: %.8f (%s)\n", balance_token, token_symbol )
}

func erc1155_get_uri(_to string, _holder_address string, _token_id string) {
    // eth erc-1155: eth_call

    // https://docs.openzeppelin.com/contracts/3.x/api/token/erc1155#IERC1155MetadataURI
    // https://eth.wiki/json-rpc/API
    // https://docs.soliditylang.org/en/latest/abi-spec.html
    // eth erc-1155: uri(uint256)
    //
    // request:
    // $ curl -X POST --data
    //  '{"jsonrpc":"2.0",
    //  "method":"eth_call",
    //  "params":[{"to": "0x1249CDA86774Bc170CAb843437DD37484F173ca8",
    //  "data":"0x0e89341c000000000000000000000000e6e55eed00218faef27eed24def9208f3878b3330000000000000000000000000000000000000000000000000000000000000000"}, "latest"],"id":0}'
    //  -H "Content-Type: application/json" http://127.0.0.1:8544/
    //
    // method name:
    // > web3.sha3('uri(uint256)')
    // "0x0e89341c5b7431e95282621bb9c54e51fb5c29234df43f9e19151d3892fb0380"
    //
    // data:
    // <method name>           + // 4 bytes
    // '0 x 24' + <to address> + // 32 bytes (64 chars)
    // '0 x X' + <token id>      // 32 bytes (64 chars)
    //
    // <0xa9059cbb> 000000000000000000000000 <to address>
    // <0 x X> + <token id>

    // result:
    // {"jsonrpc":"2.0","id":0,"result":"0x0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000003d<...>000000"}
    //
    // 2 bytes (0x) + 126 bytes (0000...20...000000) + 2 bytes (3d) + <URI> + 6 bytes (000000)
    // 0x
    // 0000000000
    // 0000000000
    // 0000000000
    // 0000000000
    // 0000000000
    // 0000000000
    // 0020000000
    // 0000000000
    // 0000000000
    // 0000000000
    // 0000000000
    // 0000000000
    // 000000
    // 3d
    // <URI>
    // 000000



    fmt.Println( "eth_call(): uri()" )

    var result types.Result

    //gas := "70000"
    //gasprice := "100"
    //value := ""
    //from := ""
    to := _to // erc-1155 contract address
    holder_address := _holder_address
    token_id := _token_id
    //to := "0x1e57f9561600b269a37437f02ce9da31e5b830ce" // erc-1155 contract address
    //holder_address := "0xe6e55eed00218faef27eed24def9208f3878b333"
    method := "0x0e89341c"
    token_id_int := new(big.Int)
    token_id_int.SetString( token_id, 10 )
    token_id_hex := "0x" + token_id_int.Text( 16 )
    data := method + "000000000000000000000000" + holder_address[2:] +
            strings.Repeat("0", 64 - len(token_id_hex[2:])) + token_id_hex[2:]

    var params []interface{}
    //request_data_param := types.RequestData_params_erc1155_transaction { From: from, To: to, Value: value, Gas: gas, Gasprice: gasprice, Data: data }
    request_data_param := types.RequestData_params_erc1155 { To: to, Data: data }
    params = append( params, request_data_param, "latest" )
    request_data := types.RequestData { Jsonrpc: "2.0", Method: "eth_call", Params: params, Id: 0 }

    message, _ := json.Marshal( request_data )
    //fmt.Println( "message: ", request_data )
    response, err := http.Post( URL, "application/json", bytes.NewBuffer(message) )
    defer response.Body.Close()
    if err != nil {
        log.Fatal( "http.Post: ", err )
    }

    //fmt.Println( "response: " )
    responseBody, err := ioutil.ReadAll( response.Body )
    if err != nil {
        log.Fatal( "ioutil.ReadAll: ", err )
    }

    //fmt.Println( string(responseBody) )
    err = json.Unmarshal( responseBody, &result )
    if err != nil {
        log.Fatal( "json.Unmarshal: ", err )
    }
    fmt.Println( "jsonrpc =" , result.Jsonrpc, ", id =", result.Id, ", result =", result.Result )


    uri_hex_str := result.Result

    // 2+126+2+6: 2 bytes (0x) + 126 bytes (0000...20...000000) + 2 byte (3d) + 6 bytes (000000)
    checks_len := 2 + 126 + 2 + 6

    fmt.Println( "checks len: ", checks_len )

    if len(uri_hex_str) <= checks_len {
        log.Fatal( "invalid length: ", len(uri_hex_str) )
    }

    fmt.Println( "uri hex str: ", uri_hex_str )
    fmt.Println( "url hex str len: ", len(uri_hex_str) )

    uri_hex := uri_hex_str[2+126+2:len(uri_hex_str) - 6]
    fmt.Println( "URI hex: ", uri_hex )

    uri_bs, err := hex.DecodeString( uri_hex )
    uri_str := string( uri_bs )

    fmt.Println( hex.Dump(uri_bs) )

    if err != nil {
        panic(err)
    }
    fmt.Println( "erc-1155 URI hex-string to str: ", uri_str )

    uri_with_token_id := strings.Replace( uri_str, "{id}", token_id, -1 )
    fmt.Println( "erc-1155 URI: ", uri_with_token_id )


    // https://docs.openzeppelin.com/contracts/3.x/erc1155#constructing_an_erc1155_token_contract
    // The uri can include the string {id} which clients must replace with the actual token ID,
    // in lowercase hexadecimal (with no 0x prefix) and leading zero padded to 64 hex characters.

    //token_id_bytes := []byte( token_id ) // from str
    //token_id_hex = hex.EncodeToString( token_id_bytes )
    //
    //token_id_hex = hex.EncodeToString( []byte(token_id) ) // from str

    //token_id_bytes := []byte( strconv.FormatInt(token_id_int, 16) ) // from int
    token_id_bytes := []byte( fmt.Sprintf("%x", token_id_int) ) // from int
    token_id_hex = hex.EncodeToString( token_id_bytes )

    fmt.Println( "token_id str: ", token_id )
    fmt.Println( "token_id hex (from str literally): ", token_id_hex )

    token_id_bs, err := hex.DecodeString( token_id_hex )
    token_id_str := string( token_id_bs )
    if err != nil {
        panic(err)
    }
    fmt.Println( "token_id ASCII: ", token_id_str )


    uri_with_token_id = strings.Repeat("0", 64 - len(token_id_hex)) + token_id_hex // from str literally
    //uri_with_token_id = strings.Repeat("0", 64 - len(token_id_str)) + token_id_str // ASCII
    uri_with_token_id = strings.Replace( uri_str, "{id}", uri_with_token_id, -1 )
    fmt.Println( "erc-1155 URI: ", uri_with_token_id )
}


//! TODO
func erc1155_set_uri(_to string, _new_uri string) {
    // eth erc-1155: eth_call

    // https://docs.openzeppelin.com/contracts/3.x/api/token/erc1155#IERC1155MetadataURI
    // https://eth.wiki/json-rpc/API
    // https://docs.soliditylang.org/en/latest/abi-spec.html
    // eth erc-1155: _setURI(string)
    //
    // request:
    // $ curl -X POST --data
    //  '{"jsonrpc":"2.0",
    //  "method":"eth_call",
    //  "params":[{"to": "0x1249CDA86774Bc170CAb843437DD37484F173ca8",
    //  "data":""}, "latest"],"id":0}'
    //  -H "Content-Type: application/json" http://127.0.0.1:8544/
    //
    // method name:
    // > web3.sha3('_setURI(string)')
    // "0xf392d4f5b83af9323820b3aa9898ce051f702e2a8c061e49eedd31b07826ce1a"
    //
    // data:
    //
    // <0xf392d4f5> 000000000000000000000000 <to address>
    // <0 x X> + <token id>

    // result:
    // 
}


func erc1155_send_transaction(_contract_address string, _from string, _to string, _token_id string, _amount string, _gas string, _gasprice string) {
    // eth erc-1155

    // https://docs.openzeppelin.com/contracts/3.x/api/token/erc1155#IERC1155MetadataURI
    // https://eth.wiki/json-rpc/API
    // https://docs.soliditylang.org/en/latest/abi-spec.html
    // eth erc-1155: safeTransferFrom(address,address,uint256,uint256,bytes)
    //
    // request:
    // $ curl -X POST --data
    //  '{"jsonrpc":"2.0",
    //  "method":"eth_sendTransaction",
    //  "params":[{
    //  "from":"0xe6e55eed00218faef27eed24def9208f3878b333",
    //  "to": "0x1249CDA86774Bc170CAb843437DD37484F173ca8",
    //  "gas":"0x11170","gasPrice":"0x174876e800",
    //  "value": "0x0",
    //  "data": "0xf242432a000000000000000000000000e6e55eed00218faef27eed24def9208f3878b3330000000000000000000000008f5b2b7608e3e3a3dc0426c3396420fbf18494540000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000000"}
    //  ],"id":0}'
    //  -H "Content-Type: application/json" http://127.0.0.1:8544/
    //
    // method name:
    // > web3.sha3('safeTransferFrom(address,address,uint256,uint256,bytes)')
    // "0xf242432a01954b0e0efb67e72c9b3b8ed77690657780385b256ac9aba0e35f0b"
    //
    // data:
    // <method name>               + // 4 bytes
    // '0 x 24' + <from address>   + // 32 bytes (64 chars)
    // '0 x 24' + <to address>     + // 32 bytes (64 chars)
    // '0 x X' + <token id>        + // 32 bytes (64 chars)
    // '0 x X' + <amount>          + // 32 bytes (64 chars)
    // '0 x x' + <a0>              + // 32 bytes (64 chars), length bytes params (a0: 32 bytes * 5(this length position))
    // '0 x x' + <bytes data>        // 32 bytes (64 chars)
    //
    // 0xf242432a
    // 000000000000000000000000e6e55eed00218faef27eed24def9208f3878b333
    // 0000000000000000000000008f5b2b7608e3e3a3dc0426c3396420fbf1849454
    // 0000000000000000000000000000000000000000000000000000000000000000
    // 0000000000000000000000000000000000000000000000000000000000000001
    // 00000000000000000000000000000000000000000000000000000000000000a0
    // 0000000000000000000000000000000000000000000000000000000000000000

    // result:
    // {"jsonrpc":"2.0","id":0,"result":"0xed7b9b146f3c9a28c2cb882097c8dd5177754216da2a92cf0793a984e60105b7"}

    fmt.Println( "eth_call(): transfer()" )

    var result types.Result

    contract_address := _contract_address
    from := _from
    to := _to
    //contract_address := "0x1e57f9561600b269a37437f02ce9da31e5b830ce"
    //from := "0xe6e55eed00218faef27eed24def9208f3878b333"
    //to := "0x8f5b2b7608e3e3a3dc0426c3396420fbf1849454"

    gas := _gas
    gasprice := _gasprice // 100 * gwei(1e9)
    value_amount := _amount // without token decimals

    value := "0" // for Ether, "0" fixed if ERC-20 transfer()
    //gas := "70000" // 70000
    //gasprice := "100" // 100 * gwei(1e9)
    //value_amount := "1" // without token decimals
    gas_hex := ""
    gasprice_hex := ""
    value_hex := ""
    value_amount_hex := ""
    token_id_hex := ""
    // ---
    {
        gas_int := new(big.Int)
        gas_float := new(big.Float)
        gasprice_int := new(big.Int)
        gasprice_float := new(big.Float)

        gas_float.SetString( gas )
        gasprice_float.SetString( gasprice )
        gasprice_decimals := big.NewFloat( math.Pow10(9) ) //new(big.Float)( math.Pow10(9) )
        gasprice_float_mul := new(big.Float).Mul( gasprice_float, gasprice_decimals ) // value * decimals(wei: 1e9)

        // float to int for hex
        // SEE: https://stackoverflow.com/questions/47545898/golang-convert-big-float-to-big-int
        gas_float.Int( gas_int )
        gasprice_float_mul.Int( gasprice_int )


        // ---


        value_amount_int := new(big.Int)
        value_amount_int.SetString( value_amount, 10 )
        token_id_int := new(big.Int)
        token_id_int.SetString( _token_id, 10 )


        gas_hex = "0x" + gas_int.Text( 16 )
        gasprice_hex = "0x" + gasprice_int.Text( 16 )
        value_hex = "0x" + value // always '0x0' for erc-20
        value_amount_hex = "0x" + value_amount_int.Text( 16 )
        token_id_hex = "0x" + token_id_int.Text( 16 )


        fmt.Printf( "gas: %s, %s\n" , _gas, gas_hex )
        fmt.Printf( "gas price: %s, %s\n" , _gasprice, gasprice_hex )
        fmt.Printf( "amount: %s, %s\n" , _amount, value_amount_hex )
        fmt.Printf( "token id: %s, %s\n" , _token_id, token_id_hex )
    }

    // ---

    method := "0xf242432a"
    data :=
            // method
            method +
            // from
            "000000000000000000000000" + from[2:] +
            // to
            "000000000000000000000000" + to[2:] +
            // token id
            strings.Repeat("0", 64 - len(token_id_hex[2:])) + token_id_hex[2:] +
            // amount
            strings.Repeat("0", 64 - len(value_amount_hex[2:])) + value_amount_hex[2:] +
            // length bytes params (a0: 32 bytes * 5(this length position))
            strings.Repeat("0", 64 - len("a0")) + "a0" +
            // bytes data
            strings.Repeat("0", 64)
    fmt.Printf( "data: %s\n", data )

    var params []interface{}
    request_data_param := types.RequestData_params_erc20_transaction {
        From: from, To: contract_address, Value: value_hex, Gas: gas_hex, Gasprice: gasprice_hex,
        Data: data,
    }
    //request_data_param := types.RequestData_params_erc20 { To: to, Data: data }
    params = append( params, request_data_param )
    request_data := types.RequestData { Jsonrpc: "2.0", Method: "eth_sendTransaction", Params: params, Id: 0 }


    {
        // unlock: personal_unlockAccount
        //
        // request:
        // $ curl -X POST --data
        //  '{"jsonrpc":"2.0",
        //  "method":"personal_unlockAccount",
        //  "params": ["0xe6e55eed00218faef27eed24def9208f3878b333","12345678",5], "id":0}'
        //  -H "Content-Type: application/json" http://127.0.0.1:8544/

        type Result struct {
            Jsonrpc string `json:"jsonrpc"`
            Id int `json:"id"`
            Result bool `json:"result"`
        }
        var result Result


        passphrase := "12345678"
        duration := 5
        var params []interface{}
        params = append( params, from, passphrase, duration )
        request_data := types.RequestData { Jsonrpc: "2.0", Method: "personal_unlockAccount", Params: params, Id: 0 }

        message, _ := json.Marshal( request_data )
        //fmt.Println( "message: ", request_data )

        response, err := http.Post( URL, "application/json", bytes.NewBuffer(message) )
        defer response.Body.Close()
        if err != nil {
            log.Fatal( "http.Post: ", err )
        }

        //fmt.Println( "response: " )
        responseBody, err := ioutil.ReadAll( response.Body )
        if err != nil {
            log.Fatal( "ioutil.ReadAll: ", err )
        }

        //fmt.Println( string(responseBody) )
        err = json.Unmarshal( responseBody, &result )
        if err != nil {
            log.Fatal( "json.Unmarshal: ", err )
        }
        //fmt.Println( "jsonrpc =" , result.Jsonrpc, ", id =", result.Id, ", result =", result.Result )
    }


    message, _ := json.Marshal( request_data )
    //fmt.Println( "message: ", request_data )

    response, err := http.Post( URL, "application/json", bytes.NewBuffer(message) )
    defer response.Body.Close()
    if err != nil {
        log.Fatal( "http.Post: ", err )
    }

    //fmt.Println( "response: " )
    responseBody, err := ioutil.ReadAll( response.Body )
    if err != nil {
        log.Fatal( "ioutil.ReadAll: ", err )
    }

    //fmt.Println( string(responseBody) )
    err = json.Unmarshal( responseBody, &result )
    if err != nil {
        log.Fatal( "json.Unmarshal: ", err )
    }
    fmt.Println( "jsonrpc =" , result.Jsonrpc, ", id =", result.Id, ", result =", result.Result )
    fmt.Println( "txid: ", result.Result )
}


// ETH, ERC-20 (transfer())
func get_blocks() {
    // eth_getBlockByNumber
    //
    // request:
    // $
    //

    fmt.Println( "eth_call(): eth_getBlockByNumber()" )

    //block_num_start_uint64 := uint64(502)
    //block_num_end_uint64 := uint64(503)
    block_num_start_uint64 := uint64(0)
    block_num_end_uint64 := uint64(0)

    {
        var result types.Result

        //var params []interface{}
        //request_data_param := types.RequestData {  }
        //params = append( params, "latest", true )
        request_data := types.RequestData { Jsonrpc: "2.0", Method: "eth_blockNumber", Id: 0 }

        message, _ := json.Marshal( request_data )
        //fmt.Println( "message: ", request_data )
        response, err := http.Post( URL, "application/json", bytes.NewBuffer(message) )
        defer response.Body.Close()
        if err != nil {
            log.Fatal( "http.Post: ", err )
        }

        //fmt.Println( "response: " )
        responseBody, err := ioutil.ReadAll( response.Body )
        if err != nil {
            log.Fatal( "ioutil.ReadAll: ", err )
        }

        //fmt.Println( string(responseBody) )
        err = json.Unmarshal( responseBody, &result )
        if err != nil {
            log.Fatal( "json.Unmarshal: ", err )
        }
        //fmt.Println( "jsonrpc =" , result.Jsonrpc, ", id =", result.Id, ", result =", result.Result )

        block_num_end_int := new(big.Int)
        block_num_end_int.SetString( result.Result[2:], 16 )
        block_num_end := block_num_end_int.String()
        //block_num_end_int32, _ = strconv.Atoi( block_num_end )
        block_num_end_uint64, _ = strconv.ParseUint( block_num_end, 10, 64 )
    }

    {
        /*
        type Result struct {
            Jsonrpc string `json:"jsonrpc"`
            Id int `json:"id"`
            Result string `json:"result"`
        }
        var result Result

        type Result_block struct {
            Jsonrpc string `json:"jsonrpc"`
            Id int `json:"id"`
            Result interface{} `json:"result"`
        }
        var result_block Result_block
        */

        var result types.Result
        var result_block types.Result_block




        fmt.Println( "block start = ", block_num_start_uint64 )
        fmt.Println( "block end = ", block_num_end_uint64 )
        fmt.Println()
        for i := block_num_start_uint64; i < uint64(block_num_end_uint64); i++ {
            request_block_num_hex := ""
            request_block_num_int := new(big.Int)
            request_block_num_int.SetUint64( uint64(i) )
            request_block_num_hex = "0x" + request_block_num_int.Text( 16 )

            //fmt.Println( i, "(" + request_block_num_hex + ")", "----------" )

            var params []interface{}
            //request_data_param := types.RequestData {  }
            //params = append( params, "latest", true )
            params = append( params, request_block_num_hex, true )
            request_data := types.RequestData { Jsonrpc: "2.0", Method: "eth_getBlockByNumber", Params: params, Id: 0 }

            message, _ := json.Marshal( request_data )
            //fmt.Println( "message: ", request_data )
            response, err := http.Post( URL, "application/json", bytes.NewBuffer(message) )
            defer response.Body.Close()
            if err != nil {
                log.Fatal( "http.Post: ", err )
            }

            //fmt.Println( "response: " )
            responseBody, err := ioutil.ReadAll( response.Body )
            if err != nil {
                log.Fatal( "ioutil.ReadAll: ", err )
            }

            //fmt.Println( string(responseBody) )
            err = json.Unmarshal( responseBody, &result_block )
            if err != nil {
                log.Fatal( "json.Unmarshal: ", err )
            }
            //fmt.Println( "jsonrpc =" , result.Jsonrpc, ", id =", result.Id, ", result =", result_block.Result )


            //fmt.Println( "data =", result_block.Result )
            //fmt.Println( reflect.TypeOf(result_block.Result) )
            _txns := result_block.Result.(map[string]interface{})["transactions"]
            fmt.Println( "transaction size = ", len(_txns.([]interface{})) )
            if len(_txns.([]interface{})) <= 0 {
                //fmt.Println( "no transaction: size = 0" )
                continue
            }

            for j := 0; j < len(_txns.([]interface{})); j++ {
                //_txns = _txns.([]interface{})[0]
                //txn := _txns.(map[string]interface{})
                _txn := _txns.([]interface{})[j]
                txn := _txn.(map[string]interface{})

                //if txn["from"] != "0x8f5b2b7608e3e3a3dc0426c3396420fbf1849454" && txn["to"] != "0x8f5b2b7608e3e3a3dc0426c3396420fbf1849454" {
                //    continue
                //}

                timestamp_hex := result_block.Result.(map[string]interface{})["timestamp"]
                timestamp_int := new(big.Int)
                timestamp_int.SetString( timestamp_hex.(string)[2:], 16 )
                timestamp_unixtime := timestamp_int.String()
                //timestamp_int32, _ := strconv.Atoi( timestamp_unixtime )
                timestamp_int64, _ := strconv.ParseInt( timestamp_unixtime, 10, 64 )
                tx_timestamp_date := time.Unix( timestamp_int64, 0 )
                //tx_timestamp_date_rfc3339 := timestamp_date.Format( time.RFC3339 )


                tx_hash := txn["hash"]
                tx_block_number_hex := txn["blockNumber"]
                tx_block_number := ""
                {
                    block_number_int := new(big.Int)
                    block_number_int.SetString( tx_block_number_hex.(string)[2:], 16 )
                    //fmt.Println( "ether hex-string to int: ", block_number_int )
                    tx_block_number = block_number_int.String()
                }
                tx_from := txn["from"]
                tx_to := txn["to"]
                tx_value_wei_hex := txn["value"]
                tx_value_wei := ""
                tx_value := "" // Ether
                tx_input := txn["input"]

                tx_token_to := "" // for ERC-20
                tx_token_name := ""
                tx_token_symbol := ""
                tx_token_decimals := ""
                tx_token_total_supply := ""
                tx_token_amount_wei_hex := ""
                tx_token_amount_wei := ""
                tx_token_amount := ""

                if tx_to == nil {
                    continue
                }

                //fmt.Println( "transaction =", _txn )
                //fmt.Println( "hash =", tx_hash )
                //fmt.Println( "timestamp =", tx_timestamp_date ) // "Y/m/d/ H:i:s"
                //fmt.Println( "block_number =", tx_block_number )
                //fmt.Println( "from =", tx_from )

                if txn["input"] == "0x" {
                    fmt.Println( "Ether" )

                    fmt.Println( "hash =", tx_hash )
                    fmt.Println( "timestamp =", tx_timestamp_date ) // "Y/m/d/ H:i:s"
                    fmt.Println( "block_number =", tx_block_number )
                    fmt.Println( "from =", tx_from )
                    fmt.Println( "to =", tx_to )

                    // SEE:
                    // - https://golang.org/pkg/math/big/
                    // - https://golang.org/pkg/strconv/
                    // - https://goethereumbook.org/account-balance/
                    amount_wei_int := new(big.Int)
                    amount_wei_int.SetString( tx_value_wei_hex.(string)[2:], 16 )
                    fmt.Println( "ether hex-string to int: ", amount_wei_int, "(wei)" )
                    amount_wei_float := new(big.Float)
                    amount_wei_float.SetString( amount_wei_int.String() )
                    tx_value_float := new(big.Float).Quo(amount_wei_float, big.NewFloat(math.Pow10(18)))
                    tx_value = fmt.Sprintf( "%.8f", tx_value_float )
                    tx_value_wei = amount_wei_int.String()

                    fmt.Println( "value_wei =", tx_value_wei, "(wei)" )
                    fmt.Println( "value_ether =", tx_value, "(ether)" )
                    fmt.Println()
                } else {
                    //fmt.Println( "ERC-xxxx" )

                    //fmt.Println( "input data =", tx_input )

                    // token to: [2: 0x] + [8: method] + [0 x 24]
                    tx_token_to = "0x" + tx_input.(string)[2 + 8 + 24:(2+8+24 + 40)]

                    // amount: 32 bytes (64 chars): [2: 0x] + [8: method] + [0 x 24] + [40: to address]
                    tx_token_amount_wei_hex = "0x" + tx_input.(string)[2 + 8 + 24 + 40:]

                    method := ""
                    data := ""

                    //fmt.Println( "method =", tx_input.(string)[:10] )
                    if tx_input.(string)[:10] != "0xa9059cbb" {
                        //fmt.Println( "Not ERC-20 transfer transaction" )
                        continue
                    }

                    {
                        _token_name := ""
                        _token_symbol := ""
                        _token_decimals := ""
                        _token_total_supply := ""

                        {
                            // Token: name
                            method = "0x06fdde03"
                            data = method + "000000000000000000000000" + tx_from.(string)[2:]

                            var params []interface{}
                            request_data_param := types.RequestData_params_erc20 { To: tx_to.(string), Data: data }
                            params = append( params, request_data_param, "latest" )
                            request_data := types.RequestData { Jsonrpc: "2.0", Method: "eth_call", Params: params, Id: 0 }

                            message, _ := json.Marshal( request_data )
                            //fmt.Println( "message: ", request_data )
                            response, err := http.Post( URL, "application/json", bytes.NewBuffer(message) )
                            defer response.Body.Close()
                            if err != nil {
                                log.Fatal( "http.Post: ", err )
                            }

                            //fmt.Println( "response: " )
                            responseBody, err := ioutil.ReadAll( response.Body )
                            if err != nil {
                                log.Fatal( "ioutil.ReadAll: ", err )
                            }

                            //fmt.Println( string(responseBody) )
                            err = json.Unmarshal( responseBody, &result )
                            if err != nil {
                                log.Fatal( "json.Unmarshal: ", err )
                            }
                            //fmt.Println( "jsonrpc =" , result.Jsonrpc, ", id =", result.Id, ", result =", result.Result )
                            _token_name = result.Result
                        }

                        {
                            // Token: symbol
                            method = "0x95d89b41"
                            data = method + "000000000000000000000000" + tx_from.(string)[2:]

                            var params []interface{}
                            request_data_param := types.RequestData_params_erc20 { To: tx_to.(string), Data: data }
                            params = append( params, request_data_param, "latest" )
                            request_data := types.RequestData { Jsonrpc: "2.0", Method: "eth_call", Params: params, Id: 0 }

                            message, _ := json.Marshal( request_data )
                            //fmt.Println( "message: ", request_data )
                            response, err := http.Post( URL, "application/json", bytes.NewBuffer(message) )
                            defer response.Body.Close()
                            if err != nil {
                                log.Fatal( "http.Post: ", err )
                            }

                            //fmt.Println( "response: " )
                            responseBody, err := ioutil.ReadAll( response.Body )
                            if err != nil {
                                log.Fatal( "ioutil.ReadAll: ", err )
                            }

                            //fmt.Println( string(responseBody) )
                            err = json.Unmarshal( responseBody, &result )
                            if err != nil {
                                log.Fatal( "json.Unmarshal: ", err )
                            }
                            //fmt.Println( "jsonrpc =" , result.Jsonrpc, ", id =", result.Id, ", result =", result.Result )
                            _token_symbol = result.Result
                        }

                        {
                            // Token: decimals
                            method = "0x313ce567"
                            data = method + "000000000000000000000000" + tx_from.(string)[2:]

                            var params []interface{}
                            request_data_param := types.RequestData_params_erc20 { To: tx_to.(string), Data: data }
                            params = append( params, request_data_param, "latest" )
                            request_data := types.RequestData { Jsonrpc: "2.0", Method: "eth_call", Params: params, Id: 0 }

                            message, _ := json.Marshal( request_data )
                            //fmt.Println( "message: ", request_data )
                            response, err := http.Post( URL, "application/json", bytes.NewBuffer(message) )
                            defer response.Body.Close()
                            if err != nil {
                                log.Fatal( "http.Post: ", err )
                            }

                            //fmt.Println( "response: " )
                            responseBody, err := ioutil.ReadAll( response.Body )
                            if err != nil {
                                log.Fatal( "ioutil.ReadAll: ", err )
                            }

                            //fmt.Println( string(responseBody) )
                            err = json.Unmarshal( responseBody, &result )
                            if err != nil {
                                log.Fatal( "json.Unmarshal: ", err )
                            }
                            //fmt.Println( "jsonrpc =" , result.Jsonrpc, ", id =", result.Id, ", result =", result.Result )
                            _token_decimals = result.Result
                        }

                        {
                            // Token: total_supply 
                            method = "0x18160ddd"
                            data = method + "000000000000000000000000" + tx_from.(string)[2:]

                            var params []interface{}
                            request_data_param := types.RequestData_params_erc20 { To: tx_to.(string), Data: data }
                            params = append( params, request_data_param, "latest" )
                            request_data := types.RequestData { Jsonrpc: "2.0", Method: "eth_call", Params: params, Id: 0 }

                            message, _ := json.Marshal( request_data )
                            //fmt.Println( "message: ", request_data )
                            response, err := http.Post( URL, "application/json", bytes.NewBuffer(message) )
                            defer response.Body.Close()
                            if err != nil {
                                log.Fatal( "http.Post: ", err )
                            }

                            //fmt.Println( "response: " )
                            responseBody, err := ioutil.ReadAll( response.Body )
                            if err != nil {
                                log.Fatal( "ioutil.ReadAll: ", err )
                            }

                            //fmt.Println( string(responseBody) )
                            err = json.Unmarshal( responseBody, &result )
                            if err != nil {
                                log.Fatal( "json.Unmarshal: ", err )
                            }
                            //fmt.Println( "jsonrpc =" , result.Jsonrpc, ", id =", result.Id, ", result =", result.Result )
                            _token_total_supply = result.Result
                        }



                        //-----{
                        // token name: 0x + [60 bytes] + [4 bytes] + [60 bytes] + [4 bytes]:
                        __token_name, _ := hex.DecodeString( _token_name[2 + 60 + 4 + 60 + 4:] )
                        tx_token_name = string(bytes.Trim(__token_name, "\x00"))

                        // token symbol: 0x + [60 bytes] + [4 bytes] + [60 bytes] + [4 bytes]:
                        __token_symbol, _ := hex.DecodeString( _token_symbol[2 + 60 + 4 + 60 + 4:] )
                        tx_token_symbol = string(bytes.Trim(__token_symbol, "\x00"))

                        // token decimals: 0x + [60 bytes] + [4 bytes]
                        token_decimals_int := new(big.Int)
                        token_decimals_int.SetString( _token_decimals[2:], 16 )
                        __token_decimals := token_decimals_int.String()
                        token_decimals_int32, _ := strconv.Atoi( __token_decimals )
                        tx_token_decimals = __token_decimals

                        // token total supply:
                        token_total_supply_int := new(big.Int)
                        token_total_supply_int.SetString( _token_total_supply[2:], 16 )
                        token_total_supply_float := new(big.Float)
                        token_total_supply_float.SetString( token_total_supply_int.String() )
                        __token_total_supply := new(big.Float).Quo(token_total_supply_float, big.NewFloat(math.Pow10(token_decimals_int32)))
                        tx_token_total_supply = fmt.Sprintf( "%.8f", __token_total_supply )

                        //fmt.Println( "token name:", string(__token_name) )
                        //fmt.Println( "token_symbol:", string(__token_symbol) )
                        //fmt.Println( "token_decimals:", __token_decimals )
                        //fmt.Printf( "token_total_supply: %f\n", __token_total_supply )
                        //-----}


                        // SEE:
                        // - https://golang.org/pkg/math/big/
                        // - https://golang.org/pkg/strconv/
                        // - https://goethereumbook.org/account-balance/
                        token_amount_wei_int := new(big.Int)
                        token_amount_wei_int.SetString( tx_token_amount_wei_hex[2:], 16 )
                        //fmt.Println( "erc-20 token amount hex-string to int: ", token_amount_wei_int, "(wei)" )
                        token_amount_wei_float := new(big.Float)
                        token_amount_wei_float.SetString( token_amount_wei_int.String() )
                        token_amount := new(big.Float).Quo(token_amount_wei_float, big.NewFloat(math.Pow10(token_decimals_int32)))
                        tx_token_amount = fmt.Sprintf( "%.8f", token_amount )
                        tx_token_amount_wei = token_amount_wei_int.String()
                        //fmt.Printf( "erc-20 token amount: %s (%s)\n", tx_token_amount, tx_token_symbol )
                    }


                    fmt.Println( "hash =", tx_hash )
                    fmt.Println( "timestamp =", tx_timestamp_date ) // "Y/m/d/ H:i:s"
                    fmt.Println( "block_number =", tx_block_number )
                    fmt.Println( "from =", tx_from )

                    fmt.Println( "token_contract address =", tx_to )
                    fmt.Println( "token_to =", tx_token_to )
                    fmt.Println( "token_name =", tx_token_name )
                    fmt.Println( "token_symbol =", tx_token_symbol )
                    fmt.Println( "token_decimals =", tx_token_decimals )
                    fmt.Println( "token_total_supply =", tx_token_total_supply )
                    fmt.Println( "token_value_wei =", tx_token_amount_wei, "(wei)" )
                    fmt.Println( "token_value_" + tx_token_symbol + " =", tx_token_amount, "(" + tx_token_symbol + ")" )
                    fmt.Println()
                }
            } // for (), transactions
        } // for (), blocks
    }
}


// ERC-1155 (safeTransferOf())
func get_blocks_erc1155() {
    // eth_getBlockByNumber
    //
    // request:
    // $
    //

    fmt.Println( "eth_call(): eth_getBlockByNumber()" )

    //block_num_start_uint64 := uint64(502)
    //block_num_end_uint64 := uint64(503)
    block_num_start_uint64 := uint64(0)
    block_num_end_uint64 := uint64(0)

    {
        var result types.Result

        //var params []interface{}
        //request_data_param := types.RequestData {  }
        //params = append( params, "latest", true )
        request_data := types.RequestData { Jsonrpc: "2.0", Method: "eth_blockNumber", Id: 0 }

        message, _ := json.Marshal( request_data )
        //fmt.Println( "message: ", request_data )
        response, err := http.Post( URL, "application/json", bytes.NewBuffer(message) )
        defer response.Body.Close()
        if err != nil {
            log.Fatal( "http.Post: ", err )
        }

        //fmt.Println( "response: " )
        responseBody, err := ioutil.ReadAll( response.Body )
        if err != nil {
            log.Fatal( "ioutil.ReadAll: ", err )
        }

        //fmt.Println( string(responseBody) )
        err = json.Unmarshal( responseBody, &result )
        if err != nil {
            log.Fatal( "json.Unmarshal: ", err )
        }
        //fmt.Println( "jsonrpc =" , result.Jsonrpc, ", id =", result.Id, ", result =", result.Result )

        block_num_end_int := new(big.Int)
        block_num_end_int.SetString( result.Result[2:], 16 )
        block_num_end := block_num_end_int.String()
        //block_num_end_int32, _ = strconv.Atoi( block_num_end )
        block_num_end_uint64, _ = strconv.ParseUint( block_num_end, 10, 64 )
    }

    {
        /*
        type Result struct {
            Jsonrpc string `json:"jsonrpc"`
            Id int `json:"id"`
            Result string `json:"result"`
        }
        var result Result

        type Result_block struct {
            Jsonrpc string `json:"jsonrpc"`
            Id int `json:"id"`
            Result interface{} `json:"result"`
        }
        var result_block Result_block
        */

        //var result types.Result
        var result_block types.Result_block




        fmt.Println( "block start = ", block_num_start_uint64 )
        fmt.Println( "block end = ", block_num_end_uint64 )
        fmt.Println()
        for i := block_num_start_uint64; i < uint64(block_num_end_uint64); i++ {
            request_block_num_hex := ""
            request_block_num_int := new(big.Int)
            request_block_num_int.SetUint64( uint64(i) )
            request_block_num_hex = "0x" + request_block_num_int.Text( 16 )

            //fmt.Println( i, "(" + request_block_num_hex + ")", "----------" )

            var params []interface{}
            //request_data_param := types.RequestData {  }
            //params = append( params, "latest", true )
            params = append( params, request_block_num_hex, true )
            request_data := types.RequestData { Jsonrpc: "2.0", Method: "eth_getBlockByNumber", Params: params, Id: 0 }

            message, _ := json.Marshal( request_data )
            //fmt.Println( "message: ", request_data )
            response, err := http.Post( URL, "application/json", bytes.NewBuffer(message) )
            defer response.Body.Close()
            if err != nil {
                log.Fatal( "http.Post: ", err )
            }

            //fmt.Println( "response: " )
            responseBody, err := ioutil.ReadAll( response.Body )
            if err != nil {
                log.Fatal( "ioutil.ReadAll: ", err )
            }

            //fmt.Println( string(responseBody) )
            err = json.Unmarshal( responseBody, &result_block )
            if err != nil {
                log.Fatal( "json.Unmarshal: ", err )
            }
            //fmt.Println( "jsonrpc =" , result.Jsonrpc, ", id =", result.Id, ", result =", result_block.Result )


            //fmt.Println( "data =", result_block.Result )
            //fmt.Println( reflect.TypeOf(result_block.Result) )
            _txns := result_block.Result.(map[string]interface{})["transactions"]
            fmt.Println( "transaction size = ", len(_txns.([]interface{})) )
            if len(_txns.([]interface{})) <= 0 {
                //fmt.Println( "no transaction: size = 0" )
                continue
            }

            for j := 0; j < len(_txns.([]interface{})); j++ {
                //_txns = _txns.([]interface{})[0]
                //txn := _txns.(map[string]interface{})
                _txn := _txns.([]interface{})[j]
                txn := _txn.(map[string]interface{})

                //if txn["from"] != "0x8f5b2b7608e3e3a3dc0426c3396420fbf1849454" && txn["to"] != "0x8f5b2b7608e3e3a3dc0426c3396420fbf1849454" {
                //if txn["from"] != "0xe6e55eed00218faef27eed24def9208f3878b333" && txn["to"] != "0x1249CDA86774Bc170CAb843437DD37484F173ca8" {
                //    continue
                //}

                timestamp_hex := result_block.Result.(map[string]interface{})["timestamp"]
                timestamp_int := new(big.Int)
                timestamp_int.SetString( timestamp_hex.(string)[2:], 16 )
                timestamp_unixtime := timestamp_int.String()
                //timestamp_int32, _ := strconv.Atoi( timestamp_unixtime )
                timestamp_int64, _ := strconv.ParseInt( timestamp_unixtime, 10, 64 )
                tx_timestamp_date := time.Unix( timestamp_int64, 0 )
                //tx_timestamp_date_rfc3339 := timestamp_date.Format( time.RFC3339 )


                tx_hash := txn["hash"]
                tx_block_number_hex := txn["blockNumber"]
                tx_block_number := ""
                {
                    block_number_int := new(big.Int)
                    block_number_int.SetString( tx_block_number_hex.(string)[2:], 16 )
                    //fmt.Println( "ether hex-string to int: ", block_number_int )
                    tx_block_number = block_number_int.String()
                }
                tx_from := txn["from"]
                tx_to := txn["to"]
                tx_value_wei_hex := txn["value"]
                tx_value_wei := ""
                tx_value := "" // Ether
                tx_input := txn["input"]

                tx_token_from := "" // for ERC-1155
                tx_token_to := "" // for ERC-1155
                tx_token_amount_hex := ""
                tx_token_amount := ""
                tx_token_id_hex := ""
                //tx_token_id := ""
                tx_token_data := ""
                tx_token_data_length := ""

                if tx_to == nil {
                    continue
                }

                //fmt.Println( "transaction =", _txn )
                //fmt.Println( "hash =", tx_hash )
                //fmt.Println( "timestamp =", tx_timestamp_date ) // "Y/m/d/ H:i:s"
                //fmt.Println( "block_number =", tx_block_number )
                //fmt.Println( "from =", tx_from )

                if txn["input"] == "0x" {
                    fmt.Println( "Ether" )

                    fmt.Println( "hash =", tx_hash )
                    fmt.Println( "timestamp =", tx_timestamp_date ) // "Y/m/d/ H:i:s"
                    fmt.Println( "block_number =", tx_block_number )
                    fmt.Println( "from =", tx_from )
                    fmt.Println( "to =", tx_to )

                    // SEE:
                    // - https://golang.org/pkg/math/big/
                    // - https://golang.org/pkg/strconv/
                    // - https://goethereumbook.org/account-balance/
                    amount_wei_int := new(big.Int)
                    amount_wei_int.SetString( tx_value_wei_hex.(string)[2:], 16 )
                    fmt.Println( "ether hex-string to int: ", amount_wei_int, "(wei)" )
                    amount_wei_float := new(big.Float)
                    amount_wei_float.SetString( amount_wei_int.String() )
                    tx_value_float := new(big.Float).Quo(amount_wei_float, big.NewFloat(math.Pow10(18)))
                    tx_value = fmt.Sprintf( "%.8f", tx_value_float )
                    tx_value_wei = amount_wei_int.String()

                    fmt.Println( "value_wei =", tx_value_wei, "(wei)" )
                    fmt.Println( "value_ether =", tx_value, "(ether)" )
                    fmt.Println()
                } else {
                    //fmt.Println( "ERC-1155" )

                    fmt.Println( "input data =", tx_input )

                    //fmt.Println( "method =", tx_input.(string)[:10] )
                    //if tx_input.(string)[:10] != "0xa9059cbb" {
                    //    fmt.Println( "Not ERC-20 transfer transaction" )
                    //    continue
                    //}


                    fmt.Println( "method =", tx_input.(string)[:10] )
                    if tx_input.(string)[:10] != "0xf242432a" {
                        fmt.Println( "Not ERC-1155 safeTransferFrom transaction" )
                        continue
                    }


                    //tx_data_offset := 0

                    //tx_data_offset = 2 + 8 + 24
                    // token from: 32 bytes (64 chars): [2: 0x] + [8: method] + [0 x 24] ~ [40]
                    tx_token_from = "0x" + tx_input.(string)[2 + 8 + 24:(2+8+24 + 40)]

                    //tx_data_offset = 2 + 8 + 24+40 + 24
                    // token to: 32 bytes (64 chars): [2: 0x] + [8: method] + [0 x 24] ~ [40]
                    tx_token_to = "0x" + tx_input.(string)[2 + 8 + 24+40 + 24:(2+8+24+40 + 64)]

                    //tx_data_offset = 2 + 8 + 64 + 64
                    // token id: 32 bytes (64 chars): [2: 0x] + [8: method] + [64] + [64]
                    tx_token_id_hex = "0x" + tx_input.(string)[2 + 8 + 64 + 64:(2+8+64+64 + 64)]

                    //tx_data_offset = 2 + 8 + 64 + 64 + 64
                    // amount: 32 bytes (64 chars): [2: 0x] + [8: method] + [64] + [64] + [64]
                    tx_token_amount_hex = "0x" + tx_input.(string)[2 + 8 + 64 + 64 + 64:(2+8+64+64+64 + 64)]

                    //tx_data_offset = 2 + 8 + 64 + 64 + 64 + 64
                    // data length bytes: 32 bytes (64 chars): [2: 0x] + [8: method] + [64] + [64] + [64] + [64]
                    tx_token_data_length = "0x" + tx_input.(string)[2 + 8 + 64 + 64 + 64 + 64:(2+8+64+64+64+64 + 64)]

                    //tx_data_offset = 2 + 8 + 64 + 64 + 64 + 64 + 64
                    // data length bytes: 32 bytes (64 chars): [2: 0x] + [8: method] + [64] + [64] + [64] + [64] + [64]
                    tx_token_data = "0x" + tx_input.(string)[2 + 8 + 64 + 64 + 64 + 64 + 64:(2+8+64+64+64+64+64 + 64)]


                    {
                        token_amount_int := new(big.Int)
                        token_amount_int.SetString( tx_token_amount_hex[2:], 16 )
                        tx_token_amount = token_amount_int.String()
                    }


                    fmt.Println( "hash =", tx_hash )
                    fmt.Println( "timestamp =", tx_timestamp_date ) // "Y/m/d/ H:i:s"
                    fmt.Println( "block_number =", tx_block_number )
                    fmt.Println( "from =", tx_from )

                    fmt.Println( "token_contract address =", tx_to )
                    fmt.Println( "token_from =", tx_token_from )
                    fmt.Println( "token_to =", tx_token_to )
                    fmt.Println( "token_id = ", tx_token_id_hex )
                    fmt.Println( "token_amount = ", tx_token_amount )
                    fmt.Println( "token_data_length = ", tx_token_data_length )
                    fmt.Println( "token_data = ", tx_token_data )
                    fmt.Println()
                }
            } // for (), transactions
        } // for (), blocks
    }
}



// ETH, ERC-20 (transfer()), ERC-1155 (safeTransferOf())
// SEE: func get_blocks_all_infinite(...) {}
func get_blocks_all() {
    TAG := "get_blocks_all(): "

    //var INTERVAL = int(1) // seconds

    count_repeat := uint64(0)
    block_height_old_uint64 := uint64(0)
    block_height_from_db := ""
    block_height_from_db_uint64 := uint64(0)

    //block_num_start_uint64 := uint64(0)
    block_num_end_uint64 := uint64(0)


    // get last block height inserted from DB
    // if duplicated, update(replace) it...
    {
        query_str := "SELECT blocks FROM txid ORDER BY idx DESC limit 1;"
        rows, err := gDB.Query( query_str )

        if err != nil {
            panic( err.Error() )
            return
        }
        //defer rows.Close()

        for rows.Next() {
            //var Block_number string
            err := rows.Scan( &block_height_from_db )

            if err != nil {
                //log.Fatal(err)
                panic( err )
            }

            break
        }
        rows.Close()

        if len(block_height_from_db) <= 0 {
            block_height_from_db = "0"
        }

        //fmt.Println( TAG, "block height from DB (str) = ", block_height_from_db )
        block_height_from_db_uint64, _ = strconv.ParseUint( block_height_from_db, 10, 64 )
        //fmt.Println( TAG, "block height from DB = ", block_height_from_db_uint64 )
    }



    for {
        //time.Sleep( time.Second * time.Duration(INTERVAL) )
        time.Sleep( time.Millisecond * 1000 * time.Duration(INTERVAL) )

        var result types.Result

        //var params []interface{}
        //request_data_param := types.RequestData {  }
        //params = append( params, "latest", true )
        request_data := types.RequestData { Jsonrpc: "2.0", Method: "eth_blockNumber", Id: 0 }

        message, _ := json.Marshal( request_data )
        //fmt.Println( "message: ", request_data )
        response, err := http.Post( URL, "application/json", bytes.NewBuffer(message) )
        defer response.Body.Close()
        if err != nil {
            log.Fatal( "http.Post: ", err )
        }

        //fmt.Println( "response: " )
        responseBody, err := ioutil.ReadAll( response.Body )
        if err != nil {
            log.Fatal( "ioutil.ReadAll: ", err )
        }

        //fmt.Println( string(responseBody) )
        err = json.Unmarshal( responseBody, &result )
        if err != nil {
            log.Fatal( "json.Unmarshal: ", err )
        }
        //fmt.Println( "jsonrpc =" , result.Jsonrpc, ", id =", result.Id, ", result =", result.Result )

        block_num_end_int := new(big.Int)
        block_num_end_int.SetString( result.Result[2:], 16 )
        block_num_end := block_num_end_int.String()
        //block_num_end_int32, _ = strconv.Atoi( block_num_end )
        block_num_end_uint64, _ = strconv.ParseUint( block_num_end, 10, 64 )


// test
//if ( block_height_from_db_uint64 < 240 ) {
//block_height_from_db_uint64 = 240
//block_height_old_uint64 = block_height_from_db_uint64
//}


        if ( block_height_old_uint64 <= 0 ) {
            //block_height_old_uint64 = block_num_end_uint64

            // block height from DB
            block_height_old_uint64 = block_height_from_db_uint64
        }

        fmt.Println( TAG, "[block height] = ", block_height_old_uint64, "/", block_num_end_uint64 )

        if ( block_num_end_uint64 == block_height_old_uint64 ) {
            count_repeat++
        } else {
            count_repeat = 0
        }

        if ( count_repeat > 1 ) {
            fmt.Println( TAG, "skip... count = ", count_repeat )
            fmt.Println()
            continue
        }

        get_blocks_all_infinite( block_height_old_uint64, block_height_old_uint64 + 1 )

        if ( block_num_end_uint64 > block_height_old_uint64 ) {
            block_height_old_uint64++
        }

        fmt.Println()
    } // for ()
}

//func _get_token_creation_txn_info(txn map[string]interface{}) {
//func _get_token_creation_txn_info(txid string ) types.Fetch_transactions_st {
func _get_token_creation_txn_info(txid string, timestamp string, datetime string) types.Fetch_transactions_st {
    TAG := "_get_token_creation_txn_info(): "

    fmt.Println( TAG, "Token creation transaction: 0x60606040 or 0x60806040" )


    var txn_receipt__contract_address = ""
    var txn_receipt__from_address = "" // owner address, // token creation: 0x0000000000000000000000000000000000000000
    var txn_receipt__to_address = ""
    var txn_receipt__token_symbol_id = ""
    var txn_receipt__total_supply = ""
    var txn_receipt__log map[string]interface{}
    //var txn map[string]interface{}


    var tx_block_number = ""
    var tx_hash = txid
    var timestamp_unixtime = ""
    //var tx_timestamp_date time.Time
    var tx_timestamp_date = ""

    timestamp_unixtime = timestamp
    tx_timestamp_date = datetime



    tx_token_type := "erc20" // default: "erc20"
    var result types.Result
    //var result_block types.Result_block

    var _data types.Fetch_transactions_st


    fmt.Println( TAG, "txid = ", txid )



    // eth.getTransactionReceipt: get contract address from 0x60606040/0x60806040 txid
    var result_block types.Result_block
    var params []interface{}
    //request_data_param := types.RequestData {  }
    //params = append( params, "latest", true )
    params = append( params, txid )
    request_data := types.RequestData { Jsonrpc: "2.0", Method: "eth_getTransactionReceipt", Params: params, Id: 0 }

    message, _ := json.Marshal( request_data )
    //fmt.Println( TAG, "message: ", request_data )
    response, err := http.Post( URL, "application/json", bytes.NewBuffer(message) )
    defer response.Body.Close()
    if err != nil {
        log.Fatal( "http.Post: ", err )
    }

    //fmt.Println( TAG, "response: " )
    responseBody, err := ioutil.ReadAll( response.Body )
    if err != nil {
        log.Fatal( "ioutil.ReadAll: ", err )
    }

    //fmt.Println( TAG, string(responseBody) )
    err = json.Unmarshal( responseBody, &result_block )
    if err != nil {
        log.Fatal( "json.Unmarshal: ", err )
    }
    fmt.Println( TAG, "jsonrpc =" , result.Jsonrpc, ", id =", result.Id, ", result =", result_block.Result )

    if result_block.Result == nil {
        fmt.Println( TAG, "result == NULL" )
        return _data
    }


    //fmt.Println( TAG, "data =", result_block.Result )
    //fmt.Println( TAG, reflect.TypeOf(result_block.Result) )
    //_txns := result_block.Result.(map[string]interface{})["transactions"]
    //fmt.Println( TAG, "transaction size = ", len(_txns.([]interface{})) )
    _txns := result_block.Result.(map[string]interface{})
    fmt.Println()

    if len(_txns["logs"].([]interface{})) <= 0 {
        fmt.Println( TAG, "no logs: size = 0; skip..." )
        return _data
    }
    fmt.Println( TAG, "logs size = ", len(_txns["logs"].([]interface{})) )
    logs_size := len( _txns["logs"].([]interface{}) )
    // eth.getTransactionReceipt( "token creation txid" )
    // {    ...
    //      contractAddress:"",
    //      logs: [{
    //          ..., data:"", topics: ["<hash...>", "<from address: NULL>", "<to address: will be received token>"], ...
    //      }, {
    //          ..., data:"", topics: ["<hash...>", "<from address: NULL>", "<to address: will be received token>"], ...
    //      }]
    //      ...
    // }
    /*
    switch logs_size {
    case 1: // [0] migration + deploy
        txn_receipt__log = _txns["logs"].([]interface{})[0].(map[string]interface{})
        break
    case 2: // [0] migration, [1] deploy (use this)
        txn_receipt__log = _txns["logs"].([]interface{})[1].(map[string]interface{})
        break
    default:
        fmt.Println( TAG, "size > 2" )
        return _data
    }
    */

    for i := 0; i < logs_size; i++ {
        txn_receipt__log = _txns["logs"].([]interface{})[i].(map[string]interface{})

        fmt.Println( txn_receipt__log )
        if txn_receipt__log == nil {
            fmt.Println( TAG, "txn_log == NULL" )
            return _data
        }


        txn_receipt__contract_address = _txns["contractAddress"].(string)
        txn_receipt__from_address = _txns["from"].(string) // for gets token info (name, symbol, decimals, total_supply)
        //txn_receipt__to_address = txn_receipt__log["topics"].([]interface{})[2].(string) // will be received token

        if len(txn_receipt__log["topics"].([]interface{})) >= 4 {
            tx_token_type = "erc1155"

            txn_receipt__to_address = "0x" + txn_receipt__log["topics"].([]interface{})[3].(string)[2+24:] // will be received token
            {
                txn_receipt__token_symbol_id = txn_receipt__log["data"].(string)[2:64] // token symbol id

                token_id_int := new(big.Int)
                token_id_int.SetString( txn_receipt__token_symbol_id[2:], 16 )
                txn_receipt__token_symbol_id = token_id_int.String()

                //txn_receipt__token_symbol_id = txn_receipt__log["data"].(string)[2:64] // total supply
                txn_receipt__total_supply = txn_receipt__log["data"].(string)[2+64:] // total supply

                token_total_supply_int := new(big.Int)
                token_total_supply_int.SetString( txn_receipt__total_supply[2:], 16 )
                txn_receipt__total_supply = token_total_supply_int.String()
            }


            fmt.Println( TAG, "----------------------------------------------" )
            fmt.Println( TAG, "ERC-1155" )
            fmt.Println( TAG, "eth_getTransactionReceipt: contract address = ", txn_receipt__contract_address )
            fmt.Println( TAG, "eth_getTransactionReceipt: from address (owner) = ", txn_receipt__from_address )
            fmt.Println( TAG, "eth_getTransactionReceipt: to address (will be received token) = ", txn_receipt__to_address )
            fmt.Println( TAG, "eth_getTransactionReceipt: token symbol id  = ", txn_receipt__token_symbol_id )
            fmt.Println( TAG, "eth_getTransactionReceipt: total supply  = ", txn_receipt__total_supply )
            fmt.Println( TAG, "----------------------------------------------" )
        } else {
            tx_token_type = "erc20"

            {
                txn_receipt__total_supply = txn_receipt__log["data"].(string)[2+24:] // total supply

                token_total_supply_int := new(big.Int)
                token_total_supply_int.SetString( txn_receipt__total_supply[2:], 16 )
                txn_receipt__total_supply = token_total_supply_int.String()
            }
        }




        //! FIXME
        tx_block_number_hex := _txns["blockNumber"]
        {
            block_number_int := new(big.Int)
            block_number_int.SetString( tx_block_number_hex.(string)[2:], 16 )
            //fmt.Println( "ether hex-string to int: ", block_number_int )
            tx_block_number = block_number_int.String()
        }

        /*
        //! FIXME
        //timestamp_hex := result_block.Result.(map[string]interface{})["timestamp"]
        timestamp_hex := _txns["timestamp"]
        timestamp_int := new(big.Int)
        timestamp_int.SetString( timestamp_hex.(string)[2:], 16 )
        timestamp_unixtime = timestamp_int.String()
        //timestamp_int32, _ := strconv.Atoi( timestamp_unixtime )
        timestamp_int64, _ := strconv.ParseInt( timestamp_unixtime, 10, 64 )
        tx_timestamp_date = time.Unix( timestamp_int64, 0 )
        //tx_timestamp_date_rfc3339 := timestamp_date.Format( time.RFC3339 )
        */




        // get token_type
        /*
        {
            // Token: symbol
            method := "0x95d89b41"
            //data = method + "000000000000000000000000" + tx_from.(string)[2:]
            data := method + "000000000000000000000000" + txn_receipt__from_address[2:]

            var params []interface{}
            //request_data_param := types.RequestData_params_erc20 { To: tx_to.(string), Data: data }
            request_data_param := types.RequestData_params_erc20 { To: txn_receipt__contract_address, Data: data }
            params = append( params, request_data_param, "latest" )
            request_data := types.RequestData { Jsonrpc: "2.0", Method: "eth_call", Params: params, Id: 0 }

            message, _ := json.Marshal( request_data )
            //fmt.Println( "message: ", request_data )
            response, err := http.Post( URL, "application/json", bytes.NewBuffer(message) )
            defer response.Body.Close()
            if err != nil {
                log.Fatal( "http.Post: ", err )
            }

            //fmt.Println( "response: " )
            responseBody, err := ioutil.ReadAll( response.Body )
            if err != nil {
                //log.Fatal( "ioutil.ReadAll: ", err )
                fmt.Println( "ioutil.ReadAll: ", err )
            }

            //fmt.Println( string(responseBody) )
            err = json.Unmarshal( responseBody, &result )
            if err != nil {
                //log.Fatal( "json.Unmarshal: ", err )
                fmt.Println( "json.Unmarshal: ", err )
            }
            fmt.Println( "jsonrpc =" , result.Jsonrpc, ", id =", result.Id, ", result =", result.Result )
            if len(result.Result) > 0 {
                tx_token_type = "erc20"
            } else {
                tx_token_type = "erc1155"
            }
        }
        */

        if tx_token_type == "erc20" {
            // ERC-20
            fmt.Println( TAG, "ERC-20 token creation transaction" )
        } else if tx_token_type == "erc1155" {
            // ERC-1155
            fmt.Println( TAG, "ERC-1155 token creation transaction" )
        } else {
            fmt.Println( TAG, "Error..." )
            fmt.Println()
            return _data
        }





        method := ""
        data := ""



        tx_from := txn_receipt__from_address//txn["from"] // token creation: 0x0000000000000000000000000000000000000000
        tx_to := txn_receipt__contract_address//txn["to"]
        //tx_value_wei_hex := txn["value"] // ETHER
        //tx_value_wei := "" // ETHER
        //tx_value := "" // Ether
        //tx_input := txn["input"]

        tx_token_from := "" // for ERC-1155
        tx_token_to := "" // for ERC-20, ERC-1155

        // find token_type
        //tx_token_type := ""

        // for ERC-20
        tx_token_name := ""
        tx_token_symbol := ""
        tx_token_decimals := ""
        tx_token_total_supply := ""
        //tx_token_amount_wei_hex := ""
        tx_token_amount_wei := ""

        tx_token_amount := ""

        // for ERC-1155
        //tx_token_amount_hex := ""
        //tx_token_amount := ""
        tx_token_id_hex := ""
        //tx_token_id := ""
        tx_token_data := ""
        tx_token_data_length := ""
        tx_token_uri_with_token_id := ""
        tx_token_uri_with_token_id_hexadecimal := ""




        if tx_token_type == "erc20" {
            _token_name := ""
            _token_symbol := ""
            _token_decimals := ""
            _token_total_supply := ""


            // token to: [2: 0x] + [8: method] + [0 x 24]
            //tx_token_to = "0x" + tx_input.(string)[2 + 8 + 24:(2+8+24 + 40)]
            tx_token_to = txn_receipt__contract_address//txn["to"]

            // amount: 32 bytes (64 chars): [2: 0x] + [8: method] + [0 x 24] + [40: to address]
            //tx_token_amount_wei_hex = "0x" + tx_input.(string)[2 + 8 + 24 + 40:]
            //tx_token_amount_wei_hex = txn_receipt__total_supply


            {
                // Token: name
                method = "0x06fdde03"
                //data = method + "000000000000000000000000" + tx_from.(string)[2:]
                data = method + "000000000000000000000000" + tx_from[2:]

                var params []interface{}
                //request_data_param := types.RequestData_params_erc20 { To: tx_to.(string), Data: data }
                request_data_param := types.RequestData_params_erc20 { To: tx_to, Data: data }
                params = append( params, request_data_param, "latest" )
                request_data := types.RequestData { Jsonrpc: "2.0", Method: "eth_call", Params: params, Id: 0 }

                message, _ := json.Marshal( request_data )
                //fmt.Println( "message: ", request_data )
                response, err := http.Post( URL, "application/json", bytes.NewBuffer(message) )
                defer response.Body.Close()
                if err != nil {
                    log.Fatal( "http.Post: ", err )
                }

                //fmt.Println( "response: " )
                responseBody, err := ioutil.ReadAll( response.Body )
                if err != nil {
                    log.Fatal( "ioutil.ReadAll: ", err )
                }

                //fmt.Println( string(responseBody) )
                err = json.Unmarshal( responseBody, &result )
                if err != nil {
                    log.Fatal( "json.Unmarshal: ", err )
                }
                //fmt.Println( "jsonrpc =" , result.Jsonrpc, ", id =", result.Id, ", result =", result.Result )
                _token_name = result.Result
            }

            {
                // Token: symbol
                method = "0x95d89b41"
                //data = method + "000000000000000000000000" + tx_from.(string)[2:]
                data = method + "000000000000000000000000" + tx_from[2:]

                var params []interface{}
                //request_data_param := types.RequestData_params_erc20 { To: tx_to.(string), Data: data }
                request_data_param := types.RequestData_params_erc20 { To: tx_to, Data: data }
                params = append( params, request_data_param, "latest" )
                request_data := types.RequestData { Jsonrpc: "2.0", Method: "eth_call", Params: params, Id: 0 }

                message, _ := json.Marshal( request_data )
                //fmt.Println( "message: ", request_data )
                response, err := http.Post( URL, "application/json", bytes.NewBuffer(message) )
                defer response.Body.Close()
                if err != nil {
                    log.Fatal( "http.Post: ", err )
                }

                //fmt.Println( "response: " )
                responseBody, err := ioutil.ReadAll( response.Body )
                if err != nil {
                    log.Fatal( "ioutil.ReadAll: ", err )
                }

                //fmt.Println( string(responseBody) )
                err = json.Unmarshal( responseBody, &result )
                if err != nil {
                    log.Fatal( "json.Unmarshal: ", err )
                }
                //fmt.Println( "jsonrpc =" , result.Jsonrpc, ", id =", result.Id, ", result =", result.Result )
                _token_symbol = result.Result
            }

            {
                // Token: decimals
                method = "0x313ce567"
                //data = method + "000000000000000000000000" + tx_from.(string)[2:]
                data = method + "000000000000000000000000" + tx_from[2:]

                var params []interface{}
                //request_data_param := types.RequestData_params_erc20 { To: tx_to.(string), Data: data }
                request_data_param := types.RequestData_params_erc20 { To: tx_to, Data: data }
                params = append( params, request_data_param, "latest" )
                request_data := types.RequestData { Jsonrpc: "2.0", Method: "eth_call", Params: params, Id: 0 }

                message, _ := json.Marshal( request_data )
                //fmt.Println( "message: ", request_data )
                response, err := http.Post( URL, "application/json", bytes.NewBuffer(message) )
                defer response.Body.Close()
                if err != nil {
                    log.Fatal( "http.Post: ", err )
                }

                //fmt.Println( "response: " )
                responseBody, err := ioutil.ReadAll( response.Body )
                if err != nil {
                    log.Fatal( "ioutil.ReadAll: ", err )
                }

                //fmt.Println( string(responseBody) )
                err = json.Unmarshal( responseBody, &result )
                if err != nil {
                    log.Fatal( "json.Unmarshal: ", err )
                }
                //fmt.Println( "jsonrpc =" , result.Jsonrpc, ", id =", result.Id, ", result =", result.Result )
                _token_decimals = result.Result
            }

            {
                // Token: total_supply 
                method = "0x18160ddd"
                //data = method + "000000000000000000000000" + tx_from.(string)[2:]
                data = method + "000000000000000000000000" + tx_from[2:]

                var params []interface{}
                //request_data_param := types.RequestData_params_erc20 { To: tx_to.(string), Data: data }
                request_data_param := types.RequestData_params_erc20 { To: tx_to, Data: data }
                params = append( params, request_data_param, "latest" )
                request_data := types.RequestData { Jsonrpc: "2.0", Method: "eth_call", Params: params, Id: 0 }

                message, _ := json.Marshal( request_data )
                //fmt.Println( "message: ", request_data )
                response, err := http.Post( URL, "application/json", bytes.NewBuffer(message) )
                defer response.Body.Close()
                if err != nil {
                    log.Fatal( "http.Post: ", err )
                }

                //fmt.Println( "response: " )
                responseBody, err := ioutil.ReadAll( response.Body )
                if err != nil {
                    log.Fatal( "ioutil.ReadAll: ", err )
                }

                //fmt.Println( string(responseBody) )
                err = json.Unmarshal( responseBody, &result )
                if err != nil {
                    log.Fatal( "json.Unmarshal: ", err )
                }
                //fmt.Println( "jsonrpc =" , result.Jsonrpc, ", id =", result.Id, ", result =", result.Result )
                _token_total_supply = result.Result
            }


            //-----{
            // token name: 0x + [60 bytes] + [4 bytes] + [60 bytes] + [4 bytes]:
            __token_name, _ := hex.DecodeString( _token_name[2 + 60 + 4 + 60 + 4:] )
            tx_token_name = string(bytes.Trim(__token_name, "\x00"))

            // token symbol: 0x + [60 bytes] + [4 bytes] + [60 bytes] + [4 bytes]:
            __token_symbol, _ := hex.DecodeString( _token_symbol[2 + 60 + 4 + 60 + 4:] )
            tx_token_symbol = string(bytes.Trim(__token_symbol, "\x00"))

            // token decimals: 0x + [60 bytes] + [4 bytes]
            token_decimals_int := new(big.Int)
            token_decimals_int.SetString( _token_decimals[2:], 16 )
            __token_decimals := token_decimals_int.String()
            token_decimals_int32, _ := strconv.Atoi( __token_decimals )
            tx_token_decimals = __token_decimals

            // token total supply:
            token_total_supply_int := new(big.Int)
            token_total_supply_int.SetString( _token_total_supply[2:], 16 )
            token_total_supply_float := new(big.Float)
            token_total_supply_float.SetString( token_total_supply_int.String() )
            __token_total_supply := new(big.Float).Quo(token_total_supply_float, big.NewFloat(math.Pow10(token_decimals_int32)))
            tx_token_total_supply = fmt.Sprintf( "%.8f", __token_total_supply )

            //fmt.Println( "token name:", string(__token_name) )
            //fmt.Println( "token_symbol:", string(__token_symbol) )
            //fmt.Println( "token_decimals:", __token_decimals )
            //fmt.Printf( "token_total_supply: %f\n", __token_total_supply )
            //-----}


            // SEE:
            // - https://golang.org/pkg/math/big/
            // - https://golang.org/pkg/strconv/
            // - https://goethereumbook.org/account-balance/
            token_amount_wei_int := new(big.Int)
            //token_amount_wei_int.SetString( tx_token_amount_wei_hex[2:], 16 )
            token_amount_wei_int.SetString( txn_receipt__total_supply, 10 ) // total supply to owner
            //fmt.Println( "erc-20 token amount hex-string to int: ", token_amount_wei_int, "(wei)" )
            token_amount_wei_float := new(big.Float)
            token_amount_wei_float.SetString( token_amount_wei_int.String() )
            token_amount := new(big.Float).Quo(token_amount_wei_float, big.NewFloat(math.Pow10(token_decimals_int32)))
            tx_token_amount = fmt.Sprintf( "%.8f", token_amount )
            tx_token_amount_wei = token_amount_wei_int.String()
            //fmt.Printf( "erc-20 token amount: %s (%s)\n", tx_token_amount, tx_token_symbol )

            fmt.Println( "token_contract address =", tx_to )
            fmt.Println( "token_from =", tx_from )
            fmt.Println( "token_to =", tx_token_to )
            fmt.Println( "token_name =", tx_token_name )
            fmt.Println( "token_symbol =", tx_token_symbol )
            fmt.Println( "token_decimals =", tx_token_decimals )
            fmt.Println( "token_total_supply =", tx_token_total_supply )
            fmt.Println( "token_value_wei =", tx_token_amount_wei, "(wei)" )
            fmt.Println( "token_value_" + tx_token_symbol + " =", tx_token_amount, "(" + tx_token_symbol + ")" ) // total supply to owner
            fmt.Println()

            ///*
            __data := types.Fetch_transactions_st {
                //Symbol: tx_token_symbol,
                //From_address: tx_from.(string),
                From_address: "0x0000000000000000000000000000000000000000", //tx_from, // token creation: 0x0000000000000000000000000000000000000000
                To_address: tx_token_to,

                //Is_send: "",

                // for SQL: ETHER
                Amount_wei: "",
                Amount_eth: "",

                // for SQL: ERC-20
                Token_type: tx_token_type,
                Token_symbol: tx_token_symbol,
                Token_decimals: tx_token_decimals,
                Token_total_supply: tx_token_total_supply,
                //Token_contract_address: tx_to.(string),
                Token_contract_address: tx_to,
                Token_amount_wei: tx_token_amount_wei,
                Token_amount_eth: tx_token_amount,

                // for SQL: ERC-1155
                Token_amount: "",
                Token_uri_ascii: "",
                Token_uri_hexadecimal: "",
                Token_data_length: "",
                Token_data: "",

                Timestamp: timestamp_unixtime,
                Datetime: fmt.Sprintf("%s", tx_timestamp_date),
                Block_number: tx_block_number,
                //Txid: tx_hash.(string),
                Txid: tx_hash,
            }
            return __data
            //*/
        } else if tx_token_type == "erc1155" {
            tx_from = "0x0000000000000000000000000000000000000000" //txn_receipt__from_address//txn["from"] // token creation: 0x0000000000000000000000000000000000000000
            tx_to = txn_receipt__contract_address//txn["to"]
            tx_token_from = "0x0000000000000000000000000000000000000000" //txn_receipt__from_address // for ERC-1155
            tx_token_to = txn_receipt__to_address // for ERC-20, ERC-1155
            //tx_token_id := txn_receipt__token_symbol_id

            tx_token_total_supply := txn_receipt__total_supply
            tx_token_amount := txn_receipt__total_supply


            //token_amount_int := new(big.Int)
            //token_amount_int.SetString( tx_token_amount_hex[2:], 16 )
            //tx_token_amount = token_amount_int.String()

            //token_id_int := new(big.Int)
            //token_id_int.SetString( tx_token_id_hex[2:], 16 )
            //tx_token_id := token_id_int.String()

            // get URI
            /*
            {
                var result types.Result

                //gas := "70000"
                //gasprice := "100"
                //value := ""
                //from := ""
                to := tx_to // erc-1155 contract address
                holder_address := tx_token_to // has been transferred already
                token_id := tx_token_id
                //to := "0x1e57f9561600b269a37437f02ce9da31e5b830ce" // erc-1155 contract address
                //holder_address := "0xe6e55eed00218faef27eed24def9208f3878b333"
                method := "0x0e89341c"
                token_id_int := new(big.Int)
                token_id_int.SetString( token_id, 10 )
                token_id_hex := "0x" + token_id_int.Text( 16 )
                data := method + "000000000000000000000000" + holder_address[2:] +
                        strings.Repeat("0", 64 - len(token_id_hex[2:])) + token_id_hex[2:]

                var params []interface{}
                //request_data_param := types.RequestData_params_erc1155_transaction { From: from, To: to, Value: value, Gas: gas, Gasprice: gasprice, Data: data }
                //request_data_param := types.RequestData_params_erc1155 { To: to.(string), Data: data }
                request_data_param := types.RequestData_params_erc1155 { To: to, Data: data }
                params = append( params, request_data_param, "latest" )
                request_data := types.RequestData { Jsonrpc: "2.0", Method: "eth_call", Params: params, Id: 0 }

                message, _ := json.Marshal( request_data )
                //fmt.Println( "message: ", request_data )
                response, err := http.Post( URL, "application/json", bytes.NewBuffer(message) )
                defer response.Body.Close()
                if err != nil {
                    log.Fatal( "http.Post: ", err )
                }

                //fmt.Println( "response: " )
                responseBody, err := ioutil.ReadAll( response.Body )
                if err != nil {
                    log.Fatal( "ioutil.ReadAll: ", err )
                }

                //fmt.Println( string(responseBody) )
                err = json.Unmarshal( responseBody, &result )
                if err != nil {
                    log.Fatal( "json.Unmarshal: ", err )
                }
                //fmt.Println( "jsonrpc =" , result.Jsonrpc, ", id =", result.Id, ", result =", result.Result )


                uri_hex_str := result.Result

                // 2+126+2+6: 2 bytes (0x) + 126 bytes (0000...20...000000) + 2 byte (3d) + 6 bytes (000000)
                checks_len := 2 + 126 + 2 + 6

                //fmt.Println( "checks len: ", checks_len )

                if len(uri_hex_str) <= checks_len {
                    log.Fatal( "invalid length: ", len(uri_hex_str) )
                }

                //fmt.Println( "uri hex str: ", uri_hex_str )
                //fmt.Println( "url hex str len: ", len(uri_hex_str) )

                uri_hex := uri_hex_str[2+126+2:len(uri_hex_str) - 6]
                //fmt.Println( "URI hex: ", uri_hex )

                uri_bs, err := hex.DecodeString( uri_hex )
                uri_str := string( uri_bs )

                //fmt.Println( hex.Dump(uri_bs) )

                if err != nil {
                    panic(err)
                }
                //fmt.Println( "erc-1155 URI hex-string to str: ", uri_str )

                uri_with_token_id := strings.Replace( uri_str, "{id}", token_id, -1 )
                //fmt.Println( "erc-1155 URI: ", uri_with_token_id )


                // https://docs.openzeppelin.com/contracts/3.x/erc1155#constructing_an_erc1155_token_contract
                // The uri can include the string {id} which clients must replace with the actual token ID,
                // in lowercase hexadecimal (with no 0x prefix) and leading zero padded to 64 hex characters.

                //token_id_bytes := []byte( token_id ) // from str
                //token_id_hex = hex.EncodeToString( token_id_bytes )
                //
                //token_id_hex = hex.EncodeToString( []byte(token_id) ) // from str

                //token_id_bytes := []byte( strconv.FormatInt(token_id_int, 16) ) // from int
                token_id_bytes := []byte( fmt.Sprintf("%x", token_id_int) ) // from int
                token_id_hex = hex.EncodeToString( token_id_bytes )

                //fmt.Println( "token_id str: ", token_id )
                //fmt.Println( "token_id hex (from str literally): ", token_id_hex )

                token_id_bs, err := hex.DecodeString( token_id_hex )
                token_id_str := string( token_id_bs )
                if err != nil {
                    panic(err)
                }
                //fmt.Println( "token_id ASCII: ", token_id_str )
                //tx_token_uri_with_token_id = token_id_str


                uri_with_token_id = strings.Repeat("0", 64 - len(token_id_hex)) + token_id_hex // from str literally
                //uri_with_token_id = strings.Repeat("0", 64 - len(token_id_str)) + token_id_str // ASCII
                uri_with_token_id = strings.Replace( uri_str, "{id}", uri_with_token_id, -1 )
                //fmt.Println( "erc-1155 URI: ", uri_with_token_id )

                // Hexadecimal
                tx_token_uri_with_token_id_hexadecimal = uri_with_token_id

                // ASCII
                tx_token_uri_with_token_id = strings.Replace( uri_str, "{id}", token_id_str, -1 )
            }
            //*/

            fmt.Println( "token_contract address =", tx_to )
            fmt.Println( "token_from =", tx_token_from )
            fmt.Println( "token_to =", tx_token_to )
            fmt.Println( "token_id = ", tx_token_id_hex )
            fmt.Println( "token_total_supply = ", tx_token_total_supply )
            fmt.Println( "token_amount = ", tx_token_amount ) // total supply to owner
            fmt.Println( "token_uri (ASCII) = ", tx_token_uri_with_token_id )
            fmt.Println( "token_uri (Hexadecimal) = ", tx_token_uri_with_token_id_hexadecimal )
            fmt.Println( "token_data_length = ", tx_token_data_length )
            fmt.Println( "token_data = ", tx_token_data )
            fmt.Println()

            ///*
            __data := types.Fetch_transactions_st {
                //Symbol: tx_token_symbol,
                //From_address: tx_from.(string),
                From_address: "0x0000000000000000000000000000000000000000", //tx_from,
                To_address: tx_token_to,

                //Is_send: "",

                // for SQL: ETHER
                Amount_wei: "",
                Amount_eth: "",

                Token_type: tx_token_type,

                // for SQL: ERC-20
                Token_symbol: tx_token_symbol,
                Token_decimals: tx_token_decimals,
                Token_total_supply: tx_token_total_supply,

                //Token_contract_address: tx_to.(string),
                Token_contract_address: tx_to,

                // for SQL: ERC-20
                Token_amount_wei: "",
                Token_amount_eth: "",

                // for SQL: ERC-1155
                Token_amount: tx_token_amount, // total supply to owner
                Token_uri_ascii: tx_token_uri_with_token_id,
                Token_uri_hexadecimal: tx_token_uri_with_token_id_hexadecimal,
                Token_data_length: tx_token_data_length,
                Token_data: tx_token_data,

                Timestamp: timestamp_unixtime,
                Datetime: fmt.Sprintf("%s", tx_timestamp_date),
                Block_number: tx_block_number,
                //Txid: tx_hash.(string),
                Txid: tx_hash,
            }
            return __data
            //*/
        }
    } // for ()

    return _data
}

// ETH, ERC-20 (transfer()), ERC-1155 (safeTransferOf())
//func get_blocks_all() {
func get_blocks_all_infinite(block_num_start_uint64 uint64, block_num_end_uint64 uint64) {
    TAG := "get_blocks_all_infinite(): "

    // eth_getBlockByNumber
    //
    // request:
    // $
    //

    fmt.Println( TAG, "eth_call(): eth_getBlockByNumber()" )

    /*
    //block_num_start_uint64 := uint64(502)
    //block_num_end_uint64 := uint64(503)
    block_num_start_uint64 := uint64(0)
    block_num_end_uint64 := uint64(0)

    {
        var result types.Result

        //var params []interface{}
        //request_data_param := types.RequestData {  }
        //params = append( params, "latest", true )
        request_data := types.RequestData { Jsonrpc: "2.0", Method: "eth_blockNumber", Id: 0 }

        message, _ := json.Marshal( request_data )
        //fmt.Println( "message: ", request_data )
        response, err := http.Post( URL, "application/json", bytes.NewBuffer(message) )
        defer response.Body.Close()
        if err != nil {
            log.Fatal( "http.Post: ", err )
        }

        //fmt.Println( "response: " )
        responseBody, err := ioutil.ReadAll( response.Body )
        if err != nil {
            log.Fatal( "ioutil.ReadAll: ", err )
        }

        //fmt.Println( string(responseBody) )
        err = json.Unmarshal( responseBody, &result )
        if err != nil {
            log.Fatal( "json.Unmarshal: ", err )
        }
        //fmt.Println( "jsonrpc =" , result.Jsonrpc, ", id =", result.Id, ", result =", result.Result )

        block_num_end_int := new(big.Int)
        block_num_end_int.SetString( result.Result[2:], 16 )
        block_num_end := block_num_end_int.String()
        //block_num_end_int32, _ = strconv.Atoi( block_num_end )
        block_num_end_uint64, _ = strconv.ParseUint( block_num_end, 10, 64 )
    }
    */

    {
        /*
        type Result struct {
            Jsonrpc string `json:"jsonrpc"`
            Id int `json:"id"`
            Result string `json:"result"`
        }
        var result Result

        type Result_block struct {
            Jsonrpc string `json:"jsonrpc"`
            Id int `json:"id"`
            Result interface{} `json:"result"`
        }
        var result_block Result_block
        */

        var result types.Result
        var result_block types.Result_block




        fmt.Println( TAG, "block start = ", block_num_start_uint64, "block end = ", block_num_end_uint64 )
        //fmt.Println()
        for i := block_num_start_uint64; i < uint64(block_num_end_uint64); i++ {
            //fmt.Println( "===== TODO =====" )
            //fmt.Println( "Token creation transaction: 0x60606040 or 0x60806040" )
            //fmt.Println( "Token creation transaction: //! TODO: get total supply, owner address" )

            request_block_num_hex := ""
            request_block_num_int := new(big.Int)
            request_block_num_int.SetUint64( uint64(i) )
            request_block_num_hex = "0x" + request_block_num_int.Text( 16 )

            //fmt.Println( TAG, i, "(" + request_block_num_hex + ")", "----------" )

            var params []interface{}
            //request_data_param := types.RequestData {  }
            //params = append( params, "latest", true )
            params = append( params, request_block_num_hex, true )
            request_data := types.RequestData { Jsonrpc: "2.0", Method: "eth_getBlockByNumber", Params: params, Id: 0 }

            message, _ := json.Marshal( request_data )
            //fmt.Println( TAG, "message: ", request_data )
            response, err := http.Post( URL, "application/json", bytes.NewBuffer(message) )
            defer response.Body.Close()
            if err != nil {
                log.Fatal( "http.Post: ", err )
            }

            //fmt.Println( TAG, "response: " )
            responseBody, err := ioutil.ReadAll( response.Body )
            if err != nil {
                log.Fatal( "ioutil.ReadAll: ", err )
            }

            //fmt.Println( TAG, string(responseBody) )
            err = json.Unmarshal( responseBody, &result_block )
            if err != nil {
                log.Fatal( "json.Unmarshal: ", err )
            }
            //fmt.Println( TAG, "jsonrpc =" , result.Jsonrpc, ", id =", result.Id, ", result =", result_block.Result )


            //fmt.Println( TAG, "data =", result_block.Result )
            //fmt.Println( TAG, reflect.TypeOf(result_block.Result) )
            _txns := result_block.Result.(map[string]interface{})["transactions"]
            fmt.Println( TAG, "transaction size = ", len(_txns.([]interface{})) )
            if len(_txns.([]interface{})) <= 0 {
                //fmt.Println( TAG, "no transaction: size = 0" )
                continue
            }
            fmt.Println()


            // block info
            {
                block_number := request_block_num_int.Text( 10 )
                block_hash := result_block.Result.(map[string]interface{})["hash"].(string)
                block_info, err_marshal := json.Marshal( result_block.Result )
                if err_marshal != nil {
                    panic( err_marshal.Error() )
                }
                //fmt.Println( string(block_info) )
                num_of_txns := len( _txns.([]interface{}) )

                db_insert_blocks_info( block_number, string(block_hash), string(block_info), num_of_txns )
            }
            fmt.Println()


            for j := 0; j < len(_txns.([]interface{})); j++ {
                //_txns = _txns.([]interface{})[0]
                //txn := _txns.(map[string]interface{})
                _txn := _txns.([]interface{})[j]
                txn := _txn.(map[string]interface{})

                //if txn["from"] != "0x8f5b2b7608e3e3a3dc0426c3396420fbf1849454" && txn["to"] != "0x8f5b2b7608e3e3a3dc0426c3396420fbf1849454" {
                //if txn["from"] != "0xe6e55eed00218faef27eed24def9208f3878b333" && txn["to"] != "0x1249CDA86774Bc170CAb843437DD37484F173ca8" {
                //    continue
                //}

                timestamp_hex := result_block.Result.(map[string]interface{})["timestamp"]
                timestamp_int := new(big.Int)
                timestamp_int.SetString( timestamp_hex.(string)[2:], 16 )
                timestamp_unixtime := timestamp_int.String()
                //timestamp_int32, _ := strconv.Atoi( timestamp_unixtime )
                timestamp_int64, _ := strconv.ParseInt( timestamp_unixtime, 10, 64 )
                tx_timestamp_date := time.Unix( timestamp_int64, 0 )
                //tx_timestamp_date_rfc3339 := timestamp_date.Format( time.RFC3339 )


                tx_hash := txn["hash"]
                tx_block_number_hex := txn["blockNumber"]
                tx_block_number := ""
                {
                    block_number_int := new(big.Int)
                    block_number_int.SetString( tx_block_number_hex.(string)[2:], 16 )
                    //fmt.Println( "ether hex-string to int: ", block_number_int )
                    tx_block_number = block_number_int.String()
                }
                tx_from := txn["from"]
                tx_to := txn["to"]
                tx_value_wei_hex := txn["value"]
                tx_value_wei := ""
                tx_value := "" // Ether
                tx_input := txn["input"]

                tx_token_from := "" // for ERC-1155
                tx_token_to := "" // for ERC-20, ERC-1155

                tx_token_type := ""

                // for ERC-20
                tx_token_name := ""
                tx_token_symbol := ""
                tx_token_decimals := ""
                tx_token_total_supply := ""
                tx_token_amount_wei_hex := ""
                tx_token_amount_wei := ""

                tx_token_amount := ""

                // for ERC-1155
                tx_token_amount_hex := ""
                //tx_token_amount := ""
                tx_token_id_hex := ""
                //tx_token_id := ""
                tx_token_data := ""
                tx_token_data_length := ""
                tx_token_uri_with_token_id := ""
                tx_token_uri_with_token_id_hexadecimal := ""

                if tx_to == nil {
                    //continue

                    //fmt.Println( "Token creation transaction: 0x60606040 or 0x60806040" )
                    //fmt.Println( "Token creation transaction: //! TODO: get total supply, owner address" )
                }

                //fmt.Println( "transaction =", _txn )
                //fmt.Println( "hash =", tx_hash )
                //fmt.Println( "timestamp =", tx_timestamp_date ) // "Y/m/d/ H:i:s"
                //fmt.Println( "block_number =", tx_block_number )
                //fmt.Println( "from =", tx_from )

                if txn["input"] == "0x" {
                    // ETH
                    fmt.Println( "Ether" )
                    tx_token_type = "ether"

                    fmt.Println( "hash =", tx_hash )
                    fmt.Println( "timestamp =", timestamp_unixtime )
                    fmt.Println( "datetime =", tx_timestamp_date ) // "Y/m/d/ H:i:s"
                    fmt.Println( "block_number =", tx_block_number )
                    fmt.Println( "from =", tx_from )
                    fmt.Println( "to =", tx_to )

                    // SEE:
                    // - https://golang.org/pkg/math/big/
                    // - https://golang.org/pkg/strconv/
                    // - https://goethereumbook.org/account-balance/
                    amount_wei_int := new(big.Int)
                    amount_wei_int.SetString( tx_value_wei_hex.(string)[2:], 16 )
                    fmt.Println( "ether hex-string to int: ", amount_wei_int, "(wei)" )
                    amount_wei_float := new(big.Float)
                    amount_wei_float.SetString( amount_wei_int.String() )
                    tx_value_float := new(big.Float).Quo(amount_wei_float, big.NewFloat(math.Pow10(18)))
                    tx_value = fmt.Sprintf( "%.8f", tx_value_float )
                    tx_value_wei = amount_wei_int.String()

                    fmt.Println( "value_wei =", tx_value_wei, "(wei)" )
                    fmt.Println( "value_ether =", tx_value, "(ether)" )
                    fmt.Println()
                } else {
                    //fmt.Println( "ERC-xxxx" )

                    tx_token_type = ""
                    //tx_data_offset := 0


                    //fmt.Println( "method =", tx_input.(string)[:10] )

                    if tx_input.(string)[:10] == "0xa9059cbb" {
                        // ERC-20
                        fmt.Println( "ERC-20 transfer() transaction" )
                        tx_token_type = "erc20"
                    } else if tx_input.(string)[:10] == "0xf242432a" {
                        // ERC-1155
                        fmt.Println( "ERC-1155 safeTransferFrom() transaction" )
                        tx_token_type = "erc1155"
                    } else if tx_input.(string)[:10] == "0x60606040" {
                        // Token creation
                        fmt.Println( "===> found transaction 0x60606040 or 0x60806040 #1" )
                        //fmt.Println( "Token creation transaction: 0x60606040 or 0x60806040" )
                        //fmt.Println( "Token creation transaction: //! TODO: get total supply, owner address" )
                        tx_token_type = "erc_token_creation"
                        //continue
                    } else if tx_input.(string)[:10] == "0x60806040" {
                        // Token creation
                        fmt.Println( "===> found transaction 0x60606040 or 0x60806040 #1" )
                        //fmt.Println( "Token creation transaction: 0x60606040 or 0x60806040" )
                        //fmt.Println( "Token creation transaction: //! TODO: get total supply, owner address" )
                        tx_token_type = "erc_token_creation"
                        //continue
                    } else {

                        //fmt.Println( "ERC-xxxx transfer transaction: No such transaction method" )
                        //fmt.Println()
                        continue
                    }



                    fmt.Println( "input data =", tx_input )

                    if tx_token_type == "erc1155" {
                        //tx_data_offset = 2 + 8 + 24
                        // token from: 32 bytes (64 chars): [2: 0x] + [8: method] + [0 x 24] ~ [40]
                        tx_token_from = "0x" + tx_input.(string)[2 + 8 + 24:(2+8+24 + 40)]

                        //tx_data_offset = 2 + 8 + 24+40 + 24
                        // token to: 32 bytes (64 chars): [2: 0x] + [8: method] + [0 x 24] ~ [40]
                        tx_token_to = "0x" + tx_input.(string)[2 + 8 + 24+40 + 24:(2+8+24+40 + 64)]

                        //tx_data_offset = 2 + 8 + 64 + 64
                        // token id: 32 bytes (64 chars): [2: 0x] + [8: method] + [64] + [64]
                        tx_token_id_hex = "0x" + tx_input.(string)[2 + 8 + 64 + 64:(2+8+64+64 + 64)]

                        //tx_data_offset = 2 + 8 + 64 + 64 + 64
                        // amount: 32 bytes (64 chars): [2: 0x] + [8: method] + [64] + [64] + [64]
                        tx_token_amount_hex = "0x" + tx_input.(string)[2 + 8 + 64 + 64 + 64:(2+8+64+64+64 + 64)]

                        //tx_data_offset = 2 + 8 + 64 + 64 + 64 + 64
                        // data length bytes: 32 bytes (64 chars): [2: 0x] + [8: method] + [64] + [64] + [64] + [64]
                        tx_token_data_length = "0x" + tx_input.(string)[2 + 8 + 64 + 64 + 64 + 64:(2+8+64+64+64+64 + 64)]

                        //tx_data_offset = 2 + 8 + 64 + 64 + 64 + 64 + 64
                        // data length bytes: 32 bytes (64 chars): [2: 0x] + [8: method] + [64] + [64] + [64] + [64] + [64]
                        tx_token_data = "0x" + tx_input.(string)[2 + 8 + 64 + 64 + 64 + 64 + 64:(2+8+64+64+64+64+64 + 64)]
                    }



                    method := ""
                    data := ""

                    if tx_token_type == "erc20" {
                        _token_name := ""
                        _token_symbol := ""
                        _token_decimals := ""
                        _token_total_supply := ""


                        // token to: [2: 0x] + [8: method] + [0 x 24]
                        tx_token_to = "0x" + tx_input.(string)[2 + 8 + 24:(2+8+24 + 40)]

                        // amount: 32 bytes (64 chars): [2: 0x] + [8: method] + [0 x 24] + [40: to address]
                        tx_token_amount_wei_hex = "0x" + tx_input.(string)[2 + 8 + 24 + 40:]


                        {
                            // Token: name
                            method = "0x06fdde03"
                            data = method + "000000000000000000000000" + tx_from.(string)[2:]

                            var params []interface{}
                            request_data_param := types.RequestData_params_erc20 { To: tx_to.(string), Data: data }
                            params = append( params, request_data_param, "latest" )
                            request_data := types.RequestData { Jsonrpc: "2.0", Method: "eth_call", Params: params, Id: 0 }

                            message, _ := json.Marshal( request_data )
                            //fmt.Println( "message: ", request_data )
                            response, err := http.Post( URL, "application/json", bytes.NewBuffer(message) )
                            defer response.Body.Close()
                            if err != nil {
                                log.Fatal( "http.Post: ", err )
                            }

                            //fmt.Println( "response: " )
                            responseBody, err := ioutil.ReadAll( response.Body )
                            if err != nil {
                                log.Fatal( "ioutil.ReadAll: ", err )
                            }

                            //fmt.Println( string(responseBody) )
                            err = json.Unmarshal( responseBody, &result )
                            if err != nil {
                                log.Fatal( "json.Unmarshal: ", err )
                            }
                            //fmt.Println( "jsonrpc =" , result.Jsonrpc, ", id =", result.Id, ", result =", result.Result )
                            _token_name = result.Result
                        }

                        {
                            // Token: symbol
                            method = "0x95d89b41"
                            data = method + "000000000000000000000000" + tx_from.(string)[2:]

                            var params []interface{}
                            request_data_param := types.RequestData_params_erc20 { To: tx_to.(string), Data: data }
                            params = append( params, request_data_param, "latest" )
                            request_data := types.RequestData { Jsonrpc: "2.0", Method: "eth_call", Params: params, Id: 0 }

                            message, _ := json.Marshal( request_data )
                            //fmt.Println( "message: ", request_data )
                            response, err := http.Post( URL, "application/json", bytes.NewBuffer(message) )
                            defer response.Body.Close()
                            if err != nil {
                                log.Fatal( "http.Post: ", err )
                            }

                            //fmt.Println( "response: " )
                            responseBody, err := ioutil.ReadAll( response.Body )
                            if err != nil {
                                log.Fatal( "ioutil.ReadAll: ", err )
                            }

                            //fmt.Println( string(responseBody) )
                            err = json.Unmarshal( responseBody, &result )
                            if err != nil {
                                log.Fatal( "json.Unmarshal: ", err )
                            }
                            //fmt.Println( "jsonrpc =" , result.Jsonrpc, ", id =", result.Id, ", result =", result.Result )
                            _token_symbol = result.Result
                        }

                        {
                            // Token: decimals
                            method = "0x313ce567"
                            data = method + "000000000000000000000000" + tx_from.(string)[2:]

                            var params []interface{}
                            request_data_param := types.RequestData_params_erc20 { To: tx_to.(string), Data: data }
                            params = append( params, request_data_param, "latest" )
                            request_data := types.RequestData { Jsonrpc: "2.0", Method: "eth_call", Params: params, Id: 0 }

                            message, _ := json.Marshal( request_data )
                            //fmt.Println( "message: ", request_data )
                            response, err := http.Post( URL, "application/json", bytes.NewBuffer(message) )
                            defer response.Body.Close()
                            if err != nil {
                                log.Fatal( "http.Post: ", err )
                            }

                            //fmt.Println( "response: " )
                            responseBody, err := ioutil.ReadAll( response.Body )
                            if err != nil {
                                log.Fatal( "ioutil.ReadAll: ", err )
                            }

                            //fmt.Println( string(responseBody) )
                            err = json.Unmarshal( responseBody, &result )
                            if err != nil {
                                log.Fatal( "json.Unmarshal: ", err )
                            }
                            //fmt.Println( "jsonrpc =" , result.Jsonrpc, ", id =", result.Id, ", result =", result.Result )
                            _token_decimals = result.Result
                        }

                        {
                            // Token: total_supply 
                            method = "0x18160ddd"
                            data = method + "000000000000000000000000" + tx_from.(string)[2:]

                            var params []interface{}
                            request_data_param := types.RequestData_params_erc20 { To: tx_to.(string), Data: data }
                            params = append( params, request_data_param, "latest" )
                            request_data := types.RequestData { Jsonrpc: "2.0", Method: "eth_call", Params: params, Id: 0 }

                            message, _ := json.Marshal( request_data )
                            //fmt.Println( "message: ", request_data )
                            response, err := http.Post( URL, "application/json", bytes.NewBuffer(message) )
                            defer response.Body.Close()
                            if err != nil {
                                log.Fatal( "http.Post: ", err )
                            }

                            //fmt.Println( "response: " )
                            responseBody, err := ioutil.ReadAll( response.Body )
                            if err != nil {
                                log.Fatal( "ioutil.ReadAll: ", err )
                            }

                            //fmt.Println( string(responseBody) )
                            err = json.Unmarshal( responseBody, &result )
                            if err != nil {
                                log.Fatal( "json.Unmarshal: ", err )
                            }
                            //fmt.Println( "jsonrpc =" , result.Jsonrpc, ", id =", result.Id, ", result =", result.Result )
                            _token_total_supply = result.Result
                        }


                        //-----{
                        // token name: 0x + [60 bytes] + [4 bytes] + [60 bytes] + [4 bytes]:
                        __token_name, _ := hex.DecodeString( _token_name[2 + 60 + 4 + 60 + 4:] )
                        tx_token_name = string(bytes.Trim(__token_name, "\x00"))

                        // token symbol: 0x + [60 bytes] + [4 bytes] + [60 bytes] + [4 bytes]:
                        __token_symbol, _ := hex.DecodeString( _token_symbol[2 + 60 + 4 + 60 + 4:] )
                        tx_token_symbol = string(bytes.Trim(__token_symbol, "\x00"))

                        // token decimals: 0x + [60 bytes] + [4 bytes]
                        token_decimals_int := new(big.Int)
                        token_decimals_int.SetString( _token_decimals[2:], 16 )
                        __token_decimals := token_decimals_int.String()
                        token_decimals_int32, _ := strconv.Atoi( __token_decimals )
                        tx_token_decimals = __token_decimals

                        // token total supply:
                        token_total_supply_int := new(big.Int)
                        token_total_supply_int.SetString( _token_total_supply[2:], 16 )
                        token_total_supply_float := new(big.Float)
                        token_total_supply_float.SetString( token_total_supply_int.String() )
                        __token_total_supply := new(big.Float).Quo(token_total_supply_float, big.NewFloat(math.Pow10(token_decimals_int32)))
                        tx_token_total_supply = fmt.Sprintf( "%.8f", __token_total_supply )

                        //fmt.Println( "token name:", string(__token_name) )
                        //fmt.Println( "token_symbol:", string(__token_symbol) )
                        //fmt.Println( "token_decimals:", __token_decimals )
                        //fmt.Printf( "token_total_supply: %f\n", __token_total_supply )
                        //-----}


                        // SEE:
                        // - https://golang.org/pkg/math/big/
                        // - https://golang.org/pkg/strconv/
                        // - https://goethereumbook.org/account-balance/
                        token_amount_wei_int := new(big.Int)
                        token_amount_wei_int.SetString( tx_token_amount_wei_hex[2:], 16 )
                        //fmt.Println( "erc-20 token amount hex-string to int: ", token_amount_wei_int, "(wei)" )
                        token_amount_wei_float := new(big.Float)
                        token_amount_wei_float.SetString( token_amount_wei_int.String() )
                        token_amount := new(big.Float).Quo(token_amount_wei_float, big.NewFloat(math.Pow10(token_decimals_int32)))
                        tx_token_amount = fmt.Sprintf( "%.8f", token_amount )
                        tx_token_amount_wei = token_amount_wei_int.String()
                        //fmt.Printf( "erc-20 token amount: %s (%s)\n", tx_token_amount, tx_token_symbol )
                    } else if tx_token_type == "erc1155" {
                        token_amount_int := new(big.Int)
                        token_amount_int.SetString( tx_token_amount_hex[2:], 16 )
                        tx_token_amount = token_amount_int.String()

                        token_id_int := new(big.Int)
                        token_id_int.SetString( tx_token_id_hex[2:], 16 )
                        tx_token_id := token_id_int.String()

                        // get URI
                        {
                            var result types.Result

                            //gas := "70000"
                            //gasprice := "100"
                            //value := ""
                            //from := ""
                            to := tx_to // erc-1155 contract address
                            holder_address := tx_token_to // has been transferred already
                            token_id := tx_token_id
                            //to := "0x1e57f9561600b269a37437f02ce9da31e5b830ce" // erc-1155 contract address
                            //holder_address := "0xe6e55eed00218faef27eed24def9208f3878b333"
                            method := "0x0e89341c"
                            token_id_int := new(big.Int)
                            token_id_int.SetString( token_id, 10 )
                            token_id_hex := "0x" + token_id_int.Text( 16 )
                            data := method + "000000000000000000000000" + holder_address[2:] +
                                    strings.Repeat("0", 64 - len(token_id_hex[2:])) + token_id_hex[2:]

                            var params []interface{}
                            //request_data_param := types.RequestData_params_erc1155_transaction { From: from, To: to, Value: value, Gas: gas, Gasprice: gasprice, Data: data }
                            request_data_param := types.RequestData_params_erc1155 { To: to.(string), Data: data }
                            params = append( params, request_data_param, "latest" )
                            request_data := types.RequestData { Jsonrpc: "2.0", Method: "eth_call", Params: params, Id: 0 }

                            message, _ := json.Marshal( request_data )
                            //fmt.Println( "message: ", request_data )
                            response, err := http.Post( URL, "application/json", bytes.NewBuffer(message) )
                            defer response.Body.Close()
                            if err != nil {
                                log.Fatal( "http.Post: ", err )
                            }

                            //fmt.Println( "response: " )
                            responseBody, err := ioutil.ReadAll( response.Body )
                            if err != nil {
                                log.Fatal( "ioutil.ReadAll: ", err )
                            }

                            //fmt.Println( string(responseBody) )
                            err = json.Unmarshal( responseBody, &result )
                            if err != nil {
                                log.Fatal( "json.Unmarshal: ", err )
                            }
                            //fmt.Println( "jsonrpc =" , result.Jsonrpc, ", id =", result.Id, ", result =", result.Result )


                            uri_hex_str := result.Result

                            // 2+126+2+6: 2 bytes (0x) + 126 bytes (0000...20...000000) + 2 byte (3d) + 6 bytes (000000)
                            checks_len := 2 + 126 + 2 + 6

                            //fmt.Println( "checks len: ", checks_len )

                            if len(uri_hex_str) <= checks_len {
                		fmt.Println( "uri_hex_str = ", uri_hex_str )
                                log.Fatal( "invalid length: ", len(uri_hex_str) )
                            }

                            //fmt.Println( "uri hex str: ", uri_hex_str )
                            //fmt.Println( "url hex str len: ", len(uri_hex_str) )

                            uri_hex := uri_hex_str[2+126+2:len(uri_hex_str) - 6]
                            //fmt.Println( "URI hex: ", uri_hex )

                            uri_bs, err := hex.DecodeString( uri_hex )
                            uri_str := string( uri_bs )

                            //fmt.Println( hex.Dump(uri_bs) )

                            if err != nil {
                                panic(err)
                            }
                            //fmt.Println( "erc-1155 URI hex-string to str: ", uri_str )

                            uri_with_token_id := strings.Replace( uri_str, "{id}", token_id, -1 )
                            //fmt.Println( "erc-1155 URI: ", uri_with_token_id )


                            // https://docs.openzeppelin.com/contracts/3.x/erc1155#constructing_an_erc1155_token_contract
                            // The uri can include the string {id} which clients must replace with the actual token ID,
                            // in lowercase hexadecimal (with no 0x prefix) and leading zero padded to 64 hex characters.

                            //token_id_bytes := []byte( token_id ) // from str
                            //token_id_hex = hex.EncodeToString( token_id_bytes )
                            //
                            //token_id_hex = hex.EncodeToString( []byte(token_id) ) // from str

                            //token_id_bytes := []byte( strconv.FormatInt(token_id_int, 16) ) // from int
                            token_id_bytes := []byte( fmt.Sprintf("%x", token_id_int) ) // from int
                            token_id_hex = hex.EncodeToString( token_id_bytes )

                            //fmt.Println( "token_id str: ", token_id )
                            //fmt.Println( "token_id hex (from str literally): ", token_id_hex )

                            token_id_bs, err := hex.DecodeString( token_id_hex )
                            token_id_str := string( token_id_bs )
                            if err != nil {
                                panic(err)
                            }
                            //fmt.Println( "token_id ASCII: ", token_id_str )
                            //tx_token_uri_with_token_id = token_id_str


                            uri_with_token_id = strings.Repeat("0", 64 - len(token_id_hex)) + token_id_hex // from str literally
                            //uri_with_token_id = strings.Repeat("0", 64 - len(token_id_str)) + token_id_str // ASCII
                            uri_with_token_id = strings.Replace( uri_str, "{id}", uri_with_token_id, -1 )
                            //fmt.Println( "erc-1155 URI: ", uri_with_token_id )

                            // Hexadecimal
                            tx_token_uri_with_token_id_hexadecimal = uri_with_token_id

                            // ASCII
                            tx_token_uri_with_token_id = strings.Replace( uri_str, "{id}", token_id_str, -1 )
                        }
                    } else if tx_token_type == "erc_token_creation" {
                        // None...
                    }

                }


                //fmt.Println( "transaction =", _txn )
                fmt.Println( "hash =", tx_hash )
                fmt.Println( "timestamp =", timestamp_unixtime )
                fmt.Println( "datetime =", tx_timestamp_date ) // "Y/m/d/ H:i:s"
                fmt.Println( "block_number =", tx_block_number )
                fmt.Println( "from =", tx_from )

                if tx_token_type == "erc20" {
                    fmt.Println( "token_contract address =", tx_to )
                    fmt.Println( "token_to =", tx_token_to )
                    fmt.Println( "token_name =", tx_token_name )
                    fmt.Println( "token_symbol =", tx_token_symbol )
                    fmt.Println( "token_decimals =", tx_token_decimals )
                    fmt.Println( "token_total_supply =", tx_token_total_supply )
                    fmt.Println( "token_value_wei =", tx_token_amount_wei, "(wei)" )
                    fmt.Println( "token_value_" + tx_token_symbol + " =", tx_token_amount, "(" + tx_token_symbol + ")" )
                    fmt.Println()

                    _data := types.Fetch_transactions_st {
                        //Symbol: tx_token_symbol,
                        From_address: tx_from.(string),
                        To_address: tx_token_to,

                        //Is_send: "",

                        // for SQL: ETHER
                        Amount_wei: "",
                        Amount_eth: "",

                        // for SQL: ERC-20
                        Token_type: tx_token_type,
                        Token_symbol: tx_token_symbol,
                        Token_decimals: tx_token_decimals,
                        Token_total_supply: tx_token_total_supply,
                        Token_contract_address: tx_to.(string),
                        Token_amount_wei: tx_token_amount_wei,
                        Token_amount_eth: tx_token_amount,

                        // for SQL: ERC-1155
                        Token_amount: "",
                        Token_uri_ascii: "",
                        Token_uri_hexadecimal: "",
                        Token_data_length: "",
                        Token_data: "",

                        Timestamp: timestamp_unixtime,
                        Datetime: fmt.Sprintf("%s", tx_timestamp_date),
                        Block_number: tx_block_number,
                        Txid: tx_hash.(string),
                    }
                    db_insert_txns(_E_TYPE__ERC20, &_data )
                } else if tx_token_type == "erc1155" {
                    fmt.Println( "token_contract address =", tx_to )
                    fmt.Println( "token_from =", tx_token_from )
                    fmt.Println( "token_to =", tx_token_to )
                    fmt.Println( "token_id = ", tx_token_id_hex )
                    fmt.Println( "token_amount = ", tx_token_amount )
                    fmt.Println( "token_uri (ASCII) = ", tx_token_uri_with_token_id )
                    fmt.Println( "token_uri (Hexadecimal) = ", tx_token_uri_with_token_id_hexadecimal )
                    fmt.Println( "token_data_length = ", tx_token_data_length )
                    fmt.Println( "token_data = ", tx_token_data )
                    fmt.Println()

                    _data := types.Fetch_transactions_st {
                        //Symbol: tx_token_symbol,
                        From_address: tx_from.(string),
                        To_address: tx_token_to,

                        //Is_send: "",

                        // for SQL: ETHER
                        Amount_wei: "",
                        Amount_eth: "",

                        Token_type: tx_token_type,

                        // for SQL: ERC-20
                        Token_symbol: tx_token_symbol,
                        Token_decimals: tx_token_decimals,
                        Token_total_supply: tx_token_total_supply,

                        Token_contract_address: tx_to.(string),

                        // for SQL: ERC-20
                        Token_amount_wei: "",
                        Token_amount_eth: "",

                        // for SQL: ERC-1155
                        Token_amount: tx_token_amount,
                        Token_uri_ascii: tx_token_uri_with_token_id,
                        Token_uri_hexadecimal: tx_token_uri_with_token_id_hexadecimal,
                        Token_data_length: tx_token_data_length,
                        Token_data: tx_token_data,

                        Timestamp: timestamp_unixtime,
                        Datetime: fmt.Sprintf("%s", tx_timestamp_date),
                        Block_number: tx_block_number,
                        Txid: tx_hash.(string),
                    }
                    db_insert_txns(_E_TYPE__ERC1155, &_data )
                } else if tx_token_type == "erc_token_creation" {
                    fmt.Println( "===> found transaction 0x60606040 or 0x60806040 #2" )
                    //fmt.Println( "Token creation transaction: 0x60606040 or 0x60806040" )
                    //fmt.Println( "Token creation transaction: //! TODO: get total supply, owner address" )

                    //_data := _get_token_creation_txn_info( txn )
                    //_data := _get_token_creation_txn_info( tx_hash.(string) )
                    _data := _get_token_creation_txn_info( tx_hash.(string), timestamp_unixtime, fmt.Sprintf("%s", tx_timestamp_date) )

                    fmt.Println( "=== _data ===" )
                    fmt.Println( _data )
                    if len(_data.From_address) <= 0 {
                        fmt.Println( "data == NULL" )
                        return
                    }
                    if len(_data.To_address) <= 0 {
                        fmt.Println( "data == NULL" )
                        return
                    }

                    db_insert_txns(_E_TYPE__ERC_TOKEN_CREATION, &_data )
                } else {
                    // ETH
                    _data := types.Fetch_transactions_st {
                        //Symbol: tx_token_symbol,
                        From_address: tx_from.(string),
                        To_address: tx_to.(string),

                        //Is_send: "",

                        // for SQL: ETHER
                        Amount_wei: tx_value_wei,
                        Amount_eth: tx_value,

                        Token_type: tx_token_type,

                        // for SQL: ERC-20
                        Token_symbol: "",
                        Token_decimals: "",
                        Token_total_supply: "",
                        Token_contract_address: "",
                        Token_amount_wei: "",
                        Token_amount_eth: "",

                        // for SQL: ERC-1155
                        Token_amount: "",
                        Token_uri_ascii: "",
                        Token_uri_hexadecimal: "",
                        Token_data_length: "",
                        Token_data: "",

                        Timestamp: timestamp_unixtime,
                        Datetime: fmt.Sprintf("%s", tx_timestamp_date),
                        Block_number: tx_block_number,
                        Txid: tx_hash.(string),
                    }
                    db_insert_txns(_E_TYPE__ETH, &_data )

                }

                fmt.Println( "------------------------------" )
            } // for (), transactions
        } // for (),  blocks
    }
}


// SEE: func get_balances_all_to_queue_infinite(...) {}
//
// goroutine: Updates_balances_all_to_queue__main_func() -> get_balances_all_to_queue()
//  - 1. get tnxs from 'txid', insert txns info 'balances'
//  - 2. get txns from 'balances', insert (insert/update query string) into 'balances_query_queue'
// goroutine: Updates_balances_all_to_queue_worker__main_func() -> updates_balances_all_to_queue_worker()
//  - 1. execute query from 'balances_query_queue'
func get_balances_all_to_queue() {
    TAG := "get_balances_all_to_queue(): "
    fmt.Println( "get_balances_all_to_queue()" )
    //return


    var TABLE_NAME_BALANCES = "balances"
    var TABLE_NAME_TXID = "txid"

    var query_str = ""
    //var query_str_update = ""


    //var INTERVAL = int(1) // seconds
    count_repeat := uint64(0)

    block_height_from_db := "" // from table 'balances': start block number
    block_height_latest_from_db := "" // from table 'txid': end block number (latest)
    block_num_start_uint64 := uint64(0)
    block_num_end_uint64 := uint64(0)




/*
1. get transactions per block, adds to new json_data
balances (
idx BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
last_blocks TEXT NOT NULL COMMENT "Block number",
blocks_hash VARCHAR(100) NOT NULL COMMENT "Block Hash",
json_data JSON NOT NULL COMMENT 'JSON Array',
# [{ "block": "", "txid": "", "from_address": "", "to_address": "", "token_type": "", "token_symbol": "", "amount": "" }, ...]
)

(SELECT A.*, B.* FROM
(SELECT blocks, blocks_hash FROM blocks WHERE blocks="1324") AS A,
(SELECT JSON_ARRAYAGG(JSON_OBJECT('blocks', blocks, 'txid', txid, 'from_address', from_address, 'to_address', to_address,
'token_type', token_type, 'token_symbol', token_symbol, 'ether_amount_eth', ether_amount_eth, 'token_amount_eth', token_amount_eth,
'erc1155_token_amount', erc1155_token_amount)) AS json_data FROM txid WHERE blocks="1324") AS B
);

INSERT INTO balances (idx, last_blocks, blocks_hash, json_data)
(SELECT 0, A.*, B.* FROM
(SELECT blocks, blocks_hash FROM blocks WHERE blocks="1324") AS A,
(SELECT JSON_ARRAYAGG(JSON_OBJECT('blocks', blocks, 'txid', txid, 'from_address', from_address, 'to_address', to_address,
'token_type', token_type, 'token_symbol', token_symbol, 'ether_amount_eth', ether_amount_eth, 'token_amount_eth', token_amount_eth,
'erc1155_token_amount', erc1155_token_amount)) AS json_data FROM txid WHERE blocks="1324") AS B
);
*/

    query_str = fmt.Sprintf(
        "SELECT last_blocks FROM %s ORDER BY last_blocks+0 DESC LIMIT 1",

        TABLE_NAME_BALANCES,
    )

    fmt.Println( "query = ", query_str )

    result, err := gDB.Query( query_str )

    if err != nil {
        panic( err.Error() )
        return
    }
    //defer result.Close()

    for result.Next() {
        //var Index sql.NullInt64
        //var Blocks string

        err := result.Scan(
            //&Index,
            &block_height_from_db,
        )

        if err != nil {
            //log.Fatal(err)
            //panic( err )
            fmt.Println( err )
        }

        break // LIMIT 1
    }
    result.Close()

    if len(block_height_from_db) <= 0 {
        block_height_from_db = "0"
    }

    //fmt.Println( TAG, "block height from DB (str) = ", block_height_from_db )
    block_num_start_uint64, _ = strconv.ParseUint( block_height_from_db, 10, 64 )
    //fmt.Println( TAG, "block height from DB = ", block_num_start_uint64 )


    for {
        //time.Sleep( time.Second * time.Duration(INTERVAL) )
        time.Sleep( time.Millisecond * 1000 * time.Duration(INTERVAL) )

        {
            query_str = fmt.Sprintf(
                "SELECT blocks FROM %s ORDER BY blocks+0 DESC LIMIT 1",

                TABLE_NAME_TXID,
            )

            result, err := gDB.Query( query_str )

            if err != nil {
                //panic( err.Error() )
                fmt.Println( err.Error() )
                continue
            }
            //defer result.Close()

            for result.Next() {
                //var Index sql.NullInt64
                //var Blocks string

                err := result.Scan(
                    //&Index,
                    &block_height_latest_from_db,
                )

                if err != nil {
                    //log.Fatal(err)
                    //panic( err )
                    fmt.Println( err )
                }

                break // LIMIT 1
            }

            //fmt.Println( TAG, "block height latest from DB (str) = ", block_height_latest_from_db )
            block_num_end_uint64, _ = strconv.ParseUint( block_height_latest_from_db, 10, 64 )
            //fmt.Println( TAG, "block height latest from DB = ", block_num_end_uint64 )

            result.Close()
        }


        //if ( block_height_old_uint64 <= 0 ) {
        //    block_height_old_uint64 = block_num_end_uint64
        //}
        //fmt.Println( TAG, "[block height] = ", block_height_old_uint64, "/", block_num_end_uint64 )
        fmt.Println( TAG, "[block height] = ", block_num_start_uint64, "/", block_num_end_uint64 )

        //if ( block_num_end_uint64 == block_height_old_uint64 ) {
        if ( block_num_end_uint64 == block_num_start_uint64 ) {
            count_repeat++
        } else {
            count_repeat = 0
        }

        if ( count_repeat > 1 ) {
            fmt.Println( TAG, "skip... count = ", count_repeat )
            fmt.Println()
            continue
        }
        fmt.Println()

        //updates_balances_all_to_queue_infinite( fmt.Sprintf("%d", block_height_old_uint64) )
        //updates_balances_all_to_queue_infinite( block_height_old_uint64, block_height_old_uint64 + 1 )
        updates_balances_all_to_queue_infinite( block_num_start_uint64, block_num_start_uint64 + 1 )

        //if ( block_num_end_uint64 > block_height_old_uint64 ) {
        //    block_height_old_uint64++
        //}
        if ( block_num_end_uint64 > block_num_start_uint64 ) {
            block_num_start_uint64++
        }


        fmt.Println()
    } // for ()
}

func updates_balances_all_to_queue_infinite(block_num_start_uint64 uint64, block_num_end_uint64 uint64) {
    TAG := "updates_balances_all_to_queue_infinite(): "
    fmt.Println( "updates_balances_all_to_queue_infinite()" )
    //return


    var blocks = "" // block number
    var TABLE_NAME = ""
    var TABLE_NAME_BALANCES = "balances"
    var TABLE_NAME_TXID = "txid"
    var TABLE_NAME_BALANCES_QUERY_QUEUE = "balances_query_queue"
    var query_str = ""
    var query_str_update__from_addr = ""
    var query_str_update__to_addr = ""
    var query_str_insert__from_addr = ""
    var query_str_insert__to_addr = ""
    var query_str_queue__from_addr = ""
    var query_str_queue__to_addr = ""


    fmt.Println( TAG, "block start = ", block_num_start_uint64, "block end = ", block_num_end_uint64 )
    //fmt.Println()
    for i := block_num_start_uint64; i < uint64(block_num_end_uint64); i++ {
        blocks = fmt.Sprintf( "%d", i )

        // makes json_data, insert into table 'balances'
        query_str = fmt.Sprintf(
            "INSERT INTO %s (idx, last_blocks, blocks_hash, json_data)" +
            " (SELECT 0, A.*, B.* FROM" +
            " (SELECT blocks, blocks_hash FROM blocks WHERE blocks='%s') AS A," +

            // array
            //" (SELECT JSON_ARRAYAGG(JSON_OBJECT('blocks', blocks, 'txid', txid, 'from_address', from_address, 'to_address', to_address," +
            //" 'token_type', token_type, 'token_symbol', token_symbol, 'token_contract_address', token_contract_address," +
            //" 'ether_amount_eth', ether_amount_eth, 'token_amount_eth', token_amount_eth," +
            //" 'erc1155_token_amount', erc1155_token_amount)) AS json_data FROM %s WHERE blocks='%s') AS B" +

            " (SELECT JSON_ARRAYAGG(JSON_OBJECT('txid', txid, 'blocks', blocks, 'from_address', from_address, 'to_address', to_address," +
            " 'token_type', token_type, 'token_symbol', token_symbol, 'token_contract_address', token_contract_address," +
            " 'ether_amount_eth', ether_amount_eth, 'token_amount_eth', token_amount_eth," +
            " 'erc1155_token_amount', erc1155_token_amount)) AS json_data FROM %s WHERE blocks='%s') AS B" +


            // object
            //" (SELECT JSON_OBJECT('blocks', blocks, 'txid', txid, 'from_address', from_address, 'to_address', to_address," +
            //" 'token_type', token_type, 'token_symbol', token_symbol, 'token_contract_address', token_contract_address," +
            //" 'ether_amount_eth', ether_amount_eth, 'token_amount_eth', token_amount_eth," +
            //" 'erc1155_token_amount', erc1155_token_amount) AS json_data FROM %s WHERE blocks='%s') AS B" +

            //" (SELECT JSON_OBJECT('txid', txid, 'blocks', blocks, 'from_address', from_address, 'to_address', to_address," +
            //" 'token_type', token_type, 'token_symbol', token_symbol, 'token_contract_address', token_contract_address," +
            //" 'ether_amount_eth', ether_amount_eth, 'token_amount_eth', token_amount_eth," +
            //" 'erc1155_token_amount', erc1155_token_amount) AS json_data FROM %s WHERE blocks='%s') AS B" +


            " )",


            TABLE_NAME_BALANCES, blocks, TABLE_NAME_TXID, blocks,
        )

        //fmt.Println( TAG, "query = ", query_str )

        result, err := gDB.Query( query_str )

        if err != nil {
            //! FIXME: Handling Error codes
            if strings.Contains( err.Error(), "Duplicate entry" ) {
                fmt.Println( TAG, "Error: Duplicate entry" )

                // update(replace) it...
                //fmt.Println( TAG, "update(replace) query = ", query_str_update )
                //result, err = gDB.Query( query_str_update )
                //if err != nil {
                //    //panic( err.Error() )
                //    fmt.Println( err.Error() )
                //    continue
                //} else {
                //    result.Close()
                //}
            } else {
                //panic( err.Error() )
                fmt.Println( err.Error() )
            }

            //return
            continue
        } else {
            //defer result.Close()
            result.Close()
        }



        // insert 'query_string' into table 'balances_query_queue'

/*
2. calculates balances and updates balances_address_suffix_0 ~ balances_address_suffix_f
balances_address_suffix_0 (
idx BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
owner_address VARCHAR(100) NOT NULL COMMENT "Owner address",
last_txid VARCHAR(100) NOT NULL COMMENT "Block Hash",
last_blocks TEXT NOT NULL COMMENT "Block number",
token_type VARCHAR(32) NOT NULL COMMENT "Token type: ETH, ERC20, ERC1155",
token_symbol VARCHAR(32) NOT NULL COMMENT "Token symbol name",
token_contract_address VARCHAR(100) NOT NULL COMMENT "ERC Contract address",
amount TEXT NOT NULL COMMENT "amount",
issend VARCHAR(1) NOT NULL DEFAULT "n" COMMENT "is send: y or n",
timestamp VARCHAR(20) NOT NULL COMMENT "Timestamp: 1643798246",
)


SELECT * FROM balances WHERE last_blocks=""


UPDATE <balances_address_suffix_?>, (SELECT amount FROM <balances_address_suffix_?> AS A)
SET last_txid, json_data.txid, last_blocks=json_data.blocks, token_type=json_data.token_type,
token_symbol=json_data.token_symbol, amount=<A.amount + json_data.amount>, issend=issend, timestamp=timestamp
WHERE owner_address = owner_address and token_type = json_token_type
*/


        // Get JSON data from the table 'balances'
        // table_name, blocks from above
        query_str = fmt.Sprintf(
            "SELECT * FROM %s WHERE last_blocks='%s'",

            TABLE_NAME_BALANCES, blocks,
        )

        result, err = gDB.Query( query_str )

        if err != nil {
            panic( err.Error() )
        }
        //defer result.Close()

        type json_data_st struct {
            Txid string `json:"txid"`
            Blocks string `json:"blocks"`
            From_address string `json:"from_address"`
            To_address string `json:"to_address"`
            Token_type string `json:"token_type"`
            Token_symbol string `json:"token_symbol"`
            Token_contract_address string `json:"token_contract_address"`
            Amount_ether_amount_eth string `json:"ether_amount_eth"`
            Amount_token_amount_eth string `json:"token_amount_eth"`
            Amount_erc1155_token_amount string `json:"erc1155_token_amount"`
        }
        //var _json_data []json_data_st // result
        var _json_data json_data_st // result
        var json_data []json_data_st // result

        for result.Next() {
            var Index sql.NullInt64
            var Last_blocks string
            var Blocks_hash string
            var Json_data_str string

            err := result.Scan(
                &Index,
                &Last_blocks,
                &Blocks_hash,
                &Json_data_str,
            )

            if err != nil {
                //log.Fatal(err)
                panic( err )
            }

            fmt.Println( TAG, "json_data (string) = " )
            fmt.Println( TAG, Json_data_str )

            /*
            //var _data = make( map[string]interface{} )
            err = json.Unmarshal( []byte(Json_data_str), &_json_data )
            if err != nil {
                //log.Fatal(err)
                panic( err.Error() )
            }
            ////json_map["txid"] = _data
            //json_map_arr["txid"] = append( json_map_arr["txid"], json_data )
            //json_arr = append( json_arr, json_data )
            json_data = append( json_data, _json_data )
            */

            var _json_data_objmap []map[string]json.RawMessage
            err = json.Unmarshal( []byte(Json_data_str), &_json_data_objmap )
            if err != nil {
                //log.Fatal(err)
                panic( err.Error() )
            }
            json_data_objmap := _json_data_objmap[0]

            json_txid := ""
            json_blocks := ""
            json_from_address := ""
            json_to_address := ""
            json_token_type := ""
            json_token_symbol := ""
            json_token_contract_address := ""
            json_amount_ether_amount_eth := ""
            json_amount_token_amount_eth := ""
            json_amount_erc1155_token_amount := ""

            err = json.Unmarshal( json_data_objmap["txid"], &json_txid )
            if err != nil {
                //log.Fatal(err)
                panic( err.Error() )
            }
            err = json.Unmarshal( json_data_objmap["blocks"], &json_blocks )
            if err != nil {
                //log.Fatal(err)
                panic( err.Error() )
            }

            err = json.Unmarshal( json_data_objmap["from_address"], &json_from_address )
            if err != nil {
                //log.Fatal(err)
                panic( err.Error() )
            }

            err = json.Unmarshal( json_data_objmap["to_address"], &json_to_address )
            if err != nil {
                //log.Fatal(err)
                panic( err.Error() )
            }

            err = json.Unmarshal( json_data_objmap["token_type"], &json_token_type )
            if err != nil {
                //log.Fatal(err)
                panic( err.Error() )
            }

            err = json.Unmarshal( json_data_objmap["token_symbol"], &json_token_symbol )
            if err != nil {
                //log.Fatal(err)
                panic( err.Error() )
            }

            err = json.Unmarshal( json_data_objmap["token_contract_address"], &json_token_contract_address )
            if err != nil {
                //log.Fatal(err)
                panic( err.Error() )
            }

            err = json.Unmarshal( json_data_objmap["ether_amount_eth"], &json_amount_ether_amount_eth )
            if err != nil {
                //log.Fatal(err)
                panic( err.Error() )
            }

            err = json.Unmarshal( json_data_objmap["token_amount_eth"], &json_amount_token_amount_eth )
            if err != nil {
                //log.Fatal(err)
                panic( err.Error() )
            }

            err = json.Unmarshal( json_data_objmap["erc1155_token_amount"], &json_amount_erc1155_token_amount )
            if err != nil {
                //log.Fatal(err)
                panic( err.Error() )
            }


            _json_data = json_data_st {
                Txid: json_txid,
                Blocks: json_blocks,
                From_address: json_from_address,
                To_address: json_to_address,
                Token_type: json_token_type,
                Token_symbol: json_token_symbol,
                Token_contract_address: json_token_contract_address,
                Amount_ether_amount_eth: json_amount_ether_amount_eth,
                Amount_token_amount_eth: json_amount_token_amount_eth,
                Amount_erc1155_token_amount: json_amount_erc1155_token_amount,
            }
            json_data = append( json_data, _json_data )


            //fmt.Println( TAG, "json_data (object) = " )
            //fmt.Println( TAG, json_data )
        } // for (), query results
        result.Close()

        fmt.Println( TAG, "json_data_len = ", len(json_data) )
        fmt.Println( TAG, "json_data (object) = " )
        fmt.Println( TAG, json_data )

        for i_json_arr := 0; i_json_arr < len(json_data); i_json_arr++ {
            var owner_address = ""
            var amount = ""
            var issend = "" // 'y' or 'n'
            //var owner_address = ""
            var timestamp = ""

            fmt.Println( TAG, "from_address = ", json_data[i_json_arr].From_address, len(json_data[i_json_arr].From_address) )
            fmt.Println( TAG, "to_address = ", json_data[i_json_arr].To_address, len(json_data[i_json_arr].To_address) )
            fmt.Println( TAG, "contract_address = ", json_data[i_json_arr].Token_contract_address, len(json_data[i_json_arr].Token_contract_address) )
            if len(json_data[i_json_arr].From_address) > 0 {
                //owner_address = json_data[i_json_arr].From_address
            } else if len(json_data[i_json_arr].To_address) > 0 {
                //owner_address = json_data[i_json_arr].To_address
                //issend = "n"
            } else {
                //panic( "from_address, to_address == NULL" )
                fmt.Println( TAG, "from_address, to_address == NULL" )
                continue
            }

            fmt.Println( TAG, "token_type = ", json_data[i_json_arr].Token_type )
            switch json_data[i_json_arr].Token_type {
            case TOKEN_TYPE_ETHER:
                amount = json_data[i_json_arr].Amount_ether_amount_eth
            case TOKEN_TYPE_ERC20:
                amount = json_data[i_json_arr].Amount_token_amount_eth
            case TOKEN_TYPE_ERC1155:
                amount = json_data[i_json_arr].Amount_erc1155_token_amount
            default:
                //panic( "amount == NULL" )
                fmt.Println( TAG, "amount == NULL" )
                continue
            }
            fmt.Println( TAG, "amount = ", amount )



            // Insert query (updates balances) string into <table 'query_queue' or Kafka>

            // for from_addr
            owner_address = json_data[i_json_arr].From_address
            issend = "y"
            fmt.Println( "owner_address = ", owner_address )
            fmt.Println( "owner_address[(len-1)] = ", owner_address[len(owner_address)-1:] )
            TABLE_NAME = "balances_address_suffix_" + owner_address[len(owner_address)-1:]
            query_str_update__from_addr = fmt.Sprintf(
                "UPDATE %s AS A, (SELECT amount FROM %s) AS B SET" +
                " A.txid='%s', A.blocks='%s', A.token_type='%s'," +
                " A.token_symbol='%s', A.token_contract_address='%s'," +
                " A.amount = (B.amount+0 + '%s'), A.issend='%s', A.timestamp='%s'" +
                //" WHERE A.owner_address='%s' and A.token_type='%s'" ,
                //
                // INSERT ONLY. UPDATES doesn't work now. SEE: query_hash
                // for UPDATE, FIXME: remove all for query_hash, remove 'balances_query_queue.query_hash' field.
                " WHERE A.owner_address='%s' and A.token_type='%s' and A.query_hash='%s%s%s%s%s%s%s%s%s'" ,

                TABLE_NAME, TABLE_NAME,
                json_data[i_json_arr].Txid, json_data[i_json_arr].Blocks, json_data[i_json_arr].Token_type,
                json_data[i_json_arr].Token_symbol, json_data[i_json_arr].Token_contract_address, amount, issend, timestamp,
                owner_address, json_data[i_json_arr].Token_type,

                // INSERT ONLY. UPDATE doesn't work now. SEE: query_hash
                // for UPDATE, FIXME: remove all for query_hash, remove 'balances_query_queue.query_hash' field.
                owner_address,
                json_data[i_json_arr].Txid, json_data[i_json_arr].Blocks, json_data[i_json_arr].Token_type,
                json_data[i_json_arr].Token_symbol, json_data[i_json_arr].Token_contract_address, amount, issend, timestamp,
            )
            query_str_insert__from_addr = fmt.Sprintf(
                "INSERT INTO %s VALUES(0, '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', " +
                //"SHA1('')" +
                //
                // INSERT ONLY. UPDATES doesn't work now. SEE: query_hash
                // for UPDATE, FIXME: remove all for query_hash, remove 'balances_query_queue.query_hash' field.
                "SHA1('%s%s%s%s%s%s%s%s%s')" +
                ")" ,

                TABLE_NAME,
                owner_address,
                json_data[i_json_arr].Txid, json_data[i_json_arr].Blocks, json_data[i_json_arr].Token_type,
                json_data[i_json_arr].Token_symbol, json_data[i_json_arr].Token_contract_address, amount, issend, timestamp,

                // INSERT ONLY. UPDATES doesn't work now. SEE: query_hash
                // for UPDATE, FIXME: remove all for query_hash, remove 'balances_query_queue.query_hash' field.
                owner_address,
                json_data[i_json_arr].Txid, json_data[i_json_arr].Blocks, json_data[i_json_arr].Token_type,
                json_data[i_json_arr].Token_symbol, json_data[i_json_arr].Token_contract_address, amount, issend, timestamp,
            )

            // for to_addr
            owner_address = json_data[i_json_arr].To_address
            issend = "n"
            TABLE_NAME = "balances_address_suffix_" + owner_address[len(owner_address)-1:]
            query_str_update__to_addr = fmt.Sprintf(
                "UPDATE %s AS A, (SELECT amount FROM %s) AS B SET" +
                " A.txid='%s', A.blocks='%s', A.token_type='%s'," +
                " A.token_symbol='%s', A.token_contract_address='%s'," +
                " A.amount = (B.amount+0 - '%s'), A.issend='%s', A.timestamp='%s'" +
                //" WHERE A.owner_address='%s' and A.token_type='%s'" ,
                //
                // INSERT ONLY. UPDATE doesn't work now. SEE: query_hash
                // for UPDATE, FIXME: remove all for query_hash, remove 'balances_query_queue.query_hash' field.
                " WHERE A.owner_address='%s' and A.token_type='%s' and A.query_hash='%s%s%s%s%s%s%s%s%s'" ,

                TABLE_NAME, TABLE_NAME,
                json_data[i_json_arr].Txid, json_data[i_json_arr].Blocks, json_data[i_json_arr].Token_type,
                json_data[i_json_arr].Token_symbol, json_data[i_json_arr].Token_contract_address, amount, issend, timestamp,
                owner_address, json_data[i_json_arr].Token_type,

                // INSERT ONLY. UPDATES doesn't work now. SEE: query_hash
                // for updates, FIXME: remove all for query_hash, remove 'balances_query_queue.query_hash' field.
                owner_address,
                json_data[i_json_arr].Txid, json_data[i_json_arr].Blocks, json_data[i_json_arr].Token_type,
                json_data[i_json_arr].Token_symbol, json_data[i_json_arr].Token_contract_address, amount, issend, timestamp,
            )
            query_str_insert__to_addr = fmt.Sprintf(
                "INSERT INTO %s VALUES(0, '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', " +
                //"SHA1('')" +
                //
                // INSERT ONLY. UPDATES doesn't work now. SEE: query_hash
                // for UPDATE, FIXME: remove all for query_hash, remove 'balances_query_queue.query_hash' field.
                "SHA1('%s%s%s%s%s%s%s%s%s')" +
                ")" ,

                TABLE_NAME,
                owner_address,
                json_data[i_json_arr].Txid, json_data[i_json_arr].Blocks, json_data[i_json_arr].Token_type,
                json_data[i_json_arr].Token_symbol, json_data[i_json_arr].Token_contract_address, amount, issend, timestamp,

                // INSERT ONLY. UPDATES doesn't work now. SEE: query_hash
                // for UPDATE, FIXME: remove all for query_hash, remove 'balances_query_queue.query_hash' field.
                owner_address,
                json_data[i_json_arr].Txid, json_data[i_json_arr].Blocks, json_data[i_json_arr].Token_type,
                json_data[i_json_arr].Token_symbol, json_data[i_json_arr].Token_contract_address, amount, issend, timestamp,
            )



            // Query string queue
            query_str_queue__from_addr = fmt.Sprintf(
                "INSERT INTO %s VALUES(0, '%s', '%s', \"%s\", \"%s\", SHA1(\"%s%s%s\")" +
                ")" ,

                TABLE_NAME_BALANCES_QUERY_QUEUE,
                json_data[i_json_arr].Blocks, json_data[i_json_arr].Txid,
                query_str_update__from_addr, query_str_insert__from_addr,
                json_data[i_json_arr].Txid, query_str_update__from_addr, query_str_insert__from_addr,
            )
            query_str_queue__to_addr = fmt.Sprintf(
                "INSERT INTO %s VALUES(0, '%s', '%s', \"%s\", \"%s\", SHA1(\"%s%s%s\")" +
                ")" ,

                TABLE_NAME_BALANCES_QUERY_QUEUE,
                json_data[i_json_arr].Blocks, json_data[i_json_arr].Txid,
                query_str_update__to_addr, query_str_insert__to_addr,
                json_data[i_json_arr].Txid, query_str_update__to_addr, query_str_insert__to_addr,
            )


            fmt.Println( TAG, "query = ", query_str_queue__from_addr )

            result, err = gDB.Query( query_str_queue__from_addr )

            if err != nil {
                //! FIXME: Handling Error codes
                if strings.Contains( err.Error(), "Duplicate entry" ) {
                    fmt.Println( TAG, "Error: Duplicate entry" )

                    // update(replace) it...
                    //fmt.Println( TAG, "update(replace) query = ", query_str_queue_update__from_addr )
                    //result, err = gDB.Query( query_str_queue_update__from_addr )
                    //if err != nil {
                    //    //panic( err.Error() )
                    //    fmt.Println( TAG, err.Error() )
                    //    fmt.Println( TAG, "skip..." )
                    //    continue
                    //} else {
                    //    result.Close()
                    //}

                    fmt.Println( TAG, "skip UPDATE... INSERT balance per transaction..." )
                } else {
                    //panic( err.Error() )
                    fmt.Println( err.Error() )
                    //return
                }

                //return
            } else {
                //defer result.Close()
                result.Close()
            }


            fmt.Println( TAG, "query = ", query_str_queue__to_addr )

            result, err = gDB.Query( query_str_queue__to_addr )

            if err != nil {
                //! FIXME: Handling Error codes
                if strings.Contains( err.Error(), "Duplicate entry" ) {
                    fmt.Println( TAG, "Error: Duplicate entry" )

                    // update(replace) it...
                    //fmt.Println( TAG, "update(replace) query = ", query_str_queue_update__to_addr )
                    //result, err = gDB.Query( query_str_queue_update__to_addr )
                    //if err != nil {
                    //    //panic( err.Error() )
                    //    fmt.Println( TAG, err.Error() )
                    //    fmt.Println( TAG, "skip..." )
                    //    continue
                    //} else {
                    //    result.Close()
                    //}

                    fmt.Println( TAG, "skip UPDATE... INSERT balance per transaction..." )
                } else {
                    //panic( err.Error() )
                    fmt.Println( err.Error() )
                    //return
                }

                //return
            } else {
                //defer result.Close()
                result.Close()
            }
        } // for (), json_data array
    } // for ()
}

// goroutine: Updates_balances_all_to_queue__main_func() -> get_balances_all_to_queue()
//  - 1. get tnxs from 'txid', insert txns info 'balances'
//  - 2. get txns from 'balances', insert (insert/update query string) into 'balances_query_queue'
// goroutine: Updates_balances_all_to_queue_worker__main_func() -> updates_balances_all_to_queue_worker()
//  - 1. execute query from 'balances_query_queue'
func updates_balances_all_to_queue_worker() {
    TAG := "updates_balances_all_to_queue_worker(): "
    fmt.Println( "updates_balances_all_to_queue_worker(): === TODO ===" )
    //return


    var TABLE_NAME_BALANCES_QUERY_QUEUE = "balances_query_queue"
    var query_str = ""
    //var query_str_update = ""


    //var INTERVAL = int(1) // seconds
    //count_repeat := uint64(0)

    for {
        //time.Sleep( time.Second * time.Duration(INTERVAL) )
        time.Sleep( time.Millisecond * 1000 * time.Duration(INTERVAL) )

        //! FIXME:
        // transaction start;
        // 1. get queue: query string
        // 2. executes the query
        // 3. removes the query row
        // transaction end;

        var query_str_executes_update = ""
        var query_str_executes_insert = ""
        var query_hash = ""
        total_rows_uint64 := uint64(0)

        {
            //! process sequentially
            query_str = fmt.Sprintf(
                "SELECT *, COUNT(*) AS TOTAL_ROW FROM %s ORDER BY blocks+0 ASC LIMIT 1",

                TABLE_NAME_BALANCES_QUERY_QUEUE,
            )

            result, err := gDB.Query( query_str )

            if err != nil {
                //panic( err.Error() )
                fmt.Println( err.Error() )

                //count_repeat++
                //fmt.Println( TAG, "skip... count = ", count_repeat )
                //fmt.Println()

                continue
            }
            //defer result.Close()

            for result.Next() {
                var Index sql.NullInt64
                var Blocks string //sql.NullString
                var Txid string //sql.NullString

                //! ignore NULL string, int, ...
                // Error message: sql: Scan error on column index ..., name "...": converting NULL to string is unsupported

                err := result.Scan(
                    &Index,
                    &Blocks,
                    &Txid,
                    &query_str_executes_update,
                    &query_str_executes_insert,
                    &query_hash,
                    &total_rows_uint64,
                )

                if err != nil {
                    //log.Fatal(err)
                    //panic( err )
                    fmt.Println( err )

                    //count_repeat++
                    fmt.Println( TAG, "// ignore NULL string, int, ..." )
                    fmt.Println( TAG, "// Error message: sql: Scan error on column index ..., name \"...\": converting NULL to string is unsupported" )
                    //fmt.Println( TAG, "skip... count = ", count_repeat )
                    fmt.Println()

                    continue
                }

                break // LIMIT 1
            }
            result.Close()

            fmt.Println( TAG, "total rows = ", total_rows_uint64 )
            if total_rows_uint64 <= 0 {
                //count_repeat++
                //fmt.Println( TAG, "skip... count = ", count_repeat )
                //fmt.Println()

                continue
            }

            fmt.Println( TAG, "query = ", query_str_executes_insert )

            fmt.Println( TAG, "========== FIXME: TODO: adds transaction { insert/update -> delete query_queue } ==========" )

            result, err = gDB.Query( query_str_executes_insert )

            if err != nil {
                //panic( err.Error() )
                fmt.Println( err.Error() )

                if strings.Contains( err.Error(), "Duplicate entry" ) {
                    fmt.Println( TAG, "Error: Duplicate entry" )

                    // update(replace) it...
                    fmt.Println( TAG, "update(replace) query = ", query_str_executes_update )
                    result, err = gDB.Query( query_str_executes_update )
                    if err != nil {
                        //panic( err.Error() )
                        fmt.Println( err.Error() )
                    } else {
                        result.Close()

                        // remove query from 'balances_query_queue'
                        fmt.Println( TAG, "remove query = ", query_str_executes_insert )

                        query_str = fmt.Sprintf(
                            "DELETE FROM %s WHERE query_hash='%s'",

                            TABLE_NAME_BALANCES_QUERY_QUEUE, query_hash,
                        )

                        result, err := gDB.Query( query_str )

                        if err != nil {
                            //panic( err.Error() )
                            fmt.Println( err.Error() )
                        } else {
                            result.Close()
                        }
                    }
                } else {
                    //panic( err.Error() )
                    fmt.Println( err.Error() )
                    //return
                }

                //count_repeat++
                //fmt.Println( TAG, "skip... count = ", count_repeat )
                //fmt.Println()

                continue
            }
            //defer result.Close()
            result.Close()
        }

        //count_repeat = 0

        fmt.Println()
    } // for ()
}



func features() {
    /*
    {
        // Ether: balance
        address := "0xe6e55eed00218faef27eed24def9208f3878b333"
        eth_get_balance( address )
    }
    */


    fmt.Println( "----------" )


    /*
    {
        // Ether: send transaction
        from := "0xe6e55eed00218faef27eed24def9208f3878b333"
        to := "0x8f5b2b7608e3e3a3dc0426c3396420fbf1849454"
        amount := "11.1357" // 10 * wei(1e18)
        gas := "70000"
        gasprice := "100"
        eth_send_transaction( from, to, amount, gas, gasprice )
    }
    */


    /*
    {
        // ERC-20: balance
        //contract_address := "0x1e57f9561600b269a37437f02ce9da31e5b830ce"
        //contract_address := "0xD0EFa91095e04B642df6846D96d6d1aD0afd05eE"
        contract_address := "0xB5AccFe1b7a59317A9F5A100dC1105Ed66b2058c"
        address := "0xe6e55eed00218faef27eed24def9208f3878b333"
        erc20_get_balance( contract_address, address )
    }
    */


    fmt.Println( "----------" )


    /*
    {
        // ERC-20: send transaction
        //contract_address := "0x1e57f9561600b269a37437f02ce9da31e5b830ce"
        contract_address := "0xB5AccFe1b7a59317A9F5A100dC1105Ed66b2058c"
        from := "0xe6e55eed00218faef27eed24def9208f3878b333"
        to := "0x8f5b2b7608e3e3a3dc0426c3396420fbf1849454"
        amount := "11.1357" // 10 * wei(1e18)
        gas := "70000"
        gasprice := "100"
        erc20_send_transaction( contract_address,  from, to, amount, gas, gasprice )
    }
    */


    fmt.Println( "----------" )


    /*
    {
        // ERC-1155: balanceOf
        //contract_address := "0x8eA78d6BfdC5B3FFf3dde2a872235D3cFaFcc203" // erc1155: TestNFT
        contract_address := "0x1249CDA86774Bc170CAb843437DD37484F173ca8" // erc1155: TestNFT (USE THIS)
        address := "0xe6e55eed00218faef27eed24def9208f3878b333"
        token_id := "0"
        //token_id := "1"
        erc1155_get_balance( contract_address, address, token_id )

        address = "0x8f5b2b7608e3e3a3dc0426c3396420fbf1849454"
        erc1155_get_balance( contract_address, address, token_id )
    }
    */


    fmt.Println( "----------" )


    /*
    {
        // ERC-1155: uri
        //contract_address := "0x8eA78d6BfdC5B3FFf3dde2a872235D3cFaFcc203" // erc1155: TestNFT
        contract_address := "0x1249CDA86774Bc170CAb843437DD37484F173ca8" // erc1155: TestNFT (USE THIS)
        address := "0xe6e55eed00218faef27eed24def9208f3878b333"
        token_id := "0"
        //token_id := "1"
        //token_id := "10"
        erc1155_get_uri( contract_address, address, token_id )


        address = "0x8f5b2b7608e3e3a3dc0426c3396420fbf1849454"
        erc1155_get_uri( contract_address, address, token_id )
    }
    */


    fmt.Println( "----------" )


    /*
    //! TODO
    {
        // ERC-1155: _setURI
        //contract_address := "0x8eA78d6BfdC5B3FFf3dde2a872235D3cFaFcc203" // erc1155: TestNFT
        contract_address := "0x1249CDA86774Bc170CAb843437DD37484F173ca8" // erc1155: TestNFT (USE THIS)
        new_uri := "http://127.0.0.1/api/v2/token/{id}.json"
        erc1155_set_uri( contract_address, new_uri )
    }
    */


    fmt.Println( "----------" )


    /*
    {
        // ERC-1155: safeTransferFrom
        //contract_address := "0x8eA78d6BfdC5B3FFf3dde2a872235D3cFaFcc203" // erc1155: TestNFT
        contract_address := "0x1249CDA86774Bc170CAb843437DD37484F173ca8" // erc1155: TestNFT (USE THIS)
        from_address := "0xe6e55eed00218faef27eed24def9208f3878b333"
        to_address := "0x8f5b2b7608e3e3a3dc0426c3396420fbf1849454"
        token_id := "0"
        //token_id := "1"
        //token_id := "10"
        amount := "1"
        gas := "70000"
        gasprice := "100"
        erc1155_send_transaction( contract_address, from_address, to_address, token_id, amount, gas, gasprice )
    }
    */


    fmt.Println( "----------" )


    /*
    {
        // Get blocks: ETH, ERC-20
        //eth_get_block_by_number
        get_blocks()
    }
    */


    fmt.Println( "----------" )


    /*
    {
        // Get blocks: ERC-1155
        //eth_get_block_by_number
        get_blocks_erc1155()
    }
    */


    fmt.Println( "----------" )


    ///*
    {
        // Get blocks: ETH, ERC-20, ERC-1155
        //eth_get_block_by_number
        get_blocks_all()
    }
    //*/
}


func Init_db_func() {
    __init_db()
}

func Release_db_func() {
    __release_db()
}

func Txns_fetcher_main_func() {
    fmt.Println( "HOST: " + URL )

    //__test_db()

    features()
}

func Updates_balances_all_to_queue__main_func() {
    get_balances_all_to_queue()
}

func Updates_balances_all_to_queue_worker__main_func() {
    updates_balances_all_to_queue_worker()
}



// replace package name
// for "package main"
func main() {
    Txns_fetcher_main_func()
}

