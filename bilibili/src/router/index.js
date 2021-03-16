import Vue from "vue";
import VueRouter from "vue-router";

import Vant from "vant";
import "vant/lib/index.css";
Vue.use(Vant);
Vue.use(VueRouter);

import Login from "../views/Login.vue";
import Register from "../views/Register.vue";

const routes = [
  { path: "/login", component: Login },
  { path: "/register", component: Register },
];

const router = new VueRouter({
  routes,
});

export default router;
