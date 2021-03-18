<!-- 用户中心页面 -->
<template>
  <div class="userinfo">
    <top-bar></top-bar>
    <div id="userinfoBackimg">
    </div>

    <user-detail :userinfo="model"></user-detail>
    <van-tabs
      v-model="active"
      background="#f4f4f4"
      class="tabs"
    >
      <van-tab title="视频">
        <div class="videolist">
          <div
            v-for="(item, index) in videoNameList"
            :key="index"
          >
            <a>
              <div class="cover">
                <img
                  :src="videoCoverList[index]"
                  alt=""
                />
              </div>
              <div class="info">
                <h3>{{ videoNameList[index] }}</h3>
              </div>
            </a>
          </div>
        </div>
      </van-tab>
      <van-tab title="动态">动态</van-tab>
    </van-tabs>

    <div class="bottomButton">
      <confirmBtn
        btnText="发布视频"
        @confirmClick="$router.push('/video/upload')"
      ></confirmBtn>
    </div>
  </div>
</template> 

<script>
import topBar from "../components/topBar.vue";
import UserDetail from "../components/userDetail.vue";
import confirmBtn from "../components/confirmButton";

export default {
  data() {
    return {
      model: {},
      active: 0,
      videoNameList: [],
      videoCoverList: [],
      videoIdList: [],
    };
  },
  components: { UserDetail, topBar, confirmBtn },
  methods: {
    UserinfoData() {
      console.log("跳转到个人中心");
      this.$request({
        methods: "get",
        url: "/api/user/info/" + this.$route.params.id,
      }).then((res) => {
        this.model = res.data;
      });
    },
    UserVideoList() {
      this.$request({
        methods: "get",
        url: "/api/user/video/list/" + this.$route.params.id,
      }).then((res) => {
        this.videoNameList = res.data.namelist;
        this.videoCoverList = res.data.coverlist;
        this.videoIdList = res.data.idlist;
      });
    },
  },
  created() {
    this.UserinfoData();
    this.UserVideoList();
  },
  watch: {
    $route() {
      this.UserinfoData();
      this.UserVideoList();
    },
  },
};
</script>

<style lang="less" scoped>
#userinfoBackimg {
  width: 100%;
  height: 20vh;
  overflow: hidden;
  background: url("../assets/img/userinfoBackimg.jpg") no-repeat 0 -10vh;
  background-size: cover;
  margin-bottom: 3vh;
}
.bottomButton {
  position: fixed;
  right: 20vw;
  bottom: 10vh;
  height: 3vh;
  width: 60vw;
}
.videolist:nth-child(1) {
  margin-top: 5vw;
}
.videolist {
  height: 70vw;
  overflow: scroll;
  padding-left: 3.2vw;
  > div {
    padding-top: 2vw;
    border-top: 1px solid #ccc;
    border-bottom: 1px solid #ccc;
  }
  .cover {
    float: left;
    width: 32vw;
    height: 20vw;
    position: relative;
    img {
      width: 100%;
      height: 100%;
    }
  }
  .info {
    margin-left: 34vw;
    height: 20vw;
    h3 {
      line-height: 4.5vw;
      font-size: 3.7vw;
      color: #212121;
      font-weight: 400;
      overflow: hidden;
    }
  }
}
</style>
