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
                  <div class="ucenter-content">
                    <UserTab></UserTab>
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
import UserTab from "@/components/user/tab.vue";
import Footer from "@/components/footer.vue";
import service from "@/utils/request";
export default {
  name: "User",
  components: {
    UserBase,
    UserTab,
    Header,
    Footer,
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
