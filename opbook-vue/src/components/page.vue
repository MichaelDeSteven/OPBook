<template>
  <el-pagination
    background
    layout="slot, prev, pager, next"
    :current-page.sync="currentPage"
    :page-size="pageSize"
    prev-text="上一页"
    next-text="下一页"
    :total="count"
    @current-change="handleCurrentChange"
  >
    <span v-html="`共 ${count} 条 &nbsp;&nbsp;&nbsp; 共 ${pages} 页`" />
  </el-pagination>
</template>

<script>
export default {
  name: "Page",
  data() {
    return {
      currentPage: 1, // 默认当前页显示第一页
    };
  },
  computed: {
    pages() {
      return Math.ceil(this.count / this.pageSize);
    },
  },
  props: {
    count: {
      type: Number,
      default: 0,
    },
    pageSize: {
      type: Number,
      default: 10,
    },
  },
  methods: {
    handleCurrentChange(page) {
      // 点击页码事件，通知父组件
      this.$emit("pageEvent", this.currentPage);
    },
  },
};
</script>
