const Alert = {
  namespaced: true,
  state: {
    type: null,
    message: null
  },

  actions: {
    success({ commit }, message) {
      commit('setSuccess', message)
    },
    error({ commit }, message) {
      commit('setError', message)
    },
    clear({ commit }) {
      commit('clearError')
    }
  },

  mutations: {
    setSuccess(state, message) {
      state.type = 'success'
      state.message = message
    },
    setError(state, message) {
      state.type = 'error'
      state.message = message
    },
    clearError(state) {
      state.type = null
      state.message = null
    }
  }
}

export default Alert
