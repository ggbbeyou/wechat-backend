import { createRouter, createWebHashHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const routes = [
  {
    path: '/friendlist',
    name: 'home',
    component: HomeView,
    // redirect: '/chatlist',
    children: [
      {
        path: '/userinfo',
        component: () => import('../components/users/UserInfo.vue')
      },
      {
        path: '/chatlist',
        redirect: '/chatf',
        component: () => import('../components/ChatView.vue'),
        children: [
          {
            path: '/chatf',
            component: () => import('../components/chat/DefaultView.vue'),
          },
          {
            path: '/chatinfo',
            component: () => import('../components/chat/ChatInfoView.vue'),
          },
        ]
      },
      {
        path: '/friendlist',
        component: () => import('../components/FriendsView.vue'),
        redirect: '/friendf',
        children: [
          {
            path: '/friendf',
            component: () => import('../components/friends/DefaultView.vue'),
          },
          {
            path: '/friendinfo/:uid',
            component: () => import('../components/friends/FriendInfoView.vue'),
          },
        ]
      },
      {
        path: '/grouplist',
        component: () => import('../components/GroupsView.vue'),
        redirect: '/groupdf',
        children: [
          {
            path: '/groupdf',
            component: () => import('../components/groups/DefaultView.vue'),
          },
          {
            path: '/groupinfo/:gid',
            component: () => import('../components/groups/GroupInfoView.vue'),
          },
        ]
      },
      {
        path: '/talklist',
        component: () => import('../components/TalkView.vue')
      },
    ],
  },
  {
    path: '/',
    redirect: '/login',
  },
  {
    path: '/login',
    component: () => import('../views/LoginView.vue')
  },
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
