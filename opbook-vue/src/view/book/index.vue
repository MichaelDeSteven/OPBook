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
              <ul class="nav nav-tabs" style="margin-top: 15px;">
                <li class>
                  <a href>私有书籍</a>
                </li>
                <li class>
                  <a href>公开书籍</a>
                </li>
              </ul>
            </div>
            <div class="box-body" id="bookList">
              <div class="book-list">
                <div class="text-center">暂无数据</div>
              </div>
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
    };
  },
  components: {
    Header,
    SettingMenu,
  },
  methods: {
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