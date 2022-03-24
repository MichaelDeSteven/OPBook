<template>
  <div class="manual-reader manual-container manual-search-reader single-search">
    <Header></Header>
    <div class="container manual-body">
      <div class="row">
        <div class="col-sm-8 col-xs-12 col-sm-offset-2">
          <form method="get" class="search-form">
            <div class="input-group input-group-lg">
              <input
                type="text"
                name="wd"
                placeholder="Search..."
                class="form-control"
                v-model="req.wd"
              />
              <span class="input-group-addon">
                <i class="fa fa-search"></i>
                <span class="hidden-xs" @click="getResult()">搜索</span>
              </span>
            </div>
            <div class="hidden-xs" style="margin-bottom: 300px;"></div>
          </form>
        </div>
      </div>
    </div>
    <Footer></Footer>
  </div>
</template>


<script>
import Header from "@/components/header.vue";
import Footer from "@/components/footer.vue";
export default {
  name: "Search",
  data() {
    return {
      req: {
        wd: "",
      },
    };
  },
  components: {
    Header,
    Footer,
  },
  methods: {
    getResult() {
      this.$router.push({
        path: "/search/result",
        query: {
          wd: this.req.wd,
        },
      });
    },
  },
};
</script>