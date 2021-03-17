import Vue from "vue";
import VueRouter from "vue-router";

import Vant from "vant";
import "vant/lib/index.css";
Vue.use(Vant);
Vue.use(VueRouter);

import Login from "../views/Login.vue";
import Register from "../views/Register.vue";
import userInfo from "../views/userInfo.vue";
import Home from "../views/Home.vue";
import usermodify from "../views/userModify.vue";
import videoUpload from "../views/videoUpload";

const routes = [
  { path: "/login", component: Login },
  { path: "/register", component: Register },
  { path: "/user/info/:id", component: userInfo },
  { path: "/", component: Home },
  { path: "/user/modify", component: usermodify },
  { path: "/video/upload", component: videoUpload },
];

const router = new VueRouter({
  routes,
});

export default router;
