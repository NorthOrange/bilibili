<!-- 首页组件 -->
<template>
  <div class="home">
    <top-bar></top-bar>
    <van-tabs v-model="active">
      <van-tab title="关注"> </van-tab>
      <van-tab title="推荐">
        <van-pull-refresh
          v-model="refreshing"
          @refresh="onRefresh"
        >
          <van-list
            v-model="loading"
            @load="onLoad"
          >
            <Video
              v-for="(item, index) in videoList"
              :key="index"
              @play="play($event)"
            ></Video>
          </van-list>
        </van-pull-refresh>
      </van-tab>
      <van-tab title="动漫">动漫</van-tab>
    </van-tabs>
    <div
      id="refreshButton"
      @click="refreshButtonClick"
    >
      <i :class="refreshButtonClass"></i>
    </div>
  </div>
</template>

<script>
import topBar from "../components/topBar.vue";
import Video from "../components/Video";
export default {
  components: { topBar, Video },
  data() {
    return {
      active: 1,
      videoList: [1, 2, 3],
      loading: false,
      refreshing: false,
      refreshButtonClass: "fa fa-refresh fa-2x fa-fw", // 控制刷新按钮的样式
      onPlayVideo: "", // 正在播放的视频对象
    };
  },
  methods: {
    onRefresh() {
      this.videoList = [];
      setTimeout(() => {
        this.videoList.push(this.videoList.length + 1);
        this.videoList.push(this.videoList.length + 1);
        this.videoList.push(this.videoList.length + 1);
        this.refreshing = false;
      }, 1000);
    },
    onLoad() {
      setTimeout(() => {
        this.videoList.push(this.videoList.length + 1);
        this.videoList.push(this.videoList.length + 1);
        this.videoList.push(this.videoList.length + 1);
        this.loading = false;
      }, 1000);
    },
    refreshButtonClick() {
      this.refreshButtonClass = "fa fa-refresh fa-spin fa-2x fa-fw";
      this.onRefresh();
      setTimeout(() => {
        this.refreshButtonClass = "fa fa-refresh fa-2x fa-fw";
      }, 1000);
    },
    play(event) {
      // 视频播放时, 暂停其它视频
      console.log(
        // event.target.getBoundingClientRect(),
        pageYOffset,
        event.target.parentNode.parentNode.getBoundingClientRect()
      );
      if (this.onPlayVideo == "") {
        this.onPlayVideo = event.target;
        return;
      } else if (event.target == this.onPlayVideo) {
        return;
      } else {
        this.onPlayVideo.pause();
        this.onPlayVideo = event.target;
      }
    },
    scrolling() {
      if (this.onPlayVideo == "") {
        return;
      } else {
        var p = this.onPlayVideo.parentNode.parentNode.getBoundingClientRect(); // 获取这个元素的位置
        if (p.top < -1600 || p.bottom > 1600) {
          // 当元素超出可视范围停止播放
          this.onPlayVideo.pause();
        }
      }
    },
  },
  mounted() {
    window.addEventListener("scroll", this.scrolling); // 监听页面滚动
  },
  destroyed() {
    window.removeEventListener("scroll", this.scrolling);
  },
};
</script>

<style lang="less" scoped>
#refreshButton {
  position: fixed;
  right: 4vw;
  bottom: 10vh;
  color: #f6869c;
  z-index: 99;
}
</style>