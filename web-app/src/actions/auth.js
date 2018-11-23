import { LOG_IN_STARTED } from '../types/auth'

export const tryToLogIn = () => (dispatch) => {
  dispatch({
    type: LOG_IN_STARTED
  })
}
