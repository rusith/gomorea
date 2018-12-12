import React, { useState } from 'react'
import {
  Button,
  Form,
  Segment,
  Message,
  Header
} from 'semantic-ui-react'
import { post } from '../helpers/req'

export default ({ history }) => {
  const [username, setUsername] = useState('')
  const [password, setPassword] = useState('')
  const [errorMessage, setErrorMessage] = useState('')

  function handleToken(t) {
    setErrorMessage(null)
    localStorage.setItem('token', t)
    history.push('/')
  }

  function handleLogInClick() {
    if (!username || !password) {
      return setErrorMessage('Username and password are required')
    }

    return post('accounts/token', {
      username,
      password
    }).then(handleToken)
      .catch(setErrorMessage)
  }

  const hdlSingUpClk = () => history.push('/sign-up')

  return (
    <div className="login">
      <Segment raised id="form-segment">
        <Header as="h2">
          Log In
          <Header.Subheader>Enter your password and sign in</Header.Subheader>
        </Header>
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
          <Button primary type="button" onClick={handleLogInClick}>Log In</Button>
          <Button type="button" onClick={hdlSingUpClk}>Sign Up</Button>
        </Form>
      </Segment>
    </div>
  )
}
