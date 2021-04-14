import Vue from "vue";
import App from "./App.vue";
import router from "./router";

Vue.config.productionTip = false;

// http 请求封装
import request from "./utils/request";
Vue.prototype.$request = request;

// 移动端组件库
import Vant from "vant";
import "vant/lib/index.css";
Vue.use(Vant);

// 一个轻提示组件
import { Toast } from "vant";
Vue.prototype.$msg = Toast;

import { Dialog } from "vant";
Vue.prototype.$Dialog = Dialog;

new Vue({
  router,
  render: (h) => h(App),
}).$mount("#app");
