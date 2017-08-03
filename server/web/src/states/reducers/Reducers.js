import { combineReducers } from 'redux'
import {
    FETCH_HOSTS_REQUEST,
    FETCH_HOSTS_SUCCESS,
    FETCH_HOSTS_FAILURE,
    INVALIDATE_HOSTS,
    REGISTER_HOST_REQUEST,
    REGISTER_HOST_SUCCESS,
    REGISTER_HOST_FAILURE
} from "../actions"

const initialStateHosts = {
    isFetching: false,
    didInvalidate: false,
    lastUpdated: null,
    data: {},
}

function hosts(state = initialStateHosts, action) {
    switch (action.type) {
        case FETCH_HOSTS_REQUEST:
            return Object.assign({}, state, {
                isFetching: true,
                didInvalidate: false,
            })
        case FETCH_HOSTS_SUCCESS:
            return Object.assign({}, state, {
                isFetching: false,
                didInvalidate: false,
                data: action.data,
                lastUpdated: action.fetchedAt
            })
        case FETCH_HOSTS_FAILURE:
            return Object.assign({}, state, {
                isFetching: false,
                didInvalidate: false,
                data: action.reason,
            })
        case INVALIDATE_HOSTS:
            return Object.assign({}, state, {
                didInvalidate: true
            })
        default:
            return state
    }
}

const initialStateNewHost = {
    isPosting: false,
    success: false,
}

function newHost(state=initialStateNewHost, action) {
    switch (action.type) {
        case REGISTER_HOST_REQUEST:
            return Object.assign({}, state, {
                isPosting: true,
                success: false
            })
        case REGISTER_HOST_SUCCESS:
            return Object.assign({}, state, {
                isPosting: false,
                success: true
        })
        case REGISTER_HOST_FAILURE:
            return Object.assign({}, state, {
                isPosting: false,
                success: false
            })
        default:
            return state
    }
}

const rootReducer = combineReducers({
    hosts,
    newHost
})

export default rootReducer