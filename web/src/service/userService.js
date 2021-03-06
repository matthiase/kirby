import Vue from "vue"
import decodeJwt from "jwt-decode"

import { handleErrorResponse } from "./serviceHelpers"

const UserService = {
  register: async (name, email, password) => {
    try {
      const response = await Vue.http.post("/users", { name, email, password })
      return handleSuccessResponse(response)
    } catch (error) {
      handleErrorResponse(error)
    }
  },

  update: async (name, email) => {
    try {
      const response = await Vue.http.put("/users", { name, email })
      return handleSuccessResponse(response)
    } catch (error) {
      handleErrorResponse(error)
    }
  },

  login: async (email, password) => {
    try {
      const response = await Vue.http.post("/tokens", { email, password })
      return handleSuccessResponse(response)
    } catch (error) {
      handleErrorResponse(error)
    }
  }
}

function handleSuccessResponse(response) {
  const { accessToken, refreshToken } = response.data.data
  const claims = (({ id, name, email }) => ({ id, name, email }))(decodeJwt(accessToken))
  return { ...claims, accessToken, refreshToken }
}

export default UserService
