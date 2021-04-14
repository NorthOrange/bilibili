<!-- 用户资料组件 -->
<template>
  <div class="userDetail">
    <div>
      <div class="user_img">
        <img
          :src="userinfo.avatar"
          v-if="userinfo.avatar"
        />
        <img
          src="../assets/img/avatar/defaultAvatar.jpg"
          v-else
        />
      </div>
      <div class="user_edit">
        <div>
          <p>
            <span>{{ userinfo.followers }}</span>
            <span class="user_text"> 粉丝 </span>
          </p>
          <p>
            <span>{{ userinfo.following }}</span>
            <span class="user_text"> 关注 </span>
          </p>
          <p>
            <span>{{ userinfo.likes }}</span>
            <span class="user_text"> 获赞 </span>
          </p>
        </div>
        <div
          v-if="localId == userinfo.id"
          class="button"
          @click="$router.push('/user/modify')"
        >
          编辑资料
        </div>
        <div
          v-else
          class="button"
          style="color: #fff; background-color: #fb7299"
          @click="follow"
        >
          <div v-if="isFollow==1">已关注</div>
          <div v-else>关注</div>
        </div>
      </div>
    </div>
    <div>
      <h2>{{ userinfo.name }} <span
          id="signout"
          @click="signout"
        >退出登录</span> </h2>
      <p
        class="introduction"
        v-if="userinfo.introduction"
      >
        {{ userinfo.introduction }}
      </p>
      <p
        class="introduction"
        v-else
      >这个人很懒, 什么都没有写</p>
    </div>

  </div>
</template>

<script>
export default {
  props: ["userinfo"],
  data() {
    return {
      localId: "",
      isFollow: 0,
      followStatusChange: false, // 判断关注状态是否改变
      timer: 0,
    };
  },
  methods: {
    signout() {
      this.$Dialog
        .confirm({
          title: "",
          confirmButtonColor: "#f6869c",
          message: "确定要退出登录吗?",
        })
        .then(() => {
          localStorage.clear();
          this.$router.push("/");
        })
        .catch(() => {
          return;
        });
    },
    follow() {
      // 点击关注触发一个定时函数, 确保不会频繁上传
      this.followStatusChange = !this.followStatusChange;
      if (this.isFollow == 0) {
        // 没关注就关注, 已关注就取消关注
        this.isFollow = 1;
      } else {
        this.isFollow = 0;
      }

      console.log("触发点击事件", this.isFollow);
      if (this.timer != 0) {
        // 如果1秒内已经点击过了, 清除上一次点击触发的函数
        clearTimeout(this.timer);
      }
      this.timer = setTimeout(() => {
        if (this.followStatusChange) {
          // 判断关注状态有没有变化
          this.$request({
            methods: "post",
            url: "/api/live/follow",
            data: {
              fromid: parseInt(this.localId),
              toid: parseInt(this.userinfo.id),
              followstatus: parseInt(this.isFollow),
            },
            type: "json",
          }).then(() => {
            this.timer = 0;
          });
        } else {
          this.timer = 0;
          return;
        }
      }, 1000);
    },
  },
  created() {
    this.localId = localStorage.getItem("id"); // 获取本地缓存的 id , 判断是否进入了登录用户的中心页
  },
  watch: {
    userinfo(n, o) {
      this.userinfo = n;
      console.log(o, n.id);
      if (this.localId != this.userinfo.id) {
        // 当登录用户进入其他用户页面, 请求关注状态
        this.$request({
          methods: "post",
          url: "/api/live/follow/query",
          data: {
            fromid: parseInt(this.localId),
            toid: parseInt(n.id),
          },
          type: "json",
        }).then((response) => {
          this.isFollow = response.data.followstatus;
        });
      }
    },
  },
};
</script>

<style lang="less" scoped>
.button {
  border-radius: 2vw;
}
.userDetail {
  background-color: #f4f4f4;
  padding: 0px 12px;
  > div:nth-child(1) {
    display: flex;
  }
  > div:nth-child(2) {
    h2 {
      margin: 10px 0px 2px 0px;
    }
    p {
      padding: 0px;
      margin: 0px;
    }
  }
  .user_img {
    margin-right: 20px;
    img {
      width: 26vw;
      height: 26vw;
      border-radius: 50%;
    }
  }
  .user_edit {
    flex: 1;
    div:nth-child(1) {
      display: flex;
      p {
        flex: 1;
        display: flex;
        justify-content: center;
        align-items: center;
        flex-direction: column;
        font-size: 4.4vw;
        .user_text {
          color: #aaa;
        }
      }
      p:nth-child(1),
      p:nth-child(2) {
        border-right: 1px solid #cccc;
      }
    }
    div:nth-child(2) {
      display: flex;
      border: 2px solid #fb7a9f;
      color: #fb7a9f;
      justify-content: center;
      align-items: center;
    }
  }
  .introduction {
    color: #aaa;
  }
  #signout {
    float: right;

    font-size: 4vw;
    color: #fb7a9f;
    border: 2px solid #fb7a9f;
    border-radius: 1vw;
  }
}
</style>
