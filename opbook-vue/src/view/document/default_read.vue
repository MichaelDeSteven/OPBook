<template>
  <div>
    <div class="m-manual manual-mode-view manual-reader">
      <header class="navbar navbar-static-top manual-head" role="banner">
        <div class="container-fluid">
          <div class="navbar-header pull-left manual-title">
            <span class="slidebar" id="slidebar">
              <i class="fa fa-align-justify"></i>
            </span>
            <a href title class="book-title">{{ book.name }}</a>
            <span style="font-size: 12px;font-weight: 100;"></span>
          </div>
          <div class="navbar-header pull-right manual-menu">
            <div class="bookstack-item">
              <a title="首页" href class="btn btn-link">
                <i class="fa fa-home"></i> 首页
              </a>
            </div>
          </div>
        </div>
      </header>
      <div class="container-fluid manual-body">
        <div class="manual-right">
          <div class="manual-article">
            <div class="article-head">
              <div class="container-fluid">
                <div class="row">
                  <div class="col-md-10 col-md-offset-1 col-xs-12 text-center">
                    <h1 id="article-title">{{ doc.name }}</h1>
                  </div>
                </div>
              </div>
            </div>
            <div class="article-content">
              <div class="help-block article-info" style="padding-bottom: 20px;">
                <!-- <span class="alink tooltips text-red" title="AdTitle">
                  <i class="fa fa-angle-right"></i>
                  <a target="_blank" rel="nofollow" href>AdTitle</a>
                  <i class="fa fa-angle-left"></i>
                </span>-->
                <span>
                  <i class="fa fa-user-o"></i>
                  <a target="_blank" title="来源" :href="book.author_url">{{ book.author }}</a>
                </span>
                <span>
                  <i class="fa fa-eye"></i> 浏览
                  <i class="read-count">{{ doc.view_count }}</i>
                </span>
                <span>
                  <a href="#" data-target="#ModalShare" data-toggle="modal">
                    <i class="fa fa-share-alt"></i> 分享
                  </a>
                </span>
                <span class="pull-right hidden-xs tooltips" title="更新时间">
                  <i class="fa fa-clock-o"></i>
                  <i
                    class="updated-at"
                  >{{ new Date(doc.modify_time).format("yyyy-MM-dd hh:mm:ss") }}</i>
                </span>
              </div>
              <article
                class="article-body markdown-body editormd-preview-container"
                id="page-content"
                v-html="doc.release"
              ></article>
              <div class="row hung-read-link">
                <div class="col-xs-12 hung-pre" style="display: none">
                  <span class="text-muted">上一篇:</span>
                  <a href="#"></a>
                </div>
                <div class="col-xs-12 hung-next" style="display: none">
                  <span class="text-muted">下一篇:</span>
                  <a href="#"></a>
                </div>
              </div>

              <div class="bookstack-bars">
                <ul>
                  <li class="visible-xs visible-sm bars-menu bars-menu-hide">
                    <a href="/" title="首页">
                      <i class="fa fa-home"></i>
                    </a>
                  </li>
                </ul>
              </div>
            </div>
          </div>
        </div>

        <div class="manual-left">
          <div class="article-search" data-bookid>
            <form id="searchForm" action method="post">
              <div class="input-group">
                <input
                  type="text"
                  name="keyword"
                  placeholder="Search..."
                  autocomplete="off"
                  class="form-control"
                />
                <span class="input-group-addon input-group-addon-clear">
                  <i class="fa fa-remove"></i>
                </span>
                <span class="input-group-addon">
                  <button type="submit">
                    <i class="fa fa-search"></i>
                  </button>
                </span>
              </div>
            </form>
            <div class="pull-right hidden-xs">
              <i class="fa fa-align-justify"></i>
            </div>
          </div>
          <div class="article-menu">
            <div class="article-menu-detail collapse-menu" v-html="menu"></div>
            <div class="search-result">
              <div class="search-empty" style="display: block;">
                <i class="fa fa-search-plus" aria-hidden="true"></i>
                <b class="text">暂无相关搜索结果！</b>
              </div>
              <ul class="search-list" id="searchList"></ul>
            </div>
          </div>
        </div>
        <div class="manual-progress">
          <b class="progress-bar"></b>
        </div>
      </div>
      <div class="manual-mask"></div>
    </div>
    <!-- <span class="article-toggle tooltips hidden-xs hidden-sm" title="展开/收起文章目录" @click="Toggle">
      <img alt="展开/收起文章目录" src="https://static.sitestack.cn/static/images/toggle.png" />
    </span>-->
    <div id="menu-hidden" style="display: none;"></div>
  </div>
</template>

<script>
import service from "@/utils/request";
export default {
  name: "READ",
  components: {},
  data() {
    return {
      doc: {
        id: "",
        name: "",
        identify: "",
        book_id: "",
        parent_id: "",
        order_sort: "",
        comment_count: 0,
        view_count: 0,
        modify_time: "",
        release: "",
        markdown: "",
      },
      menu: "",
      book: {
        name: "",
        author: "",
        author_url: "",
      },
    };
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
  methods: {
    Get() {
      service({
        url:
          "/document/read/" +
          this.$route.params.book_identify +
          "/" +
          this.$route.params.doc_identify,
        method: "get",
      }).then((res) => {
        console.log(
          "/document/read/" +
            this.$route.params.book_identify +
            "/" +
            this.$route.params.doc_identify
        );
        if (res.data.code === 0) {
          this.doc = res.data.data.doc;
          this.menu = res.data.data.menu;
          this.book = res.data.data.book;
          if (this.book.author === "") {
            this.book.author = "暂无来源";
          }
          // $(function () {
          //   editormd.markdownToHTML("page-content", {
          //     markdown: res.data.data.markdown, // Also, you can dynamic set Markdown text
          //   });
          // });
          // console.log($("#page-content").html());
        } else {
        }
      });
    },
    Toggle() {
      if ($("body").hasClass("article-menu-hide")) {
        $("body").removeClass("article-menu-hide");
      } else {
        $("body").addClass("article-menu-hide");
      }
    },
  },
};
</script>

<style scoped>
@import "https://static.sitestack.cn/static/bootstrap/css/bootstrap.min.css";
@import "https://static.sitestack.cn/static/font-awesome/css/font-awesome.min.css";
@import "https://static.sitestack.cn/static/editor.md/css/editormd.preview.css";
@import "https://static.sitestack.cn/static/css/toast.css";
@import "https://static.sitestack.cn/static/page/page.css";
@import "https://static.sitestack.cn/static/katex/katex.min.css";
@import "https://static.sitestack.cn/static/mind-map/mindmap.css";
@import "https://static.sitestack.cn/static/css/bookstack.css?version=v2.12.0-beta31008";
</style>

<style>
.editormd-preview-container ol.linenums li code,
.editormd-html-preview ol.linenums li code {
  display: block;
  white-space: pre;
}
li.L1,
li.L3,
li.L5,
li.L7,
li.L9 {
  background-color: transparent;
}
.editormd-preview-container pre.prettyprint,
.editormd-html-preview pre.prettyprint {
  border-color: transparent;
}
body {
  -webkit-overflow-scrolling: touch;
}
.alink,
.alink a {
  color: #888 !important;
}
.alink:hover,
.alink a:hover {
  color: red !important;
}
</style>
