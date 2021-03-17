<!-- 编辑栏组件 -->
<template>
  <div @click="editbarClick">
    <div class="editbar">
      <span id="leftSpan">{{ left }}</span>
      <span>
        <slot name="right"></slot>
      </span>
    </div>
  </div>
</template>

<script>
export default {
  props: ["left"],
  methods: {
    editbarClick() {
      this.$emit("editbarClick");
    },
  },
};
</script>

<style lang="less" scoped>
.editbar {
  background-color: white;
  padding: 4.4vw;
  border-bottom: 1px solid #fafafa;
  display: flex;
  justify-content: space-between;
  #leftSpan {
    flex-shrink: 0;
    padding-right: 26vw;
  }
}
</style>
