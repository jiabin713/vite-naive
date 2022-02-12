import PageLayout from '@/layouts/PageLayout';
import { createRouter, createWebHistory } from 'vue-router';

const router = createRouter({
  history: createWebHistory(),
  routes: [
    // {
    //   path: '/',
    //   redirect: 'login',
    // },
    // Login,
    {
      name: 'root',
      path: '/',
      component: PageLayout,
      //   children: appRoutes,
    },
  ],
  scrollBehavior() {
    return { top: 0 };
  },
});

export default router;
