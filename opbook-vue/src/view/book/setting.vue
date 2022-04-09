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
                <strong class="box-title">书籍设置</strong>
              </div>
            </div>
            <div class="box-body">
              <div class="form-right">
                <label>书籍封面</label>
                <label>
                  <a
                    href="javascript:;"
                    data-toggle="modal"
                    title="点击修改封面"
                    data-target="#upload-logo-panel"
                  >
                    <img
                      :src="book.cover"
                      class="border-cover-img"
                      alt="封面"
                      style="margin-left:15px;max-width: 120px;max-height: 120px;"
                      id="headimgurl"
                    />
                  </a>
                </label>
              </div>
              <div class="form-left">
                <form role="form" method="post" id="memberInfoForm">
                  <div class="form-group">
                    <label>标题</label>
                    <input
                      type="text"
                      class="form-control"
                      name="book_name"
                      id="bookName"
                      placeholder="书籍名称"
                      value
                      v-model="book.name"
                    />
                  </div>
                  <div class="form-group">
                    <label>标识</label>
                    <input
                      type="text"
                      class="form-control"
                      v-model="book.identify"
                      placeholder="书籍唯一标识"
                      disabled
                    />
                  </div>
                  <div class="form-group">
                    <label>语种</label>
                    <select name="lang" class="form-control" v-model="book.lang">
                      <option value="zh">中文</option>
                      <option value="en">英文</option>
                      <option value="other">其他</option>
                    </select>
                  </div>
                  <div class="form-group">
                    <label>来源</label>
                    <div class="input-group">
                      <span class="input-group-btn">
                        <button class="btn btn-default" type="button">来源名称</button>
                      </span>
                      <div id="ebookid" style="display:none">{{book.id}}</div>
                      <input
                        type="text"
                        placeholder="选填"
                        value
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
                        value
                        name="author_url"
                        id="author_url"
                        class="form-control"
                        v-model="book.author_url"
                      />
                    </div>
                  </div>
                  <div class="form-group">
                    <label>描述</label>
                    <textarea
                      rows="3"
                      class="form-control"
                      name="description"
                      style="height: 90px"
                      placeholder="书籍描述"
                      v-model="book.description"
                    ></textarea>
                    <p class="text help-block">描述信息不超过500个字符</p>
                  </div>
                  <div class="form-group">
                    <label>标签</label>
                    <input
                      type="text"
                      class="form-control"
                      name="label"
                      placeholder="书籍标签"
                      value
                      v-model="book.label"
                    />
                    <p class="text help-block">多个标签请用“,”分割</p>
                  </div>
                  <div class="form-group">
                    <button
                      type="button"
                      id="btnSaveBookInfo"
                      class="btn btn-success"
                      data-loading-text="保存中..."
                      @click="SetBook"
                    >保存修改</button>
                    <span id="form-error-message" class="error-message"></span>
                  </div>
                </form>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <!-- Start Modal -->
    <div
      class="modal fade"
      id="upload-logo-panel"
      tabindex="-1"
      role="dialog"
      aria-labelledby="修改封面"
      aria-hidden="true"
    >
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
              <span class="sr-only">Close</span>
            </button>
            <h4 class="modal-title">修改封面</h4>
          </div>
          <div class="modal-body">
            <div class="wraper">
              <div id="image-wraper"></div>
            </div>
            <div class="watch-crop-list">
              <div class="preview-title">预览</div>
              <ul>
                <li>
                  <div class="img-preview preview-lg"></div>
                </li>
                <li>
                  <div class="img-preview preview-sm"></div>
                </li>
              </ul>
            </div>
            <div style="clear: both"></div>
          </div>
          <div class="modal-footer">
            <span id="error-message"></span>
            <div id="filePicker" class="btn">选择</div>
            <button
              type="button"
              id="saveImage"
              class="btn btn-success"
              style="height: 40px;width: 77px;"
              data-loading-text="上传中..."
            >上传</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Header from "@/components/header.vue";
import SettingMenu from "@/components/user/menu.vue";
import $ from "jquery";
import "cropper";
import WebUploader from "webuploader";
import service from "@/utils/request";
export default {
  name: "BookSetting",
  data() {
    return {
      book: {},
    };
  },
  components: {
    Header,
    SettingMenu,
  },
  mounted: function () {
    try {
      var uploader = WebUploader.create({
        auto: false,
        swf: "../../static/webuploader/Uploader.swf",
        server: "http://localhost:8080/user/webUpload",
        pick: "#filePicker",
        fileVal: "image-file",
        fileNumLimit: 1,
        compress: false,
        accept: {
          title: "Images",
          extensions: "jpg,jpeg,png",
          mimeTypes: "image/jpg,image/jpeg,image/png",
        },
      })
        .on("beforeFileQueued", function (file) {
          uploader.reset();
        })
        .on("uploadBeforeSend", function (obj, cropper, headers) {
          console.log(obj.blob);
          console.log(cropper);
          let formData = new FormData();
          var img = obj.blob.source;
          formData.append("image-file", img);
          formData.append("x", cropper.x);
          formData.append("y", cropper.y);
          formData.append("width", cropper.width);
          formData.append("height", cropper.height);
          formData.append("book_id", $("#ebookid").text());
          service({
            url: "/book/uploadCover",
            method: "post",
            data: formData,
            headers: {
              "Content-Type": "multipart/form-data",
            },
          }).then((res) => {
            console.log(res);
            layer.msg(res.data.msg);
            if (res.data.code === 0) {
              this.book.cover = this.res.data.data.cover;
              layer.msg("封面上传成功");
            }
            $("#upload-logo-panel").modal("hide");
          });
        })
        .on("fileQueued", function (file) {
          uploader.makeThumb(
            file,
            function (error, src) {
              var v =
                '<img src="' +
                src +
                '" style="max-width: 360px;max-height: 360px;">';
              if (error) {
                v.replaceWith("<span>不能预览</span>");
                return;
              }
              $("#image-wraper").html(v);
              window.ImageCropper = $("#image-wraper>img").cropper({
                aspectRatio: 1 / 1,
                dragMode: "move",
                viewMode: 1,
                preview: ".img-preview",
              });
            },
            1,
            1
          );
        })
        .on("uploadSuccess", function (file, res) {
          if (res.errcode === 0) {
            console.log(res);
            $("#upload-logo-panel").modal("hide");
            $("#headimgurl").attr("src", res.data);
          } else {
            $("#error-message").text(res.message);
          }
        })
        .on("beforeFileQueued", function (file) {
          if (file.size > 1024 * 1024 * 2) {
            uploader.removeFile(file);
            uploader.reset();
            alert("文件必须小于2MB");
            return false;
          }
        })
        .on("uploadComplete", function () {
          $("#saveImage").button("reset");
        });

      $("#saveImage").on("click", function () {
        var files = uploader.getFiles();
        if (files.length > 0) {
          // var but = $("#saveImage")
          // but.button('loading');
          var cropper = window.ImageCropper.cropper("getData");
          uploader.option("formData", cropper);
          uploader.upload();
        } else {
          alert("请选择封面");
        }
      });
    } catch (e) {
      console.log(e);
    }
  },
  methods: {
    GetBook() {
      service({
        url: "/book/info/" + this.$route.params.bookId,
        method: "get",
      }).then((res) => {
        console.log(res);
        if (res.data.code === 0) {
          this.book = res.data.data;
        } else {
        }
      });
    },
    SetBook() {
      service({
        url: "/book/setting",
        method: "post",
        data: this.book,
      }).then((res) => {
        console.log(res);
        if (res.data.code === 0) {
          layer.msg("设置成功");
        } else {
        }
      });
    },
  },
  created() {
    // watch 路由的参数，以便再次获取数据
    this.$watch(
      () => this.$route.params,
      () => {
        this.GetBook();
      },
      // 组件创建完后获取数据，
      // 此时 data 已经被 observed 了
      { immediate: true }
    );
  },
};
</script>

<style scoped>
@import "https://static.sitestack.cn/static/bootstrap/css/bootstrap.min.css";
@import "https://static.sitestack.cn/static/font-awesome/css/font-awesome.min.css";
@import "https://static.sitestack.cn/static/css/toast.css";
@import "https://static.sitestack.cn/static/css/main.css?version=v2.12.0-beta3";
@import "https://static.sitestack.cn/static/webuploader/webuploader.css";
@import "https://static.sitestack.cn/static/cropper/2.3.4/cropper.min.css";
</style>
