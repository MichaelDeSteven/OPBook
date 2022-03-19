<template>
  <div>
    <Header></Header>
    <div class="container-fluid container-widescreen">
      <div class="row bookstack-info">
        <div class="col-xs-12">
          <h1>{{ book.name }}</h1>
          <div class="help-block subbookname">
            <span>语言：</span>
            <span style="color: red">中文</span>
            <span class="mgl-10px">评分：</span>
            <i class="bookstack-star star-45"></i>
            {{ book.score / 10 }}
            <span class="hidden-xs">
              (
              <span class="text-muted">{{ book.score_count }}个有效评分</span>)
            </span>
            <span class="mgl-10px">
              <div class="visible-xs"></div>
            </span>
            <span>
              <i class="fa fa-clock-o"></i>
              最后更新：
              {{ new Date(book.generate_time).format("yyyy-MM-dd hh:mm:ss") }}
            </span>
          </div>
        </div>
        <div class="col-sm-3 col-xs-4 col-md-3 col-lg-2">
          <div class="recommend-bookend mgt-15px">
            <img
              src
              onerror="this.src='/static/images/book.png'"
              class="img-responsive border-cover-img"
              alt=".Book.BookName"
            />
          </div>
        </div>
        <div class="col-sm-9 col-md-7 col-lg-8 col-xs-8">
          <ul class="none-listyle">
            <li>
              <span>来源：</span>
              <a
                href="https://github.com/wangdoc/javascript-tutorial"
                target="_blank"
                title="来源：阮一峰"
              >阮一峰</a>
              <span style="color: #ddd;">&nbsp;|&nbsp;</span>
              <span>整理：</span>
              <a target="_blank" class="text-primary" href="/user/admin">进击的皇虫</a>
            </li>
            <li class="bookstack-labels">
              <a target="_blank" title="javascript" href="/tag/javascript">javascript</a>
              <a target="_blank" title="阮一峰" href="/tag/阮一峰">阮一峰</a>
            </li>
            <li class="bookstack-description hidden-xs">{{ book.description }}</li>
            <li class="book-metadata hidden-xs">
              <small>章节</small>
              {{ book.doc_count }}
              <small>阅读</small>
              {{ book.view_count }}
              <small>收藏</small>
              {{ book.collect_count }}
            </li>
            <li class="hidden-xs">
              <a href target="_blank" title="阅读" class="btn btn-info btn-lg">
                <i class="fa fa-book"></i> 阅读
              </a>
              <a href="/book/star/192" rel="nofollow" class="btn btn-warning btn-lg ajax-star">
                <i class="fa fa-heart-o"></i>
                <span>加入收藏</span>
              </a>
            </li>
          </ul>
        </div>
      </div>
      <div class="row">
        <div class="col-xs-12 bookstack-menu">
          <ul class="nav nav-tabs">
            <li class="active">
              <a href>书籍目录</a>
            </li>
            <li>
              <a href>
                书籍评论 (
                <span class="text-muted">{{ book.comment_count }}</span>)
              </a>
            </li>
          </ul>
          <div class="help-block">
            <ul class="none-listyle">
              <li v-for="item in menuTop" :key="item.id">
                <a
                  :href="'read/' + book.name + '/' + item.identify"
                  target="_blank"
                  :title="item.name"
                >{{ item.name }}</a>
              </li>
            </ul>
          </div>
        </div>
      </div>
    </div>
    <Footer></Footer>
  </div>
</template>

<script>
import Header from "@/components/header.vue";
import Footer from "@/components/footer.vue";
import service from "@/utils/request";
export default {
  name: "Intro",
  data() {
    return {
      book: {
        id: "",
        name: "",
        identify: "",
        author: "",
        author_url: "",
        description: "",
        private: 0,
        doc_count: 0,
        cover: "",
        comment_count: 0,
        score_count: 0,
        score: 0.0,
        view_count: 0,
        collect_count: 0,
        generate_time: "",
      },
      menuTop: [],
    };
  },
  beforeCreate() {
    document.querySelector("body").setAttribute("id", "bookstack-intro");
  },
  created() {
    // watch 路由的参数，以便再次获取数据
    this.$watch(
      () => this.$route.params,
      () => {
        this.Get();
      },
      // 组件创建完后获取数据，
      // 此时 data 已经被 observed 了
      { immediate: true }
    );
  },
  beforeDestroy() {
    document.querySelector("body").removeAttribute("id");
  },
  components: {
    Header,
    Footer,
  },
  methods: {
    Get() {
      service({
        url: "/book/introduct/" + this.$route.params.identify,
        method: "get",
      }).then((res) => {
        console.log(res);
        if (res.data.code === 0) {
          this.book = res.data.data;
          this.GetMenuTop();
        } else {
        }
      });
    },
    GetMenuTop() {
      service({
        url: "/document/index/" + this.$route.params.identify,
        method: "get",
      }).then((res) => {
        console.log(res);
        if (res.data.code === 0) {
          this.menuTop = res.data.data;
        } else {
        }
      });
    },
  },
};
</script>

<style scoped>
@import "https://static.sitestack.cn/static/bootstrap/css/bootstrap.min.css";
@import "https://static.sitestack.cn/static/font-awesome/css/font-awesome.min.css";
@import "https://static.sitestack.cn/static/css/toast.css";
@import "https://static.sitestack.cn/static/css/main.css?version=v2.12.0-beta3";
</style>
