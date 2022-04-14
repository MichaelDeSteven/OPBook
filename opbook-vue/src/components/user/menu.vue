<template>
  <div class="page-left">
    <ul class="menu">
      <li class>
        <a href="setting" class="item">
          <i class="fa fa-cogs" aria-hidden="true"></i> 个人设置
        </a>
      </li>
      <li class>
        <a href="password" class="item">
          <i class="fa fa-user-o" aria-hidden="true"></i> 修改密码
        </a>
      </li>
      <!-- <li class=""><a href="" class="item"><i class="fa fa-sitemap" aria-hidden="true"></i> 我的书籍</a> </li> -->
      <!-- <li class>
        <a href class="item">
          <i class="fa fa-heart-o" aria-hidden="true"></i> 我的收藏
        </a>
      </li>-->
      <li class>
        <a href="book" class="item">
          <i class="fa fa-heart-o" aria-hidden="true"></i> 书籍列表
        </a>
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  name: "SettingMenu",
  data() {
    return {};
  },
};
</script>