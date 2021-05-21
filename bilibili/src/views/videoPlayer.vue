<template>
  <div class="videoplayer">
    <div class="top-bar">
      <span @click="gotoHome">返回</span>
    </div>
    <Video
      :videoid="videoid"
      @play="play($event)"
    ></Video>
  </div>
</template>

<script>
import Video from "../components/Video";
export default {
  components: { Video },
  data() {
    return {
      videoid: this.$route.params.id,
    };
  },
  methods: {
    gotoHome() {
      this.$router.push("/");
    },
  },
};
</script>

<style lang="less" scoped>
</style>