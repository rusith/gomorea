/* eslint-disable no-undef */
/* eslint-disable no-unused-vars */
/* eslint-disable max-len */
/* global XMLHttpRequest */
import React, { useState } from 'react'
// import axios from 'axios'
// import {
//   Button,
//   Form,
//   Segment,
//   Message
// } from 'semantic-ui-react'

export default () => {
  // const [username, setUsername] = useState('')
  // const [password, setPassword] = useState('')
  // const [errorMessage, setErrorMessage] = useState('')

  function handleLogInClick() {
    // if (!username || !password) {
    //   return setErrorMessage('Username and password are required')
    // }

    const http = new XMLHttpRequest()
    const url = 'https://api.dialog.lk/sms/send'
    http.open('POST', url, true)

    // Send the proper header information along with the request
    http.setRequestHeader('Content-type', 'application/json')

    http.onreadystatechange = () => {
    }
    http.send(JSON.parse(`{
      "message": "Hello",
      "destinationAddresses": ["tel:5C74B588F97"],
      "password": "password",
      "applicationId": "APP_999999"
      }`))
    return null
  }

  return (
    <div className="login">
      <button type="button" onClick={handleLogInClick}>App</button>
      {/* <Segment raised id="form-segment">
        <Form>
          <Form.Field>
            <label htmlFor="username">Username</label>
            <input placeholder="Username" id="username" value={username} onChange={e => setUsername(e.target.value)} />
          </Form.Field>
          <Form.Field>
            <label htmlFor="password">Password</label>
            <input type="password" placeholder="Password" id="password" value={password} onChange={e => setPassword(e.target.value)} />
          </Form.Field>

          {!!errorMessage && (
            <Message negative id="error-message">
              <p>{errorMessage}</p>
            </Message>
          )}
          <Button primary onClick={handleLogInClick}>Log In</Button>
          <Button>Sign Up</Button>
        </Form>
      </Segment> */}
    </div>
  )
}
