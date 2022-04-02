<template>
  <div>
    <ul class="nav nav-tabs">
      <!-- <li>
      <a href>分享</a>
      </li>-->
      <li>
        <a href="javascript:;" @click="select=0" :class="{'active':select===0}">收藏</a>
      </li>
      <li>
        <a href="javascript:;" @click="select=1" :class="{'active':select===1}">关注</a>
      </li>
      <li>
        <a href="javascript:;" @click="select=2" :class="{'active':select===2}">粉丝</a>
      </li>
    </ul>
    <div v-show="0"></div>
    <div v-show="1"></div>
    <div v-show="2">
      <Fan></Fan>
    </div>
  </div>
</template>

<script>
import Fan from "@/components/user/fans.vue";
export default {
  name: "UserTab",
  components: {
    Fan,
  },
  data() {
    return {
      select: 1,
    };
  },
};
</script>
