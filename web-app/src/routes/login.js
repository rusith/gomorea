
import React, { useState } from 'react'
import {
  Button,
  Form,
  Segment
} from 'semantic-ui-react'

export default () => {
  const [username, setUsername] = useState('')
  const [password, setPassword] = useState('')
  return (
    <div className="form">
      <Segment raised>
        <Form>
          <Form.Field>
            <label htmlFor="username">Username</label>
            <input placeholder="Username" id="username" value={username} onChange={e => setUsername(e.target.value)} />
          </Form.Field>
          <Form.Field>
            <label htmlFor="password">Password</label>
            <input placeholder="Password" id="password" value={password} onChange={e => setPassword(e.target.value)} />
          </Form.Field>
          <Button type="submit">Log In</Button>
        </Form>
      </Segment>
    </div>
  )
}
