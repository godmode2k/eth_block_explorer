import { fetch } from "./fetch";
import { fetch2 } from "./fetch";


// params: [{ "Req_block": <string> }]
export function getBlockByHeight(num) {
  return fetch2({
    url: "/rpc",
    method: "post",
    data: {
        jsonrpc: "2.0",
        id: "1",
        method: "LocalDB.JSONRPC_get_block_by_number",
        params: [{ "Req_block": String(num) }]
    }
  });
}

// params: [{ "Req_block_hash": <string> }]
export function getBlockNumberByBlockHash(block_hash) {
  return fetch2({
    url: "/rpc",
    method: "post",
    data: {
        jsonrpc: "2.0",
        id: "1",
        method: "LocalDB.JSONRPC_get_block_number_by_block_hash",
        params: [{ "Req_block_hash": String(block_hash) }]
    }
  });
}



// --------------------------------------------------------



/*
param:{
  count number
  end_height number
}
*/
export function getTipset(data) {
  return fetch({
    url: "tipset/TipSetTree",
    method: "post",
    data
  });
}

/*
param:{
  height number
}
*/
/*
export function getBlockByHeight(data) {
  return fetch({
    url: "tipset/BlockByHeight",
    method: "post",
    data
  });
}
*/

/*
param:{
  cid number
}
*/
export function getBlockByCid(data) {
  return fetch({
    url: "tipset/BlockByCid",
    method: "post",
    data
  });
}
