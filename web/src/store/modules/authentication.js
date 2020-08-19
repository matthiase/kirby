import UserService from "@/service/userService"

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
        const currentUser = await UserService.register(name, email, password)
        localStorage.setItem("currentUser", JSON.stringify(currentUser))
        commit("setSuccess", currentUser)
        return currentUser
      } catch (error) {
        dispatch("alert/error", error, { root: true })
      }
    },

    async login({ dispatch, commit }, { email, password }) {
      commit("setLoading")
      try {
        const currentUser = await UserService.login(email, password)
        localStorage.setItem("currentUser", JSON.stringify(currentUser))
        commit("setSuccess", currentUser)
        return currentUser
      } catch (error) {
        dispatch("alert/error", error, { root: true })
      }
    },

    logout({ commit }) {
      commit("setLoading")
      localStorage.removeItem("currentUser")
      commit("setLogout")
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
