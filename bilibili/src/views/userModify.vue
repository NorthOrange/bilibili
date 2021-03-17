<!-- 用户编辑资料页面 -->
<template>
  <div class="usermodify">
    <top-bar></top-bar>
    <div class="modifybox">
      <!-- 头像上传 -->
      <div class="uploadfile">
        <van-uploader
          accept="image/*"
          class="uploadimg"
          preview-size="100vw"
          :after-read="afterReadAvatar"
        />
        <edit-bar left="头像">
          <img
            :src="model.avatar"
            style="width: 16vw; height: 16vw; border-radius: 50%"
            slot="right"
          />
        </edit-bar>
      </div>

      <!-- 昵称修改 -->
      <edit-bar
        left="昵称"
        @editbarClick="nameShow = true"
      >
        <a
          href="javascript:;"
          slot="right"
        >{{ model.name }}</a>
      </edit-bar>

      <van-dialog
        v-model="nameShow"
        title="昵称"
        confirm-button-color="#0000ee"
        show-cancel-button
        @confirm="nameClick"
        @cancel="content = ''"
      >
        <van-field
          type="text"
          maxlength="7"
          v-model="content"
          autofocus
        />
      </van-dialog>

      <edit-bar left="帐号">
        <a
          href="javascript:;"
          slot="right"
        >{{ model.account }}</a>
      </edit-bar>

      <edit-bar
        left="性别"
        @editbarClick="sexShow = true"
      >
        <a
          href="javascript:;"
          slot="right"
          v-if="model.sex"
        >{{ model.sex }}</a>
        <a
          href="javascript:;"
          slot="right"
          v-else
        >外星人</a>
      </edit-bar>

      <van-action-sheet
        v-model="sexShow"
        :actions="actions"
        cancel-text="取消"
        close-on-click-action
        @select="sexSelect"
      />

      <!-- 个性签名 -->
      <edit-bar
        left="个性签名"
        @editbarClick="introShow = true"
      >
        <a
          href="javascript:;"
          slot="right"
          v-if="model.introduction"
        >{{
          model.introduction
        }}</a>
        <a
          href="javascript:;"
          slot="right"
          v-else
        >这个人很懒, 什么都没有写</a>
      </edit-bar>

      <van-dialog
        v-model="introShow"
        title="个性签名"
        confirm-button-color="#0000ee"
        show-cancel-button
        @confirm="introClick"
        @cancel="content = ''"
      >
        <van-field
          type="textarea"
          maxlength="30"
          v-model="content"
          autofocus
        />
      </van-dialog>

    </div>

  </div>
</template>

<script>
import topBar from "../components/topBar.vue";
import editBar from "../components/editBar.vue";
export default {
  data() {
    return {
      model: {},
      nameShow: false,
      sexShow: false,
      introShow: false,
      actions: [{ name: "男" }, { name: "女" }, { name: "外星人" }],
      content: "",
    };
  },
  components: {
    topBar,
    editBar,
  },
  methods: {
    UserinfoData() {
      // 请求本地用户的数据
      this.$request({
        methods: "get",
        url: "/api/user/info/" + localStorage.getItem("id"),
      }).then((res) => {
        this.model = res.data;
      });
    },
    afterReadAvatar(file) {
      // 头像上传
      console.log("拿到一个文件: " + file.file);
      if (
        file.file.type != "image/png" &&
        file.file.type != "image/jpg" &&
        file.file.type != "image/jpeg"
      ) {
        this.$msg.fail("请上传一张图片!");
        return;
      }
      var formdata = new FormData();
      formdata.append("avatar", file.file);
      formdata.append("id", localStorage.getItem("id"));
      this.$request({
        methods: "post",
        url: "/api/avatar/upload",
        data: formdata,
        type: "formdata",
      }).then((res) => {
        this.model.avatar = res.data.avatar;
      });
    },
    // 昵称修改
    nameClick() {
      if (!/^[^ ]{2,7}$/.test(this.content)) {
        this.content = "";
        this.$msg.fail("昵称不规范");
        return;
      }
      this.model.name = this.content;
      this.$request({
        methods: "post",
        url: "/api/user/modify/" + localStorage.getItem("id"),
        data: { name: this.model.name },
        type: "json",
      }).then((res) => {
        this.$msg.success(res.data["msg"]);
        this.content = "";
      });
    },

    // 个性签名修改
    introClick() {
      this.model.introduction = this.content;
      this.$request({
        methods: "post",
        url: "/api/user/modify/" + localStorage.getItem("id"),
        data: { introduction: this.model.introduction },
        type: "json",
      }).then((res) => {
        this.$msg.success(res.data["msg"]);
        this.content = "";
      });
    },

    // 性别修改
    sexSelect(data) {
      if (data.name != "男" && data.name != "女") {
        return;
      }
      this.model.sex = data.name;
      this.$request({
        methods: "post",
        url: "/api/user/modify/" + localStorage.getItem("id"),
        data: { sex: this.model.sex },
        type: "json",
      });
    },
  },
  created() {
    // 每次创建调用方法获取用户参数渲染页面
    this.UserinfoData();
  },
};
</script>

<style lang="less" scoped>
.usermodify {
  .modifybox {
    margin-top: 13vh;
    .uploadfile {
      position: relative;
      overflow: hidden;
      .uploadimg {
        position: absolute;
        opacity: 0;
      }
    }
  }
}
a:link {
  color: gray;
}
</style>
