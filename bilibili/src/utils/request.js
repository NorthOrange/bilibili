import axios from "axios";
import socket from "../config";
import router from "../router";

const request = axios.create({
  baseURL: socket,
  timeout: 10000,
});

export default function({ methods = "get", url, data = {}, type = "json" }) {
  if (methods == "get" || methods == "GET") {
    return request.get(url, data);
  }
  // json格式请求头
  const headerJSON = {
    "Content-Type": "application/json",
  };
  // FormData格式请求头
  const headerFormData = {
    "Content-Type": "multipart/form-data",
  };
  return request.post(url, type == "json" ? JSON.stringify(data) : data, {
    headers: type == "json" ? headerJSON : headerFormData,
  });
}

// 拦截请求, 加上 token
request.interceptors.request.use(
  function(config) {
    if (localStorage.getItem("token")) {
      config.headers["Authorization"] =
        "Bearer " + localStorage.getItem("token");
    }
    return config;
  },
  function(err) {
    return Promise.reject(err);
  }
);

// 拦截响应, 如果没有权限, 跳转到登陆页面
request.interceptors.response.use(
  function(response) {
    return response;
  },
  function(err) {
    if (err.response.status == 403) {
      router.push("/login");
    }
    return Promise.reject(err);
  }
);
