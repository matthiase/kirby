import Vue from 'vue'
import './plugins/axios'
import Buefy from 'buefy'
import '@mdi/font/css/materialdesignicons.css'

import App from './App'
import router from './router'
import store from './store'
import './assets/styles/index.scss'

Vue.config.productionTip = false

Vue.use(Buefy)

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
