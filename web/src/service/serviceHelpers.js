import _ from "lodash"

export function handleErrorResponse(error) {
  const { data } = error.response
  let message = _.isEmpty(data) ? error.response.statusText : data.map(e => e.message).join(", ")
  if (_.isEmpty(message)) {
    message = "Invalid server response"
  }
  throw message
}
