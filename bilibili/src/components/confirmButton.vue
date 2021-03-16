<template>
  <div class="btnparent">
    <div class="confirmBtn" @click="confirmClick">{{ btnText }}</div>
  </div>
</template>

<script>
export default {
  props: ["btnText"],
  methods: {
    confirmClick() {
      this.$emit("confirmClick");
    },
  },
};
</script>

<style lang="less" scoped>
.btnparent {
  padding: 2.7vh 2.7vw;
}
.confirmBtn {
  height: 7vh;
  background-color: #ff9db5;
  color: #f4f4f4;
  display: flex;
  justify-content: center;
  align-items: center;
  border-radius: 6.25vw;
}
</style>
