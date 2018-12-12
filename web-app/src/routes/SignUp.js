import React, { useState, useEffect } from 'react'
import {
  Button,
  Form,
  Segment,
  Header
} from 'semantic-ui-react'
import { post } from '../helpers/req'

export default ({ history }) => {
  const [firstName, setFirstName] = useState('')
  const [lastName, setLastName] = useState('')
  const [username, setUsername] = useState('')
  const [password, setPassword] = useState('')
  const [error, setError] = useState('')

  function validate() {
    const errors = {}
    if (!firstName) errors.firstName = 'Fist name is required'
    if (!lastName) errors.lastName = 'Last name is required'
    if (!username) errors.username = 'Username is required'
    if (!password) errors.password = 'Password is required'
    errors.has = Object.keys(errors).length > 0
    setError(errors)
  }

  useEffect(validate, [firstName, lastName, username, password])

  const logIn = () => history.push('/login')

  function signUp() {
    post('accounts', {
      firstName, lastName, username, password
    }).then(() => history.push('/'))
      .catch(e => setError({ e }))
  }

  return (
    <div className="sign-up">
      <Segment raised id="form-segment">
        <Header as="h2">
          Sign Up
          <Header.Subheader>Lets start by creating an account!</Header.Subheader>
        </Header>
        <Form>
          <Form.Field error={error.firstName}>
            <label htmlFor="firstName">First Name</label>
            <input placeholder="First Name" id="firstName" value={firstName} onChange={e => setFirstName(e.target.value)} />
            {error.firstName && <p className="input-err">{error.firstName}</p>}
          </Form.Field>
          <Form.Field error={error.lastName}>
            <label htmlFor="lastName">Last Name</label>
            <input placeholder="First Name" id="lastName" value={lastName} onChange={e => setLastName(e.target.value)} />
            {error.lastName && <p className="input-err">{error.lastName}</p>}
          </Form.Field>
          <Form.Field error={error.username}>
            <label htmlFor="username">Username</label>
            <input type="email" placeholder="Put an email address as the username" id="username" value={username} onChange={e => setUsername(e.target.value)} />
            {error.username && <p className="input-err">{error.username}</p>}
          </Form.Field>
          <Form.Field error={error.password}>
            <label htmlFor="password">Password</label>
            <input type="password" placeholder="Password" id="password" value={password} onChange={e => setPassword(e.target.value)} />
            {error.password && <p className="input-err">{error.password}</p>}
          </Form.Field>
          <Button primary type="button" disabled={error.has} onClick={signUp}>Sign Up</Button>
          <Button type="button" onClick={logIn}>Login</Button>
        </Form>
      </Segment>
    </div>
  )
}
