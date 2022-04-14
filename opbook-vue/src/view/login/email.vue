<template>
  <div>
    <Header></Header>
    <div class="container manual-body">
      <div class="row login">
        <div class="col-xs-12">
          <div class="login-body">
            <form role="form" class="ajax-form" data-url action method="post">
              <h3>
                用户注册
                <small class="pull-right text-muted"></small>
              </h3>
              <div class="form-group has-hidden">
                <div class="input-group">
                  <div class="input-group-addon">用户邮箱</div>
                  <input
                    type="email"
                    class="form-control"
                    placeholder
                    name="email"
                    id="email"
                    value
                    autocomplete="off"
                    v-model="user.email"
                  />
                </div>
              </div>
              <div class="form-group has-hidden">
                <div class="input-group">
                  <div class="input-group-addon">用户昵称</div>
                  <input
                    type="text"
                    class="form-control"
                    placeholder
                    name="nickname"
                    id="nickname"
                    value
                    autocomplete="off"
                    v-model="user.nickname"
                  />
                </div>
              </div>
              <div class="form-group">
                <div class="input-group">
                  <div class="input-group-addon">登录密码</div>
                  <input
                    type="password"
                    class="form-control"
                    placeholder
                    name="password1"
                    id="password1"
                    autocomplete="off"
                    v-model="user.password1"
                  />
                </div>
              </div>
              <div class="form-group has-hidden">
                <div class="input-group">
                  <div class="input-group-addon">确认密码</div>
                  <input
                    type="password"
                    class="form-control"
                    placeholder
                    name="password2"
                    id="password2"
                    autocomplete="off"
                    v-model="user.password2"
                  />
                </div>
              </div>
              <input type="hidden" name="avatar" value />
              <input type="hidden" name="id" value />
              <input type="hidden" name="oauth" value />

              <div class="help-block">
                <span class="pull-right">
                  已有账号？
                  <a href="/login" title="立即登录" class="tooltips text-primary">立即登录</a>
                </span>
              </div>

              <div class="form-group mgt-15px">
                <button
                  type="button"
                  class="btn btn-success btn-block"
                  data-has="绑定已有账号"
                  data-not="注册新账号"
                  @click="Reg"
                >注册新账号</button>
              </div>
              <!-- <div class="form-group login-by-third">
                                <a class="tooltips" href="" rel="nofollow" title="使用QQ一键登录">
                                    <img src="/static/images/qq.png" alt="QQ">
                                </a>
                                <a class="tooltips" href="" rel="nofollow" title="使用GitHub一键登录">
                                    <img src="/static/images/github.png" alt="GitHub">
                                </a>
              </div>-->
            </form>
          </div>
        </div>
      </div>
      <div class="clearfix"></div>
    </div>
  </div>
</template>

<script>
import Header from "@/components/header.vue";
import service from "@/utils/request";
export default {
  name: "Email",
  data() {
    return {
      user: {
        email: "",
        password1: "",
        password2: "",
        nickname: "",
      },
    };
  },
  components: {
    Header,
  },
  methods: {
    Reg() {
      if (this.user.email === "") {
        layer.msg("邮箱不能为空");
        return false;
      } else if (this.user.password1 === "" || this.user.password2 === "") {
        layer.msg("密码不能为空");
        return false;
      } else if (this.user.password1 !== this.user.password2) {
        layer.msg("密码输入不一致");
        return false;
      } else if (this.user.nickname === "") {
        layer.msg("昵称不能为空");
        return false;
      }
      var data = {
        email: this.user.email,
        password: this.user.password1,
        nickname: this.user.nickname,
      };
      console.log(data);
      service({
        url: "/user/reg",
        method: "post",
        data: data,
      }).then((res) => {
        console.log(res);
        if (res.code === 0) {
          layer.msg(res.data.msg);
          window.location = "/login";
        } else {
          layer.msg(res.data.msg);
        }
      });
    },
  },
};
</script>

<style>
h3 {
  font-size: 20px;
  font-weight: normal;
  margin: 15px auto;
}
.login .login-body {
  padding-bottom: 5px;
}
</style>
