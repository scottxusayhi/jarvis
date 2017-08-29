import { combineReducers } from 'redux'
import {
    FETCH_REG_HOSTS_REQUEST,
    FETCH_REG_HOSTS_SUCCESS,
    FETCH_REG_HOSTS_FAILURE,
    INVALIDATE_REG_HOSTS,
    FETCH_HOSTS_REQUEST,
    FETCH_HOSTS_SUCCESS,
    FETCH_HOSTS_FAILURE,
    INVALIDATE_HOSTS,
    FETCH_HOST_DETAIL_REQUEST,
    FETCH_HOST_DETAIL_SUCCESS,
    FETCH_HOST_DETAIL_FAILURE,
    NEW_REG_START,
    NEW_REG_DATA_SAVED,
    NEW_REG_REQUEST,
    NEW_REG_SUCCESS,
    NEW_REG_FAILURE,
    POST_REG_START,
    POST_REG_DATA_SAVED,
    POST_REG_REQUEST,
    POST_REG_SUCCESS,
    POST_REG_FAILURE,
    REG_CANCELLED,
    UPDATE_REG_TRIGGER,
    UPDATE_REG_REQUEST,
    UPDATE_REG_SUCCESS,
    UPDATE_REG_FAILURE
} from "../actions"

var merge = require('deepmerge')

const initialStateRegisteredHosts = {
    isFetching: false,
    didInvalidate: false,
    lastUpdated: null,
    data: {
        pageInfo: {
            size: 0,
            totalSize: 0,
            totalPage: 0,
            page: 1,
            perPage: 20
        }
    },
}

function registeredHosts(state = initialStateRegisteredHosts, action) {
    switch (action.type) {
        case FETCH_REG_HOSTS_REQUEST:
            return Object.assign({}, state, {
                isFetching: true,
                didInvalidate: false,
            })
        case FETCH_REG_HOSTS_SUCCESS:
            return Object.assign({}, state, {
                isFetching: false,
                didInvalidate: false,
                data: action.data,
                lastUpdated: action.fetchedAt
            })
        case FETCH_REG_HOSTS_FAILURE:
            return Object.assign({}, state, {
                isFetching: false,
                didInvalidate: false,
                data: action.reason,
            })
        case INVALIDATE_REG_HOSTS:
            return Object.assign({}, state, {
                didInvalidate: true
            })
        default:
            return state
    }
}

const initialStateHosts = {
    isFetching: false,
    didInvalidate: false,
    lastUpdated: null,
    data: {
        pageInfo: {
            size: 0,
            totalSize: 0,
            totalPage: 0,
            page: 1,
            perPage: 20
        }
    },
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
    type: "",
    postRegHostId: 0,
    isPosting: false,
    success: true, // the registration modal is closed when success==true, and vise versa
    error: {},
    newRegData: {},
    postRegData: {}
}

function regHost(state=initialStateNewHost, action) {
    switch (action.type) {
        case NEW_REG_START:
            return Object.assign({}, state, {
                type: "newReg",
                success: false,
            })
        case NEW_REG_DATA_SAVED:
            return Object.assign({}, state, {
                newRegData: action.data
            })
        case NEW_REG_REQUEST:
            return Object.assign({}, state, {
                isPosting: true,
                success: false
            })
        case NEW_REG_SUCCESS:
            return Object.assign({}, state, {
                isPosting: false,
                success: true
        })
        case NEW_REG_FAILURE:
            return Object.assign({}, state, {
                isPosting: false,
                success: false
            })
        case POST_REG_START:
            return Object.assign({}, state, {
                type: "postReg",
                postRegHostId: action.id,
                postRegData: action.initData,
                success: false,
            })
        case POST_REG_DATA_SAVED:
            return merge(state, {
                postRegData: action.data
            })
        case POST_REG_REQUEST:
            return Object.assign({}, state, {
                isPosting: true,
                success: false
            })
        case POST_REG_SUCCESS:
            return Object.assign({}, state, {
                isPosting: false,
                success: true
            })
        case POST_REG_FAILURE:
            return Object.assign({}, state, {
                isPosting: false,
                success: false
            })
        case REG_CANCELLED:
            return Object.assign({}, state, {
                success: true
            })
        default:
            return state
    }
}

const initialHostDetail = {
    id: 0,
    isFetching: false,
    error: {},
    data: {}
}

// for host detail view and inline update
function hostDetail(state=initialHostDetail, action) {
    switch (action.type) {
        case FETCH_HOST_DETAIL_REQUEST:
            return Object.assign({}, state, {
                isFetching: true,
                id: action.id,
            })
        case FETCH_HOST_DETAIL_SUCCESS:
            return Object.assign({}, state, {
                isFetching: false,
                data: action.data,
            })
        case FETCH_HOST_DETAIL_FAILURE:
            return Object.assign({}, state, {
                isFetching: false,
                error: action.error
            })
        case UPDATE_REG_SUCCESS:
            return Object.assign({}, state, {
                data: action.data
            }
        )
        default:
            return state
    }
}

const rootReducer = combineReducers({
    "registeredHosts": registeredHosts,
    "hosts": hosts,
    "regHost": regHost,
    "hostDetail": hostDetail,
})

export default rootReducer