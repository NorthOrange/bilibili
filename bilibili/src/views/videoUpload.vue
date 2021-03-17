<!-- 视频上传组件 -->
<template>
  <div class="videoupload">
    <top-bar></top-bar>
    <div class="uploadbox">
      <!-- 上传视频 -->
      <div class="uploadfile">
        <van-uploader
          class="uploadimg"
          accept="video/*"
          preview-size="100vw"
          :after-read="afterReadVideo"
        />
        <edit-bar left="上传视频">
          <a
            href="javascript:;"
            slot="right"
          >{{ VideoStatus }}</a>
        </edit-bar>
      </div>

      <!-- 视频标题 -->
      <edit-bar
        left="视频标题"
        @editbarClick="nameShow = true"
      >
        <a
          href="javascript:;"
          slot="right"
        >{{ VideoName }}</a>
      </edit-bar>
      <van-dialog
        v-model="nameShow"
        title="视频标题"
        confirm-button-color="#0000ee"
        show-cancel-button
        @confirm="nameClick"
        @cancel="VideoName = ''"
      >
        <van-field
          type="textarea"
          maxlength="50"
          v-model="VideoName"
          autofocus
        />
      </van-dialog>

      <!-- 视频封面上传 -->
      <div class="uploadfile">
        <van-uploader
          accept="image/*"
          class="uploadimg"
          preview-size="100vw"
          :after-read="afterReadCover"
        />
        <edit-bar left="视频封面">
          <img
            style="width: 32vw; height: 18vw"
            id="coverpreview"
            slot="right"
          />
        </edit-bar>
      </div>
      <!-- 视频简介 -->
      <edit-bar
        left="视频简介"
        @editbarClick="introShow = true"
      >
        <a
          href="javascript:;"
          slot="right"
        >{{ VideoIntro }}</a>
      </edit-bar>

      <van-dialog
        v-model="introShow"
        title="视频简介"
        confirm-button-color="#0000ee"
        show-cancel-button
        @confirm="introClick"
        @cancel="VideoIntro = ''"
      >
        <van-field
          type="textarea"
          maxlength="100"
          v-model="VideoIntro"
          autofocus
        />
      </van-dialog>
    </div>
    <div class="bottomButton">
      <confirmBtn
        btnText="确认上传"
        @confirmClick="confirmUpload"
      ></confirmBtn>
    </div>
  </div>
</template>

<script>
import topBar from "../components/topBar.vue";
import confirmBtn from "../components/confirmButton";
import editBar from "../components/editBar";
export default {
  data() {
    return {
      Video: "",
      VideoName: "",
      nameShow: false,
      VideoStatus: "",
      VideoIntro: "",
      introShow: false,
      VideoCover: "",
    };
  },
  components: { topBar, confirmBtn, editBar },
  methods: {
    // 视频上传
    afterReadVideo(file) {
      console.log("拿到一个文件: " + file.file);
      if (file.file.type != "video/mp4") {
        console.log(file.file.type);
        this.$msg.fail("请上传mp4格式文件.");
        this.VideoStatus = "上传失败!";
        return;
      }
      this.Video = file.file;
      this.VideoName = file.file.name.replace(".mp4", ""); // 根据文件名字填充视频名
      this.VideoStatus = "上传成功~";
    },
    // 视频标题
    nameClick() {
      if (this.VideoName.length > 50 || this.VideoName.length < 2) {
        this.$msg.fail("视频标题格式有误, 请输入2~50个字符.");
        return;
      }
    },
    // 视频封面
    afterReadCover(file) {
      if (
        file.file.type != "image/png" &&
        file.file.type != "image/jpg" &&
        file.file.type != "image/jpeg"
      ) {
        this.$msg.fail("请上传一张图片!");
        return;
      }
      this.VideoCover = file.file;
    },
    // 视频简介
    introClick() {
      if (this.VideoIntro.length > 100) {
        this.$msg.fail("视频简介格式有误, 请输入0~100个字符.");
        return;
      }
    },
    // 确认上传视频
    confirmUpload() {
      if (this.Video.type != "video/mp4") {
        console.log(this.Video.type);
        this.$msg.fail("请上传mp4格式文件.");
        this.VideoStatus = "上传失败!";
        return;
      }
      if (this.VideoIntro.length > 100) {
        this.$msg.fail("视频简介格式有误, 请输入0~100个字符.");
        return;
      }
      if (this.VideoName.length > 50 || this.VideoName.length < 2) {
        this.$msg.fail("视频标题格式有误, 请输入2~50个字符.");
        return;
      }
      var formdata = new FormData();
      formdata.append("video", this.Video);
      formdata.append("id", localStorage.getItem("id"));
      formdata.append("name", this.VideoName);
      formdata.append("introduction", this.VideoIntro);
      if (this.VideoCover != "") {
        formdata.append("cover", this.VideoCover);
      } else {
        formdata.append("nocover", "true");
      }
      // 上传
      var loading = this.$msg.loading({
        durations: 1000,
        msg: "加载中...",
        forbidClick: true,
      });
      this.$request({
        methods: "post",
        url: "/api/video/upload",
        data: formdata,
        type: "formdata",
      })
        .then((res) => {
          if (res.status == 200) {
            loading.clear();
            this.$msg.success(res.data["msg"]);
            setTimeout(() => {
              this.$router.push("/user/info/" + localStorage.getItem("id"));
            }, 1000);
          }
        })
        .catch((err) => {
          if (err.response.status == 400) {
            loading.clear();
            this.$msg.error(err.response.data["msg"]);
          } else {
            loading.clear();
            this.$msg.error("遇到了预期之外的错误,请检查你的网络");
          }
        });
    },
  },
  watch: {
    VideoCover() {
      // 上传封面时显示封面预览图
      var img = document.getElementById("coverpreview");
      img.classList.add("obj");
      img.file = this.VideoCover;
      var reader = new FileReader();
      this.coverPreview = reader.readas;
      reader.onload = (function (aImg) {
        return function (e) {
          aImg.src = e.target.result;
        };
      })(img);
      reader.readAsDataURL(this.VideoCover);
    },
  },
};
</script>

<style lang="less" scoped>
.videoupload {
  a:link {
    color: gray;
  }

  .uploadbox {
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
  .bottomButton {
    position: fixed;
    right: 20vw;
    bottom: 10vh;
    height: 3vh;
    width: 60vw;
  }
}
</style>