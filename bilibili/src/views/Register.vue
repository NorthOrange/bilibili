<!-- 注册组件 -->
<template>
  <div class="register">
    <top-bar></top-bar>
    <div class="register_title">
      <div></div>
      <div>注册</div>
      <div @click="$router.push('/login')">切换到登录</div>
    </div>
    <div class="login_box">
      <inputField
        label="昵称"
        placeholder="请输入2~7位非空字符）"
        type="text"
        maxlength="7"
        @filedInput="res => (model.name = res)"
      ></inputField>
      <inputField
        label="手机号"
        placeholder="请输入您的手机号"
        type="text"
        maxlength="11"
        @filedInput="res => (model.mobile = res)"
      >
      </inputField>
      <van-field
        v-model="smsCode"
        center
        clearable
        label="短信验证码"
        placeholder="请输入短信验证码"
      >
        <template #button>
          <van-button
            size="small"
            color="#ff9db5"
            type="primary"
            @click="sendSms"
            :text="sms_interval"
          ></van-button>
        </template>
      </van-field>
      <inputField
        label="密码"
        placeholder="请输入6-11位，可含有数字字母_-）"
        type="password"
        maxlength="11"
        @filedInput="res => (model.password = res)"
      >
      </inputField>
      <confirm-button
        btnText="注册"
        @confirmClick="confirmRegister"
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
        name: "",
        mobile: "",
        password: "",
      },
      smsCode: "",
      sms_interval: "发送验证码",
    };
  },
  methods: {
    confirmRegister() {
      // 简单正则校验用户输入
      if (!/^[^ ]{2,7}$/.test(this.model.name)) {
        this.$msg.fail("请输入规范的用户名");
      } else if (!/^1[0-9]{10}$/.test(this.model.mobile)) {
        this.$msg.fail("请输入正确的手机号");
      } else if (!/^[a-zA-Z0-9_-]{6,11}$/.test(this.model.password)) {
        this.$msg.fail("请输入规范的密码");
      } else {
        // 弹出加载框等待服务端回应
        var loading = this.$msg.loading({
          durations: 5,
          msg: "加载中...",
          forbidClick: true,
        });
        this.$request({
          methods: "post",
          url: "/api/user/register",
          data: {
            name: this.model.name,
            mobile: this.model.mobile,
            password: this.model.password,
            smsCode: this.smsCode,
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
            } else {
              this.$msg.fail(res.data["msg"]);
            }
          })
          .catch((err) => {
            loading.clear();
            this.$msg.fail(err.response.data["msg"]);
          });
      }
    },
    sendSms() {
      if (!/^1[0-9]{10}$/.test(this.model.mobile)) {
        this.$msg.fail("请输入正确的手机号");
      } else if (this.sms_interval != "发送验证码") {
        this.$msg.fail("该功能正在冷却中");
      } else {
        this.sms_interval = 60;
        var smsIn;
        smsIn = setInterval(() => {
          this.sms_interval -= 1;
          if (this.sms_interval == 0) {
            this.sms_interval = "发送验证码";
            clearInterval(smsIn);
          }
        }, 1000);
        this.$request({
          methods: "post",
          url: "/api/snedSMS",
          data: {
            mobile: this.model.mobile,
          },
          type: "json",
        })
          .then((res) => {
            if (res.status == 200) {
              this.$msg.success(res.data["msg"]);
            } else {
              this.$msg.fail(res.data["msg"]);
            }
          })
          .catch((err) => {
            this.$msg.fail(err.response.data["msg"]);
          });
      }
    },
  },
};
</script>

<style lang="less" scoped>
.register {
  .register_title {
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
