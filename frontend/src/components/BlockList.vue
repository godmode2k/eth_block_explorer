<template>
  <div class="block-list">
    <div class="table-con" v-if="!isMobile">
      <base-table
        :dataSource="blockData"
        :columns="realColumns"
        showPagination
        :total="total"
        @size-change="handleSizeChange"
        @page-change="handlePageChange"
        :labels="realLabels"
      >
        <div class="total-number border-bottom" slot="header">
          <span
            v-html="
              $t('component.blockList.total', {
                total: formatNumber(total),
                fil: formatNumber(filTotal)
              })
            "
          ></span>
        </div>
      </base-table>
    </div>
    <mb-board
      v-for="item in blockData"
      :key="item.cid"
      :dataSource="item"
      :columns="mbColumns"
      v-else
    />
    <mb-page v-if="isMobile" @page-change="handlePageChange" :total="total" />
  </div>
</template>
<script>
import { getBlockByHeight } from "@/api/tipset";
//import { getBlocksByMiner } from "@/api/account";
export default {
  name: "BlockList",
  data() {
    return {
      loading: false,
      total: 0,
      filTotal: 0,
      option: {
        begindex: "0",
        count: "25",
        miners: this.miners
      },
      blockData: [],
      /*
      columns: [
        {
          key: "height",
          isLink: true,
          target: "tipset",
          hideInMobile: true
        },
        {
          key: "hash",
          isLink: true,
          target: "tipset",
          ellipsis: true
        },
        {
          key: "time"
        },
        {
          key: "mes"
        },
        {
          key: "size"
        },
        {
          key: "miner",
          isLink: true,
          target: "address/detail",
          paramKey: "address"
        },
        {
          key: "reward",
          unit: "FIL"
        }
      ]
      */


      columns: [
        {
          key: "height",
          isLink: true,
          target: "tipset",
          hideInMobile: true
        },
        {
          key: "hash",
          isLink: true,
          target: "tipset",
          ellipsis: true
        },
        //{
        //  key: "time"
        //}
      ]

    };
  },
  props: {
    hideIndex: {
      type: Array,
      default() {
        return [];
      }
    },
    height: {
      type: Number,
      default: 0
    },
    /*
    miners: {
      type: Array,
      default() {
        return [];
      }
    }
    */
  },
  computed: {
    realColumns() {
      return this.columns.filter((item, index) => {
        return !this.hideIndex.includes(index);
      });
    },
    realLabels() {
      return this.$t("component.blockList.label").filter((item, index) => {
        return !this.hideIndex.includes(index);
      });
    },
    mbColumns() {
      return this.realColumns
        .map((item, index) => {
          return {
            ...item,
            title: this.realLabels[index]
          };
        })
        .filter(item => {
          return !item.hideInMobile;
        });
    }
  },
  methods: {
    handleSizeChange(v) {
      this.option.count = v;
    },
    handlePageChange(v) {
      this.option.begindex = (v - 1) * this.option.count;
    },
    async getBlockByHeight() {
      console.log( "========== getBlockByHeight()" );
      try {
        this.loading = true;
        /*
        let data = await getBlockByHeight({
          height: this.height
        });
        */

        let data = await getBlockByHeight( this.height );

        console.log( "========== #1" );
        console.log( data );
        console.log( "========== #2" );
        console.log( data.result );
        console.log( "========== #3" );
        const res_json = JSON.parse(data.result);
        console.log( res_json );
        console.log( "========== #4" );


        //const blockData = data.blocks.map(item => {
        //  const { height } = item;
        const blockData = JSON.parse(data.result).map(item => {
          const { block_number, block_hash, block_info/*, transactions*/ } = item;
          const block_info_json = JSON.parse(block_info)
          return {
            //height

            height: this.formatNumber(block_number),
            hash: block_hash,
            //timestamp: this.formatTime(parseInt(block_info_json.timestamp, 16)),
            timestamp: this.formatTime(block_info_json.timestamp),
          };
        });
        this.total = blockData.length;
        this.blockData = blockData;
      } catch (e) {
        if (e) {
          this.loading = false;
        }
      }
    },
    /*
    async getBlockByMiner() {
      try {
        this.loading = true;

        let data = await getBlocksByMiner(this.option);
        const dataSource = data.blocks.map(item => {
          const { block_header, size, reward, cid, msg_cids } = item;
          const { height, timestamp, miner } = block_header;
          return {
            height,
            hash: cid,
            time: this.formatTime(timestamp),
            mes: msg_cids.length,
            size,
            miner,
            reward
          };
        });
        this.blockData = Object.freeze(dataSource);
        this.total = Number(data.total);
        this.filTotal = Number(data.total_fil);
        this.loading = false;
      } catch (e) {
        if (e) {
          this.loading = false;
        }
      }
    }
    */
  },
  mounted() {
    //this.getBlockByMiner();
  },
  watch: {
    /*
    height: {
      immediate: true,
      handler(v) {
        if (v) {
          this.getBlockByHeight();
        }
      }
    },
    option: {
      deep: true,
      immediate: true,
      handler() {
        //this.getBlockByMiner();
      }
    },
    miners() {
      this.option = {
        begindex: "0",
        count: "25",
        miners: this.miners
      };
    }
    */

    height: {
      immediate: true,
      handler(v) {
        console.log("handler...");
        if (v) {
          this.getBlockByHeight();
        }
      }
    },
  }
};
</script>
<style lang="scss" scoped>
.block-list {
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
}
</style>
