import { createRouter, createWebHistory } from "vue-router";

const routers = [
  {
    "path": "/",
    "name": "Home",
    "component": () => import("../components/PortList.vue")
  },
  {
    "path": "/hello",
    "name": "HelloWorld",
    "component": () => import("../components/HelloWorld.vue")
  },
  {
    "path": "/mavendld",
    "name": "MavenDld",
    "component": () => import("../components/MavenDld.vue")
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes: routers,
});

export default router;

