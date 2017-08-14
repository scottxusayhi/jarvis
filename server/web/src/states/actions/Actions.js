import fetch from 'isomorphic-fetch'
// actions and action creators

// fetch begin
export const FETCH_HOSTS_REQUEST = 'FETCH_HOSTS_REQUEST'
export function fetchHostsRequest(filter) {
    return {
        type: FETCH_HOSTS_REQUEST,
        filter: filter
    }
}

// fetch success
export const FETCH_HOSTS_SUCCESS = 'FETCH_HOSTS_SUCCESS'
export function fetchHostsSuccess(json) {
    return {
        type: FETCH_HOSTS_SUCCESS,
        data: json,
        fetchedAt: Date.now(),
    }
}

// fetch failure
export const FETCH_HOSTS_FAILURE = 'FETCH_HOSTS_FAILURE'
export function fetchHostsFailure(json) {
    return {
        type: FETCH_HOSTS_FAILURE,
        reason: json
    }
}

// refresh fetch
export const INVALIDATE_HOSTS = 'INVALIDATE_HOSTS'

export function invalidateHosts () {
    return {
        type: INVALIDATE_HOSTS,
    }
}

// trunk action creator
// usage: store.dispatch(fetchHosts(filter))
// ses http://redux.js.org/docs/advanced/AsyncActions.html#async-action-creators
export function fetchHosts(filter) {

    return function (dispatch) {
        // helper: check http status
        var checkStatus = response => {
            if (response.status >= 200 && response.status < 300) {
                return response
            } else {
                var error = new Error(response.statusText)
                error.response = response
                throw error
            }
        }

        // helper: parse json
        var parseJson = response => {
            return response.json()
        }

        // api call begin
        dispatch(fetchHostsRequest(filter))
        // api call
        fetch('http://localhost:2999/api/v1/hosts')
            .then(checkStatus)
            .then(parseJson)
            .then(json=>{
                dispatch(fetchHostsSuccess(json))
            })
            .catch(error=>{
              console.error("api error:" + error);
              dispatch(fetchHostsFailure(error))
            })
    }
}


// register begin
export const REGISTER_HOST_REQUEST = 'REGISTER_HOST_REQUEST'
export function registerHostRequest(payload) {
    return {
        type: REGISTER_HOST_REQUEST,
        payload: payload
    }
}

// register success
export const REGISTER_HOST_SUCCESS = 'REGISTER_HOST_SUCCESS'
export function registerHostSuccess(response) {
    return {
        type: REGISTER_HOST_SUCCESS,
        response: response,
    }
}

// register failure
export const REGISTER_HOST_FAILURE = 'REGISTER_HOST_FAILURE'
export function registerHostFailure(reason) {
    return {
        type: REGISTER_HOST_FAILURE,
        reason: reason
    }
}

// trunk action creator
// usage: store.dispatch(fetchHosts(filter))
// ses http://redux.js.org/docs/advanced/AsyncActions.html#async-action-creators
export function registerHost(data) {

    return function (dispatch) {
        // helper: check http status
        var checkStatus = response => {
            if (response.status >= 200 && response.status < 300) {
                return response
            } else {
                var error = new Error(response.statusText)
                error.response = response
                throw error
            }
        }

        // api call begin
        dispatch(registerHostRequest(data))
        // api call
        fetch('http://localhost:2999/api/v1/hosts', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: data
        })
            .then(checkStatus)
            .then(result => {
                return result.json()
            })
            .then(json=>{
                dispatch(registerHostSuccess(json))
            })
            .catch(error=>{
              console.error(error);
              var p = error.response.json()
                p.then(json=> {
                    dispatch(registerHostFailure(json))
                })
            })
    }
}