<template>
    <header class="navbar navbar-static-top navbar-fixed-top manual-header" role="banner">
        <div class="container" style="position: absolute;">
            <div class="navbar-header col-sm-12 col-md-10 col-lg-10">
                <div class="btn-group dropdown-menu-right pull-right slidebar visible-xs-inline-block visible-sm-inline-block">
                    <button class="btn btn-default dropdown-toggle hidden-lg" type="button" data-toggle="dropdown"><i class="fa fa-align-justify"></i></button>
                    <ul class="dropdown-menu" role="menu">
                        <li>
                            <a href="" class="visible-xs" title="首页"><i class="fa fa-home"></i> 首页</a>
                        </li>
                        <li>
                            <a href="" class="visible-xs" title="发现"><i class="fa fa-globe"></i> 发现</a>
                        </li>
                        <li>
                            <a href="" class="visible-xs" title="标签"><i class="fa fa-tags"></i> 标签</a>
                        </li>
                        <li>
                            <a href="" class="visible-xs visible-sm" title="搜索"><i class="fa fa-search"></i> 搜索</a>
                        </li>

                        <!-- <li>
                            <a href="" title="注册"><i class="fa fa-user-plus"></i> 注册</a>
                        </li>
                        <li>
                            <a href="" title="登录"><i class="fa fa-sign-in"></i> 登录</a>
                        </li> -->
                    </ul>
                </div>

                <a href="" class="navbar-brand" title="">
                    <img class="logo" src="../../static/images/logo.png" alt="">
                </a>
                <nav class="collapse navbar-collapse">

                    <ul class="nav navbar-nav">
                        <li class="">
                            <a href="" title="首页"><i class="fa fa-home"></i> 首页</a>
                        </li>
                        <li class="">
                            <a href="" title="发现"><i class="fa fa-globe"></i> 发现</a>
                        </li>
                        <li class="">
                            <a href="" title="榜单"><img src="/static/images/rank.png" alt=""> 榜单</a>
                        </li>
                        <li class="">
                            <a href="" title="搜索"><i class="fa fa-search"></i> 搜索</a>
                        </li>
                    </ul>
                </nav>
            </div>
        </div>
        <nav class="navbar-collapse hidden-xs hidden-sm" style="position: absolute;right: 0px;" role="navigation">
            <ul class="nav navbar-nav navbar-right" style="margin-right: 15px;">
                <li><a href="" title="注册"><i class="fa fa-user-plus"></i> 注册</a></li>
                <li><a href="" title="登录"><i class="fa fa-sign-in"></i> 登录</a></li>
            </ul>
        </nav>
    </header>
</template>

<style scoped>
    @import 'https://static.sitestack.cn/static/bootstrap/css/bootstrap.min.css';
    @import 'https://static.sitestack.cn/static/font-awesome/css/font-awesome.min.css';
    @import 'https://static.sitestack.cn/static/css/toast.css';
    @import 'https://static.sitestack.cn/static/css/main.css?version=v2.12.0-beta3';
</style>

<script>
export default {
  name: 'Header',
  data () {
    return {
    }
  }
}
</script>
