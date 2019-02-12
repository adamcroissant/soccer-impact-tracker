import Vue from 'vue'
import App from './App.vue'
import axios from 'axios'


axios.defaults.baseURL = 'https://us-central1-larealfc.cloudfunctions.net/'


Vue.config.productionTip = false

new Vue({
  render: h => h(App),
}).$mount('#app')
