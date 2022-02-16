<template>
    <div class="ucenter-avatar">
        <div class="panel panel-default">
            <div class="panel-heading">个人主页</div>
            <div class="panel-body">
                <img class="thumbnail" src="https://cdn.acwing.com/media/user/profile/photo/94000_lg_5d8f5dcb89.jpg" alt="">
                <div class="ucenter-info">
                    <h3>
                        <span>{{profile.nickname}}</span>
                        <small class="text-muted"></small>
                    </h3>
                    <div class="btns">
                    </div>
                    <div class="user-stat">
                        <ul>
                            <li>
                                <span>加入网站</span> <span>{{ parseInt((Date.now() - Date.parse(profile.create_time))/(1000*60*60*24)) }}</span> <small>天</small>
                            </li>
                            <li>
                                <small>最近一次登录</small> <span>{{ new Date(profile.last_login_time).format("yyyy-MM-dd hh:mm:ss") }}</span>
                            </li>
                        </ul>
                    </div>
                </div>
            </div>
        </div>
        <div class="clearfix">

        </div>
    </div>
</template>

<script>
export default {
  name: 'UserBase',
  props: {profile : Object},
  data () {
    return {
    }
  }
}
</script>
