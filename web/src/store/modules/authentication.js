import UserService from "@/service/userService"
import _ from "lodash"

const currentUser = JSON.parse(localStorage.getItem("currentUser"))

const Authentication = {
  namespaced: true,
  state: {
    loading: false,
    authenticated: currentUser != null,
    currentUser
  },

  actions: {
    async register({ dispatch, commit }, { name, email, password }) {
      commit("setLoading", true)
      try {
        const currentUser = await UserService.register(name, email, password)
        localStorage.setItem("currentUser", JSON.stringify(currentUser))
        commit("setCurrentUser", currentUser)
        return currentUser
      } catch (error) {
        dispatch("alert/error", error, { root: true })
      }
    },

    async login({ dispatch, commit }, { email, password }) {
      commit("setLoading", true)
      try {
        const currentUser = await UserService.login(email, password)
        localStorage.setItem("currentUser", JSON.stringify(currentUser))
        commit("setCurrentUser", currentUser)
        return currentUser
      } catch (error) {
        dispatch("alert/error", error, { root: true })
      }
    },

    logout({ commit }) {
      commit("setLoading", true)
      localStorage.removeItem("currentUser")
      commit("setCurrentUser", null)
    }
  },

  mutations: {
    setLoading(state, isLoading) {
      state.loading = isLoading
    },
    setCurrentUser(state, user) {
      state.loading = false
      state.authenticated = !_.isEmpty(user)
      state.currentUser = user
    }
  }
}

export default Authentication
