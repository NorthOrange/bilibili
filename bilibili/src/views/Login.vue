<!-- 登录组件 -->
<template>
  <div class="login">
    <top-bar></top-bar>
    <div class="login_title">
      <div></div>
      <div>登录</div>
      <div @click="$router.push('/register')">切换到注册</div>
    </div>
    <div class="login_box">
      <inputField
        label="帐号"
        placeholder="请输入6-11纯位数字"
        type="text"
        maxlength="11"
        @filedInput="(res) => (model.account = res)"
      >
      </inputField>
      <inputField
        label="密码"
        placeholder="请输入6-11位，可含有数字字母_-）"
        type="password"
        maxlength="11"
        @filedInput="(res) => (model.password = res)"
      >
      </inputField>
      <confirm-button
        btnText="登录"
        @confirmClick="confirmLogin"
      ></confirm-button>
      <confirm-button
        btnText="返回首页"
        @confirmClick="$router.push('/')"
      ></confirm-button>
    </div>
  </div>
</template>

<script>
import confirmButton from "../components/confirmButton.vue";
import inputField from "../components/inputField.vue";
import topBar from "../components/topBar";
export default {
  components: { inputField, confirmButton, topBar },
  data() {
    return {
      model: {
        account: "",
        password: "",
      },
    };
  },
  methods: {
    confirmLogin() {
      // 简单正则校验用户输入
      if (!/^[0-9]{6,11}$/.test(this.model.account)) {
        this.$msg.fail("帐号不规范");
      } else if (!/^[a-zA-Z0-9-_]{6,11}$/.test(this.model.password)) {
        this.$msg.fail("密码不规范");
      } else {
        // 弹出加载框等待服务端回应
        var loading = this.$msg.loading({
          durations: 5,
          msg: "加载中...",
          forbidClick: true,
        });
        this.$request({
          methods: "post",
          url: "/api/user/login",
          data: {
            account: this.model.account,
            password: this.model.password,
          },
          type: "json",
        })
          .then((res) => {
            if (res.status == 200) {
              loading.clear();
              this.$msg.success(res.data["msg"]);
              localStorage.setItem("id", res.data["id"]);
              localStorage.setItem("token", res.data["token"]);
              setTimeout(() => {
                this.$router.push("/user/info/" + res.data["id"]);
              }, 1000);
            }
          })
          .catch((err) => {
            loading.clear();
            this.$msg.fail(err.response.data["msg"]);
          });
      }
    },
  },
};
</script>

<style lang="less" scoped>
.login {
  .login_title {
    margin: 18vh 0 0 0;
    height: 8vh;
    background-color: white;
    display: flex;
    div {
      flex: 1;
      display: flex;
      justify-content: center;
      align-items: center;
      line-height: 8h;
      font-size: 6.444vw;
    }
    :nth-child(3) {
      font-size: 4vw;
      font-weight: 100;
    }
  }
}
</style>
