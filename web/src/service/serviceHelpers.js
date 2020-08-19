import _ from "lodash"

export function handleErrorResponse(error) {
  let message = ""
  if (error.response && error.response.data) {
    const { errors } = error.response.data
    message = errors && errors.map(e => e.message).join(", ")
  }
  if (_.isEmpty(message)) {
    message = "Invalid server response"
  }
  throw message
}
