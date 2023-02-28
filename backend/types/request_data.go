/* --------------------------------------------------------------
Project:    Ethereum auto-transfer (accounts to specific address(hotwallet))
Purpose:
Author:     Ho-Jung Kim (godmode2k@hotmail.com)
Date:       Since Dec 4, 2020
Filename:   request_data.go

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
-------------------------------------------------------------- */
package types



//! Header
// ---------------------------------------------------------------

import (
    "encoding/json"
)



//! Definition
// --------------------------------------------------------------------

type RequestData struct {
    Jsonrpc string `json:"jsonrpc"`
    Method string `json:"method"`

    Params []interface{} `json:"params"`
    // ["address", "latest"]
    // [{"to": "<contract address>", "data": ""}, "latest"]
    // [{"from": "", "to": "<contract address>", "gas": "", "gasPrice": "", "data": ""}, "latest"]

    Id int `json:"id"`
}

type RequestData_params_erc20 struct {
    // [{"to": "<contract address>", "data": ""}, "latest"]
    To string `json:"to"`
    Data string `json:"data"`
}

type RequestData_params_erc1155 struct {
    // [{"to": "<contract address>", "data": ""}, "latest"]
    To string `json:"to"`
    Data string `json:"data"`
}

type RequestData_params_transaction struct {
    // [{"from": "", "to": "", "value": "<wei>", "gas": "", "gasPrice": "", "data": "", "nonce": ""}, "latest"]
    From string `json:"from"`
    To string `json:"to"`
    Value string `json:"value"`
    Gas string `json:"gas"`
    Gasprice string `json:"gasPrice"`
    //! DO NOT USE
    // except: cancel pending transaction, ...
    //Data string `json:"data"`
    //Nonce string `json:"nonce"`
}

type RequestData_params_erc20_transaction struct {
    // [{"from": "", "to": "<contract address>", "value": "<wei>", "gas": "", "gasPrice": "", "data": ""}, "latest"]
    From string `json:"from"`
    To string `json:"to"`
    Value string `json:"value"`
    Gas string `json:"gas"`
    Gasprice string `json:"gasPrice"`
    Data string `json:"data"`
}

type Result struct {
    Jsonrpc string `json:"jsonrpc"`
    Id int `json:"id"`
    Result string `json:"result"`
}

type Result_block struct {
    Jsonrpc string `json:"jsonrpc"`
    Id int `json:"id"`
    Result interface{} `json:"result"`
}


// Fetch Transactions: ETH, ERC-20, ERC-1155 (NFT, FT)
type Fetch_transactions_st struct {
    // [{"from": "", "to": "" ...}]
    //Symbol string `json:"symbol"`
    From_address string `json:"from_address"`
    To_address string `json:"to_address"`

    //Is_send string `json:"is_send"`

    // ETHER
    Amount_wei string `json:"amount_wei"`
    Amount_eth string `json:"amount_eth"`

    // Token: ERC-20
    Token_type string `json:"token_type"`
    Token_symbol string `json:"token_symbol"`
    Token_decimals string `json:"token_decimals"`
    Token_total_supply string `json:"token_total_supply"`
    Token_contract_address string `json:"token_contract_address"`
    Token_amount_wei string `json:"token_amount_wei"`
    Token_amount_eth string `json:"token_amount_eth"`

    // Token: ERC-1155 (NFT, FT)
    Token_amount string `json:"token_amount"`
    Token_uri_ascii string `json:"token_uri_ascii"`
    Token_uri_hexadecimal string `json:"token_uri_hexadecimal"`
    Token_data_length string `json:"token_data_length"`
    Token_data string `json:"token_data"`

    Timestamp string `json:"timestamp"`
    Datetime string `json:"datetime"`
    Block_number string `json:"block_number"`
    Txid string `json:"txid"`
}

type Fetch_blocks_info_st struct {
    Block_number string `json:"block_number"`
    Block_hash string `json:"block_hash"`
    Info string `json:"block_info"`
    Transactions string `json:"transactions"`
}

type Fetch_balances_st struct {
    Last_blocks string `json:"last_block_number"`
    Blocks_hash string `json:"block_hash"`
    Json_data string `json:"json_data"`
}

type Fetch_balances_by_address_st struct {
    Owner_address string `json:"owner_address"`
    Txid string `json:"txid"`
    Block_number string `json:"block_number"`
    Token_type string `json:"token_type"`
    Token_symbol string `json:"token_symbol"`
    Token_contract_address string `json:"token_contract_address"`
    Amount string `json:"amount"`
}




// ---------------------------------------------------------------


func Fetch_transactions_json(data *Fetch_transactions_st) string {
    message, err := json.Marshal( data )
    if err != nil {
        panic( err.Error() )
    }

    return string(message)
}

func Fetch_transactions_array_json(data *[]Fetch_transactions_st) string {
    message, err := json.Marshal( data )
    if err != nil {
        panic( err.Error() )
    }

    return string(message)
}

func Fetch_blocks_json(data *Fetch_blocks_info_st) string {
    message, err := json.Marshal( data )
    if err != nil {
        panic( err.Error() )
    }

    return string(message)
}

func Fetch_blocks_array_json(data *[]Fetch_blocks_info_st) string {
    message, err := json.Marshal( data )
    if err != nil {
        panic( err.Error() )
    }

    return string(message)
}

func Fetch_balances_json(data *Fetch_balances_st) string {
    message, err := json.Marshal( data )
    if err != nil {
        panic( err.Error() )
    }

    return string(message)
}

func Fetch_balances_array_json(data *[]Fetch_balances_st) string {
    message, err := json.Marshal( data )
    if err != nil {
        panic( err.Error() )
    }

    return string(message)
}


// ---------------------------------------------------------------


// HTTP JSON-RPC Server: for frontend

type RPC_DummyArgs_st struct {
    Dummy int
}

type RPC_Args_st struct {
    Req_page int
    Req_block string
    Req_block_hash string
    Req_txid string
    Req_address string
    Req_search string
}


