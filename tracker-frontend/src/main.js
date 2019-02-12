import Vue from 'vue'
import App from './App.vue'
import axios from 'axios'

axios.defaults.baseURL = 'http://localhost:8081/api/v1'


Vue.config.productionTip = false

new Vue({
  render: h => h(App),
}).$mount('#app')
