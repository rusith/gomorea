import React from 'react'
import { BrowserRouter, Route, Switch } from 'react-router-dom'
import './App.css'
import Home from './routes/Home'
import LogIn from './routes/Login'

export default () => (
  <div className="App">
    <BrowserRouter>
      <Switch>
        <Route exact path="/" component={Home} />
        <Route path="/login" component={LogIn} />
      </Switch>
    </BrowserRouter>
  </div>
)
