<!-- 输入框组件 -->
<template>
  <div>
    <van-field
      :label="label"
      :type="type"
      :placeholder="placeholder"
      :maxlength="maxlength"
      v-model="content"
      autofocus
    ></van-field>
  </div>
</template>

<script>
export default {
  data() {
    return {
      content: "",
    };
  },
  props: ["label", "type", "placeholder", "maxlength"],
  methods: {},
  watch: {
    content() {
      this.$emit("filedInput", this.content);
    },
  },
};
</script>
