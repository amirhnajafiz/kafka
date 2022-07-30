import { createRouter, createWebHistory } from 'vue-router'

// importing views
import HomeView from '@/views/HomeView/index.vue'
import AboutView from "@/views/AboutView/index.vue";
import EducationView from "@/views/EducationView/index.vue";
import ContactView from "@/views/ContactView/index.vue";
import ProjectsView from "@/views/ProjectsView/index.vue";
import PageNotFound from "@/views/PageNotFound.vue"

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: "/about",
      name: "about",
      component: AboutView
    },
    {
      path: "/edu",
      name: "edu",
      component: EducationView
    },
    {
      path: "/contact",
      name: "contact",
      component: ContactView
    },
    {
      path: "/projects",
      name: "projects",
      component: ProjectsView
    },
    { 
      path: '/:pathMatch(.*)*', 
      component: PageNotFound 
    }
  ]
})

export default router
