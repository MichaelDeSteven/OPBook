<template>
  <div class="manual-reader">
    <Header></Header>
    <div class="container-fluid container-widescreen manual-body">
      <div class="row">
        <SettingMenu></SettingMenu>
        <div class="page-right">
          <div class="m-box">
            <div class="box-head">
              <strong class="box-title">我的私信</strong>
            </div>
          </div>
          <div class="panel-body">
            <div class="chat-left">
              <ul class="list-group">
                <template v-for="item in chatList">
                  <li :class="['li-item', {'select':userInfo.id === item.id}]" @click="get(item)">
                    <div class="list-item-head">
                      <img
                        class="img-thumbnail img-circle img-responsive"
                        alt="头像"
                        src="https://opbook-1304588036.cos.ap-hongkong.myqcloud.com/file/16daaa5d9d011c88.jpg"
                      />
                    </div>
                    <a>
                      <i aria-hidden="true"></i>
                      {{item.nickname}}
                    </a>
                  </li>
                </template>
              </ul>
            </div>
            <div class="chat-right">
              <div class="chat-hd">
                <span>和{{userInfo.nickname}}对话</span>
              </div>
              <div class="chat-content">
                <div class="chat-msg-list">
                  <div class="clearfix chat-msg-item" v-for="item in conversation">
                    <a
                      target="_blank"
                      :class="[{'list-head-me':item.from_id !== userInfo.id}, {'list-head':item.from_id === userInfo.id}]"
                    >
                      <img
                        class="img-thumbnail img-circle img-responsive"
                        alt="头像"
                        src="https://opbook-1304588036.cos.ap-hongkong.myqcloud.com/file/16daaa5d9d011c88.jpg"
                      />
                    </a>
                    <div
                      :class="[{'tip-right-me':item.from_id !== userInfo.id}, {'tip-right':item.from_id === userInfo.id}]"
                    >
                      <div class="tip-inner">{{item.content}}</div>
                      <div>
                        <span>{{ new Date(item.create_time).format("yyyy-MM-dd hh:mm:ss") }}</span>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
              <div class="chat-ed">
                <textarea
                  class="form-control"
                  name="content"
                  rows="5"
                  placeholder="请输入私信内容"
                  v-model="message.content"
                ></textarea>
                <div class="form-group">
                  <span class="pull-left text-muted"></span>
                  <input
                    type="button"
                    class="btn btn-success pull-right"
                    value="发送"
                    @click="sendChat"
                  />
                </div>
              </div>
            </div>
          </div>
          <div class="container-fluid container-widescreen manual-body"></div>
        </div>
      </div>
    </div>
  </div>
</template>


<script>
import Header from "@/components/header.vue";
import SettingMenu from "@/components/user/menu.vue";
import service from "@/utils/request";
export default {
  name: "Whisper",
  data() {
    return {
      chatList: [],
      message: {
        from_id: 0,
        to_id: 0,
        content: "",
      },
      userInfo: {},
      conversation: [],
    };
  },
  created() {
    // 获取私信列表
    service({
      url: "/social/chat/getUserList",
      method: "get",
    }).then((res) => {
      console.log(res.data.data);
      this.chatList = res.data.data;
      var toid = parseInt(this.$route.params.uid);
      var inx = -1;
      for (var i = 0; i < this.chatList.length; i++) {
        if (this.chatList[i].id === toid) {
          inx = i;
          this.userInfo = this.chatList[i];
          break;
        }
      }
      // 私信用户不在列表则追加
      if (inx === -1) {
        this.pushUser(toid);
      }
    });
    // 获取与该用户的私信内容
    var u = JSON.parse(localStorage.getItem("user"));
    if (u != null) {
      this.message.from_id = parseInt(u.id);
    } else {
      alert("用户未登录！");
      return;
    }
    this.message.to_id = parseInt(this.$route.params.uid);
    this.getConversation();
  },
  components: {
    Header,
    SettingMenu,
  },
  methods: {
    sendChat() {
      if (this.message.content.length === 0) {
        layer.msg("内容不能为空！");
        return;
      }
      var m = {
        from_id: this.message.from_id,
        to_id: this.message.to_id,
        content: this.message.content,
        create_time: Date.now(),
      };
      service({
        url: "/social/chat/add",
        method: "post",
        data: m,
      }).then((res) => {
        console.log(res);
        this.message.content = "";
        this.conversation.push(m);
      });
    },
    getConversation() {
      console.log(this.message);
      service({
        url: "/social/chat/getConversation",
        method: "post",
        data: this.message,
      }).then((res) => {
        console.log(res);
        this.conversation = res.data.data;
      });
    },
    pushUser(uid) {
      service({
        url: "/user/" + uid,
        method: "get",
      }).then((res) => {
        console.log(res);
        this.chatList.push(res.data.data);
        this.userInfo = this.chatList[this.chatList.length - 1];
      });
    },
    get(item) {
      this.conversation = [];
      this.message.to_id = item.id;
      this.getConversation();
      this.userInfo = item;
    },
  },
};
</script>

<style>
.chat-left {
  position: fixed;
  overflow: auto;
  top: 110px;
  bottom: 10px;
  left: 212px;
  z-index: 100;
  width: 350px;
  background-color: #f5f5f5;
  border-right: 1px solid #eaeaea;
}
.chat-right {
  position: fixed;
  overflow: auto;
  top: 110px;
  bottom: 10px;
  left: 580px;
  z-index: 100;
  width: 1110px;
  background-color: #f5f5f5;
  border-right: 1px solid #eaeaea;
}
.chat-hd {
  border-bottom: 1px solid #ededed;
  background: #f6f6f6;
  text-align: center;
  font-size: 16px;
  padding: 10px 20px;
  border-radius: 4px 4px 0 0;
  color: #3d444d;
}
.chat-content {
  height: 450px;
  background: #fff;
  border-left: 1px solid #eaeaea;
  overflow: auto;
}
.chat-ed {
  background: #fff;
  border: 1px solid #eaeaea;
  height: 187px;
  padding: 20px;
}
.chat-msg-list {
  margin-bottom: 30px;
}
.chat-msg-item {
  margin-bottom: 20px;
  height: 40px;
  padding: 10px;
}
.list-head {
  display: block;
  float: left;
  width: 40px;
  height: 40px;
  overflow: hidden;
  border: 1px solid #ededed;
  margin-right: 15px;
  border-radius: 50%;
}
.list-head-me {
  display: block;
  float: right;
  width: 40px;
  height: 40px;
  overflow: hidden;
  border: 1px solid #ededed;
  margin-right: 15px;
  border-radius: 50%;
}
.list-item-head {
  display: block;
  float: left;
  width: 60px;
  height: 60px;
  overflow: hidden;
  border: 1px solid #ededed;
  margin-right: 15px;
  border-radius: 50%;
}
.tip-inner {
  line-height: 18px;
  border-radius: 6px;
  box-shadow: 0 0 5px #f0f0f0;
  display: inline-block;
  background: #ebf8f5;
  max-width: 537px;
  padding: 10px;
  text-align: left;
  color: #666;
}
.tip-right-me {
  padding: 0 5px;
  margin-right: 3px;
  float: right;
}
.tip-right {
  padding: 0 5px;
  margin-left: 3px;
}
.li-item {
  height: 60px;
}
.select {
  background-color: rgba(0, 10, 32, 0.03);
}
</style>