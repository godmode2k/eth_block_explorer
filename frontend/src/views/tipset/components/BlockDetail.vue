<template>
  <div class="block-detail bottom-10 top-20">

    <!-- <div> blocks = {{ realList }}</div> -->

    <div class="block-overview">
      <overview :dataList="realList" :dataLabel="$t('tipset.block.overview')" />
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
