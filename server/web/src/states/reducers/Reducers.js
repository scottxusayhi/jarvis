import { combineReducers } from 'redux'
import {
    FETCH_HOSTS_REQUEST,
    FETCH_HOSTS_SUCCESS,
    FETCH_HOSTS_FAILURE,
    INVALIDATE_HOSTS,
    FETCH_HOST_DETAIL_REQUEST,
    FETCH_HOST_DETAIL_SUCCESS,
    FETCH_HOST_DETAIL_FAILURE,
    REGISTER_HOST_REQUEST,
    REGISTER_HOST_SUCCESS,
    REGISTER_HOST_FAILURE,
    POST_REG_START,
    POST_REG_DATA_SAVED,
} from "../actions"

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

const initialHostDetail = {
    id: 0,
    isFetching: false,
    error: {},
    data: {}
}

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
        default:
            return state
    }
}


const initialPostRegHost = {
    isPosting: false,
    success: false,
    id: 0,
    data: {}
}

function postRegHost(state=initialPostRegHost, action) {
    switch (action.type) {
        case POST_REG_START:
            return Object.assign({}, state, {
                id: action.id,
                data: action.initData,
        })
        case POST_REG_DATA_SAVED:
            return Object.assign({}, state.data, action.data)
        default:
            return state
    }
}

const rootReducer = combineReducers({
    hosts,
    newHost,
    hostDetail,
    postRegHost,
})

export default rootReducer