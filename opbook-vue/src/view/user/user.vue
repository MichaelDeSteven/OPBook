<template>
  <div>
    <div class="manual-reader manual-container">
      <Header></Header>
      <div class="ucenter">
        <div class="container-fluid container-widescreen">
          <div class="row">
            <div class="col-xs-12 col-sm-4 col-md-3 col-lg-2">
              <UserBase :profile="profile"></UserBase>
            </div>
            <div class="col-xs-12 col-sm-8 col-md-9 col-lg-10">
              <div class="row">
                <div class="col-xs-12"></div>
                <div class="col-xs-12">
                  <ul class="nav nav-tabs">
                    <li :class="{'active':select===0}">
                      <a href="javascript:;" @click="select=0">收藏</a>
                    </li>
                    <li :class="{'active':select===1}">
                      <a href="javascript:;" @click="select=1">关注</a>
                    </li>
                    <li :class="{'active':select===2}">
                      <a href="javascript:;" @click="select=2">粉丝</a>
                    </li>
                  </ul>
                  <div v-if="select===0">
                    <Collection></Collection>
                  </div>
                  <div v-if="select===1">
                    <Follow></Follow>
                  </div>
                  <div v-if="select===2">
                    <Fan></Fan>
                  </div>
                </div>
                <div class="pagination-container">
                  <!-- {{.PageHtml}} -->
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <Footer></Footer>
  </div>
</template>

<script>
import Header from "@/components/header.vue";
import UserBase from "@/components/user/base.vue";
import Fan from "@/components/user/fans.vue";
import Follow from "@/components/user/follow.vue";
import Collection from "@/components/user/collection.vue";
import Footer from "@/components/footer.vue";
import service from "@/utils/request";
export default {
  name: "User",
  components: {
    UserBase,
    Header,
    Footer,
    Fan,
    Follow,
    Collection,
  },
  data() {
    return {
      select: 0,
    };
  },
  props: { profile: Object },
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
  methods: {
    GetUserProfile() {
      service({
        url: "/user/" + this.$route.params.uid,
        method: "get",
      }).then((res) => {
        // layer.msg(res.data.msg);
        if (res.data.code === 0) {
          this.profile = res.data.data;
        } else {
        }
      });
    },
  },
};
</script>
