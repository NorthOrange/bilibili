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
        >
          不编辑资料
        </div>
      </div>
    </div>
    <div>
      <h2>{{ userinfo.name }}</h2>
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
    <div
      class=""
      v-if="localId==userinfo.id"
    >

    </div>

  </div>
</template>

<script>
export default {
  props: ["userinfo"],
  data() {
    return {
      localId: "",
    };
  },
  beforeMount() {
    this.localId = localStorage.getItem("id");
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
}
</style>
