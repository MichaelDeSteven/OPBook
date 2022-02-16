<template>
  <div class="manual-reader">
    <Header></Header>
    <div class="container-fluid container-widescreen manual-body">
      <div class="row">
        <SettingMenu></SettingMenu>
        <div class="page-right">
          <div class="m-box">
            <div class="box-head">
              <strong class="box-title">修改密码</strong>
            </div>
          </div>
          <div class="box-body">
            <form role="form" method="post" id="securityForm">
              <div class="form-group">
                <label for="password1">原始密码</label>
                <input
                  type="password"
                  name="password1"
                  id="password1"
                  class="form-control disabled"
                  placeholder="原始密码"
                  v-model="pw.password1"
                />
              </div>
              <div class="form-group">
                <label for="password2">新密码</label>
                <input
                  type="password"
                  class="form-control"
                  name="password2"
                  id="password2"
                  max="50"
                  placeholder="新密码"
                  v-model="pw.password2"
                />
              </div>
              <div class="form-group">
                <label for="password3">确认密码</label>
                <input
                  type="password"
                  class="form-control"
                  id="password3"
                  name="password3"
                  placeholder="确认密码"
                  v-model="pw.password3"
                />
              </div>
              <div class="form-group">
                <button type="button" class="btn btn-success" data-loading-text="保存中..." @click="UpdateUserPassword">保存修改</button>
                <span id="form-error-message" class="error-message"></span>
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Header from "@/components/header.vue";
import SettingMenu from "@/components/user/menu.vue";
import service from "@/utils/request";
export default {
  name: "Password",
  data() {
    return {
      pw: {
        password1: "",
        password2: "",
        password3: "",
      },
    };
  },
  components: {
    Header,
    SettingMenu,
  },
  methods: {
    UpdateUserPassword() {
      var oldPasswd = this.pw.password1;
      var newPasswd = this.pw.password2;
      var confirmPassword = this.pw.password3;
      if (!oldPasswd) {
        layer.msg("原始密码不能为空");
        return false;
      }
      if (!newPasswd) {
        layer.msg("新密码不能为空");
        return false;
      }
      if (!confirmPassword) {
        layer.msg("确认密码不能为空");
        return false;
      }
      if (confirmPassword !== newPasswd) {
        layer.msg("确认密码不正确");
        return false;
      }
      var data = {
        oldPassword: oldPasswd,
        newPassword: newPasswd,
      };
      service({
        url: "/user/password",
        method: "post",
        data: data,
      }).then((res) => {
        layer.msg(res.data.msg);
        if (res.data.code === 0) {
          localStorage.setItem("user", JSON.stringify(res.data.data));
          window.location = "/password"
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
@import "https://static.sitestack.cn/static/webuploader/webuploader.css";
@import "https://static.sitestack.cn/static/cropper/2.3.4/cropper.min.css";
</style>