import Vue from "vue"
import router from '@/router'
import _ from "lodash"

const user = JSON.parse(localStorage.getItem("currentUser"))

const Authentication = {
  namespaced: true,
  state: {
    loading: false,
    error: null,
    authenticated: user != null,
    user
  },

  actions: {
    async register({ dispatch, commit }, { name, email, password }) {
      commit("setLoading", { name, email })
      try {
        const response = await Vue.axios.post("/users", { name, email, password })
        const { accessToken, refreshToken } = response.data.data
        localStorage.setItem('accessToken', accessToken)
        localStorage.setItem('refreshToken', refreshToken)
        commit('setSuccess', { accessToken, refreshToken })
        router.push('/')
      } catch (error) {
        const { errors } = error.response.data
        let message = errors.map(e => e.message).join(", ")
        if (_.isEmpty(message)) {
          message = error.message
        }
        commit("setError", message)
        dispatch("alert/error", message, { root: true })
      }
    }
  },

  mutations: {
    setLoading(state, user) {
      state.authenticated = false
      state.user = user
      state.loading = true
      state.error = null
    },
    setSuccess(state, user) {
      state.authenticated = true
      state.user = user
      state.loading = false
      state.error = null
    },
    setError(state, error) {
      state.authenticated = false
      state.loading = false
      state.error = error
    }
  }
}

export default Authentication
