import Vue from "vue"
import router from "@/router"
import _ from "lodash"
import decodeJwt from "jwt-decode"

const currentUser = JSON.parse(localStorage.getItem("currentUser"))

const Authentication = {
  namespaced: true,
  state: {
    loading: false,
    authenticated: currentUser != null,
    user: currentUser
  },

  actions: {
    async register({ dispatch, commit }, { name, email, password }) {
      commit("setLoading")
      try {
        const response = await Vue.axios.post("/users", { name, email, password })
        const { accessToken, refreshToken } = response.data.data
        const claims = (({ id, name, email }) => ({ id, name, email }))(decodeJwt(accessToken))
        const currentUser = { ...claims, accessToken, refreshToken }
        localStorage.setItem("currentUser", JSON.stringify(currentUser))
        commit("setSuccess", currentUser)
        router.push("/profile")
      } catch (error) {
        const { errors } = error.response.data
        let message = errors.map(e => e.message).join(", ")
        if (_.isEmpty(message)) {
          message = error.message
        }
        dispatch("alert/error", message, { root: true })
      }
    },

    async login({ dispatch, commit }, { email, password }) {
      commit("setLoading")
      try {
        const response = await Vue.axios.post("/tokens", { email, password })
        const { accessToken, refreshToken } = response.data.data
        const claims = (({ id, name, email }) => ({ id, name, email }))(decodeJwt(accessToken))
        const currentUser = { ...claims, accessToken, refreshToken }
        localStorage.setItem("currentUser", JSON.stringify(currentUser))
        commit("setSuccess", currentUser)
        router.push("/profile")
      } catch (error) {
        const { errors } = error.response.data
        let message = errors.map(e => e.message).join(", ")
        if (_.isEmpty(message)) {
          message = error.message
        }
        dispatch("alert/error", message, { root: true })
      }
    },

    async logout({ commit }) {
      commit("setLoading")
      localStorage.removeItem("currentUser")
      commit("setLogout")
      router.push("/")
    }
  },

  mutations: {
    setLoading(state) {
      state.loading = true
    },
    setSuccess(state, user) {
      state.authenticated = true
      state.user = user
      state.loading = false
    },
    setLogout(state) {
      state.authenticated = false
      state.user = null
      state.loading = false
    }
  }
}

export default Authentication
