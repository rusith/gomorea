import React from 'react'
import { BrowserRouter, Route, Switch } from 'react-router-dom'
import Home from './routes/Home'
import LogIn from './routes/Login'
import SignUp from './routes/SignUp'

export default () => (
  <div className="App">
    <BrowserRouter>
      <Switch>
        <Route exact path="/" component={Home} />
        <Route path="/login" component={LogIn} />
        <Route path="/sign-up" component={SignUp} />
      </Switch>
    </BrowserRouter>
  </div>
)
