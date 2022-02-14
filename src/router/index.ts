import PageLayout from '@/layouts/PageLayout';
import { createRouter, createWebHistory } from 'vue-router';

import LoginRouter from './modules/login';

const router = createRouter({
  history: createWebHistory(),
  routes: [
    // {
    //   path: '/',
    //   redirect: 'login',
    // },
    // LoginRouter,
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
