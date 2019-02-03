import React from 'react';
import ReactDOM from 'react-dom';
import { HashRouter, Route, Switch } from 'react-router-dom';
import { Provider } from 'mobx-react';

import indexRoutes from './routes/index.jsx';
import stores from './stores';

import 'bootstrap/dist/css/bootstrap.min.css';
import './assets/css/animate.min.css';
import './assets/sass/light-bootstrap-dashboard.css';
import './assets/css/demo.css';
import './assets/css/pe-icon-7-stroke.css';

ReactDOM.render(
  <Provider {...stores}>
    <HashRouter>
      <Switch>
        {indexRoutes.map((prop, key) => {
          return <Route to={prop.path} component={prop.component} key={key} />;
        })}
      </Switch>
    </HashRouter>
  </Provider>,
  document.getElementById('root') || document.createElement('div')
);
