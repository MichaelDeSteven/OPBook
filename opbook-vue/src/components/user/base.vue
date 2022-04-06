<template>
  <div class="ucenter-avatar">
    <div class="panel panel-default">
      <div class="panel-heading">个人主页</div>
      <div class="panel-body">
        <img class="thumbnail" :src="profile.avatar" alt />
        <div class="ucenter-info">
          <h3>
            <span>{{profile.nickname}}</span>
            <small class="text-muted"></small>
          </h3>
          <div class="btns">
            <template v-if="self">
              <a target href="/setting" class="btn btn-success btn-sm">
                <i class="fa fa-cogs"></i> 个人设置
              </a>
            </template>
            <template v-else>
              <template v-if="follow.follow_stat">
                <a class="btn btn-default btn-sm btn-cancel" @click="Follow">
                  <i class="fa fa-heart text-danger"></i> 取消关注
                </a>
              </template>
              <template v-else>
                <a class="btn btn-success btn-sm btn-follow" @click="Follow">
                  <i class="fa fa-heart-o"></i> 关注Ta
                </a>
              </template>
              <a target :href="'/whisper/' + profile.id" class="btn btn-success btn-sm">
                <i></i> 私信
              </a>
            </template>
          </div>
          <div class="user-stat">
            <ul>
              <li>
                <span>加入网站</span>
                <span>{{ parseInt((Date.now() - Date.parse(profile.create_time))/(1000*60*60*24)) }}</span>
                <small>天</small>
              </li>
              <li>
                <small>最近一次登录</small>
                <span>{{ new Date(profile.last_login_time).format("yyyy-MM-dd hh:mm:ss") }}</span>
              </li>
            </ul>
          </div>
        </div>
      </div>
    </div>
    <div class="clearfix"></div>
  </div>
</template>

<script>
import service from "@/utils/request";
export default {
  name: "UserBase",
  props: {
    profile: Object,
  },
  data() {
    return {
      self: false,
      follow: {
        user_id: 0,
        followee_id: 0,
        follow_stat: false,
      },
    };
  },
  created() {
    var u = JSON.parse(localStorage.getItem("user"));
    if (u != null) {
      this.follow.followee_id = parseInt(this.$route.params.uid);
      this.follow.user_id = parseInt(u.id);
      console.log(this.follow);
      if (this.follow.followee_id === this.follow.user_id) {
        this.self = true;
      } else {
        service({
          url: "/social/fan/stat",
          method: "post",
          data: this.follow,
        }).then((res) => {
          this.follow = res.data.data;
          console.log(res);
        });
      }
    }
  },
  methods: {
    Follow() {
      service({
        url: "/social/fan/follow",
        method: "post",
        data: this.follow,
      }).then((res) => {
        this.follow = res.data.data;
        console.log(res);
      });
    },
  },
};
</script>
