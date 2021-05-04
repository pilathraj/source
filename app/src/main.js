import Vue from 'vue'
import App from './App.vue'
import Buefy from 'buefy'
import axios from 'axios'
import VueAxios from 'vue-axios'

import 'buefy/dist/buefy.css'
import router from './router'

Vue.use(Buefy)
Vue.use(VueAxios, axios)

Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
