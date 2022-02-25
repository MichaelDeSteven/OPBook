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
                <strong class="box-title">个人设置</strong>
              </div>
            </div>
            <div class="box-body">
              <div class="form-right">
                <label>用户头像</label>
                <label>
                  <a
                    href="javascript:;"
                    data-toggle="modal"
                    title="点击修改头像"
                    data-target="#upload-logo-panel"
                  >
                    <img
                      :src="user.avatar"
                      class="img-circle border"
                      alt="头像"
                      style="margin-left:15px;max-width: 120px;max-height: 120px;"
                      id="headimgurl"
                    />
                  </a>
                </label>
              </div>
              <div class="form-left">
                <form role="form" method="post" id="memberInfoForm">
                  <div class="form-group">
                    <label>邮箱</label>
                    <input type="text" class="form-control" value disabled v-model="user.email" />
                  </div>
                  <div class="form-group">
                    <label>昵称</label>
                    <input
                      type="text"
                      class="form-control"
                      name="nickname"
                      value
                      v-model="user.nickname"
                    />
                  </div>
                  <div class="form-group">
                    <label>手机号</label>
                    <input
                      type="text"
                      class="form-control"
                      id="userPhone"
                      name="phone"
                      maxlength="20"
                      title="手机号码"
                      placeholder="手机号码"
                      value
                      v-model="user.mobile"
                    />
                  </div>
                  <div class="form-group">
                    <button
                      type="button"
                      class="btn btn-success"
                      data-loading-text="保存中..."
                      @click="UpdateUserProfile"
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
      aria-labelledby="修改头像"
      aria-hidden="true"
    >
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal">
              <span aria-hidden="true">&times;</span>
              <span class="sr-only">Close</span>
            </button>
            <h4 class="modal-title">修改头像</h4>
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
            <div id="filePicker" class="btn webuploader-container">选择</div>
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
  name: "Setting",
  data() {
    return {
      user: {
        email: "",
        nickname: "",
        mobile: "",
        avatar: "",
        email: "",
        password: "",
      },
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
          service({
            url: "/user/setting/upload",
            method: "post",
            data: formData,
            headers: {
              "Content-Type": "multipart/form-data",
            },
          }).then((res) => {
            console.log(res);
            layer.msg(res.data.msg);
            if (res.code === 0) {
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
          alert("请选择头像");
        }
      });
    } catch (e) {
      console.log(e);
    }
  },
  methods: {
    GetUserProfile() {
      var u = JSON.parse(localStorage.getItem("user"));
      this.user.nickname = u.nickname;
      this.user.email = u.email;
      this.user.mobile = u.mobile;
      this.user.avatar = u.avatar;
      this.user.email = u.email;
      this.user.password = u.password;
    },
    UpdateUserProfile() {
      service({
        url: "/user/update",
        method: "post",
        data: this.user,
      }).then((res) => {
        layer.msg(res.data.msg);
        if (res.data.code === 0) {
          this.user = res.data.data;
          localStorage.setItem("user", JSON.stringify(res.data.data));
          window.location = "/setting";
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
        this.GetUserProfile();
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
