import { LOG_IN_STARTED } from '../types/auth'

export const tryToLogIn = (username, password) => {
  return dispatch => {
    dispatch({
      type: LOG_IN_STARTED
    })
  }
}
