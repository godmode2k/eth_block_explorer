<template>
  <div class="address-detail bottom-10">
    <overview
      :dataList="dataList"
      :dataLabel="$t('address.detail.overview')"
      class="bottom-20"
    />

<!--
    <div class="worker-list bottom-20" v-if="workers.length">
      <span>{{ $t("address.detail.worker") }}</span>
      <span>
        <base-link
          v-for="link in workers"
          :key="link"
          :label="link"
          target="address/detail"
          :param="{
            key: 'address',
            value: link
          }"
        ></base-link>
      </span>
    </div>
    <overview
      :dataList="accountList"
      :dataLabel="$t('address.detail.ownerOverview')"
      class="bottom-20"
      v-if="isMiner"
    />
    <el-radio-group
      v-model="showMessage"
      style="margin-bottom: 20px;"
      v-if="isMiner || workers.length"
    >
      <el-radio-button :label="true">
        {{ $t("address.radio")[0] }}
      </el-radio-button>
      <el-radio-button :label="false">
        {{ $t("address.radio")[1] }}
      </el-radio-button>
    </el-radio-group>
    <block-list v-else :miners="address" />
-->


<!-- =========================================== -->


<!--
    <message-list
      v-if="showMessage"
      :address="$route.query.address"
      type="address"
    />
-->
  <div
    class="txns-list"
    v-loading="loading"
    element-loading-background="var(--board-bg-color)"
  >
    <div class="total-number border-bottom" slot="header">
    <!--
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
    -->
    </div>

    <div class="table-con" v-if="!isMobile">
      <base-table
        :dataSource="txnsList"
        :columns="columns"
        showPagination
        :total="total"
        @size-change="handleSizeChange"
        @page-change="handlePageChange"
        :labels="labels"
        :currentPage="currentPage"
      ></base-table>
    </div>
    <mb-board
      v-for="(item, index) in txnsList"
      :key="item.cid + index"
      :dataSource="item"
      :columns="mbColumns"
      v-else
    />
    <mb-page v-if="isMobile" @page-change="handlePageChange" :total="transactions" />
  </div>


  </div>
</template>



<!--
<script>
/*
import { getActorById } from "@/api/account";
import mixin from "./mixin";

export default {
  name: "AddressDetail",
  mixins: [mixin],
  data() {
    return {
      showMessage: true,
      isMiner: false,
      isOwner: false,
      messageData: [],
      address: "",
      workers: [],

      dataList: [
        {
          key: "address"
        },
        {
          key: "type"
        },
        {
          key: "balance",
          unit: "FIL"
        },
        {
          key: "token"
        },
        {
          key: "code"
        },
        {
          key: "nonce"
        }
      ],
      accountList: [
        {
          key: "owner_address",
          isLink: true,
          target: "address/detail",
          paramKey: "address"
        },
        {
          key: "peer_id",
          isLink: true,
          target: "stats/peer"
        },
        {
          key: "power"
        },
        {
          key: "sector_size",
          unit: "bytes"
        },
        {
          key: "sector_num"
        },
        {
          key: "proving_sector_num"
        },
        {
          key: "fault_num"
        }
      ]
  },
  watch: {
    "$route.query.address": {
      immediate: true,
      handler(v) {
        if (!v) {
          return;
        }
        this.address = [v];
        this.getAddressInfo(v);
      }
    }
  },
  methods: {
    async getAddressInfo(a) {
      try {
        let res = await getActorById({
          actor_id: a
        });
        const detail = this.parseAddress(res.data);
        this.dataList = this.dataList.map(item => {
          return {
            ...item,
            value: detail[item.key]
          };
        });
        this.workers = res.work_list;
        if (res.data.is_miner && res.miner.owner_address != "") {
          this.isMiner = true;
        } else {
          this.isMiner = false;
        }
        if (res.work_list.length) {
          this.address = res.work_list;
        }
        this.accountList = this.accountList.map(item => {
          let linkList;
          const originValue = res.miner[item.key];
          if (item.key === "owner_address" || item.key === "peer_id") {
            linkList = [originValue];
          } else {
            linkList = originValue;
          }
          const isNumber = parseFloat(originValue) == originValue;
          let result = {
            ...item,
            value: isNumber ? this.formatNumber(originValue, 18) : originValue,
            linkList: linkList
          };
          if (item.key === "power") {
            result.value = `${originValue} bytes (${this.unitConversion(
              originValue
            )})`;
          }
          return result;
        });
      } catch (e) {
        console.log(e);
      }
    }
    async getTransactionsByAddress() {
      try {
        let data = await getTransactionsByAddress(this.address);
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
  }
};
*/
</script>
-->


<!-- =========================================== -->


<script>
import mixin from "./mixin";
import { getTransactionsByAddress } from "@/api/home"
//import { getBalancesByAddress } from "@/api/home"
//! TODO: get balances by address
// - each tokens, total balances, contract address, ...

export default {
  name: "AddressDetail",
  mixins: [mixin],
  data() {
    return {
      showMessage: true,
      isMiner: false,
      isOwner: false,
      messageData: [],
      address: "",
      workers: [],

      loading: false,
      currentPage: 1,
      total: 0,
      //method: [],

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
          //key: "datetime",
          //hideInMobile: true
        },
        {
          key: "from",
          isLink: true,
          target: "address/detail",
          paramKey: "address",
          ellipsis: true,
          //isComponent: type === "address"
        },
        {
          key: "to",
          isLink: true,
          target: "address/detail",
          paramKey: "address",
          ellipsis: true,
          //isComponent: type === "address"
        },
        {
          key: "value"
        }
      ],

      dataList: [
        {
          key: "address"
        },
        {
          key: "balance"
        },
        {
          key: "token"
        },
      ],

      txnsList: [
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
          //key: "datetime",
          //hideInMobile: true
        },
        {
          key: "from",
          isLink: true,
          target: "address/detail",
          paramKey: "address",
          ellipsis: true,
          //isComponent: type === "address"
        },
        {
          key: "to",
          isLink: true,
          target: "address/detail",
          paramKey: "address",
          ellipsis: true,
          //isComponent: type === "address"
        },
        {
          key: "value"
        }
      ],

      labels: [],

    };
  },
  watch: {
    "$route.query.address": {
      immediate: true,
      handler(v) {
        if (!v) {
          return;
        }
        this.address = [v];
        this.getAddressInfo(v);
      }
    }
  },
  methods: {
    async getAddressInfo(a) {
      console.log( "getAddressInfo()" );
      console.log( "----------------" );
      console.log( "a = " + a );

      try {
        let res = await getTransactionsByAddress( String(a) );

        const data = res;
        const res_json = JSON.parse(data.result);

        this.total = Number(res_json.length);

        ///*
        console.log( "========== #1" );
        console.log( data );
        console.log( "========== #2" );
        console.log( data.result );
        console.log( "========== #3" );
        console.log( res_json );
        console.log( "========== #4" );
        //console.log( res_json[0].);
        console.log( "========== #5" );
        //const aaa = res_json.map(item => item);
        //console.log( aaa );
        console.log( "========== end..." );
        //*/

        console.log( "===== TODO: getAddressInfo() =====" );
        const txns = res_json.map(item => {
          /*
          const { block_number, datetime, timestamp, txid } = item;
          const { from_address, to_address, token_type, amount_eth, amount_wei } = item;
          const { token_amount, token_amount_eth, token_amount_wei, token_contract_address } = item;
          const { token_decimals, token_symbol, token_total_supply } = item;
          const { token_data, token_data_length } = item;
          const { token_uri_ascii, token_uri_hexadecimal } = item;

          let res = {
            block_number: block_number,
            datetime: datetime,
            timestamp: timestamp,
            txid: txid,

            from_address: from_address,
            to_address: to_address,
            cointype: token_type,
            amount_eth: amount_eth,
            amount_wei: amount_wei,

            token_amount: token_amount, 
            token_amount_eth: token_amount_eth,
            token_amount_wei: token_amount_wei,
            token_contract_address: token_contract_address,

            token_decimals: token_decimals, 
            token_symbol: token_symbol, 
            token_total_supply: token_total_supply,

            token_data: token_data,
            token_data_length: token_data_length,

            token_uri_ascii: token_uri_ascii, 
            token_uri_hexadecimal: token_uri_hexadecimal
          };
          */


          ///*
          const { txid, timestamp, datetime, token_type, from_address, to_address, block_number } = item;
          const { token_symbol, amount_eth, token_amount, token_amount_eth } = item;
          let { value } = "";
          let disp_token_type = "";

          console.log( "datetime = " + datetime );
          console.log( "timestamp = " + timestamp );

          if ( token_type == "ether" ) {
            disp_token_type = "Ether";
            value = amount_eth + " " + token_symbol;
          }
          else if ( token_type == "erc20" ) {
            disp_token_type = "ERC-20";
            value = token_amount_eth + " " + token_symbol;
          }
          else if ( token_type == "erc1155" ) {
            disp_token_type = "ERC-1155";
            value = token_amount + " " + token_symbol;
          }

          const ellipsisByLength = this.ellipsisByLength;

          let res = {
            cointype: disp_token_type,
            txid: txid,
            height: this.formatNumber(block_number),
            //time: this.formatTime(timestamp),
            time: this.getFormatTime(timestamp),
            //time: datetime,
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
            value: value
          };

          res.from = from_address;
          res.to = to_address;
          //*/

          return res;
        });
        this.txnsList = txns;


/*
        // account's balance list
        let dataList = res_json.map(item => {
          const { token_type, amount_eth, token_amount, token_amount_eth } = item;
          let { value } = "";

          if ( token_type == "ether" ) {
            value = amount_eth;
          }
          else if ( token_type == "erc20" ) {
            value = token_amount_eth;
          }
          else if ( token_type == "erc1155" ) {
            value = token_amount;
          }

          // tokens...

          return {
            linkList: this.address[0],
            balance: "empty",
            //token: ""
          };
        });
*/
        //this.dataList = dataList;
        console.log( "dataList = " );
        console.log( this.dataList );
        console.log( this.dataList.address );



        /*
        const detail = this.parseAddress(res.data);
        this.dataList = this.dataList.map(item => {
          return {
            ...item,
            value: detail[item.key]
          };
        });
        this.workers = res.work_list;
        if (res.data.is_miner && res.miner.owner_address != "") {
          this.isMiner = true;
        } else {
          this.isMiner = false;
        }
        if (res.work_list.length) {
          this.address = res.work_list;
        }
        */

        //this.txnsList = Object.freeze(txns);
        //this.columns = Object.freeze(txns);

      } catch (e) {
        console.log(e);
      }
    }, // getAddressInfo()
  },


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


  mounted() {
    this.labels = [...this.$t("component.mesList.label")];
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


<!-- =========================================== -->


<style lang="scss" scoped>
.address-detail {
  & ::v-deep .el-radio-group label {
    display: inline-block;
    width: 150px;
    span {
      width: 100%;
    }
  }
  .worker-list {
    min-height: 60px;
    display: flex;
    background: var(--board-bg-color);
    box-shadow: 0px 1px 5px 7px rgba(0, 0, 0, 0.03);
    border-radius: 4px;
    color: var(--main-text-color);
    span {
      line-height: 60px;
    }
    span:first-child {
      padding-left: 100px;
      min-width: 200px;
    }
    span:last-child {
      flex: 1;
      word-break: break-all;
    }
    & ::v-deep a span {
      margin-right: 10px;
    }
  }
  @media (max-width: 768px) {
    .worker-list {
      box-shadow: 0px 2px 4px 0px rgba(0, 0, 0, 0.03);
      border-radius: 4px;
      margin: 10px 0;
      span {
        height: 30px;
        line-height: 30px;
        &:first-child {
          margin-right: 10px;
        }
        a {
          margin-right: 5px;
        }
      }
    }
    & ::v-deep .el-radio-group {
      display: flex !important;
      border-radius: 4px !important;
      label {
        flex: 1;
        height: 30px !important;
        &:first-child span {
          border-radius: 4px 0 0 4px;
        }
        &:last-child span {
          border-radius: 0 4px 4px 0;
        }
        span {
          height: 100% !important;
          line-height: 30px !important;
          padding: 0;
        }
      }
    }
    & ::v-deep > div:nth-child(2).general-overview {
      margin: 10px 0 !important;
    }
  }
}
</style>


<!-- =========================================== -->


<!-- txns-list -->

<style lang="scss" scoped>
.txns-list {
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
