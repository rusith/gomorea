import {
  LOG_IN_STARTED
} from './../types/auth'

export default (state = {
  login: {
    loading: false
  }
}, action) => {
  switch (action.type) {
    case LOG_IN_STARTED:
      return { ...state,
        login: {
          ...state.login,
          loading: true
        }
      }
    default:
      return state
  }
}
