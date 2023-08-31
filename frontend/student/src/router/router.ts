import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/my-education",
      name: "ListLesson",
      component: () => import("../views/ListLessons.vue")
    },
  ]
})

export default router