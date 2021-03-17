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
        <span id="like">
          <i
            class="fa fa-thumbs-o-up"
            aria-hidden="true"
          ></i>
          {{ model.likes }}
        </span>
        <span id="dislike">
          <i
            class="fa fa-thumbs-o-down"
            aria-hidden="true"
          ></i>
          {{ model.dislike }}
        </span>
      </span>
    </div>
  </div>
</template>

<script>
import "font-awesome/css/font-awesome.min.css";
export default {
  data() {
    return {
      model: {},
    };
  },
  created() {
    this.$request({
      methods: "get",
      url: "/api/video/get",
    }).then((response) => {
      this.model = response.data;
    });
  },
  beforeMount() {
    //     this.$request(
    //       "post",
    //       "/api/relation/video/retrieve",
    //       {
    //         videoid: this.model.videoid,
    //         userid: localStorage.getItem("id"),
    //       },
    //       "json"
    //     ).then((res) => {
    //       if (res.data.relation == 0) {
    //         this.isLike = false;
    //         this.isDislike = true;
    //         this.remoteRelation = 0;
    //       } else {
    //         this.isLike = true;
    //         this.isDislike = false;
    //         this.remoteRelation = 1;
    //       }
    //     });
    //   },
    //   watch: {
    //     isLike() {
    //       // 监听属性, 改变按钮颜色
    //       const likeDom = document.getElementById("like");
    //       const dislikeDom = document.getElementById("dislike");
    //       if (this.isLike) {
    //         likeDom.style.color = " red";
    //         dislikeDom.style.color = "gray";
    //         this.model.likes += 1;
    //       } else {
    //         likeDom.style.color = "gray";
    //         this.model.likes -= 1;
    //       }
    //     },
    //     isDislike() {
    //       // 监听属性, 改变按钮颜色
    //       const likeDom = document.getElementById("like");
    //       const dislikeDom = document.getElementById("dislike");
    //       if (this.isDislike) {
    //         dislikeDom.style.color = "red";
    //         likeDom.style.color = "gray";
    //         this.model.dislike += 1;
    //       } else {
    //         dislikeDom.style.color = "gray";
    //         this.model.dislike -= 1;
    //       }
    //     },
  },
  methods: {
    goUserInfo() {
      this.$router.push("/user/info/" + this.model.fromid);
    },
    play(event) {
      // 点击播放视频, 将对象传出去
      this.$emit("play", event);
    },
    // likeClick() {
    //   this.isLike = !this.isLike;
    // },
    // dislikeClick() {
    //   this.isDislike = !this.isDislike;
    // },
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
