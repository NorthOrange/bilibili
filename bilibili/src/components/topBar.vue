<!-- 顶栏组件 -->
<template>
  <div class="topbar">
    <div
      class="logo"
      @click="$router.push('/')"
    >
      <img
        src="../assets/logo.png"
        alt=""
      />
    </div>
    <!-- <van-search
      background="#f6869c"
      disabled
      placeholder="请输入搜索关键词"
    /> -->
    <div style="background-color:#f6869c"></div>
    <div
      @click="goUserInfo"
      v-if="avatar_url"
    >
      <img
        :src="avatar_url"
        alt=""
      />
      <p>个人中心</p>
    </div>
    <div
      style="padding: 3vw 3vw 0 0"
      @click="$router.push('/login')"
      v-else
    >
      登陆
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      avatar_url: "",
    };
  },
  mounted() {
    if (localStorage.getItem("id")) {
      this.$request({
        methods: "get",
        url: "/api/user/info/" + localStorage.getItem("id"),
      }).then((res) => {
        this.avatar_url = res.data.avatar;
      });
    } else {
      this.avatar_url = "";
    }
  },
  methods: {
    goUserInfo() {
      this.$router.push("/user/info/" + localStorage.getItem("id"));
    },
  },
};
</script>

<style lang="less" scoped>
.topbar {
  position: sticky;
  top: 0;
  z-index: 999;
  height: 12vw;
  background-color: #f4f4f4;
  display: flex;
  flex: 1;
  .logo {
    display: flex;
    justify-self: center;
    align-items: center;
    width: 24vw;
    img {
      height: 100%;
      width: 100%;
    }
  }
  div:nth-child(2) {
    flex: 1;
  }
  div:nth-child(3) {
    display: flex;
    justify-content: center;
    background-color: #f6869c;
    color: white;
    img {
      width: 12vw;
      height: 12vw;
      border-radius: 50%;
    }
    p {
      padding: 0 10px;
    }
  }
}
</style>
