<template>
  <div class="block-detail bottom-10 top-20">

    <!-- <div> blocks = {{ realList }}</div> -->

    <div class="block-overview">
      <!--<overview :dataList="realList" :dataLabel="$t('tipset.block.overview')" />-->
      <overview :dataList="dataList" :dataLabel="$t('tipset.block.overview')" />
    </div>
    <!--<message-list :withType="true" :cid="hash" :blocks="realList" />-->
    <message-list :withType="true" :txid="hash" :blocks="realList" />
  </div>
</template>

<script>
/*
export default {
  name: "BlockDetail",
  props: {
    hash: {
      type: String,
      default: ""
    },
    block: {
      type: Object,
      default() {
        return {};
      }
    }
  },
  data() {
    return {

//      dataList: [
//        {
//          key: "hash",
//          style: {
//            fontWeight: "bold"
//          }
//        },
//        {
//          key: "height",
//          isLink: true,
//          target: "tipset"
//        },
//        {
//          key: "utcTime"
//        },
//        {
//          key: "size",
//          unit: "bytes"
//        },
//        {
//          key: "mesLength"
//        },
//        {
//          key: "miner",
//          target: "address/detail",
//          paramKey: "address",
//          isLink: true
//        },
//        {
//          key: "reward",
//          unit: "FIL"
//        },
//        {
//          key: "parents",
//          isLink: true,
//          target: "tipset",
//          paramKey: "hash"
//        },
//        {
//          key: "parent_weight"
//        },
//        {
//          key: "tickets"
//        },
//        {
//          key: "state_root"
//        }
//      ]



      dataList: [
        {
          key: "hash",
          style: {
            fontWeight: "bold"
          }
        },
        {
          key: "height"
        },
        {
          key: "timestamp"
        },
        {
          key: "transactions"
        }

      ]

    };
  },
  computed: {
    realList() {
      const currentBlock = this.block;
      console.log( "BlockDetail: realList()" );
      console.log( currentBlock );
      return this.dataList.map(item => {
        let linkList;
        console.log( item.key, item.value );
        if (item.key === "height" || item.key === "miner") {
          linkList = [currentBlock[item.key]];
        } else {
          linkList = currentBlock[item.key];
        }
        return {
          ...item,
          value: currentBlock[item.key],
          linkList: linkList
        };
      });
    }
  }
};
*/



// ----------------------------------------------------------



import { getTransactionsByBlockHash } from "@/api/home";

export default {
  name: "BlockDetail",
  props: {
    hash: {
      type: String,
      default: ""
    },
/*
    blocks: {
      type: Object,
      default() {
        return {};
      }
    }
*/
    //blocks: { type: Array, default() { return []; } }
    //block: { type: Array, default() { return []; } }

    block: { type: Object, default() { return {}; } }
  },
  height: {
    type: String,
    default: ""
  },

  data() {
    return {

//      dataList: [
//        {
//          key: "hash",
//          style: {
//            fontWeight: "bold"
//          }
//        },
//        {
//          key: "height",
//          isLink: true,
//          target: "tipset"
//        },
//        {
//          key: "utcTime"
//        },
//        {
//          key: "size",
//          unit: "bytes"
//        },
//        {
//          key: "mesLength"
//        },
//        {
//          key: "miner",
//          target: "address/detail",
//          paramKey: "address",
//          isLink: true
//        },
//        {
//          key: "reward",
//          unit: "FIL"
//        },
//        {
//          key: "parents",
//          isLink: true,
//          target: "tipset",
//          paramKey: "hash"
//        },
//        {
//          key: "parent_weight"
//        },
//        {
//          key: "tickets"
//        },
//        {
//          key: "state_root"
//        }
//      ]



      dataList: [
        {
          key: "hash",
          style: {
            fontWeight: "bold"
          }
        },
        {
          key: "height"
        },
        {
          key: "contract"
        },
        {
          key: "timestamp"
          //key: "datetime"
        },
        {
          key: "transactions"
        }
      ],

    };
  },

  watch: {
    option: {
      deep: true,
      immediate: true,
      handler() {
        this.getMessage();
      }
    },
  },

  methods: {
    async getMessage() {
      try {
        let data = {};
        data = await getTransactionsByBlockHash( this.hash );

        console.log( "BlockDetail::getMessage(): block data = " );
        console.log( data );

        const res_txndata_json = JSON.parse(data.result);
        console.log( "BlockDetail::getMessage(): block jSON data = " );
        console.log( res_txndata_json );

        const dataList = [
            { "key": "hash", "style": { "fontWeight": "bold" }, "value": this.hash, "linkList": this.hash },
            { "key": "height", "value": res_txndata_json[0].block_number, "linkList": [res_txndata_json[0].block_number] },
            { "key": "timestamp", "value": "" },
            { "key": "transactions", "value": res_txndata_json.length }
        ];
        console.log( "new dataList =" );
        console.log( dataList );
        this.dataList = Object.freeze( dataList );

        this.dataList.map(item => {
          let linkList;
          console.log( item.key, item.value );

          //if (item.key === "height" || item.key === "miner") {
          //  linkList = [currentBlock[item.key]];
          //} else {
          //  linkList = currentBlock[item.key];
          //}

          return {
            ...item,
            value: dataList[item.key],
            linkList: linkList
          };
        });


        console.log( "new this.dataList =" );
        console.log( this.dataList );
        return this.dataList;

      } catch (e) {
        console.log( e );
      }
    },
  },

  computed: {
    realList() {
      const currentBlock = this.block;
      console.log( "===== BlockDetail =====" );
      console.log( "BlockDetail: realList()" );
      console.log( currentBlock );
      console.log( "len = " + currentBlock.length );
      console.log( "hash = " + this.hash );
      console.log( "-----" );


      return this.dataList.map(item => {
        let linkList;
        console.log( item.key, item.value );
        if (item.key === "height" || item.key === "miner") {
          linkList = [currentBlock[item.key]];
        } else {
          linkList = currentBlock[item.key];
        }
        return {
          ...item,
          value: currentBlock[item.key],
          linkList: linkList
        };
      });

    }
  }
};

</script>



<style lang="scss" scoped>
.block-detail {
  position: relative;
  .block-hash {
    background: white;
  }
  & ::v-deep .message-list {
    margin-top: 20px;
  }
}
</style>
