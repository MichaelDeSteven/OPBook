<template>
  <div>
    <Header></Header>
    <div class="container manual-body">
      <div class="row login">
        <div class="col-xs-12">
          <div class="login-body">
            <form role="form" method="post">
              <h3>用户登录</h3>
              <div class="help-block">
                <small>分享知识，共享智慧！知识，因分享，传承久远。</small>
              </div>
              <div class="form-group">
                <div class="input-group">
                  <div class="input-group-addon">
                    <i class="fa fa-user"></i>
                  </div>
                  <input
                    type="text"
                    class="form-control"
                    placeholder="邮箱"
                    name="account"
                    id="account"
                    autocomplete="off"
                    v-model="user.email"
                  />
                </div>
              </div>
              <div class="form-group">
                <div class="input-group">
                  <div class="input-group-addon">
                    <i class="fa fa-lock"></i>
                  </div>
                  <input
                    type="password"
                    class="form-control"
                    placeholder="密码"
                    name="password"
                    id="password"
                    autocomplete="off"
                    v-model="user.password"
                  />
                </div>
              </div>

              <div class="form-group mgt-15px">
                <button
                  type="button"
                  id="btn-login"
                  class="btn btn-success"
                  style="width: 100%"
                  data-loading-text="正在登录..."
                  autocomplete="off"
                  @click="Login"
                >立即登录</button>
              </div>

              <div class="form-group">
                <div class="help-block">
                  <span>
                    没有账号？
                    <a href="\reg" title="使用邮箱注册" class="tooltips text-primary">邮箱注册</a>
                  </span>
                  <span class="pull-right">
                    忘记密码？
                    <a href title="找回密码" class="tooltips text-primary">找回密码</a>
                  </span>
                </div>
                <!-- <hr>
                <div class="help-block">您还可以使用以下方式一键登录</div>
                    <div class="login-by-third">
                        <a class="tooltips" href="" rel="nofollow" title="使用QQ一键登录">
                            <img src="/static/images/qq.png" alt="QQ">
                        </a>
                        <a class="tooltips" href="" rel="nofollow" title="使用GitHub一键登录">
                            <img src="/static/images/github.png" alt="GitHub">
                        </a>
                </div>-->
              </div>
            </form>
          </div>
        </div>
      </div>
      <div class="clearfix"></div>
    </div>
    <Footer></Footer>
  </div>
</template>

<script>
import Header from "@/components/header.vue";
import Footer from "@/components/footer.vue";
import service from "@/utils/request";
export default {
  name: "Login",
  data() {
    return {
      user: {
        email: "",
        password: "",
      },
    };
  },
  components: {
    Header,
    Footer,
  },
  methods: {
    Login() {
      if (this.user.email === "") {
        layer.msg("账号不能为空");
        return false;
      } else if (this.user.password === "") {
        layer.msg("密码不能为空");
        return false;
      }

      service({
        url: "/user/login",
        method: "post",
        data: this.user,
      }).then((res) => {
        console.log(res);
        if (res.data.code === 0) {
          layer.msg(res.data.msg);
          localStorage.setItem("user", JSON.stringify(res.data.data.user));
          localStorage.setItem("token", res.data.data.token);
          window.location = "/user";
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
