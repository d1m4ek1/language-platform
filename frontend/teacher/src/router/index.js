import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/create-course',
      name: 'CreateCourse',
      component: () => import("../views/CreateCourse.vue")
    },
    {
      path: "/list-courses",
      name: "ListCourses",
      component: () => import("../views/ListCourses.vue")
    },
    {
      path: "/profile",
      name: "Profile",
      component: () => import("../views/Profile.vue")
    },
    {
      path: "/profile/settings",
      name: "SettingsProfile",
      component: () => import("../views/SettingsProfile.vue")
    }
  ]
})

export default router
