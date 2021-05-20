<template>
  <div class="videoplayer">
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
};
</script>

<style>
</style>