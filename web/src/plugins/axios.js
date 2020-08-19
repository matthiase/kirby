'use strict'

import Vue from 'vue'
import axios from 'axios'

// Full config: https://github.com/axios/axios#request-config
// axios.defaults.baseURL = process.env.API_BASE_URL || '';
// axios.defaults.headers.common['Authorization'] = AUTH_TOKEN;
// axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded';

let config = {
  baseURL: process.env.VUE_APP_API_BASE_URL || '',
  timeout: 30 * 1000
  // withCredentials: true, // Check cross-site Access-Control
}

const _axios = axios.create(config)

_axios.interceptors.request.use(
  config => {
    const accessToken = localStorage.getItem('jwt')
    if (accessToken) {
      config.headers = { Authorization: `Bearer ${accessToken}` }
    }
    return config
  },
  error => {
    return Promise.reject(error)
  }
)

// Add a response interceptor
_axios.interceptors.response.use(
  function(response) {
    // Do something with response data
    return response
  },
  function(error) {
    // Do something with response error
    return Promise.reject(error)
  }
)

Plugin.install = function(Vue) {
  Vue.axios = _axios
  window.axios = _axios
  Object.defineProperties(Vue.prototype, {
    axios: {
      get() {
        return _axios
      }
    },
    $axios: {
      get() {
        return _axios
      }
    }
  })
}

Vue.use(Plugin)

export default Plugin