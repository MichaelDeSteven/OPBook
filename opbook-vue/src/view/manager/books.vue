<template>
  <div>
    <div class="manual-reader">
      <div class="container manual-body">
        <div class="row">
          <div class="page-left"></div>
          <div class="page-right">
            <div class="m-box">
              <div class="box-head">
                <strong class="box-title">书籍列表</strong>
                <a
                  href="javascript:void(0);"
                  data-toggle="modal"
                  data-target="#deleteBooksModal"
                  class="btn btn-danger pull-right"
                >批量删除</a>
              </div>
              <ul class="nav nav-tabs" style="margin-top: 15px;">
                <li class>
                  <a href>公开书籍</a>
                </li>
                <li class>
                  <a href>私有书籍</a>
                </li>
              </ul>
              <br />
              <form class="bs-example bs-example-form" role="form">
                <div class="input-group">
                  <input type="hidden" name="private" value />
                  <input type="text" name="wd" value placeholder="搜索书籍" class="form-control" />
                  <span class="input-group-btn">
                    <button class="btn btn-default" type="submit">搜索</button>
                  </span>
                </div>
              </form>
            </div>
            <div class="box-body" id="bookList">
              <div class="book-list">
                <div class="text-center">暂无数据</div>
              </div>
              <nav class="pagination-container"></nav>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {};
</script>
