import { createApp } from "vue";
import { createWebHistory, createRouter } from "vue-router";
import "@fontsource/sour-gummy/400.css";
import "@fontsource/sour-gummy/500.css";
import "@fontsource/sour-gummy/800.css";
import "./style.css";
import App from "./App.vue";
import routes from "./routes";

const router = createRouter({
  history: createWebHistory(),
  routes,
});

createApp(App).use(router).mount("#app");
