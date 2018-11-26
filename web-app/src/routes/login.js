
import React, { useState } from 'react'
import {
  Button,
  Form,
  Segment,
  Message
} from 'semantic-ui-react'

export default () => {
  const [username, setUsername] = useState('')
  const [password, setPassword] = useState('')
  const [errorMessage, setErrorMessage] = useState('')

  function handleLogInClick() {
    if (!username || !password) {
      return setErrorMessage('Username and password are required')
    }

    return null
  }

  return (
    <div className="login">
      <Segment raised id="form-segment">
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
      </Segment>
    </div>
  )
}
