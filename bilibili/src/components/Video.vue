<!-- 视频播放器组件 -->
<template>
  <div class="video">

    <div
      class="user"
      @click="goUserInfo"
    >
      <img
        :src="model.avatar"
        alt=""
      />
      <span class="userinfo">
        <span>{{ model.username }}</span>
        <span>{{ model.userintroduction }}</span>
      </span>
    </div>

    <div class="player">
      <span>{{ model.name }}</span>
      <video
        controls
        @play="play"
        :src="model.video"
        :poster="model.cover"
      ></video>
      <span class="videoLike">
        <span
          id="like"
          style="color: #fb7299;"
          v-show="videoLikeStatus==1"
          @click="likeClick"
        >
          <i
            class="fa fa-thumbs-up"
            aria-hidden="true"
          ></i>
          {{ model.likes + isLike }}
        </span>
        <span
          id="like"
          style="color: gray;"
          v-show="videoLikeStatus==0||videoLikeStatus==-1"
          @click="likeClick"
        >
          <i
            class="fa fa-thumbs-o-up"
            aria-hidden="true"
          ></i>
          {{ model.likes + isLike }}
        </span>
        <span
          id="dislike"
          style="color: #fb7299;"
          v-show="videoLikeStatus==-1"
          @click="dislikeClick"
        >
          <i
            class="fa fa-thumbs-down"
            aria-hidden="true"
          ></i>
          {{ model.dislike + isDislike }}
        </span>
        <span
          id="dislike"
          style="color: gray;"
          v-show="videoLikeStatus==1||videoLikeStatus==0"
          @click="dislikeClick"
        >
          <i
            class="fa fa-thumbs-o-down"
            aria-hidden="true"
          ></i>
          {{ model.dislike  + isDislike}}
        </span>
        <span>
          <i
            class="fa fa-commenting-o"
            aria-hidden="true"
          ></i>
          {{ model.comments }}
        </span>
      </span>

    </div>
  </div>
</template>

<script>
import "font-awesome/css/font-awesome.min.css";
export default {
  props: ["videoid"],
  data() {
    return {
      model: {},
      videoLikeStatus: 0,
      originVideoLikeStatus: 0,
      isLike: 0,
      isDislike: 0,
      timer: 0,
    };
  },
  created() {
    // 获取视频信息
    var req_url;
    if (this.videoid != 0) {
      req_url = "/api/video/" + this.videoid;
    } else {
      req_url = "/api/video/get";
    }
    this.$request({
      methods: "get",
      url: req_url,
    }).then((response) => {
      this.model = response.data;
      this.$request({
        // 根据视频信息, 获取用户和视频的关系
        methods: "post",
        url: "/api/live/like/query",
        data: {
          videoid: parseInt(this.model.videoid),
          userid: parseInt(localStorage.getItem("id")),
        },
        type: "json",
      }).then((response) => {
        this.videoLikeStatus = response.data.likestatus;
        this.OriginVideoLikeStatus = response.data.likestatus;
      });
    });
  },

  methods: {
    goUserInfo() {
      this.$router.push("/user/info/" + this.model.fromid);
    },
    play(event) {
      // 点击播放视频, 将对象传出去
      this.$emit("play", event);
    },
    likeClick() {
      console.log("likeClick");
      if (this.videoLikeStatus == 1) {
        this.videoLikeStatus = 0;
        this.isLike = 0;
      } else {
        this.videoLikeStatus = 1;
        this.isLike = 1;
        this.isDislike = 0;
      }
      if (this.timer != 0) {
        clearTimeout(this.timer); // 拦截重复点击
      }

      this.timer = setTimeout(() => {
        if (this.videoLikeStatus != this.OriginVideoLikeStatus) {
          // 状态发生变化
          this.$request({
            methods: "post",
            url: "/api/live/like",
            data: {
              videoid: parseInt(this.model.videoid),
              userid: parseInt(localStorage.getItem("id")),
              likestatus: parseInt(this.videoLikeStatus),
            },
            type: "json",
          }).then((res) => {
            this.originVideoLikeStatus = res.data.videolikestatus;
            this.timer = 0;
          });
        } else {
          this.timer = 0;
          return;
        }
      }, 1000);
    },
    dislikeClick() {
      console.log("dislikeClick");
      if (this.videoLikeStatus == -1) {
        this.videoLikeStatus = 0;
        this.isDislike = 0;
      } else {
        this.videoLikeStatus = -1;
        this.isDislike = 1;
        this.isLike = 0;
      }
      if (this.timer != 0) {
        clearTimeout(this.timer); // 拦截重复点击
      }

      this.timer = setTimeout(() => {
        if (this.videoLikeStatus != this.OriginVideoLikeStatus) {
          // 状态发生变化
          this.$request({
            methods: "post",
            url: "/api/live/like",
            data: {
              videoid: parseInt(this.model.videoid),
              userid: parseInt(localStorage.getItem("id")),
              likestatus: parseInt(this.videoLikeStatus),
            },
            type: "json",
          }).then((response) => {
            this.originVideoLikeStatus = response.data.videolikestatus;
            this.timer = 0;
          });
        } else {
          this.timer = 0;
          return;
        }
      }, 1000);
    },
  },
};
</script>

<style lang="less" scoped>
.video {
  border-bottom: 2vw solid #ddd;
  .user {
    background-color: white;
    display: flex;
    justify-content: left;
    align-items: center;
    img {
      width: 12vw;
      height: 12vw;
      border-radius: 50%;
    }
    .userinfo {
      display: flex;
      flex-direction: column;
      padding-left: 2vw;
      :nth-child(1) {
        color: #ff85ad !important;
        font-weight: 600;
        font-size: 4vw;
      }
      :nth-child(2) {
        color: gray;
      }
    }
  }

  .videoLike {
    display: flex;
    align-items: center;
    height: 12vw;
    background-color: white;
    span {
      padding-left: 3vw;
      color: gray;
    }
  }

  .player {
    span:nth-child(1) {
      font-size: 4vw;
      background-color: white;
    }
    display: flex;
    flex-direction: column;
    width: 100%;
    video {
      width: 100%;
    }
  }
}
</style>
