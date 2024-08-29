<template>
  <div
    class="message-list"
    v-loading="loading"
    element-loading-background="var(--board-bg-color)"
  >
    <div class="total-number border-bottom" slot="header">
      <span
        v-html="$t('component.mesList.total', { total: formatNumber(total) })"
      ></span>
      <el-select
        v-model="option.method"
        @change="handleMethodChange"
        :placeholder="$t('component.mesList.placeholder')"
      >
        <el-option label="All Method" value="" />
        <el-option
          v-for="item in method"
          :key="item.label"
          :label="item.label"
          :value="item.value"
        ></el-option>
      </el-select>
    </div>
    <div class="table-con" v-if="!isMobile">
      <base-table
        :dataSource="messageData"
        :columns="columns"
        showPagination
        :total="total"
        @size-change="handleSizeChange"
        @page-change="handlePageChange"
        :labels="labels"
        :currentPage="currentPage"
      ></base-table>
    </div>
<!--
    <mb-board
      v-for="(item, index) in messageData"
      :key="item.cid + index"
      :dataSource="item"
      :columns="mbColumns"
      v-else
    />
-->
    <mb-board
      v-for="(item, index) in messageData"
      :key="item.txid + index"
      :dataSource="item"
      :columns="mbColumns"
      v-else
    />

    <mb-page v-if="isMobile" @page-change="handlePageChange" :total="transactions" />
  </div>
</template>


<script>
/*
import {
  //getMessage,
  getMessageByAddress,
  getMessageMethods
} from "@/api/message";

export default {
  name: "MessageList",
  data() {
    const type = this.type;
    return {
      method: [],
      loading: false,
      option: {
        method: "",
        begindex: "0",
        count: "25"
      },
      currentPage: 1,
      total: 0,
      messageData: [],
      columns: [
        {
          key: "type",
          hideInMobile: true
        },
        {
          key: "cid",
          isLink: true,
          target: "message/detail",
          ellipsis: true
        },
        {
          key: "height",
          isLink: true,
          target: "tipset",
          paramKey: "height"
        },
        {
          key: "time",
          hideInMobile: true
        },
        {
          key: "from",
          isLink: true,
          target: "address/detail",
          paramKey: "address",
          ellipsis: true,
          isComponent: type === "address"
        },
        {
          key: "to",
          isLink: true,
          target: "address/detail",
          paramKey: "address",
          ellipsis: true,
          isComponent: type === "address"
        },
        {
          key: "value"
        },
        // {
        //   key: "fee",
        //   hideInMobile: true,
        //   unit: "FIL"
        // },
        {
          key: "code",
          hideInMobile: true
        },
        {
          key: "method"
        }
      ],
    };
  },
  props: {
    withType: {
      type: Boolean,
      default: true
    },
    cid: {
      type: String,
      default: ""
    },
    type: {
      type: String,
      default: "block"
    },
    address: {
      type: String,
      default: ""
    },
    blocks: {
      type: Array,
      default: () => {
        return [];
      }
    },

  },
  methods: {
    handleSizeChange(v) {
      this.option.count = v;
    },
    handlePageChange(v) {
      this.currentPage = v;
      this.option.begindex = (v - 1) * this.option.count;
    },
    handleMethodChange(v) {
      this.currentPage = 1;
      this.option = {
        method: v,
        begindex: 0,
        count: 25
      };
    },
    async getMessage() {
      try {
        this.loading = true;
        const addressHash = this.address;
        const type = this.type;
        const ellipsisByLength = this.ellipsisByLength;
        if (this.cid) {
          this.option.block_cid = this.cid;
        }
        let data = {};
        if (this.type === "block") {
          data = await getMessage(this.option);
        } else {
          this.columns;
          const res = await getMessageByAddress({
            ...this.option,
            address: this.address,
            from_to: ""
          });
          data.msgs = res.data;
          data.total = res.total;
        }
        this.total = Number(data.total);
        const messageData = data.msgs.map(item => {
          const { cid, msgcreate, msg, height, method_name, exit_code } = item;
          const { from, to, value, gasprice } = msg;
          let res = {
            cid: cid,
            time: this.formatTime(msgcreate),
            from: {
              render() {
                return from !== addressHash ? (
                  <a
                    href={`./#/address/detail?address=${from}`}
                    style={{ color: "var(--link-color)" }}
                  >
                    {ellipsisByLength(from, 6, true)}
                  </a>
                ) : (
                  <span>{ellipsisByLength(from, 6, true)}</span>
                );
              }
            },
            to: {
              render() {
                return to !== addressHash ? (
                  <a
                    href={`./#/address/detail?address=${to}`}
                    style={{ color: "var(--link-color)" }}
                  >
                    {ellipsisByLength(to, 6, true)}
                  </a>
                ) : (
                  <span>{ellipsisByLength(to, 6, true)}</span>
                );
              }
            },
            value: this.formatFilNumber(value),
            fee: gasprice,
            type: this.address !== from ? "in" : "out",
            method: method_name,
            height: this.formatNumber(height),
            code: exit_code
          };
          if (type === "block") {
            res.from = from;
            res.to = to;
          }
          return res;
        });
        this.messageData = Object.freeze(messageData);
        this.loading = false;
      } catch (e) {
        this.loading = false;
      }
    },
    async getMessageMethods() {
      try {
        let data = await getMessageMethods();
        this.method = data.method.map(item => {
          return {
            value: item,
            label: item
          };
        });
      } catch (e) {
        this.loading = false;
      }
    }
  },
  watch: {
    option: {
      deep: true,
      handler() {
        this.getMessage();
      }
    },
    cid() {
      this.option = {
        begindex: 0,
        count: 25
      };
    },
    address() {
      this.getMessage();
    }
  },
  mounted() {
    this.labels = [...this.$t("component.mesList.label")];
    if (!this.withType) {
      this.columns.shift();
      this.labels.shift();
    }
    this.getMessage();
    this.getMessageMethods();
  },
  computed: {
    mbColumns() {
      return this.columns
        .map((item, index) => {
          return {
            ...item,
            title: this.labels[index]
          };
        })
        .filter(item => {
          return !item.hideInMobile;
        });
    }
  }
};
*/



// ----------------------------------------------------------



//import {
//  //getMessage,
//  getMessageByAddress,
//  getMessageMethods
//} from "@/api/message";

//import { getBlockByHeight } from "@/api/tipset";
//import { getLatestBlockNumber } from "@/api/home";
//import { getTransactionsByBlockNumber } from "@/api/home";
import { getTransactionsByAddress } from "@/api/home";
//import { getBlockNumberByBlockHash } from "@/api/tipset";
import { getTransactionsByBlockHash } from "@/api/home";

export default {
  name: "MessageList",
  data() {
    const type = this.type;
    return {
      method: [],
      loading: false,
      option: {
        method: "",
        begindex: "0",
        count: "25"
      },
      currentPage: 1,
      total: 0,
      messageData: [],
      columns: [
        {
          key: "cointype",
          //hideInMobile: true
        },
        {
          key: "txid",
          isLink: true,
          target: "message/detail",
          ellipsis: true
        },
        {
          key: "height",
          isLink: true,
          target: "tipset",
          paramKey: "height"
        },
        {
          key: "time",
          //hideInMobile: true
        },
        {
          key: "tokencontract",
          isLink: true,
          target: "address/detail",
          paramKey: "address",
          ellipsis: true,
          isComponent: type === "address"
        },
        {
          key: "tokensymbol",
          //hideInMobile: true
        },
        {
          key: "from",
          isLink: true,
          target: "address/detail",
          paramKey: "address",
          ellipsis: true,
          isComponent: type === "address"
        },
        {
          key: "to",
          isLink: true,
          target: "address/detail",
          paramKey: "address",
          ellipsis: true,
          isComponent: type === "address"
        },
        {
          key: "value"
        },
        {
          key: "token_id"
        }
      ],

      labels: []
    };
  },
  props: {
    withType: {
      type: Boolean,
      default: true
    },
    //cid: {
    txid: {
      type: String,
      default: ""
    },
    type: {
      type: String,
      //default: "block"
      default: ""
    },
    address: {
      type: String,
      default: ""
    },
    blocks: {
      type: Array,
      default: () => {
        return [];
      }
    },

  },
  methods: {
    handleSizeChange(v) {
      this.option.count = v;
    },
    handlePageChange(v) {
      this.currentPage = v;
      this.option.begindex = (v - 1) * this.option.count;
    },
    handleMethodChange(v) {
      this.currentPage = 1;
      this.option = {
        method: v,
        begindex: 0,
        count: 25
      };
    },
    async getMessage() {
      try {
        this.loading = true;
        //const addressHash = this.address;
        //const type = this.type;
        const ellipsisByLength = this.ellipsisByLength;
        //if (this.cid) {
        if (this.txid) {
          //this.option.block_cid = this.cid;
          this.option.block_txid = this.txid;
        }
        let data = {};


        console.log( "========== message-list" );
        console.log( "type = ", this.type );
        console.log( "blocks = ", this.blocks );
        //console.log( "cid = ", this.cid );
        console.log( "txid = ", this.txid );
        console.log( "address = ", this.address );
        //console.log( "this = ", this );


        //const height = this.blocks.map((item, index) => {
        //  if ( item.key == "height" ) {
        //      return item.value;
        //  }
        //});
        let height = "";
        this.blocks.forEach(item => {
            if ( item.key == "height" ) {
                if ( typeof item.value !== "undefined" ) {
                    height = this.parseFormatNumber( item.value );
                    console.log( "height = ", height );
                }
                else {
                    console.log( "height = ", height );
                }
            }
        });


        let type_block_hash = false;
        this.blocks.forEach(item => {
            if ( item.key == "hash" ) {
                if ( typeof item.value !== "undefined" ) {
                    type_block_hash = true;
                }
            }
            if ( item.key == "height" ) {
                if ( typeof item.value !== "undefined" ) {
                    type_block_hash = true;
                }
            }
        });
        //if ( this.cid != "" ) {
        if ( this.txid != "" ) {
            type_block_hash = true;
        }
        if ( type_block_hash == true ) {
            this.type = "block";
        }
        else {
            this.type = "address";
        }

        console.log( "type = ", this.type );




        if (this.type === "block") {
          //data = await getMessage(this.option);
          //data_blocks = await getBlockByHeight( height );
          //data = await getTransactionsByBlockNumber( height );

          console.log( "block number" );
          //data = await getBlockNumberByBlockHash( this.cid );
          data = await getTransactionsByBlockHash( this.txid );

          console.log( "block number = ", data );

        }
        else {
          this.columns;
          /*
          const res = await getMessageByAddress({
            ...this.option,
            address: this.address,
            from_to: ""
          });
          data.msgs = res.data;
          data.total = res.total;
          */

          //data = await getTransactionsByAddress( this.cid );
          data = await getTransactionsByAddress( this.txid );
        }



        /*
        let data_blocks = {};
        const res_blockdata = data_blocks;
        const res_blockdata_json = JSON.parse(res_blockdata.result);

        //console.log( "block number = " + height );
        //console.log( "========== #1" );
        //console.log( res_blockdata );
        //console.log( "========== #2" );
        //console.log( res_blockdata.result );
        //console.log( "========== #3" );
        //console.log( res_blockdata_json );
        //console.log( "========== #4" );
        //console.log( res_blockdata_json[0].block_number);
        //console.log( res_blockdata_json[0].block_hash);
        //console.log( res_blockdata_json[0].transactions);
        //console.log( "========== #5" );

        //const aaa = res_blockdata_json.map(item => item);
        //console.log( aaa );
        //console.log( "========== end..." );

        const blockData = res_blockdata_json.map(item => {
          //const { cid, msgcreate, msg, height, method_name, exit_code } = item;
          //const { from, to, value, gasprice } = msg;

          const { block_number, block_hash, block_info, transactions } = item;

          console.log( "block number = " + block_number );
          console.log( "block hash = " + block_hash );
          console.log( "transactions = " + transactions );

          const res_blockinfo_json = JSON.parse(block_info);
          //console.log( res_blockinfo_json );
          //console.log( "========== #1" );

          for ( let i = 0; i < transactions; i++ ) {
            let tx_blockhash = res_blockinfo_json.transactions[i].blockHash;
            let tx_from = res_blockinfo_json.transactions[i].from;
            let tx_to = res_blockinfo_json.transactions[i].to;
            let tx_contract_address = "";
            let tx_value = res_blockinfo_json.transactions[i].value;

            console.log( "block hash = " + tx_blockhash );
            console.log( "from = " + tx_from );
            if ( tx_value == "0x0" ) {
              tx_contract_address = tx_to;
              tx_to = "";
              console.log( "contract address = " + tx_contract_address );
              console.log( "to = " + tx_to );
              console.log( "value = " + tx_value );
              //console.log( res_blockinfo_json.transactions[i].input );
            }
            else {
              tx_value = parseInt( tx_value, 16 ) / 1000000000000000000;
              console.log( "to = " + tx_to );
              console.log( "value = " + Number(tx_value).toString() + " Ether" );
            }
          }

          return res;
        });
        */



        console.log( "---------------------" );
        console.log( "transactions" );
        console.log( "---------------------" );

        const res_txndata = data;
        const res_txndata_json = JSON.parse(res_txndata.result);

        console.log( res_txndata_json );

        this.total = Number(res_txndata_json.length);
        const messageData = res_txndata_json.map(item => {
          const { from_address, to_address, amount_eth, amount_wei } = item;
          const { block_number, datetime, timestamp, token_amount, token_amount_eth } = item;
          const { token_amount_wei, token_contract_address, token_data } = item;
          const { token_data_length, token_decimals, token_symbol } = item;
          const { token_total_supply, token_type, token_id_ascii, token_uri_ascii } = item;
          const { token_uri_hexadecimal, txid } = item;
          let { value } = "";
          let disp_token_type = "";
          let disp_token_symbol = token_symbol;

          console.log( "block number = " + block_number );
          console.log( "transaction hash = " + txid );

          console.log( "datetime = " + datetime );
          console.log( "timestamp = " + timestamp );

          console.log( "token_type = " + token_type );

          console.log( "from = " + from_address );
          console.log( "to = " + to_address );


          if ( token_type == "ether" ) {
            disp_token_type = "Ether";
            disp_token_symbol = "ETH";
            value = amount_eth + " Eth";
            console.log( "amount_eth = " + amount_eth );
            console.log( "amount_wei = " + amount_wei );
          }
          else if ( token_type == "erc20" ) {
            disp_token_type = "ERC-20";
            value = token_amount_eth + " " + token_symbol;
            console.log( "token_amount_eth = " + token_amount_eth );
            console.log( "tonen_amount_wei = " + token_amount_wei );
            console.log( "token_contract_address = " + token_contract_address );
            console.log( "token_decimals = " + token_decimals );
            console.log( "token_symbol = " + token_symbol );
            console.log( "token_total_supply = " + token_total_supply );
          }
          else if ( token_type == "erc1155" ) {
            disp_token_type = "ERC-1155";
            value = token_amount;
            console.log( "token_amount = " + token_amount );
            console.log( "token_data = " + token_data );
            console.log( "token_data_length = " + token_data_length );
            console.log( "token_id_ascii = " + token_id_ascii );
            console.log( "token_url_ascii = " + token_uri_ascii );
            console.log( "token_uri_hexadecimal = " + token_uri_hexadecimal );
          }

          console.log( "---------------------" );


          let res = {
            txid: txid,
            //time: this.formatTime(timestamp),
            time: this.getFormatTime(timestamp),
            //time: datetime,
            cointype: disp_token_type,
            from: {
              render() {
                  return (
                  <a
                    href={`./#/address/detail?address=${from_address}`}
                    style={{ color: "var(--link-color)" }}
                  >
                    {ellipsisByLength(from_address, 6, true)}
                  </a>
                  );

                  //return ( <span>{ellipsisByLength(from_address, 6, true)}</span> );
              }
            },
            to: {
              render() {
                  return (
                  <a
                    href={`./#/address/detail?address=${to_address}`}
                    style={{ color: "var(--link-color)" }}
                  >
                    {ellipsisByLength(to_address, 6, true)}
                  </a>
                  );

                  //return ( <span>{ellipsisByLength(to_address, 6, true)}</span> );
              }
            },
            value: value,
            token_id: token_id_ascii,
            tokensymbol: disp_token_symbol,
            tokencontract: token_contract_address,
            //value: this.formatFilNumber(value),
            //fee: gasprice,
            //type: this.address !== from ? "in" : "out",
            //method: method_name,
            height: this.formatNumber(block_number),
            //code: exit_code
          };
          //if (type === "block") {
            res.from = from_address;
            res.to = to_address;
          //}

          return res;
        });


        this.messageData = Object.freeze(messageData);
        this.loading = false;
      } catch (e) {
        this.loading = false;
      }
    },
    /*
    async getMessageMethods() {
      try {
        let data = await getMessageMethods();
        this.method = data.method.map(item => {
          return {
            value: item,
            label: item
          };
        });
      } catch (e) {
        this.loading = false;
      }
    }
    */
  },


  watch: {
    option: {
      deep: true,
      handler() {
        this.getMessage();
      }
    },
    //cid() {
    txid() {
      this.option = {
        begindex: 0,
        count: 25
      };
    },
    address() {
      this.getMessage();
    }
  },
  mounted() {
    this.labels = [...this.$t("component.mesList.label")];
    if (!this.withType) {
      this.columns.shift();
      this.labels.shift();
    }
    this.getMessage();
    //this.getMessageMethods();
  },
  computed: {
    mbColumns() {
      return this.columns
        .map((item, index) => {
          return {
            ...item,
            title: this.labels[index]
          };
        })
        .filter(item => {
          return !item.hideInMobile;
        });
    }
  }
};


</script>



<style lang="scss" scoped>
.message-list {
  .total-number {
    height: 80px;
    align-items: center;
    padding: 0 100px;
    display: flex;
    background: var(--board-bg-color);
    color: var(--main-text-color);
    & ::v-deep > span {
      margin-right: auto;
      i {
        color: var(--link-color);
        font-size: 22px;
      }
    }
    .el-dropdown-link {
      color: white;
    }
  }
  @media (max-width: 768px) {
    .total-number {
      height: 30px;
      margin-bottom: 10px;
      & ::v-deep .el-select input {
        height: 20px;
        font-size: 12px !important;
        width: 100px;
      }
      & ::v-deep .el-input__suffix {
        i {
          line-height: 20px;
          padding-left: 5px;
          &.is-reverse {
            padding-right: 10px;
          }
        }
      }
    }
  }
}
</style>
<style lang="scss">
@media (max-width: 768px) {
  .el-select-dropdown {
    z-index: 10000 !important;
    .el-select-dropdown__wrap {
      max-height: 160px !important;
    }
    li {
      height: 20px !important;
      line-height: 20px !important;
      font-size: 12px;
    }
  }
}
</style>
