import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './sites/App';
import Login from './sites/Login';
import Admin from './sites/Admin';
import * as serviceWorker from './serviceWorker';
import { BrowserRouter, Switch, Route} from 'react-router-dom';

ReactDOM.render(
    <BrowserRouter>
        <div>
            <Switch>
                <Route exact path="/" component={App}/>
                <Route exact path="/login" component={Login}/>
                <Route exact path="/admin" component={Admin}/>
            </Switch>
        </div>
    </BrowserRouter>,
  document.getElementById('root')
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
