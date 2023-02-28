import { fetch } from "./fetch";
import { fetch2 } from "./fetch";


export function getLatestBlockNumber() {
  return fetch2({
    url: "/rpc",
    method: "post",
    data: {
        jsonrpc: "2.0",
        id: "1",
        method: "LocalDB.JSONRPC_get_latest_block_number",
        params: [{ "Req_page": 0 }]
    }
  });
}

export function getLatestBlock() {
  return fetch2({
    url: "/rpc",
    method: "post",
    data: {
        jsonrpc: "2.0",
        id: "1",
        method: "LocalDB.JSONRPC_get_blocks_info",
        params: [{ "Req_page": 0 }]
    }
  });
}

// params: [{ "Req_page": <uint> }]
export function getLatestTransactions() {
  return fetch2({
    url: "/rpc",
    method: "post",
    data: {
        jsonrpc: "2.0",
        id: "1",
        method: "LocalDB.JSONRPC_get_txns_all_mixed",
        params: [{ "Req_page": 0 }]
    }
  });
}

// params: [{ "Req_page": <uint>, "Req_block": <string> }]
//export function getLatestTransactionsByBlockNumber(num) {
export function getTransactionsByBlockNumber(num) {
  return fetch2({
    url: "/rpc",
    method: "post",
    data: {
        jsonrpc: "2.0",
        id: "1",
        method: "LocalDB.JSONRPC_get_txns_all_mixed_by_block_number",
        params: [{ "Req_page": 0, "Req_block": String(num) }]
    }
  });
}

// params: [{ "Req_page": <uint>, "Req_block_hash": <string> }]
//export function getLatestTransactionsByBlockHash(block_hash) {
export function getTransactionsByBlockHash(block_hash) {
  return fetch2({
    url: "/rpc",
    method: "post",
    data: {
        jsonrpc: "2.0",
        id: "1",
        method: "LocalDB.JSONRPC_get_txns_all_mixed_by_block_hash",
        params: [{ "Req_page": 0, "Req_block_hash": String(block_hash) }]
    }
  });
}


// params: [{ "Req_txid": <string> }]
export function getTransactionsByTxid(txid) {
  return fetch2({
    url: "/rpc",
    method: "post",
    data: {
        jsonrpc: "2.0",
        id: "1",
        method: "LocalDB.JSONRPC_get_txns_all_mixed_by_txid",
        params: [{ "Req_txid": String(txid) }]
    }
  });
}

//! TODO: add into account's balance list
// params: [{ "Req_page": <uint>, "Req_address": <string> }]
export function getTransactionsByAddress(address) {
  return fetch2({
    url: "/rpc",
    method: "post",
    data: {
        jsonrpc: "2.0",
        id: "1",
        method: "LocalDB.JSONRPC_get_txns_all_mixed_by_address",
        params: [{ "Req_page": 0, "Req_address": String(address) }]
    }
  });
}

//! TODO:
// params: [{ "Req_address": <string> }]
export function getBalancesByAddress(address) {
  return fetch2({
    url: "/rpc",
    method: "post",
    data: {
        jsonrpc: "2.0",
        id: "1",
        method: "LocalDB.JSONRPC_get_balances_by_address",
        params: [{ "Req_address": String(address) }]
    }
  });
}

export function getBoardInfo() {
  return fetch2({
    url: "/rpc",
    method: "post",
    data: {
        jsonrpc: "2.0",
        id: "1",
        method: "LocalDB.JSONRPC_get_txns_all_mixed",
        params: [{ "Dummy": 0 }]
    }
  });
}



// --------------------------------------------------------



/*
export function getLatestBlock(num) {
  return fetch({
    method: "post",
    url: "LatestBlock",
    data: {
      num
    }
  });
}

export function getLatestMessage(num) {
  return fetch({
    method: "post",
    url: "LatestMsg",
    data: {
      num
    }
  });
}

export function getBoardInfo() {
  return fetch({
    method: "post",
    url: "BaseInformation"
  });
}
*/


/*
param:{
  start_time timestamp
  end_time timestamp
}
*/
export function getBlockTimeData(data) {
  return fetch({
    method: "post",
    url: "BlocktimeGraphical",
    data
  });
}
/*
param:{
  start_time timestamp
  end_time timestamp
}
*/
export function getBlocSizeData(data) {
  return fetch({
    method: "post",
    url: "AvgBlockheaderSizeGraphical",
    data
  });
}
/*
param:{
  time timestamp
}
*/
export function getTotalPowerData(data) {
  return fetch({
    method: "post",
    url: "/TotalPowerGraphical",
    data
  });
}
/*
param:{
  key string
  filter number
}
*/
export function search(data) {
  return fetch({
    method: "post",
    url: "SearchIndex",
    data
  });
}

export function getActivePeerCount() {
  return fetch({
    method: "post",
    url: "/peer/ActivePeerCount"
  });
}
