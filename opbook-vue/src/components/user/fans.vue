<template>
  <div class="ucenter">
    <div class="container">
      <div class="row">
        <div class="col-xs-12 col-sm-4 col-md-3">
          <UserBase></UserBase>
        </div>
        <div class="col-xs-12 col-sm-8 col-md-9">
          <div class="row">
            <div class="col-xs-12">
              <UserTab></UserTab>
            </div>
            <div class="col-xs-12">
              <div class="ucenter-content"></div>
            </div>
            <div class="pagination-container"></div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "UserBase",
  data() {
    return {};
  },
};
</script>
