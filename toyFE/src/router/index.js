import { createRouter, createWebHistory } from 'vue-router'
import middleware from './middleware'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: () => import('../views/HomeView.vue'),
    },
    {
      path: '/register',
      component: () => import('../views/RegisterView.vue'),
    },
    {
      path: '/login',
      component: () => import('../views/LoginView.vue'),
    },
    {
      path: '/product/:nameProduct/:idProduct',
      component: () => import('../views/PayView.vue'),
    },
  ],
})

new middleware(router)

export default router
