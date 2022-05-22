<template>
  <div>
    <div class="ucenter-content">
      <div class="col-xs-4 col-sm-2 col-md-2 text-center fans-item" v-for="item in page.list">
        <a target="_blank" title :href="'/user/' + item.id" class="tooltips">
          <img class="thumbnail img-circle" :src="item.avatar" alt />
        </a>
        <div>
          <a
            target="_blank"
            title
            :href="'/user/' + item.id"
            class="fans-username tooltips"
          >{{item.nickname}}</a>
        </div>
      </div>
    </div>
    <div>
      <template v-if="page != null && page.list.length > 0">
        <my-pagination :pageSize="page.page_size" :count="page.total_count" @pageEvent="showPage" />
      </template>
    </div>
  </div>
</template>

<script>
import myPagination from "@/components/page";
import service from "@/utils/request";
export default {
  name: "Fan",
  components: {
    myPagination,
  },
  data() {
    return {
      page: {
        user_id: 0,
        page_index: 1,
        page_size: 10,
        total_count: 0,
        list: [],
      },
    };
  },
  created() {
    this.page.user_id = parseInt(this.$route.params.uid);
    this.page.page_index = 1;
    this.GetFollowers();
  },
  methods: {
    GetFollowers() {
      service({
        url: "/social/fan/getFollowees",
        method: "post",
        data: this.page,
      }).then((res) => {
        console.log(res);
        if (res.data.code === 0) {
          this.page.list = res.data.data.data;
          this.page.total_count = res.data.data.totalCount;
        } else {
        }
      });
    },
    showPage(nextPage) {
      if (nextPage > this.page.totolPage) {
        return;
      }
      this.page.page_index = nextPage;
      this.GetFollowers();
    },
  },
};
</script>
