<template>
  <div class="message-detail bottom-10">
    <overview
      :dataList="dataList"
      :dataLabel="$t('message.detail.overview')"
      :showLoading="loading"
    />
  </div>
</template>

<!--
<script>
/*
import { getBlockConfirmCount, getMessageDetail } from "@/api/message";

export default {
  name: "MessageDetail",
  data() {
    return {
      cid: "",
      loading: false,
      dataList: [
        {
          key: "cid",
          style: {
            fontWeight: "bold"
          }
        },
        {
          key: "height",
          isLink: true,
          target: "tipset",
          isComponent: true
        },
        {
          key: "blockHash",
          isLink: true,
          target: "tipset",
          paramKey: "hash"
        },
        {
          key: "time"
        },
        {
          key: "from",
          isLink: true,
          target: "address/detail",
          paramKey: "address"
        },
        {
          key: "to",
          isLink: true,
          target: "address/detail",
          paramKey: "address"
        },
        {
          key: "value",
          //unit: "FIL"
        },
        // {
        //   key: "fee",
        //   unit: "FIL"
        // },
        //{
        //  key: "code"
        //},
        //{
        //  key: "method"
        //},
        //{
        //  key: "nonce"
        //},
        //{
        //  key: "params",
        //  isComponent: true
        //}
      ],
      title: {
        label: "Message",
        detail: `# ${this.$route.query.cid}`
      }
    };
  },

  methods: {
    async getMessageDetail() {
      try {
        this.loading = true;
        let data = await getMessageDetail({
          msg_cid: this.cid
        });
        const {
          height,
          cid,
          msgcreate,
          msg,
          block_cid,
          method_name,
          exit_code
        } = data.msg;
        const { from, to, nonce, params, value, gaslimit } = msg;
        let blockRes = await getBlockConfirmCount({
          cid: block_cid
        });

        const paramTip = this.$t("message.detail.paramTip");
        const confirm = this.$t("message.detail.confirm");
        const sourceMap = {
          height: this.formatNumber(height),
          cid,
          confirm: this.formatNumber(blockRes.count),
          time: this.getFormatTime(msgcreate),
          from,
          to,
          method: method_name,
          nonce,
          params: params.length > 256 ? `${params.slice(0, 256)} ...` : params,
          value,
          fee: this.formatNumber(gaslimit),
          blockHash: block_cid,
          code: exit_code
        };

        this.dataList = this.dataList.map(item => {
          let linkList;
          if (item.isLink) {
            linkList = [sourceMap[item.key]];
          } else {
            linkList = sourceMap[item.key];
          }
          let res = {
            ...item,
            value: sourceMap[item.key],
            linkList
          };
          if (item.key === "height") {
            res.component = {
              render() {
                return (
                  <div class="height-link">
                    <a href={`./#/tipset?height=${res.value}`}>
                      {sourceMap.height}
                    </a>
                    <span>
                      {" "}
                      ({sourceMap.confirm} {confirm})
                    </span>
                  </div>
                );
              }
            };
          }
          if (item.key === "params") {
            res.component = {
              render() {
                return (
                  <div class="top-10 params-con">
                    <el-popover
                      placement="bottom-start"
                      width="200"
                      trigger="hover"
                      content={paramTip}
                    >
                      <i class="el-icon-warning-outline" slot="reference"></i>
                    </el-popover>
                    <span class="params-value">{sourceMap[item.key]}</span>
                  </div>
                );
              }
            };
          }
          return res;
        });
        this.loading = false;
      } catch (e) {
        if (e) {
          this.loading = false;
        }
      }
    }
  },

  watch: {
    cid: {
      immediate: true,
      handler(v) {
        if (v) {
          this.getMessageDetail();
        }
    }
  },
  mounted() {
    this.cid = this.$route.query.cid;
  }
};
*/
</script>
-->


<!-- =========================================== -->


<script>
//import { getBlockConfirmCount, getMessageDetail } from "@/api/message";
import { getTransactionsByTxid } from "@/api/home";

export default {
  name: "MessageDetail",
  data() {
    return {
      //cid: "",
      loading: false,
      dataList: [
        {
            key: "cointype",
        },
        {
          key: "txid",
          style: {
            fontWeight: "bold"
          },
          paramKey: "txid"
        },
        {
          key: "height",
          isLink: true,
          target: "tipset",
          //isComponent: true
        },
        {
          key: "time"
        },
        {
          key: "contract",
          isLink: true,
          ellipsis: true,
        },
        {
          key: "from",
          isLink: true,
          target: "address/detail",
          paramKey: "address"
        },
        {
          key: "to",
          isLink: true,
          target: "address/detail",
          paramKey: "address"
        },
        {
          key: "token_id",
        },
        {
          key: "token_uri",
          isLink: true,
          ellipsis: true,
        },
        {
          key: "value",
          //unit: "FIL"
        },

/*
        {
          key: "cid",
          style: {
            fontWeight: "bold"
          }
        },
        {
          key: "height",
          isLink: true,
          target: "tipset",
          isComponent: true
        },
        {
          key: "blockHash",
          isLink: true,
          target: "tipset",
          paramKey: "hash"
        },
        {
          key: "time"
        },
        {
          key: "from",
          isLink: true,
          target: "address/detail",
          paramKey: "address"
        },
        {
          key: "to",
          isLink: true,
          target: "address/detail",
          paramKey: "address"
        },
        {
          key: "value",
          //unit: "FIL"
        },
        // {
        //   key: "fee",
        //   unit: "FIL"
        // },
        //{
        //  key: "code"
        //},
        //{
        //  key: "method"
        //},
        //{
        //  key: "nonce"
        //},
        //{
        //  key: "params",
        //  isComponent: true
        //}
*/
      ],
      title: {
        label: "Message",
        //detail: `# ${this.$route.query.cid}`
        detail: `# ${this.$route.query.txid}`
      }
    };
  },

  props: {
    txid: {
      type: String,
      default: ""
    },
    time: {
        type: String,
        default: ""
    },
    cointype: {
        type: String,
        default: ""
    },
    contract: {
        type: String,
        default: ""
    },
    from: {
        type: String,
        default: ""
    },
    to: {
        type: String,
        default: ""
    },
    token_id: {
        type: String,
        default: ""
    },
    token_uri: {
        type: String,
        default: ""
    },
    value: {
        type: String,
        default: ""
    },
    //value: {
    //    type: String,
    //    default: ""
    //},
    //fee: {
    //    type: String,
    //    default: ""
    //},
    //type: {
    //    type: String,
    //    default: ""
    //},
    //method: {
    //    type: String,
    //    default: ""
    //},
    height: {
        type: String,
        default: ""
    },

  },

  methods: {
    //! TODO: getTransactionsByTxid
    async getMessageDetail() {
        try {
            //this.loading = true;

            const ellipsisByLength = this.ellipsisByLength;

            console.log( "===== TODO: getTransactionsByTxid =====" );
            console.log( "this.txid = ", this.txid );

            let data = await getTransactionsByTxid( this.txid );
            const res_txndata = data;
            const res_txndata_json = JSON.parse(res_txndata.result);
            console.log( res_txndata_json );


            this.total = Number(res_txndata_json.length);

            let sourceMap = {};

            res_txndata_json.map(item => {
                const { from_address, to_address, amount_eth, amount_wei } = item;
                const { block_number, datetime, timestamp, token_amount, token_amount_eth } = item;
                const { token_amount_wei, token_contract_address, token_data } = item;
                const { token_data_length, token_decimals, token_symbol } = item;
                const { token_total_supply, token_type, token_id_ascii, token_uri_ascii } = item;
                const { token_uri_hexadecimal, txid } = item;
                let { value } = "";
                let disp_token_type = "";

                console.log( "block number = " + block_number );
                console.log( "transaction hash = " + txid );

                console.log( "datetime = " + datetime );
                console.log( "timestamp = " + timestamp );

                console.log( "token_type = " + token_type );

                console.log( "from = " + from_address );
                console.log( "to = " + to_address );


                if ( token_type == "ether" ) {
                    disp_token_type = "Ether";
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
                    console.log( "token_contract_address = " + token_contract_address );
                    console.log( "token_data = " + token_data );
                    console.log( "token_data_length = " + token_data_length );
                    console.log( "token_id_ascii = #" + token_id_ascii );
                    console.log( "token_url_ascii = " + token_uri_ascii );
                    console.log( "token_uri_hexadecimal = " + token_uri_hexadecimal );
                }


            sourceMap = {
                txid: txid,
                //time: this.formatTime(timestamp),
                time: this.getFormatTime(timestamp),
                //time: datetime,
                cointype: disp_token_type,

                contract: token_contract_address,
                //! FIXME
                /*
                contract: {
                  render() {
                      return (
                      <a
                        href={`./#/address/detail?address=${token_contract_address}`}
                        style={{ color: "var(--link-color)" }}
                      >
                        {ellipsisByLength(token_contract_address, 6, true)}
                      </a>
                      );

                      //return ( <span>{ellipsisByLength(token_contract_address, 6, true)}</span> );
                  }
                },
                */

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
                token_id: token_id_ascii,
                token_uri: token_uri_ascii,
                value: value,
                //value: this.formatFilNumber(value),
                //fee: gasprice,
                //type: this.address !== from ? "in" : "out",
                //method: method_name,
                height: this.formatNumber(block_number),
                //code: exit_code
            }
            sourceMap.from = from_address;
            sourceMap.to = to_address;
            });


            this.dataList = this.dataList.map(item => {
                let linkList;
                if (item.isLink) {
                    linkList = [sourceMap[item.key]];
                } else {
                    linkList = sourceMap[item.key];
                }
                let res = {
                    ...item,
                    value: sourceMap[item.key],
                    linkList
                };

                //res.component = {
                //    render() {
                //        return (
                //        );
                //    }
                //}

                return res;
            });

            //this.loading = false;


        } catch (e) {
            if (e) {
                this.loading = false;
            }
        }
    }

  },

  watch: {
    //cid: {
    //  immediate: true,
    //  handler(v) {
    //    if (v) {
    //      this.getMessageDetail();
    //    }
    //},

    //option: {
    //  //deep: true,
    //  immediate: true,
    //  handler() {
    //    this.getMessageDetail();
    //  }
    //},

    txid() {
      //immediate: true,
      //handler() {
          this.getMessageDetail();
      //}
    }
  },
  mounted() {
    //this.cid = this.$route.query.cid;
    this.txid = this.$route.query.txid;
  }


};
</script>






<style lang="scss" scoped>
.message-detail {
  & ::v-deep .height-link {
    width: 100%;
    display: flex;
    align-items: center;
    a {
      color: var(--link-color);
      margin-right: 5px;
    }
  }
  & ::v-deep .params-con {
    margin-right: 10px;
    .params-value {
      line-height: 1.5;
      word-break: break-all;
      margin-left: 3px;
    }
  }
}
</style>
