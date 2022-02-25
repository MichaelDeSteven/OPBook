<template>
  <div>
    <!-- Modal -->
    <!-- 这个Modal主要用于显示思维导图的备注信息。 -->
    <div
      class="modal fade"
      id="ModalNote"
      tabindex="-1"
      role="dialog"
      aria-labelledby="ModalReplaceLabel"
    >
      <div class="modal-dialog" role="document">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal" aria-label="Close">
              <span aria-hidden="true">&times;</span>
            </button>
            <h1 class="modal-title" id="ModalReplaceLabel" style="align:center;">思维导图备注</h1>
          </div>
          <div class="modal-body" style="overflow-y:auto; height:400px;"></div>
          <div class="modal-footer">
            <div class="modal-footer">
              <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="m-manual manual-mode-view manual-reader">
      <header class="navbar navbar-static-top manual-head" role="banner">
        <div class="container-fluid">
          <div class="navbar-header pull-left manual-title">
            <span class="slidebar" id="slidebar">
              <i class="fa fa-align-justify"></i>
            </span>
            <a href title class="book-title">头条小程序框架文档</a>
            <span style="font-size: 12px;font-weight: 100;"></span>
          </div>
          <div class="navbar-header pull-right manual-menu">
            <div class="bookstack-item">
              <a title="首页" href class="btn btn-link">
                <i class="fa fa-home"></i> 首页
              </a>

              <!-- {{if eq $.Member.MemberId $.Model.MemberId}} -->
              <a href title="编辑" class="btn btn-link btn-edit" data-url>
                <i class="fa fa-edit"></i> 编辑文档
              </a>
            </div>
          </div>
        </div>
      </header>
      <div class="container-fluid manual-body">
        <div class="manual-right">
          <div class="manual-article">
            <div class="article-head">
              <div class="container-fluid">
                <div class="row">
                  <div class="col-md-10 col-md-offset-1 col-xs-12 text-center">
                    <h1 id="article-title">页面路由</h1>
                  </div>
                </div>
              </div>
            </div>
            <div class="article-content">
              <div class="help-block article-info" style="padding-bottom: 20px;">
                <span class="alink tooltips text-red" title="AdTitle">
                  <i class="fa fa-angle-right"></i>
                  <a target="_blank" rel="nofollow" href>AdTitle</a>
                  <i class="fa fa-angle-left"></i>
                </span>

                <span>
                  <i class="fa fa-user-o"></i>
                  <a target="_blank" title="来源" href>来源Author</a>
                </span>
                <span>
                  <i class="fa fa-eye"></i> 浏览
                  <i class="read-count">View</i>
                </span>
                <span class="hidden-xs">
                  <a href="#" data-target="#ModalQRcode" data-toggle="modal">
                    <i class="fa fa-qrcode"></i> 扫码
                  </a>
                </span>
                <span>
                  <a href="#" data-target="#ModalShare" data-toggle="modal">
                    <i class="fa fa-share-alt"></i> 分享
                  </a>
                </span>
                <span class="pull-right hidden-xs tooltips" title="更新时间">
                  <i class="fa fa-clock-o"></i>
                  <i class="updated-at">UpdatedAt</i>
                </span>
              </div>
              <article
                data-iframe="https://static.sitestack.cn/static/pdfjs/web/viewer.html?file="
                class="article-body markdown-body editormd-preview-container"
                id="page-content"
              >
                <div class="markdown-toc editormd-markdown-toc">
                  <ul class="markdown-toc-list">
                    <li>
                      <a class="toc-level-3" href="#页面路由" level="3">页面路由</a>
                      <ul>
                        <li>
                          <a class="toc-level-4" href="#页面栈" level="4">页面栈</a>
                        </li>
                        <li>
                          <a
                            class="toc-level-4"
                            href="#getCurrentPages()"
                            level="4"
                          >getCurrentPages()</a>
                        </li>
                        <li>
                          <a class="toc-level-4" href="#路由方式" level="4">路由方式</a>
                          <ul>
                            <li>
                              <a class="toc-level-5" href="#Tips:" level="5">Tips:</a>
                              <ul></ul>
                            </li>
                          </ul>
                        </li>
                      </ul>
                    </li>
                  </ul>
                </div>
                <h3 id="h3-u9875u9762u8DEFu7531">
                  <a name="页面路由" class="reference-link"></a>
                  <span class="header-link octicon octicon-link"></span>页面路由
                </h3>
                <p>框架以栈的形式维护了当前的所有页面。</p>
                <h4 id="h4-u9875u9762u6808">
                  <a name="页面栈" class="reference-link"></a>
                  <span class="header-link octicon octicon-link"></span>页面栈
                </h4>
                <p>当发生路由切换的时候，页面栈的表现如下：</p>
                <table>
                  <thead>
                    <tr>
                      <th>属性</th>
                      <th>类型</th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr>
                      <td>初始化</td>
                      <td>新页面入栈</td>
                    </tr>
                    <tr>
                      <td>打开新页面</td>
                      <td>新页面入栈</td>
                    </tr>
                    <tr>
                      <td>页面重定向</td>
                      <td>当前页面出栈，新页面入栈</td>
                    </tr>
                    <tr>
                      <td>页面返回</td>
                      <td>页面不断出栈，直到目标返回页</td>
                    </tr>
                    <tr>
                      <td>Tab 切换</td>
                      <td>页面全部出栈，只留下新的 Tab 页面</td>
                    </tr>
                    <tr>
                      <td>重加载</td>
                      <td>页面全部出栈，只留下新的页面</td>
                    </tr>
                  </tbody>
                </table>

                <h4 id="h4-getcurrentpages-">
                  <a name="getCurrentPages()" class="reference-link"></a>
                  <span class="header-link octicon octicon-link"></span>getCurrentPages()
                </h4>
                <p>
                  <code>getCurrentPages()</code> 函数用于获取当前页面栈的实例，以数组形式按栈的顺序给出，第一个元素为首页，最后一个元素为当前页面。
                </p>
                <h4 id="h4-u8DEFu7531u65B9u5F0F">
                  <a name="路由方式" class="reference-link"></a>
                  <span class="header-link octicon octicon-link"></span>路由方式
                </h4>
                <p>下面是路由的控制接口：</p>
                <table>
                  <thead>
                    <tr>
                      <th>路由方式</th>
                      <th>触发时机</th>
                      <th>路由前页面</th>
                      <th>路由后页面</th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr>
                      <td>初始化</td>
                      <td>小程序打开的第一个页面</td>
                      <td></td>
                      <td>onLoad, onShow</td>
                    </tr>
                    <tr>
                      <td>打开新页面</td>
                      <td>
                        调用 API
                        <code>tt.navigateTo</code> 或使用组件
                        <code>&lt;navigator open-type=&#34;navigateTo&#34;/&gt;</code>
                      </td>
                      <td>onHide</td>
                      <td>onLoad, onShow</td>
                    </tr>
                    <tr>
                      <td>页面重定向</td>
                      <td>
                        调用 API
                        <code>tt.redirectTo</code> 或使用组件
                        <code>&lt;navigator open-type=&#34;redirectTo&#34;/&gt;</code>
                      </td>
                      <td>onUnload</td>
                      <td>onLoad, onShow</td>
                    </tr>
                    <tr>
                      <td>页面返回</td>
                      <td>
                        调用 API
                        <code>tt.navigateBack</code> 或使用组件
                        <code>&lt;navigator open-type=&#34;navigateBack&#34;&gt;</code>或用户按左上角返回按钮
                      </td>
                      <td>onUnload</td>
                      <td>onShow</td>
                    </tr>
                    <tr>
                      <td>Tab 切换</td>
                      <td>
                        调用 API
                        <code>tt.switchTab</code> 或使用组件
                        <code>&lt;navigator open-type=&#34;switchTab&#34;/&gt;</code> 或用户切换 Tab
                      </td>
                      <td></td>
                      <td>各种情况请参考下表</td>
                    </tr>
                    <tr>
                      <td>重启动</td>
                      <td>
                        调用 API
                        <code>tt.reLaunch</code> 或使用组件
                        <code>&lt;navigator open-type=&#34;reLaunch&#34;/&gt;</code>
                      </td>
                      <td>onUnload</td>
                      <td>onLoad, onShow</td>
                    </tr>
                  </tbody>
                </table>

                <p>Tab 切换对应的生命周期（以 A、B 页面为 Tabbar 页面，C 是从 A 页面打开的页面，D 页面是从 C 页面打开的页面为例）：</p>
                <table>
                  <thead>
                    <tr>
                      <th>当前页面</th>
                      <th>路由后页面</th>
                      <th>触发的生命周期（按顺序）</th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr>
                      <td>A</td>
                      <td>A</td>
                      <td>Nothing happend</td>
                    </tr>
                    <tr>
                      <td>A</td>
                      <td>B</td>
                      <td>A.onHide(), B.onLoad(), B.onShow()</td>
                    </tr>
                    <tr>
                      <td>A</td>
                      <td>B（再次打开）</td>
                      <td>A.onHide(), B.onShow()</td>
                    </tr>
                    <tr>
                      <td>C</td>
                      <td>A</td>
                      <td>C.onUnload(), A.onShow()</td>
                    </tr>
                    <tr>
                      <td>C</td>
                      <td>B</td>
                      <td>C.onUnload(), B.onLoad(), B.onShow()</td>
                    </tr>
                    <tr>
                      <td>D</td>
                      <td>B</td>
                      <td>D.onUnload(), C.onUnload(), B.onLoad(), B.onShow()</td>
                    </tr>
                    <tr>
                      <td>D（从转发进入）</td>
                      <td>A</td>
                      <td>D.onUnload(), A.onLoad(), A.onShow()</td>
                    </tr>
                    <tr>
                      <td>D（从转发进入）</td>
                      <td>B</td>
                      <td>D.onUnload(), B.onLoad(), B.onShow()</td>
                    </tr>
                  </tbody>
                </table>

                <h5 id="h5-tips-">
                  <a name="Tips:" class="reference-link"></a>
                  <span class="header-link octicon octicon-link"></span>Tips:
                </h5>
                <ul>
                  <li>navigateTo, redirectTo 只能打开非 tabBar 页面。</li>
                  <li>switchTab 只能打开 tabBar 页面。</li>
                  <li>reLaunch 可以打开任意页面。</li>
                  <li>页面底部的 tabBar 由页面决定，即只要是定义为 tabBar 的页面，底部都有 tabBar。</li>
                  <li>调用页面路由带的参数可以在目标页面的onLoad中获取。</li>
                </ul>
                <blockquote>
                  <p>
                    原文:
                    <a
                      href="https://developer.toutiao.com/docs/framework/pageRoute.html"
                    >https://developer.toutiao.com/docs/framework/pageRoute.html</a>
                  </p>
                </blockquote>
              </article>
              <div class="row hung-read-link">
                <div class="col-xs-12 hung-pre" style="display: none">
                  <span class="text-muted">上一篇:</span>
                  <a href="#"></a>
                </div>
                <div class="col-xs-12 hung-next" style="display: none">
                  <span class="text-muted">下一篇:</span>
                  <a href="#"></a>
                </div>
              </div>

              <div class="bookstack-bars">
                <ul>
                  <li class="visible-xs visible-sm bars-menu bars-menu-hide">
                    <a href="/" title="首页">
                      <i class="fa fa-home"></i>
                    </a>
                  </li>
                </ul>
              </div>
            </div>
          </div>
        </div>

        <div class="manual-left">
          <div class="article-search" data-bookid>
            <form id="searchForm" action method="post">
              <div class="input-group">
                <input
                  type="text"
                  name="keyword"
                  placeholder="Search..."
                  autocomplete="off"
                  class="form-control"
                />
                <span class="input-group-addon input-group-addon-clear">
                  <i class="fa fa-remove"></i>
                </span>
                <span class="input-group-addon">
                  <button type="submit">
                    <i class="fa fa-search"></i>
                  </button>
                </span>
              </div>
            </form>
            <div class="pull-right hidden-xs">
              <i class="fa fa-align-justify"></i>
            </div>
          </div>
          <div class="article-menu">
            <div class="article-menu-detail collapse-menu">
              <ul>
                <li id="91373" class="readed">
                  <a href="/read/toutiao-mini-framework/intro.md" title="介绍">介绍</a>
                </li>
                <li id="91350" class="collapse-node">
                  <span></span>
                  <a href="/read/toutiao-mini-framework/base.md" title="基础说明">基础说明</a>
                  <ul>
                    <li id="91324">
                      <a href="/read/toutiao-mini-framework/1.md" title="目录结构">目录结构</a>
                    </li>
                    <li id="91325">
                      <a href="/read/toutiao-mini-framework/2.md" title="全局配置">全局配置</a>
                    </li>
                  </ul>
                </li>
                <li id="91351" class="jstree-open collapse-node">
                  <span></span>
                  <a href="/read/toutiao-mini-framework/logic.md" title="逻辑层">逻辑层</a>
                  <ul>
                    <li id="91326">
                      <a href="/read/toutiao-mini-framework/4.md" title="启动程序">启动程序</a>
                    </li>
                    <li id="91327">
                      <a href="/read/toutiao-mini-framework/5.md" title="场景值">场景值</a>
                    </li>
                    <li id="91328" class="readed">
                      <a
                        href="/read/toutiao-mini-framework/6.md"
                        title="启动页面"
                        class="jstree-clicked"
                      >启动页面</a>
                    </li>
                    <li id="91329">
                      <a href="/read/toutiao-mini-framework/7.md" title="页面路由">页面路由</a>
                    </li>
                    <li id="91330">
                      <a href="/read/toutiao-mini-framework/8.md" title="文件作用域">文件作用域</a>
                    </li>
                    <li id="91331">
                      <a href="/read/toutiao-mini-framework/9.md" title="API">API</a>
                    </li>
                  </ul>
                </li>
                <li id="91348" class="collapse-node">
                  <span></span>
                  <a href="/read/toutiao-mini-framework/view.md" title="视图层">视图层</a>
                  <ul>
                    <li id="91332">
                      <a href="/read/toutiao-mini-framework/11.md" title="TTML">TTML</a>
                    </li>
                    <li id="91333">
                      <a href="/read/toutiao-mini-framework/12.md" title="TTML-数据绑定">TTML-数据绑定</a>
                    </li>
                    <li id="91334">
                      <a href="/read/toutiao-mini-framework/13.md" title="TTML-列表渲染">TTML-列表渲染</a>
                    </li>
                    <li id="91335">
                      <a href="/read/toutiao-mini-framework/14.md" title="TTML-条件渲染">TTML-条件渲染</a>
                    </li>
                    <li id="91336">
                      <a href="/read/toutiao-mini-framework/15.md" title="TTML-模板">TTML-模板</a>
                    </li>
                    <li id="91337">
                      <a href="/read/toutiao-mini-framework/16.md" title="TTML-事件">TTML-事件</a>
                    </li>
                    <li id="91338">
                      <a href="/read/toutiao-mini-framework/17.md" title="TTML-引用">TTML-引用</a>
                    </li>
                    <li id="91339">
                      <a href="/read/toutiao-mini-framework/18.md" title="TTSS">TTSS</a>
                    </li>
                    <li id="91340">
                      <a href="/read/toutiao-mini-framework/19.md" title="组件">组件</a>
                    </li>
                  </ul>
                </li>
                <li id="91349" class="collapse-node">
                  <span></span>
                  <a href="/read/toutiao-mini-framework/custom.md" title="自定义组件">自定义组件</a>
                  <ul>
                    <li id="91341">
                      <a href="/read/toutiao-mini-framework/21.md" title="自定义组件">自定义组件</a>
                    </li>
                    <li id="91342">
                      <a href="/read/toutiao-mini-framework/22.md" title="组件模版和样式">组件模版和样式</a>
                    </li>
                    <li id="91343">
                      <a
                        href="/read/toutiao-mini-framework/23.md"
                        title="Component 构造器"
                      >Component 构造器</a>
                    </li>
                    <li id="91344">
                      <a href="/read/toutiao-mini-framework/24.md" title="组件间通信与事件">组件间通信与事件</a>
                    </li>
                    <li id="91345">
                      <a
                        href="/read/toutiao-mini-framework/25.md"
                        title="组件行为 behaviors"
                      >组件行为 behaviors</a>
                    </li>
                    <li id="91346">
                      <a href="/read/toutiao-mini-framework/26.md" title="抽象节点">抽象节点</a>
                    </li>
                    <li id="91347">
                      <a href="/read/toutiao-mini-framework/27.md" title="组件扩展">组件扩展</a>
                    </li>
                  </ul>
                </li>
              </ul>
            </div>
            <div class="search-result">
              <div class="search-empty" style="display: block;">
                <i class="fa fa-search-plus" aria-hidden="true"></i>
                <b class="text">暂无相关搜索结果！</b>
              </div>
              <ul class="search-list" id="searchList"></ul>
            </div>
          </div>

          <div class="m-copyright">
            <p>
              本文档使用
              <a
                href="//www.bookstack.cn"
                style="color: #0e8c8c;font-weight: bold"
                target="_blank"
              >BookStack</a> 构建
            </p>
          </div>
        </div>

        <div class="manual-progress">
          <b class="progress-bar"></b>
        </div>
      </div>
      <div class="manual-mask"></div>
    </div>
    <span class="article-toggle tooltips hidden-xs hidden-sm" title="展开/收起文章目录">
      <img alt="展开/收起文章目录" src="https://static.sitestack.cn/static/images/toggle.png" />
    </span>
    <div id="menu-hidden" style="display: none;"></div>
  </div>
</template>

<script>
export default {
  name: "READ",
  components: {},
};
</script>

<style scoped>
@import "https://static.sitestack.cn/static/bootstrap/css/bootstrap.min.css";
@import "https://static.sitestack.cn/static/font-awesome/css/font-awesome.min.css";
@import "https://static.sitestack.cn/static/editor.md/css/editormd.preview.css";
@import "https://static.sitestack.cn/static/css/toast.css";
@import "https://static.sitestack.cn/static/page/page.css";
@import "https://static.sitestack.cn/static/katex/katex.min.css";
@import "https://static.sitestack.cn/static/mind-map/mindmap.css";
@import "https://static.sitestack.cn/static/css/bookstack.css?version=v2.12.0-beta31008";
</style>

<style>
.editormd-preview-container ol.linenums li code,
.editormd-html-preview ol.linenums li code {
  display: block;
  white-space: pre;
}
li.L1,
li.L3,
li.L5,
li.L7,
li.L9 {
  background-color: transparent;
}
.editormd-preview-container pre.prettyprint,
.editormd-html-preview pre.prettyprint {
  border-color: transparent;
}
body {
  -webkit-overflow-scrolling: touch;
}
.alink,
.alink a {
  color: #888 !important;
}
.alink:hover,
.alink a:hover {
  color: red !important;
}
</style>
