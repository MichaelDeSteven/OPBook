<template>
  <div class="manual-reader manual-container manual-search-reader search-result">
    <Header></Header>
    <div class="container-fluid container-widescreen manual-body">
      <div class="row">
        <div class="col-sm-8 col-xs-12">
          <form method="get" action class="search-form">
            <div class="input-group input-group-lg">
              <input
                type="text"
                name="wd"
                value
                placeholder="Search..."
                class="form-control"
                v-model="req.wd"
              />
              <input type="hidden" name="tab" v-model="req.wd" />
              <span class="input-group-addon" @click="getResult">
                <i class="fa fa-search"></i>
                <span class="hidden-xs">搜索</span>
              </span>
            </div>
          </form>
        </div>
      </div>
      <div class="row">
        <div class="col-sm-8 col-xs-12 result-list">
          <div class="help-block">
            <a style="color: red;" target="_blank" rel="nofollow noreferrer" href>OPBOOK</a> 本次搜索耗时
            <span class="text-danger">{{res.spendTime}}</span> 秒，为您找到
            <span class="text-success">{{res.totalRows}}</span> 个相关结果.
          </div>
          <ul class="nav nav-tabs">
            <li :class="{active:tab==0}">
              <a class="active" @click="check(0)">搜书籍</a>
            </li>
            <li :class="{active:tab==1}">
              <a @click="check(1)">搜文档</a>
            </li>
          </ul>
          <ul :class="{'doc-result':tab==1}">
            <div v-show="tab == 0">
              <template v-if="res.books != null && res.books.length < 1">
                <li class="clearfix">
                  <div class="help-block">
                    啊哦，没搜到相关书籍，
                    <a style="font-weight: 600;color: #06f;" @click="check(1)">搜文档</a> 试下？
                  </div>
                </li>
              </template>
              <template v-else>
                <li class="clearfix" v-for="book in res.books">
                  <div class="col-sm-3 col-md-3 col-lg-2 col-xs-3" style="padding: 0px;">
                    <a :href="'/introduct/' + book.identify" target="_blank">
                      <img
                        onerror="this.src='/static/images/book.png'"
                        src
                        class="img-responsive border-cover-img"
                      />
                    </a>
                  </div>
                  <div class="col-sm-9 col-md-9 col-lg-10 col-xs-9">
                    <a :href="'/introduct/' + book.identify" target="_blank">
                      <h4>{{book.name}}</h4>
                    </a>
                    <div class="text-muted book-info hidden-xs">
                      <span title class="tooltips" data-original-title="文档数量">
                        <i class="fa fa-pie-chart"></i>
                        {{book.doc_count}}
                      </span>
                      <span title class="tooltips" data-original-title="阅读人次">
                        <i class="fa fa-eye"></i>
                        {{book.view_count}}
                      </span>
                      <span title class="tooltips" data-original-title="收藏人次">
                        <i class="fa fa-heart-o"></i>
                        {{book.score_count}}
                      </span>
                      <span title class="tooltips" data-original-title="创建时间">
                        <i class="fa fa-clock-o"></i>
                        {{ new Date(book.create_time).format("yyyy-MM-dd hh:mm:ss") }}
                      </span>
                    </div>
                    <div class="help-block book-description">{{book.description}}</div>
                  </div>
                </li>
              </template>
            </div>
            <div v-show="tab == 1">
              <template v-if="res.docs != null && res.docs.length < 1">
                <li class="clearfix">
                  <div class="help-block">很遗憾，没搜到相关文档，建议您更换关键字重新搜索</div>
                </li>
              </template>
              <template v-else>
                <li class="clearfix" v-for="doc in res.docs">
                  <div class="col-xs-12">
                    <a href target="_blank">
                      <h4>{{doc.doc_name}}</h4>
                    </a>
                    <div class="text-muted book-info">
                      <span class="tooltips" title="阅读人次">
                        <i class="fa fa-eye"></i>
                        {{doc.view_count}}
                      </span>
                      <span class="tooltips" title="发布时间">
                        <i class="fa fa-clock-o"></i>
                        {{ new Date(doc.create_time).format("yyyy-MM-dd hh:mm:ss") }}
                      </span>
                      <span>
                        <a href target="_blank" class="tooltips" title="查看书籍">
                          <i class="fa fa-book text-muted"></i>
                          《{{doc.book_name}}》
                        </a>
                      </span>
                    </div>
                    <div class="help-block book-description">{{doc.release.replace(/<[^>]+>/g,"")}}</div>
                  </div>
                </li>
              </template>
            </div>
          </ul>
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
  name: "SearchResult",
  data() {
    return {
      req: {
        wd: "",
        isSearchDoc: false,
        pageSize: 10,
        pageIndex: 1,
      },
      res: {
        wd: "",
        isSearchDoc: 0,
        books: [],
        docs: [],
        spendTime: "",
        totalRows: 0,
      },
      books: [],
      docs: [],
      tab: 0,
    };
  },
  created() {
    this.getParams();
  },
  watch: {
    $route: "getParams",
  },
  components: {
    Header,
    Footer,
  },
  methods: {
    getParams() {
      // 取到路由带过来的参数
      const wd = this.$route.query.wd; // 将数据放在当前组件的数据内
      this.req.wd = wd;
      this.getResult();
    },
    getResult() {
      if (this.req.wd === "") {
        return;
      }
      service({
        url: "/search",
        method: "post",
        data: this.req,
      }).then((res) => {
        console.log(res);
        if (res.data.code === 0) {
          this.res = res.data.data;
        } else {
        }
      });
    },
    check(inx) {
      if (this.tab === inx) return;
      this.tab = inx;
      this.req.isSearchDoc = this.tab == 1 ? true : false;
      this.getResult();
    },
  },
};
</script>

<style scoped>
.search-highlight {
  font-style: normal;
  color: red;
  font-weight: 400;
}
</style>
