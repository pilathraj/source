import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'

Vue.use(VueRouter)




const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/q1',
    name: 'Word Counter',    
    component: () => import('../components/Q1.vue')
  },
  {
    path: '/q2',
    name: 'Excel Column',    
    component: () => import('../components/Q2.vue')
  },
  {
    path: '/q3',
    name: 'CRUD',    
    component: () => import('../components/Q3.vue')
  },
  {
    path: '/q4',
    name: 'Sample Service',    
    component: () => import('../components/Q4.vue')
  },
  {
    path: '/q5',
    name: 'Goroutine',    
    component: () => import('../components/Q5.vue')
  }



]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
