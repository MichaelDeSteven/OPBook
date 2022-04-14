<template>
  <div>
    <Header></Header>
    <div class="container-fluid container-widescreen">
      <div class="row bookstack-info">
        <div class="col-xs-12">
          <h1>{{ book.name }}</h1>
          <div class="help-block subbookname">
            <span>语言：</span>
            <span
              style="color: red"
            >{{ book.lang === "ch" ? "中文" : (book.lang === "en" ? "英文" : "其它")}}</span>
            <span class="mgl-10px">评分：</span>
            <i class="bookstack-star"></i>
            {{ book.score*1.0 / 10.0 }}
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
              :src="book.cover"
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
              <a :href="book.author_url" target="_blank" title>{{book.author}}</a>
            </li>
            <li class="bookstack-labels" v-if="labels.length > 0 && labels[0].length > 0">
              <a target="_blank" :href="'/search/result?wd=' + item" v-for="item in labels">{{item}}</a>
            </li>
            <li class="bookstack-description hidden-xs">{{ book.description }}</li>
            <li class="book-metadata hidden-xs">
              {{ book.doc_count }}
              <small>章节</small>
              {{ book.view_count }}
              <small>阅读</small>
              {{ book.collect_count }}
              <small>收藏</small>
            </li>
            <li class="hidden-xs">
              <a href target="_blank" title="阅读" class="btn btn-info btn-lg">
                <i class="fa fa-book"></i> 阅读
              </a>
              <a rel="nofollow" class="btn btn-warning btn-lg ajax-star" @click="Star">
                <template v-if="star.is_deleted === 1">
                  <i class="fa fa-heart-o"></i>
                  <span>加入收藏</span>
                </template>
                <template v-else>
                  <i class="fa fa-heart"></i>
                  <span>已收藏</span>
                </template>
              </a>
            </li>
          </ul>
        </div>
      </div>
      <div class="row">
        <div class="col-xs-12 bookstack-menu">
          <ul class="nav nav-tabs">
            <li :class="{'active':select===0}">
              <a href="javascript:;" @click="select=0">书籍目录</a>
            </li>
            <li :class="{'active':select===1}">
              <a href="javascript:;" @click="DisplayComment">书籍评论</a>
            </li>
          </ul>
          <div class="help-block">
            <ul class="none-listyle">
              <div v-if="select===0">
                <li v-for="item in menuTop" :key="item.id">
                  <a
                    :href="'/read/' + book.identify + '/' + item.identify"
                    target="_blank"
                    :title="item.name"
                  >{{ item.name }}</a>
                </li>
              </div>
              <div v-if="select===1">
                <li class="comments-form clearfix">
                  <div class="score">
                    <template v-if="score.score === 0">
                      <span title="点击即可给当前书籍打分" class="cursor-pointer" data-toggle="tooltip">
                        <i class="fa fa-star-o" data-score="1" data-tips="很差" @click="AddScore(10)"></i>
                        <i class="fa fa-star-o" data-score="2" data-tips="较差" @click="AddScore(20)"></i>
                        <i class="fa fa-star-o" data-score="3" data-tips="还行" @click="AddScore(30)"></i>
                        <i class="fa fa-star-o" data-score="4" data-tips="推荐" @click="AddScore(40)"></i>
                        <i class="fa fa-star-o" data-score="5" data-tips="力荐" @click="AddScore(50)"></i>
                      </span>
                      <span class="text-muted"></span>
                    </template>
                    <template v-else>
                      <span class="text-muted">
                        我的评分:
                        <i
                          :class="['bookstack-star', {'star-10' : myScore[0]}, {'star-20' : myScore[1]}, {'star-30' : myScore[2]}, {'star-40' : myScore[3]}, {'star-50' : myScore[4]}]"
                        ></i>
                      </span>
                    </template>
                  </div>
                </li>
                <form action method="post" class="ajax-form">
                  <div class="form-group">
                    <textarea
                      class="form-control"
                      name="content"
                      rows="3"
                      placeholder="文明评论，理性发言"
                      v-model="comment.content"
                    ></textarea>
                  </div>
                  <div class="form-group">
                    <span class="pull-left text-muted"></span>
                    <input
                      type="button"
                      class="btn btn-success pull-right"
                      value="发表点评"
                      @click="AddComment(0)"
                    />
                  </div>
                </form>
                <li class="comments-list">
                  <ul>
                    <li class="clearfix" v-for="item in commentList">
                      <div class="col-xs-12">
                        <img :src="item.avatar" class="img-thumbnail img-circle img-responsive" alt />
                        <span class="username">{{item.nickname}}</span>
                        <span class="text-muted">
                          <i></i>
                          {{ new Date(item.comment_time).format("yyyy-MM-dd hh:mm:ss") }}
                        </span>
                        <span class="text-bar">
                          <i
                            class
                            :class="['like-ico', {'zan-hover' : item.is_like == true}, {'zan' : item.is_like == false}]"
                            @click="Like(item)"
                          ></i>
                          <span class="text-offset">{{item.like_count}}</span>
                        </span>
                        <span class="repley" @click="reply.id=item.id">
                          <i class="fa fa-comments-o"></i>
                          回复
                        </span>
                      </div>
                      <div v-if="item.comment_id !== 0" class="col-xs-12 comments-content">
                        <div class="reply-to">
                          <span class="text-info">{{item.reply_nickname}}</span>
                          : {{item.reply_content}}
                        </div>
                      </div>
                      <div class="col-xs-12 comments-content">{{item.content}}</div>
                      <div v-if="reply.id===item.id">
                        <form action method="post" class="ajax-form">
                          <div class="form-group">
                            <textarea
                              class="form-control"
                              name="content"
                              rows="3"
                              placeholder="文明评论，理性发言"
                              v-model="reply.content"
                            ></textarea>
                          </div>
                          <div class="form-group">
                            <span class="pull-left text-muted"></span>
                            <input
                              type="button"
                              class="btn btn-success pull-right"
                              value="发表点评"
                              @click="AddComment(item.id)"
                            />
                          </div>
                        </form>
                      </div>
                    </li>
                  </ul>
                </li>
              </div>
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
        score: 0,
        view_count: 0,
        collect_count: 0,
        generate_time: "",
      },
      menuTop: [],
      star: {
        book_id: 0,
        user_id: 0,
        is_deleted: 0,
      },
      select: 0,
      comment: {
        book_id: 0,
        user_id: 0,
        content: "",
        comment_id: 0,
      },
      commentList: [],
      reply: {
        id: 0,
        content: "",
      },
      score: {
        user_id: 0,
        book_id: 0,
        score: 0,
      },
      myScore: [false, false, false, false, false],
      labels: [],
    };
  },
  beforeCreate() {
    document.querySelector("body").setAttribute("id", "bookstack-intro");
  },
  created() {
    var u = JSON.parse(localStorage.getItem("user"));
    if (u != null) {
      this.comment.user_id = parseInt(u.id);
      this.score.user_id = parseInt(u.id);
    }
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
          if (this.book.author === "") {
            this.book.author = "暂无来源";
          }
          this.GetMenuTop();
          this.IsStar();
          this.score.book_id = this.book.id;
          var arr = document.getElementsByClassName("bookstack-star");
          arr[0].classList.add("star-" + this.book.score);
          this.GetScore();
          this.labels = this.book.label.split(",");
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
    Star() {
      service({
        url: "/book/collect/" + this.book.id,
        method: "post",
        data: this.starReq,
      }).then((res) => {
        console.log(res);
        // layer.msg(res.data.data);
        if (res.data.code === 0) {
          if (this.star.is_deleted === 1) {
            this.star.is_deleted = 0;
          } else {
            this.star.is_deleted = 1;
          }
        }
      });
    },
    IsStar() {
      service({
        url: "/book/collect/stat/" + this.book.id,
        method: "get",
      }).then((res) => {
        console.log(res);
        this.star.is_deleted = res.data.data.is_deleted;
      });
    },
    AddComment(CommentId) {
      this.comment.comment_id = CommentId;
      this.comment.book_id = this.book.id;
      if (CommentId !== 0) {
        this.comment.content = this.reply.content;
      }
      if (this.comment.user_id === 0) {
        layer.msg("用户未登录");
      }
      service({
        url: "/social/comment/add",
        method: "post",
        data: this.comment,
      }).then((res) => {
        location.reload();
        console.log(res);
      });
    },
    DisplayComment() {
      this.select = 1;
      service({
        url: "/social/comment/get/" + this.book.id,
        method: "get",
      }).then((res) => {
        console.log(res);
        if (res.data.code === 0) {
          this.commentList = res.data.data;
        } else {
        }
      });
    },
    AddScore(score) {
      this.score.score = score;
      this.myScore[score / 10 - 1] = true;
      service({
        url: "/book/score/add",
        method: "post",
        data: this.score,
      }).then((res) => {
        console.log(res);
        if (res.data.code === 0) {
        } else {
        }
      });
    },
    GetScore() {
      service({
        url: "/book/score/get",
        method: "post",
        data: this.score,
      }).then((res) => {
        console.log(res);
        if (res.data.code === 0) {
          this.score = res.data.data;
          this.myScore[this.score.score / 10 - 1] = true;
        } else {
        }
      });
    },
    Like(item) {
      service({
        url: "/social/like/" + item.id,
        method: "post",
      }).then((res) => {
        if (res.data.code === 0) {
          item.is_like = res.data.data.is_like;
          if (res.data.data.is_like === true) {
            item.like_count = item.like_count + 1;
          } else {
            item.like_count = item.like_count - 1;
          }
        }
        console.log(res);
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

.like-ico {
  width: 20px;
  height: 20px;
  display: inline-block;
  background: center/contain no-repeat;
  vertical-align: text-top;
  /* margin-right: 4px; */
  margin-top: -2px;
}
.text-offset {
  margin: auto;
}
.text-bar {
  width: 50.7px;
  height: 13.6px;
}
.zan {
  background-size: 15px;
  background-image: url("https://s1.hdslb.com/bfs/static/blive/blfe-dynamic-web/static/img/zan.0da89524.svg");
}
.zan-hover {
  background-size: 15px;
  background-image: url("https://s1.hdslb.com/bfs/static/blive/blfe-dynamic-web/static/img/zan-hover.ab577109.svg");
}
</style>
