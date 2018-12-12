import axios from 'axios'

const baseUrl = 'http://localhost:8080/'
let token = null

function getToken() {
  if (token) return token
  token = localStorage.getItem('token')
  return token
}

export function post(path, data) {
  return axios.post(`${baseUrl}${path}`, data, {
    headers: {
      Authorization: `Bearer ${getToken()}`
    }
  })
    .then(r => r.data)
    .catch((e) => {
      throw e.response.data
    })
}


export default post
