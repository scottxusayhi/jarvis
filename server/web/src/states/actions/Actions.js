import fetch from 'isomorphic-fetch'
// actions and action creators


var server=window.location.host
console.log(process.env)
if (process.env.NODE_ENV==="development") {
    server="localhost:2999"
}
console.log("set api to "+server)

// fetch registered host list begin
export const FETCH_REG_HOSTS_REQUEST = 'FETCH_REG_HOSTS_REQUEST'
export function fetchRegHostsRequest(filter) {
    return {
        type: FETCH_REG_HOSTS_REQUEST,
        filter: filter
    }
}

// fetch registered host list success
export const FETCH_REG_HOSTS_SUCCESS = 'FETCH_REG_HOSTS_SUCCESS'
export function fetchRegHostsSuccess(json) {
    return {
        type: FETCH_REG_HOSTS_SUCCESS,
        data: json,
        fetchedAt: Date.now(),
    }
}

// fetch registered host list failure
export const FETCH_REG_HOSTS_FAILURE = 'FETCH_REG_HOSTS_FAILURE'
export function fetchRegHostsFailure(json) {
    return {
        type: FETCH_REG_HOSTS_FAILURE,
        reason: json
    }
}

// refresh fetch
export const INVALIDATE_REG_HOSTS = 'INVALIDATE_REG_HOSTS'
export function invalidateRegHosts () {
    return {
        type: INVALIDATE_REG_HOSTS,
    }
}

export function fetchRegisteredHosts(filter) {

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
        dispatch(fetchRegHostsRequest(filter))
        // api call
        fetch('http://'+server+'/api/v1/hosts?'+toQueryString(filter))
            .then(checkStatus)
            .then(parseJson)
            .then(json=>{
                dispatch(fetchRegHostsSuccess(json))
            })
            .catch(error=>{
              console.error("api error:" + error);
              dispatch(fetchRegHostsFailure(error))
            })
    }
}


///////////////////////////////
//
///////////////////////////

// fetch host list begin
export const FETCH_HOSTS_REQUEST = 'FETCH_HOSTS_REQUEST'
export function fetchHostsRequest(filter) {
    return {
        type: FETCH_HOSTS_REQUEST,
        filter: filter
    }
}

// fetch host list success
export const FETCH_HOSTS_SUCCESS = 'FETCH_HOSTS_SUCCESS'
export function fetchHostsSuccess(json) {
    return {
        type: FETCH_HOSTS_SUCCESS,
        data: json,
        fetchedAt: Date.now(),
    }
}

// fetch host list failure
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

function toQueryString(obj) {
    var parts = [];
    for (var i in obj) {
        if (obj.hasOwnProperty(i)) {
            parts.push(encodeURIComponent(i) + "=" + encodeURIComponent(obj[i]));
        }
    }
    return parts.join("&");
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
        fetch('http://'+server+'/api/v1/hosts?'+toQueryString(filter))
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

///////////////
// fetch one host ( for host detail)
///////////////

export const FETCH_HOST_DETAIL_REQUEST = 'FETCH_HOST_DETAIL_REQUEST'
export function fetchHostDetailRequest(id) {
    return {
        type: FETCH_HOST_DETAIL_REQUEST,
        id: id
    }
}

export const FETCH_HOST_DETAIL_SUCCESS = 'FETCH_HOST_DETAIL_SUCCESS'
export function fetchHostDetailSuccess(json) {
    return {
        type: FETCH_HOST_DETAIL_SUCCESS,
        data: json
    }
}

export const FETCH_HOST_DETAIL_FAILURE = 'FETCH_HOST_DETAIL_FAILURE'
export function fetchHostDetailFailure(json) {
    return {
        type: FETCH_HOST_DETAIL_FAILURE,
        error: json
    }
}

export function fetchHostDetail(id) {

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

        // helper: check list length
        function extractFirst(json) {
            if (json.list.length>0) {
                return json.list[0]
            } else {
                var error = new Error("host with systemId " + id + " is not found" )
                throw error
            }
        }

        // api call begin
        dispatch(fetchHostDetailRequest(id))
        var filter = {
            systemId: id
        }
        // api call
        fetch('http://'+server+'/api/v1/hosts?'+toQueryString(filter))
            .then(checkStatus)
            .then(parseJson)
            .then(extractFirst)
            .then(json=>{
                dispatch(fetchHostDetailSuccess(json))
            })
            .catch(error=>{
              console.error("api error:" + error);
              dispatch(fetchHostDetailFailure(error))
            })
    }
}


// new-reg start
export const NEW_REG_START = "NEW_REG_START"
export function newRegStart() {
    return {
        type: NEW_REG_START,
    }
}
// new-reg data saved
export const NEW_REG_DATA_SAVED = "NEW_REG_DATA_SAVED"
export function newRegDataSaved(data) {
    return {
        type: POST_REG_DATA_SAVED,
        data: data
    }
}

// new reg api call begin
export const NEW_REG_REQUEST = 'NEW_REG_REQUEST'
export function newRegRequest(payload) {
    return {
        type: NEW_REG_REQUEST,
        payload: payload
    }
}

// new reg api call success
export const NEW_REG_SUCCESS = 'NEW_REG_SUCCESS'
export function newRegSuccess(response) {
    return {
        type: NEW_REG_SUCCESS,
        response: response,
    }
}

// new reg api call failure
export const NEW_REG_FAILURE = 'NEW_REG_FAILURE'
export function newRegFailure(reason) {
    return {
        type: NEW_REG_FAILURE,
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
        dispatch(newRegRequest(data))
        // api call
        fetch('http://'+server+'/api/v1/hosts', {
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
                dispatch(newRegSuccess(json))
            })
            .catch(error=>{
              console.error(error);
              var p = error.response.json()
                p.then(json=> {
                    dispatch(newRegFailure(json))
                })
            })
    }
}


// post-reg start
export const POST_REG_START = "POST_REG_START"
export function postRegStart(id, initData) {
    return {
        type: POST_REG_START,
        id: id,
        initData: initData
    }
}
// post-reg data saved
export const POST_REG_DATA_SAVED = "POST_REG_DATA_SAVED"
export function postRegDataSaved(data) {
    return {
        type: POST_REG_DATA_SAVED,
        data: data
    }
}

// post reg api call begin
export const POST_REG_REQUEST = 'POST_REG_REQUEST'
export function postRegRequest(payload) {
    return {
        type: POST_REG_REQUEST,
        payload: payload
    }
}

// post reg api call success
export const POST_REG_SUCCESS = 'POST_REG_SUCCESS'
export function postRegSuccess(response) {
    return {
        type: POST_REG_SUCCESS,
        response: response,
    }
}

// post reg api call failure
export const POST_REG_FAILURE = 'POST_REG_FAILURE'
export function postRegFailure(reason) {
    return {
        type: POST_REG_FAILURE,
        reason: reason
    }
}

// trunk action creator
// usage: store.dispatch(fetchHosts(filter))
// ses http://redux.js.org/docs/advanced/AsyncActions.html#async-action-creators
export function postRegHost(id, data) {

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
        dispatch(postRegRequest(data))
        // api call
        fetch('http://'+server+'/api/v1/hosts/'+id, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify(data)
        })
            .then(checkStatus)
            .then(result => {
                return result.json()
            })
            .then(json=>{
                dispatch(postRegSuccess(json))
            })
            .catch(error=>{
              console.error(error);
              var p = error.response.json()
                p.then(json=> {
                    dispatch(postRegFailure(json))
                })
            })
    }
}

// reg (any type) cancelled
export const REG_CANCELLED = 'REG_CANCELLED'
export function regCancelled() {
    return {
        type: REG_CANCELLED
    }
}


// update reg triggered, this is used to update host id to update (in the future)
export const UPDATE_REG_TRIGGER = 'UPDATE_REG_TRIGGER'
export function updateRegTrigger(id) {
    return {
        type: UPDATE_REG_TRIGGER,
        id: id
    }
}

// update reg call begin
export const UPDATE_REG_REQUEST = 'UPDATE_REG_REQUEST'
export function updateRegRequest(payload) {
    return {
        type: UPDATE_REG_REQUEST,
        payload: payload
    }
}

// update reg api call success
export const UPDATE_REG_SUCCESS = 'UPDATE_REG_SUCCESS'
export function updateRegSuccess(data) {
    return {
        type: UPDATE_REG_SUCCESS,
        data: data,
    }
}

// update reg api call failure
export const UPDATE_REG_FAILURE = 'UPDATE_REG_FAILURE'
export function updateRegFailure(reason) {
    return {
        type: UPDATE_REG_FAILURE,
        reason: reason
    }
}

// trunk action creator
// usage: store.dispatch(fetchHosts(filter))
// ses http://redux.js.org/docs/advanced/AsyncActions.html#async-action-creators
export function updateRegHost(id, data) {

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
        dispatch(updateRegRequest(data))
        // api call
        fetch('http://'+server+'/api/v1/hosts/'+id, {
          method: 'PUT',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify(data)
        })
            .then(checkStatus)
            .then(result => {
                return result.json()
            })
            .then(json=>{
                dispatch(updateRegSuccess(json))
            })
            .catch(error=>{
              console.error(error);
              // error.response && dispatch(updateRegFailure(error)) || dispatch(updateRegFailure(error.response.json()))
                dispatch(updateRegFailure(error))
              // var p = error.response.json()
              //   p.then(json=> {
              //       dispatch(updateRegFailure(json))
              //   })
            })
    }
}