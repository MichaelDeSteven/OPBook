<template>
  <div>
    <div class="ucenter-content">
      <ul>
        <li class="clearfix" v-for="item in page.list">
          <div class="col-sm-2 col-xs-3">
            <a :href="'/introduct/' + item.identify" target="_blank" title>
              <img
                onerror="this.src='/static/images/book.png'"
                :src="item.cover"
                alt
                class="img-responsive border-cover-img"
              />
            </a>
          </div>
          <div class="col-sm-10 col-xs-9">
            <a :href="'/introduct/' + item.identify" target="_blank">
              <h4>{{ item.name }}</h4>
            </a>
            <div class="text-muted book-info hidden-xs">
              <span title="文档数量" class="tooltips">
                <i class="fa fa-pie-chart">{{ item.doc_count }}</i>
              </span>
              <span title="阅读人次" class="tooltips">
                <i class="fa fa-eye">{{ item.view_count }}</i>
              </span>
              <span title="收藏人次" class="tooltips">
                <i class="fa fa-heart-o">{{ item.collect_count }}</i>
              </span>
            </div>
            <div class="help-block book-description">{{ item.description }}</div>
          </div>
        </li>
      </ul>
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
  name: "Collection",
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
    this.GetCollection();
  },
  methods: {
    GetCollection() {
      service({
        url: "/book/user/collect",
        method: "post",
        data: this.page,
      }).then((res) => {
        console.log(res);
        if (res.data.code === 0) {
          this.page.list = res.data.data.books;
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
      this.GetCollection();
    },
  },
};
</script>
