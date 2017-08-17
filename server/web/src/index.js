// import 'babel-polyfill'
import React from 'react';
import ReactDOM from 'react-dom';
import { HashRouter, Route, Switch } from 'react-router-dom'
import { createBrowserHistory } from 'history';
// redux
import thunkMiddleware from 'redux-thunk'
import { createLogger } from 'redux-logger'
import { createStore, applyMiddleware, compose } from 'redux'
import { fetchHosts } from './states/actions'
import rootReducer from './states/reducers'
// react-redux
import { Provider } from 'react-redux'

// Containers
import Full from './containers/Full/'

const loggerMiddleware = createLogger()

// const store = createStore (
//     rootReducer,
//     applyMiddleware(
//         thunkMiddleware,
//         loggerMiddleware
//     )
// )

const composeEnhancers = window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose;
const store = createStore(
    rootReducer,
    /* preloadedState, */
    composeEnhancers(applyMiddleware(
        thunkMiddleware,
        loggerMiddleware
    )
    )
)

// store.dispatch(fetchHosts({}))



const history = createBrowserHistory();

ReactDOM.render((
    <Provider store={store}>
      <HashRouter history={history}>
        <Switch>
          <Route path="/" name="Home" component={Full}/>
        </Switch>
      </HashRouter>
    </Provider>
), document.getElementById('root'))
