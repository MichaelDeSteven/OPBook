<template>
  <div>
    <div class="manual-reader">
      <Header></Header>
      <div class="container-fluid container-widescreen manual-body">
        <div class="row">
          <SettingMenu></SettingMenu>
          <div class="page-right">
            <div class="m-box">
              <div class="box-head">
                <strong class="box-title">书籍列表</strong>
                <button
                  type="button"
                  data-toggle="modal"
                  data-target="#addBookDialogModal"
                  class="btn btn-success btn-sm pull-right"
                >添加书籍</button>
              </div>
              <div class="help-block text-left" style="margin-top: 15px;">
                <span class="text-danger">凡是违反国家法律法规或恶意发布的内容，本站有权在不提前告知的情况下对内容进行删除，请须知！</span>
              </div>
              <!-- <ul class="nav nav-tabs" style="margin-top: 15px;">
                <li :class="{'active':pub == 0}" @click="checkout(0)">
                  <a href>私有书籍</a>
                </li>
                <li :class="{'active':pub != 0}" @click="checkout(1)">
                  <a href>公开书籍</a>
                </li>
              </ul>-->
            </div>
            <div class="box-body" id="bookList">
              <div class="book-list">
                <template v-if="page.list.length < 1">
                  <div class="text-center">暂无数据</div>
                </template>
                <template v-else>
                  <form
                    target="notarget"
                    style="display: none;"
                    action
                    enctype="multipart/form-data"
                    method="post"
                    id="uploadZip"
                  >
                    <input type="file" name="zipfile" accept="application/zip" />
                    <input type="text" name="identify" value />
                  </form>
                  <form
                    target="notarget"
                    style="display: none;"
                    action
                    enctype="multipart/form-data"
                    method="post"
                    id="uploadEpub"
                  >
                    <input type="file" name="zipfile" accept=".epub" />
                    <input type="text" name="identify" value />
                  </form>
                  <div class="list-item clearfix" v-for="item in page.list">
                    <div class="col-sm-2 col-xs-4" style="padding-left: 0px">
                      <a class="recommend-book" title="查看文档" data-toggle="tooltip">
                        <img
                          onerror="this.src='/static/images/book.png'"
                          class="img-responsive border-cover-img"
                          src
                          alt
                        />
                      </a>
                    </div>
                    <div class="col-sm-10 col-xs-8" style="padding-right: 0px;padding-left: 0px;">
                      <div class="book-title">
                        <div class="pull-left">
                          <a href title="查看文档" data-toggle="tooltip">
                            <i class="fa fa-unlock" aria-hidden="true"></i>
                            {{ item.name }}
                          </a>
                        </div>

                        <div class="clearfix"></div>
                      </div>
                      <div class="info hidden-xs">
                        <span title="创建者" data-toggle="tooltip" data-placement="bottom">
                          <i class="fa fa-user"></i>
                          {{ item.user_id }}
                        </span>
                        <span title="文档数量" data-toggle="tooltip" data-placement="bottom">
                          <i class="fa fa-pie-chart"></i>
                          {{ item.doc_count }}
                        </span>
                        <!-- <span title="书籍角色" data-toggle="tooltip" data-placement="bottom">
                          <i class="fa fa-user-secret"></i> role_name
                        </span>-->
                        <span title="创建时间" data-toggle="tooltip" data-placement="bottom">
                          <i class="fa fa-clock-o"></i>
                          {{ new Date(item.create_time).format("yyyy-MM-dd hh:mm:ss") }}
                        </span>
                        <span title="最后编辑" data-toggle="tooltip" data-placement="bottom">
                          <i class="fa fa-pencil"></i>
                          {{ new Date(item.modify_time).format("yyyy-MM-dd hh:mm:ss") }}
                        </span>
                      </div>
                      <div class="desc-text">{{ item.description }}</div>
                      <div class="btns">
                        <a
                          href
                          title="查看文档"
                          class="btn btn-default btn-sm"
                          data-toggle="tooltip"
                          target="_blank"
                        >
                          <i class="fa fa-eye"></i> 查看
                          <span class="hidden-xs">文档</span>
                        </a>
                        <a
                          href
                          title="书籍设置"
                          class="btn btn-default btn-sm"
                          data-toggle="tooltip"
                          target="_blank"
                        >
                          <i class="fa fa-cogs"></i>
                          <span class="hidden-xs">书籍</span>设置
                        </a>
                        <a href class="btn btn-default btn-sm ajax-get confirm">
                          <i class="fa fa-book"></i> 生成下载文档
                        </a>
                        <a href title="编辑文档" data-toggle="tooltip" class="btn btn-default btn-sm">
                          <i class="fa fa-edit" aria-hidden="true"></i> 编辑
                          <span class="hidden-xs">文档</span>
                        </a>
                        <!-- Split button -->
                        <div class="btn-group">
                          <button type="button" class="btn btn-default">
                            <i class="fa fa-cloud-upload"></i> 导入书籍
                          </button>
                          <button
                            type="button"
                            class="btn btn-default dropdown-toggle"
                            data-toggle="dropdown"
                            aria-haspopup="true"
                            aria-expanded="false"
                          >
                            <span class="caret"></span>
                            <span class="sr-only">Toggle Dropdown</span>
                          </button>
                          <ul class="dropdown-menu">
                            <!-- <li>
                              <a
                                href="javascript:void(0);"
                                class="btn-upload-epub"
                                data-toggle="tooltip"
                                title="支持epub格式电子书导入。"
                              >
                                <i class="fa fa-cloud-upload"></i> EPUB 上传导入
                              </a>
                            </li>-->
                            <li>
                              <a
                                href="javascript:void(0);"
                                class="btn-upload-zip"
                                data-toggle="tooltip"
                                title="支持任意zip压缩的markdown书籍导入。"
                                @click="upload(item.identify)"
                              >
                                <i class="fa fa-cloud-upload"></i> ZIP 上传导入
                              </a>
                            </li>
                            <!-- <li>
                              <a
                                href="javascript:void(0);"
                                class="btn-pull-by-zip"
                                data-toggle="tooltip"
                                title="从任意源拉取zip压缩的markdown书籍"
                              >
                                <i class="fa fa-link"></i> ZIP 拉取导入
                              </a>
                            </li>-->
                            <!-- <li>
                              <a
                                href="javascript:void(0);"
                                class="btn-pull-by-git"
                                data-toggle="tooltip"
                                title="从Git仓库导入markdown书籍"
                              >
                                <i class="fa fa-git"></i> Git Clone 导入
                              </a>
                            </li>-->
                          </ul>
                        </div>
                      </div>
                    </div>
                  </div>
                </template>
              </div>
              <template v-if="page.list.length >= 0">
                <nav class="pagination-container">
                  <ul class="pagination">
                    <li v-if="page.pageCurrent > 1">
                      <a href="javascript:void(0);" @click="showPage(1)">..1</a>
                    </li>
                    <li v-if="page.pageCurrent > 1">
                      <a href="javascript:void(0);" @click="showPage(page.pageCurrent - 1)">«</a>
                    </li>
                    <li class="active" v-if="page.pageCurrent >= 1">
                      <a href="javascript:void(0);">{{ page.pageCurrent }}</a>
                    </li>
                    <li v-if="page.pageCurrent + 1 <= page.totalPage">
                      <a
                        href="javascript:void(0);"
                        @click="showPage(page.pageCurrent + 1)"
                      >{{ page.pageCurrent + 1 }}</a>
                    </li>
                    <li v-if="page.pageCurrent + 2 <= page.totalPage">
                      <a
                        href="javascript:void(0);"
                        @click="showPage(page.pageCurrent + 2)"
                      >{{ page.pageCurrent + 2 }}</a>
                    </li>
                    <li v-if="page.pageCurrent + 3 <= page.totalPage">
                      <a
                        href="javascript:void(0);"
                        @click="showPage(page.pageCurrent + 3)"
                      >{{ page.pageCurrent + 3 }}</a>
                    </li>
                    <li v-if="page.pageCurrent + 4 <= page.totalPage">
                      <a
                        href="javascript:void(0);"
                        @click="showPage(page.pageCurrent + 4)"
                      >{{ page.pageCurrent + 4 }}</a>
                    </li>
                    <li v-if="page.pageCurrent < page.totalPage">
                      <a href="javascript:void(0);" @click="showPage(page.pageCurrent + 1)">»</a>
                    </li>
                    <li v-if="page.pageCurrent <= page.totalPage-5">
                      <a
                        href="javascript:void(0);"
                        @click="showPage(page.totalPage)"
                      >..{{ page.totalPage }}</a>
                    </li>
                  </ul>
                </nav>
              </template>
            </div>
          </div>
        </div>
      </div>
    </div>
    <!-- Modal -->
    <div
      class="modal fade"
      id="addBookDialogModal"
      tabindex="-1"
      role="dialog"
      aria-labelledby="addBookDialogModalLabel"
    >
      <div class="modal-dialog" role="document" style="width: 655px">
        <form method="post" autocomplete="off" action id="addBookDialogForm">
          <div class="modal-content">
            <div class="modal-header">
              <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                <span aria-hidden="true">&times;</span>
              </button>
              <h4 class="modal-title" id="myModalLabel">添加书籍</h4>
            </div>
            <div class="modal-body">
              <div class="form-group">
                <input
                  type="text"
                  class="form-control"
                  placeholder="标题(不超过100字)"
                  name="book_name"
                  id="bookName"
                  v-model="book.name"
                />
              </div>
              <div class="form-group">
                <div class="input-group">
                  <span class="input-group-btn">
                    <button class="btn btn-default" type="button">来源名称</button>
                  </span>
                  <input
                    type="text"
                    placeholder="选填"
                    name="author"
                    id="author"
                    class="form-control"
                    v-model="book.author"
                  />
                  <span class="input-group-btn">
                    <button
                      class="btn btn-default"
                      style="border-left: 0px;border-right: 0px;border-radius: 0px;"
                      type="button"
                    >来源链接</button>
                  </span>
                  <input
                    type="text"
                    placeholder="选填"
                    name="author_url"
                    id="author_url"
                    class="form-control"
                    v-model="book.author_url"
                  />
                </div>
              </div>
              <div class="form-group">
                <!-- <div class="pull-left" style="padding: 7px 5px 6px 0">BaseUrl</div> -->
                <input
                  type="text"
                  class="form-control pull-left"
                  style="width: 220px;vertical-align: middle"
                  placeholder="书籍唯一标识(不能超过50字)"
                  name="identify"
                  id="identify"
                  v-model="book.identify"
                />
                <div class="clearfix"></div>
                <p
                  class="text"
                  style="font-size: 12px;color: #999;margin-top: 6px;"
                >书籍标识只能包含字母、数字，以及“-”、"."和“_”符号，且不能是纯数字</p>
              </div>
              <div class="form-group">
                <textarea
                  name="description"
                  id="description"
                  class="form-control"
                  placeholder="描述信息不超过500个字符"
                  style="height: 90px;"
                  v-model="book.description"
                ></textarea>
              </div>
              <div class="form-group">
                <div class="col-lg-6">
                  <label>
                    <input type="radio" name="privately_owned" value="1" v-model="book.private" /> 私有
                    <span class="text">(只有参与者或使用令牌才能访问)</span>
                  </label>
                </div>
                <div class="col-lg-6">
                  <label>
                    <input
                      type="radio"
                      name="privately_owned"
                      value="0"
                      checked
                      v-model="book.private"
                    /> 公开
                    <span class="text">(任何人都可以访问)</span>
                  </label>
                </div>
                <div class="clearfix"></div>
              </div>
              <div class="clearfix"></div>
            </div>
            <div class="modal-footer">
              <span id="form-error-message"></span>
              <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
              <button
                type="button"
                class="btn btn-success"
                id="btnSaveDocument"
                data-loading-text="保存中..."
                @click="Create"
              >保存</button>
            </div>
          </div>
        </form>
      </div>
    </div>
    <!--END Modal-->
  </div>
</template>


<script>
import Header from "@/components/header.vue";
import SettingMenu from "@/components/user/menu.vue";
import service from "@/utils/request";
export default {
  name: "BookIndex",
  data() {
    return {
      book: {
        name: "",
        identify: "",
        author: "",
        author_url: "",
        description: "",
        private: 0,
      },
      pageReq: {
        pageIndex: 1,
        pageSize: 10,
      },
      page: {
        totalCount: 0,
        totalPage: 0,
        pageCurrent: 1,
        pageSize: 2,
        list: [],
      },
    };
  },
  created() {
    this.Index(1, 2);
  },
  components: {
    Header,
    SettingMenu,
  },
  methods: {
    upload(identify) {
      var form = $("form#uploadZip");
      form.find("input[type=file]").trigger("click");
      $("#uploadZip input[type=file]").change(function () {
        var _this = $(this);
        if (_this.val() && confirm("您确定要上传 " + _this.val() + " 吗？")) {
          var formData = new FormData();
          formData.append("zipfile", _this.get(0).files[0]);
          formData.append("identify", identify);
          service({
            url: "/book/upload",
            method: "post",
            data: formData,
            headers: {
              "Content-Type": "multipart/form-data",
            },
          }).then((res) => {
            console.log(res);
            if (res.data.code === 0) {
              alertTips("success", res.message, 3000, "");
              setTimeout(() => {
                location.reload();
              }, 3000);
            } else {
              alertTips("error", res.message, 3000, "");
            }
          });
        }
      });
    },
    showPage(nextPage) {
      if (nextPage > this.page.totolPage) {
        return;
      }
      this.Index(nextPage, 2);
    },
    Index(pageIndex, pageSize) {
      this.pageReq.pageSize = pageSize;
      this.pageReq.pageIndex = pageIndex;
      service({
        url: "/book/index",
        method: "post",
        data: this.pageReq,
      }).then((res) => {
        this.page.list = res.data.data.books;
        this.page.totalCount = res.data.data.totalCount;
        this.page.totalPage = Math.ceil(
          this.page.totalCount / this.page.pageSize
        );
        this.page.pageCurrent = pageIndex;
        console.log(this.page.totalPage);
      });
    },
    Create() {
      console.log(this.book);
      if (this.book.name === "") {
        layer.msg("书籍标题不能为空");
        return;
      }
      if (this.book.name.length > 100) {
        layer.msg("书籍标题必须小于100字符");
        return;
      }

      if (this.book.identify === "") {
        layer.msg("书籍标识不能为空");
        return;
      }
      if (this.book.identify.length > 50) {
        layer.msg("书籍标识必须小于50字符");
        return;
      }

      if (this.book.description.length > 500) {
        layer.msg("描述信息不超过500个字符");
        return;
      }
      service({
        url: "/book/create",
        method: "post",
        data: this.book,
      }).then((res) => {
        console.log(res);
        layer.msg(res.data.msg);
        if (res.data.code === 0) {
          window.location = "/book";
        }
      });
    },
  },
};
</script>